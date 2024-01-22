package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAInterfaceTrunk() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_interface_trunk`: VRRP-A interface trunk\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAInterfaceTrunkCreate,
		UpdateContext: resourceVrrpAInterfaceTrunkUpdate,
		ReadContext:   resourceVrrpAInterfaceTrunkRead,
		DeleteContext: resourceVrrpAInterfaceTrunkDelete,

		Schema: map[string]*schema.Schema{
			"both": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "both a router and server interface",
			},
			"no_heartbeat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "do not send out heartbeat packet from this interface",
			},
			"router_interface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "interface to upstream router",
			},
			"server_interface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "interface to real server",
			},
			"trunk_val": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet Interface",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
						},
					},
				},
			},
		},
	}
}
func resourceVrrpAInterfaceTrunkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceTrunkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceTrunk(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAInterfaceTrunkRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAInterfaceTrunkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceTrunkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceTrunk(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAInterfaceTrunkRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAInterfaceTrunkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceTrunkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceTrunk(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAInterfaceTrunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAInterfaceTrunkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAInterfaceTrunk(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVrrpAInterfaceTrunkVlanCfg(d []interface{}) []edpt.VrrpAInterfaceTrunkVlanCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAInterfaceTrunkVlanCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAInterfaceTrunkVlanCfg
		oi.Vlan = in["vlan"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAInterfaceTrunk(d *schema.ResourceData) edpt.VrrpAInterfaceTrunk {
	var ret edpt.VrrpAInterfaceTrunk
	ret.Inst.Both = d.Get("both").(int)
	ret.Inst.NoHeartbeat = d.Get("no_heartbeat").(int)
	ret.Inst.RouterInterface = d.Get("router_interface").(int)
	ret.Inst.ServerInterface = d.Get("server_interface").(int)
	ret.Inst.TrunkVal = d.Get("trunk_val").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VlanCfg = getSliceVrrpAInterfaceTrunkVlanCfg(d.Get("vlan_cfg").([]interface{}))
	return ret
}
