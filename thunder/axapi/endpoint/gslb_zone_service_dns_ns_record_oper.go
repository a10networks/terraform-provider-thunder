package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type GslbZoneServiceDnsNsRecordOper struct {
	NsName string `json:"ns-name"`

	Oper GslbZoneServiceDnsNsRecordOperOper `json:"oper"`

	ServicePort string

	ServiceName string

	Zone_name string
}
type DataGslbZoneServiceDnsNsRecordOper struct {
	DtGslbZoneServiceDnsNsRecordOper GslbZoneServiceDnsNsRecordOper `json:"dns-ns-record"`
}

type GslbZoneServiceDnsNsRecordOperOper struct {
	LastServer string `json:"last-server"`
	Hits       int    `json:"hits"`
}

func (p *GslbZoneServiceDnsNsRecordOper) GetId() string {
	return "1"
}

func (p *GslbZoneServiceDnsNsRecordOper) getPath() string {

	return "gslb/zone/" + p.Zone_name + "/service/" + p.ServicePort + "+" + p.ServiceName + "/dns-ns-record/" + p.NsName + "/oper"
}

func (p *GslbZoneServiceDnsNsRecordOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsNsRecordOper, error) {
	logger.Println("GslbZoneServiceDnsNsRecordOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataGslbZoneServiceDnsNsRecordOper
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
