package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P6_74
type GslbGroup struct {
	Inst struct {
		AutoMapLearn    int                        `json:"auto-map-learn" dval:"1"`
		AutoMapPrimary  int                        `json:"auto-map-primary" dval:"1"`
		AutoMapSmart    int                        `json:"auto-map-smart" dval:"1"`
		ConfigAnywhere  int                        `json:"config-anywhere"`
		ConfigMerge     int                        `json:"config-merge"`
		ConfigSave      int                        `json:"config-save" dval:"1"`
		DataInterface   int                        `json:"data-interface" dval:"1"`
		DnsDiscover     int                        `json:"dns-discover" dval:"1"`
		Enable          int                        `json:"enable"`
		Learn           int                        `json:"learn" dval:"1"`
		MgmtInterface   int                        `json:"mgmt-interface" dval:"1"`
		Name            string                     `json:"name"`
		PrimaryIpv6List []GslbGroupPrimaryIpv6List `json:"primary-ipv6-list"`
		PrimaryList     []GslbGroupPrimaryList     `json:"primary-list"`
		Priority        int                        `json:"priority" dval:"100"`
		ResolveAs       string                     `json:"resolve-as" dval:"resolve-to-ipv4"`
		Standalone      int                        `json:"standalone"`
		Suffix          string                     `json:"suffix"`
		UserTag         string                     `json:"user-tag"`
		Uuid            string                     `json:"uuid"`
	} `json:"group"`
}

type GslbGroupPrimaryIpv6List struct {
	PrimaryIpv6 string `json:"primary-ipv6"`
}

type GslbGroupPrimaryList struct {
	Primary string `json:"primary"`
}

func (p *GslbGroup) GetId() string {
	return p.Inst.Name
}

func (p *GslbGroup) getPath() string {
	return "gslb/group"
}

func (p *GslbGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbGroup::Post")
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

func (p *GslbGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbGroup::Get")
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

func (p *GslbGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbGroup::Put")
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

func (p *GslbGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbGroup::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
