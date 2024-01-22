

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOneSharedPoolGroup struct {
	Inst struct {

    Members Cgnv6OneToOneSharedPoolGroupMembers109 `json:"members"`
    Uuid string `json:"uuid"`

	} `json:"shared-pool-group"`
}


type Cgnv6OneToOneSharedPoolGroupMembers109 struct {
    Uuid string `json:"uuid"`
}

func (p *Cgnv6OneToOneSharedPoolGroup) GetId() string{
    return "1"
}

func (p *Cgnv6OneToOneSharedPoolGroup) getPath() string{
    return "cgnv6/one-to-one/shared-pool-group"
}

func (p *Cgnv6OneToOneSharedPoolGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOneSharedPoolGroup::Post")
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

func (p *Cgnv6OneToOneSharedPoolGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOneSharedPoolGroup::Get")
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
func (p *Cgnv6OneToOneSharedPoolGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOneSharedPoolGroup::Put")
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

func (p *Cgnv6OneToOneSharedPoolGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOneSharedPoolGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
