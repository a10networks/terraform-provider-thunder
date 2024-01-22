

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDblb struct {
	Inst struct {

    CalcSha1 SlbTemplateDblbCalcSha11415 `json:"calc-sha1"`
    ClassList string `json:"class-list"`
    Name string `json:"name"`
    ServerVersion string `json:"server-version"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dblb"`
}


type SlbTemplateDblbCalcSha11415 struct {
    Sha1Value string `json:"sha1-value"`
}

func (p *SlbTemplateDblb) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDblb) getPath() string{
    return "slb/template/dblb"
}

func (p *SlbTemplateDblb) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDblb::Post")
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

func (p *SlbTemplateDblb) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDblb::Get")
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
func (p *SlbTemplateDblb) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDblb::Put")
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

func (p *SlbTemplateDblb) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDblb::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
