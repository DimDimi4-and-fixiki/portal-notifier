package entity

import (
	"fmt"
	"github.com/sethvargo/go-password/password"
	"notify/internal/config"
)

type JSONSerializable any

type ReqWithAuth[DataT JSONSerializable] struct {
	Auth AuthData `json:"auth"`
	Data DataT    `json:"data"`
}

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// CreateServiceReq is data for create_service request
type CreateServiceReq struct {
	ServiceName string `json:"service_name"`
	Login       string `json:"login"`
}

func genPass(length int) (string, error) {
	return password.Generate(length, 10, 0, false, true)
}

func (c *CreateServiceReq) ToUser(cfg config.CryptConfig) (*User, error) {
	apiKey, err := genPass(cfg.APIKeyLength)
	if err != nil {
		return nil, fmt.Errorf("error generating password for service: %v", err)
	}
	user := User{
		Info: UserCommonInfo{Name: c.ServiceName, Role: Service},
		Cred: UserCred{Login: c.Login, Password: apiKey},
	}
	return &user, nil
}
