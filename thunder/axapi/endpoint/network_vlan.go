

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVlan struct {
	Inst struct {

    Name string `json:"name"`
    SamplingEnable []NetworkVlanSamplingEnable `json:"sampling-enable"`
    SharedVlan int `json:"shared-vlan"`
    TaggedEthList []NetworkVlanTaggedEthList `json:"tagged-eth-list"`
    TaggedTrunkList []NetworkVlanTaggedTrunkList `json:"tagged-trunk-list"`
    TrafficDistributionMode string `json:"traffic-distribution-mode"`
    UntaggedEthList []NetworkVlanUntaggedEthList `json:"untagged-eth-list"`
    UntaggedLif string `json:"untagged-lif"`
    UntaggedTrunkList []NetworkVlanUntaggedTrunkList `json:"untagged-trunk-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Ve int `json:"ve"`
    VlanNum int `json:"vlan-num"`

	} `json:"vlan"`
}


type NetworkVlanSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type NetworkVlanTaggedEthList struct {
    TaggedEthernetStart int `json:"tagged-ethernet-start"`
    TaggedEthernetEnd int `json:"tagged-ethernet-end"`
}


type NetworkVlanTaggedTrunkList struct {
    TaggedTrunkStart int `json:"tagged-trunk-start"`
    TaggedTrunkEnd int `json:"tagged-trunk-end"`
}


type NetworkVlanUntaggedEthList struct {
    UntaggedEthernetStart int `json:"untagged-ethernet-start"`
    UntaggedEthernetEnd int `json:"untagged-ethernet-end"`
}


type NetworkVlanUntaggedTrunkList struct {
    UntaggedTrunkStart int `json:"untagged-trunk-start"`
    UntaggedTrunkEnd int `json:"untagged-trunk-end"`
}

func (p *NetworkVlan) GetId() string{
    return strconv.Itoa(p.Inst.VlanNum)
}

func (p *NetworkVlan) getPath() string{
    return "network/vlan"
}

func (p *NetworkVlan) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlan::Post")
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

func (p *NetworkVlan) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlan::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *NetworkVlan) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlan::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *NetworkVlan) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlan::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
