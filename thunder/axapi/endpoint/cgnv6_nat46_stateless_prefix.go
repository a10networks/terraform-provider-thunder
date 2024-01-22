

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat46StatelessPrefix struct {
	Inst struct {

    Ipv6Prefix string `json:"ipv6-prefix"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"prefix"`
}

func (p *Cgnv6Nat46StatelessPrefix) GetId() string{
    return "1"
}

func (p *Cgnv6Nat46StatelessPrefix) getPath() string{
    return "cgnv6/nat46-stateless/prefix"
}

func (p *Cgnv6Nat46StatelessPrefix) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat46StatelessPrefix::Post")
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

func (p *Cgnv6Nat46StatelessPrefix) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat46StatelessPrefix::Get")
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
func (p *Cgnv6Nat46StatelessPrefix) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat46StatelessPrefix::Put")
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

func (p *Cgnv6Nat46StatelessPrefix) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat46StatelessPrefix::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
