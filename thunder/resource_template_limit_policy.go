package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateLimitPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_limit_policy`: Create a Limit Policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateLimitPolicyCreate,
		UpdateContext: resourceTemplateLimitPolicyUpdate,
		ReadContext:   resourceTemplateLimitPolicyRead,
		DeleteContext: resourceTemplateLimitPolicyDelete,

		Schema: map[string]*schema.Schema{
			"limit_concurrent_sessions": {
				Type: schema.TypeInt, Optional: true, Description: "Enable Concurrent Session Limit (Number of Concurrent Sessions)",
			},
			"limit_cps": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeInt, Optional: true, Description: "Connections Per Second Rate Limit (Number of Connections per second)",
						},
						"burstsize": {
							Type: schema.TypeInt, Optional: true, Description: "CPS Token Bucket Size (Must Exceed Configured Rate) (In Connections per second)",
						},
						"relaxed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"limit_pps": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uplink": {
							Type: schema.TypeInt, Optional: true, Description: "Uplink PPS limit (Number of Packets per second)",
						},
						"uplink_burstsize": {
							Type: schema.TypeInt, Optional: true, Description: "PPS Token Bucket Size (Must Exceed Configured Rate) (In Packets)",
						},
						"uplink_relaxed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
						},
						"downlink": {
							Type: schema.TypeInt, Optional: true, Description: "Downlink PPS limit (Number of Packets per second)",
						},
						"ddos_protection_factor": {
							Type: schema.TypeInt, Optional: true, Description: "Enable DDoS Protection (Multiplier of the downlink PPS)",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"limit_scope": {
				Type: schema.TypeString, Optional: true, Default: "subscriber-ip", Description: "'aggregate': Rule Level; 'subscriber-ip': Subscriber IP Level; 'subscriber-prefix': Subscriber Prefix Level;",
			},
			"limit_throughput": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uplink": {
							Type: schema.TypeInt, Optional: true, Description: "Uplink Throughput limit (Mega Bits per second)",
						},
						"uplink_burstsize": {
							Type: schema.TypeInt, Optional: true, Description: "Token Bucket Size (Must Exceed Configured Rate) (In Mega Bits per second)",
						},
						"uplink_relaxed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
						},
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when Session Limit is exceeded",
			},
			"max_min_fair": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable max-min-fairness",
			},
			"parent": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the parent of limit-policy",
			},
			"policy_number": {
				Type: schema.TypeInt, Required: true, Description: "Limit Policy Number",
			},
			"prefix_length": {
				Type: schema.TypeInt, Optional: true, Description: "Prefix length",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTemplateLimitPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateLimitPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateLimitPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateLimitPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTemplateLimitPolicyLimitCps1903(d []interface{}) edpt.TemplateLimitPolicyLimitCps1903 {

	count1 := len(d)
	var ret edpt.TemplateLimitPolicyLimitCps1903
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Burstsize = in["burstsize"].(int)
		ret.Relaxed = in["relaxed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectTemplateLimitPolicyLimitPps1904(d []interface{}) edpt.TemplateLimitPolicyLimitPps1904 {

	count1 := len(d)
	var ret edpt.TemplateLimitPolicyLimitPps1904
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Uplink = in["uplink"].(int)
		ret.UplinkBurstsize = in["uplink_burstsize"].(int)
		ret.UplinkRelaxed = in["uplink_relaxed"].(int)
		ret.Downlink = in["downlink"].(int)
		ret.DdosProtectionFactor = in["ddos_protection_factor"].(int)
		ret.DownlinkBurstsize = in["downlink_burstsize"].(int)
		ret.DownlinkRelaxed = in["downlink_relaxed"].(int)
		ret.Total = in["total"].(int)
		ret.TotalBurstsize = in["total_burstsize"].(int)
		ret.TotalRelaxed = in["total_relaxed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectTemplateLimitPolicyLimitThroughput1905(d []interface{}) edpt.TemplateLimitPolicyLimitThroughput1905 {

	count1 := len(d)
	var ret edpt.TemplateLimitPolicyLimitThroughput1905
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Uplink = in["uplink"].(int)
		ret.UplinkBurstsize = in["uplink_burstsize"].(int)
		ret.UplinkRelaxed = in["uplink_relaxed"].(int)
		ret.Downlink = in["downlink"].(int)
		ret.DownlinkBurstsize = in["downlink_burstsize"].(int)
		ret.DownlinkRelaxed = in["downlink_relaxed"].(int)
		ret.Total = in["total"].(int)
		ret.TotalBurstsize = in["total_burstsize"].(int)
		ret.TotalRelaxed = in["total_relaxed"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointTemplateLimitPolicy(d *schema.ResourceData) edpt.TemplateLimitPolicy {
	var ret edpt.TemplateLimitPolicy
	ret.Inst.LimitConcurrentSessions = d.Get("limit_concurrent_sessions").(int)
	ret.Inst.LimitCps = getObjectTemplateLimitPolicyLimitCps1903(d.Get("limit_cps").([]interface{}))
	ret.Inst.LimitPps = getObjectTemplateLimitPolicyLimitPps1904(d.Get("limit_pps").([]interface{}))
	ret.Inst.LimitScope = d.Get("limit_scope").(string)
	ret.Inst.LimitThroughput = getObjectTemplateLimitPolicyLimitThroughput1905(d.Get("limit_throughput").([]interface{}))
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.MaxMinFair = d.Get("max_min_fair").(int)
	ret.Inst.Parent = d.Get("parent").(int)
	ret.Inst.PolicyNumber = d.Get("policy_number").(int)
	ret.Inst.PrefixLength = d.Get("prefix_length").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
