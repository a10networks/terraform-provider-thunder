package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_loopback_oper`: Operational Status for the object loopback\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceLoopbackOperRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Loopback interface number",
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
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
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

func resourceInterfaceLoopbackOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceLoopbackOperOper := setObjectInterfaceLoopbackOperOper(res)
		d.Set("oper", InterfaceLoopbackOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceLoopbackOperOper(ret edpt.DataInterfaceLoopbackOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                  ret.DtInterfaceLoopbackOper.Oper.State,
			"line_protocol":          ret.DtInterfaceLoopbackOper.Oper.Line_protocol,
			"ipv4_address":           ret.DtInterfaceLoopbackOper.Oper.Ipv4Address,
			"ipv4_netmask":           ret.DtInterfaceLoopbackOper.Oper.Ipv4Netmask,
			"ipv6_link_local":        ret.DtInterfaceLoopbackOper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix": ret.DtInterfaceLoopbackOper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_link_local_type":   ret.DtInterfaceLoopbackOper.Oper.Ipv6LinkLocalType,
			"ipv6_link_local_scope":  ret.DtInterfaceLoopbackOper.Oper.Ipv6LinkLocalScope,
			"ipv4_addr_count":        ret.DtInterfaceLoopbackOper.Oper.Ipv4_addr_count,
			"ipv4_list":              setSliceInterfaceLoopbackOperOperIpv4_list(ret.DtInterfaceLoopbackOper.Oper.Ipv4_list),
			"ipv6_addr_count":        ret.DtInterfaceLoopbackOper.Oper.Ipv6_addr_count,
			"ipv6_list":              setSliceInterfaceLoopbackOperOperIpv6_list(ret.DtInterfaceLoopbackOper.Oper.Ipv6_list),
		},
	}
}

func setSliceInterfaceLoopbackOperOperIpv4_list(d []edpt.InterfaceLoopbackOperOperIpv4_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr"] = item.Addr
		in["mask"] = item.Mask
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceLoopbackOperOperIpv6_list(d []edpt.InterfaceLoopbackOperOperIpv6_list) []map[string]interface{} {
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

func getObjectInterfaceLoopbackOperOper(d []interface{}) edpt.InterfaceLoopbackOperOper {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Line_protocol = in["line_protocol"].(string)
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.Ipv6LinkLocal = in["ipv6_link_local"].(string)
		ret.Ipv6LinkLocalPrefix = in["ipv6_link_local_prefix"].(string)
		ret.Ipv6LinkLocalType = in["ipv6_link_local_type"].(string)
		ret.Ipv6LinkLocalScope = in["ipv6_link_local_scope"].(string)
		ret.Ipv4_addr_count = in["ipv4_addr_count"].(int)
		ret.Ipv4_list = getSliceInterfaceLoopbackOperOperIpv4_list(in["ipv4_list"].([]interface{}))
		ret.Ipv6_addr_count = in["ipv6_addr_count"].(int)
		ret.Ipv6_list = getSliceInterfaceLoopbackOperOperIpv6_list(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackOperOperIpv4_list(d []interface{}) []edpt.InterfaceLoopbackOperOperIpv4_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackOperOperIpv4_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackOperOperIpv4_list
		oi.Addr = in["addr"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackOperOperIpv6_list(d []interface{}) []edpt.InterfaceLoopbackOperOperIpv6_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackOperOperIpv6_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackOperOperIpv6_list
		oi.Addr = in["addr"].(string)
		oi.Prefix = in["prefix"].(string)
		oi.Is_anycast = in["is_anycast"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLoopbackOper(d *schema.ResourceData) edpt.InterfaceLoopbackOper {
	var ret edpt.InterfaceLoopbackOper

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Oper = getObjectInterfaceLoopbackOperOper(d.Get("oper").([]interface{}))
	return ret
}
