package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAny() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_source_destination_any`: Default destination match rule\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "Action to be performed",
			},
			"dual_stack_action": {
				Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests matching this destination rule;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAny(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAny(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAny(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAny(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAny(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationAny {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationAny
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DualStackAction = d.Get("dual_stack_action").(string)
	ret.Inst.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
