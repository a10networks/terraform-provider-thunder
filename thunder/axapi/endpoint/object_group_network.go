

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type ObjectGroupNetwork struct {
	Inst struct {

    Description string `json:"description"`
    IpVersion string `json:"ip-version"`
    NetName string `json:"net-name"`
    Rules []ObjectGroupNetworkRules `json:"rules"`
    Usage string `json:"usage" dval:"acl"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"network"`
}


type ObjectGroupNetworkRules struct {
    SeqNum int `json:"seq-num"`
    HostV4 string `json:"host-v4"`
    HostV6 string `json:"host-v6"`
    IpRangeStart string `json:"ip-range-start"`
    IpRangeEnd string `json:"ip-range-end"`
    Ipv6RangeStart string `json:"ipv6-range-start"`
    Ipv6RangeEnd string `json:"ipv6-range-end"`
    Any int `json:"any"`
    Subnet string `json:"subnet"`
    RevSubnetMask string `json:"rev-subnet-mask"`
    FwIpv4Address string `json:"fw-ipv4-address"`
    Ipv6Subnet string `json:"ipv6-subnet"`
    FwIpv6Subnet string `json:"fw-ipv6-subnet"`
    ObjNetwork string `json:"obj-network"`
    SlbServer string `json:"slb-server"`
    SlbVserver string `json:"slb-vserver"`
}

func (p *ObjectGroupNetwork) GetId() string{
    return url.QueryEscape(p.Inst.NetName)
}

func (p *ObjectGroupNetwork) getPath() string{
    return "object-group/network"
}

func (p *ObjectGroupNetwork) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupNetwork::Post")
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

func (p *ObjectGroupNetwork) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupNetwork::Get")
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
func (p *ObjectGroupNetwork) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupNetwork::Put")
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

func (p *ObjectGroupNetwork) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupNetwork::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
