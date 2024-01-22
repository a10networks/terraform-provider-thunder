

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PartitionAllOper struct {
    
    Oper PartitionAllOperOper `json:"oper"`

}
type DataPartitionAllOper struct {
    DtPartitionAllOper PartitionAllOper `json:"partition-all"`
}


type PartitionAllOperOper struct {
    PartitionList []PartitionAllOperOperPartitionList `json:"partition-list"`
    ActivePartitionCount int `json:"active-partition-count"`
    Manageable int `json:"manageable"`
}


type PartitionAllOperOperPartitionList struct {
    PartitionName string `json:"partition-name"`
    PartitionId int `json:"partition-id"`
    PartitionType string `json:"partition-type"`
    ParentL3v string `json:"parent-l3v"`
    AppType string `json:"app-Type"`
    AdminCount int `json:"admin-Count"`
    Status string `json:"status"`
}

func (p *PartitionAllOper) GetId() string{
    return "1"
}

func (p *PartitionAllOper) getPath() string{
    return "partition-all/oper"
}

func (p *PartitionAllOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataPartitionAllOper,error) {
logger.Println("PartitionAllOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataPartitionAllOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
