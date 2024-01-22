package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnelOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_tunnel_oper`: Operational Status for the object tunnel\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceTunnelOperRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Tunnel interface number",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"line_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"link_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mac": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"config_speed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 address",
						},
						"ipv4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 subnet mask",
						},
						"ipv6_link_local": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_link_local_prefix": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_link_local_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_link_local_scope": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4_addr_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"mask": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
									},
								},
							},
						},
						"ipv6_addr_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"prefix": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
									},
									"is_anycast": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceInterfaceTunnelOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceTunnelOperOper := setObjectInterfaceTunnelOperOper(res)
		d.Set("oper", InterfaceTunnelOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceTunnelOperOper(ret edpt.DataInterfaceTunnelOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                  ret.DtInterfaceTunnelOper.Oper.State,
			"line_protocol":          ret.DtInterfaceTunnelOper.Oper.LineProtocol,
			"link_type":              ret.DtInterfaceTunnelOper.Oper.LinkType,
			"mac":                    ret.DtInterfaceTunnelOper.Oper.Mac,
			"config_speed":           ret.DtInterfaceTunnelOper.Oper.ConfigSpeed,
			"ipv4_address":           ret.DtInterfaceTunnelOper.Oper.Ipv4Address,
			"ipv4_netmask":           ret.DtInterfaceTunnelOper.Oper.Ipv4Netmask,
			"ipv6_link_local":        ret.DtInterfaceTunnelOper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix": ret.DtInterfaceTunnelOper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_link_local_type":   ret.DtInterfaceTunnelOper.Oper.Ipv6LinkLocalType,
			"ipv6_link_local_scope":  ret.DtInterfaceTunnelOper.Oper.Ipv6LinkLocalScope,
			"ipv4_addr_count":        ret.DtInterfaceTunnelOper.Oper.Ipv4_addr_count,
			"ipv4_list":              setSliceInterfaceTunnelOperOperIpv4_list(ret.DtInterfaceTunnelOper.Oper.Ipv4_list),
			"ipv6_addr_count":        ret.DtInterfaceTunnelOper.Oper.Ipv6_addr_count,
			"ipv6_list":              setSliceInterfaceTunnelOperOperIpv6_list(ret.DtInterfaceTunnelOper.Oper.Ipv6_list),
		},
	}
}

func setSliceInterfaceTunnelOperOperIpv4_list(d []edpt.InterfaceTunnelOperOperIpv4_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr"] = item.Addr
		in["mask"] = item.Mask
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceTunnelOperOperIpv6_list(d []edpt.InterfaceTunnelOperOperIpv6_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr"] = item.Addr
		in["prefix"] = item.Prefix
		in["is_anycast"] = item.Is_anycast
		result = append(result, in)
	}
	return result
}

func getObjectInterfaceTunnelOperOper(d []interface{}) edpt.InterfaceTunnelOperOper {

	count1 := len(d)
	var ret edpt.InterfaceTunnelOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.LineProtocol = in["line_protocol"].(string)
		ret.LinkType = in["link_type"].(string)
		ret.Mac = in["mac"].(string)
		ret.ConfigSpeed = in["config_speed"].(int)
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.Ipv6LinkLocal = in["ipv6_link_local"].(string)
		ret.Ipv6LinkLocalPrefix = in["ipv6_link_local_prefix"].(string)
		ret.Ipv6LinkLocalType = in["ipv6_link_local_type"].(string)
		ret.Ipv6LinkLocalScope = in["ipv6_link_local_scope"].(string)
		ret.Ipv4_addr_count = in["ipv4_addr_count"].(int)
		ret.Ipv4_list = getSliceInterfaceTunnelOperOperIpv4_list(in["ipv4_list"].([]interface{}))
		ret.Ipv6_addr_count = in["ipv6_addr_count"].(int)
		ret.Ipv6_list = getSliceInterfaceTunnelOperOperIpv6_list(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTunnelOperOperIpv4_list(d []interface{}) []edpt.InterfaceTunnelOperOperIpv4_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelOperOperIpv4_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelOperOperIpv4_list
		oi.Addr = in["addr"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelOperOperIpv6_list(d []interface{}) []edpt.InterfaceTunnelOperOperIpv6_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelOperOperIpv6_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelOperOperIpv6_list
		oi.Addr = in["addr"].(string)
		oi.Prefix = in["prefix"].(string)
		oi.Is_anycast = in["is_anycast"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTunnelOper(d *schema.ResourceData) edpt.InterfaceTunnelOper {
	var ret edpt.InterfaceTunnelOper

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Oper = getObjectInterfaceTunnelOperOper(d.Get("oper").([]interface{}))
	return ret
}
