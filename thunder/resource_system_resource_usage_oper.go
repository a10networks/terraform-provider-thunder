package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_session_count_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_session_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_session_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l4_session_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool_addr_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool_addr_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool_addr_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_pool_addr_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ipv6_addr_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ipv6_addr_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ipv6_addr_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ipv6_addr_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ac_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ac_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ac_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_ac_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_html_file_size_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_html_file_size_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_html_file_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_html_file_size_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_image_file_size_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_image_file_size_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_image_file_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_portal_image_file_size_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_file_size_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_file_size_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_file_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_file_size_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_table_entry_count_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_table_entry_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_table_entry_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_table_entry_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_authz_collection_number_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_authz_collection_number_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_authz_collection_number_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_authz_collection_number_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"visibility_mon_entity_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"visibility_mon_entity_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"visibility_mon_entity_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"visibility_mon_entity_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"authz_policy_number_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"authz_policy_number_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"authz_policy_number_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"authz_policy_number_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_sa_number_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_sa_number_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_sa_number_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_sa_number_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ram_cache_memory_limit_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ram_cache_memory_limit_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ram_cache_memory_limit_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ram_cache_memory_limit_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"waf_template_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"waf_template_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"waf_template_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"waf_template_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_session_count_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_session_count_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_session_count_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_session_count_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_entry_cur": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_entry_min": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_entry_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list_entry_default": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemResourceUsageOperOper := setObjectSystemResourceUsageOperOper(res)
		d.Set("oper", SystemResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemResourceUsageOperOper(ret edpt.DataSystemResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"l4_session_count_cur":                  ret.DtSystemResourceUsageOper.Oper.L4SessionCountCur,
			"l4_session_count_min":                  ret.DtSystemResourceUsageOper.Oper.L4SessionCountMin,
			"l4_session_count_max":                  ret.DtSystemResourceUsageOper.Oper.L4SessionCountMax,
			"l4_session_count_default":              ret.DtSystemResourceUsageOper.Oper.L4SessionCountDefault,
			"nat_pool_addr_cur":                     ret.DtSystemResourceUsageOper.Oper.NatPoolAddrCur,
			"nat_pool_addr_min":                     ret.DtSystemResourceUsageOper.Oper.NatPoolAddrMin,
			"nat_pool_addr_max":                     ret.DtSystemResourceUsageOper.Oper.NatPoolAddrMax,
			"nat_pool_addr_default":                 ret.DtSystemResourceUsageOper.Oper.NatPoolAddrDefault,
			"class_list_ipv6_addr_cur":              ret.DtSystemResourceUsageOper.Oper.ClassListIpv6AddrCur,
			"class_list_ipv6_addr_min":              ret.DtSystemResourceUsageOper.Oper.ClassListIpv6AddrMin,
			"class_list_ipv6_addr_max":              ret.DtSystemResourceUsageOper.Oper.ClassListIpv6AddrMax,
			"class_list_ipv6_addr_default":          ret.DtSystemResourceUsageOper.Oper.ClassListIpv6AddrDefault,
			"class_list_ac_cur":                     ret.DtSystemResourceUsageOper.Oper.ClassListAcCur,
			"class_list_ac_min":                     ret.DtSystemResourceUsageOper.Oper.ClassListAcMin,
			"class_list_ac_max":                     ret.DtSystemResourceUsageOper.Oper.ClassListAcMax,
			"class_list_ac_default":                 ret.DtSystemResourceUsageOper.Oper.ClassListAcDefault,
			"auth_portal_html_file_size_cur":        ret.DtSystemResourceUsageOper.Oper.AuthPortalHtmlFileSizeCur,
			"auth_portal_html_file_size_min":        ret.DtSystemResourceUsageOper.Oper.AuthPortalHtmlFileSizeMin,
			"auth_portal_html_file_size_max":        ret.DtSystemResourceUsageOper.Oper.AuthPortalHtmlFileSizeMax,
			"auth_portal_html_file_size_default":    ret.DtSystemResourceUsageOper.Oper.AuthPortalHtmlFileSizeDefault,
			"auth_portal_image_file_size_cur":       ret.DtSystemResourceUsageOper.Oper.AuthPortalImageFileSizeCur,
			"auth_portal_image_file_size_min":       ret.DtSystemResourceUsageOper.Oper.AuthPortalImageFileSizeMin,
			"auth_portal_image_file_size_max":       ret.DtSystemResourceUsageOper.Oper.AuthPortalImageFileSizeMax,
			"auth_portal_image_file_size_default":   ret.DtSystemResourceUsageOper.Oper.AuthPortalImageFileSizeDefault,
			"aflex_file_size_cur":                   ret.DtSystemResourceUsageOper.Oper.AflexFileSizeCur,
			"aflex_file_size_min":                   ret.DtSystemResourceUsageOper.Oper.AflexFileSizeMin,
			"aflex_file_size_max":                   ret.DtSystemResourceUsageOper.Oper.AflexFileSizeMax,
			"aflex_file_size_default":               ret.DtSystemResourceUsageOper.Oper.AflexFileSizeDefault,
			"aflex_table_entry_count_cur":           ret.DtSystemResourceUsageOper.Oper.AflexTableEntryCountCur,
			"aflex_table_entry_count_min":           ret.DtSystemResourceUsageOper.Oper.AflexTableEntryCountMin,
			"aflex_table_entry_count_max":           ret.DtSystemResourceUsageOper.Oper.AflexTableEntryCountMax,
			"aflex_table_entry_count_default":       ret.DtSystemResourceUsageOper.Oper.AflexTableEntryCountDefault,
			"aflex_authz_collection_number_cur":     ret.DtSystemResourceUsageOper.Oper.AflexAuthzCollectionNumberCur,
			"aflex_authz_collection_number_min":     ret.DtSystemResourceUsageOper.Oper.AflexAuthzCollectionNumberMin,
			"aflex_authz_collection_number_max":     ret.DtSystemResourceUsageOper.Oper.AflexAuthzCollectionNumberMax,
			"aflex_authz_collection_number_default": ret.DtSystemResourceUsageOper.Oper.AflexAuthzCollectionNumberDefault,
			"radius_table_size_cur":                 ret.DtSystemResourceUsageOper.Oper.RadiusTableSizeCur,
			"radius_table_size_min":                 ret.DtSystemResourceUsageOper.Oper.RadiusTableSizeMin,
			"radius_table_size_max":                 ret.DtSystemResourceUsageOper.Oper.RadiusTableSizeMax,
			"radius_table_size_default":             ret.DtSystemResourceUsageOper.Oper.RadiusTableSizeDefault,
			"visibility_mon_entity_cur":             ret.DtSystemResourceUsageOper.Oper.VisibilityMonEntityCur,
			"visibility_mon_entity_min":             ret.DtSystemResourceUsageOper.Oper.VisibilityMonEntityMin,
			"visibility_mon_entity_max":             ret.DtSystemResourceUsageOper.Oper.VisibilityMonEntityMax,
			"visibility_mon_entity_default":         ret.DtSystemResourceUsageOper.Oper.VisibilityMonEntityDefault,
			"authz_policy_number_cur":               ret.DtSystemResourceUsageOper.Oper.AuthzPolicyNumberCur,
			"authz_policy_number_min":               ret.DtSystemResourceUsageOper.Oper.AuthzPolicyNumberMin,
			"authz_policy_number_max":               ret.DtSystemResourceUsageOper.Oper.AuthzPolicyNumberMax,
			"authz_policy_number_default":           ret.DtSystemResourceUsageOper.Oper.AuthzPolicyNumberDefault,
			"ipsec_sa_number_cur":                   ret.DtSystemResourceUsageOper.Oper.IpsecSaNumberCur,
			"ipsec_sa_number_min":                   ret.DtSystemResourceUsageOper.Oper.IpsecSaNumberMin,
			"ipsec_sa_number_max":                   ret.DtSystemResourceUsageOper.Oper.IpsecSaNumberMax,
			"ipsec_sa_number_default":               ret.DtSystemResourceUsageOper.Oper.IpsecSaNumberDefault,
			"ram_cache_memory_limit_cur":            ret.DtSystemResourceUsageOper.Oper.RamCacheMemoryLimitCur,
			"ram_cache_memory_limit_min":            ret.DtSystemResourceUsageOper.Oper.RamCacheMemoryLimitMin,
			"ram_cache_memory_limit_max":            ret.DtSystemResourceUsageOper.Oper.RamCacheMemoryLimitMax,
			"ram_cache_memory_limit_default":        ret.DtSystemResourceUsageOper.Oper.RamCacheMemoryLimitDefault,
			"waf_template_cur":                      ret.DtSystemResourceUsageOper.Oper.WafTemplateCur,
			"waf_template_min":                      ret.DtSystemResourceUsageOper.Oper.WafTemplateMin,
			"waf_template_max":                      ret.DtSystemResourceUsageOper.Oper.WafTemplateMax,
			"waf_template_default":                  ret.DtSystemResourceUsageOper.Oper.WafTemplateDefault,
			"auth_session_count_cur":                ret.DtSystemResourceUsageOper.Oper.AuthSessionCountCur,
			"auth_session_count_min":                ret.DtSystemResourceUsageOper.Oper.AuthSessionCountMin,
			"auth_session_count_max":                ret.DtSystemResourceUsageOper.Oper.AuthSessionCountMax,
			"auth_session_count_default":            ret.DtSystemResourceUsageOper.Oper.AuthSessionCountDefault,
			"class_list_entry_cur":                  ret.DtSystemResourceUsageOper.Oper.ClassListEntryCur,
			"class_list_entry_min":                  ret.DtSystemResourceUsageOper.Oper.ClassListEntryMin,
			"class_list_entry_max":                  ret.DtSystemResourceUsageOper.Oper.ClassListEntryMax,
			"class_list_entry_default":              ret.DtSystemResourceUsageOper.Oper.ClassListEntryDefault,
		},
	}
}

func getObjectSystemResourceUsageOperOper(d []interface{}) edpt.SystemResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.SystemResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4SessionCountCur = in["l4_session_count_cur"].(int)
		ret.L4SessionCountMin = in["l4_session_count_min"].(int)
		ret.L4SessionCountMax = in["l4_session_count_max"].(int)
		ret.L4SessionCountDefault = in["l4_session_count_default"].(int)
		ret.NatPoolAddrCur = in["nat_pool_addr_cur"].(int)
		ret.NatPoolAddrMin = in["nat_pool_addr_min"].(int)
		ret.NatPoolAddrMax = in["nat_pool_addr_max"].(int)
		ret.NatPoolAddrDefault = in["nat_pool_addr_default"].(int)
		ret.ClassListIpv6AddrCur = in["class_list_ipv6_addr_cur"].(int)
		ret.ClassListIpv6AddrMin = in["class_list_ipv6_addr_min"].(int)
		ret.ClassListIpv6AddrMax = in["class_list_ipv6_addr_max"].(int)
		ret.ClassListIpv6AddrDefault = in["class_list_ipv6_addr_default"].(int)
		ret.ClassListAcCur = in["class_list_ac_cur"].(int)
		ret.ClassListAcMin = in["class_list_ac_min"].(int)
		ret.ClassListAcMax = in["class_list_ac_max"].(int)
		ret.ClassListAcDefault = in["class_list_ac_default"].(int)
		ret.AuthPortalHtmlFileSizeCur = in["auth_portal_html_file_size_cur"].(int)
		ret.AuthPortalHtmlFileSizeMin = in["auth_portal_html_file_size_min"].(int)
		ret.AuthPortalHtmlFileSizeMax = in["auth_portal_html_file_size_max"].(int)
		ret.AuthPortalHtmlFileSizeDefault = in["auth_portal_html_file_size_default"].(int)
		ret.AuthPortalImageFileSizeCur = in["auth_portal_image_file_size_cur"].(int)
		ret.AuthPortalImageFileSizeMin = in["auth_portal_image_file_size_min"].(int)
		ret.AuthPortalImageFileSizeMax = in["auth_portal_image_file_size_max"].(int)
		ret.AuthPortalImageFileSizeDefault = in["auth_portal_image_file_size_default"].(int)
		ret.AflexFileSizeCur = in["aflex_file_size_cur"].(int)
		ret.AflexFileSizeMin = in["aflex_file_size_min"].(int)
		ret.AflexFileSizeMax = in["aflex_file_size_max"].(int)
		ret.AflexFileSizeDefault = in["aflex_file_size_default"].(int)
		ret.AflexTableEntryCountCur = in["aflex_table_entry_count_cur"].(int)
		ret.AflexTableEntryCountMin = in["aflex_table_entry_count_min"].(int)
		ret.AflexTableEntryCountMax = in["aflex_table_entry_count_max"].(int)
		ret.AflexTableEntryCountDefault = in["aflex_table_entry_count_default"].(int)
		ret.AflexAuthzCollectionNumberCur = in["aflex_authz_collection_number_cur"].(int)
		ret.AflexAuthzCollectionNumberMin = in["aflex_authz_collection_number_min"].(int)
		ret.AflexAuthzCollectionNumberMax = in["aflex_authz_collection_number_max"].(int)
		ret.AflexAuthzCollectionNumberDefault = in["aflex_authz_collection_number_default"].(int)
		ret.RadiusTableSizeCur = in["radius_table_size_cur"].(int)
		ret.RadiusTableSizeMin = in["radius_table_size_min"].(int)
		ret.RadiusTableSizeMax = in["radius_table_size_max"].(int)
		ret.RadiusTableSizeDefault = in["radius_table_size_default"].(int)
		ret.VisibilityMonEntityCur = in["visibility_mon_entity_cur"].(int)
		ret.VisibilityMonEntityMin = in["visibility_mon_entity_min"].(int)
		ret.VisibilityMonEntityMax = in["visibility_mon_entity_max"].(int)
		ret.VisibilityMonEntityDefault = in["visibility_mon_entity_default"].(int)
		ret.AuthzPolicyNumberCur = in["authz_policy_number_cur"].(int)
		ret.AuthzPolicyNumberMin = in["authz_policy_number_min"].(int)
		ret.AuthzPolicyNumberMax = in["authz_policy_number_max"].(int)
		ret.AuthzPolicyNumberDefault = in["authz_policy_number_default"].(int)
		ret.IpsecSaNumberCur = in["ipsec_sa_number_cur"].(int)
		ret.IpsecSaNumberMin = in["ipsec_sa_number_min"].(int)
		ret.IpsecSaNumberMax = in["ipsec_sa_number_max"].(int)
		ret.IpsecSaNumberDefault = in["ipsec_sa_number_default"].(int)
		ret.RamCacheMemoryLimitCur = in["ram_cache_memory_limit_cur"].(int)
		ret.RamCacheMemoryLimitMin = in["ram_cache_memory_limit_min"].(int)
		ret.RamCacheMemoryLimitMax = in["ram_cache_memory_limit_max"].(int)
		ret.RamCacheMemoryLimitDefault = in["ram_cache_memory_limit_default"].(int)
		ret.WafTemplateCur = in["waf_template_cur"].(int)
		ret.WafTemplateMin = in["waf_template_min"].(int)
		ret.WafTemplateMax = in["waf_template_max"].(int)
		ret.WafTemplateDefault = in["waf_template_default"].(int)
		ret.AuthSessionCountCur = in["auth_session_count_cur"].(int)
		ret.AuthSessionCountMin = in["auth_session_count_min"].(int)
		ret.AuthSessionCountMax = in["auth_session_count_max"].(int)
		ret.AuthSessionCountDefault = in["auth_session_count_default"].(int)
		ret.ClassListEntryCur = in["class_list_entry_cur"].(int)
		ret.ClassListEntryMin = in["class_list_entry_min"].(int)
		ret.ClassListEntryMax = in["class_list_entry_max"].(int)
		ret.ClassListEntryDefault = in["class_list_entry_default"].(int)
	}
	return ret
}

func dataToEndpointSystemResourceUsageOper(d *schema.ResourceData) edpt.SystemResourceUsageOper {
	var ret edpt.SystemResourceUsageOper

	ret.Oper = getObjectSystemResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
