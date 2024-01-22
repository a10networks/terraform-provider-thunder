

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkSpanningTreeModeInterfaceOper struct {
    
    Oper NetworkSpanningTreeModeInterfaceOperOper `json:"oper"`

}
type DataNetworkSpanningTreeModeInterfaceOper struct {
    DtNetworkSpanningTreeModeInterfaceOper NetworkSpanningTreeModeInterfaceOper `json:"interface"`
}


type NetworkSpanningTreeModeInterfaceOperOper struct {
    Ports []NetworkSpanningTreeModeInterfaceOperOperPorts `json:"ports"`
}


type NetworkSpanningTreeModeInterfaceOperOperPorts struct {
    Instances []NetworkSpanningTreeModeInterfaceOperOperPortsInstances `json:"instances"`
    Port_name string `json:"port_name"`
    Bridge string `json:"bridge"`
    Mode int `json:"mode"`
    Enabled string `json:"enabled"`
    Ex_port_cost int `json:"ex_port_cost"`
    Ad_ex_port_cost int `json:"ad_ex_port_cost"`
    Des_root string `json:"des_root"`
    Des_root_priority int `json:"des_root_priority"`
    Des_root_ext_priority int `json:"des_root_ext_priority"`
    Des_ext_cost int `json:"des_ext_cost"`
    Adm_edge_port string `json:"adm_edge_port"`
    Auto_edge_port string `json:"auto_edge_port"`
    Oper_edge_port string `json:"oper_edge_port"`
    Tc_ack string `json:"tc_ack"`
    Oper_p2p string `json:"oper_p2p"`
    Adm_p2p string `json:"adm_p2p"`
    Rest_role string `json:"rest_role"`
    Rest_tcn string `json:"rest_tcn"`
    Port_hello_time int `json:"port_hello_time"`
    Bpdu_guard_port string `json:"bpdu_guard_port"`
    Bpdu_guard_err string `json:"bpdu_guard_err"`
    Net_port string `json:"net_port"`
    Ba_incon string `json:"ba_incon"`
    Tx_bpdu int `json:"tx_bpdu"`
    Tx_tcn int `json:"tx_tcn"`
    Rx_bpdu int `json:"rx_bpdu"`
    Rx_tcn int `json:"rx_tcn"`
    Trns_fwd int `json:"trns_fwd"`
    Trns_blk int `json:"trns_blk"`
    Rcvd_bpdu string `json:"rcvd_bpdu"`
    Rcvd_stp string `json:"rcvd_stp"`
    Rcvd_rstp string `json:"rcvd_rstp"`
    Snd_rstp string `json:"snd_rstp"`
    Rcvd_ack string `json:"rcvd_ack"`
    Rcvd_tcn string `json:"rcvd_tcn"`
}


type NetworkSpanningTreeModeInterfaceOperOperPortsInstances struct {
    Instance_num int `json:"instance_num"`
    Role string `json:"role"`
    Port_id string `json:"port_id"`
    State string `json:"state"`
    In_port_cost int `json:"in_port_cost"`
    Ad_in_port_cost int `json:"ad_in_port_cost"`
    Des_reg_root string `json:"des_reg_root"`
    Des_reg_root_priority int `json:"des_reg_root_priority"`
    Des_reg_root_ext_priority int `json:"des_reg_root_ext_priority"`
    Des_int_cost int `json:"des_int_cost"`
    Des_bridge string `json:"des_bridge"`
    Des_bridge_priority int `json:"des_bridge_priority"`
    Des_bridge_ext_priority int `json:"des_bridge_ext_priority"`
    Des_port string `json:"des_port"`
    Disputed string `json:"disputed"`
}

func (p *NetworkSpanningTreeModeInterfaceOper) GetId() string{
    return "1"
}

func (p *NetworkSpanningTreeModeInterfaceOper) getPath() string{
    return "network/spanning-tree/mode/interface/oper"
}

func (p *NetworkSpanningTreeModeInterfaceOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkSpanningTreeModeInterfaceOper,error) {
logger.Println("NetworkSpanningTreeModeInterfaceOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkSpanningTreeModeInterfaceOper
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
