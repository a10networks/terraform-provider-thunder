package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNotificationTemplateDebugLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_notification_template_debug_log_oper`: Operational Status for the object notification-template-debug-log\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNotificationTemplateDebugLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_template_debug_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification_template_debug_log": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceDdosNotificationTemplateDebugLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateDebugLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateDebugLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNotificationTemplateDebugLogOperOper := setObjectDdosNotificationTemplateDebugLogOperOper(res)
		d.Set("oper", DdosNotificationTemplateDebugLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNotificationTemplateDebugLogOperOper(ret edpt.DataDdosNotificationTemplateDebugLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"notification_template_debug_list": setSliceDdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList(ret.DtDdosNotificationTemplateDebugLogOper.Oper.NotificationTemplateDebugList),
		},
	}
}

func setSliceDdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList(d []edpt.DdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["notification_template_debug_log"] = item.NotificationTemplateDebugLog
		result = append(result, in)
	}
	return result
}

func getObjectDdosNotificationTemplateDebugLogOperOper(d []interface{}) edpt.DdosNotificationTemplateDebugLogOperOper {

	count1 := len(d)
	var ret edpt.DdosNotificationTemplateDebugLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NotificationTemplateDebugList = getSliceDdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList(in["notification_template_debug_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList(d []interface{}) []edpt.DdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList {

	count1 := len(d)
	ret := make([]edpt.DdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList
		oi.NotificationTemplateDebugLog = in["notification_template_debug_log"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosNotificationTemplateDebugLogOper(d *schema.ResourceData) edpt.DdosNotificationTemplateDebugLogOper {
	var ret edpt.DdosNotificationTemplateDebugLogOper

	ret.Oper = getObjectDdosNotificationTemplateDebugLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
