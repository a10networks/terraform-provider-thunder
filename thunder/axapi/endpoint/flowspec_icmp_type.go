

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FlowspecIcmpType struct {
	Inst struct {

    IcmpTypeAttribute string `json:"icmp-type-attribute"`
    Type int `json:"type"`
    TypeEnd int `json:"type-end"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"icmp-type"`
}

func (p *FlowspecIcmpType) GetId() string{
    return p.Inst.IcmpTypeAttribute+"+"+strconv.Itoa(p.Inst.Type)
}

func (p *FlowspecIcmpType) getPath() string{
    return "flowspec/" +p.Inst.Name + "/icmp-type"
}

func (p *FlowspecIcmpType) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpType::Post")
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

func (p *FlowspecIcmpType) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpType::Get")
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
func (p *FlowspecIcmpType) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpType::Put")
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

func (p *FlowspecIcmpType) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpType::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
