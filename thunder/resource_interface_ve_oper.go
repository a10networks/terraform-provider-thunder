package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_ve_oper`: Operational Status for the object ve\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceVeOperRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Virtual ethernet interface number",
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
						"user_trunk_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip_unnumbered_oper": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip_unnumbered_enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip_unnumbered_mac_learned": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ip_unnumbered_peer_lla": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"mtu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceVeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceVeOperOper := setObjectInterfaceVeOperOper(res)
		d.Set("oper", InterfaceVeOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceVeOperOper(ret edpt.DataInterfaceVeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                      ret.DtInterfaceVeOper.Oper.State,
			"line_protocol":              ret.DtInterfaceVeOper.Oper.Line_protocol,
			"link_type":                  ret.DtInterfaceVeOper.Oper.Link_type,
			"mac":                        ret.DtInterfaceVeOper.Oper.Mac,
			"ipv4_address":               ret.DtInterfaceVeOper.Oper.Ipv4Address,
			"ipv4_netmask":               ret.DtInterfaceVeOper.Oper.Ipv4Netmask,
			"ipv6_link_local":            ret.DtInterfaceVeOper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix":     ret.DtInterfaceVeOper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_link_local_type":       ret.DtInterfaceVeOper.Oper.Ipv6LinkLocalType,
			"ipv6_link_local_scope":      ret.DtInterfaceVeOper.Oper.Ipv6LinkLocalScope,
			"ipv4_addr_count":            ret.DtInterfaceVeOper.Oper.Ipv4_addr_count,
			"ipv4_list":                  setSliceInterfaceVeOperOperIpv4_list(ret.DtInterfaceVeOper.Oper.Ipv4_list),
			"ipv6_addr_count":            ret.DtInterfaceVeOper.Oper.Ipv6_addr_count,
			"ipv6_list":                  setSliceInterfaceVeOperOperIpv6_list(ret.DtInterfaceVeOper.Oper.Ipv6_list),
			"igmp_query_sent":            ret.DtInterfaceVeOper.Oper.IgmpQuerySent,
			"icmp_rate_limit_current":    ret.DtInterfaceVeOper.Oper.IcmpRateLimitCurrent,
			"icmp_rate_over_limit_drop":  ret.DtInterfaceVeOper.Oper.IcmpRateOverLimitDrop,
			"icmp6_rate_limit_current":   ret.DtInterfaceVeOper.Oper.Icmp6RateLimitCurrent,
			"icmp6_rate_over_limit_drop": ret.DtInterfaceVeOper.Oper.Icmp6RateOverLimitDrop,
			"user_trunk_id":              ret.DtInterfaceVeOper.Oper.UserTrunkId,
			"ip_unnumbered_oper":         ret.DtInterfaceVeOper.Oper.Ip_unnumbered_oper,
			"ip_unnumbered_enabled":      ret.DtInterfaceVeOper.Oper.Ip_unnumbered_enabled,
			"ip_unnumbered_mac_learned":  ret.DtInterfaceVeOper.Oper.Ip_unnumbered_mac_learned,
			"ip_unnumbered_peer_lla":     ret.DtInterfaceVeOper.Oper.Ip_unnumbered_peer_lla,
			"mtu":                        ret.DtInterfaceVeOper.Oper.Mtu,
		},
	}
}

func setSliceInterfaceVeOperOperIpv4_list(d []edpt.InterfaceVeOperOperIpv4_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr"] = item.Addr
		in["mask"] = item.Mask
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceVeOperOperIpv6_list(d []edpt.InterfaceVeOperOperIpv6_list) []map[string]interface{} {
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

func getObjectInterfaceVeOperOper(d []interface{}) edpt.InterfaceVeOperOper {

	count1 := len(d)
	var ret edpt.InterfaceVeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Line_protocol = in["line_protocol"].(string)
		ret.Link_type = in["link_type"].(string)
		ret.Mac = in["mac"].(string)
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.Ipv6LinkLocal = in["ipv6_link_local"].(string)
		ret.Ipv6LinkLocalPrefix = in["ipv6_link_local_prefix"].(string)
		ret.Ipv6LinkLocalType = in["ipv6_link_local_type"].(string)
		ret.Ipv6LinkLocalScope = in["ipv6_link_local_scope"].(string)
		ret.Ipv4_addr_count = in["ipv4_addr_count"].(int)
		ret.Ipv4_list = getSliceInterfaceVeOperOperIpv4_list(in["ipv4_list"].([]interface{}))
		ret.Ipv6_addr_count = in["ipv6_addr_count"].(int)
		ret.Ipv6_list = getSliceInterfaceVeOperOperIpv6_list(in["ipv6_list"].([]interface{}))
		ret.IgmpQuerySent = in["igmp_query_sent"].(int)
		ret.IcmpRateLimitCurrent = in["icmp_rate_limit_current"].(int)
		ret.IcmpRateOverLimitDrop = in["icmp_rate_over_limit_drop"].(int)
		ret.Icmp6RateLimitCurrent = in["icmp6_rate_limit_current"].(int)
		ret.Icmp6RateOverLimitDrop = in["icmp6_rate_over_limit_drop"].(int)
		ret.UserTrunkId = in["user_trunk_id"].(int)
		ret.Ip_unnumbered_oper = in["ip_unnumbered_oper"].(int)
		ret.Ip_unnumbered_enabled = in["ip_unnumbered_enabled"].(int)
		ret.Ip_unnumbered_mac_learned = in["ip_unnumbered_mac_learned"].(int)
		ret.Ip_unnumbered_peer_lla = in["ip_unnumbered_peer_lla"].(string)
		ret.Mtu = in["mtu"].(int)
	}
	return ret
}

func getSliceInterfaceVeOperOperIpv4_list(d []interface{}) []edpt.InterfaceVeOperOperIpv4_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeOperOperIpv4_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeOperOperIpv4_list
		oi.Addr = in["addr"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeOperOperIpv6_list(d []interface{}) []edpt.InterfaceVeOperOperIpv6_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeOperOperIpv6_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeOperOperIpv6_list
		oi.Addr = in["addr"].(string)
		oi.Prefix = in["prefix"].(string)
		oi.Is_anycast = in["is_anycast"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceVeOper(d *schema.ResourceData) edpt.InterfaceVeOper {
	var ret edpt.InterfaceVeOper

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Oper = getObjectInterfaceVeOperOper(d.Get("oper").([]interface{}))
	return ret
}
