

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6RipDistributeListPrefix struct {
	Inst struct {

    PrefixCfg []RouterIpv6RipDistributeListPrefixPrefixCfg `json:"prefix-cfg"`
    Uuid string `json:"uuid"`

	} `json:"prefix"`
}


type RouterIpv6RipDistributeListPrefixPrefixCfg struct {
    PrefixList string `json:"prefix-list"`
    PrefixListDirection string `json:"prefix-list-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}

func (p *RouterIpv6RipDistributeListPrefix) GetId() string{
    return "1"
}

func (p *RouterIpv6RipDistributeListPrefix) getPath() string{
    return "router/ipv6/rip/distribute-list/prefix"
}

func (p *RouterIpv6RipDistributeListPrefix) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeListPrefix::Post")
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

func (p *RouterIpv6RipDistributeListPrefix) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeListPrefix::Get")
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
func (p *RouterIpv6RipDistributeListPrefix) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeListPrefix::Put")
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

func (p *RouterIpv6RipDistributeListPrefix) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeListPrefix::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
