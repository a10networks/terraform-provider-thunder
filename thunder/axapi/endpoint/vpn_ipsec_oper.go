

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIpsecOper struct {
    
    Name string `json:"name"`
    Oper VpnIpsecOperOper `json:"oper"`

}
type DataVpnIpsecOper struct {
    DtVpnIpsecOper VpnIpsecOper `json:"ipsec"`
}


type VpnIpsecOperOper struct {
    RemoteTsFilter string `json:"remote-ts-filter"`
    RemoteTsV6Filter string `json:"remote-ts-v6-filter"`
    InSpiFilter string `json:"in-spi-filter"`
    OutSpiFilter string `json:"out-spi-filter"`
    SaList []VpnIpsecOperOperSaList `json:"SA-List"`
}


type VpnIpsecOperOperSaList struct {
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
    PacketsErrLifetimeLifebytes VpnIpsecOperOperSaListPacketsErrLifetimeLifebytes `json:"packets-err-lifetime-lifebytes"`
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


type VpnIpsecOperOperSaListPacketsErrLifetimeLifebytes struct {
}

func (p *VpnIpsecOper) GetId() string{
    return "1"
}

func (p *VpnIpsecOper) getPath() string{
    
    return "vpn/ipsec/"+p.Name+"/oper"
}

func (p *VpnIpsecOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIpsecOper,error) {
logger.Println("VpnIpsecOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIpsecOper
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
