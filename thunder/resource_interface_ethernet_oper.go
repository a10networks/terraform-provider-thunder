package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_ethernet_oper`: Operational Status for the object ethernet\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceEthernetOperRead,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
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
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"actual_speed": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"config_duplexity": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"actual_duplexity": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"media_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
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
						"is_tagged": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vlan_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tagged_vlan_list": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"span_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"span_port_state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"is_lead_member": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_blocked": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"current_vnp_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_vnp_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_pristine": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rate_byte_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rate_byte_sent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rate_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rate_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"input_utilization": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"output_utilization": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"is_device_transparent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"incoming_pkts_mirrored": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"outgoing_pkts_mirrored": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"incoming_pkts_monitored": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"outgoing_pkts_monitored": {
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
						"last_count_clear_at": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"last_up_event_at": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"last_down_event_at": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceEthernetOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceEthernetOperOper := setObjectInterfaceEthernetOperOper(res)
		d.Set("oper", InterfaceEthernetOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceEthernetOperOper(ret edpt.DataInterfaceEthernetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                      ret.DtInterfaceEthernetOper.Oper.State,
			"line_protocol":              ret.DtInterfaceEthernetOper.Oper.Line_protocol,
			"link_type":                  ret.DtInterfaceEthernetOper.Oper.Link_type,
			"mac":                        ret.DtInterfaceEthernetOper.Oper.Mac,
			"config_speed":               ret.DtInterfaceEthernetOper.Oper.Config_speed,
			"actual_speed":               ret.DtInterfaceEthernetOper.Oper.Actual_speed,
			"config_duplexity":           ret.DtInterfaceEthernetOper.Oper.Config_duplexity,
			"actual_duplexity":           ret.DtInterfaceEthernetOper.Oper.Actual_duplexity,
			"media_type":                 ret.DtInterfaceEthernetOper.Oper.Media_type,
			"ipv4_address":               ret.DtInterfaceEthernetOper.Oper.Ipv4Address,
			"ipv4_netmask":               ret.DtInterfaceEthernetOper.Oper.Ipv4Netmask,
			"ipv4_addr_count":            ret.DtInterfaceEthernetOper.Oper.Ipv4_addr_count,
			"ipv4_list":                  setSliceInterfaceEthernetOperOperIpv4_list(ret.DtInterfaceEthernetOper.Oper.Ipv4_list),
			"ipv6_addr_count":            ret.DtInterfaceEthernetOper.Oper.Ipv6_addr_count,
			"ipv6_list":                  setSliceInterfaceEthernetOperOperIpv6_list(ret.DtInterfaceEthernetOper.Oper.Ipv6_list),
			"ipv6_link_local":            ret.DtInterfaceEthernetOper.Oper.Ipv6LinkLocal,
			"ipv6_link_local_prefix":     ret.DtInterfaceEthernetOper.Oper.Ipv6LinkLocalPrefix,
			"ipv6_link_local_type":       ret.DtInterfaceEthernetOper.Oper.Ipv6LinkLocalType,
			"ipv6_link_local_scope":      ret.DtInterfaceEthernetOper.Oper.Ipv6LinkLocalScope,
			"igmp_query_sent":            ret.DtInterfaceEthernetOper.Oper.IgmpQuerySent,
			"icmp_rate_limit_current":    ret.DtInterfaceEthernetOper.Oper.IcmpRateLimitCurrent,
			"icmp_rate_over_limit_drop":  ret.DtInterfaceEthernetOper.Oper.IcmpRateOverLimitDrop,
			"icmp6_rate_limit_current":   ret.DtInterfaceEthernetOper.Oper.Icmp6RateLimitCurrent,
			"icmp6_rate_over_limit_drop": ret.DtInterfaceEthernetOper.Oper.Icmp6RateOverLimitDrop,
			"is_tagged":                  ret.DtInterfaceEthernetOper.Oper.IsTagged,
			"vlan_id":                    ret.DtInterfaceEthernetOper.Oper.VlanId,
			"tagged_vlan_list":           ret.DtInterfaceEthernetOper.Oper.TaggedVlanList,
			"span_mode":                  ret.DtInterfaceEthernetOper.Oper.Span_mode,
			"span_port_state":            ret.DtInterfaceEthernetOper.Oper.Span_port_state,
			"is_lead_member":             ret.DtInterfaceEthernetOper.Oper.Is_lead_member,
			"is_blocked":                 ret.DtInterfaceEthernetOper.Oper.Is_blocked,
			"current_vnp_id":             ret.DtInterfaceEthernetOper.Oper.Current_vnp_id,
			"port_vnp_id":                ret.DtInterfaceEthernetOper.Oper.Port_vnp_id,
			"is_pristine":                ret.DtInterfaceEthernetOper.Oper.Is_pristine,
			"rate_byte_rcvd":             ret.DtInterfaceEthernetOper.Oper.Rate_byte_rcvd,
			"rate_byte_sent":             ret.DtInterfaceEthernetOper.Oper.Rate_byte_sent,
			"rate_pkt_rcvd":              ret.DtInterfaceEthernetOper.Oper.Rate_pkt_rcvd,
			"rate_pkt_sent":              ret.DtInterfaceEthernetOper.Oper.Rate_pkt_sent,
			"input_utilization":          ret.DtInterfaceEthernetOper.Oper.Input_utilization,
			"output_utilization":         ret.DtInterfaceEthernetOper.Oper.Output_utilization,
			"is_device_transparent":      ret.DtInterfaceEthernetOper.Oper.Is_device_transparent,
			"incoming_pkts_mirrored":     ret.DtInterfaceEthernetOper.Oper.Incoming_pkts_mirrored,
			"outgoing_pkts_mirrored":     ret.DtInterfaceEthernetOper.Oper.Outgoing_pkts_mirrored,
			"incoming_pkts_monitored":    ret.DtInterfaceEthernetOper.Oper.Incoming_pkts_monitored,
			"outgoing_pkts_monitored":    ret.DtInterfaceEthernetOper.Oper.Outgoing_pkts_monitored,
			"ip_unnumbered_oper":         ret.DtInterfaceEthernetOper.Oper.Ip_unnumbered_oper,
			"ip_unnumbered_enabled":      ret.DtInterfaceEthernetOper.Oper.Ip_unnumbered_enabled,
			"ip_unnumbered_mac_learned":  ret.DtInterfaceEthernetOper.Oper.Ip_unnumbered_mac_learned,
			"ip_unnumbered_peer_lla":     ret.DtInterfaceEthernetOper.Oper.Ip_unnumbered_peer_lla,
			"last_count_clear_at":        ret.DtInterfaceEthernetOper.Oper.Last_count_clear_at,
			"last_up_event_at":           ret.DtInterfaceEthernetOper.Oper.Last_up_event_at,
			"last_down_event_at":         ret.DtInterfaceEthernetOper.Oper.Last_down_event_at,
		},
	}
}

func setSliceInterfaceEthernetOperOperIpv4_list(d []edpt.InterfaceEthernetOperOperIpv4_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["addr"] = item.Addr
		in["mask"] = item.Mask
		result = append(result, in)
	}
	return result
}

func setSliceInterfaceEthernetOperOperIpv6_list(d []edpt.InterfaceEthernetOperOperIpv6_list) []map[string]interface{} {
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

func getObjectInterfaceEthernetOperOper(d []interface{}) edpt.InterfaceEthernetOperOper {

	count1 := len(d)
	var ret edpt.InterfaceEthernetOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Line_protocol = in["line_protocol"].(string)
		ret.Link_type = in["link_type"].(string)
		ret.Mac = in["mac"].(string)
		ret.Config_speed = in["config_speed"].(string)
		ret.Actual_speed = in["actual_speed"].(string)
		ret.Config_duplexity = in["config_duplexity"].(string)
		ret.Actual_duplexity = in["actual_duplexity"].(string)
		ret.Media_type = in["media_type"].(string)
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.Ipv4_addr_count = in["ipv4_addr_count"].(int)
		ret.Ipv4_list = getSliceInterfaceEthernetOperOperIpv4_list(in["ipv4_list"].([]interface{}))
		ret.Ipv6_addr_count = in["ipv6_addr_count"].(int)
		ret.Ipv6_list = getSliceInterfaceEthernetOperOperIpv6_list(in["ipv6_list"].([]interface{}))
		ret.Ipv6LinkLocal = in["ipv6_link_local"].(string)
		ret.Ipv6LinkLocalPrefix = in["ipv6_link_local_prefix"].(string)
		ret.Ipv6LinkLocalType = in["ipv6_link_local_type"].(string)
		ret.Ipv6LinkLocalScope = in["ipv6_link_local_scope"].(string)
		ret.IgmpQuerySent = in["igmp_query_sent"].(int)
		ret.IcmpRateLimitCurrent = in["icmp_rate_limit_current"].(int)
		ret.IcmpRateOverLimitDrop = in["icmp_rate_over_limit_drop"].(int)
		ret.Icmp6RateLimitCurrent = in["icmp6_rate_limit_current"].(int)
		ret.Icmp6RateOverLimitDrop = in["icmp6_rate_over_limit_drop"].(int)
		ret.IsTagged = in["is_tagged"].(int)
		ret.VlanId = in["vlan_id"].(int)
		ret.TaggedVlanList = in["tagged_vlan_list"].(string)
		ret.Span_mode = in["span_mode"].(string)
		ret.Span_port_state = in["span_port_state"].(string)
		ret.Is_lead_member = in["is_lead_member"].(int)
		ret.Is_blocked = in["is_blocked"].(int)
		ret.Current_vnp_id = in["current_vnp_id"].(int)
		ret.Port_vnp_id = in["port_vnp_id"].(int)
		ret.Is_pristine = in["is_pristine"].(int)
		ret.Rate_byte_rcvd = in["rate_byte_rcvd"].(int)
		ret.Rate_byte_sent = in["rate_byte_sent"].(int)
		ret.Rate_pkt_rcvd = in["rate_pkt_rcvd"].(int)
		ret.Rate_pkt_sent = in["rate_pkt_sent"].(int)
		ret.Input_utilization = in["input_utilization"].(int)
		ret.Output_utilization = in["output_utilization"].(int)
		ret.Is_device_transparent = in["is_device_transparent"].(int)
		ret.Incoming_pkts_mirrored = in["incoming_pkts_mirrored"].(int)
		ret.Outgoing_pkts_mirrored = in["outgoing_pkts_mirrored"].(int)
		ret.Incoming_pkts_monitored = in["incoming_pkts_monitored"].(int)
		ret.Outgoing_pkts_monitored = in["outgoing_pkts_monitored"].(int)
		ret.Ip_unnumbered_oper = in["ip_unnumbered_oper"].(int)
		ret.Ip_unnumbered_enabled = in["ip_unnumbered_enabled"].(int)
		ret.Ip_unnumbered_mac_learned = in["ip_unnumbered_mac_learned"].(int)
		ret.Ip_unnumbered_peer_lla = in["ip_unnumbered_peer_lla"].(string)
		ret.Last_count_clear_at = in["last_count_clear_at"].(string)
		ret.Last_up_event_at = in["last_up_event_at"].(string)
		ret.Last_down_event_at = in["last_down_event_at"].(string)
	}
	return ret
}

func getSliceInterfaceEthernetOperOperIpv4_list(d []interface{}) []edpt.InterfaceEthernetOperOperIpv4_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetOperOperIpv4_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetOperOperIpv4_list
		oi.Addr = in["addr"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetOperOperIpv6_list(d []interface{}) []edpt.InterfaceEthernetOperOperIpv6_list {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetOperOperIpv6_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetOperOperIpv6_list
		oi.Addr = in["addr"].(string)
		oi.Prefix = in["prefix"].(string)
		oi.Is_anycast = in["is_anycast"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceEthernetOper(d *schema.ResourceData) edpt.InterfaceEthernetOper {
	var ret edpt.InterfaceEthernetOper

	ret.Ifnum = d.Get("ifnum").(int)

	ret.Oper = getObjectInterfaceEthernetOperOper(d.Get("oper").([]interface{}))
	return ret
}
