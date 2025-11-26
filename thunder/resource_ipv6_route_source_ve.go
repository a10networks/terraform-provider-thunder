package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RouteSourceVe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_route_source_ve`: Source ve\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RouteSourceVeCreate,
		UpdateContext: resourceIpv6RouteSourceVeUpdate,
		ReadContext:   resourceIpv6RouteSourceVeRead,
		DeleteContext: resourceIpv6RouteSourceVeDelete,

		Schema: map[string]*schema.Schema{
			"nexthop_ip": {
				Type: schema.TypeString, Required: true, Description: "Nexthop IPv6 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve_num": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
		},
	}
}
func resourceIpv6RouteSourceVeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceVeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceVe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteSourceVeRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteSourceVeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceVeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceVe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteSourceVeRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RouteSourceVeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceVeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceVe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteSourceVeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceVeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceVe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteSourceVe(d *schema.ResourceData) edpt.Ipv6RouteSourceVe {
	var ret edpt.Ipv6RouteSourceVe
	ret.Inst.NexthopIp = d.Get("nexthop_ip").(string)
	//omit uuid
	ret.Inst.VeNum = d.Get("ve_num").(int)
	return ret
}
