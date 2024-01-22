package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateLimitPolicyLimitCps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_limit_policy_limit_cps`: Enable Connections Per Second Rate Limit\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateLimitPolicyLimitCpsCreate,
		UpdateContext: resourceTemplateLimitPolicyLimitCpsUpdate,
		ReadContext:   resourceTemplateLimitPolicyLimitCpsRead,
		DeleteContext: resourceTemplateLimitPolicyLimitCpsDelete,

		Schema: map[string]*schema.Schema{
			"burstsize": {
				Type: schema.TypeInt, Optional: true, Description: "CPS Token Bucket Size (Must Exceed Configured Rate) (In Connections per second)",
			},
			"relaxed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Relax the limitation when the policy has more tokens from the parent of policy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Connections Per Second Rate Limit (Number of Connections per second)",
			},
			"policy_number": {
				Type: schema.TypeString, Required: true, Description: "PolicyNumber",
			},
		},
	}
}
func resourceTemplateLimitPolicyLimitCpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitCpsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitCps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyLimitCpsRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateLimitPolicyLimitCpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitCpsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitCps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateLimitPolicyLimitCpsRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateLimitPolicyLimitCpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitCpsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitCps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateLimitPolicyLimitCpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateLimitPolicyLimitCpsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateLimitPolicyLimitCps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateLimitPolicyLimitCps(d *schema.ResourceData) edpt.TemplateLimitPolicyLimitCps {
	var ret edpt.TemplateLimitPolicyLimitCps
	ret.Inst.Burstsize = d.Get("burstsize").(int)
	ret.Inst.Relaxed = d.Get("relaxed").(int)
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.PolicyNumber = d.Get("policy_number").(string)
	return ret
}
