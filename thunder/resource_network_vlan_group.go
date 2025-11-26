package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVlanGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_vlan_group`: Configure group of VLANs\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVlanGroupCreate,
		UpdateContext: resourceNetworkVlanGroupUpdate,
		ReadContext:   resourceNetworkVlanGroupRead,
		DeleteContext: resourceNetworkVlanGroupDelete,

		Schema: map[string]*schema.Schema{
			"eth_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
						},
						"ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
						},
					},
				},
			},
			"trunk_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"trunk_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan_num_end": {
				Type: schema.TypeInt, Required: true, Description: "End of VLAN range",
			},
			"vlan_num_start": {
				Type: schema.TypeInt, Required: true, Description: "Start of VLAN range",
			},
		},
	}
}
func resourceNetworkVlanGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVlanGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVlanGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVlanGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVlanGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVlanGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkVlanGroupEthList(d []interface{}) []edpt.NetworkVlanGroupEthList {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanGroupEthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanGroupEthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVlanGroupTrunkList(d []interface{}) []edpt.NetworkVlanGroupTrunkList {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanGroupTrunkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanGroupTrunkList
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVlanGroup(d *schema.ResourceData) edpt.NetworkVlanGroup {
	var ret edpt.NetworkVlanGroup
	ret.Inst.EthList = getSliceNetworkVlanGroupEthList(d.Get("eth_list").([]interface{}))
	ret.Inst.TrunkList = getSliceNetworkVlanGroupTrunkList(d.Get("trunk_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VlanNumEnd = d.Get("vlan_num_end").(int)
	ret.Inst.VlanNumStart = d.Get("vlan_num_start").(int)
	return ret
}
