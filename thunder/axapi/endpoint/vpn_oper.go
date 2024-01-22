

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnOper struct {
    
    Crl VpnOperCrl `json:"crl"`
    Default VpnOperDefault `json:"default"`
    Errordump VpnOperErrordump `json:"errordump"`
    GroupList VpnOperGroupList `json:"group-list"`
    IkeGatewayList []VpnOperIkeGatewayList `json:"ike-gateway-list"`
    IkeSa VpnOperIkeSa `json:"ike-sa"`
    IkeSaBrief VpnOperIkeSaBrief `json:"ike-sa-brief"`
    IkeSaClients VpnOperIkeSaClients `json:"ike-sa-clients"`
    IkeStatsByGw VpnOperIkeStatsByGw `json:"ike-stats-by-gw"`
    IpsecList []VpnOperIpsecList `json:"ipsec-list"`
    IpsecSa VpnOperIpsecSa `json:"ipsec-sa"`
    IpsecSaByGw VpnOperIpsecSaByGw `json:"ipsec-sa-by-gw"`
    IpsecSaClients VpnOperIpsecSaClients `json:"ipsec-sa-clients"`
    Log VpnOperLog `json:"log"`
    Ocsp VpnOperOcsp `json:"ocsp"`
    Oper VpnOperOper `json:"oper"`

}
type DataVpnOper struct {
    DtVpnOper VpnOper `json:"vpn"`
}


type VpnOperCrl struct {
    Oper VpnOperCrlOper `json:"oper"`
}


type VpnOperCrlOper struct {
    CrlList []VpnOperCrlOperCrlList `json:"crl-list"`
    TotalCrls int `json:"total-crls"`
}


type VpnOperCrlOperCrlList struct {
    Subject string `json:"subject"`
    Issuer string `json:"issuer"`
    Updates string `json:"updates"`
    Serial string `json:"serial"`
    Revoked string `json:"revoked"`
    StorageType string `json:"storage-type"`
}


type VpnOperDefault struct {
    Oper VpnOperDefaultOper `json:"oper"`
}


type VpnOperDefaultOper struct {
    IkeVersion string `json:"ike-version"`
    IkeMode string `json:"ike-mode"`
    IkeDhGroup string `json:"ike-dh-group"`
    IkeAuthMethod string `json:"ike-auth-method"`
    IkeEncryption string `json:"ike-encryption"`
    IkeHash string `json:"ike-hash"`
    IkePriority int `json:"ike-priority"`
    IkeLifetime int `json:"ike-lifetime"`
    IkeNatTraversal string `json:"ike-nat-traversal"`
    IkeLocalAddress string `json:"ike-local-address"`
    IkeRemoteAddress string `json:"ike-remote-address"`
    IkeDpdInterval int `json:"ike-dpd-interval"`
    IpsecMode string `json:"IPsec-mode"`
    IpsecProtocol string `json:"IPsec-protocol"`
    IpsecDhGroup string `json:"IPsec-dh-group"`
    IpsecEncryption string `json:"IPsec-encryption"`
    IpsecHash string `json:"IPsec-hash"`
    IpsecPriority int `json:"IPsec-priority"`
    IpsecLifetime int `json:"IPsec-lifetime"`
    IpsecLifebytes int `json:"IPsec-lifebytes"`
    IpsecTrafficSelector string `json:"IPsec-traffic-selector"`
    IpsecLocalSubnet string `json:"IPsec-local-subnet"`
    IpsecLocalPort int `json:"IPsec-local-port"`
    IpsecLocalProtocol int `json:"IPsec-local-protocol"`
    IpsecRemoteSubnet string `json:"IPsec-remote-subnet"`
    IpsecRemotePort int `json:"IPsec-remote-port"`
    IpsecRemoteProtocol int `json:"IPsec-remote-protocol"`
    IpsecAntiReplayWindow int `json:"IPsec-anti-replay-window"`
}


type VpnOperErrordump struct {
    Oper VpnOperErrordumpOper `json:"oper"`
}


type VpnOperErrordumpOper struct {
    IpsecErrorDumpPath string `json:"IPsec-error-dump-path"`
}


type VpnOperGroupList struct {
    Oper VpnOperGroupListOper `json:"oper"`
}


type VpnOperGroupListOper struct {
    GroupName string `json:"group-name"`
    GroupList []VpnOperGroupListOperGroupList `json:"group-list"`
}


type VpnOperGroupListOperGroupList struct {
    Name string `json:"Name"`
    IpsecSaName string `json:"Ipsec-sa-name"`
    IkeGatewayName string `json:"Ike-gateway-name"`
    Priority int `json:"Priority"`
    Status string `json:"Status"`
    Role string `json:"Role"`
    IsNewGroup int `json:"Is-new-group"`
    GrpMemberCount int `json:"Grp-member-count"`
}


type VpnOperIkeGatewayList struct {
    Name string `json:"name"`
    Oper VpnOperIkeGatewayListOper `json:"oper"`
}


type VpnOperIkeGatewayListOper struct {
    RemoteIpFilter string `json:"remote-ip-filter"`
    RemoteIdFilter string `json:"remote-id-filter"`
    BriefFilter string `json:"brief-filter"`
    SaList []VpnOperIkeGatewayListOperSaList `json:"SA-List"`
}


type VpnOperIkeGatewayListOperSaList struct {
    InitiatorSpi string `json:"Initiator-SPI"`
    ResponderSpi string `json:"Responder-SPI"`
    LocalIp string `json:"Local-IP"`
    RemoteIp string `json:"Remote-IP"`
    Encryption string `json:"Encryption"`
    Hash string `json:"Hash"`
    SignHash string `json:"Sign-hash"`
    Lifetime int `json:"Lifetime"`
    Status string `json:"Status"`
    NatTraversal int `json:"NAT-Traversal"`
    RemoteId string `json:"Remote-ID"`
    DhGroup int `json:"DH-Group"`
    FragmentMessageGenerated int `json:"Fragment-message-generated"`
    FragmentMessageReceived int `json:"Fragment-message-received"`
    FragmentationError int `json:"Fragmentation-error"`
    FragmentReassembleError int `json:"Fragment-reassemble-error"`
}


type VpnOperIkeSa struct {
    Oper VpnOperIkeSaOper `json:"oper"`
}


type VpnOperIkeSaOper struct {
    IkeSaList []VpnOperIkeSaOperIkeSaList `json:"ike-sa-list"`
}


type VpnOperIkeSaOperIkeSaList struct {
    Name string `json:"Name"`
    InitiatorSpi string `json:"Initiator-SPI"`
    ResponderSpi string `json:"Responder-SPI"`
    LocalIp string `json:"Local-IP"`
    RemoteIp string `json:"Remote-IP"`
    Encryption string `json:"Encryption"`
    Hash string `json:"Hash"`
    Lifetime int `json:"Lifetime"`
    Status string `json:"Status"`
    NatTraversal int `json:"NAT-Traversal"`
}


type VpnOperIkeSaBrief struct {
    Oper VpnOperIkeSaBriefOper `json:"oper"`
}


type VpnOperIkeSaBriefOper struct {
    Name string `json:"name"`
    LocalIp string `json:"local-ip"`
    IkeSaBriefRemoteGw []VpnOperIkeSaBriefOperIkeSaBriefRemoteGw `json:"ike-sa-brief-remote-gw"`
}


type VpnOperIkeSaBriefOperIkeSaBriefRemoteGw struct {
    IkeSaBriefRemoteGwIp string `json:"ike-sa-brief-remote-gw-ip"`
    IkeSaBriefRemoteGwId string `json:"ike-sa-brief-remote-gw-id"`
    IkeSaBriefRemoteGwLifetime string `json:"ike-sa-brief-remote-gw-lifetime"`
    IkeSaBriefRemoteGwStatus string `json:"ike-sa-brief-remote-gw-status"`
}


type VpnOperIkeSaClients struct {
    Oper VpnOperIkeSaClientsOper `json:"oper"`
}


type VpnOperIkeSaClientsOper struct {
    Name string `json:"name"`
    IkeSaClientsLocalIp string `json:"ike-sa-clients-local-ip"`
    IkeSaClientsRemoteGw []VpnOperIkeSaClientsOperIkeSaClientsRemoteGw `json:"ike-sa-clients-remote-gw"`
}


type VpnOperIkeSaClientsOperIkeSaClientsRemoteGw struct {
    IkeSaClientsRemoteGwIp string `json:"ike-sa-clients-remote-gw-ip"`
    IkeSaClientsRemoteGwRemoteId string `json:"ike-sa-clients-remote-gw-remote-id"`
    IkeSaClientsRemoteGwUserId string `json:"ike-sa-clients-remote-gw-user-id"`
    IkeSaClientsRemoteGwIdleTime string `json:"ike-sa-clients-remote-gw-idle-time"`
    IkeSaClientsRemoteGwSessionTime string `json:"ike-sa-clients-remote-gw-session-time"`
    IkeSaClientsRemoteGwBytes string `json:"ike-sa-clients-remote-gw-bytes"`
}


type VpnOperIkeStatsByGw struct {
    Oper VpnOperIkeStatsByGwOper `json:"oper"`
}


type VpnOperIkeStatsByGwOper struct {
    GatewayNameFilter string `json:"gateway-name-filter"`
    RemoteIpFilter string `json:"remote-ip-filter"`
    RemoteIdFilter string `json:"remote-id-filter"`
    DisplayAllFilter int `json:"display-all-filter"`
    IkeStatsList []VpnOperIkeStatsByGwOperIkeStatsList `json:"ike-stats-list"`
}


type VpnOperIkeStatsByGwOperIkeStatsList struct {
    Name string `json:"name"`
    RemoteId string `json:"remote-id"`
    RemoteIp string `json:"remote-ip"`
    IkeVersion string `json:"ike-version"`
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
    V2ChildSaInvalidSpi int `json:"v2-child-sa-invalid-spi"`
}


type VpnOperIpsecList struct {
    Name string `json:"name"`
    Oper VpnOperIpsecListOper `json:"oper"`
}


type VpnOperIpsecListOper struct {
    RemoteTsFilter string `json:"remote-ts-filter"`
    RemoteTsV6Filter string `json:"remote-ts-v6-filter"`
    InSpiFilter string `json:"in-spi-filter"`
    OutSpiFilter string `json:"out-spi-filter"`
    SaList []VpnOperIpsecListOperSaList `json:"SA-List"`
}


type VpnOperIpsecListOperSaList struct {
    Status string `json:"Status"`
    SaIndex int `json:"SA-Index"`
    TsProto int `json:"TS-Proto"`
    LocalIp string `json:"Local-IP"`
    LocalPort int `json:"Local-Port"`
    PeerIp string `json:"Peer-IP"`
    PeerPort int `json:"Peer-Port"`
    LocalSpi string `json:"Local-SPI"`
    RemoteSpi string `json:"Remote-SPI"`
    Protocol string `json:"Protocol"`
    Mode string `json:"Mode"`
    EncryptionAlgorithm string `json:"Encryption-Algorithm"`
    HashAlgorithm string `json:"Hash-Algorithm"`
    Lifetime int `json:"Lifetime"`
    Lifebytes string `json:"Lifebytes"`
    DhGroup int `json:"DH-Group"`
    NatTraversal int `json:"NAT-Traversal"`
    AntiReplay string `json:"Anti-Replay"`
    PacketsEncrypted int `json:"packets-encrypted"`
    PacketsDecrypted int `json:"packets-decrypted"`
    AntiReplayNum int `json:"anti-replay-num"`
    RekeyNum int `json:"rekey-num"`
    PacketsErrInactive int `json:"packets-err-inactive"`
    PacketsErrEncryption int `json:"packets-err-encryption"`
    PacketsErrPadCheck int `json:"packets-err-pad-check"`
    PacketsErrPktSanity int `json:"packets-err-pkt-sanity"`
    PacketsErrIcvCheck int `json:"packets-err-icv-check"`
    PacketsErrLifetimeLifebytes VpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes `json:"packets-err-lifetime-lifebytes"`
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
    EnforceTsEncapDrop int `json:"enforce-ts-encap-drop"`
    EnforceTsDecapDrop int `json:"enforce-ts-decap-drop"`
}


type VpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes struct {
}


type VpnOperIpsecSa struct {
    Oper VpnOperIpsecSaOper `json:"oper"`
}


type VpnOperIpsecSaOper struct {
    IpsecSaList []VpnOperIpsecSaOperIpsecSaList `json:"ipsec-sa-list"`
}


type VpnOperIpsecSaOperIpsecSaList struct {
    IpsecSaName string `json:"ipsec-sa-name"`
    IkeGatewayName string `json:"ike-gateway-name"`
    LocalTs string `json:"local-ts"`
    RemoteTs string `json:"remote-ts"`
    InSpi string `json:"in-spi"`
    OutSpi string `json:"out-spi"`
    Protocol string `json:"protocol"`
    Mode string `json:"mode"`
    Encryption string `json:"encryption"`
    Hash string `json:"hash"`
    Lifetime int `json:"lifetime"`
    Lifebytes string `json:"lifebytes"`
}


type VpnOperIpsecSaByGw struct {
    Oper VpnOperIpsecSaByGwOper `json:"oper"`
}


type VpnOperIpsecSaByGwOper struct {
    IkeGatewayName string `json:"ike-gateway-name"`
    LocalIp string `json:"local-ip"`
    PeerIp string `json:"peer-ip"`
    IpsecSaList []VpnOperIpsecSaByGwOperIpsecSaList `json:"ipsec-sa-list"`
}


type VpnOperIpsecSaByGwOperIpsecSaList struct {
    IpsecSaName string `json:"ipsec-sa-name"`
    LocalTs string `json:"local-ts"`
    RemoteTs string `json:"remote-ts"`
    InSpi string `json:"in-spi"`
    OutSpi string `json:"out-spi"`
    Protocol string `json:"protocol"`
    Mode string `json:"mode"`
    Encryption string `json:"encryption"`
    Hash string `json:"hash"`
    Lifetime int `json:"lifetime"`
    Lifebytes string `json:"lifebytes"`
}


type VpnOperIpsecSaClients struct {
    Oper VpnOperIpsecSaClientsOper `json:"oper"`
}


type VpnOperIpsecSaClientsOper struct {
    IpsecClients []VpnOperIpsecSaClientsOperIpsecClients `json:"ipsec-clients"`
}


type VpnOperIpsecSaClientsOperIpsecClients struct {
    IpsecClientsIp string `json:"ipsec-clients-ip"`
    SaList []VpnOperIpsecSaClientsOperIpsecClientsSaList `json:"sa-list"`
}


type VpnOperIpsecSaClientsOperIpsecClientsSaList struct {
    Name string `json:"name"`
    LocalTs string `json:"local-ts"`
    InSpi string `json:"in-spi"`
    OutSpi string `json:"out-spi"`
    Lifetime string `json:"lifetime"`
    Lifebytes string `json:"lifebytes"`
}


type VpnOperLog struct {
    Oper VpnOperLogOper `json:"oper"`
}


type VpnOperLogOper struct {
    VpnLogList []VpnOperLogOperVpnLogList `json:"vpn-log-list"`
    VpnLogOffset int `json:"vpn-log-offset"`
    VpnLogOver int `json:"vpn-log-over"`
    Follow int `json:"follow"`
    FromStart int `json:"from-start"`
    NumLines int `json:"num-lines"`
}


type VpnOperLogOperVpnLogList struct {
    VpnLogData string `json:"vpn-log-data"`
}


type VpnOperOcsp struct {
    Oper VpnOperOcspOper `json:"oper"`
}


type VpnOperOcspOper struct {
    OcspList []VpnOperOcspOperOcspList `json:"ocsp-list"`
    TotalOcsps int `json:"total-ocsps"`
}


type VpnOperOcspOperOcspList struct {
    Subject string `json:"subject"`
    Issuer string `json:"issuer"`
    Validity string `json:"validity"`
    CertificateStatus string `json:"certificate-status"`
}


type VpnOperOper struct {
    IkeGatewayTotal int `json:"IKE-Gateway-total"`
    IpsecTotal int `json:"IPsec-total"`
    IkeSaTotal int `json:"IKE-SA-total"`
    IpsecSaTotal int `json:"IPsec-SA-total"`
    IpsecMode string `json:"IPsec-mode"`
    NumHardwareDevices int `json:"Num-hardware-devices"`
    CryptoCoresTotal int `json:"Crypto-cores-total"`
    CryptoCoresAssignedToIpsec int `json:"Crypto-cores-assigned-to-IPsec"`
    CryptoMem int `json:"Crypto-mem"`
    AllPartitionList []VpnOperOperAllPartitionList `json:"all-partition-list"`
    AllPartitions int `json:"all-partitions"`
    Shared int `json:"shared"`
    SpecificPartition string `json:"specific-partition"`
}


type VpnOperOperAllPartitionList struct {
    IkeGatewayTotal int `json:"IKE-Gateway-total"`
    IpsecTotal int `json:"IPsec-total"`
    IkeSaTotal int `json:"IKE-SA-total"`
    IpsecSaTotal int `json:"IPsec-SA-total"`
    IpsecStateless int `json:"IPsec-stateless"`
    IpsecMode string `json:"IPsec-mode"`
    IpsecHardwareType string `json:"IPsec-hardware-type"`
    NumHardwareDevices int `json:"Num-hardware-devices"`
    IkeHardwareAccelerate string `json:"IKE-hardware-accelerate"`
    CryptoCoresTotal int `json:"Crypto-cores-total"`
    CryptoCoresAssignedToIpsec int `json:"Crypto-cores-assigned-to-IPsec"`
    CryptoMem int `json:"Crypto-mem"`
    CryptoHwErr int `json:"Crypto-hw-err"`
    CryptoHwErrReqAllocFail int `json:"Crypto-hw-err-req-alloc-fail"`
    CryptoHwErrEnqueueFail int `json:"Crypto-hw-err-enqueue-fail"`
    CryptoHwErrSgBuffAllocFail int `json:"Crypto-hw-err-sg-buff-alloc-fail"`
    CryptoHwErrBadPointer int `json:"Crypto-hw-err-bad-pointer"`
    CryptoHwErrBadCtxPointer int `json:"Crypto-hw-err-bad-ctx-pointer"`
    CryptoHwErrReqError int `json:"Crypto-hw-err-req-error"`
    CryptoHwErrStateError int `json:"Crypto-hw-err-state-error"`
    CryptoHwErrState string `json:"Crypto-hw-err-state"`
    CryptoHwErrTimeOut int `json:"Crypto-hw-err-time-out"`
    CryptoHwErrTimeOutState int `json:"Crypto-hw-err-time-out-state"`
    CryptoHwErrBuffAllocError int `json:"Crypto-hw-err-buff-alloc-error"`
    PassthroughTotal int `json:"passthrough-total"`
    VpnList []VpnOperOperAllPartitionListVpnList `json:"vpn-list"`
    StandbyDrop int `json:"standby-drop"`
    PartitionName string `json:"partition-name"`
    SignedAuthFlag int `json:"Signed-auth-flag"`
}


type VpnOperOperAllPartitionListVpnList struct {
    Passthrough int `json:"passthrough"`
    CpuId int `json:"cpu-id"`
}

func (p *VpnOper) GetId() string{
    return "1"
}

func (p *VpnOper) getPath() string{
    return "vpn/oper"
}

func (p *VpnOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnOper,error) {
logger.Println("VpnOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnOper
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
