

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosNetworkObjectSubNetwork struct {
	Inst struct {

    HostAnomalyThreshold DdosNetworkObjectSubNetworkHostAnomalyThreshold `json:"host-anomaly-threshold"`
    SubNetworkAnomalyThreshold DdosNetworkObjectSubNetworkSubNetworkAnomalyThreshold `json:"sub-network-anomaly-threshold"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ObjectName string 

	} `json:"sub-network"`
}


type DdosNetworkObjectSubNetworkHostAnomalyThreshold struct {
    StaticPktRateThreshold int `json:"static-pkt-rate-threshold"`
    StaticByteRateThreshold int `json:"static-byte-rate-threshold"`
}


type DdosNetworkObjectSubNetworkSubNetworkAnomalyThreshold struct {
    StaticSubNetworkPktRate int `json:"static-sub-network-pkt-rate"`
    StaticSubNetworkByteRate int `json:"static-sub-network-byte-rate"`
}

func (p *DdosNetworkObjectSubNetwork) GetId() string{
    return url.QueryEscape(p.Inst.SubnetIpAddr)
}

func (p *DdosNetworkObjectSubNetwork) getPath() string{
    return "ddos/network-object/" +p.Inst.ObjectName + "/sub-network"
}

func (p *DdosNetworkObjectSubNetwork) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectSubNetwork::Post")
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

func (p *DdosNetworkObjectSubNetwork) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectSubNetwork::Get")
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
func (p *DdosNetworkObjectSubNetwork) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectSubNetwork::Put")
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

func (p *DdosNetworkObjectSubNetwork) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectSubNetwork::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
