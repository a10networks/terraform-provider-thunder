

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterRipDistributeListPrefix struct {
	Inst struct {

    PrefixCfg []RouterRipDistributeListPrefixPrefixCfg `json:"prefix-cfg"`
    Uuid string `json:"uuid"`

	} `json:"prefix"`
}


type RouterRipDistributeListPrefixPrefixCfg struct {
    PrefixList string `json:"prefix-list"`
    PrefixListDirection string `json:"prefix-list-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}

func (p *RouterRipDistributeListPrefix) GetId() string{
    return "1"
}

func (p *RouterRipDistributeListPrefix) getPath() string{
    return "router/rip/distribute-list/prefix"
}

func (p *RouterRipDistributeListPrefix) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeListPrefix::Post")
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

func (p *RouterRipDistributeListPrefix) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeListPrefix::Get")
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
func (p *RouterRipDistributeListPrefix) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeListPrefix::Put")
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

func (p *RouterRipDistributeListPrefix) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRipDistributeListPrefix::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
