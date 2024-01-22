package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_status_oper`: Operational Status for the object status\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_role": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"role": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"device_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_local": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_master": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cluster_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"follow_shared_redirection": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"follow_shared_session_sync": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l2redirect": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l2redirect_valid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l2redirect_operational": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l2redirect_eth": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l2redirect_trunk": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"l2redirect_vlan": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"advertised_redirect_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"dest_redirect_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"direction": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"active_interface_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"advertised_session_sync_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"dest_session_sync_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"exclude_interface_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
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

func resourceScaleoutStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutStatusOperOper := setObjectScaleoutStatusOperOper(res)
		d.Set("oper", ScaleoutStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutStatusOperOper(ret edpt.DataScaleoutStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"db_role":                         ret.DtScaleoutStatusOper.Oper.Db_role,
			"role":                            ret.DtScaleoutStatusOper.Oper.Role,
			"device_list":                     setSliceScaleoutStatusOperOperDeviceList(ret.DtScaleoutStatusOper.Oper.DeviceList),
			"cluster_mode":                    ret.DtScaleoutStatusOper.Oper.ClusterMode,
			"follow_shared_redirection":       ret.DtScaleoutStatusOper.Oper.FollowSharedRedirection,
			"follow_shared_session_sync":      ret.DtScaleoutStatusOper.Oper.FollowSharedSessionSync,
			"l2redirect":                      ret.DtScaleoutStatusOper.Oper.L2redirect,
			"l2redirect_valid":                ret.DtScaleoutStatusOper.Oper.L2redirectValid,
			"l2redirect_operational":          ret.DtScaleoutStatusOper.Oper.L2redirectOperational,
			"l2redirect_eth":                  ret.DtScaleoutStatusOper.Oper.L2redirectEth,
			"l2redirect_trunk":                ret.DtScaleoutStatusOper.Oper.L2redirectTrunk,
			"l2redirect_vlan":                 ret.DtScaleoutStatusOper.Oper.L2redirectVlan,
			"advertised_redirect_ip_list":     setSliceScaleoutStatusOperOperAdvertisedRedirectIpList(ret.DtScaleoutStatusOper.Oper.AdvertisedRedirectIpList),
			"dest_redirect_ip_list":           setSliceScaleoutStatusOperOperDestRedirectIpList(ret.DtScaleoutStatusOper.Oper.DestRedirectIpList),
			"active_interface_list":           setSliceScaleoutStatusOperOperActiveInterfaceList(ret.DtScaleoutStatusOper.Oper.ActiveInterfaceList),
			"advertised_session_sync_ip_list": setSliceScaleoutStatusOperOperAdvertisedSessionSyncIpList(ret.DtScaleoutStatusOper.Oper.AdvertisedSessionSyncIpList),
			"dest_session_sync_ip_list":       setSliceScaleoutStatusOperOperDestSessionSyncIpList(ret.DtScaleoutStatusOper.Oper.DestSessionSyncIpList),
			"exclude_interface_ip_list":       setSliceScaleoutStatusOperOperExcludeInterfaceIpList(ret.DtScaleoutStatusOper.Oper.ExcludeInterfaceIpList),
		},
	}
}

func setSliceScaleoutStatusOperOperDeviceList(d []edpt.ScaleoutStatusOperOperDeviceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["id1"] = item.Id1
		in["state"] = item.State
		in["is_local"] = item.IsLocal
		in["is_master"] = item.IsMaster
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutStatusOperOperAdvertisedRedirectIpList(d []edpt.ScaleoutStatusOperOperAdvertisedRedirectIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutStatusOperOperDestRedirectIpList(d []edpt.ScaleoutStatusOperOperDestRedirectIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["device_id"] = item.DeviceId
		in["direction"] = item.Direction
		in["ip"] = item.Ip
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutStatusOperOperActiveInterfaceList(d []edpt.ScaleoutStatusOperOperActiveInterfaceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["interface"] = item.Interface
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutStatusOperOperAdvertisedSessionSyncIpList(d []edpt.ScaleoutStatusOperOperAdvertisedSessionSyncIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutStatusOperOperDestSessionSyncIpList(d []edpt.ScaleoutStatusOperOperDestSessionSyncIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["device_id"] = item.DeviceId
		in["ip"] = item.Ip
		in["ipv6"] = item.Ipv6
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutStatusOperOperExcludeInterfaceIpList(d []edpt.ScaleoutStatusOperOperExcludeInterfaceIpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutStatusOperOper(d []interface{}) edpt.ScaleoutStatusOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Db_role = in["db_role"].(string)
		ret.Role = in["role"].(string)
		ret.DeviceList = getSliceScaleoutStatusOperOperDeviceList(in["device_list"].([]interface{}))
		ret.ClusterMode = in["cluster_mode"].(string)
		ret.FollowSharedRedirection = in["follow_shared_redirection"].(int)
		ret.FollowSharedSessionSync = in["follow_shared_session_sync"].(int)
		ret.L2redirect = in["l2redirect"].(int)
		ret.L2redirectValid = in["l2redirect_valid"].(int)
		ret.L2redirectOperational = in["l2redirect_operational"].(int)
		ret.L2redirectEth = in["l2redirect_eth"].(int)
		ret.L2redirectTrunk = in["l2redirect_trunk"].(int)
		ret.L2redirectVlan = in["l2redirect_vlan"].(int)
		ret.AdvertisedRedirectIpList = getSliceScaleoutStatusOperOperAdvertisedRedirectIpList(in["advertised_redirect_ip_list"].([]interface{}))
		ret.DestRedirectIpList = getSliceScaleoutStatusOperOperDestRedirectIpList(in["dest_redirect_ip_list"].([]interface{}))
		ret.ActiveInterfaceList = getSliceScaleoutStatusOperOperActiveInterfaceList(in["active_interface_list"].([]interface{}))
		ret.AdvertisedSessionSyncIpList = getSliceScaleoutStatusOperOperAdvertisedSessionSyncIpList(in["advertised_session_sync_ip_list"].([]interface{}))
		ret.DestSessionSyncIpList = getSliceScaleoutStatusOperOperDestSessionSyncIpList(in["dest_session_sync_ip_list"].([]interface{}))
		ret.ExcludeInterfaceIpList = getSliceScaleoutStatusOperOperExcludeInterfaceIpList(in["exclude_interface_ip_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutStatusOperOperDeviceList(d []interface{}) []edpt.ScaleoutStatusOperOperDeviceList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutStatusOperOperDeviceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutStatusOperOperDeviceList
		oi.Id1 = in["id1"].(int)
		oi.State = in["state"].(string)
		oi.IsLocal = in["is_local"].(int)
		oi.IsMaster = in["is_master"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutStatusOperOperAdvertisedRedirectIpList(d []interface{}) []edpt.ScaleoutStatusOperOperAdvertisedRedirectIpList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutStatusOperOperAdvertisedRedirectIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutStatusOperOperAdvertisedRedirectIpList
		oi.Ip = in["ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutStatusOperOperDestRedirectIpList(d []interface{}) []edpt.ScaleoutStatusOperOperDestRedirectIpList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutStatusOperOperDestRedirectIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutStatusOperOperDestRedirectIpList
		oi.DeviceId = in["device_id"].(int)
		oi.Direction = in["direction"].(string)
		oi.Ip = in["ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutStatusOperOperActiveInterfaceList(d []interface{}) []edpt.ScaleoutStatusOperOperActiveInterfaceList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutStatusOperOperActiveInterfaceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutStatusOperOperActiveInterfaceList
		oi.Interface = in["interface"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutStatusOperOperAdvertisedSessionSyncIpList(d []interface{}) []edpt.ScaleoutStatusOperOperAdvertisedSessionSyncIpList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutStatusOperOperAdvertisedSessionSyncIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutStatusOperOperAdvertisedSessionSyncIpList
		oi.Ip = in["ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutStatusOperOperDestSessionSyncIpList(d []interface{}) []edpt.ScaleoutStatusOperOperDestSessionSyncIpList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutStatusOperOperDestSessionSyncIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutStatusOperOperDestSessionSyncIpList
		oi.DeviceId = in["device_id"].(int)
		oi.Ip = in["ip"].(string)
		oi.Ipv6 = in["ipv6"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutStatusOperOperExcludeInterfaceIpList(d []interface{}) []edpt.ScaleoutStatusOperOperExcludeInterfaceIpList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutStatusOperOperExcludeInterfaceIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutStatusOperOperExcludeInterfaceIpList
		oi.Ip = in["ip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutStatusOper(d *schema.ResourceData) edpt.ScaleoutStatusOper {
	var ret edpt.ScaleoutStatusOper

	ret.Oper = getObjectScaleoutStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
