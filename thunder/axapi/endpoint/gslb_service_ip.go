

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceIp struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    ExternalIp string `json:"external-ip"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckProtocolDisable int `json:"health-check-protocol-disable"`
    IpAddress string `json:"ip-address"`
    Ipv6 string `json:"ipv6"`
    Ipv6Address string `json:"ipv6-address"`
    NodeName string `json:"node-name"`
    PortList []GslbServiceIpPortList `json:"port-list"`
    SamplingEnable []GslbServiceIpSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service-ip"`
}


type GslbServiceIpPortList struct {
    PortNum int `json:"port-num"`
    PortProto string `json:"port-proto"`
    Action string `json:"action" dval:"enable"`
    HealthCheck string `json:"health-check"`
    HealthCheckFollowPort int `json:"health-check-follow-port"`
    FollowPortProtocol string `json:"follow-port-protocol"`
    HealthCheckProtocolDisable int `json:"health-check-protocol-disable"`
    HealthCheckDisable int `json:"health-check-disable"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []GslbServiceIpPortListSamplingEnable `json:"sampling-enable"`
}


type GslbServiceIpPortListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbServiceIpSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbServiceIp) GetId() string{
    return p.Inst.NodeName
}

func (p *GslbServiceIp) getPath() string{
    return "gslb/service-ip"
}

func (p *GslbServiceIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIp::Post")
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

func (p *GslbServiceIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIp::Get")
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
func (p *GslbServiceIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIp::Put")
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

func (p *GslbServiceIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
