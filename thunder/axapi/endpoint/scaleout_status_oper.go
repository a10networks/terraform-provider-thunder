

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutStatusOper struct {
    
    Oper ScaleoutStatusOperOper `json:"oper"`

}
type DataScaleoutStatusOper struct {
    DtScaleoutStatusOper ScaleoutStatusOper `json:"status"`
}


type ScaleoutStatusOperOper struct {
    Db_role string `json:"db_role"`
    Role string `json:"role"`
    DeviceList []ScaleoutStatusOperOperDeviceList `json:"device-list"`
    ClusterMode string `json:"cluster-mode"`
    FollowSharedRedirection int `json:"follow-shared-redirection"`
    FollowSharedSessionSync int `json:"follow-shared-session-sync"`
    L2redirect int `json:"l2redirect"`
    L2redirectValid int `json:"l2redirect-valid"`
    L2redirectOperational int `json:"l2redirect-operational"`
    L2redirectEth int `json:"l2redirect-eth"`
    L2redirectTrunk int `json:"l2redirect-trunk"`
    L2redirectVlan int `json:"l2redirect-vlan"`
    AdvertisedRedirectIpList []ScaleoutStatusOperOperAdvertisedRedirectIpList `json:"advertised-redirect-ip-list"`
    DestRedirectIpList []ScaleoutStatusOperOperDestRedirectIpList `json:"dest-redirect-ip-list"`
    ActiveInterfaceList []ScaleoutStatusOperOperActiveInterfaceList `json:"active-interface-list"`
    AdvertisedSessionSyncIpList []ScaleoutStatusOperOperAdvertisedSessionSyncIpList `json:"advertised-session-sync-ip-list"`
    DestSessionSyncIpList []ScaleoutStatusOperOperDestSessionSyncIpList `json:"dest-session-sync-ip-list"`
    ExcludeInterfaceIpList []ScaleoutStatusOperOperExcludeInterfaceIpList `json:"exclude-interface-ip-list"`
}


type ScaleoutStatusOperOperDeviceList struct {
    Id1 int `json:"id1"`
    State string `json:"state"`
    IsLocal int `json:"is-local"`
    IsMaster int `json:"is-master"`
}


type ScaleoutStatusOperOperAdvertisedRedirectIpList struct {
    Ip string `json:"ip"`
}


type ScaleoutStatusOperOperDestRedirectIpList struct {
    DeviceId int `json:"device-id"`
    Direction string `json:"direction"`
    Ip string `json:"ip"`
}


type ScaleoutStatusOperOperActiveInterfaceList struct {
    Interface string `json:"interface"`
}


type ScaleoutStatusOperOperAdvertisedSessionSyncIpList struct {
    Ip string `json:"ip"`
}


type ScaleoutStatusOperOperDestSessionSyncIpList struct {
    DeviceId int `json:"device-id"`
    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
}


type ScaleoutStatusOperOperExcludeInterfaceIpList struct {
    Ip string `json:"ip"`
}

func (p *ScaleoutStatusOper) GetId() string{
    return "1"
}

func (p *ScaleoutStatusOper) getPath() string{
    return "scaleout/status/oper"
}

func (p *ScaleoutStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutStatusOper,error) {
logger.Println("ScaleoutStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutStatusOper
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
