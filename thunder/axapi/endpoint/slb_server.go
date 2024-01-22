

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbServer struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    AlternateServer []SlbServerAlternateServer `json:"alternate-server"`
    ConnLimit int `json:"conn-limit" dval:"64000000"`
    ConnResume int `json:"conn-resume"`
    Ethernet int `json:"ethernet"`
    ExtendedStats int `json:"extended-stats"`
    ExternalIp string `json:"external-ip"`
    FqdnName string `json:"fqdn-name"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckShared string `json:"health-check-shared"`
    Host string `json:"host"`
    Ipv6 string `json:"ipv6"`
    L2HealthCheckPath string `json:"l2-health-check-path"`
    Name string `json:"name"`
    NoLogging int `json:"no-logging"`
    PortList []SlbServerPortList `json:"port-list"`
    ResolveAs string `json:"resolve-as" dval:"resolve-to-ipv4"`
    SamplingEnable []SlbServerSamplingEnable `json:"sampling-enable"`
    ServerIpv6Addr string `json:"server-ipv6-addr"`
    SharedPartitionHealthCheck int `json:"shared-partition-health-check"`
    SharedPartitionServerTemplate int `json:"shared-partition-server-template"`
    SlowStart int `json:"slow-start"`
    SpoofingCache int `json:"spoofing-cache"`
    StatsDataAction string `json:"stats-data-action" dval:"stats-data-enable"`
    TemplateLinkCost string `json:"template-link-cost"`
    TemplateServer string `json:"template-server" dval:"default"`
    TemplateServerShared string `json:"template-server-shared"`
    Trunk int `json:"trunk"`
    UseAamServer int `json:"use-aam-server"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight" dval:"1"`

	} `json:"server"`
}


type SlbServerAlternateServer struct {
    Alternate int `json:"alternate"`
    AlternateName string `json:"alternate-name"`
}


type SlbServerPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Range int `json:"range"`
    TemplatePort string `json:"template-port" dval:"default"`
    SharedPartitionPortTemplate int `json:"shared-partition-port-template"`
    TemplatePortShared string `json:"template-port-shared"`
    TemplateServerSsl string `json:"template-server-ssl"`
    Action string `json:"action" dval:"enable"`
    NoSsl int `json:"no-ssl"`
    HealthCheck string `json:"health-check"`
    SharedRportHealthCheck int `json:"shared-rport-health-check"`
    RportHealthCheckShared string `json:"rport-health-check-shared"`
    HealthCheckFollowPort int `json:"health-check-follow-port"`
    FollowPortProtocol string `json:"follow-port-protocol"`
    HealthCheckDisable int `json:"health-check-disable"`
    SupportHttp2 int `json:"support-http2"`
    Only int `json:"only"`
    Weight int `json:"weight" dval:"1"`
    ConnLimit int `json:"conn-limit" dval:"64000000"`
    NoLogging int `json:"no-logging"`
    ConnResume int `json:"conn-resume"`
    StatsDataAction string `json:"stats-data-action" dval:"stats-data-enable"`
    ExtendedStats int `json:"extended-stats"`
    AlternatePort []SlbServerPortListAlternatePort `json:"alternate-port"`
    AuthCfg SlbServerPortListAuthCfg `json:"auth-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbServerPortListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type SlbServerPortListAlternatePort struct {
    Alternate int `json:"alternate"`
    AlternateName string `json:"alternate-name"`
    AlternateServerPort int `json:"alternate-server-port"`
}


type SlbServerPortListAuthCfg struct {
    ServicePrincipalName string `json:"service-principal-name"`
}


type SlbServerPortListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbServerSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbServer) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbServer) getPath() string{
    return "slb/server"
}

func (p *SlbServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServer::Post")
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

func (p *SlbServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServer::Get")
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
func (p *SlbServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServer::Put")
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

func (p *SlbServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
