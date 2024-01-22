

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwSessionAgingTcp struct {
	Inst struct {

    ForceDeleteTimeout int `json:"force-delete-timeout"`
    ForceDeleteTimeout100ms int `json:"force-delete-timeout-100ms"`
    HalfCloseIdleTimeout int `json:"half-close-idle-timeout"`
    HalfOpenIdleTimeout int `json:"half-open-idle-timeout"`
    PortCfg []FwSessionAgingTcpPortCfg `json:"port-cfg"`
    TcpIdleTimeout int `json:"tcp-idle-timeout" dval:"600"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"tcp"`
}


type FwSessionAgingTcpPortCfg struct {
    TcpPort int `json:"tcp-port"`
    TcpIdleTimeout int `json:"tcp-idle-timeout"`
    HalfOpenIdleTimeout int `json:"half-open-idle-timeout"`
    HalfCloseIdleTimeout int `json:"half-close-idle-timeout"`
    ForceDeleteTimeout int `json:"force-delete-timeout"`
    ForceDeleteTimeout100ms int `json:"force-delete-timeout-100ms"`
}

func (p *FwSessionAgingTcp) GetId() string{
    return "1"
}

func (p *FwSessionAgingTcp) getPath() string{
    return "fw/session-aging/" +p.Inst.Name + "/tcp"
}

func (p *FwSessionAgingTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingTcp::Post")
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

func (p *FwSessionAgingTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingTcp::Get")
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
func (p *FwSessionAgingTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingTcp::Put")
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

func (p *FwSessionAgingTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
