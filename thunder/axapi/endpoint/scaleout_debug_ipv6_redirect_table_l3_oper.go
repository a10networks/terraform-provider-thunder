

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugIpv6RedirectTableL3Oper struct {
    
    Oper ScaleoutDebugIpv6RedirectTableL3OperOper `json:"oper"`

}
type DataScaleoutDebugIpv6RedirectTableL3Oper struct {
    DtScaleoutDebugIpv6RedirectTableL3Oper ScaleoutDebugIpv6RedirectTableL3Oper `json:"l3"`
}


type ScaleoutDebugIpv6RedirectTableL3OperOper struct {
    Part_name string `json:"part_name"`
    Device_list []ScaleoutDebugIpv6RedirectTableL3OperOperDevice_list `json:"device_list"`
}


type ScaleoutDebugIpv6RedirectTableL3OperOperDevice_list struct {
    Device_id int `json:"device_id"`
    Red_follow_shared int `json:"red_follow_shared"`
    Red_table_list []ScaleoutDebugIpv6RedirectTableL3OperOperDevice_listRed_table_list `json:"red_table_list"`
}


type ScaleoutDebugIpv6RedirectTableL3OperOperDevice_listRed_table_list struct {
    Id1 int `json:"id1"`
    Dir string `json:"dir"`
    Dst_ip string `json:"dst_ip"`
    Total_count int `json:"total_count"`
    Mac string `json:"mac"`
    Real_port int `json:"real_port"`
    Count_list []ScaleoutDebugIpv6RedirectTableL3OperOperDevice_listRed_table_listCount_list `json:"count_list"`
}


type ScaleoutDebugIpv6RedirectTableL3OperOperDevice_listRed_table_listCount_list struct {
    Dst_index int `json:"dst_index"`
    Dst_index_valid int `json:"dst_index_valid"`
}

func (p *ScaleoutDebugIpv6RedirectTableL3Oper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugIpv6RedirectTableL3Oper) getPath() string{
    return "scaleout/debug/ipv6/redirect-table/l3/oper"
}

func (p *ScaleoutDebugIpv6RedirectTableL3Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugIpv6RedirectTableL3Oper,error) {
logger.Println("ScaleoutDebugIpv6RedirectTableL3Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugIpv6RedirectTableL3Oper
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
