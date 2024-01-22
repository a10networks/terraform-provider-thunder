package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicyDualStackAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_dual_stack_action`: Action to forward requests via IPv4 or IPv6 networks\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicyDualStackActionCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicyDualStackActionUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicyDualStackActionRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicyDualStackActionDelete,

		Schema: map[string]*schema.Schema{
			"fall_back": {
				Type: schema.TypeString, Optional: true, Description: "Fallback service group",
			},
			"fall_back_snat": {
				Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group for fallback",
			},
			"ipv4": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 service group to forward",
			},
			"ipv4_snat": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 source NAT pool or pool group",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 service group to forward",
			},
			"ipv6_snat": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 source NAT pool or pool group",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable logging",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Action name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests forward by this action;",
						},
					},
				},
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
func resourceSlbTemplatePolicyForwardPolicyDualStackActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyDualStackActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyDualStackAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicyDualStackActionRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicyDualStackActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyDualStackActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyDualStackAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicyDualStackActionRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicyDualStackActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyDualStackActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyDualStackAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicyDualStackActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyDualStackActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyDualStackAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplatePolicyForwardPolicyDualStackActionSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyDualStackActionSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyDualStackActionSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyDualStackActionSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicyDualStackAction(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicyDualStackAction {
	var ret edpt.SlbTemplatePolicyForwardPolicyDualStackAction
	ret.Inst.FallBack = d.Get("fall_back").(string)
	ret.Inst.FallBackSnat = d.Get("fall_back_snat").(string)
	ret.Inst.Ipv4 = d.Get("ipv4").(string)
	ret.Inst.Ipv4Snat = d.Get("ipv4_snat").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	ret.Inst.Ipv6Snat = d.Get("ipv6_snat").(string)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicyDualStackActionSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
