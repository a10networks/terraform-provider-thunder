package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterRipOffsetList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_rip_offset_list`: Modify RIP metric\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterRipOffsetListCreate,
		UpdateContext: resourceRouterRipOffsetListUpdate,
		ReadContext:   resourceRouterRipOffsetListRead,
		DeleteContext: resourceRouterRipOffsetListDelete,

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
func resourceRouterRipOffsetListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipOffsetListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipOffsetList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipOffsetListRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterRipOffsetListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipOffsetListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipOffsetList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterRipOffsetListRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterRipOffsetListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipOffsetListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipOffsetList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterRipOffsetListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterRipOffsetListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterRipOffsetList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterRipOffsetListAclCfg(d []interface{}) []edpt.RouterRipOffsetListAclCfg {

	count1 := len(d)
	ret := make([]edpt.RouterRipOffsetListAclCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterRipOffsetListAclCfg
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

func dataToEndpointRouterRipOffsetList(d *schema.ResourceData) edpt.RouterRipOffsetList {
	var ret edpt.RouterRipOffsetList
	ret.Inst.AclCfg = getSliceRouterRipOffsetListAclCfg(d.Get("acl_cfg").([]interface{}))
	//omit uuid
	return ret
}
