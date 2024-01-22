

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTableStats struct {
    
    Stats DdosTableStatsStats `json:"stats"`

}
type DataDdosTableStats struct {
    DtDdosTableStats DdosTableStats `json:"table"`
}


type DdosTableStatsStats struct {
    Dst_learn int `json:"dst_learn"`
    Dst_hit int `json:"dst_hit"`
    Dst_miss int `json:"dst_miss"`
    Dst_entry_aged int `json:"dst_entry_aged"`
    Src_learn int `json:"src_learn"`
    Src_hit int `json:"src_hit"`
    Src_miss int `json:"src_miss"`
    Src_entry_aged int `json:"src_entry_aged"`
    Src_dst_learn int `json:"src_dst_learn"`
    Src_dst_hit int `json:"src_dst_hit"`
    Src_dst_miss int `json:"src_dst_miss"`
    Src_dst_entry_aged int `json:"src_dst_entry_aged"`
    Telem_err_misc int `json:"telem_err_misc"`
    Telem_route_add_rcvd int `json:"telem_route_add_rcvd"`
    Telem_route_del_rcvd int `json:"telem_route_del_rcvd"`
    Telem_entry_created int `json:"telem_entry_created"`
    Telem_entry_cleared int `json:"telem_entry_cleared"`
    Telem_err_telem_entry_pre_exist int `json:"telem_err_telem_entry_pre_exist"`
    Telem_err_conflict_with_static int `json:"telem_err_conflict_with_static"`
    Telem_err_fail_to_create int `json:"telem_err_fail_to_create"`
    Telem_err_fail_to_delete int `json:"telem_err_fail_to_delete"`
    Src_zone_service_learn int `json:"src_zone_service_learn"`
    Src_zone_service_hit int `json:"src_zone_service_hit"`
    Src_zone_service_miss int `json:"src_zone_service_miss"`
    Src_zone_service_entry_aged int `json:"src_zone_service_entry_aged"`
    Dst_white_list int `json:"dst_white_list"`
    Src_white_list int `json:"src_white_list"`
    Src_dst_white_list int `json:"src_dst_white_list"`
    Src_zone_service_white_list int `json:"src_zone_service_white_list"`
    Dst_black_list int `json:"dst_black_list"`
    Src_black_list int `json:"src_black_list"`
    Src_dst_black_list int `json:"src_dst_black_list"`
    Src_zone_service_black_list int `json:"src_zone_service_black_list"`
    Dst_learning_thre_exceed int `json:"dst_learning_thre_exceed"`
    Dst_over_thre_policy_at_learning int `json:"dst_over_thre_policy_at_learning"`
    Src_learning_thre_exceed int `json:"src_learning_thre_exceed"`
    Src_over_thre_policy_at_lookup int `json:"src_over_thre_policy_at_lookup"`
    Src_over_thre_policy_at_learning int `json:"src_over_thre_policy_at_learning"`
    Src_dst_learning_thre_exceed int `json:"src_dst_learning_thre_exceed"`
    Src_dst_over_thre_policy_at_lookup int `json:"src_dst_over_thre_policy_at_lookup"`
    Src_dst_over_thre_policy_at_learning int `json:"src_dst_over_thre_policy_at_learning"`
    Src_zone_service_learning_thre_exceed int `json:"src_zone_service_learning_thre_exceed"`
    Src_zone_service_over_thre_policy_at_lookup int `json:"src_zone_service_over_thre_policy_at_lookup"`
    Src_zone_service_over_thre_policy_at_learning int `json:"src_zone_service_over_thre_policy_at_learning"`
    Entry_oom int `json:"entry_oom"`
    Entry_ext_oom int `json:"entry_ext_oom"`
    Src_dst_classlist_overflow_policy_at_learning int `json:"src_dst_classlist_overflow_policy_at_learning"`
    Src_zone_service_classlist_overflow_policy_at_learning int `json:"src_zone_service_classlist_overflow_policy_at_learning"`
}

func (p *DdosTableStats) GetId() string{
    return "1"
}

func (p *DdosTableStats) getPath() string{
    return "ddos/table/stats"
}

func (p *DdosTableStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosTableStats,error) {
logger.Println("DdosTableStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosTableStats
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
