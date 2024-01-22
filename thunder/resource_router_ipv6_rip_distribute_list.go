package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6RipDistributeList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_rip_distribute_list`: Filter networks in routing updates\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6RipDistributeListCreate,
		UpdateContext: resourceRouterIpv6RipDistributeListUpdate,
		ReadContext:   resourceRouterIpv6RipDistributeListRead,
		DeleteContext: resourceRouterIpv6RipDistributeListDelete,

		Schema: map[string]*schema.Schema{
			"acl_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl": {
							Type: schema.TypeString, Optional: true, Description: "Access-list name",
						},
						"acl_direction": {
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
			"prefix": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRouterIpv6RipDistributeListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipDistributeListRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6RipDistributeListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipDistributeListRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6RipDistributeListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6RipDistributeListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipDistributeListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipDistributeList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6RipDistributeListAclCfg(d []interface{}) []edpt.RouterIpv6RipDistributeListAclCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipDistributeListAclCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipDistributeListAclCfg
		oi.Acl = in["acl"].(string)
		oi.AclDirection = in["acl_direction"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIpv6RipDistributeListPrefix1255(d []interface{}) edpt.RouterIpv6RipDistributeListPrefix1255 {

	count1 := len(d)
	var ret edpt.RouterIpv6RipDistributeListPrefix1255
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixCfg = getSliceRouterIpv6RipDistributeListPrefixPrefixCfg1256(in["prefix_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIpv6RipDistributeListPrefixPrefixCfg1256(d []interface{}) []edpt.RouterIpv6RipDistributeListPrefixPrefixCfg1256 {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipDistributeListPrefixPrefixCfg1256, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipDistributeListPrefixPrefixCfg1256
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

func dataToEndpointRouterIpv6RipDistributeList(d *schema.ResourceData) edpt.RouterIpv6RipDistributeList {
	var ret edpt.RouterIpv6RipDistributeList
	ret.Inst.AclCfg = getSliceRouterIpv6RipDistributeListAclCfg(d.Get("acl_cfg").([]interface{}))
	ret.Inst.Prefix = getObjectRouterIpv6RipDistributeListPrefix1255(d.Get("prefix").([]interface{}))
	//omit uuid
	return ret
}
