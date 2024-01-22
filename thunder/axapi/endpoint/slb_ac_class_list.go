

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbAcClassList struct {
	Inst struct {

    AcList []SlbAcClassListAcList `json:"ac-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`

	} `json:"ac-class-list"`
}


type SlbAcClassListAcList struct {
    Action string `json:"action"`
    AcMatchType string `json:"ac-match-type"`
    AcKeyString string `json:"ac-key-string"`
    AcKeyValue string `json:"ac-key-value"`
}

func (p *SlbAcClassList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbAcClassList) getPath() string{
    return "slb/ac-class-list"
}

func (p *SlbAcClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbAcClassList::Post")
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

func (p *SlbAcClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbAcClassList::Get")
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
func (p *SlbAcClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbAcClassList::Put")
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

func (p *SlbAcClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbAcClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
