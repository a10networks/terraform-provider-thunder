

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPlayerIdEpOper struct {
    
    Oper SlbPlayerIdEpOperOper `json:"oper"`

}
type DataSlbPlayerIdEpOper struct {
    DtSlbPlayerIdEpOper SlbPlayerIdEpOper `json:"player-id-ep"`
}


type SlbPlayerIdEpOperOper struct {
    PlayerIdEpList []SlbPlayerIdEpOperOperPlayerIdEpList `json:"player-id-ep-list"`
    TotalPlayers int `json:"total-players"`
    Filter_type string `json:"filter_type"`
    PlayerId int `json:"player-id"`
}


type SlbPlayerIdEpOperOperPlayerIdEpList struct {
    PlayerId int `json:"player-id"`
    GameServerAddress string `json:"game-server-address"`
    GameServerPort int `json:"game-server-port"`
    Age int `json:"age"`
    UserSessionCount int `json:"user-session-count"`
    IdleTime int `json:"idle-time"`
}

func (p *SlbPlayerIdEpOper) GetId() string{
    return "1"
}

func (p *SlbPlayerIdEpOper) getPath() string{
    return "slb/player-id-ep/oper"
}

func (p *SlbPlayerIdEpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPlayerIdEpOper,error) {
logger.Println("SlbPlayerIdEpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPlayerIdEpOper
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
