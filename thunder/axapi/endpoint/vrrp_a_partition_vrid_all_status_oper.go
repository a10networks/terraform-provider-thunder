

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAPartitionVridAllStatusOper struct {
    
    Oper VrrpAPartitionVridAllStatusOperOper `json:"oper"`

}
type DataVrrpAPartitionVridAllStatusOper struct {
    DtVrrpAPartitionVridAllStatusOper VrrpAPartitionVridAllStatusOper `json:"partition-vrid-all-status"`
}


type VrrpAPartitionVridAllStatusOperOper struct {
    AllPartitionList []VrrpAPartitionVridAllStatusOperOperAllPartitionList `json:"all-partition-list"`
}


type VrrpAPartitionVridAllStatusOperOperAllPartitionList struct {
    Local_device_id int `json:"local_device_ID"`
    PartitionName string `json:"partition-name"`
    Vrid int `json:"vrid"`
    Active_device_id int `json:"active_device_id"`
    Active_priority int `json:"active_priority"`
    Active_weight int `json:"active_weight"`
    Standby_device_id int `json:"standby_device_id"`
    Standby_priority int `json:"standby_priority"`
    Standby_weight int `json:"standby_weight"`
}

func (p *VrrpAPartitionVridAllStatusOper) GetId() string{
    return "1"
}

func (p *VrrpAPartitionVridAllStatusOper) getPath() string{
    return "vrrp-a/partition-vrid-all-status/oper"
}

func (p *VrrpAPartitionVridAllStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpAPartitionVridAllStatusOper,error) {
logger.Println("VrrpAPartitionVridAllStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpAPartitionVridAllStatusOper
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
