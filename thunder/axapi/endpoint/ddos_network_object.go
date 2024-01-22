

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosNetworkObject struct {
	Inst struct {

    AnomalyDetectionTrigger string `json:"anomaly-detection-trigger" dval:"all"`
    HistogramEnable int `json:"histogram-enable"`
    HostAnomalyThreshold DdosNetworkObjectHostAnomalyThreshold `json:"host-anomaly-threshold"`
    Ip []DdosNetworkObjectIp `json:"ip"`
    Ipv6 []DdosNetworkObjectIpv6 `json:"ipv6"`
    NetworkObjectAnomalyThreshold DdosNetworkObjectNetworkObjectAnomalyThreshold `json:"network-object-anomaly-threshold"`
    Notification DdosNetworkObjectNotification282 `json:"notification"`
    ObjectName string `json:"object-name"`
    OperationalMode string `json:"operational-mode" dval:"learning"`
    RelativeAutoBreakDownThreshold DdosNetworkObjectRelativeAutoBreakDownThreshold `json:"relative-auto-break-down-threshold"`
    SamplingEnable []DdosNetworkObjectSamplingEnable `json:"sampling-enable"`
    ServiceBreakDownThresholdLocal DdosNetworkObjectServiceBreakDownThresholdLocal `json:"service-break-down-threshold-local"`
    ServiceDiscovery string `json:"service-discovery"`
    StaticAutoBreakDownThreshold DdosNetworkObjectStaticAutoBreakDownThreshold `json:"static-auto-break-down-threshold"`
    SubNetworkList []DdosNetworkObjectSubNetworkList `json:"sub-network-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"network-object"`
}


type DdosNetworkObjectHostAnomalyThreshold struct {
    HostPktRate int `json:"host-pkt-rate"`
    HostByteRate int `json:"host-byte-rate"`
}


type DdosNetworkObjectIp struct {
    SubnetIpAddr string `json:"subnet-ip-addr"`
}


type DdosNetworkObjectIpv6 struct {
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
}


type DdosNetworkObjectNetworkObjectAnomalyThreshold struct {
    NetworkObjectPktRate int `json:"network-object-pkt-rate"`
    NetworkObjectByteRate int `json:"network-object-byte-rate"`
}


type DdosNetworkObjectNotification282 struct {
    Configuration string `json:"configuration"`
    Notification []DdosNetworkObjectNotificationNotification283 `json:"notification"`
    Uuid string `json:"uuid"`
}


type DdosNetworkObjectNotificationNotification283 struct {
    NotificationTemplateName string `json:"notification-template-name"`
}


type DdosNetworkObjectRelativeAutoBreakDownThreshold struct {
    NetworkPercentage int `json:"network-percentage"`
    Permil int `json:"permil"`
}


type DdosNetworkObjectSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosNetworkObjectServiceBreakDownThresholdLocal struct {
    SvcPercentage int `json:"svc-percentage"`
}


type DdosNetworkObjectStaticAutoBreakDownThreshold struct {
    NetworkPktRate int `json:"network-pkt-rate"`
}


type DdosNetworkObjectSubNetworkList struct {
    SubnetIpAddr string `json:"subnet-ip-addr"`
    HostAnomalyThreshold DdosNetworkObjectSubNetworkListHostAnomalyThreshold `json:"host-anomaly-threshold"`
    SubNetworkAnomalyThreshold DdosNetworkObjectSubNetworkListSubNetworkAnomalyThreshold `json:"sub-network-anomaly-threshold"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosNetworkObjectSubNetworkListHostAnomalyThreshold struct {
    StaticPktRateThreshold int `json:"static-pkt-rate-threshold"`
    StaticByteRateThreshold int `json:"static-byte-rate-threshold"`
}


type DdosNetworkObjectSubNetworkListSubNetworkAnomalyThreshold struct {
    StaticSubNetworkPktRate int `json:"static-sub-network-pkt-rate"`
    StaticSubNetworkByteRate int `json:"static-sub-network-byte-rate"`
}

func (p *DdosNetworkObject) GetId() string{
    return url.QueryEscape(p.Inst.ObjectName)
}

func (p *DdosNetworkObject) getPath() string{
    return "ddos/network-object"
}

func (p *DdosNetworkObject) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObject::Post")
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

func (p *DdosNetworkObject) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObject::Get")
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
func (p *DdosNetworkObject) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObject::Put")
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

func (p *DdosNetworkObject) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObject::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
