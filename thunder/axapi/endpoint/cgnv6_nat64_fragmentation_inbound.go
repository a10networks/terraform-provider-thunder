

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat64FragmentationInbound struct {
	Inst struct {

    Count1 int `json:"count1" dval:"1"`
    DfSet string `json:"df-set" dval:"send-icmp"`
    FragAction string `json:"frag-action" dval:"ipv6"`
    Uuid string `json:"uuid"`

	} `json:"inbound"`
}

func (p *Cgnv6Nat64FragmentationInbound) GetId() string{
    return "1"
}

func (p *Cgnv6Nat64FragmentationInbound) getPath() string{
    return "cgnv6/nat64/fragmentation/inbound"
}

func (p *Cgnv6Nat64FragmentationInbound) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64FragmentationInbound::Post")
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

func (p *Cgnv6Nat64FragmentationInbound) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64FragmentationInbound::Get")
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
func (p *Cgnv6Nat64FragmentationInbound) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64FragmentationInbound::Put")
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

func (p *Cgnv6Nat64FragmentationInbound) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64FragmentationInbound::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
