

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsChassisSlotSummaryOper struct {
    
    Oper VcsChassisSlotSummaryOperOper `json:"oper"`

}
type DataVcsChassisSlotSummaryOper struct {
    DtVcsChassisSlotSummaryOper VcsChassisSlotSummaryOper `json:"slot-summary"`
}


type VcsChassisSlotSummaryOperOper struct {
    VcsEnabled string `json:"vcs-enabled"`
    ChassisId int `json:"chassis-id"`
    MulticastAddr string `json:"multicast-addr"`
    MulticastPort int `json:"multicast-port"`
    Version string `json:"version"`
    VmasterMaintenanceInterval int `json:"vmaster-maintenance-interval"`
    VmasterMaintenanceLeft int `json:"vmaster-maintenance-left"`
    VcsHandshakeCompletedList []VcsChassisSlotSummaryOperOperVcsHandshakeCompletedList `json:"vcs-handshake-completed-list"`
    FloatingIpv4List []VcsChassisSlotSummaryOperOperFloatingIpv4List `json:"floating-ipv4-list"`
    FloatingIpv6List []VcsChassisSlotSummaryOperOperFloatingIpv6List `json:"floating-ipv6-list"`
    MemberList []VcsChassisSlotSummaryOperOperMemberList `json:"member-list"`
}


type VcsChassisSlotSummaryOperOperVcsHandshakeCompletedList struct {
    VcsHandshakeCompletedId int `json:"vcs-handshake-completed-id"`
    VcsHandshakeCompleted int `json:"vcs-handshake-completed"`
}


type VcsChassisSlotSummaryOperOperFloatingIpv4List struct {
    FloatingIpv4 string `json:"floating-ipv4"`
    FloatingIpv4Mask string `json:"floating-ipv4-mask"`
}


type VcsChassisSlotSummaryOperOperFloatingIpv6List struct {
    FloatingIpv6 string `json:"floating-ipv6"`
    FloatingIpv6Prefix int `json:"floating-ipv6-prefix"`
}


type VcsChassisSlotSummaryOperOperMemberList struct {
    Id1 int `json:"id1"`
    State string `json:"state"`
    Priority int `json:"priority"`
    IpList []VcsChassisSlotSummaryOperOperMemberListIpList `json:"ip-list"`
    Port int `json:"port"`
    Location string `json:"location"`
}


type VcsChassisSlotSummaryOperOperMemberListIpList struct {
    Ip string `json:"ip"`
}

func (p *VcsChassisSlotSummaryOper) GetId() string{
    return "1"
}

func (p *VcsChassisSlotSummaryOper) getPath() string{
    return "vcs-chassis/slot-summary/oper"
}

func (p *VcsChassisSlotSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVcsChassisSlotSummaryOper,error) {
logger.Println("VcsChassisSlotSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVcsChassisSlotSummaryOper
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
