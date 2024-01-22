package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutTrafficMapOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_traffic_map_oper`: Operational Status for the object traffic-map\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutTrafficMapOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_server": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"virtual_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"map_entries_list_head": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"service_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user_grp_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"running_device_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"map_entries_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_group": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cur_active": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cur_standby": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"new_active": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"new_standby": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"tbl_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceScaleoutTrafficMapOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutTrafficMapOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutTrafficMapOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutTrafficMapOperOper := setObjectScaleoutTrafficMapOperOper(res)
		d.Set("oper", ScaleoutTrafficMapOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutTrafficMapOperOper(ret edpt.DataScaleoutTrafficMapOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"virtual_server":        ret.DtScaleoutTrafficMapOper.Oper.VirtualServer,
			"virtual_port":          ret.DtScaleoutTrafficMapOper.Oper.VirtualPort,
			"src_ip":                ret.DtScaleoutTrafficMapOper.Oper.SrcIp,
			"src_ipv6":              ret.DtScaleoutTrafficMapOper.Oper.SrcIpv6,
			"map_entries_list_head": setSliceScaleoutTrafficMapOperOperMapEntriesListHead(ret.DtScaleoutTrafficMapOper.Oper.MapEntriesListHead),
			"tbl_num":               ret.DtScaleoutTrafficMapOper.Oper.TblNum,
		},
	}
}

func setSliceScaleoutTrafficMapOperOperMapEntriesListHead(d []edpt.ScaleoutTrafficMapOperOperMapEntriesListHead) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["service_type"] = item.ServiceType
		in["service_name"] = item.ServiceName
		in["user_grp_num"] = item.UserGrpNum
		in["running_device_num"] = item.RunningDeviceNum
		in["map_entries_list"] = setSliceScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList(item.MapEntriesList)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList(d []edpt.ScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["user_group"] = item.UserGroup
		in["cur_active"] = item.CurActive
		in["cur_standby"] = item.CurStandby
		in["new_active"] = item.NewActive
		in["new_standby"] = item.NewStandby
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutTrafficMapOperOper(d []interface{}) edpt.ScaleoutTrafficMapOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutTrafficMapOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VirtualServer = in["virtual_server"].(string)
		ret.VirtualPort = in["virtual_port"].(int)
		ret.SrcIp = in["src_ip"].(string)
		ret.SrcIpv6 = in["src_ipv6"].(string)
		ret.MapEntriesListHead = getSliceScaleoutTrafficMapOperOperMapEntriesListHead(in["map_entries_list_head"].([]interface{}))
		ret.TblNum = in["tbl_num"].(int)
	}
	return ret
}

func getSliceScaleoutTrafficMapOperOperMapEntriesListHead(d []interface{}) []edpt.ScaleoutTrafficMapOperOperMapEntriesListHead {

	count1 := len(d)
	ret := make([]edpt.ScaleoutTrafficMapOperOperMapEntriesListHead, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutTrafficMapOperOperMapEntriesListHead
		oi.ServiceType = in["service_type"].(string)
		oi.ServiceName = in["service_name"].(string)
		oi.UserGrpNum = in["user_grp_num"].(int)
		oi.RunningDeviceNum = in["running_device_num"].(int)
		oi.MapEntriesList = getSliceScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList(in["map_entries_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList(d []interface{}) []edpt.ScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutTrafficMapOperOperMapEntriesListHeadMapEntriesList
		oi.UserGroup = in["user_group"].(int)
		oi.CurActive = in["cur_active"].(int)
		oi.CurStandby = in["cur_standby"].(int)
		oi.NewActive = in["new_active"].(int)
		oi.NewStandby = in["new_standby"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutTrafficMapOper(d *schema.ResourceData) edpt.ScaleoutTrafficMapOper {
	var ret edpt.ScaleoutTrafficMapOper

	ret.Oper = getObjectScaleoutTrafficMapOperOper(d.Get("oper").([]interface{}))
	return ret
}
