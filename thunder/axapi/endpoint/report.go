

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Report struct {
	Inst struct {

    Debug ReportDebug1095 `json:"debug"`

	} `json:"report"`
}


type ReportDebug1095 struct {
    Log int `json:"log"`
    Sflow ReportDebugSflow1096 `json:"sflow"`
}


type ReportDebugSflow1096 struct {
    Parser int `json:"parser"`
    StatsOid int `json:"stats-oid"`
}

func (p *Report) GetId() string{
    return "1"
}

func (p *Report) getPath() string{
    return "report"
}

func (p *Report) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Report::Post")
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

func (p *Report) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Report::Get")
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
func (p *Report) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Report::Put")
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

func (p *Report) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Report::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
