package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper struct {
	Oper Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper `json:"oper"`
}
type DataCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper struct {
	DtCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper `json:"inside-user-or-nat-ip"`
}

type Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOperOper struct {
	UserGroup        int    `json:"user-group"`
	ActiveNode       int    `json:"active-node"`
	StandbyNode      int    `json:"standby-node"`
	ServiceTemplate  string `json:"service-template"`
	Application_type string `json:"application_type"`
	Ip               string `json:"ip"`
	Ipv6             string `json:"ipv6"`
	NatIp            string `json:"nat-ip"`
	Application      string `json:"application"`
}

func (p *Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper) GetId() string {
	return "1"
}

func (p *Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper) getPath() string {
	return "cgnv6/scaleout/address-mapping/inside-user-or-nat-ip/oper"
}

func (p *Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper, error) {
	logger.Println("Cgnv6ScaleoutAddressMappingInsideUserOrNatIpOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataCgnv6ScaleoutAddressMappingInsideUserOrNatIpOper
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
