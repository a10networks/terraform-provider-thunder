package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionSettingsStandaloneSettings() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_settings_standalone_settings`: Configure ddos standalone detection settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionSettingsStandaloneSettingsCreate,
		UpdateContext: resourceDdosDetectionSettingsStandaloneSettingsUpdate,
		ReadContext:   resourceDdosDetectionSettingsStandaloneSettingsRead,
		DeleteContext: resourceDdosDetectionSettingsStandaloneSettingsDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable standalone detector; 'disable': Disable standalone detector (default);",
			},
			"de_escalation_quiet_time": {
				Type: schema.TypeInt, Optional: true, Description: "Configure de-escalation needed time in minutes from level 1 to 0.(legacy)",
			},
			"netflow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"listening_port": {
							Type: schema.TypeInt, Optional: true, Default: 9996, Description: "Netflow port to receive packets (Netflow port number(default 9996))",
						},
						"template_active_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Configure active timeout of the netflow templates received in mins (Template active timeout(mins)(default 30mins))",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sflow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"listening_port": {
							Type: schema.TypeInt, Optional: true, Default: 6343, Description: "sFlow port to receive packets (sFlow port number(default 6343))",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDetectionSettingsStandaloneSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettings(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsStandaloneSettingsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionSettingsStandaloneSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettings(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsStandaloneSettingsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionSettingsStandaloneSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettings(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionSettingsStandaloneSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettings(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDetectionSettingsStandaloneSettingsNetflow130(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsNetflow130 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsNetflow130
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		ret.TemplateActiveTimeout = in["template_active_timeout"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettingsSflow131(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsSflow131 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsSflow131
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosDetectionSettingsStandaloneSettings(d *schema.ResourceData) edpt.DdosDetectionSettingsStandaloneSettings {
	var ret edpt.DdosDetectionSettingsStandaloneSettings
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DeEscalationQuietTime = d.Get("de_escalation_quiet_time").(int)
	ret.Inst.Netflow = getObjectDdosDetectionSettingsStandaloneSettingsNetflow130(d.Get("netflow").([]interface{}))
	ret.Inst.Sflow = getObjectDdosDetectionSettingsStandaloneSettingsSflow131(d.Get("sflow").([]interface{}))
	//omit uuid
	return ret
}
