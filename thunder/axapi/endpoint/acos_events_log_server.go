

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsLogServer struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    Host string `json:"host"`
    Name string `json:"name"`
    PortList []AcosEventsLogServerPortList `json:"port-list"`
    ResolveAs string `json:"resolve-as" dval:"resolve-to-ipv4"`
    SamplingEnable []AcosEventsLogServerSamplingEnable `json:"sampling-enable"`
    ServerIpv6Addr string `json:"server-ipv6-addr"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"log-server"`
}


type AcosEventsLogServerPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Action string `json:"action" dval:"enable"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []AcosEventsLogServerPortListSamplingEnable `json:"sampling-enable"`
}


type AcosEventsLogServerPortListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AcosEventsLogServerSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AcosEventsLogServer) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AcosEventsLogServer) getPath() string{
    return "acos-events/log-server"
}

func (p *AcosEventsLogServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogServer::Post")
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

func (p *AcosEventsLogServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogServer::Get")
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
func (p *AcosEventsLogServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogServer::Put")
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

func (p *AcosEventsLogServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
