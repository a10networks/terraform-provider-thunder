

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkTrunkOper struct {
    
    Oper NetworkTrunkOperOper `json:"oper"`

}
type DataNetworkTrunkOper struct {
    DtNetworkTrunkOper NetworkTrunkOper `json:"trunk"`
}


type NetworkTrunkOperOper struct {
    Trunk []NetworkTrunkOperOperTrunk `json:"trunk"`
}


type NetworkTrunkOperOperTrunk struct {
    Trunk_id int `json:"trunk_id"`
    Member_count int `json:"member_count"`
    Trunk_name string `json:"trunk_name"`
    Trunk_status string `json:"trunk_status"`
    Trunk_type string `json:"trunk_type"`
    Admin_key int `json:"admin_key"`
    TrunkMemberStatus []NetworkTrunkOperOperTrunkTrunkMemberStatus `json:"trunk-member-status"`
    Ports_threshold int `json:"ports_threshold"`
    Timer int `json:"timer"`
    Timer_running string `json:"timer_running"`
    Ports_threshold_block string `json:"ports_threshold_block"`
    Working_lead int `json:"working_lead"`
}


type NetworkTrunkOperOperTrunkTrunkMemberStatus struct {
    Members int `json:"members"`
    Cfg_status string `json:"cfg_status"`
    Oper_status string `json:"oper_status"`
}

func (p *NetworkTrunkOper) GetId() string{
    return "1"
}

func (p *NetworkTrunkOper) getPath() string{
    return "network/trunk/oper"
}

func (p *NetworkTrunkOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkTrunkOper,error) {
logger.Println("NetworkTrunkOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkTrunkOper
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
