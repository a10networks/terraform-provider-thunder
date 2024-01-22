

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FlowspecIcmpCode struct {
	Inst struct {

    Code int `json:"code"`
    CodeEnd int `json:"code-end"`
    IcmpCodeAttribute string `json:"icmp-code-attribute"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"icmp-code"`
}

func (p *FlowspecIcmpCode) GetId() string{
    return p.Inst.IcmpCodeAttribute+"+"+strconv.Itoa(p.Inst.Code)
}

func (p *FlowspecIcmpCode) getPath() string{
    return "flowspec/" +p.Inst.Name + "/icmp-code"
}

func (p *FlowspecIcmpCode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpCode::Post")
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

func (p *FlowspecIcmpCode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpCode::Get")
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
func (p *FlowspecIcmpCode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpCode::Put")
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

func (p *FlowspecIcmpCode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecIcmpCode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
