

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemBfdStats struct {
    
    Stats SystemBfdStatsStats `json:"stats"`

}
type DataSystemBfdStats struct {
    DtSystemBfdStats SystemBfdStats `json:"bfd"`
}


type SystemBfdStatsStats struct {
    Ip_checksum_error int `json:"ip_checksum_error"`
    Udp_checksum_error int `json:"udp_checksum_error"`
    Session_not_found int `json:"session_not_found"`
    Multihop_mismatch int `json:"multihop_mismatch"`
    Version_mismatch int `json:"version_mismatch"`
    Length_too_small int `json:"length_too_small"`
    Data_is_short int `json:"data_is_short"`
    Invalid_detect_mult int `json:"invalid_detect_mult"`
    Invalid_multipoint int `json:"invalid_multipoint"`
    Invalid_my_disc int `json:"invalid_my_disc"`
    Invalid_ttl int `json:"invalid_ttl"`
    Auth_length_invalid int `json:"auth_length_invalid"`
    Auth_mismatch int `json:"auth_mismatch"`
    Auth_type_mismatch int `json:"auth_type_mismatch"`
    Auth_key_id_mismatch int `json:"auth_key_id_mismatch"`
    Auth_key_mismatch int `json:"auth_key_mismatch"`
    Auth_seqnum_invalid int `json:"auth_seqnum_invalid"`
    Auth_failed int `json:"auth_failed"`
    Local_state_admin_down int `json:"local_state_admin_down"`
    Dest_unreachable int `json:"dest_unreachable"`
    No_ipv6_enable int `json:"no_ipv6_enable"`
    Other_error int `json:"other_error"`
}

func (p *SystemBfdStats) GetId() string{
    return "1"
}

func (p *SystemBfdStats) getPath() string{
    return "system/bfd/stats"
}

func (p *SystemBfdStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemBfdStats,error) {
logger.Println("SystemBfdStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemBfdStats
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
