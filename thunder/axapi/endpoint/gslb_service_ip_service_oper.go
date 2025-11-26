package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type GslbServiceIpServiceOper struct {
	Label string `json:"label"`

	Oper GslbServiceIpServiceOperOper `json:"oper"`

	PortNum int `json:"port-num"`

	PortProto string `json:"port-proto"`

	NodeName string
}
type DataGslbServiceIpServiceOper struct {
	DtGslbServiceIpServiceOper GslbServiceIpServiceOper `json:"service"`
}

type GslbServiceIpServiceOperOper struct {
	ServicePort         int    `json:"service-port"`
	State               string `json:"state"`
	Disabled            int    `json:"disabled"`
	GslbProtocol        int    `json:"gslb-protocol"`
	LocalProtocol       int    `json:"local-protocol"`
	Tcp                 int    `json:"tcp"`
	ManuallyHealthCheck int    `json:"manually-health-check"`
	Use_gslb_state      int    `json:"use_gslb_state"`
	Dynamic             int    `json:"dynamic"`
}

func (p *GslbServiceIpServiceOper) GetId() string {
	return "1"
}

func (p *GslbServiceIpServiceOper) getPath() string {

	return "gslb/service-ip/" + p.NodeName + "/service/" + strconv.Itoa(p.PortNum) + "+" + p.PortProto + "+" + p.Label + "/oper"
}

func (p *GslbServiceIpServiceOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbServiceIpServiceOper, error) {
	logger.Println("GslbServiceIpServiceOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataGslbServiceIpServiceOper
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
