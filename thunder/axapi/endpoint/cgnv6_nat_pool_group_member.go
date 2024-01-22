

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatPoolGroupMember struct {
	Inst struct {

    PoolName string `json:"pool-name"`
    Uuid string `json:"uuid"`
    PoolGroupName string 

	} `json:"member"`
}

func (p *Cgnv6NatPoolGroupMember) GetId() string{
    return url.QueryEscape(p.Inst.PoolName)
}

func (p *Cgnv6NatPoolGroupMember) getPath() string{
    return "cgnv6/nat/pool-group/" +p.Inst.PoolGroupName + "/member"
}

func (p *Cgnv6NatPoolGroupMember) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPoolGroupMember::Post")
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

func (p *Cgnv6NatPoolGroupMember) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPoolGroupMember::Get")
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
func (p *Cgnv6NatPoolGroupMember) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPoolGroupMember::Put")
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

func (p *Cgnv6NatPoolGroupMember) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatPoolGroupMember::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
