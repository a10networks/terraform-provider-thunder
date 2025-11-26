package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyDnsStickyOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_dns_sticky_options`: Specify sticky session options\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyDnsStickyOptionsCreate,
		UpdateContext: resourceGslbPolicyDnsStickyOptionsUpdate,
		ReadContext:   resourceGslbPolicyDnsStickyOptionsRead,
		DeleteContext: resourceGslbPolicyDnsStickyOptionsDelete,

		Schema: map[string]*schema.Schema{
			"edns_client_subnet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use ECS for sticky creation and lookup",
			},
			"only_ecs": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only use ECS for session creation",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"policy_name": {
				Type: schema.TypeString, Required: true, Description: "Policy_name",
			},
		},
	}
}
func resourceGslbPolicyDnsStickyOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsStickyOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDnsStickyOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyDnsStickyOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyDnsStickyOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsStickyOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDnsStickyOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyDnsStickyOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyDnsStickyOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsStickyOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDnsStickyOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyDnsStickyOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyDnsStickyOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyDnsStickyOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyDnsStickyOptions(d *schema.ResourceData) edpt.GslbPolicyDnsStickyOptions {
	var ret edpt.GslbPolicyDnsStickyOptions
	ret.Inst.EdnsClientSubnet = d.Get("edns_client_subnet").(int)
	ret.Inst.OnlyEcs = d.Get("only_ecs").(int)
	//omit uuid
	ret.Inst.Policy_name = d.Get("policy_name").(string)
	return ret
}
