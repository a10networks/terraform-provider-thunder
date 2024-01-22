

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Server struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    FqdnName string `json:"fqdn-name"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    Host string `json:"host"`
    Name string `json:"name"`
    PortList []Cgnv6ServerPortList `json:"port-list"`
    ResolveAs string `json:"resolve-as" dval:"resolve-to-ipv4"`
    SamplingEnable []Cgnv6ServerSamplingEnable `json:"sampling-enable"`
    ServerIpv6Addr string `json:"server-ipv6-addr"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"server"`
}


type Cgnv6ServerPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Action string `json:"action" dval:"enable"`
    HealthCheck string `json:"health-check"`
    HealthCheckFollowPort int `json:"health-check-follow-port"`
    FollowPortProtocol string `json:"follow-port-protocol"`
    HealthCheckDisable int `json:"health-check-disable"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []Cgnv6ServerPortListSamplingEnable `json:"sampling-enable"`
}


type Cgnv6ServerPortListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6ServerSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6Server) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6Server) getPath() string{
    return "cgnv6/server"
}

func (p *Cgnv6Server) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Server::Post")
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

func (p *Cgnv6Server) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Server::Get")
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
func (p *Cgnv6Server) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Server::Put")
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

func (p *Cgnv6Server) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Server::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
