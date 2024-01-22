

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbIpv6ClassList struct {
	Inst struct {

    Ipv6List []SlbIpv6ClassListIpv6List `json:"ipv6-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`

	} `json:"ipv6-class-list"`
}


type SlbIpv6ClassListIpv6List struct {
    Action string `json:"action"`
    Ipv6Addr string `json:"ipv6-addr"`
    Lid int `json:"lid"`
    Glid string `json:"glid"`
    LsnLid int `json:"lsn-lid"`
    LsnRadiusProfile int `json:"lsn-radius-profile"`
}

func (p *SlbIpv6ClassList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbIpv6ClassList) getPath() string{
    return "slb/ipv6-class-list"
}

func (p *SlbIpv6ClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIpv6ClassList::Post")
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

func (p *SlbIpv6ClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIpv6ClassList::Get")
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
func (p *SlbIpv6ClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIpv6ClassList::Put")
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

func (p *SlbIpv6ClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIpv6ClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
