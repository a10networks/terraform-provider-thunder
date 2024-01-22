

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAPeerGroup struct {
	Inst struct {

    Peer VrrpAPeerGroupPeer `json:"peer"`
    Uuid string `json:"uuid"`

	} `json:"peer-group"`
}


type VrrpAPeerGroupPeer struct {
    IpPeerAddressCfg []VrrpAPeerGroupPeerIpPeerAddressCfg `json:"ip-peer-address-cfg"`
    Ipv6PeerAddressCfg []VrrpAPeerGroupPeerIpv6PeerAddressCfg `json:"ipv6-peer-address-cfg"`
}


type VrrpAPeerGroupPeerIpPeerAddressCfg struct {
    IpPeerAddress string `json:"ip-peer-address"`
}


type VrrpAPeerGroupPeerIpv6PeerAddressCfg struct {
    Ipv6PeerAddress string `json:"ipv6-peer-address"`
}

func (p *VrrpAPeerGroup) GetId() string{
    return "1"
}

func (p *VrrpAPeerGroup) getPath() string{
    return "vrrp-a/peer-group"
}

func (p *VrrpAPeerGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPeerGroup::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *VrrpAPeerGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPeerGroup::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *VrrpAPeerGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPeerGroup::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *VrrpAPeerGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPeerGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
