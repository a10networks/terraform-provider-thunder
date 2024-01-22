

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplatePcp struct {
	Inst struct {

    AllowThirdPartyFromLan int `json:"allow-third-party-from-lan"`
    AllowThirdPartyFromWan int `json:"allow-third-party-from-wan"`
    Announce int `json:"announce"`
    CheckClientNonce int `json:"check-client-nonce"`
    DisableMapFilter int `json:"disable-map-filter"`
    Map int `json:"map"`
    Maximum int `json:"maximum" dval:"1440"`
    Minimum int `json:"minimum" dval:"2"`
    Name string `json:"name"`
    PcpServerPort int `json:"pcp-server-port" dval:"5351"`
    Peer int `json:"peer"`
    SourceIp string `json:"source-ip"`
    SourceIpv6 string `json:"source-ipv6"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"pcp"`
}

func (p *Cgnv6TemplatePcp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6TemplatePcp) getPath() string{
    return "cgnv6/template/pcp"
}

func (p *Cgnv6TemplatePcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePcp::Post")
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

func (p *Cgnv6TemplatePcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePcp::Get")
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
func (p *Cgnv6TemplatePcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePcp::Put")
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

func (p *Cgnv6TemplatePcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
