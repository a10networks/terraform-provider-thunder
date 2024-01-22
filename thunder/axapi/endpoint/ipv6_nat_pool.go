

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NatPool struct {
	Inst struct {

    ChunkSharing int `json:"chunk-sharing"`
    EndAddress string `json:"end-address"`
    Gateway string `json:"gateway"`
    IpRr int `json:"ip-rr"`
    Netmask int `json:"netmask"`
    PoolName string `json:"pool-name"`
    PortOverload int `json:"port-overload"`
    SamplingEnable []Ipv6NatPoolSamplingEnable `json:"sampling-enable"`
    ScaleoutDeviceId int `json:"scaleout-device-id"`
    StartAddress string `json:"start-address"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"pool"`
}


type Ipv6NatPoolSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Ipv6NatPool) GetId() string{
    return url.QueryEscape(p.Inst.PoolName)
}

func (p *Ipv6NatPool) getPath() string{
    return "ipv6/nat/pool"
}

func (p *Ipv6NatPool) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatPool::Post")
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

func (p *Ipv6NatPool) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatPool::Get")
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
func (p *Ipv6NatPool) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatPool::Put")
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

func (p *Ipv6NatPool) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatPool::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
