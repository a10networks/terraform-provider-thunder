

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpRouteSourceVe struct {
	Inst struct {

    NexthopIp string `json:"nexthop-ip"`
    Uuid string `json:"uuid"`
    VeNum int `json:"ve-num"`

	} `json:"ve"`
}

func (p *IpRouteSourceVe) GetId() string{
    return strconv.Itoa(p.Inst.VeNum)+"+"+p.Inst.NexthopIp
}

func (p *IpRouteSourceVe) getPath() string{
    return "ip/route/source/ve"
}

func (p *IpRouteSourceVe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteSourceVe::Post")
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

func (p *IpRouteSourceVe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteSourceVe::Get")
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
func (p *IpRouteSourceVe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteSourceVe::Put")
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

func (p *IpRouteSourceVe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteSourceVe::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
