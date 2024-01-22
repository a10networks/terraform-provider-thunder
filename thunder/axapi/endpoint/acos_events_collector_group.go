

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsCollectorGroup struct {
	Inst struct {

    Facility string `json:"facility" dval:"local0"`
    Format string `json:"format" dval:"syslog"`
    HealthCheck string `json:"health-check"`
    LogDistribution string `json:"log-distribution" dval:"round-robin"`
    LogServerList []AcosEventsCollectorGroupLogServerList `json:"log-server-list"`
    Name string `json:"name"`
    Protocol string `json:"protocol"`
    Rate int `json:"rate" dval:"500"`
    SamplingEnable []AcosEventsCollectorGroupSamplingEnable `json:"sampling-enable"`
    ServerDistributionHash string `json:"server-distribution-hash" dval:"name"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"collector-group"`
}


type AcosEventsCollectorGroupLogServerList struct {
    Name string `json:"name"`
    Port int `json:"port"`
    Uuid string `json:"uuid"`
}


type AcosEventsCollectorGroupSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AcosEventsCollectorGroup) GetId() string{
    return p.Inst.Name
}

func (p *AcosEventsCollectorGroup) getPath() string{
    return "acos-events/collector-group"
}

func (p *AcosEventsCollectorGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroup::Post")
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

func (p *AcosEventsCollectorGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroup::Get")
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
func (p *AcosEventsCollectorGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroup::Put")
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

func (p *AcosEventsCollectorGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsCollectorGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
