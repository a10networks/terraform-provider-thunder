package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionSettingsZoneNotifications() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_settings_zone_notifications`: Configure ddos zone detection notification settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionSettingsZoneNotificationsCreate,
		UpdateContext: resourceDdosDetectionSettingsZoneNotificationsUpdate,
		ReadContext:   resourceDdosDetectionSettingsZoneNotificationsRead,
		DeleteContext: resourceDdosDetectionSettingsZoneNotificationsDelete,

		Schema: map[string]*schema.Schema{
			"source_entry": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable source entry detection notification; 'disable': Disable source entry detection notification(default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDetectionSettingsZoneNotificationsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsZoneNotificationsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsZoneNotifications(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsZoneNotificationsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionSettingsZoneNotificationsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsZoneNotificationsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsZoneNotifications(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsZoneNotificationsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionSettingsZoneNotificationsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsZoneNotificationsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsZoneNotifications(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionSettingsZoneNotificationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsZoneNotificationsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsZoneNotifications(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionSettingsZoneNotifications(d *schema.ResourceData) edpt.DdosDetectionSettingsZoneNotifications {
	var ret edpt.DdosDetectionSettingsZoneNotifications
	ret.Inst.SourceEntry = d.Get("source_entry").(string)
	//omit uuid
	return ret
}
