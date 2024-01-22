package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNotificationTemplateCommon() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_notification_template_common`: Global parameters for the ddos notification templates\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNotificationTemplateCommonCreate,
		UpdateContext: resourceDdosNotificationTemplateCommonUpdate,
		ReadContext:   resourceDdosNotificationTemplateCommonRead,
		DeleteContext: resourceDdosNotificationTemplateCommonDelete,

		Schema: map[string]*schema.Schema{
			"default_template": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_notification_template": {
							Type: schema.TypeString, Optional: true, Description: "Specify the notification template name (Default notification template name)",
						},
					},
				},
			},
			"on_box_gui_notification": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosNotificationTemplateCommonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateCommonCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateCommon(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateCommonRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNotificationTemplateCommonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateCommonUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateCommon(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateCommonRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNotificationTemplateCommonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateCommonDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateCommon(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNotificationTemplateCommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateCommonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateCommon(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosNotificationTemplateCommonDefaultTemplate(d []interface{}) []edpt.DdosNotificationTemplateCommonDefaultTemplate {

	count1 := len(d)
	ret := make([]edpt.DdosNotificationTemplateCommonDefaultTemplate, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNotificationTemplateCommonDefaultTemplate
		oi.DefaultNotificationTemplate = in["default_notification_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosNotificationTemplateCommon(d *schema.ResourceData) edpt.DdosNotificationTemplateCommon {
	var ret edpt.DdosNotificationTemplateCommon
	ret.Inst.DefaultTemplate = getSliceDdosNotificationTemplateCommonDefaultTemplate(d.Get("default_template").([]interface{}))
	ret.Inst.OnBoxGuiNotification = d.Get("on_box_gui_notification").(string)
	//omit uuid
	return ret
}
