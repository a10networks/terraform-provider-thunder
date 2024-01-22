

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDoh struct {
	Inst struct {

    ConnReuse string `json:"conn-reuse" dval:"disable"`
    Dns string `json:"dns" dval:"default"`
    DnsRetry SlbTemplateDohDnsRetry1445 `json:"dns-retry"`
    Forwarder SlbTemplateDohForwarder1446 `json:"forwarder"`
    Name string `json:"name"`
    NonDnsRequest string `json:"non-dns-request" dval:"reject"`
    RejectStatusCode string `json:"reject-status-code" dval:"400"`
    SharedPartitionDnsTemplate int `json:"shared-partition-dns-template"`
    SharedPartitionTcpProxyTemplate int `json:"shared-partition-tcp-proxy-template"`
    SnatPool string `json:"snat-pool"`
    SourceNat string `json:"source-nat" dval:"auto"`
    TcpProxy string `json:"tcp-proxy" dval:"default"`
    TemplateDnsShared string `json:"template-dns-shared"`
    TemplateTcpProxyShared string `json:"template-tcp-proxy-shared"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"doh"`
}


type SlbTemplateDohDnsRetry1445 struct {
    RetryInterval int `json:"retry-interval" dval:"10"`
    AfterTimeout string `json:"after-timeout" dval:"close"`
    MaxTrials int `json:"max-trials" dval:"3"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDohForwarder1446 struct {
    ForwardingIpv4 string `json:"forwarding-ipv4"`
    V4Internal int `json:"v4-internal"`
    V4Port int `json:"v4-port" dval:"53"`
    V4L4Proto string `json:"v4-l4-proto" dval:"both"`
    ForwardingIpv6 string `json:"forwarding-ipv6"`
    V6Internal int `json:"v6-internal"`
    V6Port int `json:"v6-port" dval:"53"`
    V6L4Proto string `json:"v6-l4-proto" dval:"both"`
    TcpServiceGroup string `json:"tcp-service-group"`
    UdpServiceGroup string `json:"udp-service-group"`
    BypassDoh int `json:"bypass-doh"`
    Uuid string `json:"uuid"`
}

func (p *SlbTemplateDoh) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDoh) getPath() string{
    return "slb/template/doh"
}

func (p *SlbTemplateDoh) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDoh::Post")
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

func (p *SlbTemplateDoh) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDoh::Get")
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
func (p *SlbTemplateDoh) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDoh::Put")
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

func (p *SlbTemplateDoh) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDoh::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
