package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbSslJa4Oper struct {
	Oper SlbSslJa4OperOper `json:"oper"`
}
type DataSlbSslJa4Oper struct {
	DtSlbSslJa4Oper SlbSslJa4Oper `json:"ssl-ja4"`
}

type SlbSslJa4OperOper struct {
	Record []SlbSslJa4OperOperRecord `json:"record"`
}

type SlbSslJa4OperOperRecord struct {
	AddrV4 string `json:"addr-v4"`
	AddrV6 string `json:"addr-v6"`
	Amount int    `json:"amount"`
}

func (p *SlbSslJa4Oper) GetId() string {
	return "1"
}

func (p *SlbSslJa4Oper) getPath() string {
	return "slb/ssl-ja4/oper"
}

func (p *SlbSslJa4Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslJa4Oper, error) {
	logger.Println("SlbSslJa4Oper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbSslJa4Oper
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
