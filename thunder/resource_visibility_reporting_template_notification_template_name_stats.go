package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityReportingTemplateNotificationTemplateNameStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_reporting_template_notification_template_name_stats`: Statistics for the object template-name\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityReportingTemplateNotificationTemplateNameStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Notification template name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sent_successful": {
							Type: schema.TypeInt, Optional: true, Description: "Sent successful",
						},
						"send_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Send failures",
						},
						"response_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Response failures",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityReportingTemplateNotificationTemplateNameStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTemplateNotificationTemplateNameStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTemplateNotificationTemplateNameStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityReportingTemplateNotificationTemplateNameStatsStats := setObjectVisibilityReportingTemplateNotificationTemplateNameStatsStats(res)
		d.Set("stats", VisibilityReportingTemplateNotificationTemplateNameStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityReportingTemplateNotificationTemplateNameStatsStats(ret edpt.DataVisibilityReportingTemplateNotificationTemplateNameStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sent_successful": ret.DtVisibilityReportingTemplateNotificationTemplateNameStats.Stats.Sent_successful,
			"send_fail":       ret.DtVisibilityReportingTemplateNotificationTemplateNameStats.Stats.Send_fail,
			"response_fail":   ret.DtVisibilityReportingTemplateNotificationTemplateNameStats.Stats.Response_fail,
		},
	}
}

func getObjectVisibilityReportingTemplateNotificationTemplateNameStatsStats(d []interface{}) edpt.VisibilityReportingTemplateNotificationTemplateNameStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityReportingTemplateNotificationTemplateNameStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sent_successful = in["sent_successful"].(int)
		ret.Send_fail = in["send_fail"].(int)
		ret.Response_fail = in["response_fail"].(int)
	}
	return ret
}

func dataToEndpointVisibilityReportingTemplateNotificationTemplateNameStats(d *schema.ResourceData) edpt.VisibilityReportingTemplateNotificationTemplateNameStats {
	var ret edpt.VisibilityReportingTemplateNotificationTemplateNameStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectVisibilityReportingTemplateNotificationTemplateNameStatsStats(d.Get("stats").([]interface{}))
	return ret
}
