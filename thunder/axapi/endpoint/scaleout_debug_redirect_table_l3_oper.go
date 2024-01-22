

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugRedirectTableL3Oper struct {
    
    Oper ScaleoutDebugRedirectTableL3OperOper `json:"oper"`

}
type DataScaleoutDebugRedirectTableL3Oper struct {
    DtScaleoutDebugRedirectTableL3Oper ScaleoutDebugRedirectTableL3Oper `json:"l3"`
}


type ScaleoutDebugRedirectTableL3OperOper struct {
    Part_name string `json:"part_name"`
    Device_list []ScaleoutDebugRedirectTableL3OperOperDevice_list `json:"device_list"`
}


type ScaleoutDebugRedirectTableL3OperOperDevice_list struct {
    Device_id int `json:"device_id"`
    Red_follow_shared int `json:"red_follow_shared"`
    Red_table_list []ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list `json:"red_table_list"`
}


type ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_list struct {
    Id1 int `json:"id1"`
    Dir string `json:"dir"`
    Dst_ip string `json:"dst_ip"`
    Total_count int `json:"total_count"`
    Mac string `json:"mac"`
    Real_port int `json:"real_port"`
    Count_list []ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list `json:"count_list"`
}


type ScaleoutDebugRedirectTableL3OperOperDevice_listRed_table_listCount_list struct {
    Dst_index int `json:"dst_index"`
    Dst_index_valid int `json:"dst_index_valid"`
}

func (p *ScaleoutDebugRedirectTableL3Oper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugRedirectTableL3Oper) getPath() string{
    return "scaleout/debug/redirect-table/l3/oper"
}

func (p *ScaleoutDebugRedirectTableL3Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugRedirectTableL3Oper,error) {
logger.Println("ScaleoutDebugRedirectTableL3Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugRedirectTableL3Oper
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
