

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatPool struct {
	Inst struct {

    All int `json:"all"`
    EndAddress string `json:"end-address"`
    ExcludeIp []Cgnv6NatPoolExcludeIp `json:"exclude-ip"`
    Group string `json:"group"`
    MaxUsersPerIp int `json:"max-users-per-ip"`
    Netmask string `json:"netmask"`
    Partition string `json:"partition"`
    PerBatchPortUsageWarningThreshold int `json:"per-batch-port-usage-warning-threshold"`
    PoolName string `json:"pool-name"`
    PortBatchV2Size string `json:"port-batch-v2-size"`
    Shared int `json:"shared"`
    SimultaneousBatchAllocation int `json:"simultaneous-batch-allocation"`
    StartAddress string `json:"start-address"`
    TcpTimeWaitInterval int `json:"tcp-time-wait-interval"`
    UsableNatPorts int `json:"usable-nat-ports"`
    UsableNatPortsEnd int `json:"usable-nat-ports-end"`
    UsableNatPortsStart int `json:"usable-nat-ports-start"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"pool"`
}


type Cgnv6NatPoolExcludeIp struct {
    ExcludeIpStart string `json:"exclude-ip-start"`
    ExcludeIpEnd string `json:"exclude-ip-end"`
}

func (p *Cgnv6NatPool) GetId() string{
    return url.QueryEscape(p.Inst.PoolName)
}

func (p *Cgnv6NatPool) getPath() string{
    return "cgnv6/nat/pool"
}

func (p *Cgnv6NatPool) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPool::Post")
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

func (p *Cgnv6NatPool) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPool::Get")
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
func (p *Cgnv6NatPool) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPool::Put")
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

func (p *Cgnv6NatPool) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPool::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
