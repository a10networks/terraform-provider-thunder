

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPlayerIdList struct {
	Inst struct {

    PlayerRecord []SlbPlayerIdListPlayerRecord `json:"player-record"`

	} `json:"player-id-list"`
}


type SlbPlayerIdListPlayerRecord struct {
    PlayerId int `json:"player-id"`
    GameServerIpv4 string `json:"game-server-ipv4"`
    GameServerIpv6 string `json:"game-server-ipv6"`
    GameServerPortV4 int `json:"game-server-port-v4"`
    GameServerPortV6 int `json:"game-server-port-v6"`
}

func (p *SlbPlayerIdList) GetId() string{
    return "1"
}

func (p *SlbPlayerIdList) getPath() string{
    return "slb/player-id-list"
}

func (p *SlbPlayerIdList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdList::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *SlbPlayerIdList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdList::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SlbPlayerIdList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdList::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SlbPlayerIdList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
