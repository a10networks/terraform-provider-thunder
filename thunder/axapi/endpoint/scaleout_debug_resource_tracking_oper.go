

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugResourceTrackingOper struct {
    
    Oper ScaleoutDebugResourceTrackingOperOper `json:"oper"`

}
type DataScaleoutDebugResourceTrackingOper struct {
    DtScaleoutDebugResourceTrackingOper ScaleoutDebugResourceTrackingOper `json:"resource-tracking"`
}


type ScaleoutDebugResourceTrackingOperOper struct {
    Scaleout_cluster_list []ScaleoutDebugResourceTrackingOperOperScaleout_cluster_list `json:"scaleout_cluster_list"`
    Track_template_list []ScaleoutDebugResourceTrackingOperOperTrack_template_list `json:"track_template_list"`
}


type ScaleoutDebugResourceTrackingOperOperScaleout_cluster_list struct {
    Cluster_id int `json:"cluster_id"`
    Template_count int `json:"template_count"`
    Last_action string `json:"last_action"`
    Current_state string `json:"current_state"`
    So_track_cleanup int `json:"so_track_cleanup"`
    Template_count_list []ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list `json:"template_count_list"`
}


type ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list struct {
    Template_name string `json:"template_name"`
    Template_valid int `json:"template_valid"`
    Threshold_count int `json:"threshold_count"`
    Threshold_count_list []ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list `json:"threshold_count_list"`
}


type ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list struct {
    Thresholds int `json:"thresholds"`
    Action string `json:"action"`
}


type ScaleoutDebugResourceTrackingOperOperTrack_template_list struct {
    Track_name string `json:"track_name"`
    User_index int `json:"user_index"`
    Ref_cnt int `json:"ref_cnt"`
    Current_cost int `json:"current_cost"`
    Old_cost int `json:"old_cost"`
}

func (p *ScaleoutDebugResourceTrackingOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugResourceTrackingOper) getPath() string{
    return "scaleout/debug/resource-tracking/oper"
}

func (p *ScaleoutDebugResourceTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugResourceTrackingOper,error) {
logger.Println("ScaleoutDebugResourceTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugResourceTrackingOper
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
