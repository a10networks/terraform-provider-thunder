

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatRangeList struct {
	Inst struct {

    GlobalNetmaskv4 string `json:"global-netmaskv4"`
    GlobalStartIpv4Addr string `json:"global-start-ipv4-addr"`
    LocalNetmaskv4 string `json:"local-netmaskv4"`
    LocalStartIpv4Addr string `json:"local-start-ipv4-addr"`
    Name string `json:"name"`
    Partition string `json:"partition"`
    Uuid string `json:"uuid"`
    V4Count int `json:"v4-count"`
    V4Vrid int `json:"v4-vrid"`

	} `json:"range-list"`
}

func (p *Cgnv6NatRangeList) GetId() string{
    return url.QueryEscape(p.Inst.Name)+"+"+p.Inst.Partition
}

func (p *Cgnv6NatRangeList) getPath() string{
    return "cgnv6/nat/range-list"
}

func (p *Cgnv6NatRangeList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatRangeList::Post")
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

func (p *Cgnv6NatRangeList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatRangeList::Get")
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
func (p *Cgnv6NatRangeList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatRangeList::Put")
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

func (p *Cgnv6NatRangeList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatRangeList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
