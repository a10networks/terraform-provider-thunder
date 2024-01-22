

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDiameter struct {
	Inst struct {

    AvpCode int `json:"avp-code"`
    AvpList []SlbTemplateDiameterAvpList `json:"avp-list"`
    AvpString string `json:"avp-string"`
    CustomizeCea int `json:"customize-cea"`
    DwrTime int `json:"dwr-time" dval:"100"`
    DwrUpRetry int `json:"dwr-up-retry" dval:"3"`
    ForwardToLatestServer int `json:"forward-to-latest-server"`
    ForwardUnknownSessionId int `json:"forward-unknown-session-id"`
    IdleTimeout int `json:"idle-timeout" dval:"5"`
    LoadBalanceOnSessionId int `json:"load-balance-on-session-id"`
    MessageCodeList []SlbTemplateDiameterMessageCodeList `json:"message-code-list"`
    MultipleOriginHost int `json:"multiple-origin-host"`
    Name string `json:"name"`
    OriginHost SlbTemplateDiameterOriginHost1416 `json:"origin-host"`
    OriginRealm string `json:"origin-realm"`
    ProductName string `json:"product-name"`
    RelaxedOriginHost int `json:"relaxed-origin-host"`
    ServiceGroupName string `json:"service-group-name"`
    SessionAge int `json:"session-age" dval:"10"`
    TerminateOnCcaT int `json:"terminate-on-cca-t"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VendorId int `json:"vendor-id"`

	} `json:"diameter"`
}


type SlbTemplateDiameterAvpList struct {
    Avp int `json:"avp"`
    Int32 int `json:"int32"`
    Int64 int `json:"int64"`
    String string `json:"string"`
    Mandatory int `json:"mandatory"`
}


type SlbTemplateDiameterMessageCodeList struct {
    MessageCode int `json:"message-code"`
}


type SlbTemplateDiameterOriginHost1416 struct {
    OriginHostName string `json:"origin-host-name"`
    Uuid string `json:"uuid"`
}

func (p *SlbTemplateDiameter) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDiameter) getPath() string{
    return "slb/template/diameter"
}

func (p *SlbTemplateDiameter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDiameter::Post")
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

func (p *SlbTemplateDiameter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDiameter::Get")
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
func (p *SlbTemplateDiameter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDiameter::Put")
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

func (p *SlbTemplateDiameter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDiameter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
