

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SessionsSmpTableOper struct {
    
    Oper SessionsSmpTableOperOper `json:"oper"`

}
type DataSessionsSmpTableOper struct {
    DtSessionsSmpTableOper SessionsSmpTableOper `json:"smp-table"`
}


type SessionsSmpTableOperOper struct {
    EntryList []SessionsSmpTableOperOperEntryList `json:"entry-list"`
}


type SessionsSmpTableOperOperEntryList struct {
    Src4 string `json:"src4"`
    Src6 string `json:"src6"`
    Dst4 string `json:"dst4"`
    Dst6 string `json:"dst6"`
    Srcport int `json:"srcport"`
    Dstport int `json:"dstport"`
    Ttl int `json:"ttl"`
    Type string `json:"type"`
    Payload string `json:"payload"`
}

func (p *SessionsSmpTableOper) GetId() string{
    return "1"
}

func (p *SessionsSmpTableOper) getPath() string{
    return "sessions/smp-table/oper"
}

func (p *SessionsSmpTableOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSessionsSmpTableOper,error) {
logger.Println("SessionsSmpTableOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSessionsSmpTableOper
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
