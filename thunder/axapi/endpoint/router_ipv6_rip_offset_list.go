

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6RipOffsetList struct {
	Inst struct {

    AclCfg []RouterIpv6RipOffsetListAclCfg `json:"acl-cfg"`
    Uuid string `json:"uuid"`

	} `json:"offset-list"`
}


type RouterIpv6RipOffsetListAclCfg struct {
    Acl string `json:"acl"`
    OffsetListDirection string `json:"offset-list-direction"`
    Metric int `json:"metric"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}

func (p *RouterIpv6RipOffsetList) GetId() string{
    return "1"
}

func (p *RouterIpv6RipOffsetList) getPath() string{
    return "router/ipv6/rip/offset-list"
}

func (p *RouterIpv6RipOffsetList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipOffsetList::Post")
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

func (p *RouterIpv6RipOffsetList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipOffsetList::Get")
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
func (p *RouterIpv6RipOffsetList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipOffsetList::Put")
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

func (p *RouterIpv6RipOffsetList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipOffsetList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
