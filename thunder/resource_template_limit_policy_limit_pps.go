package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateLimitPolicyLimitPps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_limit_policy_limit_pps`: Enable Packets Per Second Rate Limit\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateLimitPolicyLimitPpsCreate,
		UpdateContext: resourceTemplateLimitPolicyLimitPpsUpdate,
		ReadContext:   resourceTemplateLimitPolicyLimitPpsRead,
		DeleteContext: resourceTemplateLimitPolicyLimitPpsDelete,

		Schema: map[string]*schema.Schema{
			"ddos_protection_factor": {
				Type: schema.TypeInt, Optional: true, Description: "Enable DDoS Protection (Multiplier of the downlink PPS)",
			},
			"downlink": {
				Type: schema.TypeInt, Optional: true, Description: "Downlink PPS limit (Number of Packets per second)",
			},
			"downlink_burstsize": {
				Type: schema.TypeInt, Optional: true, Description: "PPS Token Bucket Size (Must Exceed Configured Rate) (In Packets)",
			},
			"downlink_relaxed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
			},
			"total": {
				Type: schema.TypeInt, Optional: true, Description: "Total PPS limit (Number of Packets per second)",
			},
			"total_burstsize": {
				Type: schema.TypeInt, Optional: true, Description: "PPS Token Bucket Size (Must Exceed Configured Rate) (In Packets)",
			},
			"total_relaxed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
			},
			"uplink": {
				Type: schema.TypeInt, Optional: true, Description: "Uplink PPS limit (Number of Packets per second)",
			},
			"uplink_burstsize": {
				Type: schema.TypeInt, Optional: true, Description: "PPS Token Bucket Size (Must Exceed Configured Rate) (In Packets)",
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
func resourceTemplateLimitPolicyLimitPpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitPpsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitPps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyLimitPpsRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateLimitPolicyLimitPpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitPpsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitPps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyLimitPpsRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateLimitPolicyLimitPpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitPpsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitPps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateLimitPolicyLimitPpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitPpsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitPps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateLimitPolicyLimitPps(d *schema.ResourceData) edpt.TemplateLimitPolicyLimitPps {
	var ret edpt.TemplateLimitPolicyLimitPps
	ret.Inst.DdosProtectionFactor = d.Get("ddos_protection_factor").(int)
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
