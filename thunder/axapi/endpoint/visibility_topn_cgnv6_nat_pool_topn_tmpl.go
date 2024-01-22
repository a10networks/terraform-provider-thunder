

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityTopnCgnv6NatPoolTopnTmpl struct {
	Inst struct {

    Interval string `json:"interval"`
    Metrics VisibilityTopnCgnv6NatPoolTopnTmplMetrics3131 `json:"metrics"`
    Name string `json:"name"`
    TopnSize int `json:"topn-size"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"cgnv6-nat-pool-topn-tmpl"`
}


type VisibilityTopnCgnv6NatPoolTopnTmplMetrics3131 struct {
    UdpTotal int `json:"udp-total"`
    TcpTotal int `json:"tcp-total"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityTopnCgnv6NatPoolTopnTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityTopnCgnv6NatPoolTopnTmpl) getPath() string{
    return "visibility/topn/cgnv6-nat-pool-topn-tmpl"
}

func (p *VisibilityTopnCgnv6NatPoolTopnTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnTmpl::Post")
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

func (p *VisibilityTopnCgnv6NatPoolTopnTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnTmpl::Get")
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
func (p *VisibilityTopnCgnv6NatPoolTopnTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnTmpl::Put")
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

func (p *VisibilityTopnCgnv6NatPoolTopnTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnCgnv6NatPoolTopnTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
