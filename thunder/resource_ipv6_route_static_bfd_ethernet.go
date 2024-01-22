package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RouteStaticBfdEthernet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_route_static_bfd_ethernet`: Ethernet interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RouteStaticBfdEthernetCreate,
		UpdateContext: resourceIpv6RouteStaticBfdEthernetUpdate,
		ReadContext:   resourceIpv6RouteStaticBfdEthernetRead,
		DeleteContext: resourceIpv6RouteStaticBfdEthernetDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'down': BFD down;  (BFD state)",
			},
			"eth_num": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet (not a member of vlan or trunk)",
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
		},
	}
}
func resourceIpv6RouteStaticBfdEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdEthernetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdEthernet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdEthernetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdEthernet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdEthernetRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RouteStaticBfdEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdEthernetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdEthernet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteStaticBfdEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteStaticBfdEthernetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdEthernet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteStaticBfdEthernet(d *schema.ResourceData) edpt.Ipv6RouteStaticBfdEthernet {
	var ret edpt.Ipv6RouteStaticBfdEthernet
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.EthNum = d.Get("eth_num").(int)
	ret.Inst.NexthopIpv6Ll = d.Get("nexthop_ipv6_ll").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	return ret
}
