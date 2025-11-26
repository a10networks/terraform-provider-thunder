package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemConfigMgmtNotification() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_config_mgmt_notification`: cm notification cconfigure option\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemConfigMgmtNotificationCreate,
		UpdateContext: resourceSystemConfigMgmtNotificationUpdate,
		ReadContext:   resourceSystemConfigMgmtNotificationRead,
		DeleteContext: resourceSystemConfigMgmtNotificationDelete,

		Schema: map[string]*schema.Schema{
			"period": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "Time interval (seconds) for kafka notification. Default is 15 seconds.",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemConfigMgmtNotificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtNotificationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtNotification(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtNotificationRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemConfigMgmtNotificationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtNotificationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtNotification(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtNotificationRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemConfigMgmtNotificationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtNotificationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtNotification(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemConfigMgmtNotificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtNotificationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtNotification(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemConfigMgmtNotification(d *schema.ResourceData) edpt.SystemConfigMgmtNotification {
	var ret edpt.SystemConfigMgmtNotification
	ret.Inst.Period = d.Get("period").(int)
	//omit uuid
	return ret
}
