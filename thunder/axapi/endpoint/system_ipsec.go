

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpsec struct {
	Inst struct {

    CryptoCore int `json:"crypto-core"`
    CryptoMem int `json:"crypto-mem"`
    FpgaDecrypt SystemIpsecFpgaDecrypt1592 `json:"fpga-decrypt"`
    PacketRoundRobin int `json:"packet-round-robin"`
    Qat int `json:"QAT"`
    Uuid string `json:"uuid"`

	} `json:"ipsec"`
}


type SystemIpsecFpgaDecrypt1592 struct {
    Action string `json:"action" dval:"disable"`
}

func (p *SystemIpsec) GetId() string{
    return "1"
}

func (p *SystemIpsec) getPath() string{
    return "system/ipsec"
}

func (p *SystemIpsec) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsec::Post")
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

func (p *SystemIpsec) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsec::Get")
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
func (p *SystemIpsec) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsec::Put")
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

func (p *SystemIpsec) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsec::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
