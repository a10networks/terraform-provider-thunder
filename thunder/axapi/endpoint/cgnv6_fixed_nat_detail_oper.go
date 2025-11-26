package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type Cgnv6FixedNatDetailOper struct {
	Oper Cgnv6FixedNatDetailOperOper `json:"oper"`
}
type DataCgnv6FixedNatDetailOper struct {
	DtCgnv6FixedNatDetailOper Cgnv6FixedNatDetailOper `json:"detail"`
}

type Cgnv6FixedNatDetailOperOper struct {
	FixedNatConfigList []Cgnv6FixedNatDetailOperOperFixedNatConfigList `json:"fixed-nat-config-list"`
}

type Cgnv6FixedNatDetailOperOperFixedNatConfigList struct {
	InsideUser string `json:"inside-user"`
	Index      int    `json:"index"`
}

func (p *Cgnv6FixedNatDetailOper) GetId() string {
	return "1"
}

func (p *Cgnv6FixedNatDetailOper) getPath() string {
	return "cgnv6/fixed-nat/detail/oper"
}

func (p *Cgnv6FixedNatDetailOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatDetailOper, error) {
	logger.Println("Cgnv6FixedNatDetailOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataCgnv6FixedNatDetailOper
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
