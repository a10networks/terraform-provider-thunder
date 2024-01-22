

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayKerberosInstanceOper struct {
    
    Name string `json:"name"`
    Oper AamAuthenticationRelayKerberosInstanceOperOper `json:"oper"`

}
type DataAamAuthenticationRelayKerberosInstanceOper struct {
    DtAamAuthenticationRelayKerberosInstanceOper AamAuthenticationRelayKerberosInstanceOper `json:"instance"`
}


type AamAuthenticationRelayKerberosInstanceOperOper struct {
    TicketCache string `json:"ticket-cache"`
    DefaultPrincipal string `json:"default-principal"`
    ItemList []AamAuthenticationRelayKerberosInstanceOperOperItemList `json:"item-list"`
}


type AamAuthenticationRelayKerberosInstanceOperOperItemList struct {
    ServicePrincipal string `json:"service-principal"`
    ClientPrincipal string `json:"client-principal"`
    StartTime string `json:"start-time"`
    EndTime string `json:"end-time"`
    RenewTime string `json:"renew-time"`
    Flags string `json:"flags"`
}

func (p *AamAuthenticationRelayKerberosInstanceOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayKerberosInstanceOper) getPath() string{
    
    return "aam/authentication/relay/kerberos/instance/"+p.Name+"/oper"
}

func (p *AamAuthenticationRelayKerberosInstanceOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayKerberosInstanceOper,error) {
logger.Println("AamAuthenticationRelayKerberosInstanceOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayKerberosInstanceOper
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
