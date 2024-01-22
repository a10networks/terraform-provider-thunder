

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerCommunityReadOid struct {
	Inst struct {

    OidVal string `json:"oid-val"`
    Remote SnmpServerCommunityReadOidRemote `json:"remote"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    User string 

	} `json:"oid"`
}


type SnmpServerCommunityReadOidRemote struct {
    HostList []SnmpServerCommunityReadOidRemoteHostList `json:"host-list"`
    Ipv4List []SnmpServerCommunityReadOidRemoteIpv4List `json:"ipv4-list"`
    Ipv6List []SnmpServerCommunityReadOidRemoteIpv6List `json:"ipv6-list"`
}


type SnmpServerCommunityReadOidRemoteHostList struct {
    DnsHost string `json:"dns-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerCommunityReadOidRemoteIpv4List struct {
    Ipv4Host string `json:"ipv4-host"`
    Ipv4Mask string `json:"ipv4-mask"`
}


type SnmpServerCommunityReadOidRemoteIpv6List struct {
    Ipv6Host string `json:"ipv6-host"`
    Ipv6Mask int `json:"ipv6-mask"`
}

func (p *SnmpServerCommunityReadOid) GetId() string{
    return p.Inst.OidVal
}

func (p *SnmpServerCommunityReadOid) getPath() string{
    return "snmp-server/community/read/" +p.Inst.User + "/oid"
}

func (p *SnmpServerCommunityReadOid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityReadOid::Post")
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

func (p *SnmpServerCommunityReadOid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityReadOid::Get")
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
func (p *SnmpServerCommunityReadOid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityReadOid::Put")
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

func (p *SnmpServerCommunityReadOid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerCommunityReadOid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
