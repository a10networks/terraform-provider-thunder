

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceIpPort struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    FollowPortProtocol string `json:"follow-port-protocol"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckFollowPort int `json:"health-check-follow-port"`
    HealthCheckProtocolDisable int `json:"health-check-protocol-disable"`
    PortNum int `json:"port-num"`
    PortProto string `json:"port-proto"`
    SamplingEnable []GslbServiceIpPortSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    NodeName string 

	} `json:"port"`
}


type GslbServiceIpPortSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbServiceIpPort) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)+"+"+p.Inst.PortProto
}

func (p *GslbServiceIpPort) getPath() string{
    return "gslb/service-ip/" +p.Inst.NodeName + "/port"
}

func (p *GslbServiceIpPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIpPort::Post")
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

func (p *GslbServiceIpPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIpPort::Get")
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
func (p *GslbServiceIpPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIpPort::Put")
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

func (p *GslbServiceIpPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbServiceIpPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
