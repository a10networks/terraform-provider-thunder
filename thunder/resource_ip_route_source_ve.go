package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpRouteSourceVe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_route_source_ve`: Source ve\n\n__PLACEHOLDER__",
		CreateContext: resourceIpRouteSourceVeCreate,
		UpdateContext: resourceIpRouteSourceVeUpdate,
		ReadContext:   resourceIpRouteSourceVeRead,
		DeleteContext: resourceIpRouteSourceVeDelete,

		Schema: map[string]*schema.Schema{
			"nexthop_ip": {
				Type: schema.TypeString, Required: true, Description: "Nexthop IP address",
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
func resourceIpRouteSourceVeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceVeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceVe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteSourceVeRead(ctx, d, meta)
	}
	return diags
}

func resourceIpRouteSourceVeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceVeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceVe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteSourceVeRead(ctx, d, meta)
	}
	return diags
}
func resourceIpRouteSourceVeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceVeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceVe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpRouteSourceVeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceVeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceVe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpRouteSourceVe(d *schema.ResourceData) edpt.IpRouteSourceVe {
	var ret edpt.IpRouteSourceVe
	ret.Inst.NexthopIp = d.Get("nexthop_ip").(string)
	//omit uuid
	ret.Inst.VeNum = d.Get("ve_num").(int)
	return ret
}
