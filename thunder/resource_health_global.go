package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_global`: Define the Health Monitor global default\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthGlobalCreate,
		UpdateContext: resourceHealthGlobalUpdate,
		ReadContext:   resourceHealthGlobalRead,
		DeleteContext: resourceHealthGlobalDelete,

		Schema: map[string]*schema.Schema{
			"check_rate": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Health Check Rate per 500ms (default 1000)",
			},
			"disable_auto_adjust": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the Health Check Rate Auto Adjustment",
			},
			"external_rate": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Define the External Health Check Rate (Number of External Script Programs (default 2))",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the Healthcheck Interval (Interval Value, in seconds (default 5))",
			},
			"multi_process": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Start Health Monitoring in Multi-Process Mode (Specify the number of multiple processes (default 1))",
			},
			"per": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Specify the Unit Time for the rate (Specify the Unit Time, multiple of 100ms)",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the Healthcheck Retries (Retry Count (default 3))",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the Healthcheck Timeout (Timeout Value, in seconds (default 5), Timeout should be less than or equal to interval)",
			},
			"up_retry": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify the Healthcheck Retries before declaring target up (Up-retry count (default 1))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceHealthGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthGlobal(d *schema.ResourceData) edpt.HealthGlobal {
	var ret edpt.HealthGlobal
	ret.Inst.CheckRate = d.Get("check_rate").(int)
	ret.Inst.DisableAutoAdjust = d.Get("disable_auto_adjust").(int)
	ret.Inst.ExternalRate = d.Get("external_rate").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.MultiProcess = d.Get("multi_process").(int)
	ret.Inst.Per = d.Get("per").(int)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UpRetry = d.Get("up_retry").(int)
	//omit uuid
	return ret
}
