package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateLimitPolicyLimitThroughput() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_limit_policy_limit_throughput`: Enable Throughput Rate Limit\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateLimitPolicyLimitThroughputCreate,
		UpdateContext: resourceTemplateLimitPolicyLimitThroughputUpdate,
		ReadContext:   resourceTemplateLimitPolicyLimitThroughputRead,
		DeleteContext: resourceTemplateLimitPolicyLimitThroughputDelete,

		Schema: map[string]*schema.Schema{
			"downlink": {
				Type: schema.TypeInt, Optional: true, Description: "Downlink Throughput limit (Mega Bits per second)",
			},
			"downlink_burstsize": {
				Type: schema.TypeInt, Optional: true, Description: "Token Bucket Size (Must Exceed Configured Rate) (In Mega Bits per second)",
			},
			"downlink_relaxed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
			},
			"total": {
				Type: schema.TypeInt, Optional: true, Description: "Total Throughput limit (Mega Bits per second)",
			},
			"total_burstsize": {
				Type: schema.TypeInt, Optional: true, Description: "Token Bucket Size (Must Exceed Configured Rate) (In Mega Bits per second)",
			},
			"total_relaxed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
			},
			"uplink": {
				Type: schema.TypeInt, Optional: true, Description: "Uplink Throughput limit (Mega Bits per second)",
			},
			"uplink_burstsize": {
				Type: schema.TypeInt, Optional: true, Description: "Token Bucket Size (Must Exceed Configured Rate) (In Mega Bits per second)",
			},
			"uplink_relaxed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"policy_number": {
				Type: schema.TypeString, Required: true, Description: "PolicyNumber",
			},
		},
	}
}
func resourceTemplateLimitPolicyLimitThroughputCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitThroughputCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitThroughput(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyLimitThroughputRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateLimitPolicyLimitThroughputUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitThroughputUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitThroughput(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyLimitThroughputRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateLimitPolicyLimitThroughputDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitThroughputDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitThroughput(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateLimitPolicyLimitThroughputRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitThroughputRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitThroughput(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateLimitPolicyLimitThroughput(d *schema.ResourceData) edpt.TemplateLimitPolicyLimitThroughput {
	var ret edpt.TemplateLimitPolicyLimitThroughput
	ret.Inst.Downlink = d.Get("downlink").(int)
	ret.Inst.DownlinkBurstsize = d.Get("downlink_burstsize").(int)
	ret.Inst.DownlinkRelaxed = d.Get("downlink_relaxed").(int)
	ret.Inst.Total = d.Get("total").(int)
	ret.Inst.TotalBurstsize = d.Get("total_burstsize").(int)
	ret.Inst.TotalRelaxed = d.Get("total_relaxed").(int)
	ret.Inst.Uplink = d.Get("uplink").(int)
	ret.Inst.UplinkBurstsize = d.Get("uplink_burstsize").(int)
	ret.Inst.UplinkRelaxed = d.Get("uplink_relaxed").(int)
	//omit uuid
	ret.Inst.PolicyNumber = d.Get("policy_number").(string)
	return ret
}
