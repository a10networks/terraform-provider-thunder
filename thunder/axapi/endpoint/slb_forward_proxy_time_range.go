

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbForwardProxyTimeRange struct {
	Inst struct {

    DailyBegin string `json:"daily-begin"`
    DailyEnd string `json:"daily-end"`
    DaysOfMonthBegin int `json:"days-of-month-begin"`
    DaysOfMonthEnd int `json:"days-of-month-end"`
    MonthsOfYearBegin int `json:"months-of-year-begin"`
    MonthsOfYearEnd int `json:"months-of-year-end"`
    Name string `json:"name"`
    UseUtcTime int `json:"use-utc-time"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    WeekdayFr int `json:"weekday-fr"`
    WeekdayMo int `json:"weekday-mo"`
    WeekdaySa int `json:"weekday-sa"`
    WeekdaySu int `json:"weekday-su"`
    WeekdayTh int `json:"weekday-th"`
    WeekdayTu int `json:"weekday-tu"`
    WeekdayWe int `json:"weekday-we"`

	} `json:"time-range"`
}

func (p *SlbForwardProxyTimeRange) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbForwardProxyTimeRange) getPath() string{
    return "slb/forward-proxy/time-range"
}

func (p *SlbForwardProxyTimeRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyTimeRange::Post")
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

func (p *SlbForwardProxyTimeRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyTimeRange::Get")
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
func (p *SlbForwardProxyTimeRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyTimeRange::Put")
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

func (p *SlbForwardProxyTimeRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyTimeRange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
