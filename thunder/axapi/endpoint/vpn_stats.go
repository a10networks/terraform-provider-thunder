

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnStats struct {
    
    Error VpnStatsError `json:"error"`
    IkeGatewayList []VpnStatsIkeGatewayList `json:"ike-gateway-list"`
    IkeStatsGlobal VpnStatsIkeStatsGlobal `json:"ike-stats-global"`
    IpsecList []VpnStatsIpsecList `json:"ipsec-list"`
    IpsecSaStatsList []VpnStatsIpsecSaStatsList `json:"ipsec-sa-stats-list"`
    Stats VpnStatsStats `json:"stats"`

}
type DataVpnStats struct {
    DtVpnStats VpnStats `json:"vpn"`
}


type VpnStatsError struct {
    Stats VpnStatsErrorStats `json:"stats"`
}


type VpnStatsErrorStats struct {
    Bad_opcode int `json:"bad_opcode"`
    Bad_sg_write_len int `json:"bad_sg_write_len"`
    Bad_len int `json:"bad_len"`
    Bad_ipsec_protocol int `json:"bad_ipsec_protocol"`
    Bad_ipsec_auth int `json:"bad_ipsec_auth"`
    Bad_ipsec_padding int `json:"bad_ipsec_padding"`
    Bad_ip_version int `json:"bad_ip_version"`
    Bad_auth_type int `json:"bad_auth_type"`
    Bad_encrypt_type int `json:"bad_encrypt_type"`
    Bad_ipsec_spi int `json:"bad_ipsec_spi"`
    Bad_checksum int `json:"bad_checksum"`
    Bad_ipsec_context int `json:"bad_ipsec_context"`
    Bad_ipsec_context_direction int `json:"bad_ipsec_context_direction"`
    Bad_ipsec_context_flag_mismatch int `json:"bad_ipsec_context_flag_mismatch"`
    Ipcomp_payload int `json:"ipcomp_payload"`
    Bad_selector_match int `json:"bad_selector_match"`
    Bad_fragment_size int `json:"bad_fragment_size"`
    Bad_inline_data int `json:"bad_inline_data"`
    Bad_frag_size_configuration int `json:"bad_frag_size_configuration"`
    Dummy_payload int `json:"dummy_payload"`
    Bad_ip_payload_type int `json:"bad_ip_payload_type"`
    Bad_min_frag_size_auth_sha384_512 int `json:"bad_min_frag_size_auth_sha384_512"`
    Bad_esp_next_header int `json:"bad_esp_next_header"`
    Bad_gre_header int `json:"bad_gre_header"`
    Bad_gre_protocol int `json:"bad_gre_protocol"`
    Ipv6_extension_headers_too_big int `json:"ipv6_extension_headers_too_big"`
    Ipv6_hop_by_hop_error int `json:"ipv6_hop_by_hop_error"`
    Error_ipv6_decrypt_rh_segs_left_error int `json:"error_ipv6_decrypt_rh_segs_left_error"`
    Ipv6_rh_length_error int `json:"ipv6_rh_length_error"`
    Ipv6_outbound_rh_copy_addr_error int `json:"ipv6_outbound_rh_copy_addr_error"`
    Error_ipv6_extension_header_bad int `json:"error_IPv6_extension_header_bad"`
    Bad_encrypt_type_ctr_gcm int `json:"bad_encrypt_type_ctr_gcm"`
    Ah_not_supported_with_gcm_gmac_sha2 int `json:"ah_not_supported_with_gcm_gmac_sha2"`
    Tfc_padding_with_prefrag_not_supported int `json:"tfc_padding_with_prefrag_not_supported"`
    Bad_srtp_auth_tag int `json:"bad_srtp_auth_tag"`
    Bad_ipcomp_configuration int `json:"bad_ipcomp_configuration"`
    Dsiv_incorrect_param int `json:"dsiv_incorrect_param"`
    Bad_ipsec_unknown int `json:"bad_ipsec_unknown"`
}


type VpnStatsIkeGatewayList struct {
    Name string `json:"name"`
    Stats VpnStatsIkeGatewayListStats `json:"stats"`
}


type VpnStatsIkeGatewayListStats struct {
    V2InitRekey int `json:"v2-init-rekey"`
    V2RspRekey int `json:"v2-rsp-rekey"`
    V2ChildSaRekey int `json:"v2-child-sa-rekey"`
    V2InInvalid int `json:"v2-in-invalid"`
    V2InInvalidSpi int `json:"v2-in-invalid-spi"`
    V2InInitReq int `json:"v2-in-init-req"`
    V2InInitRsp int `json:"v2-in-init-rsp"`
    V2OutInitReq int `json:"v2-out-init-req"`
    V2OutInitRsp int `json:"v2-out-init-rsp"`
    V2InAuthReq int `json:"v2-in-auth-req"`
    V2InAuthRsp int `json:"v2-in-auth-rsp"`
    V2OutAuthReq int `json:"v2-out-auth-req"`
    V2OutAuthRsp int `json:"v2-out-auth-rsp"`
    V2InCreateChildReq int `json:"v2-in-create-child-req"`
    V2InCreateChildRsp int `json:"v2-in-create-child-rsp"`
    V2OutCreateChildReq int `json:"v2-out-create-child-req"`
    V2OutCreateChildRsp int `json:"v2-out-create-child-rsp"`
    V2InInfoReq int `json:"v2-in-info-req"`
    V2InInfoRsp int `json:"v2-in-info-rsp"`
    V2OutInfoReq int `json:"v2-out-info-req"`
    V2OutInfoRsp int `json:"v2-out-info-rsp"`
    V1InIdProtReq int `json:"v1-in-id-prot-req"`
    V1InIdProtRsp int `json:"v1-in-id-prot-rsp"`
    V1OutIdProtReq int `json:"v1-out-id-prot-req"`
    V1OutIdProtRsp int `json:"v1-out-id-prot-rsp"`
    V1InAuthOnlyReq int `json:"v1-in-auth-only-req"`
    V1InAuthOnlyRsp int `json:"v1-in-auth-only-rsp"`
    V1OutAuthOnlyReq int `json:"v1-out-auth-only-req"`
    V1OutAuthOnlyRsp int `json:"v1-out-auth-only-rsp"`
    V1InAggressiveReq int `json:"v1-in-aggressive-req"`
    V1InAggressiveRsp int `json:"v1-in-aggressive-rsp"`
    V1OutAggressiveReq int `json:"v1-out-aggressive-req"`
    V1OutAggressiveRsp int `json:"v1-out-aggressive-rsp"`
    V1InInfoV1Req int `json:"v1-in-info-v1-req"`
    V1InInfoV1Rsp int `json:"v1-in-info-v1-rsp"`
    V1OutInfoV1Req int `json:"v1-out-info-v1-req"`
    V1OutInfoV1Rsp int `json:"v1-out-info-v1-rsp"`
    V1InTransactionReq int `json:"v1-in-transaction-req"`
    V1InTransactionRsp int `json:"v1-in-transaction-rsp"`
    V1OutTransactionReq int `json:"v1-out-transaction-req"`
    V1OutTransactionRsp int `json:"v1-out-transaction-rsp"`
    V1InQuickModeReq int `json:"v1-in-quick-mode-req"`
    V1InQuickModeRsp int `json:"v1-in-quick-mode-rsp"`
    V1OutQuickModeReq int `json:"v1-out-quick-mode-req"`
    V1OutQuickModeRsp int `json:"v1-out-quick-mode-rsp"`
    V1InNewGroupModeReq int `json:"v1-in-new-group-mode-req"`
    V1InNewGroupModeRsp int `json:"v1-in-new-group-mode-rsp"`
    V1OutNewGroupModeReq int `json:"v1-out-new-group-mode-req"`
    V1OutNewGroupModeRsp int `json:"v1-out-new-group-mode-rsp"`
    V1ChildSaInvalidSpi int `json:"v1-child-sa-invalid-spi"`
    V2ChildSaInvalidSpi int `json:"v2-child-sa-invalid-spi"`
    IkeCurrentVersion int `json:"ike-current-version"`
}


type VpnStatsIkeStatsGlobal struct {
    Stats VpnStatsIkeStatsGlobalStats `json:"stats"`
}


type VpnStatsIkeStatsGlobalStats struct {
    V2InitRekey int `json:"v2-init-rekey"`
    V2RspRekey int `json:"v2-rsp-rekey"`
    V2ChildSaRekey int `json:"v2-child-sa-rekey"`
    V2InInvalid int `json:"v2-in-invalid"`
    V2InInvalidSpi int `json:"v2-in-invalid-spi"`
    V2InInitReq int `json:"v2-in-init-req"`
    V2InInitRsp int `json:"v2-in-init-rsp"`
    V2OutInitReq int `json:"v2-out-init-req"`
    V2OutInitRsp int `json:"v2-out-init-rsp"`
    V2InAuthReq int `json:"v2-in-auth-req"`
    V2InAuthRsp int `json:"v2-in-auth-rsp"`
    V2OutAuthReq int `json:"v2-out-auth-req"`
    V2OutAuthRsp int `json:"v2-out-auth-rsp"`
    V2InCreateChildReq int `json:"v2-in-create-child-req"`
    V2InCreateChildRsp int `json:"v2-in-create-child-rsp"`
    V2OutCreateChildReq int `json:"v2-out-create-child-req"`
    V2OutCreateChildRsp int `json:"v2-out-create-child-rsp"`
    V2InInfoReq int `json:"v2-in-info-req"`
    V2InInfoRsp int `json:"v2-in-info-rsp"`
    V2OutInfoReq int `json:"v2-out-info-req"`
    V2OutInfoRsp int `json:"v2-out-info-rsp"`
    V1InIdProtReq int `json:"v1-in-id-prot-req"`
    V1InIdProtRsp int `json:"v1-in-id-prot-rsp"`
    V1OutIdProtReq int `json:"v1-out-id-prot-req"`
    V1OutIdProtRsp int `json:"v1-out-id-prot-rsp"`
    V1InAuthOnlyReq int `json:"v1-in-auth-only-req"`
    V1InAuthOnlyRsp int `json:"v1-in-auth-only-rsp"`
    V1OutAuthOnlyReq int `json:"v1-out-auth-only-req"`
    V1OutAuthOnlyRsp int `json:"v1-out-auth-only-rsp"`
    V1InAggressiveReq int `json:"v1-in-aggressive-req"`
    V1InAggressiveRsp int `json:"v1-in-aggressive-rsp"`
    V1OutAggressiveReq int `json:"v1-out-aggressive-req"`
    V1OutAggressiveRsp int `json:"v1-out-aggressive-rsp"`
    V1InInfoV1Req int `json:"v1-in-info-v1-req"`
    V1InInfoV1Rsp int `json:"v1-in-info-v1-rsp"`
    V1OutInfoV1Req int `json:"v1-out-info-v1-req"`
    V1OutInfoV1Rsp int `json:"v1-out-info-v1-rsp"`
    V1InTransactionReq int `json:"v1-in-transaction-req"`
    V1InTransactionRsp int `json:"v1-in-transaction-rsp"`
    V1OutTransactionReq int `json:"v1-out-transaction-req"`
    V1OutTransactionRsp int `json:"v1-out-transaction-rsp"`
    V1InQuickModeReq int `json:"v1-in-quick-mode-req"`
    V1InQuickModeRsp int `json:"v1-in-quick-mode-rsp"`
    V1OutQuickModeReq int `json:"v1-out-quick-mode-req"`
    V1OutQuickModeRsp int `json:"v1-out-quick-mode-rsp"`
    V1InNewGroupModeReq int `json:"v1-in-new-group-mode-req"`
    V1InNewGroupModeRsp int `json:"v1-in-new-group-mode-rsp"`
    V1OutNewGroupModeReq int `json:"v1-out-new-group-mode-req"`
    V1OutNewGroupModeRsp int `json:"v1-out-new-group-mode-rsp"`
}


type VpnStatsIpsecList struct {
    Name string `json:"name"`
    Stats VpnStatsIpsecListStats `json:"stats"`
}


type VpnStatsIpsecListStats struct {
    PacketsEncrypted int `json:"packets-encrypted"`
    PacketsDecrypted int `json:"packets-decrypted"`
    AntiReplayNum int `json:"anti-replay-num"`
    RekeyNum int `json:"rekey-num"`
    PacketsErrInactive int `json:"packets-err-inactive"`
    PacketsErrEncryption int `json:"packets-err-encryption"`
    PacketsErrPadCheck int `json:"packets-err-pad-check"`
    PacketsErrPktSanity int `json:"packets-err-pkt-sanity"`
    PacketsErrIcvCheck int `json:"packets-err-icv-check"`
    PacketsErrLifetimeLifebytes int `json:"packets-err-lifetime-lifebytes"`
    BytesEncrypted int `json:"bytes-encrypted"`
    BytesDecrypted int `json:"bytes-decrypted"`
    PrefragSuccess int `json:"prefrag-success"`
    PrefragError int `json:"prefrag-error"`
    CaviumBytesEncrypted int `json:"cavium-bytes-encrypted"`
    CaviumBytesDecrypted int `json:"cavium-bytes-decrypted"`
    CaviumPacketsEncrypted int `json:"cavium-packets-encrypted"`
    CaviumPacketsDecrypted int `json:"cavium-packets-decrypted"`
    QatBytesEncrypted int `json:"qat-bytes-encrypted"`
    QatBytesDecrypted int `json:"qat-bytes-decrypted"`
    QatPacketsEncrypted int `json:"qat-packets-encrypted"`
    QatPacketsDecrypted int `json:"qat-packets-decrypted"`
    TunnelIntfDown int `json:"tunnel-intf-down"`
    PktFailPrepToSend int `json:"pkt-fail-prep-to-send"`
    NoNextHop int `json:"no-next-hop"`
    InvalidTunnelId int `json:"invalid-tunnel-id"`
    NoTunnelFound int `json:"no-tunnel-found"`
    PktFailToSend int `json:"pkt-fail-to-send"`
    FragAfterEncapFragPackets int `json:"frag-after-encap-frag-packets"`
    FragReceived int `json:"frag-received"`
    SequenceNum int `json:"sequence-num"`
    SequenceNumRollover int `json:"sequence-num-rollover"`
    PacketsErrNhCheck int `json:"packets-err-nh-check"`
}


type VpnStatsIpsecSaStatsList struct {
    Stats VpnStatsIpsecSaStatsListStats `json:"stats"`
}


type VpnStatsIpsecSaStatsListStats struct {
    PacketsEncrypted int `json:"packets-encrypted"`
    PacketsDecrypted int `json:"packets-decrypted"`
    AntiReplayNum int `json:"anti-replay-num"`
    RekeyNum int `json:"rekey-num"`
    PacketsErrInactive int `json:"packets-err-inactive"`
    PacketsErrEncryption int `json:"packets-err-encryption"`
    PacketsErrPadCheck int `json:"packets-err-pad-check"`
    PacketsErrPktSanity int `json:"packets-err-pkt-sanity"`
    PacketsErrIcvCheck int `json:"packets-err-icv-check"`
    PacketsErrLifetimeLifebytes int `json:"packets-err-lifetime-lifebytes"`
    BytesEncrypted int `json:"bytes-encrypted"`
    BytesDecrypted int `json:"bytes-decrypted"`
    PrefragSuccess int `json:"prefrag-success"`
    PrefragError int `json:"prefrag-error"`
    CaviumBytesEncrypted int `json:"cavium-bytes-encrypted"`
    CaviumBytesDecrypted int `json:"cavium-bytes-decrypted"`
    CaviumPacketsEncrypted int `json:"cavium-packets-encrypted"`
    CaviumPacketsDecrypted int `json:"cavium-packets-decrypted"`
    QatBytesEncrypted int `json:"qat-bytes-encrypted"`
    QatBytesDecrypted int `json:"qat-bytes-decrypted"`
    QatPacketsEncrypted int `json:"qat-packets-encrypted"`
    QatPacketsDecrypted int `json:"qat-packets-decrypted"`
    TunnelIntfDown int `json:"tunnel-intf-down"`
    PktFailPrepToSend int `json:"pkt-fail-prep-to-send"`
    NoNextHop int `json:"no-next-hop"`
    InvalidTunnelId int `json:"invalid-tunnel-id"`
    NoTunnelFound int `json:"no-tunnel-found"`
    PktFailToSend int `json:"pkt-fail-to-send"`
    FragAfterEncapFragPackets int `json:"frag-after-encap-frag-packets"`
    FragReceived int `json:"frag-received"`
    SequenceNum int `json:"sequence-num"`
    SequenceNumRollover int `json:"sequence-num-rollover"`
    PacketsErrNhCheck int `json:"packets-err-nh-check"`
}


type VpnStatsStats struct {
    Passthrough int `json:"passthrough"`
    HaStandbyDrop int `json:"ha-standby-drop"`
}

func (p *VpnStats) GetId() string{
    return "1"
}

func (p *VpnStats) getPath() string{
    return "vpn/stats"
}

func (p *VpnStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnStats,error) {
logger.Println("VpnStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnStats
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
