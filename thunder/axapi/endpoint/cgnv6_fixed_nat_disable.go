

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatDisable struct {
	Inst struct {

    ClearSession int `json:"clear-session"`
    InsideEndV4address string `json:"inside-end-v4address"`
    InsideEndV6address string `json:"inside-end-v6address"`
    InsideStartV4address string `json:"inside-start-v4address"`
    InsideStartV6address string `json:"inside-start-v6address"`
    IpList string `json:"ip-list"`
    Partition string `json:"partition"`
    V4Netmask string `json:"v4-netmask"`
    V6Netmask int `json:"v6-netmask"`

	} `json:"disable"`
}

func (p *Cgnv6FixedNatDisable) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatDisable) getPath() string{
    return "cgnv6/fixed-nat/disable"
}

func (p *Cgnv6FixedNatDisable) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatDisable::Post")
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

func (p *Cgnv6FixedNatDisable) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatDisable::Get")
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
func (p *Cgnv6FixedNatDisable) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatDisable::Put")
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

func (p *Cgnv6FixedNatDisable) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatDisable::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
