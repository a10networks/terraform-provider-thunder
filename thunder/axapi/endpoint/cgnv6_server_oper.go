

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ServerOper struct {
    
    Name string `json:"name"`
    Oper Cgnv6ServerOperOper `json:"oper"`
    PortList []Cgnv6ServerOperPortList `json:"port-list"`

}
type DataCgnv6ServerOper struct {
    DtCgnv6ServerOper Cgnv6ServerOper `json:"server"`
}


type Cgnv6ServerOperOper struct {
    State string `json:"state"`
    Creation_type string `json:"creation_type"`
    Dns_update_time string `json:"dns_update_time"`
    Server_ttl int `json:"server_ttl"`
    Srv_gateway_arp string `json:"srv_gateway_arp"`
    Is_autocreate int `json:"is_autocreate"`
    Slow_start_conn_limit int `json:"slow_start_conn_limit"`
    Curr_conn_rate int `json:"curr_conn_rate"`
    Conn_rate_unit string `json:"conn_rate_unit"`
    Curr_observe_rate int `json:"curr_observe_rate"`
    Disable int `json:"disable"`
    Weight int `json:"weight"`
    DrsList []Cgnv6ServerOperOperDrsList `json:"drs-list"`
}


type Cgnv6ServerOperOperDrsList struct {
    DrsName string `json:"drs-name"`
    DrsHost string `json:"drs-host"`
    DrsServerIpv6Addr string `json:"drs-server-ipv6-addr"`
    DrsState string `json:"drs-state"`
    DrsCreation_type string `json:"drs-creation_type"`
    DrsDns_update_time string `json:"drs-dns_update_time"`
    DrsServer_ttl int `json:"drs-server_ttl"`
    DrsSrv_gateway_arp string `json:"drs-srv_gateway_arp"`
    DrsIs_autocreate int `json:"drs-is_autocreate"`
    DrsSlow_start_conn_limit int `json:"drs-slow_start_conn_limit"`
    DrsCurr_conn_rate int `json:"drs-curr_conn_rate"`
    DrsConn_rate_unit string `json:"drs-conn_rate_unit"`
    DrsCurr_observe_rate int `json:"drs-curr_observe_rate"`
    DrsDisable int `json:"drs-disable"`
    DrsCurrConn int `json:"drs-curr-conn"`
    DrsCurrReq int `json:"drs-curr-req"`
    DrsTotConn int `json:"drs-tot-conn"`
    DrsTotReq int `json:"drs-tot-req"`
    DrsTotReqSuc int `json:"drs-tot-req-suc"`
    DrsTotFwdBytes int `json:"drs-tot-fwd-bytes"`
    DrsTotFwdPkts int `json:"drs-tot-fwd-pkts"`
    DrsTotRevBytes int `json:"drs-tot-rev-bytes"`
    DrsTotRevPkts int `json:"drs-tot-rev-pkts"`
    DrsPeakConn int `json:"drs-peak-conn"`
    DrsWeight int `json:"drs-weight"`
}


type Cgnv6ServerOperPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Oper Cgnv6ServerOperPortListOper `json:"oper"`
}


type Cgnv6ServerOperPortListOper struct {
    State string `json:"state"`
    Curr_conn_rate int `json:"curr_conn_rate"`
    Conn_rate_unit string `json:"conn_rate_unit"`
    Slow_start_conn_limit int `json:"slow_start_conn_limit"`
    Curr_observe_rate int `json:"curr_observe_rate"`
    Down_grace_period_allowed int `json:"down_grace_period_allowed"`
    Current_time int `json:"current_time"`
    Down_time_grace_period int `json:"down_time_grace_period"`
    Diameter_enabled int `json:"diameter_enabled"`
    Es_resp_time int `json:"es_resp_time"`
    Inband_hm_reassign_num int `json:"inband_hm_reassign_num"`
    Disable int `json:"disable"`
    HmKey int `json:"hm-key"`
    HmIndex int `json:"hm-index"`
    Soft_down_time int `json:"soft_down_time"`
    Aflow_conn_limit int `json:"aflow_conn_limit"`
    Aflow_queue_size int `json:"aflow_queue_size"`
    Resv_conn int `json:"resv_conn"`
    AutoNatAddrList []Cgnv6ServerOperPortListOperAutoNatAddrList `json:"auto-nat-addr-list"`
    DrsAutoNatList []Cgnv6ServerOperPortListOperDrsAutoNatList `json:"drs-auto-nat-list"`
    Pool_name string `json:"pool_name"`
    NatPoolAddrList []Cgnv6ServerOperPortListOperNatPoolAddrList `json:"nat-pool-addr-list"`
    DrsIpNatList []Cgnv6ServerOperPortListOperDrsIpNatList `json:"drs-ip-nat-list"`
}


type Cgnv6ServerOperPortListOperAutoNatAddrList struct {
    Auto_nat_ip string `json:"auto_nat_ip"`
    Vrid int `json:"vrid"`
    Ha_group_id int `json:"ha_group_id"`
    Ip_rr int `json:"ip_rr"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}


type Cgnv6ServerOperPortListOperDrsAutoNatList struct {
    Drs_name string `json:"drs_name"`
    Drs_port int `json:"drs_port"`
    DrsAutoNatAddressList []Cgnv6ServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList `json:"drs-auto-nat-address-list"`
}


type Cgnv6ServerOperPortListOperDrsAutoNatListDrsAutoNatAddressList struct {
    Auto_nat_ip string `json:"auto_nat_ip"`
    Vrid int `json:"vrid"`
    Ha_group_id int `json:"ha_group_id"`
    Ip_rr int `json:"ip_rr"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}


type Cgnv6ServerOperPortListOperNatPoolAddrList struct {
    Nat_ip string `json:"nat_ip"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}


type Cgnv6ServerOperPortListOperDrsIpNatList struct {
    Drs_name string `json:"drs_name"`
    Drs_port int `json:"drs_port"`
    Pool_name string `json:"pool_name"`
    NatPoolAddrList []Cgnv6ServerOperPortListOperDrsIpNatListNatPoolAddrList `json:"nat-pool-addr-list"`
}


type Cgnv6ServerOperPortListOperDrsIpNatListNatPoolAddrList struct {
    Nat_ip string `json:"nat_ip"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}

func (p *Cgnv6ServerOper) GetId() string{
    return "1"
}

func (p *Cgnv6ServerOper) getPath() string{
    
    return "cgnv6/server/"+p.Name+"/oper"
}

func (p *Cgnv6ServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ServerOper,error) {
logger.Println("Cgnv6ServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6ServerOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
