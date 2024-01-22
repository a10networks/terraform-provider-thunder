

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LdapServerHostIpv6 struct {
	Inst struct {

    Base string `json:"base"`
    CnValue string `json:"cn-value"`
    DnValue string `json:"dn-value"`
    Domain string `json:"domain"`
    DomainCfg LdapServerHostIpv6DomainCfg `json:"domain-cfg"`
    Group string `json:"group"`
    Ipv6Addr string `json:"ipv6-addr"`
    Ipv6AddrCfg LdapServerHostIpv6Ipv6AddrCfg `json:"ipv6-addr-cfg"`
    PortCfg LdapServerHostIpv6PortCfg `json:"port-cfg"`
    Uuid string `json:"uuid"`

	} `json:"ipv6"`
}


type LdapServerHostIpv6DomainCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}


type LdapServerHostIpv6Ipv6AddrCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}


type LdapServerHostIpv6PortCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}

func (p *LdapServerHostIpv6) GetId() string{
    return p.Inst.Ipv6Addr
}

func (p *LdapServerHostIpv6) getPath() string{
    return "ldap-server/host/ipv6"
}

func (p *LdapServerHostIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv6::Post")
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

func (p *LdapServerHostIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv6::Get")
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
func (p *LdapServerHostIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv6::Put")
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

func (p *LdapServerHostIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
