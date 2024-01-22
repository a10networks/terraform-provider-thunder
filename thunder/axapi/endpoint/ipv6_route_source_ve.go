

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Ipv6RouteSourceVe struct {
	Inst struct {

    NexthopIp string `json:"nexthop-ip"`
    Uuid string `json:"uuid"`
    VeNum int `json:"ve-num"`

	} `json:"ve"`
}

func (p *Ipv6RouteSourceVe) GetId() string{
    return strconv.Itoa(p.Inst.VeNum)+"+"+p.Inst.NexthopIp
}

func (p *Ipv6RouteSourceVe) getPath() string{
    return "ipv6/route/source/ve"
}

func (p *Ipv6RouteSourceVe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteSourceVe::Post")
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

func (p *Ipv6RouteSourceVe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteSourceVe::Get")
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
func (p *Ipv6RouteSourceVe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteSourceVe::Put")
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

func (p *Ipv6RouteSourceVe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteSourceVe::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
