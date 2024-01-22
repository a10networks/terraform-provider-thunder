

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwGtp struct {
	Inst struct {

    ApnLogPeriodicity int `json:"apn-log-periodicity"`
    ApnPrefix FwGtpApnPrefix369 `json:"apn-prefix"`
    ApnPrefixList string `json:"apn-prefix-list"`
    EchoTimeout int `json:"echo-timeout" dval:"120"`
    GtpValue string `json:"gtp-value"`
    InsertionMode string `json:"insertion-mode"`
    NeV4LogPeriodicity int `json:"ne-v4-log-periodicity"`
    NeV6LogPeriodicity int `json:"ne-v6-log-periodicity"`
    NetworkElement FwGtpNetworkElement370 `json:"network-element"`
    NetworkElementListV4 string `json:"network-element-list-v4"`
    NetworkElementListV6 string `json:"network-element-list-v6"`
    PathMgmtLogging string `json:"path-mgmt-logging"`
    SamplingEnable []FwGtpSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"gtp"`
}


type FwGtpApnPrefix369 struct {
    Uuid string `json:"uuid"`
}


type FwGtpNetworkElement370 struct {
    Uuid string `json:"uuid"`
}


type FwGtpSamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
}

func (p *FwGtp) GetId() string{
    return "1"
}

func (p *FwGtp) getPath() string{
    return "fw/gtp"
}

func (p *FwGtp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwGtp::Post")
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

func (p *FwGtp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwGtp::Get")
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
func (p *FwGtp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwGtp::Put")
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

func (p *FwGtp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwGtp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
