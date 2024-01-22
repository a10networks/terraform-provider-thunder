

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatInsideSourceStatic struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    NatAddress string `json:"nat-address"`
    SrcAddress string `json:"src-address"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"static"`
}

func (p *IpNatInsideSourceStatic) GetId() string{
    return p.Inst.SrcAddress+"+"+p.Inst.NatAddress
}

func (p *IpNatInsideSourceStatic) getPath() string{
    return "ip/nat/inside/source/static"
}

func (p *IpNatInsideSourceStatic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceStatic::Post")
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

func (p *IpNatInsideSourceStatic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceStatic::Get")
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
func (p *IpNatInsideSourceStatic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceStatic::Put")
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

func (p *IpNatInsideSourceStatic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatInsideSourceStatic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
