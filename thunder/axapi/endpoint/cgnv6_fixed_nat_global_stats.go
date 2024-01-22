

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatGlobalStats struct {
    
    Stats Cgnv6FixedNatGlobalStatsStats `json:"stats"`

}
type DataCgnv6FixedNatGlobalStats struct {
    DtCgnv6FixedNatGlobalStats Cgnv6FixedNatGlobalStats `json:"global"`
}


type Cgnv6FixedNatGlobalStatsStats struct {
    TotalNatInUse int `json:"total-nat-in-use"`
    TotalTcpAllocated int `json:"total-tcp-allocated"`
    TotalTcpFreed int `json:"total-tcp-freed"`
    TotalUdpAllocated int `json:"total-udp-allocated"`
    TotalUdpFreed int `json:"total-udp-freed"`
    TotalIcmpAllocated int `json:"total-icmp-allocated"`
    TotalIcmpFreed int `json:"total-icmp-freed"`
    Nat44DataSessionCreated int `json:"nat44-data-session-created"`
    Nat44DataSessionFreed int `json:"nat44-data-session-freed"`
    Nat64DataSessionCreated int `json:"nat64-data-session-created"`
    Nat64DataSessionFreed int `json:"nat64-data-session-freed"`
    DsliteDataSessionCreated int `json:"dslite-data-session-created"`
    DsliteDataSessionFreed int `json:"dslite-data-session-freed"`
    NatPortUnavailableTcp int `json:"nat-port-unavailable-tcp"`
    NatPortUnavailableUdp int `json:"nat-port-unavailable-udp"`
    NatPortUnavailableIcmp int `json:"nat-port-unavailable-icmp"`
    SessionUserQuotaExceeded int `json:"session-user-quota-exceeded"`
    Nat44TcpFullconeCreated int `json:"nat44-tcp-fullcone-created"`
    Nat44TcpFullconeFreed int `json:"nat44-tcp-fullcone-freed"`
    Nat44UdpFullconeCreated int `json:"nat44-udp-fullcone-created"`
    Nat44UdpFullconeFreed int `json:"nat44-udp-fullcone-freed"`
    Nat44UdpAlgFullconeCreated int `json:"nat44-udp-alg-fullcone-created"`
    Nat44UdpAlgFullconeFreed int `json:"nat44-udp-alg-fullcone-freed"`
    Nat64TcpFullconeCreated int `json:"nat64-tcp-fullcone-created"`
    Nat64TcpFullconeFreed int `json:"nat64-tcp-fullcone-freed"`
    Nat64UdpFullconeCreated int `json:"nat64-udp-fullcone-created"`
    Nat64UdpFullconeFreed int `json:"nat64-udp-fullcone-freed"`
    Nat64UdpAlgFullconeCreated int `json:"nat64-udp-alg-fullcone-created"`
    Nat64UdpAlgFullconeFreed int `json:"nat64-udp-alg-fullcone-freed"`
    DsliteTcpFullconeCreated int `json:"dslite-tcp-fullcone-created"`
    DsliteTcpFullconeFreed int `json:"dslite-tcp-fullcone-freed"`
    DsliteUdpFullconeCreated int `json:"dslite-udp-fullcone-created"`
    DsliteUdpFullconeFreed int `json:"dslite-udp-fullcone-freed"`
    DsliteUdpAlgFullconeCreated int `json:"dslite-udp-alg-fullcone-created"`
    DsliteUdpAlgFullconeFreed int `json:"dslite-udp-alg-fullcone-freed"`
    FullconeFailure int `json:"fullcone-failure"`
    Nat44EimMatch int `json:"nat44-eim-match"`
    Nat64EimMatch int `json:"nat64-eim-match"`
    DsliteEimMatch int `json:"dslite-eim-match"`
    Nat44EifMatch int `json:"nat44-eif-match"`
    Nat64EifMatch int `json:"nat64-eif-match"`
    DsliteEifMatch int `json:"dslite-eif-match"`
    Nat44InboundFiltered int `json:"nat44-inbound-filtered"`
    Nat64InboundFiltered int `json:"nat64-inbound-filtered"`
    DsliteInboundFiltered int `json:"dslite-inbound-filtered"`
    Nat44EifLimitExceeded int `json:"nat44-eif-limit-exceeded"`
    Nat64EifLimitExceeded int `json:"nat64-eif-limit-exceeded"`
    DsliteEifLimitExceeded int `json:"dslite-eif-limit-exceeded"`
    Nat44Hairpin int `json:"nat44-hairpin"`
    Nat64Hairpin int `json:"nat64-hairpin"`
    DsliteHairpin int `json:"dslite-hairpin"`
    StandbyDrop int `json:"standby-drop"`
    FixedNatFullconeSelfHairpinningDrop int `json:"fixed-nat-fullcone-self-hairpinning-drop"`
    SixrdDrop int `json:"sixrd-drop"`
    DestRlistDrop int `json:"dest-rlist-drop"`
    DestRlistPassThrough int `json:"dest-rlist-pass-through"`
    DestRlistSnatDrop int `json:"dest-rlist-snat-drop"`
    ConfigNotFound int `json:"config-not-found"`
    TotalTcpOverloadAcquired int `json:"total-tcp-overload-acquired"`
    TotalUdpOverloadAcquired int `json:"total-udp-overload-acquired"`
    TotalTcpOverloadReleased int `json:"total-tcp-overload-released"`
    TotalUdpOverloadReleased int `json:"total-udp-overload-released"`
    TotalTcpAllocOverload int `json:"total-tcp-alloc-overload"`
    TotalUdpAllocOverload int `json:"total-udp-alloc-overload"`
    TotalTcpFreeOverload int `json:"total-tcp-free-overload"`
    TotalUdpFreeOverload int `json:"total-udp-free-overload"`
    PortOverloadSmpDeleteScheduled int `json:"port-overload-smp-delete-scheduled"`
    PortOverloadFailed int `json:"port-overload-failed"`
    HaSessionUserQuotaExceeded int `json:"ha-session-user-quota-exceeded"`
    Fnat44_fwd_ingress_packets_tcp int `json:"fnat44_fwd_ingress_packets_tcp"`
    Fnat44_fwd_egress_packets_tcp int `json:"fnat44_fwd_egress_packets_tcp"`
    Fnat44_rev_ingress_packets_tcp int `json:"fnat44_rev_ingress_packets_tcp"`
    Fnat44_rev_egress_packets_tcp int `json:"fnat44_rev_egress_packets_tcp"`
    Fnat44_fwd_ingress_bytes_tcp int `json:"fnat44_fwd_ingress_bytes_tcp"`
    Fnat44_fwd_egress_bytes_tcp int `json:"fnat44_fwd_egress_bytes_tcp"`
    Fnat44_rev_ingress_bytes_tcp int `json:"fnat44_rev_ingress_bytes_tcp"`
    Fnat44_rev_egress_bytes_tcp int `json:"fnat44_rev_egress_bytes_tcp"`
    Fnat44_fwd_ingress_packets_udp int `json:"fnat44_fwd_ingress_packets_udp"`
    Fnat44_fwd_egress_packets_udp int `json:"fnat44_fwd_egress_packets_udp"`
    Fnat44_rev_ingress_packets_udp int `json:"fnat44_rev_ingress_packets_udp"`
    Fnat44_rev_egress_packets_udp int `json:"fnat44_rev_egress_packets_udp"`
    Fnat44_fwd_ingress_bytes_udp int `json:"fnat44_fwd_ingress_bytes_udp"`
    Fnat44_fwd_egress_bytes_udp int `json:"fnat44_fwd_egress_bytes_udp"`
    Fnat44_rev_ingress_bytes_udp int `json:"fnat44_rev_ingress_bytes_udp"`
    Fnat44_rev_egress_bytes_udp int `json:"fnat44_rev_egress_bytes_udp"`
    Fnat44_fwd_ingress_packets_icmp int `json:"fnat44_fwd_ingress_packets_icmp"`
    Fnat44_fwd_egress_packets_icmp int `json:"fnat44_fwd_egress_packets_icmp"`
    Fnat44_rev_ingress_packets_icmp int `json:"fnat44_rev_ingress_packets_icmp"`
    Fnat44_rev_egress_packets_icmp int `json:"fnat44_rev_egress_packets_icmp"`
    Fnat44_fwd_ingress_bytes_icmp int `json:"fnat44_fwd_ingress_bytes_icmp"`
    Fnat44_fwd_egress_bytes_icmp int `json:"fnat44_fwd_egress_bytes_icmp"`
    Fnat44_rev_ingress_bytes_icmp int `json:"fnat44_rev_ingress_bytes_icmp"`
    Fnat44_rev_egress_bytes_icmp int `json:"fnat44_rev_egress_bytes_icmp"`
    Fnat44_fwd_ingress_packets_others int `json:"fnat44_fwd_ingress_packets_others"`
    Fnat44_fwd_egress_packets_others int `json:"fnat44_fwd_egress_packets_others"`
    Fnat44_rev_ingress_packets_others int `json:"fnat44_rev_ingress_packets_others"`
    Fnat44_rev_egress_packets_others int `json:"fnat44_rev_egress_packets_others"`
    Fnat44_fwd_ingress_bytes_others int `json:"fnat44_fwd_ingress_bytes_others"`
    Fnat44_fwd_egress_bytes_others int `json:"fnat44_fwd_egress_bytes_others"`
    Fnat44_rev_ingress_bytes_others int `json:"fnat44_rev_ingress_bytes_others"`
    Fnat44_rev_egress_bytes_others int `json:"fnat44_rev_egress_bytes_others"`
    Fnat44_fwd_ingress_pkt_size_range1 int `json:"fnat44_fwd_ingress_pkt_size_range1"`
    Fnat44_fwd_ingress_pkt_size_range2 int `json:"fnat44_fwd_ingress_pkt_size_range2"`
    Fnat44_fwd_ingress_pkt_size_range3 int `json:"fnat44_fwd_ingress_pkt_size_range3"`
    Fnat44_fwd_ingress_pkt_size_range4 int `json:"fnat44_fwd_ingress_pkt_size_range4"`
    Fnat44_fwd_egress_pkt_size_range1 int `json:"fnat44_fwd_egress_pkt_size_range1"`
    Fnat44_fwd_egress_pkt_size_range2 int `json:"fnat44_fwd_egress_pkt_size_range2"`
    Fnat44_fwd_egress_pkt_size_range3 int `json:"fnat44_fwd_egress_pkt_size_range3"`
    Fnat44_fwd_egress_pkt_size_range4 int `json:"fnat44_fwd_egress_pkt_size_range4"`
    Fnat44_rev_ingress_pkt_size_range1 int `json:"fnat44_rev_ingress_pkt_size_range1"`
    Fnat44_rev_ingress_pkt_size_range2 int `json:"fnat44_rev_ingress_pkt_size_range2"`
    Fnat44_rev_ingress_pkt_size_range3 int `json:"fnat44_rev_ingress_pkt_size_range3"`
    Fnat44_rev_ingress_pkt_size_range4 int `json:"fnat44_rev_ingress_pkt_size_range4"`
    Fnat44_rev_egress_pkt_size_range1 int `json:"fnat44_rev_egress_pkt_size_range1"`
    Fnat44_rev_egress_pkt_size_range2 int `json:"fnat44_rev_egress_pkt_size_range2"`
    Fnat44_rev_egress_pkt_size_range3 int `json:"fnat44_rev_egress_pkt_size_range3"`
    Fnat44_rev_egress_pkt_size_range4 int `json:"fnat44_rev_egress_pkt_size_range4"`
    Fnat64_fwd_ingress_packets_tcp int `json:"fnat64_fwd_ingress_packets_tcp"`
    Fnat64_fwd_egress_packets_tcp int `json:"fnat64_fwd_egress_packets_tcp"`
    Fnat64_rev_ingress_packets_tcp int `json:"fnat64_rev_ingress_packets_tcp"`
    Fnat64_rev_egress_packets_tcp int `json:"fnat64_rev_egress_packets_tcp"`
    Fnat64_fwd_ingress_bytes_tcp int `json:"fnat64_fwd_ingress_bytes_tcp"`
    Fnat64_fwd_egress_bytes_tcp int `json:"fnat64_fwd_egress_bytes_tcp"`
    Fnat64_rev_ingress_bytes_tcp int `json:"fnat64_rev_ingress_bytes_tcp"`
    Fnat64_rev_egress_bytes_tcp int `json:"fnat64_rev_egress_bytes_tcp"`
    Fnat64_fwd_ingress_packets_udp int `json:"fnat64_fwd_ingress_packets_udp"`
    Fnat64_fwd_egress_packets_udp int `json:"fnat64_fwd_egress_packets_udp"`
    Fnat64_rev_ingress_packets_udp int `json:"fnat64_rev_ingress_packets_udp"`
    Fnat64_rev_egress_packets_udp int `json:"fnat64_rev_egress_packets_udp"`
    Fnat64_fwd_ingress_bytes_udp int `json:"fnat64_fwd_ingress_bytes_udp"`
    Fnat64_fwd_egress_bytes_udp int `json:"fnat64_fwd_egress_bytes_udp"`
    Fnat64_rev_ingress_bytes_udp int `json:"fnat64_rev_ingress_bytes_udp"`
    Fnat64_rev_egress_bytes_udp int `json:"fnat64_rev_egress_bytes_udp"`
    Fnat64_fwd_ingress_packets_icmp int `json:"fnat64_fwd_ingress_packets_icmp"`
    Fnat64_fwd_egress_packets_icmp int `json:"fnat64_fwd_egress_packets_icmp"`
    Fnat64_rev_ingress_packets_icmp int `json:"fnat64_rev_ingress_packets_icmp"`
    Fnat64_rev_egress_packets_icmp int `json:"fnat64_rev_egress_packets_icmp"`
    Fnat64_fwd_ingress_bytes_icmp int `json:"fnat64_fwd_ingress_bytes_icmp"`
    Fnat64_fwd_egress_bytes_icmp int `json:"fnat64_fwd_egress_bytes_icmp"`
    Fnat64_rev_ingress_bytes_icmp int `json:"fnat64_rev_ingress_bytes_icmp"`
    Fnat64_rev_egress_bytes_icmp int `json:"fnat64_rev_egress_bytes_icmp"`
    Fnat64_fwd_ingress_packets_others int `json:"fnat64_fwd_ingress_packets_others"`
    Fnat64_fwd_egress_packets_others int `json:"fnat64_fwd_egress_packets_others"`
    Fnat64_rev_ingress_packets_others int `json:"fnat64_rev_ingress_packets_others"`
    Fnat64_rev_egress_packets_others int `json:"fnat64_rev_egress_packets_others"`
    Fnat64_fwd_ingress_bytes_others int `json:"fnat64_fwd_ingress_bytes_others"`
    Fnat64_fwd_egress_bytes_others int `json:"fnat64_fwd_egress_bytes_others"`
    Fnat64_rev_ingress_bytes_others int `json:"fnat64_rev_ingress_bytes_others"`
    Fnat64_rev_egress_bytes_others int `json:"fnat64_rev_egress_bytes_others"`
    Fnat64_fwd_ingress_pkt_size_range1 int `json:"fnat64_fwd_ingress_pkt_size_range1"`
    Fnat64_fwd_ingress_pkt_size_range2 int `json:"fnat64_fwd_ingress_pkt_size_range2"`
    Fnat64_fwd_ingress_pkt_size_range3 int `json:"fnat64_fwd_ingress_pkt_size_range3"`
    Fnat64_fwd_ingress_pkt_size_range4 int `json:"fnat64_fwd_ingress_pkt_size_range4"`
    Fnat64_fwd_egress_pkt_size_range1 int `json:"fnat64_fwd_egress_pkt_size_range1"`
    Fnat64_fwd_egress_pkt_size_range2 int `json:"fnat64_fwd_egress_pkt_size_range2"`
    Fnat64_fwd_egress_pkt_size_range3 int `json:"fnat64_fwd_egress_pkt_size_range3"`
    Fnat64_fwd_egress_pkt_size_range4 int `json:"fnat64_fwd_egress_pkt_size_range4"`
    Fnat64_rev_ingress_pkt_size_range1 int `json:"fnat64_rev_ingress_pkt_size_range1"`
    Fnat64_rev_ingress_pkt_size_range2 int `json:"fnat64_rev_ingress_pkt_size_range2"`
    Fnat64_rev_ingress_pkt_size_range3 int `json:"fnat64_rev_ingress_pkt_size_range3"`
    Fnat64_rev_ingress_pkt_size_range4 int `json:"fnat64_rev_ingress_pkt_size_range4"`
    Fnat64_rev_egress_pkt_size_range1 int `json:"fnat64_rev_egress_pkt_size_range1"`
    Fnat64_rev_egress_pkt_size_range2 int `json:"fnat64_rev_egress_pkt_size_range2"`
    Fnat64_rev_egress_pkt_size_range3 int `json:"fnat64_rev_egress_pkt_size_range3"`
    Fnat64_rev_egress_pkt_size_range4 int `json:"fnat64_rev_egress_pkt_size_range4"`
    Fnatdslite_fwd_ingress_packets_tcp int `json:"fnatdslite_fwd_ingress_packets_tcp"`
    Fnatdslite_fwd_egress_packets_tcp int `json:"fnatdslite_fwd_egress_packets_tcp"`
    Fnatdslite_rev_ingress_packets_tcp int `json:"fnatdslite_rev_ingress_packets_tcp"`
    Fnatdslite_rev_egress_packets_tcp int `json:"fnatdslite_rev_egress_packets_tcp"`
    Fnatdslite_fwd_ingress_bytes_tcp int `json:"fnatdslite_fwd_ingress_bytes_tcp"`
    Fnatdslite_fwd_egress_bytes_tcp int `json:"fnatdslite_fwd_egress_bytes_tcp"`
    Fnatdslite_rev_ingress_bytes_tcp int `json:"fnatdslite_rev_ingress_bytes_tcp"`
    Fnatdslite_rev_egress_bytes_tcp int `json:"fnatdslite_rev_egress_bytes_tcp"`
    Fnatdslite_fwd_ingress_packets_udp int `json:"fnatdslite_fwd_ingress_packets_udp"`
    Fnatdslite_fwd_egress_packets_udp int `json:"fnatdslite_fwd_egress_packets_udp"`
    Fnatdslite_rev_ingress_packets_udp int `json:"fnatdslite_rev_ingress_packets_udp"`
    Fnatdslite_rev_egress_packets_udp int `json:"fnatdslite_rev_egress_packets_udp"`
    Fnatdslite_fwd_ingress_bytes_udp int `json:"fnatdslite_fwd_ingress_bytes_udp"`
    Fnatdslite_fwd_egress_bytes_udp int `json:"fnatdslite_fwd_egress_bytes_udp"`
    Fnatdslite_rev_ingress_bytes_udp int `json:"fnatdslite_rev_ingress_bytes_udp"`
    Fnatdslite_rev_egress_bytes_udp int `json:"fnatdslite_rev_egress_bytes_udp"`
    Fnatdslite_fwd_ingress_packets_icmp int `json:"fnatdslite_fwd_ingress_packets_icmp"`
    Fnatdslite_fwd_egress_packets_icmp int `json:"fnatdslite_fwd_egress_packets_icmp"`
    Fnatdslite_rev_ingress_packets_icmp int `json:"fnatdslite_rev_ingress_packets_icmp"`
    Fnatdslite_rev_egress_packets_icmp int `json:"fnatdslite_rev_egress_packets_icmp"`
    Fnatdslite_fwd_ingress_bytes_icmp int `json:"fnatdslite_fwd_ingress_bytes_icmp"`
    Fnatdslite_fwd_egress_bytes_icmp int `json:"fnatdslite_fwd_egress_bytes_icmp"`
    Fnatdslite_rev_ingress_bytes_icmp int `json:"fnatdslite_rev_ingress_bytes_icmp"`
    Fnatdslite_rev_egress_bytes_icmp int `json:"fnatdslite_rev_egress_bytes_icmp"`
    Fnatdslite_fwd_ingress_packets_others int `json:"fnatdslite_fwd_ingress_packets_others"`
    Fnatdslite_fwd_egress_packets_others int `json:"fnatdslite_fwd_egress_packets_others"`
    Fnatdslite_rev_ingress_packets_others int `json:"fnatdslite_rev_ingress_packets_others"`
    Fnatdslite_rev_egress_packets_others int `json:"fnatdslite_rev_egress_packets_others"`
    Fnatdslite_fwd_ingress_bytes_others int `json:"fnatdslite_fwd_ingress_bytes_others"`
    Fnatdslite_fwd_egress_bytes_others int `json:"fnatdslite_fwd_egress_bytes_others"`
    Fnatdslite_rev_ingress_bytes_others int `json:"fnatdslite_rev_ingress_bytes_others"`
    Fnatdslite_rev_egress_bytes_others int `json:"fnatdslite_rev_egress_bytes_others"`
    Fnatdslite_fwd_ingress_pkt_size_range1 int `json:"fnatdslite_fwd_ingress_pkt_size_range1"`
    Fnatdslite_fwd_ingress_pkt_size_range2 int `json:"fnatdslite_fwd_ingress_pkt_size_range2"`
    Fnatdslite_fwd_ingress_pkt_size_range3 int `json:"fnatdslite_fwd_ingress_pkt_size_range3"`
    Fnatdslite_fwd_ingress_pkt_size_range4 int `json:"fnatdslite_fwd_ingress_pkt_size_range4"`
    Fnatdslite_fwd_egress_pkt_size_range1 int `json:"fnatdslite_fwd_egress_pkt_size_range1"`
    Fnatdslite_fwd_egress_pkt_size_range2 int `json:"fnatdslite_fwd_egress_pkt_size_range2"`
    Fnatdslite_fwd_egress_pkt_size_range3 int `json:"fnatdslite_fwd_egress_pkt_size_range3"`
    Fnatdslite_fwd_egress_pkt_size_range4 int `json:"fnatdslite_fwd_egress_pkt_size_range4"`
    Fnatdslite_rev_ingress_pkt_size_range1 int `json:"fnatdslite_rev_ingress_pkt_size_range1"`
    Fnatdslite_rev_ingress_pkt_size_range2 int `json:"fnatdslite_rev_ingress_pkt_size_range2"`
    Fnatdslite_rev_ingress_pkt_size_range3 int `json:"fnatdslite_rev_ingress_pkt_size_range3"`
    Fnatdslite_rev_ingress_pkt_size_range4 int `json:"fnatdslite_rev_ingress_pkt_size_range4"`
    Fnatdslite_rev_egress_pkt_size_range1 int `json:"fnatdslite_rev_egress_pkt_size_range1"`
    Fnatdslite_rev_egress_pkt_size_range2 int `json:"fnatdslite_rev_egress_pkt_size_range2"`
    Fnatdslite_rev_egress_pkt_size_range3 int `json:"fnatdslite_rev_egress_pkt_size_range3"`
    Fnatdslite_rev_egress_pkt_size_range4 int `json:"fnatdslite_rev_egress_pkt_size_range4"`
    ActiveSubscriberAdded int `json:"active-subscriber-added"`
    ActiveSubscriberRemoved int `json:"active-subscriber-removed"`
}

func (p *Cgnv6FixedNatGlobalStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatGlobalStats) getPath() string{
    return "cgnv6/fixed-nat/global/stats"
}

func (p *Cgnv6FixedNatGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatGlobalStats,error) {
logger.Println("Cgnv6FixedNatGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatGlobalStats
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
