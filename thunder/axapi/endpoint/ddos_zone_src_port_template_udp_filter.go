

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneSrcPortTemplateUdpFilter struct {
	Inst struct {

    ByteOffsetFilter string `json:"byte-offset-filter"`
    UdpFilterAction string `json:"udp-filter-action"`
    UdpFilterActionListName string `json:"udp-filter-action-list-name"`
    UdpFilterInverseMatch int `json:"udp-filter-inverse-match"`
    UdpFilterName string `json:"udp-filter-name"`
    UdpFilterRegex string `json:"udp-filter-regex"`
    UdpFilterSeq int `json:"udp-filter-seq"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"filter"`
}

func (p *DdosZoneSrcPortTemplateUdpFilter) GetId() string{
    return url.QueryEscape(p.Inst.UdpFilterName)
}

func (p *DdosZoneSrcPortTemplateUdpFilter) getPath() string{
    return "ddos/zone-src-port-template/udp/" +p.Inst.Name + "/filter"
}

func (p *DdosZoneSrcPortTemplateUdpFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdpFilter::Post")
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

func (p *DdosZoneSrcPortTemplateUdpFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdpFilter::Get")
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
func (p *DdosZoneSrcPortTemplateUdpFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdpFilter::Put")
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

func (p *DdosZoneSrcPortTemplateUdpFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdpFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
