package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceBriefOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_brief_oper`: Operational Status for the object brief\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceBriefOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interfaces": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_num": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"speed": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"duplexity": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"trunk_group": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vlan_info": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"encap_info": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"mac": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv4_addr_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"ipv4_mask": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
									},
									"ipv6_addr_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"ipv6_prefix": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
									},
									"intf_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"unnumbered_oper": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceBriefOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceBriefOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceBriefOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceBriefOperOper := setObjectInterfaceBriefOperOper(res)
		d.Set("oper", InterfaceBriefOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceBriefOperOper(ret edpt.DataInterfaceBriefOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"interfaces": setSliceInterfaceBriefOperOperInterfaces(ret.DtInterfaceBriefOper.Oper.Interfaces),
		},
	}
}

func setSliceInterfaceBriefOperOperInterfaces(d []edpt.InterfaceBriefOperOperInterfaces) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port_num"] = item.Port_num
		in["speed"] = item.Speed
		in["duplexity"] = item.Duplexity
		in["trunk_group"] = item.Trunk_group
		in["vlan_info"] = item.Vlan_info
		in["encap_info"] = item.Encap_info
		in["state"] = item.State
		in["mac"] = item.Mac
		in["ipv4_addr_count"] = item.Ipv4_addr_count
		in["ipv4_addr"] = item.Ipv4_addr
		in["ipv4_mask"] = item.Ipv4_mask
		in["ipv6_addr_count"] = item.Ipv6_addr_count
		in["ipv6_addr"] = item.Ipv6_addr
		in["ipv6_prefix"] = item.Ipv6_prefix
		in["intf_name"] = item.Intf_name
		in["unnumbered_oper"] = item.Unnumbered_oper
		result = append(result, in)
	}
	return result
}

func getObjectInterfaceBriefOperOper(d []interface{}) edpt.InterfaceBriefOperOper {

	count1 := len(d)
	var ret edpt.InterfaceBriefOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interfaces = getSliceInterfaceBriefOperOperInterfaces(in["interfaces"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceBriefOperOperInterfaces(d []interface{}) []edpt.InterfaceBriefOperOperInterfaces {

	count1 := len(d)
	ret := make([]edpt.InterfaceBriefOperOperInterfaces, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceBriefOperOperInterfaces
		oi.Port_num = in["port_num"].(string)
		oi.Speed = in["speed"].(string)
		oi.Duplexity = in["duplexity"].(string)
		oi.Trunk_group = in["trunk_group"].(string)
		oi.Vlan_info = in["vlan_info"].(string)
		oi.Encap_info = in["encap_info"].(string)
		oi.State = in["state"].(string)
		oi.Mac = in["mac"].(string)
		oi.Ipv4_addr_count = in["ipv4_addr_count"].(int)
		oi.Ipv4_addr = in["ipv4_addr"].(string)
		oi.Ipv4_mask = in["ipv4_mask"].(string)
		oi.Ipv6_addr_count = in["ipv6_addr_count"].(int)
		oi.Ipv6_addr = in["ipv6_addr"].(string)
		oi.Ipv6_prefix = in["ipv6_prefix"].(string)
		oi.Intf_name = in["intf_name"].(string)
		oi.Unnumbered_oper = in["unnumbered_oper"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceBriefOper(d *schema.ResourceData) edpt.InterfaceBriefOper {
	var ret edpt.InterfaceBriefOper

	ret.Oper = getObjectInterfaceBriefOperOper(d.Get("oper").([]interface{}))
	return ret
}
