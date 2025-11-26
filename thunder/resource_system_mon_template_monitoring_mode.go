package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMonTemplateMonitoringMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_mon_template_monitoring_mode`: Monitoring behavior mode interdependent/and\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMonTemplateMonitoringModeCreate,
		UpdateContext: resourceSystemMonTemplateMonitoringModeUpdate,
		ReadContext:   resourceSystemMonTemplateMonitoringModeRead,
		DeleteContext: resourceSystemMonTemplateMonitoringModeDelete,

		Schema: map[string]*schema.Schema{
			"mmode": {
				Type: schema.TypeString, Optional: true, Default: "and", Description: "'interdependent': INTERDEPENDENT monitoring behaviour; 'and': AND monitoring behaviour, Default;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemMonTemplateMonitoringModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitoringModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitoringMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateMonitoringModeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMonTemplateMonitoringModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitoringModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitoringMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMonTemplateMonitoringModeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMonTemplateMonitoringModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitoringModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitoringMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMonTemplateMonitoringModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMonTemplateMonitoringModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMonTemplateMonitoringMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemMonTemplateMonitoringMode(d *schema.ResourceData) edpt.SystemMonTemplateMonitoringMode {
	var ret edpt.SystemMonTemplateMonitoringMode
	ret.Inst.Mmode = d.Get("mmode").(string)
	//omit uuid
	return ret
}
