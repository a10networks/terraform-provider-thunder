package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RouteStaticBfdBfdIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_route_static_bfd_bfd_ipv6`: Bidirectional Forwarding Detection\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RouteStaticBfdBfdIpv6Create,
		UpdateContext: resourceIpv6RouteStaticBfdBfdIpv6Update,
		ReadContext:   resourceIpv6RouteStaticBfdBfdIpv6Read,
		DeleteContext: resourceIpv6RouteStaticBfdBfdIpv6Delete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'down': BFD down;  (BFD state)",
			},
			"local_ipv6": {
				Type: schema.TypeString, Required: true, Description: "Local IPv6 address",
			},
			"nexthop_ipv6": {
				Type: schema.TypeString, Required: true, Description: "Nexthop IPv6 address",
			},
			"template": {
				Type: schema.TypeString, Optional: true, Description: "Configure tracking template (bind tracking template name)",
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6RouteStaticBfdBfdIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdBfdIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdBfdIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdBfdIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdBfdIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdBfdIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RouteStaticBfdBfdIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdBfdIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteStaticBfdBfdIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdBfdIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteStaticBfdBfdIpv6(d *schema.ResourceData) edpt.Ipv6RouteStaticBfdBfdIpv6 {
	var ret edpt.Ipv6RouteStaticBfdBfdIpv6
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.LocalIpv6 = d.Get("local_ipv6").(string)
	ret.Inst.NexthopIpv6 = d.Get("nexthop_ipv6").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	return ret
}
