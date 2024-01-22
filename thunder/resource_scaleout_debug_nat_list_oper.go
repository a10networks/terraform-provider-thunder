package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugNatListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_nat_list_oper`: Operational Status for the object nat-list\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugNatListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vnp_id_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nat_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"device": {
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

func resourceScaleoutDebugNatListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugNatListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugNatListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugNatListOperOper := setObjectScaleoutDebugNatListOperOper(res)
		d.Set("oper", ScaleoutDebugNatListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugNatListOperOper(ret edpt.DataScaleoutDebugNatListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vnp_id_list": setSliceScaleoutDebugNatListOperOperVnp_id_list(ret.DtScaleoutDebugNatListOper.Oper.Vnp_id_list),
		},
	}
}

func setSliceScaleoutDebugNatListOperOperVnp_id_list(d []edpt.ScaleoutDebugNatListOperOperVnp_id_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["nat_list"] = setSliceScaleoutDebugNatListOperOperVnp_id_listNat_list(item.Nat_list)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugNatListOperOperVnp_id_listNat_list(d []edpt.ScaleoutDebugNatListOperOperVnp_id_listNat_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		in["device"] = item.Device
		in["active"] = item.Active
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugNatListOperOper(d []interface{}) edpt.ScaleoutDebugNatListOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugNatListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vnp_id_list = getSliceScaleoutDebugNatListOperOperVnp_id_list(in["vnp_id_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugNatListOperOperVnp_id_list(d []interface{}) []edpt.ScaleoutDebugNatListOperOperVnp_id_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugNatListOperOperVnp_id_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugNatListOperOperVnp_id_list
		oi.Nat_list = getSliceScaleoutDebugNatListOperOperVnp_id_listNat_list(in["nat_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugNatListOperOperVnp_id_listNat_list(d []interface{}) []edpt.ScaleoutDebugNatListOperOperVnp_id_listNat_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugNatListOperOperVnp_id_listNat_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugNatListOperOperVnp_id_listNat_list
		oi.Ip = in["ip"].(string)
		oi.Device = in["device"].(int)
		oi.Active = in["active"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugNatListOper(d *schema.ResourceData) edpt.ScaleoutDebugNatListOper {
	var ret edpt.ScaleoutDebugNatListOper

	ret.Oper = getObjectScaleoutDebugNatListOperOper(d.Get("oper").([]interface{}))
	return ret
}
