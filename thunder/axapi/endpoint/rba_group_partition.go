

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RbaGroupPartition struct {
	Inst struct {

    PartitionName string `json:"partition-name"`
    RoleList []RbaGroupPartitionRoleList `json:"role-list"`
    RuleList []RbaGroupPartitionRuleList `json:"rule-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"partition"`
}


type RbaGroupPartitionRoleList struct {
    Role string `json:"role"`
}


type RbaGroupPartitionRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}

func (p *RbaGroupPartition) GetId() string{
    return p.Inst.PartitionName
}

func (p *RbaGroupPartition) getPath() string{
    return "rba/group/" +p.Inst.Name + "/partition"
}

func (p *RbaGroupPartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroupPartition::Post")
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

func (p *RbaGroupPartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroupPartition::Get")
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
func (p *RbaGroupPartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroupPartition::Put")
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

func (p *RbaGroupPartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaGroupPartition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
