

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFtpCtl struct {
	Inst struct {

    SamplingEnable []SlbFtpCtlSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"ftp-ctl"`
}


type SlbFtpCtlSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbFtpCtl) GetId() string{
    return "1"
}

func (p *SlbFtpCtl) getPath() string{
    return "slb/ftp-ctl"
}

func (p *SlbFtpCtl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbFtpCtl::Post")
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

func (p *SlbFtpCtl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbFtpCtl::Get")
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
func (p *SlbFtpCtl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbFtpCtl::Put")
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

func (p *SlbFtpCtl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbFtpCtl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
