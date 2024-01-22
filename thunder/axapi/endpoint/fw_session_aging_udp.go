

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwSessionAgingUdp struct {
	Inst struct {

    PortCfg []FwSessionAgingUdpPortCfg `json:"port-cfg"`
    UdpIdleTimeout int `json:"udp-idle-timeout" dval:"120"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"udp"`
}


type FwSessionAgingUdpPortCfg struct {
    UdpPort int `json:"udp-port"`
    UdpIdleTimeout int `json:"udp-idle-timeout" dval:"120"`
}

func (p *FwSessionAgingUdp) GetId() string{
    return "1"
}

func (p *FwSessionAgingUdp) getPath() string{
    return "fw/session-aging/" +p.Inst.Name + "/udp"
}

func (p *FwSessionAgingUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingUdp::Post")
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

func (p *FwSessionAgingUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingUdp::Get")
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
func (p *FwSessionAgingUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingUdp::Put")
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

func (p *FwSessionAgingUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwSessionAgingUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
