

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgTftp struct {
	Inst struct {

    SamplingEnable []Cgnv6LsnAlgTftpSamplingEnable `json:"sampling-enable"`
    TftpValue string `json:"tftp-value"`
    Uuid string `json:"uuid"`

	} `json:"tftp"`
}


type Cgnv6LsnAlgTftpSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6LsnAlgTftp) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgTftp) getPath() string{
    return "cgnv6/lsn/alg/tftp"
}

func (p *Cgnv6LsnAlgTftp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgTftp::Post")
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

func (p *Cgnv6LsnAlgTftp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgTftp::Get")
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
func (p *Cgnv6LsnAlgTftp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgTftp::Put")
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

func (p *Cgnv6LsnAlgTftp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgTftp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
