

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwDdosProtectionOper struct {
    
    Oper FwDdosProtectionOperOper `json:"oper"`

}
type DataFwDdosProtectionOper struct {
    DtFwDdosProtectionOper FwDdosProtectionOper `json:"ddos-protection"`
}


type FwDdosProtectionOperOper struct {
    EntriesList []FwDdosProtectionOperOperEntriesList `json:"entries-list"`
    Details int `json:"details"`
    V4Address string `json:"v4-address"`
    V4Netmask string `json:"v4-netmask"`
    V6Prefix string `json:"v6-prefix"`
}


type FwDdosProtectionOperOperEntriesList struct {
    Ip string `json:"ip"`
    Prefix int `json:"prefix"`
    RuleName string `json:"rule-name"`
    Pps int `json:"pps"`
    Expiration int `json:"expiration"`
    Hints string `json:"hints"`
    Hash int `json:"hash"`
    Lid int `json:"lid"`
    Rate int `json:"rate"`
}

func (p *FwDdosProtectionOper) GetId() string{
    return "1"
}

func (p *FwDdosProtectionOper) getPath() string{
    return "fw/ddos-protection/oper"
}

func (p *FwDdosProtectionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwDdosProtectionOper,error) {
logger.Println("FwDdosProtectionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwDdosProtectionOper
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
