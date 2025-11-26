package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMemoryBlockDebugOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_memory_block_debug_oper`: Operational Status for the object memory-block-debug\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemMemoryBlockDebugOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"a10_mem_id_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"block_debug_arr_zero": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"block_debug_arr_one": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"block_debug_arr_two": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"block_debug_arr_three": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"all_everything": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"module_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"alloc_zero": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_one": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_two": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_three": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"free_four": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"free_five": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"free_six": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"free_seven": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"block_debug_alloc_over": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"block_debug_free_over": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"block_debug_alloc_max_over": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"debug_decision_over_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"debug_decision_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"debug_decision_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"thread_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"a10_mem": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"t_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_module_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"t_alloc_zero": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_alloc_one": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_alloc_two": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_alloc_three": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_free_four": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_free_five": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_free_six": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_free_seven": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_block_debug_alloc_over": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_block_debug_free_over": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_block_debug_alloc_max_over": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_debug_decision_over_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_debug_decision_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"t_debug_decision_free": {
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

func resourceSystemMemoryBlockDebugOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryBlockDebugOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemoryBlockDebugOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemMemoryBlockDebugOperOper := setObjectSystemMemoryBlockDebugOperOper(res)
		d.Set("oper", SystemMemoryBlockDebugOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemMemoryBlockDebugOperOper(ret edpt.DataSystemMemoryBlockDebugOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"a10_mem_id_max":        ret.DtSystemMemoryBlockDebugOper.Oper.A10MemIdMax,
			"block_debug_arr_zero":  ret.DtSystemMemoryBlockDebugOper.Oper.BlockDebugArrZero,
			"block_debug_arr_one":   ret.DtSystemMemoryBlockDebugOper.Oper.BlockDebugArrOne,
			"block_debug_arr_two":   ret.DtSystemMemoryBlockDebugOper.Oper.BlockDebugArrTwo,
			"block_debug_arr_three": ret.DtSystemMemoryBlockDebugOper.Oper.BlockDebugArrThree,
			"all_everything":        setSliceSystemMemoryBlockDebugOperOperAllEverything(ret.DtSystemMemoryBlockDebugOper.Oper.AllEverything),
			"thread_id":             ret.DtSystemMemoryBlockDebugOper.Oper.ThreadId,
			"a10_mem":               setSliceSystemMemoryBlockDebugOperOperA10Mem(ret.DtSystemMemoryBlockDebugOper.Oper.A10Mem),
		},
	}
}

func setSliceSystemMemoryBlockDebugOperOperAllEverything(d []edpt.SystemMemoryBlockDebugOperOperAllEverything) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["id1"] = item.Id1
		in["module_name"] = item.ModuleName
		in["alloc_zero"] = item.AllocZero
		in["alloc_one"] = item.AllocOne
		in["alloc_two"] = item.AllocTwo
		in["alloc_three"] = item.AllocThree
		in["free_four"] = item.FreeFour
		in["free_five"] = item.FreeFive
		in["free_six"] = item.FreeSix
		in["free_seven"] = item.FreeSeven
		in["block_debug_alloc_over"] = item.BlockDebugAllocOver
		in["block_debug_free_over"] = item.BlockDebugFreeOver
		in["block_debug_alloc_max_over"] = item.BlockDebugAllocMaxOver
		in["debug_decision_over_size"] = item.DebugDecisionOverSize
		in["debug_decision_alloc"] = item.DebugDecisionAlloc
		in["debug_decision_free"] = item.DebugDecisionFree
		result = append(result, in)
	}
	return result
}

func setSliceSystemMemoryBlockDebugOperOperA10Mem(d []edpt.SystemMemoryBlockDebugOperOperA10Mem) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["t_id"] = item.TId
		in["t_module_name"] = item.TModuleName
		in["t_alloc_zero"] = item.TAllocZero
		in["t_alloc_one"] = item.TAllocOne
		in["t_alloc_two"] = item.TAllocTwo
		in["t_alloc_three"] = item.TAllocThree
		in["t_free_four"] = item.TFreeFour
		in["t_free_five"] = item.TFreeFive
		in["t_free_six"] = item.TFreeSix
		in["t_free_seven"] = item.TFreeSeven
		in["t_block_debug_alloc_over"] = item.TBlockDebugAllocOver
		in["t_block_debug_free_over"] = item.TBlockDebugFreeOver
		in["t_block_debug_alloc_max_over"] = item.TBlockDebugAllocMaxOver
		in["t_debug_decision_over_size"] = item.TDebugDecisionOverSize
		in["t_debug_decision_alloc"] = item.TDebugDecisionAlloc
		in["t_debug_decision_free"] = item.TDebugDecisionFree
		result = append(result, in)
	}
	return result
}

func getObjectSystemMemoryBlockDebugOperOper(d []interface{}) edpt.SystemMemoryBlockDebugOperOper {

	count1 := len(d)
	var ret edpt.SystemMemoryBlockDebugOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.A10MemIdMax = in["a10_mem_id_max"].(int)
		ret.BlockDebugArrZero = in["block_debug_arr_zero"].(string)
		ret.BlockDebugArrOne = in["block_debug_arr_one"].(string)
		ret.BlockDebugArrTwo = in["block_debug_arr_two"].(string)
		ret.BlockDebugArrThree = in["block_debug_arr_three"].(string)
		ret.AllEverything = getSliceSystemMemoryBlockDebugOperOperAllEverything(in["all_everything"].([]interface{}))
		ret.ThreadId = in["thread_id"].(int)
		ret.A10Mem = getSliceSystemMemoryBlockDebugOperOperA10Mem(in["a10_mem"].([]interface{}))
	}
	return ret
}

func getSliceSystemMemoryBlockDebugOperOperAllEverything(d []interface{}) []edpt.SystemMemoryBlockDebugOperOperAllEverything {

	count1 := len(d)
	ret := make([]edpt.SystemMemoryBlockDebugOperOperAllEverything, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemoryBlockDebugOperOperAllEverything
		oi.Id1 = in["id1"].(int)
		oi.ModuleName = in["module_name"].(string)
		oi.AllocZero = in["alloc_zero"].(int)
		oi.AllocOne = in["alloc_one"].(int)
		oi.AllocTwo = in["alloc_two"].(int)
		oi.AllocThree = in["alloc_three"].(int)
		oi.FreeFour = in["free_four"].(int)
		oi.FreeFive = in["free_five"].(int)
		oi.FreeSix = in["free_six"].(int)
		oi.FreeSeven = in["free_seven"].(int)
		oi.BlockDebugAllocOver = in["block_debug_alloc_over"].(int)
		oi.BlockDebugFreeOver = in["block_debug_free_over"].(int)
		oi.BlockDebugAllocMaxOver = in["block_debug_alloc_max_over"].(int)
		oi.DebugDecisionOverSize = in["debug_decision_over_size"].(int)
		oi.DebugDecisionAlloc = in["debug_decision_alloc"].(int)
		oi.DebugDecisionFree = in["debug_decision_free"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMemoryBlockDebugOperOperA10Mem(d []interface{}) []edpt.SystemMemoryBlockDebugOperOperA10Mem {

	count1 := len(d)
	ret := make([]edpt.SystemMemoryBlockDebugOperOperA10Mem, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemoryBlockDebugOperOperA10Mem
		oi.TId = in["t_id"].(int)
		oi.TModuleName = in["t_module_name"].(string)
		oi.TAllocZero = in["t_alloc_zero"].(int)
		oi.TAllocOne = in["t_alloc_one"].(int)
		oi.TAllocTwo = in["t_alloc_two"].(int)
		oi.TAllocThree = in["t_alloc_three"].(int)
		oi.TFreeFour = in["t_free_four"].(int)
		oi.TFreeFive = in["t_free_five"].(int)
		oi.TFreeSix = in["t_free_six"].(int)
		oi.TFreeSeven = in["t_free_seven"].(int)
		oi.TBlockDebugAllocOver = in["t_block_debug_alloc_over"].(int)
		oi.TBlockDebugFreeOver = in["t_block_debug_free_over"].(int)
		oi.TBlockDebugAllocMaxOver = in["t_block_debug_alloc_max_over"].(int)
		oi.TDebugDecisionOverSize = in["t_debug_decision_over_size"].(int)
		oi.TDebugDecisionAlloc = in["t_debug_decision_alloc"].(int)
		oi.TDebugDecisionFree = in["t_debug_decision_free"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemMemoryBlockDebugOper(d *schema.ResourceData) edpt.SystemMemoryBlockDebugOper {
	var ret edpt.SystemMemoryBlockDebugOper

	ret.Oper = getObjectSystemMemoryBlockDebugOperOper(d.Get("oper").([]interface{}))
	return ret
}
