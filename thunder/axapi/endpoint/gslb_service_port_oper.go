package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type GslbServicePortOper struct {
	Oper GslbServicePortOperOper `json:"oper"`
}
type DataGslbServicePortOper struct {
	DtGslbServicePortOper GslbServicePortOper `json:"service-port"`
}

type GslbServicePortOperOper struct {
	ServiceHash     int                                      `json:"service-hash"`
	ServicePortList []GslbServicePortOperOperServicePortList `json:"service-port-list"`
	Label           string                                   `json:"label"`
	PortCount       int                                      `json:"port-count"`
}

type GslbServicePortOperOperServicePortList struct {
	ServicePortName    string `json:"service-port-name"`
	Attributes         string `json:"attributes"`
	State              string `json:"state"`
	ActiveRealServer   int    `json:"active-real-server"`
	CurrentConnections int    `json:"current-connections"`
	ServiceLabel       string `json:"service-label"`
	ServiceHcode       int    `json:"service-hcode"`
}

func (p *GslbServicePortOper) GetId() string {
	return "1"
}

func (p *GslbServicePortOper) getPath() string {
	return "gslb/service-port/oper"
}

func (p *GslbServicePortOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbServicePortOper, error) {
	logger.Println("GslbServicePortOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataGslbServicePortOper
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
