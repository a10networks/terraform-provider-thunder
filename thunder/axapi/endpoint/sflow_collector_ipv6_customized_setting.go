

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SflowCollectorIpv6CustomizedSetting struct {
	Inst struct {

    A10ProprietaryPolling int `json:"a10-proprietary-polling"`
    CounterPolling int `json:"counter-polling"`
    EventNotification int `json:"event-notification"`
    ExportEnable string `json:"export-enable"`
    PacketSampling int `json:"packet-sampling"`
    Uuid string `json:"uuid"`
    Port string 
    Addr string 

	} `json:"customized-setting"`
}

func (p *SflowCollectorIpv6CustomizedSetting) GetId() string{
    return "1"
}

func (p *SflowCollectorIpv6CustomizedSetting) getPath() string{
    return "sflow/collector/ipv6/" +p.Inst.Addr + "+" +p.Inst.Port + "/customized-setting"
}

func (p *SflowCollectorIpv6CustomizedSetting) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6CustomizedSetting::Post")
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

func (p *SflowCollectorIpv6CustomizedSetting) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6CustomizedSetting::Get")
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
func (p *SflowCollectorIpv6CustomizedSetting) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6CustomizedSetting::Put")
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

func (p *SflowCollectorIpv6CustomizedSetting) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorIpv6CustomizedSetting::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
