

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LdapServerHostIpv4 struct {
	Inst struct {

    Base string `json:"base"`
    CnValue string `json:"cn-value"`
    DnValue string `json:"dn-value"`
    Domain string `json:"domain"`
    DomainCfg LdapServerHostIpv4DomainCfg `json:"domain-cfg"`
    Group string `json:"group"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv4AddrCfg LdapServerHostIpv4Ipv4AddrCfg `json:"ipv4-addr-cfg"`
    PortCfg LdapServerHostIpv4PortCfg `json:"port-cfg"`
    Uuid string `json:"uuid"`

	} `json:"ipv4"`
}


type LdapServerHostIpv4DomainCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}


type LdapServerHostIpv4Ipv4AddrCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}


type LdapServerHostIpv4PortCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}

func (p *LdapServerHostIpv4) GetId() string{
    return p.Inst.Ipv4Addr
}

func (p *LdapServerHostIpv4) getPath() string{
    return "ldap-server/host/ipv4"
}

func (p *LdapServerHostIpv4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv4::Post")
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

func (p *LdapServerHostIpv4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv4::Get")
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
func (p *LdapServerHostIpv4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv4::Put")
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

func (p *LdapServerHostIpv4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostIpv4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
