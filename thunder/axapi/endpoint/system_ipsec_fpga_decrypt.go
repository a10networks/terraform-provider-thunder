

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpsecFpgaDecrypt struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`

	} `json:"fpga-decrypt"`
}

func (p *SystemIpsecFpgaDecrypt) GetId() string{
    return "1"
}

func (p *SystemIpsecFpgaDecrypt) getPath() string{
    return "system/ipsec/fpga-decrypt"
}

func (p *SystemIpsecFpgaDecrypt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsecFpgaDecrypt::Post")
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

func (p *SystemIpsecFpgaDecrypt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsecFpgaDecrypt::Get")
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
func (p *SystemIpsecFpgaDecrypt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsecFpgaDecrypt::Put")
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

func (p *SystemIpsecFpgaDecrypt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpsecFpgaDecrypt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
