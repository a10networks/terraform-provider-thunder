package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type FwTcpSynCookieOper struct {
	Oper FwTcpSynCookieOperOper `json:"oper"`
}
type DataFwTcpSynCookieOper struct {
	DtFwTcpSynCookieOper FwTcpSynCookieOper `json:"syn-cookie"`
}

type FwTcpSynCookieOperOper struct {
	Syn_cookie_on int `json:"syn_cookie_on"`
}

func (p *FwTcpSynCookieOper) GetId() string {
	return "1"
}

func (p *FwTcpSynCookieOper) getPath() string {
	return "fw/tcp/syn-cookie/oper"
}

func (p *FwTcpSynCookieOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwTcpSynCookieOper, error) {
	logger.Println("FwTcpSynCookieOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataFwTcpSynCookieOper
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
