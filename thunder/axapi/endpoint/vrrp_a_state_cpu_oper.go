

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAStateCpuOper struct {
    
    Oper VrrpAStateCpuOperOper `json:"oper"`

}
type DataVrrpAStateCpuOper struct {
    DtVrrpAStateCpuOper VrrpAStateCpuOper `json:"state-cpu"`
}


type VrrpAStateCpuOperOper struct {
    CpuUsage []VrrpAStateCpuOperOperCpuUsage `json:"cpu-usage"`
}


type VrrpAStateCpuOperOperCpuUsage struct {
    CpuId int `json:"cpu-id"`
    MaxSyncMsgPerPacket int `json:"Max-Sync-Msg-Per-Packet"`
    MinSyncMsgPerPacket int `json:"Min-Sync-Msg-Per-Packet"`
    MaxQueryMsgPerPacket int `json:"Max-Query-Msg-Per-Packet"`
    MinQueryMsgPerPacket int `json:"Min-Query-Msg-Per-Packet"`
}

func (p *VrrpAStateCpuOper) GetId() string{
    return "1"
}

func (p *VrrpAStateCpuOper) getPath() string{
    return "vrrp-a/state-cpu/oper"
}

func (p *VrrpAStateCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpAStateCpuOper,error) {
logger.Println("VrrpAStateCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpAStateCpuOper
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
