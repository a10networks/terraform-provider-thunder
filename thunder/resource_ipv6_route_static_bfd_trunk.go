package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RouteStaticBfdTrunk() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_route_static_bfd_trunk`: Trunk interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RouteStaticBfdTrunkCreate,
		UpdateContext: resourceIpv6RouteStaticBfdTrunkUpdate,
		ReadContext:   resourceIpv6RouteStaticBfdTrunkRead,
		DeleteContext: resourceIpv6RouteStaticBfdTrunkDelete,

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
			"trunk_num": {
				Type: schema.TypeInt, Required: true, Description: "Trunk interface",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6RouteStaticBfdTrunkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdTrunkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdTrunk(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdTrunkRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdTrunkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdTrunkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdTrunk(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdTrunkRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RouteStaticBfdTrunkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdTrunkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdTrunk(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteStaticBfdTrunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdTrunkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdTrunk(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteStaticBfdTrunk(d *schema.ResourceData) edpt.Ipv6RouteStaticBfdTrunk {
	var ret edpt.Ipv6RouteStaticBfdTrunk
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.NexthopIpv6Ll = d.Get("nexthop_ipv6_ll").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	ret.Inst.TrunkNum = d.Get("trunk_num").(int)
	//omit uuid
	return ret
}
