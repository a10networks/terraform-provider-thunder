

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PartitionGroup struct {
	Inst struct {

    MemberList []PartitionGroupMemberList `json:"member-list"`
    PartitionGroupName string `json:"partition-group-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"partition-group"`
}


type PartitionGroupMemberList struct {
    Member string `json:"member"`
}

func (p *PartitionGroup) GetId() string{
    return p.Inst.PartitionGroupName
}

func (p *PartitionGroup) getPath() string{
    return "partition-group"
}

func (p *PartitionGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionGroup::Post")
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

func (p *PartitionGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionGroup::Get")
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
func (p *PartitionGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionGroup::Put")
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

func (p *PartitionGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PartitionGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
