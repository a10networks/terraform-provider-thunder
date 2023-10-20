package thunder

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"util"
)

func resourceIpv6RouteStaticBfdEthernet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpv6RouteStaticBfdEthernetCreate,
		UpdateContext: resourceIpv6RouteStaticBfdEthernetUpdate,
		ReadContext:   resourceIpv6RouteStaticBfdEthernetRead,
		DeleteContext: resourceIpv6RouteStaticBfdEthernetDelete,
		Schema: map[string]*schema.Schema{
			"eth_num": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Ethernet (not a member of vlan or trunk)",
				ValidateFunc: validation.IntAtLeast(1),
			},
			"nexthop_ipv6_ll": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Nexthop IPv6 address (Link_local)",
				ValidateFunc: validation.IsIPv6Address,
			},
			"template": {
				Type: schema.TypeString, Optional: true, Description: "Configure tracking template (bind tracking template name)",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
				ValidateFunc: validation.IntBetween(1, 255),
			},
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'down': BFD down;  (BFD state)",
				ValidateFunc: validation.StringInSlice([]string{"down"}, false),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceIpv6RouteStaticBfdEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdEthernetCreate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdEthernet(d)
		d.SetId(obj.GetId())
		err := PostEndpointIpv6RouteStaticBfdEthernet(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdEthernetRead()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		_, err := GetEndpointIpv6RouteStaticBfdEthernet(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteStaticBfdEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdEthernetUpdate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdEthernet(d)
		err := PutEndpointIpv6RouteStaticBfdEthernet(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdEthernetDelete()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		err := DeleteEndpointIpv6RouteStaticBfdEthernet(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteStaticBfdEthernet(d *schema.ResourceData) EndpointIpv6RouteStaticBfdEthernet {
	var ret EndpointIpv6RouteStaticBfdEthernet
	ret.Inst.EthNum = d.Get("eth_num").(int)
	ret.Inst.NexthopIpv6Ll = d.Get("nexthop_ipv6_ll").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
