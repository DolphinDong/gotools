package wechat

import (
	thttp "github.com/DolphinDong/gotools/http"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	_ "github.com/tidwall/gjson"
	"strings"
	"time"
)

const (
	getTokenUrl = "/cgi-bin/gettoken"
)

type weChatClient struct {
	Corpid          string `json:"corpid"`
	Corpsecret      string `json:"corpsecret"`
	AccessToken     string `json:"access_token"`
	TokenExpireTime int64  `json:"token_expire_time"`
	HttpClient      *thttp.Client
}

func NewWeCatClient(baseUrl, corpid, corpsecret string) *weChatClient {
	baseUrl = strings.TrimRight(baseUrl, "/")
	httpClient := &thttp.Client{
		BaseUrl: baseUrl,
	}
	return &weChatClient{
		Corpid:     corpid,
		Corpsecret: corpsecret,
		HttpClient: httpClient,
	}
}

func (c *weChatClient) Get(fullPath string, param, header map[string]string) (response []byte, err error) {
	if time.Now().Unix() > c.TokenExpireTime {
		err = c.Auth()
		if err != nil {
			err = errors.WithMessage(err, "auth failed")
			return
		}
	}

}

// 重新获取token
func (c *weChatClient) Auth() error {
	param := map[string]string{
		"corpid":     c.Corpid,
		"corpsecret": c.Corpsecret,
	}

	response, err := c.HttpClient.Get(getTokenUrl, param, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	errCode := gjson.Get(string(response), "errcode").Int()
	if errCode != 0 {
		return errors.Errorf("get access token failed: %v", errors.WithStack(err))
	}

	token := gjson.Get(string(response), "access_token").String()
	expiresIn := gjson.Get(string(response), "expires_in").Int()
	if token == "" || expiresIn == 0 {
		return errors.Errorf("get access token failed: %v", errors.WithStack(err))
	}
	// 为了保险起见减少一分钟
	tokenExpireTime := time.Now().Unix() + expiresIn - 60
	c.AccessToken = token
	c.TokenExpireTime = tokenExpireTime
	return nil
}
