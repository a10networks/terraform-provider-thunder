package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionSettingsStandaloneSettingsNetflow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_settings_standalone_settings_netflow`: Configure Netflow parameters for flow based detection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionSettingsStandaloneSettingsNetflowCreate,
		UpdateContext: resourceDdosDetectionSettingsStandaloneSettingsNetflowUpdate,
		ReadContext:   resourceDdosDetectionSettingsStandaloneSettingsNetflowRead,
		DeleteContext: resourceDdosDetectionSettingsStandaloneSettingsNetflowDelete,

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
	}
}
func resourceDdosDetectionSettingsStandaloneSettingsNetflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsNetflowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsNetflow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsStandaloneSettingsNetflowRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionSettingsStandaloneSettingsNetflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsNetflowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsNetflow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsStandaloneSettingsNetflowRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionSettingsStandaloneSettingsNetflowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsNetflowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsNetflow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionSettingsStandaloneSettingsNetflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsStandaloneSettingsNetflowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettingsStandaloneSettingsNetflow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionSettingsStandaloneSettingsNetflow(d *schema.ResourceData) edpt.DdosDetectionSettingsStandaloneSettingsNetflow {
	var ret edpt.DdosDetectionSettingsStandaloneSettingsNetflow
	ret.Inst.ListeningPort = d.Get("listening_port").(int)
	ret.Inst.TemplateActiveTimeout = d.Get("template_active_timeout").(int)
	//omit uuid
	return ret
}
