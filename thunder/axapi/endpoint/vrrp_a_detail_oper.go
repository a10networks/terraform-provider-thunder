

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpADetailOper struct {
    
    Oper VrrpADetailOperOper `json:"oper"`

}
type DataVrrpADetailOper struct {
    DtVrrpADetailOper VrrpADetailOper `json:"detail"`
}


type VrrpADetailOperOper struct {
    Dup_id_rcv int `json:"dup_id_rcv"`
    Bad_group_rcv int `json:"bad_group_rcv"`
    Set_id_mismatch int `json:"set_id_mismatch"`
    Vrrp_version_mismatch int `json:"vrrp_version_mismatch"`
    Time_inaccurate_count int `json:"time_inaccurate_count"`
    Ip_pools_exceeded int `json:"ip_pools_exceeded"`
    Ip6_pools_exceeded int `json:"ip6_pools_exceeded"`
    Err_port int `json:"err_port"`
    Err_devid int `json:"err_devid"`
    Err_parid int `json:"err_parid"`
    Lock_try int `json:"lock_try"`
    L2_no_route int `json:"l2_no_route"`
    Peer_info_list []VrrpADetailOperOperPeer_info_list `json:"peer_info_list"`
    Local_info_list []VrrpADetailOperOperLocal_info_list `json:"local_info_list"`
}


type VrrpADetailOperOperPeer_info_list struct {
    Peer_ip string `json:"peer_ip"`
    Peer_pkt_recv int `json:"peer_pkt_recv"`
    Peer_list []VrrpADetailOperOperPeer_info_listPeer_list `json:"peer_list"`
}


type VrrpADetailOperOperPeer_info_listPeer_list struct {
    Peer_id int `json:"peer_id"`
    Vrid int `json:"vrid"`
    Missing_heartbeat int `json:"missing_heartbeat"`
    Peer_port_list []VrrpADetailOperOperPeer_info_listPeer_listPeer_port_list `json:"peer_port_list"`
}


type VrrpADetailOperOperPeer_info_listPeer_listPeer_port_list struct {
    Eth int `json:"eth"`
    Vrrp_pkt_recv int `json:"vrrp_pkt_recv"`
    Eth_miss int `json:"eth_miss"`
}


type VrrpADetailOperOperLocal_info_list struct {
    Switch_to_active int `json:"switch_to_active"`
    Switch_to_standby int `json:"switch_to_standby"`
    Vrid int `json:"vrid"`
    Vrrp_pkt_send int `json:"vrrp_pkt_send"`
    Local_eth_send_list []VrrpADetailOperOperLocal_info_listLocal_eth_send_list `json:"local_eth_send_list"`
}


type VrrpADetailOperOperLocal_info_listLocal_eth_send_list struct {
    Eth_port int `json:"eth_port"`
    Eth_pkt_send int `json:"eth_pkt_send"`
}

func (p *VrrpADetailOper) GetId() string{
    return "1"
}

func (p *VrrpADetailOper) getPath() string{
    return "vrrp-a/detail/oper"
}

func (p *VrrpADetailOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpADetailOper,error) {
logger.Println("VrrpADetailOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpADetailOper
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
