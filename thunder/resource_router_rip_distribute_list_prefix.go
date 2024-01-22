package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterRipDistributeListPrefix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_rip_distribute_list_prefix`: Filter prefixes in routing updates\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterRipDistributeListPrefixCreate,
		UpdateContext: resourceRouterRipDistributeListPrefixUpdate,
		ReadContext:   resourceRouterRipDistributeListPrefixRead,
		DeleteContext: resourceRouterRipDistributeListPrefixDelete,

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
func resourceRouterRipDistributeListPrefixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListPrefixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeListPrefix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipDistributeListPrefixRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterRipDistributeListPrefixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListPrefixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeListPrefix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipDistributeListPrefixRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterRipDistributeListPrefixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListPrefixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeListPrefix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterRipDistributeListPrefixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListPrefixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeListPrefix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterRipDistributeListPrefixPrefixCfg(d []interface{}) []edpt.RouterRipDistributeListPrefixPrefixCfg {

	count1 := len(d)
	ret := make([]edpt.RouterRipDistributeListPrefixPrefixCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipDistributeListPrefixPrefixCfg
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

func dataToEndpointRouterRipDistributeListPrefix(d *schema.ResourceData) edpt.RouterRipDistributeListPrefix {
	var ret edpt.RouterRipDistributeListPrefix
	ret.Inst.PrefixCfg = getSliceRouterRipDistributeListPrefixPrefixCfg(d.Get("prefix_cfg").([]interface{}))
	//omit uuid
	return ret
}
