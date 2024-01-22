

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL4SyncStats struct {
    
    Stats DdosL4SyncStatsStats `json:"stats"`

}
type DataDdosL4SyncStats struct {
    DtDdosL4SyncStats DdosL4SyncStats `json:"l4-sync"`
}


type DdosL4SyncStatsStats struct {
    Sync_hello_rcv int `json:"sync_hello_rcv"`
    Sync_hello_sent int `json:"sync_hello_sent"`
    Sync_dst_sent int `json:"sync_dst_sent"`
    Sync_dst_rcv int `json:"sync_dst_rcv"`
    Sync_src_sent int `json:"sync_src_sent"`
    Sync_src_rcv int `json:"sync_src_rcv"`
    Sync_src_dst_sent int `json:"sync_src_dst_sent"`
    Sync_src_dst_rcv int `json:"sync_src_dst_rcv"`
    Sync_szp_sent int `json:"sync_szp_sent"`
    Sync_szp_rcv int `json:"sync_szp_rcv"`
    Sync_sent_no_peer int `json:"sync_sent_no_peer"`
    Sync_sent_fail int `json:"sync_sent_fail"`
    Sync_rcv_hello_unk_peer int `json:"sync_rcv_hello_unk_peer"`
    Sync_src_dst_no_dst_drop int `json:"sync_src_dst_no_dst_drop"`
    Sync_szp_no_dst_drop int `json:"sync_szp_no_dst_drop"`
    Sync_rcv_entry_create_fail int `json:"sync_rcv_entry_create_fail"`
    Sync_rcv_entry_conflict_static int `json:"sync_rcv_entry_conflict_static"`
    Sync_rcv_unk_msg int `json:"sync_rcv_unk_msg"`
    Sync_rcv_fail int `json:"sync_rcv_fail"`
    Sync_rcv_hello_unk_subtype int `json:"sync_rcv_hello_unk_subtype"`
    Sync_rcv_entry_unk_subtype int `json:"sync_rcv_entry_unk_subtype"`
    Sync_rcv_entry_unk int `json:"sync_rcv_entry_unk"`
    Sync_src_udp_wl_sent int `json:"sync_src_udp_wl_sent"`
    Sync_src_tcp_wl_sent int `json:"sync_src_tcp_wl_sent"`
    Sync_src_icmp_wl_sent int `json:"sync_src_icmp_wl_sent"`
    Sync_src_other_wl_sent int `json:"sync_src_other_wl_sent"`
    Sync_src_dns_udp_auth_sent int `json:"sync_src_dns_udp_auth_sent"`
    Sync_src_dns_tcp_auth_sent int `json:"sync_src_dns_tcp_auth_sent"`
    Sync_src_bl_sent int `json:"sync_src_bl_sent"`
    Sync_src_dst_udp_wl_sent int `json:"sync_src_dst_udp_wl_sent"`
    Sync_src_dst_tcp_wl_sent int `json:"sync_src_dst_tcp_wl_sent"`
    Sync_src_dst_icmp_wl_sent int `json:"sync_src_dst_icmp_wl_sent"`
    Sync_src_dst_other_wl_sent int `json:"sync_src_dst_other_wl_sent"`
    Sync_src_dst_dns_udp_auth_sent int `json:"sync_src_dst_dns_udp_auth_sent"`
    Sync_src_dst_dns_tcp_auth_sent int `json:"sync_src_dst_dns_tcp_auth_sent"`
    Sync_src_dst_bl_sent int `json:"sync_src_dst_bl_sent"`
    Sync_szp_tcp_auth_sent int `json:"sync_szp_tcp_auth_sent"`
    Sync_szp_udp_auth_sent int `json:"sync_szp_udp_auth_sent"`
    Sync_szp_icmp_auth_sent int `json:"sync_szp_icmp_auth_sent"`
    Sync_szp_other_auth_sent int `json:"sync_szp_other_auth_sent"`
    Sync_szp_dns_udp_auth_sent int `json:"sync_szp_dns_udp_auth_sent"`
    Sync_szp_dns_tcp_auth_sent int `json:"sync_szp_dns_tcp_auth_sent"`
    Sync_szp_bl_sent int `json:"sync_szp_bl_sent"`
    Sync_src_udp_wl_rcvd int `json:"sync_src_udp_wl_rcvd"`
    Sync_src_tcp_wl_rcvd int `json:"sync_src_tcp_wl_rcvd"`
    Sync_src_icmp_wl_rcvd int `json:"sync_src_icmp_wl_rcvd"`
    Sync_src_other_wl_rcvd int `json:"sync_src_other_wl_rcvd"`
    Sync_src_dns_udp_auth_rcvd int `json:"sync_src_dns_udp_auth_rcvd"`
    Sync_src_dns_tcp_auth_rcvd int `json:"sync_src_dns_tcp_auth_rcvd"`
    Sync_src_bl_rcvd int `json:"sync_src_bl_rcvd"`
    Sync_src_dst_udp_wl_rcvd int `json:"sync_src_dst_udp_wl_rcvd"`
    Sync_src_dst_tcp_wl_rcvd int `json:"sync_src_dst_tcp_wl_rcvd"`
    Sync_src_dst_icmp_wl_rcvd int `json:"sync_src_dst_icmp_wl_rcvd"`
    Sync_src_dst_other_wl_rcvd int `json:"sync_src_dst_other_wl_rcvd"`
    Sync_src_dst_dns_udp_auth_rcvd int `json:"sync_src_dst_dns_udp_auth_rcvd"`
    Sync_src_dst_dns_tcp_auth_rcvd int `json:"sync_src_dst_dns_tcp_auth_rcvd"`
    Sync_src_dst_bl_rcvd int `json:"sync_src_dst_bl_rcvd"`
    Sync_szp_tcp_auth_rcvd int `json:"sync_szp_tcp_auth_rcvd"`
    Sync_szp_udp_auth_rcvd int `json:"sync_szp_udp_auth_rcvd"`
    Sync_szp_icmp_auth_rcvd int `json:"sync_szp_icmp_auth_rcvd"`
    Sync_szp_other_auth_rcvd int `json:"sync_szp_other_auth_rcvd"`
    Sync_szp_dns_udp_auth_rcvd int `json:"sync_szp_dns_udp_auth_rcvd"`
    Sync_szp_dns_tcp_auth_rcvd int `json:"sync_szp_dns_tcp_auth_rcvd"`
    Sync_szp_bl_rcvd int `json:"sync_szp_bl_rcvd"`
}

func (p *DdosL4SyncStats) GetId() string{
    return "1"
}

func (p *DdosL4SyncStats) getPath() string{
    return "ddos/l4-sync/stats"
}

func (p *DdosL4SyncStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL4SyncStats,error) {
logger.Println("DdosL4SyncStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL4SyncStats
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
