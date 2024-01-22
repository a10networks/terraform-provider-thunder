package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6RipOffsetList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_rip_offset_list`: Modify RIP metric\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6RipOffsetListCreate,
		UpdateContext: resourceRouterIpv6RipOffsetListUpdate,
		ReadContext:   resourceRouterIpv6RipOffsetListRead,
		DeleteContext: resourceRouterIpv6RipOffsetListDelete,

		Schema: map[string]*schema.Schema{
			"acl_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl": {
							Type: schema.TypeString, Optional: true, Description: "Access-list name",
						},
						"offset_list_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Filter incoming updates; 'out': Filter outgoing updates;",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Description: "Metric value",
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
func resourceRouterIpv6RipOffsetListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipOffsetListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipOffsetList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipOffsetListRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6RipOffsetListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipOffsetListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipOffsetList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipOffsetListRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6RipOffsetListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipOffsetListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipOffsetList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6RipOffsetListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipOffsetListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipOffsetList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6RipOffsetListAclCfg(d []interface{}) []edpt.RouterIpv6RipOffsetListAclCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipOffsetListAclCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipOffsetListAclCfg
		oi.Acl = in["acl"].(string)
		oi.OffsetListDirection = in["offset_list_direction"].(string)
		oi.Metric = in["metric"].(int)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIpv6RipOffsetList(d *schema.ResourceData) edpt.RouterIpv6RipOffsetList {
	var ret edpt.RouterIpv6RipOffsetList
	ret.Inst.AclCfg = getSliceRouterIpv6RipOffsetListAclCfg(d.Get("acl_cfg").([]interface{}))
	//omit uuid
	return ret
}
