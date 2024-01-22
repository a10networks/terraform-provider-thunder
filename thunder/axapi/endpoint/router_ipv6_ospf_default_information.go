

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6OspfDefaultInformation struct {
	Inst struct {

    Always int `json:"always"`
    Metric int `json:"metric"`
    MetricType int `json:"metric-type" dval:"2"`
    Originate int `json:"originate"`
    RouteMap string `json:"route-map"`
    Uuid string `json:"uuid"`
    ProcessId string 

	} `json:"default-information"`
}

func (p *RouterIpv6OspfDefaultInformation) GetId() string{
    return "1"
}

func (p *RouterIpv6OspfDefaultInformation) getPath() string{
    return "router/ipv6/ospf/" +p.Inst.ProcessId + "/default-information"
}

func (p *RouterIpv6OspfDefaultInformation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfDefaultInformation::Post")
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

func (p *RouterIpv6OspfDefaultInformation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfDefaultInformation::Get")
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
func (p *RouterIpv6OspfDefaultInformation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfDefaultInformation::Put")
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

func (p *RouterIpv6OspfDefaultInformation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfDefaultInformation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
