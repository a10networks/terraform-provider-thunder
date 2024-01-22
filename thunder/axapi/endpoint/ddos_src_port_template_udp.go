

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcPortTemplateUdp struct {
	Inst struct {

    DropNtpMonlist int `json:"drop-ntp-monlist"`
    FilterList []DdosSrcPortTemplateUdpFilterList `json:"filter-list"`
    MaxPayloadSize int `json:"max-payload-size"`
    MinPayloadSize int `json:"min-payload-size"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"udp"`
}


type DdosSrcPortTemplateUdpFilterList struct {
    UdpFilterSeq int `json:"udp-filter-seq"`
    UdpFilterRegex string `json:"udp-filter-regex"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    UdpFilterUnmatched int `json:"udp-filter-unmatched"`
    UdpFilterAction string `json:"udp-filter-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosSrcPortTemplateUdp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosSrcPortTemplateUdp) getPath() string{
    return "ddos/src-port-template/udp"
}

func (p *DdosSrcPortTemplateUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateUdp::Post")
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

func (p *DdosSrcPortTemplateUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateUdp::Get")
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
func (p *DdosSrcPortTemplateUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateUdp::Put")
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

func (p *DdosSrcPortTemplateUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
