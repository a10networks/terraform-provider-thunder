

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAPartitionVridStatusOper struct {
    
    Oper VrrpAPartitionVridStatusOperOper `json:"oper"`

}
type DataVrrpAPartitionVridStatusOper struct {
    DtVrrpAPartitionVridStatusOper VrrpAPartitionVridStatusOper `json:"partition-vrid-status"`
}


type VrrpAPartitionVridStatusOperOper struct {
    AllPartitionList []VrrpAPartitionVridStatusOperOperAllPartitionList `json:"all-partition-list"`
}


type VrrpAPartitionVridStatusOperOperAllPartitionList struct {
    Local_device_id int `json:"local_device_ID"`
    PartitionName string `json:"partition-name"`
    Vrid int `json:"vrid"`
    Active_device_id int `json:"active_device_id"`
    Active_priority int `json:"active_priority"`
    Active_weight int `json:"active_weight"`
}

func (p *VrrpAPartitionVridStatusOper) GetId() string{
    return "1"
}

func (p *VrrpAPartitionVridStatusOper) getPath() string{
    return "vrrp-a/partition-vrid-status/oper"
}

func (p *VrrpAPartitionVridStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpAPartitionVridStatusOper,error) {
logger.Println("VrrpAPartitionVridStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpAPartitionVridStatusOper
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
