

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbServerPortOper struct {
    
    Oper SlbServerPortOperOper `json:"oper"`
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Name string 

}
type DataSlbServerPortOper struct {
    DtSlbServerPortOper SlbServerPortOper `json:"port"`
}


type SlbServerPortOperOper struct {
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
    AutoNatAddrList []SlbServerPortOperOperAutoNatAddrList `json:"auto-nat-addr-list"`
    DrsAutoNatList []SlbServerPortOperOperDrsAutoNatList `json:"drs-auto-nat-list"`
    Pool_name string `json:"pool_name"`
    NatPoolAddrList []SlbServerPortOperOperNatPoolAddrList `json:"nat-pool-addr-list"`
    DrsIpNatList []SlbServerPortOperOperDrsIpNatList `json:"drs-ip-nat-list"`
}


type SlbServerPortOperOperAutoNatAddrList struct {
    Auto_nat_ip string `json:"auto_nat_ip"`
    Vrid int `json:"vrid"`
    Ha_group_id int `json:"ha_group_id"`
    Ip_rr int `json:"ip_rr"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}


type SlbServerPortOperOperDrsAutoNatList struct {
    Drs_name string `json:"drs_name"`
    Drs_port int `json:"drs_port"`
    DrsAutoNatAddressList []SlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList `json:"drs-auto-nat-address-list"`
}


type SlbServerPortOperOperDrsAutoNatListDrsAutoNatAddressList struct {
    Auto_nat_ip string `json:"auto_nat_ip"`
    Vrid int `json:"vrid"`
    Ha_group_id int `json:"ha_group_id"`
    Ip_rr int `json:"ip_rr"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}


type SlbServerPortOperOperNatPoolAddrList struct {
    Nat_ip string `json:"nat_ip"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}


type SlbServerPortOperOperDrsIpNatList struct {
    Drs_name string `json:"drs_name"`
    Drs_port int `json:"drs_port"`
    Pool_name string `json:"pool_name"`
    NatPoolAddrList []SlbServerPortOperOperDrsIpNatListNatPoolAddrList `json:"nat-pool-addr-list"`
}


type SlbServerPortOperOperDrsIpNatListNatPoolAddrList struct {
    Nat_ip string `json:"nat_ip"`
    Ports_consumed int `json:"ports_consumed"`
    Ports_consumed_total int `json:"ports_consumed_total"`
    Ports_freed_total int `json:"ports_freed_total"`
    Alloc_failed int `json:"alloc_failed"`
}

func (p *SlbServerPortOper) GetId() string{
    return "1"
}

func (p *SlbServerPortOper) getPath() string{
    
    return "slb/server/" +p.Name + "/port/" +strconv.Itoa(p.PortNumber)+"+"+p.Protocol+"/oper"
}

func (p *SlbServerPortOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbServerPortOper,error) {
logger.Println("SlbServerPortOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbServerPortOper
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
