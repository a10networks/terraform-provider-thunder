

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkIcmpv6RateLimit struct {
	Inst struct {

    Icmpv6Lockup int `json:"icmpv6-lockup"`
    Icmpv6LockupPeriod int `json:"icmpv6-lockup-period"`
    Icmpv6NormalRateLimit int `json:"icmpv6-normal-rate-limit"`
    Uuid string `json:"uuid"`

	} `json:"icmpv6-rate-limit"`
}

func (p *NetworkIcmpv6RateLimit) GetId() string{
    return "1"
}

func (p *NetworkIcmpv6RateLimit) getPath() string{
    return "network/icmpv6-rate-limit"
}

func (p *NetworkIcmpv6RateLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkIcmpv6RateLimit::Post")
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

func (p *NetworkIcmpv6RateLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkIcmpv6RateLimit::Get")
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
func (p *NetworkIcmpv6RateLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkIcmpv6RateLimit::Put")
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

func (p *NetworkIcmpv6RateLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkIcmpv6RateLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
