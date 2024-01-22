

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SflowCollectorIpv6 struct {
	Inst struct {

    Addr string `json:"addr"`
    CustomizedSetting SflowCollectorIpv6CustomizedSetting1401 `json:"customized-setting"`
    Port int `json:"port"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ipv6"`
}


type SflowCollectorIpv6CustomizedSetting1401 struct {
    ExportEnable string `json:"export-enable"`
    PacketSampling int `json:"packet-sampling"`
    CounterPolling int `json:"counter-polling"`
    A10ProprietaryPolling int `json:"a10-proprietary-polling"`
    EventNotification int `json:"event-notification"`
    Uuid string `json:"uuid"`
}

func (p *SflowCollectorIpv6) GetId() string{
    return p.Inst.Addr+"+"+strconv.Itoa(p.Inst.Port)
}

func (p *SflowCollectorIpv6) getPath() string{
    return "sflow/collector/ipv6"
}

func (p *SflowCollectorIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6::Post")
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

func (p *SflowCollectorIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6::Get")
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
func (p *SflowCollectorIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6::Put")
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

func (p *SflowCollectorIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
