package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwResourceUsageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_resource_usage_oper`: Operational Status for the object resource-usage\n\n__PLACEHOLDER__",
		ReadContext: resourceFwResourceUsageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fw_object_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_object_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_object_group_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_object_group_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_rule_set_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_rule_set_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_rule_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_rule_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_zone_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_zone_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_ip_range_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_ip_range_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_helper_sessions_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_helper_sessions_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_current_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"clause_per_obj_grp_current_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"clause_per_obj_grp_total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"object": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"object_group": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rule_set": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rule": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"zone": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip_range": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"helper_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"radius_table_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"clause_per_obj_grp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwResourceUsageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwResourceUsageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwResourceUsageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwResourceUsageOperOper := setObjectFwResourceUsageOperOper(res)
		d.Set("oper", FwResourceUsageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwResourceUsageOperOper(ret edpt.DataFwResourceUsageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"fw_object_current_count":          ret.DtFwResourceUsageOper.Oper.FwObjectCurrentCount,
			"fw_object_total_count":            ret.DtFwResourceUsageOper.Oper.FwObjectTotalCount,
			"fw_object_group_current_count":    ret.DtFwResourceUsageOper.Oper.FwObjectGroupCurrentCount,
			"fw_object_group_total_count":      ret.DtFwResourceUsageOper.Oper.FwObjectGroupTotalCount,
			"fw_rule_set_current_count":        ret.DtFwResourceUsageOper.Oper.FwRuleSetCurrentCount,
			"fw_rule_set_total_count":          ret.DtFwResourceUsageOper.Oper.FwRuleSetTotalCount,
			"fw_rule_current_count":            ret.DtFwResourceUsageOper.Oper.FwRuleCurrentCount,
			"fw_rule_total_count":              ret.DtFwResourceUsageOper.Oper.FwRuleTotalCount,
			"fw_zone_current_count":            ret.DtFwResourceUsageOper.Oper.FwZoneCurrentCount,
			"fw_zone_total_count":              ret.DtFwResourceUsageOper.Oper.FwZoneTotalCount,
			"fw_ip_range_current_count":        ret.DtFwResourceUsageOper.Oper.FwIpRangeCurrentCount,
			"fw_ip_range_total_count":          ret.DtFwResourceUsageOper.Oper.FwIpRangeTotalCount,
			"fw_helper_sessions_current_count": ret.DtFwResourceUsageOper.Oper.FwHelperSessionsCurrentCount,
			"fw_helper_sessions_total_count":   ret.DtFwResourceUsageOper.Oper.FwHelperSessionsTotalCount,
			"radius_table_current_count":       ret.DtFwResourceUsageOper.Oper.RadiusTableCurrentCount,
			"radius_table_total_count":         ret.DtFwResourceUsageOper.Oper.RadiusTableTotalCount,
			"clause_per_obj_grp_current_count": ret.DtFwResourceUsageOper.Oper.ClausePerObjGrpCurrentCount,
			"clause_per_obj_grp_total_count":   ret.DtFwResourceUsageOper.Oper.ClausePerObjGrpTotalCount,
			"object":                           ret.DtFwResourceUsageOper.Oper.Object,
			"object_group":                     ret.DtFwResourceUsageOper.Oper.ObjectGroup,
			"rule_set":                         ret.DtFwResourceUsageOper.Oper.RuleSet,
			"rule":                             ret.DtFwResourceUsageOper.Oper.Rule,
			"zone":                             ret.DtFwResourceUsageOper.Oper.Zone,
			"ip_range":                         ret.DtFwResourceUsageOper.Oper.IpRange,
			"helper_sessions":                  ret.DtFwResourceUsageOper.Oper.HelperSessions,
			"radius_table_size":                ret.DtFwResourceUsageOper.Oper.RadiusTableSize,
			"clause_per_obj_grp":               ret.DtFwResourceUsageOper.Oper.ClausePerObjGrp,
		},
	}
}

func getObjectFwResourceUsageOperOper(d []interface{}) edpt.FwResourceUsageOperOper {

	count1 := len(d)
	var ret edpt.FwResourceUsageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FwObjectCurrentCount = in["fw_object_current_count"].(int)
		ret.FwObjectTotalCount = in["fw_object_total_count"].(int)
		ret.FwObjectGroupCurrentCount = in["fw_object_group_current_count"].(int)
		ret.FwObjectGroupTotalCount = in["fw_object_group_total_count"].(int)
		ret.FwRuleSetCurrentCount = in["fw_rule_set_current_count"].(int)
		ret.FwRuleSetTotalCount = in["fw_rule_set_total_count"].(int)
		ret.FwRuleCurrentCount = in["fw_rule_current_count"].(int)
		ret.FwRuleTotalCount = in["fw_rule_total_count"].(int)
		ret.FwZoneCurrentCount = in["fw_zone_current_count"].(int)
		ret.FwZoneTotalCount = in["fw_zone_total_count"].(int)
		ret.FwIpRangeCurrentCount = in["fw_ip_range_current_count"].(int)
		ret.FwIpRangeTotalCount = in["fw_ip_range_total_count"].(int)
		ret.FwHelperSessionsCurrentCount = in["fw_helper_sessions_current_count"].(int)
		ret.FwHelperSessionsTotalCount = in["fw_helper_sessions_total_count"].(int)
		ret.RadiusTableCurrentCount = in["radius_table_current_count"].(int)
		ret.RadiusTableTotalCount = in["radius_table_total_count"].(int)
		ret.ClausePerObjGrpCurrentCount = in["clause_per_obj_grp_current_count"].(string)
		ret.ClausePerObjGrpTotalCount = in["clause_per_obj_grp_total_count"].(int)
		ret.Object = in["object"].(int)
		ret.ObjectGroup = in["object_group"].(int)
		ret.RuleSet = in["rule_set"].(int)
		ret.Rule = in["rule"].(int)
		ret.Zone = in["zone"].(int)
		ret.IpRange = in["ip_range"].(int)
		ret.HelperSessions = in["helper_sessions"].(int)
		ret.RadiusTableSize = in["radius_table_size"].(int)
		ret.ClausePerObjGrp = in["clause_per_obj_grp"].(int)
	}
	return ret
}

func dataToEndpointFwResourceUsageOper(d *schema.ResourceData) edpt.FwResourceUsageOper {
	var ret edpt.FwResourceUsageOper

	ret.Oper = getObjectFwResourceUsageOperOper(d.Get("oper").([]interface{}))
	return ret
}
