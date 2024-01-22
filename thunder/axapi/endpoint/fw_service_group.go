

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type FwServiceGroup struct {
	Inst struct {

    HealthCheck string `json:"health-check"`
    MemberList []FwServiceGroupMemberList `json:"member-list"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Protocol string `json:"protocol"`
    SamplingEnable []FwServiceGroupSamplingEnable `json:"sampling-enable"`
    TrafficReplicationMirrorIpRepl int `json:"traffic-replication-mirror-ip-repl"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service-group"`
}


type FwServiceGroupMemberList struct {
    Name string `json:"name"`
    Port int `json:"port"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []FwServiceGroupMemberListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type FwServiceGroupMemberListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type FwServiceGroupSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *FwServiceGroup) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *FwServiceGroup) getPath() string{
    return "fw/service-group"
}

func (p *FwServiceGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwServiceGroup::Post")
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

func (p *FwServiceGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwServiceGroup::Get")
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
func (p *FwServiceGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwServiceGroup::Put")
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

func (p *FwServiceGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwServiceGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
