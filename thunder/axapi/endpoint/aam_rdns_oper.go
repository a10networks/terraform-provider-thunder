

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamRdnsOper struct {
    
    Oper AamRdnsOperOper `json:"oper"`

}
type DataAamRdnsOper struct {
    DtAamRdnsOper AamRdnsOper `json:"rdns"`
}


type AamRdnsOperOper struct {
    EntryList []AamRdnsOperOperEntryList `json:"entry-list"`
    Mode string `json:"mode"`
    Addr string `json:"addr"`
}


type AamRdnsOperOperEntryList struct {
    Type string `json:"type"`
    Address string `json:"address"`
    Domain string `json:"domain"`
    Ttl int `json:"ttl"`
}

func (p *AamRdnsOper) GetId() string{
    return "1"
}

func (p *AamRdnsOper) getPath() string{
    return "aam/rdns/oper"
}

func (p *AamRdnsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamRdnsOper,error) {
logger.Println("AamRdnsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamRdnsOper
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
