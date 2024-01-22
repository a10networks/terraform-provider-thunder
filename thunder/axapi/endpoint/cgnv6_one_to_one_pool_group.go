

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOnePoolGroup struct {
	Inst struct {

    MemberList []Cgnv6OneToOnePoolGroupMemberList `json:"member-list"`
    PoolGroupName string `json:"pool-group-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"pool-group"`
}


type Cgnv6OneToOnePoolGroupMemberList struct {
    PoolName string `json:"pool-name"`
    Uuid string `json:"uuid"`
}

func (p *Cgnv6OneToOnePoolGroup) GetId() string{
    return url.QueryEscape(p.Inst.PoolGroupName)
}

func (p *Cgnv6OneToOnePoolGroup) getPath() string{
    return "cgnv6/one-to-one/pool-group"
}

func (p *Cgnv6OneToOnePoolGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroup::Post")
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

func (p *Cgnv6OneToOnePoolGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroup::Get")
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
func (p *Cgnv6OneToOnePoolGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroup::Put")
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

func (p *Cgnv6OneToOnePoolGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
