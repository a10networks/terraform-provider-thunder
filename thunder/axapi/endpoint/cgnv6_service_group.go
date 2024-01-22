

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ServiceGroup struct {
	Inst struct {

    HealthCheck string `json:"health-check"`
    MemberList []Cgnv6ServiceGroupMemberList `json:"member-list"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Protocol string `json:"protocol"`
    SamplingEnable []Cgnv6ServiceGroupSamplingEnable `json:"sampling-enable"`
    Shared int `json:"shared"`
    SharedGroup string `json:"shared-group"`
    SharedPartition string `json:"shared-partition"`
    TrafficReplicationMirrorIpRepl int `json:"traffic-replication-mirror-ip-repl"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service-group"`
}


type Cgnv6ServiceGroupMemberList struct {
    Name string `json:"name"`
    Port int `json:"port"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []Cgnv6ServiceGroupMemberListSamplingEnable `json:"sampling-enable"`
}


type Cgnv6ServiceGroupMemberListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6ServiceGroupSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6ServiceGroup) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6ServiceGroup) getPath() string{
    return "cgnv6/service-group"
}

func (p *Cgnv6ServiceGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroup::Post")
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

func (p *Cgnv6ServiceGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroup::Get")
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
func (p *Cgnv6ServiceGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroup::Put")
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

func (p *Cgnv6ServiceGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
