

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type IpList struct {
	Inst struct {

    Ipv4Config []IpListIpv4Config `json:"ipv4-config"`
    Ipv6Config []IpListIpv6Config `json:"ipv6-config"`
    Ipv6PrefixConfig []IpListIpv6PrefixConfig `json:"ipv6-prefix-config"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ip-list"`
}


type IpListIpv4Config struct {
    Ipv4StartAddr string `json:"ipv4-start-addr"`
    Ipv4EndAddr string `json:"ipv4-end-addr"`
}


type IpListIpv6Config struct {
    Ipv6StartAddr string `json:"ipv6-start-addr"`
    Ipv6EndAddr string `json:"ipv6-end-addr"`
}


type IpListIpv6PrefixConfig struct {
    Ipv6AddrPrefix string `json:"ipv6-addr-prefix"`
    Ipv6PrefixTo string `json:"ipv6-prefix-to"`
    Count1 int `json:"count1"`
}

func (p *IpList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *IpList) getPath() string{
    return "ip-list"
}

func (p *IpList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpList::Post")
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

func (p *IpList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpList::Get")
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
func (p *IpList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpList::Put")
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

func (p *IpList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
