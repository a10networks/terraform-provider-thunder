package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type SlbServerGroupMember struct {
	Inst struct {
		Name string `json:"name"`

		Uuid string `json:"uuid"`

		Server_group_name string
	} `json:"member"`
}

func (p *SlbServerGroupMember) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *SlbServerGroupMember) getPath() string {
	return "slb/server-group/" + p.Inst.Server_group_name + "/member"
}

func (p *SlbServerGroupMember) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerGroupMember::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *SlbServerGroupMember) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerGroupMember::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *SlbServerGroupMember) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerGroupMember::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *SlbServerGroupMember) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerGroupMember::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
