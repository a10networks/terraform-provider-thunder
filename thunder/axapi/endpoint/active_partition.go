package endpoint

//Operational

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type ActivePartition struct {
	Inst struct {
		ChangeTo string `json:"curr_part_name,omitempty"`
		NowAt    string `json:"partition-name,omitempty"`
		Shared   int    `json:"shared,omitempty"`
	} `json:"active-partition"`
}

func (p *ActivePartition) ChangeTo(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ActivePartition::ChangeTo")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := json.Marshal(p)
	if err != nil {
		logger.Println("json.Marshal() failed with error ", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "active-partition", payloadBytes, headers, logger)
	if err != nil {
		logger.Println("ActivePartition::ChangeTo:", err)
		return err
	}

	logger.Println("ActivePartition::ChangeTo, checking current partition")
	_, axapiResp, err2 := axapi.SendGet(host, "active-partition", "", nil, headers, logger)
	if err2 == nil {
		obj := ActivePartition{}
		err = json.Unmarshal(axapiResp, &obj)
		if obj.Inst.NowAt != p.Inst.ChangeTo {
			logger.Println("ActivePartition::ChangeTo, current partition is " + obj.Inst.NowAt)
			err2 = errors.New("Failed. current partition is " + obj.Inst.NowAt)
		}
	}
	return err2
}
