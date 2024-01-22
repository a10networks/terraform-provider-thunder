

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbServiceGroupMember struct {
	Inst struct {

    FqdnName string `json:"fqdn-name"`
    Host string `json:"host"`
    MemberPriority int `json:"member-priority" dval:"1"`
    MemberState string `json:"member-state" dval:"enable"`
    MemberStatsDataDisable int `json:"member-stats-data-disable"`
    MemberTemplate string `json:"member-template"`
    Name string `json:"name"`
    Port int `json:"port"`
    ResolveAs string `json:"resolve-as" dval:"resolve-to-ipv4"`
    SamplingEnable []SlbServiceGroupMemberSamplingEnable `json:"sampling-enable"`
    ServerIpv6Addr string `json:"server-ipv6-addr"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ServiceGroupName string `json:"-"`
	} `json:"member"`
}


type SlbServiceGroupMemberSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbServiceGroupMember) GetId() string{
    return p.Inst.Name+"+"+strconv.Itoa(p.Inst.Port)
}

func (p *SlbServiceGroupMember) getPath() string{
    return "slb/service-group/"+p.Inst.ServiceGroupName+"/member"
}

func (p *SlbServiceGroupMember) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroupMember::Post")
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

func (p *SlbServiceGroupMember) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroupMember::Get")
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
func (p *SlbServiceGroupMember) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroupMember::Put")
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

func (p *SlbServiceGroupMember) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroupMember::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
