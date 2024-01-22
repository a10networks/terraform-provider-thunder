package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugNatMapOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_nat_map_oper`: Operational Status for the object nat-map\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugNatMapOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vnp_id_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat_map_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vnp_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"owner": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"active": {
													Type: schema.TypeInt, Optional: true, Description: "",
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
		},
	}
}

func resourceScaleoutDebugNatMapOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugNatMapOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugNatMapOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugNatMapOperOper := setObjectScaleoutDebugNatMapOperOper(res)
		d.Set("oper", ScaleoutDebugNatMapOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugNatMapOperOper(ret edpt.DataScaleoutDebugNatMapOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"device_id":   ret.DtScaleoutDebugNatMapOper.Oper.DeviceId,
			"vnp_id_list": setSliceScaleoutDebugNatMapOperOperVnp_id_list(ret.DtScaleoutDebugNatMapOper.Oper.Vnp_id_list),
		},
	}
}

func setSliceScaleoutDebugNatMapOperOperVnp_id_list(d []edpt.ScaleoutDebugNatMapOperOperVnp_id_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["nat_map_list"] = setSliceScaleoutDebugNatMapOperOperVnp_id_listNat_map_list(item.Nat_map_list)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugNatMapOperOperVnp_id_listNat_map_list(d []edpt.ScaleoutDebugNatMapOperOperVnp_id_listNat_map_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vnp_id"] = item.Vnp_id
		in["ip"] = item.Ip
		in["owner"] = item.Owner
		in["active"] = item.Active
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugNatMapOperOper(d []interface{}) edpt.ScaleoutDebugNatMapOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugNatMapOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DeviceId = in["device_id"].(int)
		ret.Vnp_id_list = getSliceScaleoutDebugNatMapOperOperVnp_id_list(in["vnp_id_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugNatMapOperOperVnp_id_list(d []interface{}) []edpt.ScaleoutDebugNatMapOperOperVnp_id_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugNatMapOperOperVnp_id_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugNatMapOperOperVnp_id_list
		oi.Nat_map_list = getSliceScaleoutDebugNatMapOperOperVnp_id_listNat_map_list(in["nat_map_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugNatMapOperOperVnp_id_listNat_map_list(d []interface{}) []edpt.ScaleoutDebugNatMapOperOperVnp_id_listNat_map_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugNatMapOperOperVnp_id_listNat_map_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugNatMapOperOperVnp_id_listNat_map_list
		oi.Vnp_id = in["vnp_id"].(int)
		oi.Ip = in["ip"].(string)
		oi.Owner = in["owner"].(int)
		oi.Active = in["active"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugNatMapOper(d *schema.ResourceData) edpt.ScaleoutDebugNatMapOper {
	var ret edpt.ScaleoutDebugNatMapOper

	ret.Oper = getObjectScaleoutDebugNatMapOperOper(d.Get("oper").([]interface{}))
	return ret
}
