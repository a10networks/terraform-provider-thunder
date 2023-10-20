package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type ServicePartition struct {
	Inst struct {
		ApplicationType string `json:"application-type"`
		FollowVrid      int    `json:"follow-vrid"`
		Id              int    `json:"id"`
		PartitionName   string `json:"partition-name"`
		UserTag         string `json:"user-tag"`
		Uuid            string `json:"uuid"`
	} `json:"service-partition"`
}

func (p *ServicePartition) GetId() string {
	return p.Inst.PartitionName
}

func (p *ServicePartition) getPath() string {
	return "service-partition"
}

func (p *ServicePartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ServicePartition::Post")
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

func (p *ServicePartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ServicePartition::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *ServicePartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ServicePartition::Put")
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

func (p *ServicePartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ServicePartition::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
