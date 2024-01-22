

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateTcp struct {
	Inst struct {

    AliveIfActive int `json:"alive-if-active"`
    DelSessionOnServerDown int `json:"del-session-on-server-down"`
    Disable int `json:"disable"`
    Down int `json:"down"`
    ForceDeleteTimeout int `json:"force-delete-timeout"`
    ForceDeleteTimeout100ms int `json:"force-delete-timeout-100ms"`
    HalfCloseIdleTimeout int `json:"half-close-idle-timeout"`
    HalfOpenIdleTimeout int `json:"half-open-idle-timeout"`
    IdleTimeout int `json:"idle-timeout" dval:"120"`
    InitialWindowSize int `json:"initial-window-size"`
    InsertClientIp int `json:"insert-client-ip"`
    LanFastAck int `json:"lan-fast-ack"`
    Logging string `json:"logging"`
    Name string `json:"name"`
    ProxyHeader SlbTemplateTcpProxyHeader `json:"proxy-header"`
    Qos int `json:"qos"`
    ReSelectIfServerDown int `json:"re-select-if-server-down"`
    ResetFollowFin int `json:"reset-follow-fin"`
    ResetFwd int `json:"reset-fwd"`
    ResetRev int `json:"reset-rev"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type SlbTemplateTcpProxyHeader struct {
    ProxyHeaderAction string `json:"proxy-header-action"`
    ProxyHeaderVersion string `json:"proxy-header-version"`
}

func (p *SlbTemplateTcp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateTcp) getPath() string{
    return "slb/template/tcp"
}

func (p *SlbTemplateTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcp::Post")
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

func (p *SlbTemplateTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcp::Get")
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
func (p *SlbTemplateTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcp::Put")
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

func (p *SlbTemplateTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
