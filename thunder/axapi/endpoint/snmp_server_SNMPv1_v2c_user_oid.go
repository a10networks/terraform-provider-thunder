

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerSnmpv1V2cUserOid struct {
	Inst struct {

    OidVal string `json:"oid-val"`
    Remote SnmpServerSnmpv1V2cUserOidRemote `json:"remote"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    User string 

	} `json:"oid"`
}


type SnmpServerSnmpv1V2cUserOidRemote struct {
    HostList []SnmpServerSnmpv1V2cUserOidRemoteHostList `json:"host-list"`
    Ipv4List []SnmpServerSnmpv1V2cUserOidRemoteIpv4List `json:"ipv4-list"`
    Ipv6List []SnmpServerSnmpv1V2cUserOidRemoteIpv6List `json:"ipv6-list"`
}


type SnmpServerSnmpv1V2cUserOidRemoteHostList struct {
    DnsHost string `json:"dns-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerSnmpv1V2cUserOidRemoteIpv4List struct {
    Ipv4Host string `json:"ipv4-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerSnmpv1V2cUserOidRemoteIpv6List struct {
    Ipv6Host string `json:"ipv6-host"`
    Ipv6Mask int `json:"ipv6-mask"`
}

func (p *SnmpServerSnmpv1V2cUserOid) GetId() string{
    return p.Inst.OidVal
}

func (p *SnmpServerSnmpv1V2cUserOid) getPath() string{
    return "snmp-server/SNMPv1-v2c/user/" +p.Inst.User + "/oid"
}

func (p *SnmpServerSnmpv1V2cUserOid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUserOid::Post")
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

func (p *SnmpServerSnmpv1V2cUserOid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUserOid::Get")
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
func (p *SnmpServerSnmpv1V2cUserOid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUserOid::Put")
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

func (p *SnmpServerSnmpv1V2cUserOid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerSnmpv1V2cUserOid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
