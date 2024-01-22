

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPlayerIdGlobalOper struct {
    
    Oper SlbPlayerIdGlobalOperOper `json:"oper"`

}
type DataSlbPlayerIdGlobalOper struct {
    DtSlbPlayerIdGlobalOper SlbPlayerIdGlobalOper `json:"player-id-global"`
}


type SlbPlayerIdGlobalOperOper struct {
    State string `json:"state"`
    Time_to_active int `json:"time_to_active"`
    Table_count int `json:"table_count"`
}

func (p *SlbPlayerIdGlobalOper) GetId() string{
    return "1"
}

func (p *SlbPlayerIdGlobalOper) getPath() string{
    return "slb/player-id-global/oper"
}

func (p *SlbPlayerIdGlobalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPlayerIdGlobalOper,error) {
logger.Println("SlbPlayerIdGlobalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPlayerIdGlobalOper
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
