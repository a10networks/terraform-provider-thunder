

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutTrafficMapOper struct {
    
    Oper ScaleoutTrafficMapOperOper `json:"oper"`

}
type DataScaleoutTrafficMapOper struct {
    DtScaleoutTrafficMapOper ScaleoutTrafficMapOper `json:"traffic-map"`
}


type ScaleoutTrafficMapOperOper struct {
    VirtualServer string `json:"virtual-server"`
    VirtualPort int `json:"virtual-port"`
    SrcIp string `json:"src-ip"`
    SrcIpv6 string `json:"src-ipv6"`
    MapEntriesListHead []ScaleoutTrafficMapOperOperMapEntriesListHead `json:"map-entries-list-head"`
    TblNum int `json:"tbl-num"`
}


type ScaleoutTrafficMapOperOperMapEntriesListHead struct {
    ServiceType string `json:"service-type"`
    ServiceName string `json:"service-name"`
    UserGrpNum int `json:"user-grp-num"`
    RunningDeviceNum int `json:"running-device-num"`
    MapEntriesList []ScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList `json:"map-entries-list"`
}


type ScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList struct {
    UserGroup int `json:"user-group"`
    CurActive int `json:"cur-active"`
    CurStandby int `json:"cur-standby"`
    NewActive int `json:"new-active"`
    NewStandby int `json:"new-standby"`
}

func (p *ScaleoutTrafficMapOper) GetId() string{
    return "1"
}

func (p *ScaleoutTrafficMapOper) getPath() string{
    return "scaleout/traffic-map/oper"
}

func (p *ScaleoutTrafficMapOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutTrafficMapOper,error) {
logger.Println("ScaleoutTrafficMapOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutTrafficMapOper
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
