

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneSrcPortTemplateUdp struct {
	Inst struct {

    FilterList []DdosZoneSrcPortTemplateUdpFilterList `json:"filter-list"`
    FilterMatchType string `json:"filter-match-type" dval:"default"`
    MaxPayloadSizeCfg DdosZoneSrcPortTemplateUdpMaxPayloadSizeCfg `json:"max-payload-size-cfg"`
    MinPayloadSizeCfg DdosZoneSrcPortTemplateUdpMinPayloadSizeCfg `json:"min-payload-size-cfg"`
    Name string `json:"name"`
    NtpMonlistCfg DdosZoneSrcPortTemplateUdpNtpMonlistCfg `json:"ntp-monlist-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"udp"`
}


type DdosZoneSrcPortTemplateUdpFilterList struct {
    UdpFilterName string `json:"udp-filter-name"`
    UdpFilterSeq int `json:"udp-filter-seq"`
    UdpFilterRegex string `json:"udp-filter-regex"`
    UdpFilterInverseMatch int `json:"udp-filter-inverse-match"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    UdpFilterActionListName string `json:"udp-filter-action-list-name"`
    UdpFilterAction string `json:"udp-filter-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneSrcPortTemplateUdpMaxPayloadSizeCfg struct {
    MaxPayloadSize int `json:"max-payload-size"`
    MaxPayloadSizeActionListName string `json:"max-payload-size-action-list-name"`
    MaxPayloadSizeAction string `json:"max-payload-size-action"`
}


type DdosZoneSrcPortTemplateUdpMinPayloadSizeCfg struct {
    MinPayloadSize int `json:"min-payload-size"`
    MinPayloadSizeActionListName string `json:"min-payload-size-action-list-name"`
    MinPayloadSizeAction string `json:"min-payload-size-action"`
}


type DdosZoneSrcPortTemplateUdpNtpMonlistCfg struct {
    NtpMonlist int `json:"ntp-monlist"`
    NtpMonlistActionListName string `json:"ntp-monlist-action-list-name"`
    NtpMonlistAction string `json:"ntp-monlist-action"`
}

func (p *DdosZoneSrcPortTemplateUdp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosZoneSrcPortTemplateUdp) getPath() string{
    return "ddos/zone-src-port-template/udp"
}

func (p *DdosZoneSrcPortTemplateUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdp::Post")
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

func (p *DdosZoneSrcPortTemplateUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdp::Get")
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
func (p *DdosZoneSrcPortTemplateUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdp::Put")
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

func (p *DdosZoneSrcPortTemplateUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
