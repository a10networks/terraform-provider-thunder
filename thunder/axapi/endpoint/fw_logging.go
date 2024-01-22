

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwLogging struct {
	Inst struct {

    Gtp FwLoggingGtp371 `json:"gtp"`
    Name string `json:"name"`
    SamplingEnable []FwLoggingSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}


type FwLoggingGtp371 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []FwLoggingGtpSamplingEnable372 `json:"sampling-enable"`
}


type FwLoggingGtpSamplingEnable372 struct {
    Counters1 string `json:"counters1"`
}


type FwLoggingSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *FwLogging) GetId() string{
    return "1"
}

func (p *FwLogging) getPath() string{
    return "fw/logging"
}

func (p *FwLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwLogging::Post")
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

func (p *FwLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwLogging::Get")
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
func (p *FwLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwLogging::Put")
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

func (p *FwLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
