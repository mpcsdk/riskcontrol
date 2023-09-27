package common

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

type UserTokenInfoGeter struct {
	url   string
	c     *resty.Client
	cache *gcache.Cache
}
type respUserInfo struct {
	Status  int       `json:"status"`
	ErrCode int       `json:"errorCode"`
	Msg     string    `json:"msg"`
	Data    *UserInfo `json:"data"`
}
type UserInfo struct {
	Id         int    `json:"id"`
	UserId     string `json:"appPubKey"`
	Email      string `json:"email"`
	LoginType  string `json:"loginType"`
	Address    string `json:"address"`
	KeyHash    string `json:"keyHash"`
	CreateTime int64  `json:"create_time"`
}

func (s *UserTokenInfoGeter) GetUserInfo(ctx context.Context, token string) (*UserInfo, error) {
	resp, err := s.c.R().
		SetQueryParams(map[string]string{
			"token": token,
		}).
		// EnableTrace().
		Get(s.url)
	fmt.Println(resp)
	if err != nil {
		return nil, err
	}
	userInfo := respUserInfo{}
	err = json.Unmarshal(resp.Body(), &userInfo)
	if err != nil {
		g.Log().Error(ctx, "getUserInfo:", err, token, userInfo)
		return nil, err
	}
	return userInfo.Data, nil
}

func NewUserInfoGeter(url string) (*UserTokenInfoGeter, error) {
	c := resty.New()
	s := &UserTokenInfoGeter{
		c:   c,
		url: url,
	}
	_, err := s.GetUserInfo(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBQdWJLZXkiOiIwMjI1YmI1MmU5NTcyMDUwZmZjMGM4MGRjZDBhYTBmNjQyNDFjMDk5ZDAzZjFlYTFjODEzMmZkMzViY2Q3MDBiMWMiLCJpYXQiOjE2OTQ0Mjk5OTEsImV4cCI6MTcyNTk2NTk5MX0.8YaF5spnD1SjI-NNbBCIBj9H5pspXMMkPJrKk23LdnM")
	return s, err
}
