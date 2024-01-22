

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateEncap struct {
	Inst struct {

    EncapTmplName string `json:"encap-tmpl-name"`
    PreserveSourceIp int `json:"preserve-source-ip"`
    TunnelEncap DdosTemplateEncapTunnelEncap `json:"tunnel-encap"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"encap"`
}


type DdosTemplateEncapTunnelEncap struct {
    IpCfg DdosTemplateEncapTunnelEncapIpCfg `json:"ip-cfg"`
    GreCfg DdosTemplateEncapTunnelEncapGreCfg `json:"gre-cfg"`
}


type DdosTemplateEncapTunnelEncapIpCfg struct {
    IpEncap int `json:"ip-encap"`
    Always DdosTemplateEncapTunnelEncapIpCfgAlways `json:"always"`
}


type DdosTemplateEncapTunnelEncapIpCfgAlways struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
}


type DdosTemplateEncapTunnelEncapGreCfg struct {
    GreEncap int `json:"gre-encap"`
    GreAlways DdosTemplateEncapTunnelEncapGreCfgGreAlways `json:"gre-always"`
}


type DdosTemplateEncapTunnelEncapGreCfgGreAlways struct {
    GreIpv4 string `json:"gre-ipv4"`
    KeyIpv4 string `json:"key-ipv4"`
    GreIpv6 string `json:"gre-ipv6"`
    KeyIpv6 string `json:"key-ipv6"`
}

func (p *DdosTemplateEncap) GetId() string{
    return url.QueryEscape(p.Inst.EncapTmplName)
}

func (p *DdosTemplateEncap) getPath() string{
    return "ddos/template/encap"
}

func (p *DdosTemplateEncap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateEncap::Post")
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

func (p *DdosTemplateEncap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateEncap::Get")
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
func (p *DdosTemplateEncap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateEncap::Put")
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

func (p *DdosTemplateEncap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateEncap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
