

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatSharedPoolGroup struct {
	Inst struct {

    Members Cgnv6NatSharedPoolGroupMembers108 `json:"members"`
    Uuid string `json:"uuid"`

	} `json:"shared-pool-group"`
}


type Cgnv6NatSharedPoolGroupMembers108 struct {
    Uuid string `json:"uuid"`
}

func (p *Cgnv6NatSharedPoolGroup) GetId() string{
    return "1"
}

func (p *Cgnv6NatSharedPoolGroup) getPath() string{
    return "cgnv6/nat/shared-pool-group"
}

func (p *Cgnv6NatSharedPoolGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatSharedPoolGroup::Post")
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

func (p *Cgnv6NatSharedPoolGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatSharedPoolGroup::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *Cgnv6NatSharedPoolGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatSharedPoolGroup::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *Cgnv6NatSharedPoolGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6NatSharedPoolGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
