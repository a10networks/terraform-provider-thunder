package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZoneInterface() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_zone_interface`: Interface\n\n__PLACEHOLDER__",
		CreateContext: resourceZoneInterfaceCreate,
		UpdateContext: resourceZoneInterfaceUpdate,
		ReadContext:   resourceZoneInterfaceRead,
		DeleteContext: resourceZoneInterfaceDelete,

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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceZoneInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneInterfaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneInterface(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneInterfaceRead(ctx, d, meta)
	}
	return diags
}

func resourceZoneInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneInterfaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneInterface(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneInterfaceRead(ctx, d, meta)
	}
	return diags
}
func resourceZoneInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneInterfaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneInterface(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceZoneInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneInterfaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneInterface(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceZoneInterfaceEthernetList(d []interface{}) []edpt.ZoneInterfaceEthernetList {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceEthernetList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceEthernetList
		oi.InterfaceEthernetStart = in["interface_ethernet_start"].(int)
		oi.InterfaceEthernetEnd = in["interface_ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceLifList(d []interface{}) []edpt.ZoneInterfaceLifList {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceLifList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceLifList
		oi.InterfaceLifStart = in["interface_lif_start"].(int)
		oi.InterfaceLifEnd = in["interface_lif_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceTrunkList(d []interface{}) []edpt.ZoneInterfaceTrunkList {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceTrunkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceTrunkList
		oi.InterfaceTrunkStart = in["interface_trunk_start"].(int)
		oi.InterfaceTrunkEnd = in["interface_trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceTunnelList(d []interface{}) []edpt.ZoneInterfaceTunnelList {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceTunnelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceTunnelList
		oi.InterfaceTunnelStart = in["interface_tunnel_start"].(int)
		oi.InterfaceTunnelEnd = in["interface_tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceZoneInterfaceVeList(d []interface{}) []edpt.ZoneInterfaceVeList {

	count1 := len(d)
	ret := make([]edpt.ZoneInterfaceVeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ZoneInterfaceVeList
		oi.InterfaceVeStart = in["interface_ve_start"].(int)
		oi.InterfaceVeEnd = in["interface_ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointZoneInterface(d *schema.ResourceData) edpt.ZoneInterface {
	var ret edpt.ZoneInterface
	ret.Inst.EthernetList = getSliceZoneInterfaceEthernetList(d.Get("ethernet_list").([]interface{}))
	ret.Inst.LifList = getSliceZoneInterfaceLifList(d.Get("lif_list").([]interface{}))
	ret.Inst.TrunkList = getSliceZoneInterfaceTrunkList(d.Get("trunk_list").([]interface{}))
	ret.Inst.TunnelList = getSliceZoneInterfaceTunnelList(d.Get("tunnel_list").([]interface{}))
	//omit uuid
	ret.Inst.VeList = getSliceZoneInterfaceVeList(d.Get("ve_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
