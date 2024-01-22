package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkBridgeVlanGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_bridge_vlan_group`: Bridge VLAN Group Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkBridgeVlanGroupCreate,
		UpdateContext: resourceNetworkBridgeVlanGroupUpdate,
		ReadContext:   resourceNetworkBridgeVlanGroupRead,
		DeleteContext: resourceNetworkBridgeVlanGroupDelete,

		Schema: map[string]*schema.Schema{
			"bridge_vlan_group_number": {
				Type: schema.TypeInt, Required: true, Description: "Bridge VLAN Group Number",
			},
			"forward_traffic": {
				Type: schema.TypeString, Optional: true, Default: "forward-ip-traffic", Description: "'forward-all-traffic': Forward all traffic between bridge members; 'forward-ip-traffic': Forward only IP traffic between bridge members (default);",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Bridge Group Name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve": {
				Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet Port (Virtual Ethernet Port number)",
			},
			"vlan_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_start": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN id",
						},
						"vlan_end": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN id",
						},
					},
				},
			},
		},
	}
}
func resourceNetworkBridgeVlanGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBridgeVlanGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBridgeVlanGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkBridgeVlanGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkBridgeVlanGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBridgeVlanGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBridgeVlanGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkBridgeVlanGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkBridgeVlanGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBridgeVlanGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBridgeVlanGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkBridgeVlanGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBridgeVlanGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBridgeVlanGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkBridgeVlanGroupVlanList(d []interface{}) []edpt.NetworkBridgeVlanGroupVlanList {

	count1 := len(d)
	ret := make([]edpt.NetworkBridgeVlanGroupVlanList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkBridgeVlanGroupVlanList
		oi.VlanStart = in["vlan_start"].(int)
		oi.VlanEnd = in["vlan_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkBridgeVlanGroup(d *schema.ResourceData) edpt.NetworkBridgeVlanGroup {
	var ret edpt.NetworkBridgeVlanGroup
	ret.Inst.BridgeVlanGroupNumber = d.Get("bridge_vlan_group_number").(int)
	ret.Inst.ForwardTraffic = d.Get("forward_traffic").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.VlanList = getSliceNetworkBridgeVlanGroupVlanList(d.Get("vlan_list").([]interface{}))
	return ret
}
