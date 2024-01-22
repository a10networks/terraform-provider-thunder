

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemResourceUsageOper struct {
    
    Oper SystemResourceUsageOperOper `json:"oper"`

}
type DataSystemResourceUsageOper struct {
    DtSystemResourceUsageOper SystemResourceUsageOper `json:"resource-usage"`
}


type SystemResourceUsageOperOper struct {
    L4SessionCountCur int `json:"l4-session-count-cur"`
    L4SessionCountMin int `json:"l4-session-count-min"`
    L4SessionCountMax int `json:"l4-session-count-max"`
    L4SessionCountDefault int `json:"l4-session-count-default"`
    NatPoolAddrCur int `json:"nat-pool-addr-cur"`
    NatPoolAddrMin int `json:"nat-pool-addr-min"`
    NatPoolAddrMax int `json:"nat-pool-addr-max"`
    NatPoolAddrDefault int `json:"nat-pool-addr-default"`
    ClassListIpv6AddrCur int `json:"class-list-ipv6-addr-cur"`
    ClassListIpv6AddrMin int `json:"class-list-ipv6-addr-min"`
    ClassListIpv6AddrMax int `json:"class-list-ipv6-addr-max"`
    ClassListIpv6AddrDefault int `json:"class-list-ipv6-addr-default"`
    ClassListAcCur int `json:"class-list-ac-cur"`
    ClassListAcMin int `json:"class-list-ac-min"`
    ClassListAcMax int `json:"class-list-ac-max"`
    ClassListAcDefault int `json:"class-list-ac-default"`
    AuthPortalHtmlFileSizeCur int `json:"auth-portal-html-file-size-cur"`
    AuthPortalHtmlFileSizeMin int `json:"auth-portal-html-file-size-min"`
    AuthPortalHtmlFileSizeMax int `json:"auth-portal-html-file-size-max"`
    AuthPortalHtmlFileSizeDefault int `json:"auth-portal-html-file-size-default"`
    AuthPortalImageFileSizeCur int `json:"auth-portal-image-file-size-cur"`
    AuthPortalImageFileSizeMin int `json:"auth-portal-image-file-size-min"`
    AuthPortalImageFileSizeMax int `json:"auth-portal-image-file-size-max"`
    AuthPortalImageFileSizeDefault int `json:"auth-portal-image-file-size-default"`
    AflexFileSizeCur int `json:"aflex-file-size-cur"`
    AflexFileSizeMin int `json:"aflex-file-size-min"`
    AflexFileSizeMax int `json:"aflex-file-size-max"`
    AflexFileSizeDefault int `json:"aflex-file-size-default"`
    AflexTableEntryCountCur int `json:"aflex-table-entry-count-cur"`
    AflexTableEntryCountMin int `json:"aflex-table-entry-count-min"`
    AflexTableEntryCountMax int `json:"aflex-table-entry-count-max"`
    AflexTableEntryCountDefault int `json:"aflex-table-entry-count-default"`
    AflexAuthzCollectionNumberCur int `json:"aflex-authz-collection-number-cur"`
    AflexAuthzCollectionNumberMin int `json:"aflex-authz-collection-number-min"`
    AflexAuthzCollectionNumberMax int `json:"aflex-authz-collection-number-max"`
    AflexAuthzCollectionNumberDefault int `json:"aflex-authz-collection-number-default"`
    RadiusTableSizeCur int `json:"radius-table-size-cur"`
    RadiusTableSizeMin int `json:"radius-table-size-min"`
    RadiusTableSizeMax int `json:"radius-table-size-max"`
    RadiusTableSizeDefault int `json:"radius-table-size-default"`
    VisibilityMonEntityCur int `json:"visibility-mon-entity-cur"`
    VisibilityMonEntityMin int `json:"visibility-mon-entity-min"`
    VisibilityMonEntityMax int `json:"visibility-mon-entity-max"`
    VisibilityMonEntityDefault int `json:"visibility-mon-entity-default"`
    AuthzPolicyNumberCur int `json:"authz-policy-number-cur"`
    AuthzPolicyNumberMin int `json:"authz-policy-number-min"`
    AuthzPolicyNumberMax int `json:"authz-policy-number-max"`
    AuthzPolicyNumberDefault int `json:"authz-policy-number-default"`
    IpsecSaNumberCur int `json:"ipsec-sa-number-cur"`
    IpsecSaNumberMin int `json:"ipsec-sa-number-min"`
    IpsecSaNumberMax int `json:"ipsec-sa-number-max"`
    IpsecSaNumberDefault int `json:"ipsec-sa-number-default"`
    RamCacheMemoryLimitCur int `json:"ram-cache-memory-limit-cur"`
    RamCacheMemoryLimitMin int `json:"ram-cache-memory-limit-min"`
    RamCacheMemoryLimitMax int `json:"ram-cache-memory-limit-max"`
    RamCacheMemoryLimitDefault int `json:"ram-cache-memory-limit-default"`
    WafTemplateCur int `json:"waf-template-cur"`
    WafTemplateMin int `json:"waf-template-min"`
    WafTemplateMax int `json:"waf-template-max"`
    WafTemplateDefault int `json:"waf-template-default"`
    AuthSessionCountCur int `json:"auth-session-count-cur"`
    AuthSessionCountMin int `json:"auth-session-count-min"`
    AuthSessionCountMax int `json:"auth-session-count-max"`
    AuthSessionCountDefault int `json:"auth-session-count-default"`
    ClassListEntryCur int `json:"class-list-entry-cur"`
    ClassListEntryMin int `json:"class-list-entry-min"`
    ClassListEntryMax int `json:"class-list-entry-max"`
    ClassListEntryDefault int `json:"class-list-entry-default"`
}

func (p *SystemResourceUsageOper) GetId() string{
    return "1"
}

func (p *SystemResourceUsageOper) getPath() string{
    return "system/resource-usage/oper"
}

func (p *SystemResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemResourceUsageOper,error) {
logger.Println("SystemResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemResourceUsageOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
