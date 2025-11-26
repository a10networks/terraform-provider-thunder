package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugRateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_rate_limit`: Debug Rate Limiting\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugRateLimitCreate,
		UpdateContext: resourceDebugRateLimitUpdate,
		ReadContext:   resourceDebugRateLimitRead,
		DeleteContext: resourceDebugRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"cpu_id": {
				Type: schema.TypeInt, Optional: true, Description: "Starting From 0",
			},
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"entry_ip": {
				Type: schema.TypeString, Optional: true, Description: "Source Address",
			},
			"life_cycle_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print entry life cycle info",
			},
			"metric": {
				Type: schema.TypeInt, Optional: true, Description: "CPS 0 Uplink 1 Downlink 2 Total 3",
			},
			"per_pkt_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print entry info based on code path of each pkt",
			},
			"raw_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print entry info per second",
			},
			"tpl_id": {
				Type: schema.TypeInt, Optional: true, Description: "Template ID",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugRateLimit(d *schema.ResourceData) edpt.DebugRateLimit {
	var ret edpt.DebugRateLimit
	ret.Inst.CpuId = d.Get("cpu_id").(int)
	ret.Inst.Dumy = d.Get("dumy").(int)
	ret.Inst.EntryIp = d.Get("entry_ip").(string)
	ret.Inst.LifeCycleLog = d.Get("life_cycle_log").(int)
	ret.Inst.Metric = d.Get("metric").(int)
	ret.Inst.PerPktLog = d.Get("per_pkt_log").(int)
	ret.Inst.RawLog = d.Get("raw_log").(int)
	ret.Inst.TplId = d.Get("tpl_id").(int)
	//omit uuid
	return ret
}
