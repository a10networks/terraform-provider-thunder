package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireEthernetGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_virtual_wire_ethernet_group`: \n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVirtualWireEthernetGroupCreate,
		UpdateContext: resourceNetworkVirtualWireEthernetGroupUpdate,
		ReadContext:   resourceNetworkVirtualWireEthernetGroupRead,
		DeleteContext: resourceNetworkVirtualWireEthernetGroupDelete,

		Schema: map[string]*schema.Schema{
			"eth_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eth_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"eth_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"group_id": {
				Type: schema.TypeInt, Required: true, Description: "",
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
		},
	}
}
func resourceNetworkVirtualWireEthernetGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireEthernetGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireEthernetGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireEthernetGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVirtualWireEthernetGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireEthernetGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireEthernetGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireEthernetGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVirtualWireEthernetGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireEthernetGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireEthernetGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVirtualWireEthernetGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireEthernetGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireEthernetGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkVirtualWireEthernetGroupEthList(d []interface{}) []edpt.NetworkVirtualWireEthernetGroupEthList {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireEthernetGroupEthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireEthernetGroupEthList
		oi.EthStart = in["eth_start"].(int)
		oi.EthEnd = in["eth_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVirtualWireEthernetGroupTrunkList(d []interface{}) []edpt.NetworkVirtualWireEthernetGroupTrunkList {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireEthernetGroupTrunkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireEthernetGroupTrunkList
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireEthernetGroup(d *schema.ResourceData) edpt.NetworkVirtualWireEthernetGroup {
	var ret edpt.NetworkVirtualWireEthernetGroup
	ret.Inst.EthList = getSliceNetworkVirtualWireEthernetGroupEthList(d.Get("eth_list").([]interface{}))
	ret.Inst.GroupId = d.Get("group_id").(int)
	ret.Inst.TrunkList = getSliceNetworkVirtualWireEthernetGroupTrunkList(d.Get("trunk_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
