package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_lif_oper`: Operational Status for the object lif\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceLifOperRead,

		Schema: map[string]*schema.Schema{
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Lif interface name",
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
						"encapsulation_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"member_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"keep_alive": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mac": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"igmp_query_sent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"icmp_rate_limit_current": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"icmp_rate_over_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"icmp6_rate_limit_current": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"icmp6_rate_over_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "",
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
						"ip_unnumbered_enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mtu": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceLifOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceLifOperOper := setObjectInterfaceLifOperOper(res)
		d.Set("oper", InterfaceLifOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceLifOperOper(ret edpt.DataInterfaceLifOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                      ret.DtInterfaceLifOper.Oper.State,
			"line_protocol":              ret.DtInterfaceLifOper.Oper.Line_protocol,
			"link_type":                  ret.DtInterfaceLifOper.Oper.Link_type,
			"encapsulation_type":         ret.DtInterfaceLifOper.Oper.Encapsulation_type,
			"member_id":                  ret.DtInterfaceLifOper.Oper.Member_id,
			"keep_alive":                 ret.DtInterfaceLifOper.Oper.Keep_alive,
			"mac":                        ret.DtInterfaceLifOper.Oper.Mac,
			"igmp_query_sent":            ret.DtInterfaceLifOper.Oper.IgmpQuerySent,
			"icmp_rate_limit_current":    ret.DtInterfaceLifOper.Oper.IcmpRateLimitCurrent,
			"icmp_rate_over_limit_drop":  ret.DtInterfaceLifOper.Oper.IcmpRateOverLimitDrop,
			"icmp6_rate_limit_current":   ret.DtInterfaceLifOper.Oper.Icmp6RateLimitCurrent,
			"icmp6_rate_over_limit_drop": ret.DtInterfaceLifOper.Oper.Icmp6RateOverLimitDrop,
			"ipv4_addr_count":            ret.DtInterfaceLifOper.Oper.Ipv4_addr_count,
			"ipv4_list":                  setSliceInterfaceLifOperOperIpv4_list(ret.DtInterfaceLifOper.Oper.Ipv4_list),
			"ipv6_addr_count":            ret.DtInterfaceLifOper.Oper.Ipv6_addr_count,
			"ipv6_list":                  setSliceInterfaceLifOperOperIpv6_list(ret.DtInterfaceLifOper.Oper.Ipv6_list),
			"ipv6_link_local":            ret.DtInterfaceLifOper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix":     ret.DtInterfaceLifOper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_link_local_type":       ret.DtInterfaceLifOper.Oper.Ipv6LinkLocalType,
			"ipv6_link_local_scope":      ret.DtInterfaceLifOper.Oper.Ipv6LinkLocalScope,
			"ip_unnumbered_enabled":      ret.DtInterfaceLifOper.Oper.Ip_unnumbered_enabled,
			"mtu":                        ret.DtInterfaceLifOper.Oper.Mtu,
		},
	}
}

func setSliceInterfaceLifOperOperIpv4_list(d []edpt.InterfaceLifOperOperIpv4_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr"] = item.Addr
		in["mask"] = item.Mask
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceLifOperOperIpv6_list(d []edpt.InterfaceLifOperOperIpv6_list) []map[string]interface{} {
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

func getObjectInterfaceLifOperOper(d []interface{}) edpt.InterfaceLifOperOper {

	count1 := len(d)
	var ret edpt.InterfaceLifOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Line_protocol = in["line_protocol"].(string)
		ret.Link_type = in["link_type"].(string)
		ret.Encapsulation_type = in["encapsulation_type"].(string)
		ret.Member_id = in["member_id"].(string)
		ret.Keep_alive = in["keep_alive"].(string)
		ret.Mac = in["mac"].(string)
		ret.IgmpQuerySent = in["igmp_query_sent"].(int)
		ret.IcmpRateLimitCurrent = in["icmp_rate_limit_current"].(int)
		ret.IcmpRateOverLimitDrop = in["icmp_rate_over_limit_drop"].(int)
		ret.Icmp6RateLimitCurrent = in["icmp6_rate_limit_current"].(int)
		ret.Icmp6RateOverLimitDrop = in["icmp6_rate_over_limit_drop"].(int)
		ret.Ipv4_addr_count = in["ipv4_addr_count"].(int)
		ret.Ipv4_list = getSliceInterfaceLifOperOperIpv4_list(in["ipv4_list"].([]interface{}))
		ret.Ipv6_addr_count = in["ipv6_addr_count"].(int)
		ret.Ipv6_list = getSliceInterfaceLifOperOperIpv6_list(in["ipv6_list"].([]interface{}))
		ret.Ipv6LinkLocal = in["ipv6_link_local"].(string)
		ret.Ipv6LinkLocalPrefix = in["ipv6_link_local_prefix"].(string)
		ret.Ipv6LinkLocalType = in["ipv6_link_local_type"].(string)
		ret.Ipv6LinkLocalScope = in["ipv6_link_local_scope"].(string)
		ret.Ip_unnumbered_enabled = in["ip_unnumbered_enabled"].(int)
		ret.Mtu = in["mtu"].(string)
	}
	return ret
}

func getSliceInterfaceLifOperOperIpv4_list(d []interface{}) []edpt.InterfaceLifOperOperIpv4_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifOperOperIpv4_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifOperOperIpv4_list
		oi.Addr = in["addr"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifOperOperIpv6_list(d []interface{}) []edpt.InterfaceLifOperOperIpv6_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifOperOperIpv6_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifOperOperIpv6_list
		oi.Addr = in["addr"].(string)
		oi.Prefix = in["prefix"].(string)
		oi.Is_anycast = in["is_anycast"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLifOper(d *schema.ResourceData) edpt.InterfaceLifOper {
	var ret edpt.InterfaceLifOper

	ret.Ifname = d.Get("ifname").(string)

	ret.Oper = getObjectInterfaceLifOperOper(d.Get("oper").([]interface{}))
	return ret
}
