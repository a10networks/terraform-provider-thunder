

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugClusterStatusOper struct {
    
    Oper ScaleoutDebugClusterStatusOperOper `json:"oper"`

}
type DataScaleoutDebugClusterStatusOper struct {
    DtScaleoutDebugClusterStatusOper ScaleoutDebugClusterStatusOper `json:"status"`
}


type ScaleoutDebugClusterStatusOperOper struct {
    Scaleout_device_id int `json:"scaleout_device_id"`
    Scaleout_priority int `json:"scaleout_priority"`
    Scaleout_current_role int `json:"scaleout_current_role"`
    Scaleout_enabled int `json:"scaleout_enabled"`
    G_scaleout int `json:"g_scaleout"`
    Min_node_num int `json:"min_node_num"`
    Cluster_disc_timer_running int `json:"cluster_disc_timer_running"`
    Explicitly_stop_service int `json:"explicitly_stop_service"`
    Maintain_mode_configured int `json:"maintain_mode_configured"`
    Device_id int `json:"device_id"`
    Priority int `json:"priority"`
    Vcs_device_id int `json:"vcs_device_id"`
    Vcs_priority int `json:"vcs_priority"`
    Follow_vcs int `json:"follow_vcs"`
    Active_device_count int `json:"active_device_count"`
    Active_device_count_list []ScaleoutDebugClusterStatusOperOperActive_device_count_list `json:"active_device_count_list"`
    Scaleout_device_list []ScaleoutDebugClusterStatusOperOperScaleout_device_list `json:"scaleout_device_list"`
    Scaleout_device_group_list []ScaleoutDebugClusterStatusOperOperScaleout_device_group_list `json:"scaleout_device_group_list"`
}


type ScaleoutDebugClusterStatusOperOperActive_device_count_list struct {
    Device_id int `json:"device_id"`
    State int `json:"state"`
}


type ScaleoutDebugClusterStatusOperOperScaleout_device_list struct {
    Device_id int `json:"device_id"`
    State int `json:"state"`
}


type ScaleoutDebugClusterStatusOperOperScaleout_device_group_list struct {
    Index int `json:"index"`
    Members int `json:"members"`
    Members_low64 int `json:"members_low64"`
}

func (p *ScaleoutDebugClusterStatusOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugClusterStatusOper) getPath() string{
    return "scaleout/debug/cluster/status/oper"
}

func (p *ScaleoutDebugClusterStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugClusterStatusOper,error) {
logger.Println("ScaleoutDebugClusterStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugClusterStatusOper
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
