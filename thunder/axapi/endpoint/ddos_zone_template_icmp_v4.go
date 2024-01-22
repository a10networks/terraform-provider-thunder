

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateIcmpV4 struct {
	Inst struct {

    FilterList []DdosZoneTemplateIcmpV4FilterList `json:"filter-list"`
    FilterMatchType string `json:"filter-match-type" dval:"default"`
    IcmpTmplName string `json:"icmp-tmpl-name"`
    TypeList []DdosZoneTemplateIcmpV4TypeList `json:"type-list"`
    TypeOther DdosZoneTemplateIcmpV4TypeOther309 `json:"type-other"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"icmp-v4"`
}


type DdosZoneTemplateIcmpV4FilterList struct {
    IcmpFilterName string `json:"icmp-filter-name"`
    IcmpFilterSeq int `json:"icmp-filter-seq"`
    IcmpFilterRegex string `json:"icmp-filter-regex"`
    IcmpFilterInverseMatch int `json:"icmp-filter-inverse-match"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    IcmpFilterActionListName string `json:"icmp-filter-action-list-name"`
    IcmpFilterAction string `json:"icmp-filter-action" dval:"drop"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneTemplateIcmpV4TypeList struct {
    TypeNumber int `json:"type-number"`
    IcmpTypeActionListName string `json:"icmp-type-action-list-name"`
    IcmpTypeAction string `json:"icmp-type-action"`
    V4SrcRateCfg DdosZoneTemplateIcmpV4TypeListV4SrcRateCfg `json:"v4-src-rate-cfg"`
    V4SrcCodeCfg []DdosZoneTemplateIcmpV4TypeListV4SrcCodeCfg `json:"v4-src-code-cfg"`
    SrcCodeOtherRate int `json:"src-code-other-rate"`
    SrcCodeOtherRateActionListName string `json:"src-code-other-rate-action-list-name"`
    SrcCodeOtherRateAction string `json:"src-code-other-rate-action"`
    V4DstRateCfg DdosZoneTemplateIcmpV4TypeListV4DstRateCfg `json:"v4-dst-rate-cfg"`
    V4DstCodeCfg []DdosZoneTemplateIcmpV4TypeListV4DstCodeCfg `json:"v4-dst-code-cfg"`
    DstCodeOtherRate int `json:"dst-code-other-rate"`
    DstCodeOtherRateActionListName string `json:"dst-code-other-rate-action-list-name"`
    DstCodeOtherRateAction string `json:"dst-code-other-rate-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneTemplateIcmpV4TypeListV4SrcRateCfg struct {
    SrcTypeRate int `json:"src-type-rate"`
    SrcTypeRateActionListName string `json:"src-type-rate-action-list-name"`
    SrcTypeRateAction string `json:"src-type-rate-action"`
}


type DdosZoneTemplateIcmpV4TypeListV4SrcCodeCfg struct {
    SrcCodeNumber int `json:"src-code-number"`
    SrcCodeRate int `json:"src-code-rate"`
    SrcCodeRateActionListName string `json:"src-code-rate-action-list-name"`
    SrcCodeRateAction string `json:"src-code-rate-action"`
}


type DdosZoneTemplateIcmpV4TypeListV4DstRateCfg struct {
    DstTypeRate int `json:"dst-type-rate"`
    DstTypeRateActionListName string `json:"dst-type-rate-action-list-name"`
    DstTypeRateAction string `json:"dst-type-rate-action"`
}


type DdosZoneTemplateIcmpV4TypeListV4DstCodeCfg struct {
    DstCodeNumber int `json:"dst-code-number"`
    DstCodeRate int `json:"dst-code-rate"`
    DstCodeRateActionListName string `json:"dst-code-rate-action-list-name"`
    DstCodeRateAction string `json:"dst-code-rate-action"`
}


type DdosZoneTemplateIcmpV4TypeOther309 struct {
    IcmpTypeOtherActionListName string `json:"icmp-type-other-action-list-name"`
    IcmpTypeOtherAction string `json:"icmp-type-other-action"`
    Src DdosZoneTemplateIcmpV4TypeOtherSrc310 `json:"src"`
    Dst DdosZoneTemplateIcmpV4TypeOtherDst311 `json:"dst"`
    Uuid string `json:"uuid"`
}


type DdosZoneTemplateIcmpV4TypeOtherSrc310 struct {
    SrcTypeOtherRate int `json:"src-type-other-rate"`
    SrcTypeOtherRateActionListName string `json:"src-type-other-rate-action-list-name"`
    SrcTypeOtherRateAction string `json:"src-type-other-rate-action"`
}


type DdosZoneTemplateIcmpV4TypeOtherDst311 struct {
    DstTypeOtherRate int `json:"dst-type-other-rate"`
    DstTypeOtherRateActionListName string `json:"dst-type-other-rate-action-list-name"`
    DstTypeOtherRateAction string `json:"dst-type-other-rate-action"`
}

func (p *DdosZoneTemplateIcmpV4) GetId() string{
    return url.QueryEscape(p.Inst.IcmpTmplName)
}

func (p *DdosZoneTemplateIcmpV4) getPath() string{
    return "ddos/zone-template/icmp-v4"
}

func (p *DdosZoneTemplateIcmpV4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4::Post")
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

func (p *DdosZoneTemplateIcmpV4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4::Get")
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
func (p *DdosZoneTemplateIcmpV4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4::Put")
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

func (p *DdosZoneTemplateIcmpV4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
