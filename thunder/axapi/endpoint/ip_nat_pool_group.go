

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type IpNatPoolGroup struct {
	Inst struct {

    MemberList []IpNatPoolGroupMemberList `json:"member-list"`
    PoolGroupName string `json:"pool-group-name"`
    SamplingEnable []IpNatPoolGroupSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"pool-group"`
}


type IpNatPoolGroupMemberList struct {
    PoolName string `json:"pool-name"`
    Uuid string `json:"uuid"`
}


type IpNatPoolGroupSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *IpNatPoolGroup) GetId() string{
    return url.QueryEscape(p.Inst.PoolGroupName)
}

func (p *IpNatPoolGroup) getPath() string{
    return "ip/nat/pool-group"
}

func (p *IpNatPoolGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatPoolGroup::Post")
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

func (p *IpNatPoolGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatPoolGroup::Get")
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
func (p *IpNatPoolGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatPoolGroup::Put")
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

func (p *IpNatPoolGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatPoolGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
