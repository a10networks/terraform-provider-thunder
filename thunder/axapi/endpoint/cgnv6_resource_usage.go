

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ResourceUsage struct {
	Inst struct {

    FixedNatInsideUserCount int `json:"fixed-nat-inside-user-count"`
    FixedNatIpAddrCount int `json:"fixed-nat-ip-addr-count"`
    LsnNatAddrCount int `json:"lsn-nat-addr-count"`
    RadiusTableSize int `json:"radius-table-size"`
    StatelessEntries Cgnv6ResourceUsageStatelessEntries110 `json:"stateless-entries"`
    Uuid string `json:"uuid"`

	} `json:"resource-usage"`
}


type Cgnv6ResourceUsageStatelessEntries110 struct {
    L4SessionCount int `json:"l4-session-count"`
    Uuid string `json:"uuid"`
}

func (p *Cgnv6ResourceUsage) GetId() string{
    return "1"
}

func (p *Cgnv6ResourceUsage) getPath() string{
    return "cgnv6/resource-usage"
}

func (p *Cgnv6ResourceUsage) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ResourceUsage::Post")
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

func (p *Cgnv6ResourceUsage) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ResourceUsage::Get")
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
func (p *Cgnv6ResourceUsage) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ResourceUsage::Put")
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

func (p *Cgnv6ResourceUsage) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ResourceUsage::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
