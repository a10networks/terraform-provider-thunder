

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDistributedForwardingSlb struct {
	Inst struct {

    SlbValue string `json:"slb-value" dval:"disable"`
    TcpThreshold int `json:"tcp-threshold" dval:"5"`
    UdpThreshold int `json:"udp-threshold" dval:"5"`
    Uuid string `json:"uuid"`

	} `json:"slb"`
}

func (p *ScaleoutDistributedForwardingSlb) GetId() string{
    return "1"
}

func (p *ScaleoutDistributedForwardingSlb) getPath() string{
    return "scaleout/distributed-forwarding/slb"
}

func (p *ScaleoutDistributedForwardingSlb) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingSlb::Post")
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

func (p *ScaleoutDistributedForwardingSlb) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingSlb::Get")
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
func (p *ScaleoutDistributedForwardingSlb) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingSlb::Put")
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

func (p *ScaleoutDistributedForwardingSlb) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ScaleoutDistributedForwardingSlb::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
