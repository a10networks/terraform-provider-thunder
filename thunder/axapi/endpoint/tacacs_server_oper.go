package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type TacacsServerOper struct {
	Oper TacacsServerOperOper `json:"oper"`
}
type DataTacacsServerOper struct {
	DtTacacsServerOper TacacsServerOper `json:"tacacs-server"`
}

type TacacsServerOperOper struct {
	TacacsServerList []TacacsServerOperOperTacacsServerList `json:"tacacs-server-list"`
}

type TacacsServerOperOperTacacsServerList struct {
	Name              string `json:"name"`
	Port              int    `json:"port"`
	Socket_open       int    `json:"socket_open"`
	Socket_close      int    `json:"socket_close"`
	Socket_aborts     int    `json:"socket_aborts"`
	Socket_errors     int    `json:"socket_errors"`
	Socket_timeout    int    `json:"socket_timeout"`
	Socket_failconn   int    `json:"socket_failconn"`
	Socket_rev        int    `json:"socket_rev"`
	Socket_send       int    `json:"socket_send"`
	Monitor_oper      int    `json:"monitor_oper"`
	Con_fail_attempts int    `json:"con_fail_attempts"`
	Total_fail_conn   int    `json:"total_fail_conn"`
	Total_fail_auth   int    `json:"total_fail_auth"`
	Last_available    int    `json:"last_available"`
}

func (p *TacacsServerOper) GetId() string {
	return "1"
}

func (p *TacacsServerOper) getPath() string {
	return "tacacs-server/oper"
}

func (p *TacacsServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataTacacsServerOper, error) {
	logger.Println("TacacsServerOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataTacacsServerOper
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
