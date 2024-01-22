

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6OspfArea struct {
	Inst struct {

    AreaIpv4 string `json:"area-ipv4"`
    AreaNum int `json:"area-num"`
    DefaultCost int `json:"default-cost" dval:"1"`
    NoSummary int `json:"no-summary"`
    RangeList []RouterIpv6OspfAreaRangeList `json:"range-list"`
    Stub int `json:"stub"`
    Uuid string `json:"uuid"`
    VirtualLinkList []RouterIpv6OspfAreaVirtualLinkList `json:"virtual-link-list"`
    ProcessId string 

	} `json:"area"`
}


type RouterIpv6OspfAreaRangeList struct {
    Value string `json:"value"`
    Option string `json:"option"`
}


type RouterIpv6OspfAreaVirtualLinkList struct {
    Value string `json:"value"`
    DeadInterval int `json:"dead-interval"`
    Bfd int `json:"bfd"`
    HelloInterval int `json:"hello-interval"`
    RetransmitInterval int `json:"retransmit-interval"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}

func (p *RouterIpv6OspfArea) GetId() string{
    return p.Inst.AreaIpv4+"+"+strconv.Itoa(p.Inst.AreaNum)
}

func (p *RouterIpv6OspfArea) getPath() string{
    return "router/ipv6/ospf/" +p.Inst.ProcessId + "/area"
}

func (p *RouterIpv6OspfArea) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfArea::Post")
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

func (p *RouterIpv6OspfArea) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfArea::Get")
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
func (p *RouterIpv6OspfArea) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfArea::Put")
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

func (p *RouterIpv6OspfArea) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfArea::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
