

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerCommunityRead struct {
	Inst struct {

    OidList []SnmpServerCommunityReadOidList `json:"oid-list"`
    Remote SnmpServerCommunityReadRemote `json:"remote"`
    User string `json:"user"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"read"`
}


type SnmpServerCommunityReadOidList struct {
    OidVal string `json:"oid-val"`
    Remote SnmpServerCommunityReadOidListRemote `json:"remote"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SnmpServerCommunityReadOidListRemote struct {
    HostList []SnmpServerCommunityReadOidListRemoteHostList `json:"host-list"`
    Ipv4List []SnmpServerCommunityReadOidListRemoteIpv4List `json:"ipv4-list"`
    Ipv6List []SnmpServerCommunityReadOidListRemoteIpv6List `json:"ipv6-list"`
}


type SnmpServerCommunityReadOidListRemoteHostList struct {
    DnsHost string `json:"dns-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerCommunityReadOidListRemoteIpv4List struct {
    Ipv4Host string `json:"ipv4-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerCommunityReadOidListRemoteIpv6List struct {
    Ipv6Host string `json:"ipv6-host"`
    Ipv6Mask int `json:"ipv6-mask"`
}


type SnmpServerCommunityReadRemote struct {
    HostList []SnmpServerCommunityReadRemoteHostList `json:"host-list"`
    Ipv4List []SnmpServerCommunityReadRemoteIpv4List `json:"ipv4-list"`
    Ipv6List []SnmpServerCommunityReadRemoteIpv6List `json:"ipv6-list"`
}


type SnmpServerCommunityReadRemoteHostList struct {
    DnsHost string `json:"dns-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerCommunityReadRemoteIpv4List struct {
    Ipv4Host string `json:"ipv4-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerCommunityReadRemoteIpv6List struct {
    Ipv6Host string `json:"ipv6-host"`
    Ipv6Mask int `json:"ipv6-mask"`
}

func (p *SnmpServerCommunityRead) GetId() string{
    return url.QueryEscape(p.Inst.User)
}

func (p *SnmpServerCommunityRead) getPath() string{
    return "snmp-server/community/read"
}

func (p *SnmpServerCommunityRead) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityRead::Post")
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

func (p *SnmpServerCommunityRead) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityRead::Get")
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
func (p *SnmpServerCommunityRead) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityRead::Put")
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

func (p *SnmpServerCommunityRead) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityRead::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
