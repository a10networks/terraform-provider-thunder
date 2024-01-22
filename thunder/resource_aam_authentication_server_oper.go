package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_oper`: Operational Status for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerOperRead,

		Schema: map[string]*schema.Schema{
			"ldap": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ldaps_server_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ldap_uri": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ldaps_idle_conn_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ldaps_idle_conn_fd_list": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ldaps_inuse_conn_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ldaps_inuse_conn_fd_list": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ocsp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stats_clear_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rserver_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rport_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rserver_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"host": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hm": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"max_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rport_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"protocol": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_hm": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"port_max_conn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"port_weight": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sg_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"sg_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"sg_state": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"part_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"get_count": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"windows": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stats_clear_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"name": {
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

func resourceAamAuthenticationServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerOperLdap := setObjectAamAuthenticationServerOperLdap(res)
		d.Set("ldap", AamAuthenticationServerOperLdap)
		AamAuthenticationServerOperOcsp := setObjectAamAuthenticationServerOperOcsp(res)
		d.Set("ocsp", AamAuthenticationServerOperOcsp)
		AamAuthenticationServerOperOper := setObjectAamAuthenticationServerOperOper(res)
		d.Set("oper", AamAuthenticationServerOperOper)
		AamAuthenticationServerOperWindows := setObjectAamAuthenticationServerOperWindows(res)
		d.Set("windows", AamAuthenticationServerOperWindows)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerOperLdap(ret edpt.DataAamAuthenticationServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectAamAuthenticationServerOperLdapOper(ret.DtAamAuthenticationServerOper.Ldap.Oper),
		},
	}
}

func setObjectAamAuthenticationServerOperLdapOper(d edpt.AamAuthenticationServerOperLdapOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ldaps_server_list"] = setSliceAamAuthenticationServerOperLdapOperLdapsServerList(d.LdapsServerList)
	result = append(result, in)
	return result
}

func setSliceAamAuthenticationServerOperLdapOperLdapsServerList(d []edpt.AamAuthenticationServerOperLdapOperLdapsServerList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ldap_uri"] = item.LdapUri
		in["ldaps_idle_conn_num"] = item.LdapsIdleConnNum
		in["ldaps_idle_conn_fd_list"] = item.LdapsIdleConnFdList
		in["ldaps_inuse_conn_num"] = item.LdapsInuseConnNum
		in["ldaps_inuse_conn_fd_list"] = item.LdapsInuseConnFdList
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationServerOperOcsp(ret edpt.DataAamAuthenticationServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectAamAuthenticationServerOperOcspOper(ret.DtAamAuthenticationServerOper.Ocsp.Oper),
		},
	}
}

func setObjectAamAuthenticationServerOperOcspOper(d edpt.AamAuthenticationServerOperOcspOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["stats_clear_type"] = d.StatsClearType

	in["name"] = d.Name
	result = append(result, in)
	return result
}

func setObjectAamAuthenticationServerOperOper(ret edpt.DataAamAuthenticationServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rserver_count": ret.DtAamAuthenticationServerOper.Oper.RserverCount,
			"rport_count":   ret.DtAamAuthenticationServerOper.Oper.RportCount,
			"rserver_list":  setSliceAamAuthenticationServerOperOperRserverList(ret.DtAamAuthenticationServerOper.Oper.RserverList),
			"name":          ret.DtAamAuthenticationServerOper.Oper.Name,
			"part_id":       ret.DtAamAuthenticationServerOper.Oper.PartId,
			"get_count":     ret.DtAamAuthenticationServerOper.Oper.GetCount,
		},
	}
}

func setSliceAamAuthenticationServerOperOperRserverList(d []edpt.AamAuthenticationServerOperOperRserverList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["server_name"] = item.ServerName
		in["host"] = item.Host
		in["ip"] = item.Ip
		in["hm"] = item.Hm
		in["status"] = item.Status
		in["max_conn"] = item.MaxConn
		in["weight"] = item.Weight
		in["rport_list"] = setSliceAamAuthenticationServerOperOperRserverListRportList(item.RportList)
		result = append(result, in)
	}
	return result
}

func setSliceAamAuthenticationServerOperOperRserverListRportList(d []edpt.AamAuthenticationServerOperOperRserverListRportList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["port"] = item.Port
		in["protocol"] = item.Protocol
		in["port_state"] = item.PortState
		in["port_hm"] = item.PortHm
		in["port_status"] = item.PortStatus
		in["port_max_conn"] = item.PortMaxConn
		in["port_weight"] = item.PortWeight
		in["sg_list"] = setSliceAamAuthenticationServerOperOperRserverListRportListSgList(item.SgList)
		result = append(result, in)
	}
	return result
}

func setSliceAamAuthenticationServerOperOperRserverListRportListSgList(d []edpt.AamAuthenticationServerOperOperRserverListRportListSgList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sg_name"] = item.SgName
		in["sg_state"] = item.SgState
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationServerOperWindows(ret edpt.DataAamAuthenticationServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectAamAuthenticationServerOperWindowsOper(ret.DtAamAuthenticationServerOper.Windows.Oper),
		},
	}
}

func setObjectAamAuthenticationServerOperWindowsOper(d edpt.AamAuthenticationServerOperWindowsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["stats_clear_type"] = d.StatsClearType

	in["name"] = d.Name
	result = append(result, in)
	return result
}

func getObjectAamAuthenticationServerOperLdap(d []interface{}) edpt.AamAuthenticationServerOperLdap {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOperLdap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectAamAuthenticationServerOperLdapOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationServerOperLdapOper(d []interface{}) edpt.AamAuthenticationServerOperLdapOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOperLdapOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LdapsServerList = getSliceAamAuthenticationServerOperLdapOperLdapsServerList(in["ldaps_server_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationServerOperLdapOperLdapsServerList(d []interface{}) []edpt.AamAuthenticationServerOperLdapOperLdapsServerList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOperLdapOperLdapsServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOperLdapOperLdapsServerList
		oi.LdapUri = in["ldap_uri"].(string)
		oi.LdapsIdleConnNum = in["ldaps_idle_conn_num"].(int)
		oi.LdapsIdleConnFdList = in["ldaps_idle_conn_fd_list"].(string)
		oi.LdapsInuseConnNum = in["ldaps_inuse_conn_num"].(int)
		oi.LdapsInuseConnFdList = in["ldaps_inuse_conn_fd_list"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerOperOcsp(d []interface{}) edpt.AamAuthenticationServerOperOcsp {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOperOcsp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectAamAuthenticationServerOperOcspOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationServerOperOcspOper(d []interface{}) edpt.AamAuthenticationServerOperOcspOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOperOcspOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatsClearType = in["stats_clear_type"].(string)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectAamAuthenticationServerOperOper(d []interface{}) edpt.AamAuthenticationServerOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RserverCount = in["rserver_count"].(int)
		ret.RportCount = in["rport_count"].(int)
		ret.RserverList = getSliceAamAuthenticationServerOperOperRserverList(in["rserver_list"].([]interface{}))
		ret.Name = in["name"].(string)
		ret.PartId = in["part_id"].(int)
		ret.GetCount = in["get_count"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerOperOperRserverList(d []interface{}) []edpt.AamAuthenticationServerOperOperRserverList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOperOperRserverList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOperOperRserverList
		oi.ServerName = in["server_name"].(string)
		oi.Host = in["host"].(string)
		oi.Ip = in["ip"].(string)
		oi.Hm = in["hm"].(string)
		oi.Status = in["status"].(string)
		oi.MaxConn = in["max_conn"].(int)
		oi.Weight = in["weight"].(int)
		oi.RportList = getSliceAamAuthenticationServerOperOperRserverListRportList(in["rport_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerOperOperRserverListRportList(d []interface{}) []edpt.AamAuthenticationServerOperOperRserverListRportList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOperOperRserverListRportList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOperOperRserverListRportList
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.PortState = in["port_state"].(string)
		oi.PortHm = in["port_hm"].(string)
		oi.PortStatus = in["port_status"].(string)
		oi.PortMaxConn = in["port_max_conn"].(int)
		oi.PortWeight = in["port_weight"].(int)
		oi.SgList = getSliceAamAuthenticationServerOperOperRserverListRportListSgList(in["sg_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerOperOperRserverListRportListSgList(d []interface{}) []edpt.AamAuthenticationServerOperOperRserverListRportListSgList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOperOperRserverListRportListSgList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOperOperRserverListRportListSgList
		oi.SgName = in["sg_name"].(string)
		oi.SgState = in["sg_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerOperWindows(d []interface{}) edpt.AamAuthenticationServerOperWindows {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOperWindows
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectAamAuthenticationServerOperWindowsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationServerOperWindowsOper(d []interface{}) edpt.AamAuthenticationServerOperWindowsOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOperWindowsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatsClearType = in["stats_clear_type"].(string)
		ret.Name = in["name"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerOper(d *schema.ResourceData) edpt.AamAuthenticationServerOper {
	var ret edpt.AamAuthenticationServerOper

	ret.Ldap = getObjectAamAuthenticationServerOperLdap(d.Get("ldap").([]interface{}))

	ret.Ocsp = getObjectAamAuthenticationServerOperOcsp(d.Get("ocsp").([]interface{}))

	ret.Oper = getObjectAamAuthenticationServerOperOper(d.Get("oper").([]interface{}))

	ret.Windows = getObjectAamAuthenticationServerOperWindows(d.Get("windows").([]interface{}))
	return ret
}
