package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_trunk_oper`: Operational Status for the object trunk\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceTrunkOperRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Trunk interface number",
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
						"trunk_member": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"members": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"hardware": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"address": {
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
						"vlan_id": {
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

func resourceInterfaceTrunkOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceTrunkOperOper := setObjectInterfaceTrunkOperOper(res)
		d.Set("oper", InterfaceTrunkOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceTrunkOperOper(ret edpt.DataInterfaceTrunkOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                      ret.DtInterfaceTrunkOper.Oper.State,
			"line_protocol":              ret.DtInterfaceTrunkOper.Oper.LineProtocol,
			"trunk_member":               setSliceInterfaceTrunkOperOperTrunkMember(ret.DtInterfaceTrunkOper.Oper.TrunkMember),
			"hardware":                   ret.DtInterfaceTrunkOper.Oper.Hardware,
			"address":                    ret.DtInterfaceTrunkOper.Oper.Address,
			"ipv4_address":               ret.DtInterfaceTrunkOper.Oper.Ipv4Address,
			"ipv4_netmask":               ret.DtInterfaceTrunkOper.Oper.Ipv4Netmask,
			"ipv6_link_local":            ret.DtInterfaceTrunkOper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix":     ret.DtInterfaceTrunkOper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_link_local_type":       ret.DtInterfaceTrunkOper.Oper.Ipv6LinkLocalType,
			"ipv6_link_local_scope":      ret.DtInterfaceTrunkOper.Oper.Ipv6LinkLocalScope,
			"ipv4_addr_count":            ret.DtInterfaceTrunkOper.Oper.Ipv4_addr_count,
			"ipv4_list":                  setSliceInterfaceTrunkOperOperIpv4_list(ret.DtInterfaceTrunkOper.Oper.Ipv4_list),
			"ipv6_addr_count":            ret.DtInterfaceTrunkOper.Oper.Ipv6_addr_count,
			"ipv6_list":                  setSliceInterfaceTrunkOperOperIpv6_list(ret.DtInterfaceTrunkOper.Oper.Ipv6_list),
			"igmp_query_sent":            ret.DtInterfaceTrunkOper.Oper.IgmpQuerySent,
			"icmp_rate_limit_current":    ret.DtInterfaceTrunkOper.Oper.IcmpRateLimitCurrent,
			"icmp_rate_over_limit_drop":  ret.DtInterfaceTrunkOper.Oper.IcmpRateOverLimitDrop,
			"icmp6_rate_limit_current":   ret.DtInterfaceTrunkOper.Oper.Icmp6RateLimitCurrent,
			"icmp6_rate_over_limit_drop": ret.DtInterfaceTrunkOper.Oper.Icmp6RateOverLimitDrop,
			"vlan_id":                    ret.DtInterfaceTrunkOper.Oper.VlanId,
			"ip_unnumbered_oper":         ret.DtInterfaceTrunkOper.Oper.Ip_unnumbered_oper,
			"ip_unnumbered_enabled":      ret.DtInterfaceTrunkOper.Oper.Ip_unnumbered_enabled,
			"ip_unnumbered_mac_learned":  ret.DtInterfaceTrunkOper.Oper.Ip_unnumbered_mac_learned,
			"ip_unnumbered_peer_lla":     ret.DtInterfaceTrunkOper.Oper.Ip_unnumbered_peer_lla,
			"mtu":                        ret.DtInterfaceTrunkOper.Oper.Mtu,
		},
	}
}

func setSliceInterfaceTrunkOperOperTrunkMember(d []edpt.InterfaceTrunkOperOperTrunkMember) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["members"] = item.Members
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceTrunkOperOperIpv4_list(d []edpt.InterfaceTrunkOperOperIpv4_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr"] = item.Addr
		in["mask"] = item.Mask
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceTrunkOperOperIpv6_list(d []edpt.InterfaceTrunkOperOperIpv6_list) []map[string]interface{} {
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

func getObjectInterfaceTrunkOperOper(d []interface{}) edpt.InterfaceTrunkOperOper {

	count1 := len(d)
	var ret edpt.InterfaceTrunkOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.LineProtocol = in["line_protocol"].(string)
		ret.TrunkMember = getSliceInterfaceTrunkOperOperTrunkMember(in["trunk_member"].([]interface{}))
		ret.Hardware = in["hardware"].(string)
		ret.Address = in["address"].(string)
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.Ipv6LinkLocal = in["ipv6_link_local"].(string)
		ret.Ipv6LinkLocalPrefix = in["ipv6_link_local_prefix"].(string)
		ret.Ipv6LinkLocalType = in["ipv6_link_local_type"].(string)
		ret.Ipv6LinkLocalScope = in["ipv6_link_local_scope"].(string)
		ret.Ipv4_addr_count = in["ipv4_addr_count"].(int)
		ret.Ipv4_list = getSliceInterfaceTrunkOperOperIpv4_list(in["ipv4_list"].([]interface{}))
		ret.Ipv6_addr_count = in["ipv6_addr_count"].(int)
		ret.Ipv6_list = getSliceInterfaceTrunkOperOperIpv6_list(in["ipv6_list"].([]interface{}))
		ret.IgmpQuerySent = in["igmp_query_sent"].(int)
		ret.IcmpRateLimitCurrent = in["icmp_rate_limit_current"].(int)
		ret.IcmpRateOverLimitDrop = in["icmp_rate_over_limit_drop"].(int)
		ret.Icmp6RateLimitCurrent = in["icmp6_rate_limit_current"].(int)
		ret.Icmp6RateOverLimitDrop = in["icmp6_rate_over_limit_drop"].(int)
		ret.VlanId = in["vlan_id"].(int)
		ret.Ip_unnumbered_oper = in["ip_unnumbered_oper"].(int)
		ret.Ip_unnumbered_enabled = in["ip_unnumbered_enabled"].(int)
		ret.Ip_unnumbered_mac_learned = in["ip_unnumbered_mac_learned"].(int)
		ret.Ip_unnumbered_peer_lla = in["ip_unnumbered_peer_lla"].(string)
		ret.Mtu = in["mtu"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkOperOperTrunkMember(d []interface{}) []edpt.InterfaceTrunkOperOperTrunkMember {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkOperOperTrunkMember, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkOperOperTrunkMember
		oi.Members = in["members"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkOperOperIpv4_list(d []interface{}) []edpt.InterfaceTrunkOperOperIpv4_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkOperOperIpv4_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkOperOperIpv4_list
		oi.Addr = in["addr"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkOperOperIpv6_list(d []interface{}) []edpt.InterfaceTrunkOperOperIpv6_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkOperOperIpv6_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkOperOperIpv6_list
		oi.Addr = in["addr"].(string)
		oi.Prefix = in["prefix"].(string)
		oi.Is_anycast = in["is_anycast"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTrunkOper(d *schema.ResourceData) edpt.InterfaceTrunkOper {
	var ret edpt.InterfaceTrunkOper

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Oper = getObjectInterfaceTrunkOperOper(d.Get("oper").([]interface{}))
	return ret
}
