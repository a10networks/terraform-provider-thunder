

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodDns struct {
	Inst struct {

    Dns int `json:"dns"`
    DnsDomain string `json:"dns-domain"`
    DnsDomainExpect HealthMonitorMethodDnsDnsDomainExpect `json:"dns-domain-expect"`
    DnsDomainPort int `json:"dns-domain-port" dval:"53"`
    DnsDomainRecurse string `json:"dns-domain-recurse" dval:"enabled"`
    DnsDomainTcp int `json:"dns-domain-tcp"`
    DnsDomainType string `json:"dns-domain-type" dval:"A"`
    DnsIpKey int `json:"dns-ip-key"`
    DnsIpv4Addr string `json:"dns-ipv4-addr"`
    DnsIpv4Expect HealthMonitorMethodDnsDnsIpv4Expect `json:"dns-ipv4-expect"`
    DnsIpv4Port int `json:"dns-ipv4-port" dval:"53"`
    DnsIpv4Recurse string `json:"dns-ipv4-recurse" dval:"enabled"`
    DnsIpv4Tcp int `json:"dns-ipv4-tcp"`
    DnsIpv6Addr string `json:"dns-ipv6-addr"`
    DnsIpv6Expect HealthMonitorMethodDnsDnsIpv6Expect `json:"dns-ipv6-expect"`
    DnsIpv6Port int `json:"dns-ipv6-port" dval:"53"`
    DnsIpv6Recurse string `json:"dns-ipv6-recurse" dval:"enabled"`
    DnsIpv6Tcp int `json:"dns-ipv6-tcp"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"dns"`
}


type HealthMonitorMethodDnsDnsDomainExpect struct {
    DnsDomainResponse string `json:"dns-domain-response"`
    DnsDomainFqdn string `json:"dns-domain-fqdn"`
    DnsDomainIpv4 string `json:"dns-domain-ipv4"`
    DnsDomainIpv6 string `json:"dns-domain-ipv6"`
}


type HealthMonitorMethodDnsDnsIpv4Expect struct {
    DnsIpv4Response string `json:"dns-ipv4-response"`
    DnsIpv4Fqdn string `json:"dns-ipv4-fqdn"`
}


type HealthMonitorMethodDnsDnsIpv6Expect struct {
    DnsIpv6Response string `json:"dns-ipv6-response"`
    DnsIpv6Fqdn string `json:"dns-ipv6-fqdn"`
}

func (p *HealthMonitorMethodDns) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodDns) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/dns"
}

func (p *HealthMonitorMethodDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDns::Post")
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

func (p *HealthMonitorMethodDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDns::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *HealthMonitorMethodDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDns::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *HealthMonitorMethodDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
