

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SflowSampling struct {
	Inst struct {

    EthList []SflowSamplingEthList `json:"eth-list"`
    Uuid string `json:"uuid"`
    VeList []SflowSamplingVeList `json:"ve-list"`

	} `json:"sampling"`
}


type SflowSamplingEthList struct {
    EthStart int `json:"eth-start"`
    EthEnd int `json:"eth-end"`
}


type SflowSamplingVeList struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *SflowSampling) GetId() string{
    return "1"
}

func (p *SflowSampling) getPath() string{
    return "sflow/sampling"
}

func (p *SflowSampling) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSampling::Post")
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

func (p *SflowSampling) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSampling::Get")
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
func (p *SflowSampling) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSampling::Put")
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

func (p *SflowSampling) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSampling::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
