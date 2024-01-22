

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayFormBasedInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationRelayFormBasedInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationRelayFormBasedInstanceStats struct {
    DtAamAuthenticationRelayFormBasedInstanceStats AamAuthenticationRelayFormBasedInstanceStats `json:"instance"`
}


type AamAuthenticationRelayFormBasedInstanceStatsStats struct {
    Request int `json:"request"`
    Invalid_srv_rsp int `json:"invalid_srv_rsp"`
    Post_fail int `json:"post_fail"`
    Invalid_cred int `json:"invalid_cred"`
    Bad_req int `json:"bad_req"`
    Not_fnd int `json:"not_fnd"`
    Error int `json:"error"`
    Other_error int `json:"other_error"`
}

func (p *AamAuthenticationRelayFormBasedInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayFormBasedInstanceStats) getPath() string{
    
    return "aam/authentication/relay/form-based/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationRelayFormBasedInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayFormBasedInstanceStats,error) {
logger.Println("AamAuthenticationRelayFormBasedInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayFormBasedInstanceStats
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
