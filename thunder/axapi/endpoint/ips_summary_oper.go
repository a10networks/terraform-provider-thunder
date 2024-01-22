

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpsSummaryOper struct {
    
    Oper IpsSummaryOperOper `json:"oper"`

}
type DataIpsSummaryOper struct {
    DtIpsSummaryOper IpsSummaryOper `json:"summary"`
}


type IpsSummaryOperOper struct {
    Version string `json:"version"`
    A10SignatureSource string `json:"a10-signature-source"`
    A10SignatureCount int `json:"a10-signature-count"`
    CustomSignatureCount int `json:"custom-signature-count"`
    Protection string `json:"protection"`
}

func (p *IpsSummaryOper) GetId() string{
    return "1"
}

func (p *IpsSummaryOper) getPath() string{
    return "ips/summary/oper"
}

func (p *IpsSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpsSummaryOper,error) {
logger.Println("IpsSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpsSummaryOper
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
