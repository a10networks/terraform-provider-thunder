package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterRipDistributeList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_rip_distribute_list`: Filter networks in routing updates\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterRipDistributeListCreate,
		UpdateContext: resourceRouterRipDistributeListUpdate,
		ReadContext:   resourceRouterRipDistributeListRead,
		DeleteContext: resourceRouterRipDistributeListDelete,

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
func resourceRouterRipDistributeListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipDistributeListRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterRipDistributeListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipDistributeListRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterRipDistributeListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterRipDistributeListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipDistributeListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipDistributeList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterRipDistributeListAclCfg(d []interface{}) []edpt.RouterRipDistributeListAclCfg {

	count1 := len(d)
	ret := make([]edpt.RouterRipDistributeListAclCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipDistributeListAclCfg
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

func getObjectRouterRipDistributeListPrefix1304(d []interface{}) edpt.RouterRipDistributeListPrefix1304 {

	count1 := len(d)
	var ret edpt.RouterRipDistributeListPrefix1304
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixCfg = getSliceRouterRipDistributeListPrefixPrefixCfg1305(in["prefix_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterRipDistributeListPrefixPrefixCfg1305(d []interface{}) []edpt.RouterRipDistributeListPrefixPrefixCfg1305 {

	count1 := len(d)
	ret := make([]edpt.RouterRipDistributeListPrefixPrefixCfg1305, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipDistributeListPrefixPrefixCfg1305
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

func dataToEndpointRouterRipDistributeList(d *schema.ResourceData) edpt.RouterRipDistributeList {
	var ret edpt.RouterRipDistributeList
	ret.Inst.AclCfg = getSliceRouterRipDistributeListAclCfg(d.Get("acl_cfg").([]interface{}))
	ret.Inst.Prefix = getObjectRouterRipDistributeListPrefix1304(d.Get("prefix").([]interface{}))
	//omit uuid
	return ret
}
