package thunder

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"util"
)

func resourceIpv6RouteStaticBfdBfdIpv6() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpv6RouteStaticBfdBfdIpv6Create,
		UpdateContext: resourceIpv6RouteStaticBfdBfdIpv6Update,
		ReadContext:   resourceIpv6RouteStaticBfdBfdIpv6Read,
		DeleteContext: resourceIpv6RouteStaticBfdBfdIpv6Delete,
		Schema: map[string]*schema.Schema{
			"local_ipv6": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Local IPv6 address",
				ValidateFunc: validation.IsIPv6Address,
			},
			"nexthop_ipv6": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Nexthop IPv6 address",
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

func resourceIpv6RouteStaticBfdBfdIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Create()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdBfdIpv6(d)
		d.SetId(obj.GetId())
		err := PostEndpointIpv6RouteStaticBfdBfdIpv6(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdBfdIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdBfdIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Read()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		_, err := GetEndpointIpv6RouteStaticBfdBfdIpv6(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteStaticBfdBfdIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Update()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteStaticBfdBfdIpv6(d)
		err := PutEndpointIpv6RouteStaticBfdBfdIpv6(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteStaticBfdBfdIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteStaticBfdBfdIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceIpv6RouteStaticBfdBfdIpv6Delete()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		err := DeleteEndpointIpv6RouteStaticBfdBfdIpv6(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteStaticBfdBfdIpv6(d *schema.ResourceData) EndpointIpv6RouteStaticBfdBfdIpv6 {
	var ret EndpointIpv6RouteStaticBfdBfdIpv6
	ret.Inst.LocalIpv6 = d.Get("local_ipv6").(string)
	ret.Inst.NexthopIpv6 = d.Get("nexthop_ipv6").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
