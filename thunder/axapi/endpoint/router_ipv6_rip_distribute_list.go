

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6RipDistributeList struct {
	Inst struct {

    AclCfg []RouterIpv6RipDistributeListAclCfg `json:"acl-cfg"`
    Prefix RouterIpv6RipDistributeListPrefix1255 `json:"prefix"`
    Uuid string `json:"uuid"`

	} `json:"distribute-list"`
}


type RouterIpv6RipDistributeListAclCfg struct {
    Acl string `json:"acl"`
    AclDirection string `json:"acl-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterIpv6RipDistributeListPrefix1255 struct {
    PrefixCfg []RouterIpv6RipDistributeListPrefixPrefixCfg1256 `json:"prefix-cfg"`
    Uuid string `json:"uuid"`
}


type RouterIpv6RipDistributeListPrefixPrefixCfg1256 struct {
    PrefixList string `json:"prefix-list"`
    PrefixListDirection string `json:"prefix-list-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}

func (p *RouterIpv6RipDistributeList) GetId() string{
    return "1"
}

func (p *RouterIpv6RipDistributeList) getPath() string{
    return "router/ipv6/rip/distribute-list"
}

func (p *RouterIpv6RipDistributeList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeList::Post")
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

func (p *RouterIpv6RipDistributeList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeList::Get")
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
func (p *RouterIpv6RipDistributeList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeList::Put")
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

func (p *RouterIpv6RipDistributeList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipDistributeList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
