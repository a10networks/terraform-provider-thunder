package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type SlbVirtualServerPortStats53 struct {
	Inst struct {
		PortNumber int `json:"port-number"`

		Protocol string `json:"protocol"`

		Stats SlbVirtualServerPortStats53Stats `json:"stats"`

		Virtual_server_name string
	} `json:"port"`
}

type SlbVirtualServerPortStats53Stats struct {
	Pop3_vport SlbVirtualServerPortStats53StatsPop3_vport `json:"pop3_vport"`
}

type SlbVirtualServerPortStats53StatsPop3_vport struct {
	Num                   int `json:"num"`
	Curr                  int `json:"curr"`
	Total                 int `json:"total"`
	Svrsel_fail           int `json:"svrsel_fail"`
	No_route              int `json:"no_route"`
	Snat_fail             int `json:"snat_fail"`
	Line_too_long         int `json:"line_too_long"`
	Line_mem_freed        int `json:"line_mem_freed"`
	Invalid_start_line    int `json:"invalid_start_line"`
	Stls                  int `json:"stls"`
	Request_dont_care     int `json:"request_dont_care"`
	Unsupported_command   int `json:"unsupported_command"`
	Bad_sequence          int `json:"bad_sequence"`
	Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
	Smp_v6_fail           int `json:"smp_v6_fail"`
	Smp_v4_fail           int `json:"smp_v4_fail"`
	Insert_tuple_fail     int `json:"insert_tuple_fail"`
	Cl_est_err            int `json:"cl_est_err"`
	Ser_connecting_err    int `json:"ser_connecting_err"`
	Server_response_err   int `json:"server_response_err"`
	Cl_request_err        int `json:"cl_request_err"`
	Request               int `json:"request"`
	Control_to_ssl        int `json:"control_to_ssl"`
}

func (p *SlbVirtualServerPortStats53) GetId() string {
	return "1"
}

func (p *SlbVirtualServerPortStats53) getPath() string {
	return "slb/virtual-server/" + p.Inst.Virtual_server_name + "/port/" + strconv.Itoa(p.Inst.PortNumber) + "+" + p.Inst.Protocol + "/stats?pop3_vport=true"
}

func (p *SlbVirtualServerPortStats53) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPortStats53::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *SlbVirtualServerPortStats53) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPortStats53::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *SlbVirtualServerPortStats53) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPortStats53::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *SlbVirtualServerPortStats53) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPortStats53::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
