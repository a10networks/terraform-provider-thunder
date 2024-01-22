package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionSettingsStandaloneSettingsSflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_settings_standalone_settings_sflow`: Configure sFlow parameters for flow based detection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionSettingsStandaloneSettingsSflowCreate,
		UpdateContext: resourceDdosDetectionSettingsStandaloneSettingsSflowUpdate,
		ReadContext:   resourceDdosDetectionSettingsStandaloneSettingsSflowRead,
		DeleteContext: resourceDdosDetectionSettingsStandaloneSettingsSflowDelete,

		Schema: map[string]*schema.Schema{
			"listening_port": {
				Type: schema.TypeInt, Optional: true, Default: 6343, Description: "sFlow port to receive packets (sFlow port number(default 6343))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDetectionSettingsStandaloneSettingsSflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsSflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsSflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsStandaloneSettingsSflowRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionSettingsStandaloneSettingsSflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsSflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsSflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsStandaloneSettingsSflowRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionSettingsStandaloneSettingsSflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsSflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsSflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionSettingsStandaloneSettingsSflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsSflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsSflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionSettingsStandaloneSettingsSflow(d *schema.ResourceData) edpt.DdosDetectionSettingsStandaloneSettingsSflow {
	var ret edpt.DdosDetectionSettingsStandaloneSettingsSflow
	ret.Inst.ListeningPort = d.Get("listening_port").(int)
	//omit uuid
	return ret
}
