package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityFlowCollectorNetflowStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_flow_collector_netflow_stats`: Statistics for the object netflow\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityFlowCollectorNetflowStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkts_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Total nflow packets received",
						},
						"v9_templates_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total v9 templates created",
						},
						"v9_templates_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Total v9 templates deleted",
						},
						"v10_templates_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total v10(IPFIX) templates created",
						},
						"v10_templates_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Total v10(IPFIX) templates deleted",
						},
						"template_drop_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Total templates dropped because of maximum limit",
						},
						"template_drop_out_of_memory": {
							Type: schema.TypeInt, Optional: true, Description: "Total templates dropped becuase of out of memory",
						},
						"frag_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total nflow fragment packets dropped",
						},
						"agent_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "nflow pkts from not configured agents",
						},
						"version_not_supported": {
							Type: schema.TypeInt, Optional: true, Description: "nflow version not supported",
						},
						"unknown_dir": {
							Type: schema.TypeInt, Optional: true, Description: "nflow sample direction is unknown",
						},
					},
				},
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"templates_added_to_delq": {
										Type: schema.TypeInt, Optional: true, Description: "Netflow templates added to the delete queue",
									},
									"templates_removed_from_delq": {
										Type: schema.TypeInt, Optional: true, Description: "Netflow templates removed from the delete queue",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityFlowCollectorNetflowStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityFlowCollectorNetflowStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityFlowCollectorNetflowStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityFlowCollectorNetflowStatsStats := setObjectVisibilityFlowCollectorNetflowStatsStats(res)
		d.Set("stats", VisibilityFlowCollectorNetflowStatsStats)
		VisibilityFlowCollectorNetflowStatsTemplate := setObjectVisibilityFlowCollectorNetflowStatsTemplate(res)
		d.Set("template", VisibilityFlowCollectorNetflowStatsTemplate)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityFlowCollectorNetflowStatsStats(ret edpt.DataVisibilityFlowCollectorNetflowStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pkts_rcvd":                   ret.DtVisibilityFlowCollectorNetflowStats.Stats.PktsRcvd,
			"v9_templates_created":        ret.DtVisibilityFlowCollectorNetflowStats.Stats.V9TemplatesCreated,
			"v9_templates_deleted":        ret.DtVisibilityFlowCollectorNetflowStats.Stats.V9TemplatesDeleted,
			"v10_templates_created":       ret.DtVisibilityFlowCollectorNetflowStats.Stats.V10TemplatesCreated,
			"v10_templates_deleted":       ret.DtVisibilityFlowCollectorNetflowStats.Stats.V10TemplatesDeleted,
			"template_drop_exceeded":      ret.DtVisibilityFlowCollectorNetflowStats.Stats.TemplateDropExceeded,
			"template_drop_out_of_memory": ret.DtVisibilityFlowCollectorNetflowStats.Stats.TemplateDropOutOfMemory,
			"frag_dropped":                ret.DtVisibilityFlowCollectorNetflowStats.Stats.FragDropped,
			"agent_not_found":             ret.DtVisibilityFlowCollectorNetflowStats.Stats.AgentNotFound,
			"version_not_supported":       ret.DtVisibilityFlowCollectorNetflowStats.Stats.VersionNotSupported,
			"unknown_dir":                 ret.DtVisibilityFlowCollectorNetflowStats.Stats.UnknownDir,
		},
	}
}

func setObjectVisibilityFlowCollectorNetflowStatsTemplate(ret edpt.DataVisibilityFlowCollectorNetflowStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectVisibilityFlowCollectorNetflowStatsTemplateStats(ret.DtVisibilityFlowCollectorNetflowStats.Template.Stats),
		},
	}
}

func setObjectVisibilityFlowCollectorNetflowStatsTemplateStats(d edpt.VisibilityFlowCollectorNetflowStatsTemplateStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["templates_added_to_delq"] = d.TemplatesAddedToDelq

	in["templates_removed_from_delq"] = d.TemplatesRemovedFromDelq
	result = append(result, in)
	return result
}

func getObjectVisibilityFlowCollectorNetflowStatsStats(d []interface{}) edpt.VisibilityFlowCollectorNetflowStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktsRcvd = in["pkts_rcvd"].(int)
		ret.V9TemplatesCreated = in["v9_templates_created"].(int)
		ret.V9TemplatesDeleted = in["v9_templates_deleted"].(int)
		ret.V10TemplatesCreated = in["v10_templates_created"].(int)
		ret.V10TemplatesDeleted = in["v10_templates_deleted"].(int)
		ret.TemplateDropExceeded = in["template_drop_exceeded"].(int)
		ret.TemplateDropOutOfMemory = in["template_drop_out_of_memory"].(int)
		ret.FragDropped = in["frag_dropped"].(int)
		ret.AgentNotFound = in["agent_not_found"].(int)
		ret.VersionNotSupported = in["version_not_supported"].(int)
		ret.UnknownDir = in["unknown_dir"].(int)
	}
	return ret
}

func getObjectVisibilityFlowCollectorNetflowStatsTemplate(d []interface{}) edpt.VisibilityFlowCollectorNetflowStatsTemplate {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowStatsTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVisibilityFlowCollectorNetflowStatsTemplateStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityFlowCollectorNetflowStatsTemplateStats(d []interface{}) edpt.VisibilityFlowCollectorNetflowStatsTemplateStats {

	count1 := len(d)
	var ret edpt.VisibilityFlowCollectorNetflowStatsTemplateStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplatesAddedToDelq = in["templates_added_to_delq"].(int)
		ret.TemplatesRemovedFromDelq = in["templates_removed_from_delq"].(int)
	}
	return ret
}

func dataToEndpointVisibilityFlowCollectorNetflowStats(d *schema.ResourceData) edpt.VisibilityFlowCollectorNetflowStats {
	var ret edpt.VisibilityFlowCollectorNetflowStats

	ret.Stats = getObjectVisibilityFlowCollectorNetflowStatsStats(d.Get("stats").([]interface{}))

	ret.Template = getObjectVisibilityFlowCollectorNetflowStatsTemplate(d.Get("template").([]interface{}))
	return ret
}
