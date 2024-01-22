package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionSettingsEntrySaving() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_settings_entry_saving`: Detection entries and indicators saving configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionSettingsEntrySavingCreate,
		UpdateContext: resourceDdosDetectionSettingsEntrySavingUpdate,
		ReadContext:   resourceDdosDetectionSettingsEntrySavingRead,
		DeleteContext: resourceDdosDetectionSettingsEntrySavingDelete,

		Schema: map[string]*schema.Schema{
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure periodical auto-saving interval in minutes(default: 0) and 0 to disable.",
			},
			"manual_restore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually restore network-object-based detection entries and learned indicators",
			},
			"manual_save": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually save network-object-based detection entries and learned indicators",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDetectionSettingsEntrySavingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsEntrySavingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsEntrySaving(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsEntrySavingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionSettingsEntrySavingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsEntrySavingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsEntrySaving(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsEntrySavingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionSettingsEntrySavingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsEntrySavingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsEntrySaving(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionSettingsEntrySavingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsEntrySavingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsEntrySaving(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionSettingsEntrySaving(d *schema.ResourceData) edpt.DdosDetectionSettingsEntrySaving {
	var ret edpt.DdosDetectionSettingsEntrySaving
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.ManualRestore = d.Get("manual_restore").(int)
	ret.Inst.ManualSave = d.Get("manual_save").(int)
	//omit uuid
	return ret
}
