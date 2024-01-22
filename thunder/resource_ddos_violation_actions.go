package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosViolationActions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_violation_actions`: Violation Actions Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosViolationActionsCreate,
		UpdateContext: resourceDdosViolationActionsUpdate,
		ReadContext:   resourceDdosViolationActionsRead,
		DeleteContext: resourceDdosViolationActionsDelete,

		Schema: map[string]*schema.Schema{
			"blackhole": {
				Type: schema.TypeInt, Optional: true, Description: "Blackhole the zone (in minute, 0 means infinite)",
			},
			"blacklist_src": {
				Type: schema.TypeInt, Optional: true, Description: "Blacklist-src (in min) (applied only for source action)",
			},
			"execute_script": {
				Type: schema.TypeString, Optional: true, Description: "Specify DDOS script to run (applied only for zone action)",
			},
			"execute_script_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Timeout for script execution (in seconds) (applied only for zone action)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS violation-actions name",
			},
			"notification": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_template": {
							Type: schema.TypeString, Optional: true, Description: "Specify the notification template name",
						},
					},
				},
			},
			"send_notification_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forces TPS to only send out notification for the violation-action",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosViolationActionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosViolationActionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosViolationActions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosViolationActionsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosViolationActionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosViolationActionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosViolationActions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosViolationActionsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosViolationActionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosViolationActionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosViolationActions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosViolationActionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosViolationActionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosViolationActions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosViolationActionsNotification(d []interface{}) []edpt.DdosViolationActionsNotification {

	count1 := len(d)
	ret := make([]edpt.DdosViolationActionsNotification, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosViolationActionsNotification
		oi.NotificationTemplate = in["notification_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosViolationActions(d *schema.ResourceData) edpt.DdosViolationActions {
	var ret edpt.DdosViolationActions
	ret.Inst.Blackhole = d.Get("blackhole").(int)
	ret.Inst.BlacklistSrc = d.Get("blacklist_src").(int)
	ret.Inst.ExecuteScript = d.Get("execute_script").(string)
	ret.Inst.ExecuteScriptTimeout = d.Get("execute_script_timeout").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Notification = getSliceDdosViolationActionsNotification(d.Get("notification").([]interface{}))
	ret.Inst.SendNotificationOnly = d.Get("send_notification_only").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
