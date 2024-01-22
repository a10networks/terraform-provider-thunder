

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemMemoryBlockDebug struct {
	Inst struct {

    AssertBlock int `json:"assert-block" dval:"65536"`
    FirstBlk int `json:"first-blk" dval:"8192"`
    FourthBlk int `json:"fourth-blk" dval:"65536"`
    PktdumpBlock int `json:"pktdump-block"`
    SecondBlk int `json:"second-blk" dval:"16384"`
    ThirdBlk int `json:"third-blk" dval:"32768"`
    Uuid string `json:"uuid"`

	} `json:"memory-block-debug"`
}

func (p *SystemMemoryBlockDebug) GetId() string{
    return "1"
}

func (p *SystemMemoryBlockDebug) getPath() string{
    return "system/memory-block-debug"
}

func (p *SystemMemoryBlockDebug) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMemoryBlockDebug::Post")
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

func (p *SystemMemoryBlockDebug) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMemoryBlockDebug::Get")
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
func (p *SystemMemoryBlockDebug) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMemoryBlockDebug::Put")
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

func (p *SystemMemoryBlockDebug) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMemoryBlockDebug::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
