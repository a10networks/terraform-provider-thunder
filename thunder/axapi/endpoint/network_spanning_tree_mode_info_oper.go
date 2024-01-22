

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkSpanningTreeModeInfoOper struct {
    
    Oper NetworkSpanningTreeModeInfoOperOper `json:"oper"`

}
type DataNetworkSpanningTreeModeInfoOper struct {
    DtNetworkSpanningTreeModeInfoOper NetworkSpanningTreeModeInfoOper `json:"info"`
}


type NetworkSpanningTreeModeInfoOperOper struct {
    Mode string `json:"mode"`
    Packets_input int `json:"packets_input"`
    Packets_output int `json:"packets_output"`
    Instances []NetworkSpanningTreeModeInfoOperOperInstances `json:"instances"`
}


type NetworkSpanningTreeModeInfoOperOperInstances struct {
    Instance_num int `json:"instance_num"`
    Fwd_state_port string `json:"fwd_state_port"`
    Fwd_state_trunk string `json:"fwd_state_trunk"`
    Blk_state_port string `json:"blk_state_port"`
    Blk_state_trunk string `json:"blk_state_trunk"`
    Bridge_id string `json:"bridge_id"`
    Bridge_priority int `json:"bridge_priority"`
    Bridge_ext_priority int `json:"bridge_ext_priority"`
    Des_root string `json:"des_root"`
    Des_root_priority int `json:"des_root_priority"`
    Des_root_ext_priority int `json:"des_root_ext_priority"`
    Regional_root string `json:"regional_root"`
    Regional_root_priority int `json:"regional_root_priority"`
    Regional_root_ext_priority int `json:"regional_root_ext_priority"`
    Root_port string `json:"root_port"`
    Path_cost int `json:"path_cost"`
    Int_path_cost int `json:"int_path_cost"`
    Max_age int `json:"max_age"`
    Int_max_age int `json:"int_max_age"`
    Root_fwd_delay int `json:"root_fwd_delay"`
    Bridge_fwd_delay int `json:"bridge_fwd_delay"`
    Tx_hold_count int `json:"tx_hold_count"`
    Max_hops int `json:"max_hops"`
    Bridge_hello_time int `json:"bridge_hello_time"`
    Age_time int `json:"age_time"`
    Time_since_topo_change int `json:"time_since_topo_change"`
    Topo_change_count int `json:"topo_change_count"`
    Topo_change_port string `json:"topo_change_port"`
    Last_topo_change_port string `json:"last_topo_change_port"`
    Vlans string `json:"vlans"`
}

func (p *NetworkSpanningTreeModeInfoOper) GetId() string{
    return "1"
}

func (p *NetworkSpanningTreeModeInfoOper) getPath() string{
    return "network/spanning-tree/mode/info/oper"
}

func (p *NetworkSpanningTreeModeInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkSpanningTreeModeInfoOper,error) {
logger.Println("NetworkSpanningTreeModeInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkSpanningTreeModeInfoOper
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
