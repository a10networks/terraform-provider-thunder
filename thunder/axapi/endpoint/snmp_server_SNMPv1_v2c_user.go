

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerSnmpv1V2cUser struct {
	Inst struct {

    Encrypted string `json:"encrypted"`
    OidList []SnmpServerSnmpv1V2cUserOidList `json:"oid-list"`
    Passwd string `json:"passwd"`
    Remote SnmpServerSnmpv1V2cUserRemote `json:"remote"`
    User string `json:"user"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"user"`
}


type SnmpServerSnmpv1V2cUserOidList struct {
    OidVal string `json:"oid-val"`
    Remote SnmpServerSnmpv1V2cUserOidListRemote `json:"remote"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SnmpServerSnmpv1V2cUserOidListRemote struct {
    HostList []SnmpServerSnmpv1V2cUserOidListRemoteHostList `json:"host-list"`
    Ipv4List []SnmpServerSnmpv1V2cUserOidListRemoteIpv4List `json:"ipv4-list"`
    Ipv6List []SnmpServerSnmpv1V2cUserOidListRemoteIpv6List `json:"ipv6-list"`
}


type SnmpServerSnmpv1V2cUserOidListRemoteHostList struct {
    DnsHost string `json:"dns-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerSnmpv1V2cUserOidListRemoteIpv4List struct {
    Ipv4Host string `json:"ipv4-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerSnmpv1V2cUserOidListRemoteIpv6List struct {
    Ipv6Host string `json:"ipv6-host"`
    Ipv6Mask int `json:"ipv6-mask"`
}


type SnmpServerSnmpv1V2cUserRemote struct {
    HostList []SnmpServerSnmpv1V2cUserRemoteHostList `json:"host-list"`
    Ipv4List []SnmpServerSnmpv1V2cUserRemoteIpv4List `json:"ipv4-list"`
    Ipv6List []SnmpServerSnmpv1V2cUserRemoteIpv6List `json:"ipv6-list"`
}


type SnmpServerSnmpv1V2cUserRemoteHostList struct {
    DnsHost string `json:"dns-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerSnmpv1V2cUserRemoteIpv4List struct {
    Ipv4Host string `json:"ipv4-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerSnmpv1V2cUserRemoteIpv6List struct {
    Ipv6Host string `json:"ipv6-host"`
    Ipv6Mask int `json:"ipv6-mask"`
}

func (p *SnmpServerSnmpv1V2cUser) GetId() string{
    return url.QueryEscape(p.Inst.User)
}

func (p *SnmpServerSnmpv1V2cUser) getPath() string{
    return "snmp-server/SNMPv1-v2c/user"
}

func (p *SnmpServerSnmpv1V2cUser) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUser::Post")
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

func (p *SnmpServerSnmpv1V2cUser) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUser::Get")
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
func (p *SnmpServerSnmpv1V2cUser) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUser::Put")
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

func (p *SnmpServerSnmpv1V2cUser) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUser::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
