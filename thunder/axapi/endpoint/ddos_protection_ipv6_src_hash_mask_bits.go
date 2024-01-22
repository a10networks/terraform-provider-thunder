

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosProtectionIpv6SrcHashMaskBits struct {
	Inst struct {

    MaskBitOffset1 int `json:"mask-bit-offset-1"`
    MaskBitOffset2 int `json:"mask-bit-offset-2"`
    MaskBitOffset3 int `json:"mask-bit-offset-3"`
    MaskBitOffset4 int `json:"mask-bit-offset-4"`
    MaskBitOffset5 int `json:"mask-bit-offset-5"`
    Uuid string `json:"uuid"`

	} `json:"ipv6-src-hash-mask-bits"`
}

func (p *DdosProtectionIpv6SrcHashMaskBits) GetId() string{
    return "1"
}

func (p *DdosProtectionIpv6SrcHashMaskBits) getPath() string{
    return "ddos/protection/ipv6-src-hash-mask-bits"
}

func (p *DdosProtectionIpv6SrcHashMaskBits) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionIpv6SrcHashMaskBits::Post")
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

func (p *DdosProtectionIpv6SrcHashMaskBits) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionIpv6SrcHashMaskBits::Get")
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
func (p *DdosProtectionIpv6SrcHashMaskBits) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionIpv6SrcHashMaskBits::Put")
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

func (p *DdosProtectionIpv6SrcHashMaskBits) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionIpv6SrcHashMaskBits::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
