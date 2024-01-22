package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsVcsSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vcs_vcs_summary_oper`: Operational Status for the object vcs-summary\n\n__PLACEHOLDER__",
		ReadContext: resourceVcsVcsSummaryOperRead,

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

func resourceVcsVcsSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVcsSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVcsSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VcsVcsSummaryOperOper := setObjectVcsVcsSummaryOperOper(res)
		d.Set("oper", VcsVcsSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVcsVcsSummaryOperOper(ret edpt.DataVcsVcsSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vcs_enabled":                  ret.DtVcsVcsSummaryOper.Oper.VcsEnabled,
			"chassis_id":                   ret.DtVcsVcsSummaryOper.Oper.ChassisId,
			"multicast_addr":               ret.DtVcsVcsSummaryOper.Oper.MulticastAddr,
			"multicast_port":               ret.DtVcsVcsSummaryOper.Oper.MulticastPort,
			"version":                      ret.DtVcsVcsSummaryOper.Oper.Version,
			"vmaster_maintenance_interval": ret.DtVcsVcsSummaryOper.Oper.VmasterMaintenanceInterval,
			"vmaster_maintenance_left":     ret.DtVcsVcsSummaryOper.Oper.VmasterMaintenanceLeft,
			"vcs_handshake_completed_list": setSliceVcsVcsSummaryOperOperVcsHandshakeCompletedList(ret.DtVcsVcsSummaryOper.Oper.VcsHandshakeCompletedList),
			"floating_ipv4_list":           setSliceVcsVcsSummaryOperOperFloatingIpv4List(ret.DtVcsVcsSummaryOper.Oper.FloatingIpv4List),
			"floating_ipv6_list":           setSliceVcsVcsSummaryOperOperFloatingIpv6List(ret.DtVcsVcsSummaryOper.Oper.FloatingIpv6List),
			"member_list":                  setSliceVcsVcsSummaryOperOperMemberList(ret.DtVcsVcsSummaryOper.Oper.MemberList),
		},
	}
}

func setSliceVcsVcsSummaryOperOperVcsHandshakeCompletedList(d []edpt.VcsVcsSummaryOperOperVcsHandshakeCompletedList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vcs_handshake_completed_id"] = item.VcsHandshakeCompletedId
		in["vcs_handshake_completed"] = item.VcsHandshakeCompleted
		result = append(result, in)
	}
	return result
}

func setSliceVcsVcsSummaryOperOperFloatingIpv4List(d []edpt.VcsVcsSummaryOperOperFloatingIpv4List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["floating_ipv4"] = item.FloatingIpv4
		in["floating_ipv4_mask"] = item.FloatingIpv4Mask
		result = append(result, in)
	}
	return result
}

func setSliceVcsVcsSummaryOperOperFloatingIpv6List(d []edpt.VcsVcsSummaryOperOperFloatingIpv6List) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["floating_ipv6"] = item.FloatingIpv6
		in["floating_ipv6_prefix"] = item.FloatingIpv6Prefix
		result = append(result, in)
	}
	return result
}

func setSliceVcsVcsSummaryOperOperMemberList(d []edpt.VcsVcsSummaryOperOperMemberList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["id1"] = item.Id1
		in["state"] = item.State
		in["priority"] = item.Priority
		in["ip_list"] = setSliceVcsVcsSummaryOperOperMemberListIpList(item.IpList)
		in["port"] = item.Port
		in["location"] = item.Location
		result = append(result, in)
	}
	return result
}

func setSliceVcsVcsSummaryOperOperMemberListIpList(d []edpt.VcsVcsSummaryOperOperMemberListIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		result = append(result, in)
	}
	return result
}

func getObjectVcsVcsSummaryOperOper(d []interface{}) edpt.VcsVcsSummaryOperOper {

	count1 := len(d)
	var ret edpt.VcsVcsSummaryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VcsEnabled = in["vcs_enabled"].(string)
		ret.ChassisId = in["chassis_id"].(int)
		ret.MulticastAddr = in["multicast_addr"].(string)
		ret.MulticastPort = in["multicast_port"].(int)
		ret.Version = in["version"].(string)
		ret.VmasterMaintenanceInterval = in["vmaster_maintenance_interval"].(int)
		ret.VmasterMaintenanceLeft = in["vmaster_maintenance_left"].(int)
		ret.VcsHandshakeCompletedList = getSliceVcsVcsSummaryOperOperVcsHandshakeCompletedList(in["vcs_handshake_completed_list"].([]interface{}))
		ret.FloatingIpv4List = getSliceVcsVcsSummaryOperOperFloatingIpv4List(in["floating_ipv4_list"].([]interface{}))
		ret.FloatingIpv6List = getSliceVcsVcsSummaryOperOperFloatingIpv6List(in["floating_ipv6_list"].([]interface{}))
		ret.MemberList = getSliceVcsVcsSummaryOperOperMemberList(in["member_list"].([]interface{}))
	}
	return ret
}

func getSliceVcsVcsSummaryOperOperVcsHandshakeCompletedList(d []interface{}) []edpt.VcsVcsSummaryOperOperVcsHandshakeCompletedList {

	count1 := len(d)
	ret := make([]edpt.VcsVcsSummaryOperOperVcsHandshakeCompletedList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVcsSummaryOperOperVcsHandshakeCompletedList
		oi.VcsHandshakeCompletedId = in["vcs_handshake_completed_id"].(int)
		oi.VcsHandshakeCompleted = in["vcs_handshake_completed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsVcsSummaryOperOperFloatingIpv4List(d []interface{}) []edpt.VcsVcsSummaryOperOperFloatingIpv4List {

	count1 := len(d)
	ret := make([]edpt.VcsVcsSummaryOperOperFloatingIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVcsSummaryOperOperFloatingIpv4List
		oi.FloatingIpv4 = in["floating_ipv4"].(string)
		oi.FloatingIpv4Mask = in["floating_ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsVcsSummaryOperOperFloatingIpv6List(d []interface{}) []edpt.VcsVcsSummaryOperOperFloatingIpv6List {

	count1 := len(d)
	ret := make([]edpt.VcsVcsSummaryOperOperFloatingIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVcsSummaryOperOperFloatingIpv6List
		oi.FloatingIpv6 = in["floating_ipv6"].(string)
		oi.FloatingIpv6Prefix = in["floating_ipv6_prefix"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsVcsSummaryOperOperMemberList(d []interface{}) []edpt.VcsVcsSummaryOperOperMemberList {

	count1 := len(d)
	ret := make([]edpt.VcsVcsSummaryOperOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVcsSummaryOperOperMemberList
		oi.Id1 = in["id1"].(int)
		oi.State = in["state"].(string)
		oi.Priority = in["priority"].(int)
		oi.IpList = getSliceVcsVcsSummaryOperOperMemberListIpList(in["ip_list"].([]interface{}))
		oi.Port = in["port"].(int)
		oi.Location = in["location"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsVcsSummaryOperOperMemberListIpList(d []interface{}) []edpt.VcsVcsSummaryOperOperMemberListIpList {

	count1 := len(d)
	ret := make([]edpt.VcsVcsSummaryOperOperMemberListIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsVcsSummaryOperOperMemberListIpList
		oi.Ip = in["ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsVcsSummaryOper(d *schema.ResourceData) edpt.VcsVcsSummaryOper {
	var ret edpt.VcsVcsSummaryOper

	ret.Oper = getObjectVcsVcsSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
