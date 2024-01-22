package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6RipDistributeListPrefix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_rip_distribute_list_prefix`: Filter prefixes in routing updates\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6RipDistributeListPrefixCreate,
		UpdateContext: resourceRouterIpv6RipDistributeListPrefixUpdate,
		ReadContext:   resourceRouterIpv6RipDistributeListPrefixRead,
		DeleteContext: resourceRouterIpv6RipDistributeListPrefixDelete,

		Schema: map[string]*schema.Schema{
			"prefix_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix_list": {
							Type: schema.TypeString, Optional: true, Description: "Filter prefixes in routing updates (Name of a prefix list)",
						},
						"prefix_list_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Filter incoming routing updates; 'out': Filter outgoing routing updates;",
						},
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"loopback": {
							Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRouterIpv6RipDistributeListPrefixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListPrefixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeListPrefix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipDistributeListPrefixRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6RipDistributeListPrefixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListPrefixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeListPrefix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipDistributeListPrefixRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6RipDistributeListPrefixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListPrefixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeListPrefix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6RipDistributeListPrefixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListPrefixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeListPrefix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6RipDistributeListPrefixPrefixCfg(d []interface{}) []edpt.RouterIpv6RipDistributeListPrefixPrefixCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipDistributeListPrefixPrefixCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipDistributeListPrefixPrefixCfg
		oi.PrefixList = in["prefix_list"].(string)
		oi.PrefixListDirection = in["prefix_list_direction"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIpv6RipDistributeListPrefix(d *schema.ResourceData) edpt.RouterIpv6RipDistributeListPrefix {
	var ret edpt.RouterIpv6RipDistributeListPrefix
	ret.Inst.PrefixCfg = getSliceRouterIpv6RipDistributeListPrefixPrefixCfg(d.Get("prefix_cfg").([]interface{}))
	//omit uuid
	return ret
}
