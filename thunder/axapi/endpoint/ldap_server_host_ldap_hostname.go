

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type LdapServerHostLdapHostname struct {
	Inst struct {

    Base string `json:"base"`
    CnValue string `json:"cn-value"`
    DnValue string `json:"dn-value"`
    Domain string `json:"domain"`
    DomainCfg LdapServerHostLdapHostnameDomainCfg `json:"domain-cfg"`
    Group string `json:"group"`
    Hostname string `json:"hostname"`
    HostnameCfg LdapServerHostLdapHostnameHostnameCfg `json:"hostname-cfg"`
    PortCfg LdapServerHostLdapHostnamePortCfg `json:"port-cfg"`
    Uuid string `json:"uuid"`

	} `json:"ldap-hostname"`
}


type LdapServerHostLdapHostnameDomainCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}


type LdapServerHostLdapHostnameHostnameCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}


type LdapServerHostLdapHostnamePortCfg struct {
    Port int `json:"port"`
    Ssl int `json:"ssl"`
    Timeout int `json:"timeout"`
}

func (p *LdapServerHostLdapHostname) GetId() string{
    return url.QueryEscape(p.Inst.Hostname)
}

func (p *LdapServerHostLdapHostname) getPath() string{
    return "ldap-server/host/ldap-hostname"
}

func (p *LdapServerHostLdapHostname) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostLdapHostname::Post")
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

func (p *LdapServerHostLdapHostname) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostLdapHostname::Get")
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
func (p *LdapServerHostLdapHostname) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostLdapHostname::Put")
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

func (p *LdapServerHostLdapHostname) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LdapServerHostLdapHostname::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
