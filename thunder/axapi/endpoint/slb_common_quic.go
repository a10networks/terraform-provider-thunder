

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbCommonQuic struct {
	Inst struct {

    CidLen int `json:"cid-len" dval:"4"`
    CpuOffset int `json:"cpu-offset" dval:"12"`
    EnableHash int `json:"enable-hash"`
    EnableSignature int `json:"enable-signature"`
    QuicLbOffset int `json:"quic-lb-offset" dval:"8"`
    Signature string `json:"signature"`
    SignatureLen int `json:"signature-len" dval:"3"`
    SignatureOffset int `json:"signature-offset" dval:"4"`
    Uuid string `json:"uuid"`

	} `json:"quic"`
}

func (p *SlbCommonQuic) GetId() string{
    return "1"
}

func (p *SlbCommonQuic) getPath() string{
    return "slb/common/quic"
}

func (p *SlbCommonQuic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonQuic::Post")
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

func (p *SlbCommonQuic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonQuic::Get")
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
func (p *SlbCommonQuic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonQuic::Put")
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

func (p *SlbCommonQuic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonQuic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
