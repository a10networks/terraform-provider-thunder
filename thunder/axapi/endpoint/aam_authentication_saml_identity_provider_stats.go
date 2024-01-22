

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSamlIdentityProviderStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationSamlIdentityProviderStatsStats `json:"stats"`

}
type DataAamAuthenticationSamlIdentityProviderStats struct {
    DtAamAuthenticationSamlIdentityProviderStats AamAuthenticationSamlIdentityProviderStats `json:"identity-provider"`
}


type AamAuthenticationSamlIdentityProviderStatsStats struct {
    ValidStatus int `json:"valid-status"`
    MdState int `json:"md-state"`
    MdUpdate int `json:"md-update"`
    MdFail int `json:"md-fail"`
    AcsState int `json:"acs-state"`
    AcsReq int `json:"acs-req"`
    AcsPass int `json:"acs-pass"`
    AcsFail int `json:"acs-fail"`
}

func (p *AamAuthenticationSamlIdentityProviderStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationSamlIdentityProviderStats) getPath() string{
    
    return "aam/authentication/saml/identity-provider/"+p.Name+"/stats"
}

func (p *AamAuthenticationSamlIdentityProviderStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationSamlIdentityProviderStats,error) {
logger.Println("AamAuthenticationSamlIdentityProviderStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationSamlIdentityProviderStats
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
