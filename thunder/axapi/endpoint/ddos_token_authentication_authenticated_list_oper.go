

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTokenAuthenticationAuthenticatedListOper struct {
    
    Oper DdosTokenAuthenticationAuthenticatedListOperOper `json:"oper"`

}
type DataDdosTokenAuthenticationAuthenticatedListOper struct {
    DtDdosTokenAuthenticationAuthenticatedListOper DdosTokenAuthenticationAuthenticatedListOper `json:"authenticated-list"`
}


type DdosTokenAuthenticationAuthenticatedListOperOper struct {
    PlayerList []DdosTokenAuthenticationAuthenticatedListOperOperPlayerList `json:"player-list"`
}


type DdosTokenAuthenticationAuthenticatedListOperOperPlayerList struct {
    SrcIpStr string `json:"src-ip-str"`
    DstIpStr string `json:"dst-ip-str"`
    SrcPort int `json:"src-port"`
    DstPort int `json:"dst-port"`
    Token int `json:"token"`
}

func (p *DdosTokenAuthenticationAuthenticatedListOper) GetId() string{
    return "1"
}

func (p *DdosTokenAuthenticationAuthenticatedListOper) getPath() string{
    return "ddos/token-authentication/authenticated-list/oper"
}

func (p *DdosTokenAuthenticationAuthenticatedListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosTokenAuthenticationAuthenticatedListOper,error) {
logger.Println("DdosTokenAuthenticationAuthenticatedListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosTokenAuthenticationAuthenticatedListOper
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
