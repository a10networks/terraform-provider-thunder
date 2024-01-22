

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ReportDebug struct {
	Inst struct {

    Log int `json:"log"`
    Sflow ReportDebugSflow1094 `json:"sflow"`

	} `json:"debug"`
}


type ReportDebugSflow1094 struct {
    Parser int `json:"parser"`
    StatsOid int `json:"stats-oid"`
}

func (p *ReportDebug) GetId() string{
    return "1"
}

func (p *ReportDebug) getPath() string{
    return "report/debug"
}

func (p *ReportDebug) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ReportDebug::Post")
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

func (p *ReportDebug) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ReportDebug::Get")
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
func (p *ReportDebug) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ReportDebug::Put")
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

func (p *ReportDebug) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ReportDebug::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
