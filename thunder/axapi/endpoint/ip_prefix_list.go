

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpPrefixList struct {
	Inst struct {

    Name string `json:"name"`
    Rules []IpPrefixListRules `json:"rules"`
    Uuid string `json:"uuid"`

	} `json:"prefix-list"`
}


type IpPrefixListRules struct {
    Seq int `json:"seq"`
    Description string `json:"description"`
    Action string `json:"action"`
    Any int `json:"any"`
    Ipaddr string `json:"ipaddr"`
    Ge int `json:"ge"`
    Le int `json:"le"`
}

func (p *IpPrefixList) GetId() string{
    return p.Inst.Name
}

func (p *IpPrefixList) getPath() string{
    return "ip/prefix-list"
}

func (p *IpPrefixList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpPrefixList::Post")
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

func (p *IpPrefixList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpPrefixList::Get")
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
func (p *IpPrefixList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpPrefixList::Put")
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

func (p *IpPrefixList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpPrefixList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
