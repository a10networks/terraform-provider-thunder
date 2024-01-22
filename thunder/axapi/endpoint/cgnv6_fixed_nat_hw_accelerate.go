

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatHwAccelerate struct {
	Inst struct {

    SamplingEnable []Cgnv6FixedNatHwAccelerateSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"hw-accelerate"`
}


type Cgnv6FixedNatHwAccelerateSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6FixedNatHwAccelerate) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatHwAccelerate) getPath() string{
    return "cgnv6/fixed-nat/hw-accelerate"
}

func (p *Cgnv6FixedNatHwAccelerate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatHwAccelerate::Post")
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

func (p *Cgnv6FixedNatHwAccelerate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatHwAccelerate::Get")
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
func (p *Cgnv6FixedNatHwAccelerate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatHwAccelerate::Put")
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

func (p *Cgnv6FixedNatHwAccelerate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatHwAccelerate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
