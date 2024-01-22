

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ServerPort struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    FollowPortProtocol string `json:"follow-port-protocol"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckFollowPort int `json:"health-check-follow-port"`
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    SamplingEnable []Cgnv6ServerPortSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"port"`
}


type Cgnv6ServerPortSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6ServerPort) GetId() string{
    return strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol
}

func (p *Cgnv6ServerPort) getPath() string{
    return "cgnv6/server/" +p.Inst.Name + "/port"
}

func (p *Cgnv6ServerPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServerPort::Post")
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

func (p *Cgnv6ServerPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServerPort::Get")
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
func (p *Cgnv6ServerPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServerPort::Put")
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

func (p *Cgnv6ServerPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServerPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
