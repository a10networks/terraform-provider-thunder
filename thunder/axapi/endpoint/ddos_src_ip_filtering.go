package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type DdosSrcIpFiltering struct {
	Inst struct {
		FilterClassListList []DdosSrcIpFilteringFilterClassListList `json:"filter-class-list-list"`

		Name string `json:"name"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"src-ip-filtering"`
}

type DdosSrcIpFilteringFilterClassListList struct {
	ClassListName string `json:"class-list-name"`
	Action        string `json:"action"`
	Uuid          string `json:"uuid"`
}

func (p *DdosSrcIpFiltering) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *DdosSrcIpFiltering) getPath() string {
	return "ddos/src-ip-filtering"
}

func (p *DdosSrcIpFiltering) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFiltering::Post")
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

func (p *DdosSrcIpFiltering) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFiltering::Get")
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
func (p *DdosSrcIpFiltering) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFiltering::Put")
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

func (p *DdosSrcIpFiltering) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFiltering::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
