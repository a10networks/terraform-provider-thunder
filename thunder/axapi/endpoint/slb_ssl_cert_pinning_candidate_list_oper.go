

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCertPinningCandidateListOper struct {
    
    Oper SlbSslCertPinningCandidateListOperOper `json:"oper"`

}
type DataSlbSslCertPinningCandidateListOper struct {
    DtSlbSslCertPinningCandidateListOper SlbSslCertPinningCandidateListOper `json:"ssl-cert-pinning-candidate-list"`
}


type SlbSslCertPinningCandidateListOperOper struct {
    CandidateList []SlbSslCertPinningCandidateListOperOperCandidateList `json:"candidate-list"`
    MidRequest int `json:"mid-request"`
    ServerName string `json:"server-name"`
    CentralList int `json:"central-list"`
    LocalList int `json:"local-list"`
    Bypassed int `json:"bypassed"`
    AlphabetOrder int `json:"alphabet-order"`
    All int `json:"all"`
    StatsOnly int `json:"stats-only"`
}


type SlbSslCertPinningCandidateListOperOperCandidateList struct {
    Servername string `json:"servername"`
    ConnFailureCount int `json:"conn-failure-count"`
    BypassCount int `json:"bypass-count"`
    ListType string `json:"list-type"`
    Ttl int `json:"ttl"`
}

func (p *SlbSslCertPinningCandidateListOper) GetId() string{
    return "1"
}

func (p *SlbSslCertPinningCandidateListOper) getPath() string{
    return "slb/ssl-cert-pinning-candidate-list/oper"
}

func (p *SlbSslCertPinningCandidateListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCertPinningCandidateListOper,error) {
logger.Println("SlbSslCertPinningCandidateListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCertPinningCandidateListOper
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
