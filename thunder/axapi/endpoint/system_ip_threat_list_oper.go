

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpThreatListOper struct {
    
    Oper SystemIpThreatListOperOper `json:"oper"`

}
type DataSystemIpThreatListOper struct {
    DtSystemIpThreatListOper SystemIpThreatListOper `json:"ip-threat-list"`
}


type SystemIpThreatListOperOper struct {
    EntriesList []SystemIpThreatListOperOperEntriesList `json:"entries-list"`
    V4Address string `json:"v4-address"`
    V4Netmask string `json:"v4-netmask"`
    V6Prefix string `json:"v6-prefix"`
}


type SystemIpThreatListOperOperEntriesList struct {
    Ip string `json:"ip"`
    MatchType string `json:"match-type"`
    InSpe string `json:"in-spe"`
    Age int `json:"age"`
}

func (p *SystemIpThreatListOper) GetId() string{
    return "1"
}

func (p *SystemIpThreatListOper) getPath() string{
    return "system/ip-threat-list/oper"
}

func (p *SystemIpThreatListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIpThreatListOper,error) {
logger.Println("SystemIpThreatListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIpThreatListOper
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
