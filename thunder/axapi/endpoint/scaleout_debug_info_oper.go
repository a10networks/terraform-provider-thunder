

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugInfoOper struct {
    
    Oper ScaleoutDebugInfoOperOper `json:"oper"`

}
type DataScaleoutDebugInfoOper struct {
    DtScaleoutDebugInfoOper ScaleoutDebugInfoOper `json:"info"`
}


type ScaleoutDebugInfoOperOper struct {
    Scaleout_enabled int `json:"scaleout_enabled"`
    Db_process_running int `json:"db_process_running"`
    G_scaleout int `json:"g_scaleout"`
    Scaleout_current_role int `json:"scaleout_current_role"`
    Traffic_map_update int `json:"traffic_map_update"`
    Explicitly_stop_service int `json:"explicitly_stop_service"`
    Cluster_disc_timer_running int `json:"cluster_disc_timer_running"`
    Pending_scaleout_exit int `json:"pending_scaleout_exit"`
    Perform_tracking_work int `json:"perform_tracking_work"`
    Node_disable_in_prog int `json:"node_disable_in_prog"`
    Node_enable_in_prog int `json:"node_enable_in_prog"`
    Trigger_cluster_exit int `json:"trigger_cluster_exit"`
    Trigger_enable int `json:"trigger_enable"`
    Trigger_disable int `json:"trigger_disable"`
    Cluster_discovery_timeout int `json:"cluster_discovery_timeout"`
    Cluster_discovery_start_timestamp int `json:"cluster_discovery_start_timestamp"`
    Db_operation_max_retry int `json:"db_operation_max_retry"`
    So_device_count int `json:"so_device_count"`
    Active_device_count int `json:"active_device_count"`
}

func (p *ScaleoutDebugInfoOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugInfoOper) getPath() string{
    return "scaleout/debug/info/oper"
}

func (p *ScaleoutDebugInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugInfoOper,error) {
logger.Println("ScaleoutDebugInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugInfoOper
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
