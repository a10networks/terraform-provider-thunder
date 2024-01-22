package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireVlanPair() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_virtual_wire_vlan_pair`: Virtual Wire VLAN Pair Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVirtualWireVlanPairCreate,
		UpdateContext: resourceNetworkVirtualWireVlanPairUpdate,
		ReadContext:   resourceNetworkVirtualWireVlanPairRead,
		DeleteContext: resourceNetworkVirtualWireVlanPairDelete,

		Schema: map[string]*schema.Schema{
			"pair_id": {
				Type: schema.TypeInt, Required: true, Description: "virtual wire vlan pair id",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan1": {
				Type: schema.TypeInt, Optional: true, Description: "first vlan in the pair",
			},
			"vlan2": {
				Type: schema.TypeInt, Optional: true, Description: "second vlan in the pair",
			},
		},
	}
}
func resourceNetworkVirtualWireVlanPairCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPair(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireVlanPairRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVirtualWireVlanPairUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPair(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireVlanPairRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVirtualWireVlanPairDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPair(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVirtualWireVlanPairRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireVlanPairRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireVlanPair(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkVirtualWireVlanPair(d *schema.ResourceData) edpt.NetworkVirtualWireVlanPair {
	var ret edpt.NetworkVirtualWireVlanPair
	ret.Inst.PairId = d.Get("pair_id").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vlan1 = d.Get("vlan1").(int)
	ret.Inst.Vlan2 = d.Get("vlan2").(int)
	return ret
}
