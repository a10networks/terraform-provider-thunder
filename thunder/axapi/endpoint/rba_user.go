

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RbaUser struct {
	Inst struct {

    Name string `json:"name"`
    PartitionList []RbaUserPartitionList `json:"partition-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"user"`
}


type RbaUserPartitionList struct {
    PartitionName string `json:"partition-name"`
    RoleList []RbaUserPartitionListRoleList `json:"role-list"`
    RuleList []RbaUserPartitionListRuleList `json:"rule-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type RbaUserPartitionListRoleList struct {
    Role string `json:"role"`
}


type RbaUserPartitionListRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}

func (p *RbaUser) GetId() string{
    return p.Inst.Name
}

func (p *RbaUser) getPath() string{
    return "rba/user"
}

func (p *RbaUser) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUser::Post")
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

func (p *RbaUser) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUser::Get")
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
func (p *RbaUser) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUser::Put")
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

func (p *RbaUser) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaUser::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
