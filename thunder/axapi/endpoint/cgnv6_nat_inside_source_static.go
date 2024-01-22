

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatInsideSourceStatic struct {
	Inst struct {

    NatAddress string `json:"nat-address"`
    Partition string `json:"partition"`
    SrcAddress string `json:"src-address"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"static"`
}

func (p *Cgnv6NatInsideSourceStatic) GetId() string{
    return p.Inst.SrcAddress+"+"+p.Inst.Partition
}

func (p *Cgnv6NatInsideSourceStatic) getPath() string{
    return "cgnv6/nat/inside/source/static"
}

func (p *Cgnv6NatInsideSourceStatic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatInsideSourceStatic::Post")
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

func (p *Cgnv6NatInsideSourceStatic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatInsideSourceStatic::Get")
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
func (p *Cgnv6NatInsideSourceStatic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatInsideSourceStatic::Put")
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

func (p *Cgnv6NatInsideSourceStatic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatInsideSourceStatic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
