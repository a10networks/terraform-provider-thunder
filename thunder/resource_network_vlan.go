package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVlan() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_vlan`: Configure VLAN\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVlanCreate,
		UpdateContext: resourceNetworkVlanUpdate,
		ReadContext:   resourceNetworkVlanRead,
		DeleteContext: resourceNetworkVlanDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "VLAN name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'broadcast_count': Broadcast counter; 'multicast_count': Multicast counter; 'ip_multicast_count': IP Multicast counter; 'unknown_unicast_count': Unknown Unicast counter; 'mac_movement_count': Mac Movement counter; 'shared_vlan_partition_switched_counter': SVLAN Partition switched counter;",
						},
					},
				},
			},
			"shared_vlan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure VLAN as a shared VLAN",
			},
			"tagged_eth_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tagged_ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
						},
						"tagged_ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
						},
					},
				},
			},
			"tagged_trunk_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tagged_trunk_start": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
						},
						"tagged_trunk_end": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
						},
					},
				},
			},
			"traffic_distribution_mode": {
				Type: schema.TypeString, Optional: true, Description: "'sip': sip; 'dip': dip; 'primary': primary; 'blade': blade; 'l4-src-port': l4-src-port; 'l4-dst-port': l4-dst-port;",
			},
			"untagged_eth_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"untagged_ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
						},
						"untagged_ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
						},
					},
				},
			},
			"untagged_lif": {
				Type: schema.TypeString, Optional: true, Description: "Logical tunnel interface (Logical tunnel interface name)",
			},
			"untagged_trunk_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"untagged_trunk_start": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
						},
						"untagged_trunk_end": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
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
			"ve": {
				Type: schema.TypeInt, Optional: true, Description: "ve number",
			},
			"vlan_num": {
				Type: schema.TypeInt, Required: true, Description: "VLAN number",
			},
		},
	}
}
func resourceNetworkVlanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlan(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVlanRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVlanUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlan(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVlanRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVlanDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlan(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVlanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlan(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkVlanSamplingEnable(d []interface{}) []edpt.NetworkVlanSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVlanTaggedEthList(d []interface{}) []edpt.NetworkVlanTaggedEthList {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanTaggedEthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanTaggedEthList
		oi.TaggedEthernetStart = in["tagged_ethernet_start"].(int)
		oi.TaggedEthernetEnd = in["tagged_ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVlanTaggedTrunkList(d []interface{}) []edpt.NetworkVlanTaggedTrunkList {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanTaggedTrunkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanTaggedTrunkList
		oi.TaggedTrunkStart = in["tagged_trunk_start"].(int)
		oi.TaggedTrunkEnd = in["tagged_trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVlanUntaggedEthList(d []interface{}) []edpt.NetworkVlanUntaggedEthList {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanUntaggedEthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanUntaggedEthList
		oi.UntaggedEthernetStart = in["untagged_ethernet_start"].(int)
		oi.UntaggedEthernetEnd = in["untagged_ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkVlanUntaggedTrunkList(d []interface{}) []edpt.NetworkVlanUntaggedTrunkList {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanUntaggedTrunkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanUntaggedTrunkList
		oi.UntaggedTrunkStart = in["untagged_trunk_start"].(int)
		oi.UntaggedTrunkEnd = in["untagged_trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVlan(d *schema.ResourceData) edpt.NetworkVlan {
	var ret edpt.NetworkVlan
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceNetworkVlanSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SharedVlan = d.Get("shared_vlan").(int)
	ret.Inst.TaggedEthList = getSliceNetworkVlanTaggedEthList(d.Get("tagged_eth_list").([]interface{}))
	ret.Inst.TaggedTrunkList = getSliceNetworkVlanTaggedTrunkList(d.Get("tagged_trunk_list").([]interface{}))
	ret.Inst.TrafficDistributionMode = d.Get("traffic_distribution_mode").(string)
	ret.Inst.UntaggedEthList = getSliceNetworkVlanUntaggedEthList(d.Get("untagged_eth_list").([]interface{}))
	ret.Inst.UntaggedLif = d.Get("untagged_lif").(string)
	ret.Inst.UntaggedTrunkList = getSliceNetworkVlanUntaggedTrunkList(d.Get("untagged_trunk_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	ret.Inst.VlanNum = d.Get("vlan_num").(int)
	return ret
}
