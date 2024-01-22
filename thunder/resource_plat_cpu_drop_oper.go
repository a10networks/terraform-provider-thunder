package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePlatCpuDropOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_plat_cpu_drop_oper`: Operational Status for the object plat-cpu-drop\n\n__PLACEHOLDER__",
		ReadContext: resourcePlatCpuDropOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fpga_seg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fpga_seg_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"drop_seg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"drop_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"drop_cnt": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"drop_count": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rate_limit_drp": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rate_limit_drop": {
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

func resourcePlatCpuDropOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePlatCpuDropOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPlatCpuDropOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		PlatCpuDropOperOper := setObjectPlatCpuDropOperOper(res)
		d.Set("oper", PlatCpuDropOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPlatCpuDropOperOper(ret edpt.DataPlatCpuDropOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"fpga_seg":       setSlicePlatCpuDropOperOperFpgaSeg(ret.DtPlatCpuDropOper.Oper.FpgaSeg),
			"drop_seg":       setSlicePlatCpuDropOperOperDropSeg(ret.DtPlatCpuDropOper.Oper.DropSeg),
			"rate_limit":     ret.DtPlatCpuDropOper.Oper.RateLimit,
			"rate_limit_drp": setSlicePlatCpuDropOperOperRateLimitDrp(ret.DtPlatCpuDropOper.Oper.RateLimitDrp),
		},
	}
}

func setSlicePlatCpuDropOperOperFpgaSeg(d []edpt.PlatCpuDropOperOperFpgaSeg) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fpga_seg_name"] = item.FpgaSegName
		result = append(result, in)
	}
	return result
}

func setSlicePlatCpuDropOperOperDropSeg(d []edpt.PlatCpuDropOperOperDropSeg) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drop_name"] = item.DropName
		in["drop_cnt"] = setSlicePlatCpuDropOperOperDropSegDropCnt(item.DropCnt)
		result = append(result, in)
	}
	return result
}

func setSlicePlatCpuDropOperOperDropSegDropCnt(d []edpt.PlatCpuDropOperOperDropSegDropCnt) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drop_count"] = item.DropCount
		result = append(result, in)
	}
	return result
}

func setSlicePlatCpuDropOperOperRateLimitDrp(d []edpt.PlatCpuDropOperOperRateLimitDrp) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rate_limit_drop"] = item.RateLimitDrop
		result = append(result, in)
	}
	return result
}

func getObjectPlatCpuDropOperOper(d []interface{}) edpt.PlatCpuDropOperOper {

	count1 := len(d)
	var ret edpt.PlatCpuDropOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FpgaSeg = getSlicePlatCpuDropOperOperFpgaSeg(in["fpga_seg"].([]interface{}))
		ret.DropSeg = getSlicePlatCpuDropOperOperDropSeg(in["drop_seg"].([]interface{}))
		ret.RateLimit = in["rate_limit"].(int)
		ret.RateLimitDrp = getSlicePlatCpuDropOperOperRateLimitDrp(in["rate_limit_drp"].([]interface{}))
	}
	return ret
}

func getSlicePlatCpuDropOperOperFpgaSeg(d []interface{}) []edpt.PlatCpuDropOperOperFpgaSeg {

	count1 := len(d)
	ret := make([]edpt.PlatCpuDropOperOperFpgaSeg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatCpuDropOperOperFpgaSeg
		oi.FpgaSegName = in["fpga_seg_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatCpuDropOperOperDropSeg(d []interface{}) []edpt.PlatCpuDropOperOperDropSeg {

	count1 := len(d)
	ret := make([]edpt.PlatCpuDropOperOperDropSeg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatCpuDropOperOperDropSeg
		oi.DropName = in["drop_name"].(string)
		oi.DropCnt = getSlicePlatCpuDropOperOperDropSegDropCnt(in["drop_cnt"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatCpuDropOperOperDropSegDropCnt(d []interface{}) []edpt.PlatCpuDropOperOperDropSegDropCnt {

	count1 := len(d)
	ret := make([]edpt.PlatCpuDropOperOperDropSegDropCnt, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatCpuDropOperOperDropSegDropCnt
		oi.DropCount = in["drop_count"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSlicePlatCpuDropOperOperRateLimitDrp(d []interface{}) []edpt.PlatCpuDropOperOperRateLimitDrp {

	count1 := len(d)
	ret := make([]edpt.PlatCpuDropOperOperRateLimitDrp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatCpuDropOperOperRateLimitDrp
		oi.RateLimitDrop = in["rate_limit_drop"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointPlatCpuDropOper(d *schema.ResourceData) edpt.PlatCpuDropOper {
	var ret edpt.PlatCpuDropOper

	ret.Oper = getObjectPlatCpuDropOperOper(d.Get("oper").([]interface{}))
	return ret
}
