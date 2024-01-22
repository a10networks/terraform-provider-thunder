package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityReportingTemplateNotificationDebugOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_reporting_template_notification_debug_oper`: Operational Status for the object debug\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityReportingTemplateNotificationDebugOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"debug_log": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"log": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityReportingTemplateNotificationDebugOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTemplateNotificationDebugOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTemplateNotificationDebugOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityReportingTemplateNotificationDebugOperOper := setObjectVisibilityReportingTemplateNotificationDebugOperOper(res)
		d.Set("oper", VisibilityReportingTemplateNotificationDebugOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityReportingTemplateNotificationDebugOperOper(ret edpt.DataVisibilityReportingTemplateNotificationDebugOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"debug_list": setSliceVisibilityReportingTemplateNotificationDebugOperOperDebugList(ret.DtVisibilityReportingTemplateNotificationDebugOper.Oper.DebugList),
			"log":        ret.DtVisibilityReportingTemplateNotificationDebugOper.Oper.Log,
		},
	}
}

func setSliceVisibilityReportingTemplateNotificationDebugOperOperDebugList(d []edpt.VisibilityReportingTemplateNotificationDebugOperOperDebugList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["debug_log"] = item.DebugLog
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityReportingTemplateNotificationDebugOperOper(d []interface{}) edpt.VisibilityReportingTemplateNotificationDebugOperOper {

	count1 := len(d)
	var ret edpt.VisibilityReportingTemplateNotificationDebugOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugList = getSliceVisibilityReportingTemplateNotificationDebugOperOperDebugList(in["debug_list"].([]interface{}))
		ret.Log = in["log"].(int)
	}
	return ret
}

func getSliceVisibilityReportingTemplateNotificationDebugOperOperDebugList(d []interface{}) []edpt.VisibilityReportingTemplateNotificationDebugOperOperDebugList {

	count1 := len(d)
	ret := make([]edpt.VisibilityReportingTemplateNotificationDebugOperOperDebugList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityReportingTemplateNotificationDebugOperOperDebugList
		oi.DebugLog = in["debug_log"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityReportingTemplateNotificationDebugOper(d *schema.ResourceData) edpt.VisibilityReportingTemplateNotificationDebugOper {
	var ret edpt.VisibilityReportingTemplateNotificationDebugOper

	ret.Oper = getObjectVisibilityReportingTemplateNotificationDebugOperOper(d.Get("oper").([]interface{}))
	return ret
}
