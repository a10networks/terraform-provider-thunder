

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterRipDistributeList struct {
	Inst struct {

    AclCfg []RouterRipDistributeListAclCfg `json:"acl-cfg"`
    Prefix RouterRipDistributeListPrefix1304 `json:"prefix"`
    Uuid string `json:"uuid"`

	} `json:"distribute-list"`
}


type RouterRipDistributeListAclCfg struct {
    Acl string `json:"acl"`
    AclDirection string `json:"acl-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterRipDistributeListPrefix1304 struct {
    PrefixCfg []RouterRipDistributeListPrefixPrefixCfg1305 `json:"prefix-cfg"`
    Uuid string `json:"uuid"`
}


type RouterRipDistributeListPrefixPrefixCfg1305 struct {
    PrefixList string `json:"prefix-list"`
    PrefixListDirection string `json:"prefix-list-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}

func (p *RouterRipDistributeList) GetId() string{
    return "1"
}

func (p *RouterRipDistributeList) getPath() string{
    return "router/rip/distribute-list"
}

func (p *RouterRipDistributeList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeList::Post")
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

func (p *RouterRipDistributeList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeList::Get")
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
func (p *RouterRipDistributeList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeList::Put")
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

func (p *RouterRipDistributeList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
