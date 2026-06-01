package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type DdosSrcIpFilteringFilterClassList struct {
	Inst struct {
		Action string `json:"action"`

		ClassListName string `json:"class-list-name"`

		Uuid string `json:"uuid"`

		Src_ip_filtering_name string
	} `json:"filter-class-list"`
}

func (p *DdosSrcIpFilteringFilterClassList) GetId() string {
	return url.QueryEscape(p.Inst.ClassListName)
}

func (p *DdosSrcIpFilteringFilterClassList) getPath() string {
	return "ddos/src-ip-filtering/" + p.Inst.Src_ip_filtering_name + "/filter-class-list"
}

func (p *DdosSrcIpFilteringFilterClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFilteringFilterClassList::Post")
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

func (p *DdosSrcIpFilteringFilterClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFilteringFilterClassList::Get")
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
func (p *DdosSrcIpFilteringFilterClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFilteringFilterClassList::Put")
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

func (p *DdosSrcIpFilteringFilterClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosSrcIpFilteringFilterClassList::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
