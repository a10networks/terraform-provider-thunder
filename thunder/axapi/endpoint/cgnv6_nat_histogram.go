

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatHistogram struct {
	Inst struct {

    BinCount int `json:"bin-count" dval:"50"`
    BinSkew int `json:"bin-skew" dval:"75"`
    DataSkew int `json:"data-skew" dval:"25"`
    Uuid string `json:"uuid"`

	} `json:"histogram"`
}

func (p *Cgnv6NatHistogram) GetId() string{
    return "1"
}

func (p *Cgnv6NatHistogram) getPath() string{
    return "cgnv6/nat/histogram"
}

func (p *Cgnv6NatHistogram) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatHistogram::Post")
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

func (p *Cgnv6NatHistogram) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatHistogram::Get")
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
func (p *Cgnv6NatHistogram) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatHistogram::Put")
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

func (p *Cgnv6NatHistogram) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatHistogram::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
