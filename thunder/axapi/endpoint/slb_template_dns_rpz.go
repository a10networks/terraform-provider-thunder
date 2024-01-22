

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsRpz struct {
	Inst struct {

    Logging SlbTemplateDnsRpzLogging1422 `json:"logging"`
    Name string `json:"name"`
    SeqId int `json:"seq-id"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"rpz"`
}


type SlbTemplateDnsRpzLogging1422 struct {
    Enable int `json:"enable"`
    RpzAction []SlbTemplateDnsRpzLoggingRpzAction1423 `json:"rpz-action"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsRpzLoggingRpzAction1423 struct {
    StrRpzAction string `json:"str-rpz-action"`
}

func (p *SlbTemplateDnsRpz) GetId() string{
    return strconv.Itoa(p.Inst.SeqId)
}

func (p *SlbTemplateDnsRpz) getPath() string{
    return "slb/template/dns/"+p.Inst.Name+"/rpz"
}

func (p *SlbTemplateDnsRpz) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRpz::Post")
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

func (p *SlbTemplateDnsRpz) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRpz::Get")
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
func (p *SlbTemplateDnsRpz) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRpz::Put")
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

func (p *SlbTemplateDnsRpz) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRpz::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
