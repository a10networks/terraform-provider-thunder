package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P6_74
type AdminPassword struct {
	Inst struct {
		EncryptedInModule string `json:"encrypted-in-module"`
		PasswordInModule  string `json:"password-in-module"`
		Uuid              string `json:"uuid"`
	} `json:"password"`
}

func (p *AdminPassword) GetId() string {
	return "1"
}

func (p *AdminPassword) getPath(user string) string {
	return "admin/" + user + "/password"
}

func (p *AdminPassword) Post(authToken string, host string, user string, logger *axapi.ThunderLog) error {
	logger.Println("AdminPassword::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(user), payloadBytes, headers, logger)
	return err
}

func (p *AdminPassword) Get(authToken string, host string, user string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("AdminPassword::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(user), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *AdminPassword) Put(authToken string, host string, user string, logger *axapi.ThunderLog) error {
	logger.Println("AdminPassword::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(user), "", payloadBytes, headers, logger)
	return err
}

func (p *AdminPassword) Delete(authToken string, host string, user string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("AdminPassword::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(user), "", nil, headers, logger)
	return err
}
