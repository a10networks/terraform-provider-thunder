

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwSessionAging struct {
	Inst struct {

    IcmpIdleTimeout int `json:"icmp-idle-timeout" dval:"2"`
    IpIdleTimeout int `json:"ip-idle-timeout" dval:"30"`
    Name string `json:"name"`
    Tcp FwSessionAgingTcp374 `json:"tcp"`
    Udp FwSessionAgingUdp376 `json:"udp"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"session-aging"`
}


type FwSessionAgingTcp374 struct {
    TcpIdleTimeout int `json:"tcp-idle-timeout" dval:"600"`
    HalfOpenIdleTimeout int `json:"half-open-idle-timeout"`
    HalfCloseIdleTimeout int `json:"half-close-idle-timeout"`
    ForceDeleteTimeout int `json:"force-delete-timeout"`
    ForceDeleteTimeout100ms int `json:"force-delete-timeout-100ms"`
    PortCfg []FwSessionAgingTcpPortCfg375 `json:"port-cfg"`
    Uuid string `json:"uuid"`
}


type FwSessionAgingTcpPortCfg375 struct {
    TcpPort int `json:"tcp-port"`
    TcpIdleTimeout int `json:"tcp-idle-timeout"`
    HalfOpenIdleTimeout int `json:"half-open-idle-timeout"`
    HalfCloseIdleTimeout int `json:"half-close-idle-timeout"`
    ForceDeleteTimeout int `json:"force-delete-timeout"`
    ForceDeleteTimeout100ms int `json:"force-delete-timeout-100ms"`
}


type FwSessionAgingUdp376 struct {
    UdpIdleTimeout int `json:"udp-idle-timeout" dval:"120"`
    PortCfg []FwSessionAgingUdpPortCfg377 `json:"port-cfg"`
    Uuid string `json:"uuid"`
}


type FwSessionAgingUdpPortCfg377 struct {
    UdpPort int `json:"udp-port"`
    UdpIdleTimeout int `json:"udp-idle-timeout" dval:"120"`
}

func (p *FwSessionAging) GetId() string{
    return p.Inst.Name
}

func (p *FwSessionAging) getPath() string{
    return "fw/session-aging"
}

func (p *FwSessionAging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAging::Post")
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

func (p *FwSessionAging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAging::Get")
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
func (p *FwSessionAging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAging::Put")
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

func (p *FwSessionAging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
