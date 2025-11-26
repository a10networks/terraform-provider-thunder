package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectNotification() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_notification`: DDOS Network-object Notification\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectNotificationCreate,
		UpdateContext: resourceDdosNetworkObjectNotificationUpdate,
		ReadContext:   resourceDdosNetworkObjectNotificationRead,
		DeleteContext: resourceDdosNetworkObjectNotificationDelete,

		Schema: map[string]*schema.Schema{
			"configuration": {
				Type: schema.TypeString, Required: true, Description: "'configuration': configuration;",
			},
			"notification": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification_template_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the notification template name",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectNotificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectNotificationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectNotification(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectNotificationRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectNotificationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectNotificationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectNotification(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectNotificationRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectNotificationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectNotificationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectNotification(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectNotificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectNotificationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectNotification(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosNetworkObjectNotificationNotification(d []interface{}) []edpt.DdosNetworkObjectNotificationNotification {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectNotificationNotification, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectNotificationNotification
		oi.NotificationTemplateName = in["notification_template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectNotification(d *schema.ResourceData) edpt.DdosNetworkObjectNotification {
	var ret edpt.DdosNetworkObjectNotification
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.Notification = getSliceDdosNetworkObjectNotificationNotification(d.Get("notification").([]interface{}))
	//omit uuid
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
