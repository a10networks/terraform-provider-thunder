package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RouteStaticBfdVe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_route_static_bfd_ve`: Virtual ethernet interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RouteStaticBfdVeCreate,
		UpdateContext: resourceIpv6RouteStaticBfdVeUpdate,
		ReadContext:   resourceIpv6RouteStaticBfdVeRead,
		DeleteContext: resourceIpv6RouteStaticBfdVeDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'down': BFD down;  (BFD state)",
			},
			"nexthop_ipv6_ll": {
				Type: schema.TypeString, Required: true, Description: "Nexthop IPv6 address (Link-local)",
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
			"ve_num": {
				Type: schema.TypeInt, Required: true, Description: "Virtual ethernet interface",
			},
		},
	}
}
func resourceIpv6RouteStaticBfdVeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdVeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdVe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdVeRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdVeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdVeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdVe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdVeRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RouteStaticBfdVeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdVeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdVe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteStaticBfdVeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdVeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdVe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteStaticBfdVe(d *schema.ResourceData) edpt.Ipv6RouteStaticBfdVe {
	var ret edpt.Ipv6RouteStaticBfdVe
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.NexthopIpv6Ll = d.Get("nexthop_ipv6_ll").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	ret.Inst.VeNum = d.Get("ve_num").(int)
	return ret
}
