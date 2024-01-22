package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugResourceTrackingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_resource_tracking_oper`: Operational Status for the object resource-tracking\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugResourceTrackingOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scaleout_cluster_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"template_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"so_track_cleanup": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"template_count_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"template_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"template_valid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"threshold_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"threshold_count_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"thresholds": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"action": {
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
						"track_template_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"track_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"user_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ref_cnt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"current_cost": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"old_cost": {
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

func resourceScaleoutDebugResourceTrackingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugResourceTrackingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugResourceTrackingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugResourceTrackingOperOper := setObjectScaleoutDebugResourceTrackingOperOper(res)
		d.Set("oper", ScaleoutDebugResourceTrackingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugResourceTrackingOperOper(ret edpt.DataScaleoutDebugResourceTrackingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"scaleout_cluster_list": setSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_list(ret.DtScaleoutDebugResourceTrackingOper.Oper.Scaleout_cluster_list),
			"track_template_list":   setSliceScaleoutDebugResourceTrackingOperOperTrack_template_list(ret.DtScaleoutDebugResourceTrackingOper.Oper.Track_template_list),
		},
	}
}

func setSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_list(d []edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cluster_id"] = item.Cluster_id
		in["template_count"] = item.Template_count
		in["last_action"] = item.Last_action
		in["current_state"] = item.Current_state
		in["so_track_cleanup"] = item.So_track_cleanup
		in["template_count_list"] = setSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list(item.Template_count_list)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list(d []edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["template_name"] = item.Template_name
		in["template_valid"] = item.Template_valid
		in["threshold_count"] = item.Threshold_count
		in["threshold_count_list"] = setSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list(item.Threshold_count_list)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list(d []edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["thresholds"] = item.Thresholds
		in["action"] = item.Action
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugResourceTrackingOperOperTrack_template_list(d []edpt.ScaleoutDebugResourceTrackingOperOperTrack_template_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["track_name"] = item.Track_name
		in["user_index"] = item.User_index
		in["ref_cnt"] = item.Ref_cnt
		in["current_cost"] = item.Current_cost
		in["old_cost"] = item.Old_cost
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugResourceTrackingOperOper(d []interface{}) edpt.ScaleoutDebugResourceTrackingOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugResourceTrackingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Scaleout_cluster_list = getSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_list(in["scaleout_cluster_list"].([]interface{}))
		ret.Track_template_list = getSliceScaleoutDebugResourceTrackingOperOperTrack_template_list(in["track_template_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_list(d []interface{}) []edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_list
		oi.Cluster_id = in["cluster_id"].(int)
		oi.Template_count = in["template_count"].(int)
		oi.Last_action = in["last_action"].(string)
		oi.Current_state = in["current_state"].(string)
		oi.So_track_cleanup = in["so_track_cleanup"].(int)
		oi.Template_count_list = getSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list(in["template_count_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list(d []interface{}) []edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_list
		oi.Template_name = in["template_name"].(string)
		oi.Template_valid = in["template_valid"].(int)
		oi.Threshold_count = in["threshold_count"].(int)
		oi.Threshold_count_list = getSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list(in["threshold_count_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list(d []interface{}) []edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugResourceTrackingOperOperScaleout_cluster_listTemplate_count_listThreshold_count_list
		oi.Thresholds = in["thresholds"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugResourceTrackingOperOperTrack_template_list(d []interface{}) []edpt.ScaleoutDebugResourceTrackingOperOperTrack_template_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugResourceTrackingOperOperTrack_template_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugResourceTrackingOperOperTrack_template_list
		oi.Track_name = in["track_name"].(string)
		oi.User_index = in["user_index"].(int)
		oi.Ref_cnt = in["ref_cnt"].(int)
		oi.Current_cost = in["current_cost"].(int)
		oi.Old_cost = in["old_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugResourceTrackingOper(d *schema.ResourceData) edpt.ScaleoutDebugResourceTrackingOper {
	var ret edpt.ScaleoutDebugResourceTrackingOper

	ret.Oper = getObjectScaleoutDebugResourceTrackingOperOper(d.Get("oper").([]interface{}))
	return ret
}
