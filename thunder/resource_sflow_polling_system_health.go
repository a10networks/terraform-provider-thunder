package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowPollingSystemHealth() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_polling_system_health`: Poll system-health counters\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowPollingSystemHealthCreate,
		UpdateContext: resourceSflowPollingSystemHealthUpdate,
		ReadContext:   resourceSflowPollingSystemHealthRead,
		DeleteContext: resourceSflowPollingSystemHealthDelete,

		Schema: map[string]*schema.Schema{
			"license_statistics": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling license statistics; 'disable': Disable polling license statistics;",
			},
			"per_control_cpu_usage": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling control cpu; 'disable': Disable polling control cpu usage;",
			},
			"per_data_cpu_usage": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling data cpu; 'disable': Disable polling data cpu usage;",
			},
			"system_health_usage": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable polling system health information; 'disable': Disable polling system health information;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowPollingSystemHealthCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingSystemHealthCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingSystemHealth(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingSystemHealthRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowPollingSystemHealthUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingSystemHealthUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingSystemHealth(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingSystemHealthRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowPollingSystemHealthDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingSystemHealthDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingSystemHealth(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowPollingSystemHealthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingSystemHealthRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingSystemHealth(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowPollingSystemHealth(d *schema.ResourceData) edpt.SflowPollingSystemHealth {
	var ret edpt.SflowPollingSystemHealth
	ret.Inst.LicenseStatistics = d.Get("license_statistics").(string)
	ret.Inst.PerControlCpuUsage = d.Get("per_control_cpu_usage").(string)
	ret.Inst.PerDataCpuUsage = d.Get("per_data_cpu_usage").(string)
	ret.Inst.SystemHealthUsage = d.Get("system_health_usage").(string)
	//omit uuid
	return ret
}
