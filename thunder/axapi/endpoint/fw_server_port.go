

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type FwServerPort struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    SamplingEnable []FwServerPortSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"port"`
}


type FwServerPortSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *FwServerPort) GetId() string{
    return strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol
}

func (p *FwServerPort) getPath() string{
    return "fw/server/" +p.Inst.Name + "/port"
}

func (p *FwServerPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwServerPort::Post")
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

func (p *FwServerPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwServerPort::Get")
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
func (p *FwServerPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwServerPort::Put")
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

func (p *FwServerPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwServerPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
