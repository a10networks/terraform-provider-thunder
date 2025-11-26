package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZone() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_zone`: Security zone\n\n__PLACEHOLDER__",
		CreateContext: resourceZoneCreate,
		UpdateContext: resourceZoneUpdate,
		ReadContext:   resourceZoneRead,
		DeleteContext: resourceZoneDelete,

		Schema: map[string]*schema.Schema{
			"interface": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_ethernet_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interface_ethernet_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"trunk_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_trunk_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interface_trunk_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"ve_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_ve_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interface_ve_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"lif_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_lif_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interface_lif_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"tunnel_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_tunnel_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"interface_tunnel_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
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
			"local_zone_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set to local zone",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "name of zone object",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_start": {
										Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
									},
									"vlan_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
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
		},
	}
}
func resourceZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZone(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneRead(ctx, d, meta)
	}
	return diags
}

func resourceZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZone(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneRead(ctx, d, meta)
	}
	return diags
}
func resourceZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZone(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZone(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectZoneInterface3780(d []interface{}) edpt.ZoneInterface3780 {

	count1 := len(d)
	var ret edpt.ZoneInterface3780
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthernetList = getSliceZoneInterfaceEthernetList3781(in["ethernet_list"].([]interface{}))
		ret.TrunkList = getSliceZoneInterfaceTrunkList3782(in["trunk_list"].([]interface{}))
		ret.VeList = getSliceZoneInterfaceVeList3783(in["ve_list"].([]interface{}))
		ret.LifList = getSliceZoneInterfaceLifList3784(in["lif_list"].([]interface{}))
		ret.TunnelList = getSliceZoneInterfaceTunnelList3785(in["tunnel_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceZoneInterfaceEthernetList3781(d []interface{}) []edpt.ZoneInterfaceEthernetList3781 {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceEthernetList3781, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceEthernetList3781
		oi.InterfaceEthernetStart = in["interface_ethernet_start"].(int)
		oi.InterfaceEthernetEnd = in["interface_ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceTrunkList3782(d []interface{}) []edpt.ZoneInterfaceTrunkList3782 {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceTrunkList3782, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceTrunkList3782
		oi.InterfaceTrunkStart = in["interface_trunk_start"].(int)
		oi.InterfaceTrunkEnd = in["interface_trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceVeList3783(d []interface{}) []edpt.ZoneInterfaceVeList3783 {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceVeList3783, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceVeList3783
		oi.InterfaceVeStart = in["interface_ve_start"].(int)
		oi.InterfaceVeEnd = in["interface_ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceLifList3784(d []interface{}) []edpt.ZoneInterfaceLifList3784 {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceLifList3784, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceLifList3784
		oi.InterfaceLifStart = in["interface_lif_start"].(int)
		oi.InterfaceLifEnd = in["interface_lif_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceTunnelList3785(d []interface{}) []edpt.ZoneInterfaceTunnelList3785 {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceTunnelList3785, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceTunnelList3785
		oi.InterfaceTunnelStart = in["interface_tunnel_start"].(int)
		oi.InterfaceTunnelEnd = in["interface_tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectZoneLocalZoneCfg3786(d []interface{}) edpt.ZoneLocalZoneCfg3786 {

	count1 := len(d)
	var ret edpt.ZoneLocalZoneCfg3786
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalType = in["local_type"].(int)
		//omit uuid
	}
	return ret
}

func getObjectZoneVlan3787(d []interface{}) edpt.ZoneVlan3787 {

	count1 := len(d)
	var ret edpt.ZoneVlan3787
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VlanList = getSliceZoneVlanVlanList3788(in["vlan_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceZoneVlanVlanList3788(d []interface{}) []edpt.ZoneVlanVlanList3788 {

	count1 := len(d)
	ret := make([]edpt.ZoneVlanVlanList3788, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneVlanVlanList3788
		oi.VlanStart = in["vlan_start"].(int)
		oi.VlanEnd = in["vlan_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointZone(d *schema.ResourceData) edpt.Zone {
	var ret edpt.Zone
	ret.Inst.Interface = getObjectZoneInterface3780(d.Get("interface").([]interface{}))
	ret.Inst.LocalZoneCfg = getObjectZoneLocalZoneCfg3786(d.Get("local_zone_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vlan = getObjectZoneVlan3787(d.Get("vlan").([]interface{}))
	return ret
}
