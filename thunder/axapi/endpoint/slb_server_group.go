

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbServerGroup struct {
	Inst struct {

    MemberList []SlbServerGroupMemberList `json:"member-list"`
    Name string `json:"name"`
    SamplingEnable []SlbServerGroupSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"server-group"`
}


type SlbServerGroupMemberList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
}


type SlbServerGroupSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbServerGroup) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbServerGroup) getPath() string{
    return "slb/server-group"
}

func (p *SlbServerGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerGroup::Post")
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

func (p *SlbServerGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerGroup::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *SlbServerGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerGroup::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *SlbServerGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
