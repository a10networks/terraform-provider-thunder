

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateLinkProbe struct {
	Inst struct {

    Destination SlbTemplateLinkProbeDestination1449 `json:"destination"`
    Disable int `json:"disable"`
    ExpectedStatusCode string `json:"expected-status-code" dval:"200"`
    Name string `json:"name"`
    ProbeInterval int `json:"probe-interval" dval:"5"`
    ProbesPerTest int `json:"probes-per-test" dval:"5"`
    RttMethod string `json:"rtt-method" dval:"http-rtt"`
    SelectionRule string `json:"selection-rule" dval:"fastest-link-always"`
    TestInterval int `json:"test-interval" dval:"60"`
    ThresholdValue int `json:"threshold-value"`
    Url string `json:"url" dval:"/"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"link-probe"`
}


type SlbTemplateLinkProbeDestination1449 struct {
    Hostname string `json:"hostname"`
    ResolveAs string `json:"resolve-as"`
    StaticIpv4Addr string `json:"static-ipv4-addr"`
    StaticIpv6Addr string `json:"static-ipv6-addr"`
    Uuid string `json:"uuid"`
}

func (p *SlbTemplateLinkProbe) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateLinkProbe) getPath() string{
    return "slb/template/link-probe"
}

func (p *SlbTemplateLinkProbe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbe::Post")
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

func (p *SlbTemplateLinkProbe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbe::Get")
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
func (p *SlbTemplateLinkProbe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbe::Put")
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

func (p *SlbTemplateLinkProbe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbe::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
