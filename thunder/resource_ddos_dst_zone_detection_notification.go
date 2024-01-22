package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionNotification() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_notification`: DDOS Detection Notification\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionNotificationCreate,
		UpdateContext: resourceDdosDstZoneDetectionNotificationUpdate,
		ReadContext:   resourceDdosDstZoneDetectionNotificationRead,
		DeleteContext: resourceDdosDstZoneDetectionNotificationDelete,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneDetectionNotificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionNotificationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionNotification(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionNotificationRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionNotificationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionNotificationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionNotification(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionNotificationRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionNotificationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionNotificationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionNotification(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionNotificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionNotificationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionNotification(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneDetectionNotificationNotification(d []interface{}) []edpt.DdosDstZoneDetectionNotificationNotification {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionNotificationNotification, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionNotificationNotification
		oi.NotificationTemplateName = in["notification_template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetectionNotification(d *schema.ResourceData) edpt.DdosDstZoneDetectionNotification {
	var ret edpt.DdosDstZoneDetectionNotification
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.Notification = getSliceDdosDstZoneDetectionNotificationNotification(d.Get("notification").([]interface{}))
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
