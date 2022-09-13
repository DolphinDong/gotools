package wechat

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"strconv"
)

type WeChatUser struct {
	Name       string `json:"name"`
	Department []int  `json:"department"`
	Userid     string `json:"userid"`
}

// 获取部门下的成员
//department id
// fetchChild 是否查找子部门
func (c *weChatClient) GetUsersByDepartmentID(departmentId int, fetchChild bool) (users []*WeChatUser, err error) {
	url := "/cgi-bin/user/simplelist"
	fetch_child := "0"
	if fetchChild {
		fetch_child = "1"
	}
	param := map[string]string{
		"department_id": strconv.Itoa(departmentId),
		"fetch_child":   fetch_child,
	}
	response, err := c.Get(url, param, nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	responseStr := string(response)
	if gjson.Get(responseStr, "errcode").Int() != 0 {
		err = errors.Errorf("get department failed:%v", responseStr)
		return
	}
	userStr := gjson.Get(responseStr, "userlist").String()
	users = make([]*WeChatUser, 0)
	err = json.Unmarshal([]byte(userStr), &users)
	if err != nil {
		err = errors.Errorf("get department failed:%+v", errors.WithStack(err))
	}
	return
}

// 通过名字查询用户ID
// departmentId 部门ID
// fetchChild 是否查找子部门
func (c *weChatClient) GetUserIDbyName(userName string, departmentId int, fetchChild bool) (userId string, err error) {

	users, err := c.GetUsersByDepartmentID(departmentId, fetchChild)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	for _, u := range users {
		if u.Name == userName {
			userId = u.Userid
			return
		}
	}
	return
}
