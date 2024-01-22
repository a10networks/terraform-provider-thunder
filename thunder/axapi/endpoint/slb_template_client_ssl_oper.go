

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateClientSslOper struct {
    
    Name string `json:"name"`
    Oper SlbTemplateClientSslOperOper `json:"oper"`

}
type DataSlbTemplateClientSslOper struct {
    DtSlbTemplateClientSslOper SlbTemplateClientSslOper `json:"client-ssl"`
}


type SlbTemplateClientSslOperOper struct {
    CertStatusList []SlbTemplateClientSslOperOperCertStatusList `json:"cert-status-list"`
}


type SlbTemplateClientSslOperOperCertStatusList struct {
    CertStatusName string `json:"cert-status-name"`
    CertStatusStatus string `json:"cert-status-status"`
    CertStatusAge int `json:"cert-status-age"`
    CertStatusNextUpdate string `json:"cert-status-next-update"`
    CertStatusResponder string `json:"cert-status-responder"`
}

func (p *SlbTemplateClientSslOper) GetId() string{
    return "1"
}

func (p *SlbTemplateClientSslOper) getPath() string{
    
    return "slb/template/client-ssl/"+p.Name+"/oper"
}

func (p *SlbTemplateClientSslOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplateClientSslOper,error) {
logger.Println("SlbTemplateClientSslOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplateClientSslOper
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
