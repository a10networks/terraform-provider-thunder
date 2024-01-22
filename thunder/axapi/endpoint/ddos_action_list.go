

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosActionList struct {
	Inst struct {

    Action DdosActionListAction `json:"action"`
    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosActionListZoneTemplate `json:"zone-template"`

	} `json:"action-list"`
}


type DdosActionListAction struct {
    Action string `json:"action"`
    BlacklistSrcValue int `json:"blacklist-src-value"`
    Stateless int `json:"stateless"`
    ScrubPacket int `json:"scrub-packet"`
}


type DdosActionListZoneTemplate struct {
    Logging string `json:"logging"`
    Encap string `json:"encap"`
}

func (p *DdosActionList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosActionList) getPath() string{
    return "ddos/action-list"
}

func (p *DdosActionList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosActionList::Post")
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

func (p *DdosActionList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosActionList::Get")
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
func (p *DdosActionList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosActionList::Put")
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

func (p *DdosActionList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosActionList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
