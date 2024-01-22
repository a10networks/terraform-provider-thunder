

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugTrafficMapOper struct {
    
    Oper ScaleoutDebugTrafficMapOperOper `json:"oper"`

}
type DataScaleoutDebugTrafficMapOper struct {
    DtScaleoutDebugTrafficMapOper ScaleoutDebugTrafficMapOper `json:"traffic-map"`
}


type ScaleoutDebugTrafficMapOperOper struct {
    DeviceGroupList []ScaleoutDebugTrafficMapOperOperDeviceGroupList `json:"device-group-list"`
}


type ScaleoutDebugTrafficMapOperOperDeviceGroupList struct {
    Rc int `json:"rc"`
    Cmd string `json:"cmd"`
    Buffer_len int `json:"buffer_len"`
    BucketsList []ScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList `json:"buckets-list"`
}


type ScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList struct {
    User_group int `json:"user_group"`
    Active_device int `json:"active_device"`
    Standby_device int `json:"standby_device"`
}

func (p *ScaleoutDebugTrafficMapOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugTrafficMapOper) getPath() string{
    return "scaleout/debug/traffic-map/oper"
}

func (p *ScaleoutDebugTrafficMapOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugTrafficMapOper,error) {
logger.Println("ScaleoutDebugTrafficMapOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugTrafficMapOper
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
