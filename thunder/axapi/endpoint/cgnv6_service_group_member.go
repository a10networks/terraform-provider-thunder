

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ServiceGroupMember struct {
	Inst struct {

    Name string `json:"name"`
    Port int `json:"port"`
    SamplingEnable []Cgnv6ServiceGroupMemberSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"member"`
}


type Cgnv6ServiceGroupMemberSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6ServiceGroupMember) GetId() string{
    return p.Inst.Name+"+"+strconv.Itoa(p.Inst.Port)
}

func (p *Cgnv6ServiceGroupMember) getPath() string{
    return "cgnv6/service-group/"+p.Inst.Name+"/member"
}

func (p *Cgnv6ServiceGroupMember) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroupMember::Post")
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

func (p *Cgnv6ServiceGroupMember) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroupMember::Get")
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
func (p *Cgnv6ServiceGroupMember) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroupMember::Put")
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

func (p *Cgnv6ServiceGroupMember) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6ServiceGroupMember::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
