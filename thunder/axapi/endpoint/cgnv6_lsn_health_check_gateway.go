

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnHealthCheckGateway struct {
	Inst struct {

    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    Uuid string `json:"uuid"`

	} `json:"health-check-gateway"`
}

func (p *Cgnv6LsnHealthCheckGateway) GetId() string{
    return p.Inst.Ipv4Addr+"+"+p.Inst.Ipv6Addr
}

func (p *Cgnv6LsnHealthCheckGateway) getPath() string{
    return "cgnv6/lsn/health-check-gateway"
}

func (p *Cgnv6LsnHealthCheckGateway) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnHealthCheckGateway::Post")
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

func (p *Cgnv6LsnHealthCheckGateway) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnHealthCheckGateway::Get")
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
func (p *Cgnv6LsnHealthCheckGateway) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnHealthCheckGateway::Put")
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

func (p *Cgnv6LsnHealthCheckGateway) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnHealthCheckGateway::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
