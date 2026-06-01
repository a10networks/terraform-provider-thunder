package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type SlbServerServiceOper struct {
	Label string `json:"label"`

	Oper SlbServerServiceOperOper `json:"oper"`

	PortNumber int `json:"port-number"`

	Protocol string `json:"protocol"`

	Server_name string
}
type DataSlbServerServiceOper struct {
	DtSlbServerServiceOper SlbServerServiceOper `json:"service"`
}

type SlbServerServiceOperOper struct {
	State                     string                                    `json:"state"`
	Curr_conn_rate            int                                       `json:"curr_conn_rate"`
	Conn_rate_unit            string                                    `json:"conn_rate_unit"`
	Slow_start_conn_limit     int                                       `json:"slow_start_conn_limit"`
	Curr_observe_rate         int                                       `json:"curr_observe_rate"`
	Down_grace_period_allowed int                                       `json:"down_grace_period_allowed"`
	Current_time              int                                       `json:"current_time"`
	Down_time_grace_period    int                                       `json:"down_time_grace_period"`
	Diameter_enabled          int                                       `json:"diameter_enabled"`
	Es_resp_time              int                                       `json:"es_resp_time"`
	Inband_hm_reassign_num    int                                       `json:"inband_hm_reassign_num"`
	Disable                   int                                       `json:"disable"`
	HmKey                     int                                       `json:"hm-key"`
	HmIndex                   int                                       `json:"hm-index"`
	Soft_down_time            int                                       `json:"soft_down_time"`
	Aflow_conn_limit          int                                       `json:"aflow_conn_limit"`
	Aflow_queue_size          int                                       `json:"aflow_queue_size"`
	Resv_conn                 int                                       `json:"resv_conn"`
	AutoNatAddrList           []SlbServerServiceOperOperAutoNatAddrList `json:"auto-nat-addr-list"`
	DrsAutoNatList            []SlbServerServiceOperOperDrsAutoNatList  `json:"drs-auto-nat-list"`
	Pool_name                 string                                    `json:"pool_name"`
	NatPoolAddrList           []SlbServerServiceOperOperNatPoolAddrList `json:"nat-pool-addr-list"`
	DrsIpNatList              []SlbServerServiceOperOperDrsIpNatList    `json:"drs-ip-nat-list"`
}

type SlbServerServiceOperOperAutoNatAddrList struct {
	Auto_nat_ip          string `json:"auto_nat_ip"`
	Vrid                 int    `json:"vrid"`
	Ha_group_id          int    `json:"ha_group_id"`
	Ip_rr                int    `json:"ip_rr"`
	Ports_consumed       int    `json:"ports_consumed"`
	Ports_consumed_total int    `json:"ports_consumed_total"`
	Ports_freed_total    int    `json:"ports_freed_total"`
	Alloc_failed         int    `json:"alloc_failed"`
}

type SlbServerServiceOperOperDrsAutoNatList struct {
	Drs_name              string                                                        `json:"drs_name"`
	Drs_port              int                                                           `json:"drs_port"`
	DrsAutoNatAddressList []SlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList `json:"drs-auto-nat-address-list"`
}

type SlbServerServiceOperOperDrsAutoNatListDrsAutoNatAddressList struct {
	Auto_nat_ip          string `json:"auto_nat_ip"`
	Vrid                 int    `json:"vrid"`
	Ha_group_id          int    `json:"ha_group_id"`
	Ip_rr                int    `json:"ip_rr"`
	Ports_consumed       int    `json:"ports_consumed"`
	Ports_consumed_total int    `json:"ports_consumed_total"`
	Ports_freed_total    int    `json:"ports_freed_total"`
	Alloc_failed         int    `json:"alloc_failed"`
}

type SlbServerServiceOperOperNatPoolAddrList struct {
	Nat_ip               string `json:"nat_ip"`
	Ports_consumed       int    `json:"ports_consumed"`
	Ports_consumed_total int    `json:"ports_consumed_total"`
	Ports_freed_total    int    `json:"ports_freed_total"`
	Alloc_failed         int    `json:"alloc_failed"`
}

type SlbServerServiceOperOperDrsIpNatList struct {
	Drs_name        string                                                `json:"drs_name"`
	Drs_port        int                                                   `json:"drs_port"`
	Pool_name       string                                                `json:"pool_name"`
	NatPoolAddrList []SlbServerServiceOperOperDrsIpNatListNatPoolAddrList `json:"nat-pool-addr-list"`
}

type SlbServerServiceOperOperDrsIpNatListNatPoolAddrList struct {
	Nat_ip               string `json:"nat_ip"`
	Ports_consumed       int    `json:"ports_consumed"`
	Ports_consumed_total int    `json:"ports_consumed_total"`
	Ports_freed_total    int    `json:"ports_freed_total"`
	Alloc_failed         int    `json:"alloc_failed"`
}

func (p *SlbServerServiceOper) GetId() string {
	return "1"
}

func (p *SlbServerServiceOper) getPath() string {

	return "slb/server/" + p.Server_name + "/service/" + strconv.Itoa(p.PortNumber) + "+" + p.Protocol + "+" + p.Label + "/oper"
}

func (p *SlbServerServiceOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbServerServiceOper, error) {
	logger.Println("SlbServerServiceOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbServerServiceOper
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
