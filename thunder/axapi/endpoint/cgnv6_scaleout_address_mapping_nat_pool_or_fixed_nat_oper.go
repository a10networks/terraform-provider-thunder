package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper struct {
	Oper Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper `json:"oper"`
}
type DataCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper struct {
	DtCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper `json:"nat-pool-or-fixed-nat"`
}

type Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOper struct {
	ServiceTemplate string                                                       `json:"service-template"`
	IpList          []Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList `json:"ip-list"`
	NatPool         string                                                       `json:"nat-pool"`
	Index           int                                                          `json:"index"`
}

type Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOperOperIpList struct {
	UserGroup   int    `json:"user-group"`
	ActiveNode  int    `json:"active-node"`
	StandbyNode int    `json:"standby-node"`
	NatIp       string `json:"nat-ip"`
}

func (p *Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper) GetId() string {
	return "1"
}

func (p *Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper) getPath() string {
	return "cgnv6/scaleout/address-mapping/nat-pool-or-fixed-nat/oper"
}

func (p *Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper, error) {
	logger.Println("Cgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataCgnv6ScaleoutAddressMappingNatPoolOrFixedNatOper
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
