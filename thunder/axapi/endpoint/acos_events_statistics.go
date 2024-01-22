

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsStatistics struct {
	Inst struct {

    SamplingEnable []AcosEventsStatisticsSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"statistics"`
}


type AcosEventsStatisticsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AcosEventsStatistics) GetId() string{
    return "1"
}

func (p *AcosEventsStatistics) getPath() string{
    return "acos-events/statistics"
}

func (p *AcosEventsStatistics) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsStatistics::Post")
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

func (p *AcosEventsStatistics) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsStatistics::Get")
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
func (p *AcosEventsStatistics) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsStatistics::Put")
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

func (p *AcosEventsStatistics) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsStatistics::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
