

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FlowspecProtocol struct {
	Inst struct {

    ProtoAttribute string `json:"proto-attribute"`
    ProtoNum int `json:"proto-num"`
    ProtoNumEnd int `json:"proto-num-end"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"protocol"`
}

func (p *FlowspecProtocol) GetId() string{
    return p.Inst.ProtoAttribute+"+"+strconv.Itoa(p.Inst.ProtoNum)
}

func (p *FlowspecProtocol) getPath() string{
    return "flowspec/" +p.Inst.Name + "/protocol"
}

func (p *FlowspecProtocol) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecProtocol::Post")
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

func (p *FlowspecProtocol) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecProtocol::Get")
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
func (p *FlowspecProtocol) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecProtocol::Put")
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

func (p *FlowspecProtocol) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecProtocol::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
