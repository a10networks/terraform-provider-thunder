package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpRouteStaticBfd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_route_static_bfd`: Bidirectional Forwarding Detection\n\n__PLACEHOLDER__",
		CreateContext: resourceIpRouteStaticBfdCreate,
		UpdateContext: resourceIpRouteStaticBfdUpdate,
		ReadContext:   resourceIpRouteStaticBfdRead,
		DeleteContext: resourceIpRouteStaticBfdDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'down': BFD down;  (BFD state)",
			},
			"local_ip": {
				Type: schema.TypeString, Required: true, Description: "Local IP address",
			},
			"nexthop_ip": {
				Type: schema.TypeString, Required: true, Description: "Nexthop IP address",
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
func resourceIpRouteStaticBfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteStaticBfdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteStaticBfd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteStaticBfdRead(ctx, d, meta)
	}
	return diags
}

func resourceIpRouteStaticBfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteStaticBfdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteStaticBfd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteStaticBfdRead(ctx, d, meta)
	}
	return diags
}
func resourceIpRouteStaticBfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteStaticBfdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteStaticBfd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpRouteStaticBfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteStaticBfdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteStaticBfd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpRouteStaticBfd(d *schema.ResourceData) edpt.IpRouteStaticBfd {
	var ret edpt.IpRouteStaticBfd
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.LocalIp = d.Get("local_ip").(string)
	ret.Inst.NexthopIp = d.Get("nexthop_ip").(string)
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	return ret
}
