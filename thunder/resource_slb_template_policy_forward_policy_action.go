package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicyAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_action`: action list\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicyActionCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicyActionUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicyActionRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicyActionDelete,

		Schema: map[string]*schema.Schema{
			"action1": {
				Type: schema.TypeString, Optional: true, Description: "'forward-to-internet': Forward request to Internet; 'forward-to-service-group': Forward request to service group; 'forward-to-proxy': Forward request to HTTP proxy server; 'drop': Drop request;",
			},
			"drop_message": {
				Type: schema.TypeString, Optional: true, Description: "drop-message sent to the client as webpage(html tags are included and quotation marks are required for white spaces)",
			},
			"drop_redirect_url": {
				Type: schema.TypeString, Optional: true, Description: "Specify URL to which client request is redirected upon being dropped",
			},
			"drop_response_code": {
				Type: schema.TypeInt, Optional: true, Description: "Specify response code for drop action",
			},
			"fake_sg": {
				Type: schema.TypeString, Optional: true, Description: "service group to forward the packets to Internet",
			},
			"fall_back": {
				Type: schema.TypeString, Optional: true, Description: "Fallback service group for Internet",
			},
			"fall_back_snat": {
				Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group for fallback server",
			},
			"fall_back_snat_pt_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source port translation only for fallback server",
			},
			"forward_snat": {
				Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group",
			},
			"forward_snat_pt_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source port translation only",
			},
			"http_status_code": {
				Type: schema.TypeString, Optional: true, Default: "302", Description: "'301': Moved permanently; '302': Found;",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable logging",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Action policy name",
			},
			"proxy_chaining": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable proxy chaining feature",
			},
			"proxy_chaining_bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward all https packets to upstream proxy",
			},
			"real_sg": {
				Type: schema.TypeString, Optional: true, Description: "service group to forward the packets",
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
			"support_cert_fetch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fetch server certificate by upstream proxy",
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
func resourceSlbTemplatePolicyForwardPolicyActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicyActionRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicyActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicyActionRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicyActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicyActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplatePolicyForwardPolicyActionSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyActionSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyActionSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyActionSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicyAction(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicyAction {
	var ret edpt.SlbTemplatePolicyForwardPolicyAction
	ret.Inst.Action1 = d.Get("action1").(string)
	ret.Inst.DropMessage = d.Get("drop_message").(string)
	ret.Inst.DropRedirectUrl = d.Get("drop_redirect_url").(string)
	ret.Inst.DropResponseCode = d.Get("drop_response_code").(int)
	ret.Inst.FakeSg = d.Get("fake_sg").(string)
	ret.Inst.FallBack = d.Get("fall_back").(string)
	ret.Inst.FallBackSnat = d.Get("fall_back_snat").(string)
	ret.Inst.FallBackSnatPtOnly = d.Get("fall_back_snat_pt_only").(int)
	ret.Inst.ForwardSnat = d.Get("forward_snat").(string)
	ret.Inst.ForwardSnatPtOnly = d.Get("forward_snat_pt_only").(int)
	ret.Inst.HttpStatusCode = d.Get("http_status_code").(string)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ProxyChaining = d.Get("proxy_chaining").(int)
	ret.Inst.ProxyChainingBypass = d.Get("proxy_chaining_bypass").(int)
	ret.Inst.RealSg = d.Get("real_sg").(string)
	ret.Inst.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicyActionSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SupportCertFetch = d.Get("support_cert_fetch").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
