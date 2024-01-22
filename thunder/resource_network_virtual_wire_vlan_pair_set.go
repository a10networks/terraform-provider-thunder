package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireVlanPairSet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_virtual_wire_vlan_pair_set`: Virtual Wire VLAN Pair Set Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVirtualWireVlanPairSetCreate,
		UpdateContext: resourceNetworkVirtualWireVlanPairSetUpdate,
		ReadContext:   resourceNetworkVirtualWireVlanPairSetRead,
		DeleteContext: resourceNetworkVirtualWireVlanPairSetDelete,

		Schema: map[string]*schema.Schema{
			"set_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_wire_vlan_pair": {
							Type: schema.TypeInt, Optional: true, Description: "virtual wire vlan pair id",
						},
					},
				},
			},
			"set_id": {
				Type: schema.TypeInt, Required: true, Description: "virtual wire vlan set id",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkVirtualWireVlanPairSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairSetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPairSet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireVlanPairSetRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVirtualWireVlanPairSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairSetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPairSet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireVlanPairSetRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVirtualWireVlanPairSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairSetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPairSet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVirtualWireVlanPairSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairSetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPairSet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkVirtualWireVlanPairSetSetCfg(d []interface{}) []edpt.NetworkVirtualWireVlanPairSetSetCfg {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireVlanPairSetSetCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireVlanPairSetSetCfg
		oi.VirtualWireVlanPair = in["virtual_wire_vlan_pair"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireVlanPairSet(d *schema.ResourceData) edpt.NetworkVirtualWireVlanPairSet {
	var ret edpt.NetworkVirtualWireVlanPairSet
	ret.Inst.SetCfg = getSliceNetworkVirtualWireVlanPairSetSetCfg(d.Get("set_cfg").([]interface{}))
	ret.Inst.SetId = d.Get("set_id").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
