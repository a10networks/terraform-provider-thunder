

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatRangeList struct {
	Inst struct {

    GlobalNetmaskv4 string `json:"global-netmaskv4"`
    GlobalStartIpv4Addr string `json:"global-start-ipv4-addr"`
    GlobalStartIpv6Addr string `json:"global-start-ipv6-addr"`
    LocalNetmaskv4 string `json:"local-netmaskv4"`
    LocalStartIpv4Addr string `json:"local-start-ipv4-addr"`
    LocalStartIpv6Addr string `json:"local-start-ipv6-addr"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    V4AclId int `json:"v4-acl-id"`
    V4AclName string `json:"v4-acl-name"`
    V4Count int `json:"v4-count"`
    V4Vrid int `json:"v4-vrid"`
    V6AclName string `json:"v6-acl-name"`
    V6Count int `json:"v6-count"`
    V6Vrid int `json:"v6-vrid"`

	} `json:"range-list"`
}

func (p *IpNatRangeList) GetId() string{
    return p.Inst.Name
}

func (p *IpNatRangeList) getPath() string{
    return "ip/nat/range-list"
}

func (p *IpNatRangeList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatRangeList::Post")
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

func (p *IpNatRangeList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatRangeList::Get")
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
func (p *IpNatRangeList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatRangeList::Put")
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

func (p *IpNatRangeList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatRangeList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
