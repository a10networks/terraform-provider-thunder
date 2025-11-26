package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type GslbZoneDnsCaaRecordOper struct {
	CriticalFlag int `json:"critical-flag"`

	Oper GslbZoneDnsCaaRecordOperOper `json:"oper"`

	PropertyTag string `json:"property-tag"`

	Rdata string `json:"rdata"`

	Zone_name string
}
type DataGslbZoneDnsCaaRecordOper struct {
	DtGslbZoneDnsCaaRecordOper GslbZoneDnsCaaRecordOper `json:"dns-caa-record"`
}

type GslbZoneDnsCaaRecordOperOper struct {
	LastServer string `json:"last-server"`
}

func (p *GslbZoneDnsCaaRecordOper) GetId() string {
	return "1"
}

func (p *GslbZoneDnsCaaRecordOper) getPath() string {

	return "gslb/zone/" + p.Zone_name + "/dns-caa-record/" + strconv.Itoa(p.CriticalFlag) + "+" + p.PropertyTag + "+" + p.Rdata + "/oper"
}

func (p *GslbZoneDnsCaaRecordOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneDnsCaaRecordOper, error) {
	logger.Println("GslbZoneDnsCaaRecordOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataGslbZoneDnsCaaRecordOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
