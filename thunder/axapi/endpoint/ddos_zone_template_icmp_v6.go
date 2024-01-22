

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateIcmpV6 struct {
	Inst struct {

    FilterList []DdosZoneTemplateIcmpV6FilterList `json:"filter-list"`
    FilterMatchType string `json:"filter-match-type" dval:"default"`
    IcmpTmplName string `json:"icmp-tmpl-name"`
    TypeList []DdosZoneTemplateIcmpV6TypeList `json:"type-list"`
    TypeOther DdosZoneTemplateIcmpV6TypeOther312 `json:"type-other"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"icmp-v6"`
}


type DdosZoneTemplateIcmpV6FilterList struct {
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


type DdosZoneTemplateIcmpV6TypeList struct {
    TypeNumber int `json:"type-number"`
    IcmpTypeActionListName string `json:"icmp-type-action-list-name"`
    IcmpTypeAction string `json:"icmp-type-action"`
    V6SrcRateCfg DdosZoneTemplateIcmpV6TypeListV6SrcRateCfg `json:"v6-src-rate-cfg"`
    V6SrcCodeCfg []DdosZoneTemplateIcmpV6TypeListV6SrcCodeCfg `json:"v6-src-code-cfg"`
    SrcCodeOtherRate int `json:"src-code-other-rate"`
    SrcCodeOtherRateActionListName string `json:"src-code-other-rate-action-list-name"`
    SrcCodeOtherRateAction string `json:"src-code-other-rate-action"`
    V6DstRateCfg DdosZoneTemplateIcmpV6TypeListV6DstRateCfg `json:"v6-dst-rate-cfg"`
    V6DstCodeCfg []DdosZoneTemplateIcmpV6TypeListV6DstCodeCfg `json:"v6-dst-code-cfg"`
    DstCodeOtherRate int `json:"dst-code-other-rate"`
    DstCodeOtherRateActionListName string `json:"dst-code-other-rate-action-list-name"`
    DstCodeOtherRateAction string `json:"dst-code-other-rate-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneTemplateIcmpV6TypeListV6SrcRateCfg struct {
    SrcTypeRate int `json:"src-type-rate"`
    SrcTypeRateActionListName string `json:"src-type-rate-action-list-name"`
    SrcTypeRateAction string `json:"src-type-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeListV6SrcCodeCfg struct {
    SrcCodeNumber int `json:"src-code-number"`
    SrcCodeRate int `json:"src-code-rate"`
    SrcCodeRateActionListName string `json:"src-code-rate-action-list-name"`
    SrcCodeRateAction string `json:"src-code-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeListV6DstRateCfg struct {
    DstTypeRate int `json:"dst-type-rate"`
    DstTypeRateActionListName string `json:"dst-type-rate-action-list-name"`
    DstTypeRateAction string `json:"dst-type-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeListV6DstCodeCfg struct {
    DstCodeNumber int `json:"dst-code-number"`
    DstCodeRate int `json:"dst-code-rate"`
    DstCodeRateActionListName string `json:"dst-code-rate-action-list-name"`
    DstCodeRateAction string `json:"dst-code-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeOther312 struct {
    IcmpTypeOtherActionListName string `json:"icmp-type-other-action-list-name"`
    IcmpTypeOtherAction string `json:"icmp-type-other-action"`
    Src DdosZoneTemplateIcmpV6TypeOtherSrc313 `json:"src"`
    Dst DdosZoneTemplateIcmpV6TypeOtherDst314 `json:"dst"`
    Uuid string `json:"uuid"`
}


type DdosZoneTemplateIcmpV6TypeOtherSrc313 struct {
    SrcTypeOtherRate int `json:"src-type-other-rate"`
    SrcTypeOtherRateActionListName string `json:"src-type-other-rate-action-list-name"`
    SrcTypeOtherRateAction string `json:"src-type-other-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeOtherDst314 struct {
    DstTypeOtherRate int `json:"dst-type-other-rate"`
    DstTypeOtherRateActionListName string `json:"dst-type-other-rate-action-list-name"`
    DstTypeOtherRateAction string `json:"dst-type-other-rate-action"`
}

func (p *DdosZoneTemplateIcmpV6) GetId() string{
    return url.QueryEscape(p.Inst.IcmpTmplName)
}

func (p *DdosZoneTemplateIcmpV6) getPath() string{
    return "ddos/zone-template/icmp-v6"
}

func (p *DdosZoneTemplateIcmpV6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6::Post")
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

func (p *DdosZoneTemplateIcmpV6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6::Get")
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
func (p *DdosZoneTemplateIcmpV6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6::Put")
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

func (p *DdosZoneTemplateIcmpV6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
