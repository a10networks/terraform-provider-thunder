package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsChassisSlotSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vcs_chassis_slot_summary_oper`: Operational Status for the object slot-summary\n\n__PLACEHOLDER__",
		ReadContext: resourceVcsChassisSlotSummaryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcs_enabled": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"chassis_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"multicast_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"multicast_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vmaster_maintenance_interval": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vmaster_maintenance_left": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vcs_handshake_completed_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vcs_handshake_completed_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vcs_handshake_completed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"floating_ipv4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"floating_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"floating_ipv4_mask": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"floating_ipv6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"floating_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"floating_ipv6_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"member_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"location": {
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

func resourceVcsChassisSlotSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsChassisSlotSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsChassisSlotSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VcsChassisSlotSummaryOperOper := setObjectVcsChassisSlotSummaryOperOper(res)
		d.Set("oper", VcsChassisSlotSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVcsChassisSlotSummaryOperOper(ret edpt.DataVcsChassisSlotSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vcs_enabled":                  ret.DtVcsChassisSlotSummaryOper.Oper.VcsEnabled,
			"chassis_id":                   ret.DtVcsChassisSlotSummaryOper.Oper.ChassisId,
			"multicast_addr":               ret.DtVcsChassisSlotSummaryOper.Oper.MulticastAddr,
			"multicast_port":               ret.DtVcsChassisSlotSummaryOper.Oper.MulticastPort,
			"version":                      ret.DtVcsChassisSlotSummaryOper.Oper.Version,
			"vmaster_maintenance_interval": ret.DtVcsChassisSlotSummaryOper.Oper.VmasterMaintenanceInterval,
			"vmaster_maintenance_left":     ret.DtVcsChassisSlotSummaryOper.Oper.VmasterMaintenanceLeft,
			"vcs_handshake_completed_list": setSliceVcsChassisSlotSummaryOperOperVcsHandshakeCompletedList(ret.DtVcsChassisSlotSummaryOper.Oper.VcsHandshakeCompletedList),
			"floating_ipv4_list":           setSliceVcsChassisSlotSummaryOperOperFloatingIpv4List(ret.DtVcsChassisSlotSummaryOper.Oper.FloatingIpv4List),
			"floating_ipv6_list":           setSliceVcsChassisSlotSummaryOperOperFloatingIpv6List(ret.DtVcsChassisSlotSummaryOper.Oper.FloatingIpv6List),
			"member_list":                  setSliceVcsChassisSlotSummaryOperOperMemberList(ret.DtVcsChassisSlotSummaryOper.Oper.MemberList),
		},
	}
}

func setSliceVcsChassisSlotSummaryOperOperVcsHandshakeCompletedList(d []edpt.VcsChassisSlotSummaryOperOperVcsHandshakeCompletedList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vcs_handshake_completed_id"] = item.VcsHandshakeCompletedId
		in["vcs_handshake_completed"] = item.VcsHandshakeCompleted
		result = append(result, in)
	}
	return result
}

func setSliceVcsChassisSlotSummaryOperOperFloatingIpv4List(d []edpt.VcsChassisSlotSummaryOperOperFloatingIpv4List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["floating_ipv4"] = item.FloatingIpv4
		in["floating_ipv4_mask"] = item.FloatingIpv4Mask
		result = append(result, in)
	}
	return result
}

func setSliceVcsChassisSlotSummaryOperOperFloatingIpv6List(d []edpt.VcsChassisSlotSummaryOperOperFloatingIpv6List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["floating_ipv6"] = item.FloatingIpv6
		in["floating_ipv6_prefix"] = item.FloatingIpv6Prefix
		result = append(result, in)
	}
	return result
}

func setSliceVcsChassisSlotSummaryOperOperMemberList(d []edpt.VcsChassisSlotSummaryOperOperMemberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["id1"] = item.Id1
		in["state"] = item.State
		in["priority"] = item.Priority
		in["ip_list"] = setSliceVcsChassisSlotSummaryOperOperMemberListIpList(item.IpList)
		in["port"] = item.Port
		in["location"] = item.Location
		result = append(result, in)
	}
	return result
}

func setSliceVcsChassisSlotSummaryOperOperMemberListIpList(d []edpt.VcsChassisSlotSummaryOperOperMemberListIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		result = append(result, in)
	}
	return result
}

func getObjectVcsChassisSlotSummaryOperOper(d []interface{}) edpt.VcsChassisSlotSummaryOperOper {

	count1 := len(d)
	var ret edpt.VcsChassisSlotSummaryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VcsEnabled = in["vcs_enabled"].(string)
		ret.ChassisId = in["chassis_id"].(int)
		ret.MulticastAddr = in["multicast_addr"].(string)
		ret.MulticastPort = in["multicast_port"].(int)
		ret.Version = in["version"].(string)
		ret.VmasterMaintenanceInterval = in["vmaster_maintenance_interval"].(int)
		ret.VmasterMaintenanceLeft = in["vmaster_maintenance_left"].(int)
		ret.VcsHandshakeCompletedList = getSliceVcsChassisSlotSummaryOperOperVcsHandshakeCompletedList(in["vcs_handshake_completed_list"].([]interface{}))
		ret.FloatingIpv4List = getSliceVcsChassisSlotSummaryOperOperFloatingIpv4List(in["floating_ipv4_list"].([]interface{}))
		ret.FloatingIpv6List = getSliceVcsChassisSlotSummaryOperOperFloatingIpv6List(in["floating_ipv6_list"].([]interface{}))
		ret.MemberList = getSliceVcsChassisSlotSummaryOperOperMemberList(in["member_list"].([]interface{}))
	}
	return ret
}

func getSliceVcsChassisSlotSummaryOperOperVcsHandshakeCompletedList(d []interface{}) []edpt.VcsChassisSlotSummaryOperOperVcsHandshakeCompletedList {

	count1 := len(d)
	ret := make([]edpt.VcsChassisSlotSummaryOperOperVcsHandshakeCompletedList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsChassisSlotSummaryOperOperVcsHandshakeCompletedList
		oi.VcsHandshakeCompletedId = in["vcs_handshake_completed_id"].(int)
		oi.VcsHandshakeCompleted = in["vcs_handshake_completed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsChassisSlotSummaryOperOperFloatingIpv4List(d []interface{}) []edpt.VcsChassisSlotSummaryOperOperFloatingIpv4List {

	count1 := len(d)
	ret := make([]edpt.VcsChassisSlotSummaryOperOperFloatingIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsChassisSlotSummaryOperOperFloatingIpv4List
		oi.FloatingIpv4 = in["floating_ipv4"].(string)
		oi.FloatingIpv4Mask = in["floating_ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsChassisSlotSummaryOperOperFloatingIpv6List(d []interface{}) []edpt.VcsChassisSlotSummaryOperOperFloatingIpv6List {

	count1 := len(d)
	ret := make([]edpt.VcsChassisSlotSummaryOperOperFloatingIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsChassisSlotSummaryOperOperFloatingIpv6List
		oi.FloatingIpv6 = in["floating_ipv6"].(string)
		oi.FloatingIpv6Prefix = in["floating_ipv6_prefix"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsChassisSlotSummaryOperOperMemberList(d []interface{}) []edpt.VcsChassisSlotSummaryOperOperMemberList {

	count1 := len(d)
	ret := make([]edpt.VcsChassisSlotSummaryOperOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsChassisSlotSummaryOperOperMemberList
		oi.Id1 = in["id1"].(int)
		oi.State = in["state"].(string)
		oi.Priority = in["priority"].(int)
		oi.IpList = getSliceVcsChassisSlotSummaryOperOperMemberListIpList(in["ip_list"].([]interface{}))
		oi.Port = in["port"].(int)
		oi.Location = in["location"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsChassisSlotSummaryOperOperMemberListIpList(d []interface{}) []edpt.VcsChassisSlotSummaryOperOperMemberListIpList {

	count1 := len(d)
	ret := make([]edpt.VcsChassisSlotSummaryOperOperMemberListIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsChassisSlotSummaryOperOperMemberListIpList
		oi.Ip = in["ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsChassisSlotSummaryOper(d *schema.ResourceData) edpt.VcsChassisSlotSummaryOper {
	var ret edpt.VcsChassisSlotSummaryOper

	ret.Oper = getObjectVcsChassisSlotSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
