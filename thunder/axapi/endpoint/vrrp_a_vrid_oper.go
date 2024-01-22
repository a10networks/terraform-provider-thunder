

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAVridOper struct {
    
    Oper VrrpAVridOperOper `json:"oper"`
    VridVal int `json:"vrid-val"`

}
type DataVrrpAVridOper struct {
    DtVrrpAVridOper VrrpAVridOper `json:"vrid"`
}


type VrrpAVridOperOper struct {
    Unit int `json:"unit"`
    State string `json:"state"`
    Weight int `json:"weight"`
    Priority int `json:"priority"`
    Force_standby string `json:"force_standby"`
    Init_status string `json:"init_status"`
    Became_active string `json:"became_active"`
    Vrid_lead string `json:"vrid_lead"`
    Active_standby_local string `json:"active_standby_local"`
    PeerList []VrrpAVridOperOperPeerList `json:"peer-list"`
}


type VrrpAVridOperOperPeerList struct {
    Peer_state string `json:"peer_state"`
    Peer_unit int `json:"peer_unit"`
    Peer_weight int `json:"peer_weight"`
    Peer_priority int `json:"peer_priority"`
    Active_standby_peer string `json:"active_standby_peer"`
    Peer_vrid int `json:"peer_vrid"`
}

func (p *VrrpAVridOper) GetId() string{
    return "1"
}

func (p *VrrpAVridOper) getPath() string{
    
    return "vrrp-a/vrid/" +strconv.Itoa(p.VridVal)+"/oper"
}

func (p *VrrpAVridOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpAVridOper,error) {
logger.Println("VrrpAVridOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpAVridOper
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
