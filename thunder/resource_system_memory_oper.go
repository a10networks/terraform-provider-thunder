package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMemoryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_memory_oper`: Operational Status for the object memory\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemMemoryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"system_memory": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"allocated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"system_memory_counts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aflex_memory": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"allocated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"aflex_memory_counts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ssl_memory": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"allocated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"ssl_memory_counts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"n2_memory": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"allocated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"n2_memory_counts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tcp_memory": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"allocated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"tcp_memory_counts": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"usage": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"shared": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"buffers": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cached": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemMemoryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemoryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemMemoryOperOper := setObjectSystemMemoryOperOper(res)
		d.Set("oper", SystemMemoryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemMemoryOperOper(ret edpt.DataSystemMemoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total":                ret.DtSystemMemoryOper.Oper.Total,
			"system_memory":        setSliceSystemMemoryOperOperSystemMemory(ret.DtSystemMemoryOper.Oper.SystemMemory),
			"system_memory_counts": ret.DtSystemMemoryOper.Oper.SystemMemoryCounts,
			"aflex_memory":         setSliceSystemMemoryOperOperAflexMemory(ret.DtSystemMemoryOper.Oper.AflexMemory),
			"aflex_memory_counts":  ret.DtSystemMemoryOper.Oper.AflexMemoryCounts,
			"ssl_memory":           setSliceSystemMemoryOperOperSslMemory(ret.DtSystemMemoryOper.Oper.SslMemory),
			"ssl_memory_counts":    ret.DtSystemMemoryOper.Oper.SslMemoryCounts,
			"n2_memory":            setSliceSystemMemoryOperOperN2Memory(ret.DtSystemMemoryOper.Oper.N2Memory),
			"n2_memory_counts":     ret.DtSystemMemoryOper.Oper.N2MemoryCounts,
			"tcp_memory":           setSliceSystemMemoryOperOperTcpMemory(ret.DtSystemMemoryOper.Oper.TcpMemory),
			"tcp_memory_counts":    ret.DtSystemMemoryOper.Oper.TcpMemoryCounts,
			"usage":                ret.DtSystemMemoryOper.Oper.Usage,
			"used":                 ret.DtSystemMemoryOper.Oper.Used,
			"free":                 ret.DtSystemMemoryOper.Oper.Free,
			"shared":               ret.DtSystemMemoryOper.Oper.Shared,
			"buffers":              ret.DtSystemMemoryOper.Oper.Buffers,
			"cached":               ret.DtSystemMemoryOper.Oper.Cached,
		},
	}
}

func setSliceSystemMemoryOperOperSystemMemory(d []edpt.SystemMemoryOperOperSystemMemory) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["object_size"] = item.ObjectSize
		in["allocated"] = item.Allocated
		in["max"] = item.Max
		result = append(result, in)
	}
	return result
}

func setSliceSystemMemoryOperOperAflexMemory(d []edpt.SystemMemoryOperOperAflexMemory) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["object_size"] = item.ObjectSize
		in["allocated"] = item.Allocated
		in["max"] = item.Max
		result = append(result, in)
	}
	return result
}

func setSliceSystemMemoryOperOperSslMemory(d []edpt.SystemMemoryOperOperSslMemory) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["object_size"] = item.ObjectSize
		in["allocated"] = item.Allocated
		in["max"] = item.Max
		result = append(result, in)
	}
	return result
}

func setSliceSystemMemoryOperOperN2Memory(d []edpt.SystemMemoryOperOperN2Memory) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["object_size"] = item.ObjectSize
		in["allocated"] = item.Allocated
		in["max"] = item.Max
		result = append(result, in)
	}
	return result
}

func setSliceSystemMemoryOperOperTcpMemory(d []edpt.SystemMemoryOperOperTcpMemory) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["object_size"] = item.ObjectSize
		in["allocated"] = item.Allocated
		in["max"] = item.Max
		result = append(result, in)
	}
	return result
}

func getObjectSystemMemoryOperOper(d []interface{}) edpt.SystemMemoryOperOper {

	count1 := len(d)
	var ret edpt.SystemMemoryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total = in["total"].(int)
		ret.SystemMemory = getSliceSystemMemoryOperOperSystemMemory(in["system_memory"].([]interface{}))
		ret.SystemMemoryCounts = in["system_memory_counts"].(int)
		ret.AflexMemory = getSliceSystemMemoryOperOperAflexMemory(in["aflex_memory"].([]interface{}))
		ret.AflexMemoryCounts = in["aflex_memory_counts"].(int)
		ret.SslMemory = getSliceSystemMemoryOperOperSslMemory(in["ssl_memory"].([]interface{}))
		ret.SslMemoryCounts = in["ssl_memory_counts"].(int)
		ret.N2Memory = getSliceSystemMemoryOperOperN2Memory(in["n2_memory"].([]interface{}))
		ret.N2MemoryCounts = in["n2_memory_counts"].(int)
		ret.TcpMemory = getSliceSystemMemoryOperOperTcpMemory(in["tcp_memory"].([]interface{}))
		ret.TcpMemoryCounts = in["tcp_memory_counts"].(int)
		ret.Usage = in["usage"].(string)
		ret.Used = in["used"].(int)
		ret.Free = in["free"].(int)
		ret.Shared = in["shared"].(int)
		ret.Buffers = in["buffers"].(int)
		ret.Cached = in["cached"].(int)
	}
	return ret
}

func getSliceSystemMemoryOperOperSystemMemory(d []interface{}) []edpt.SystemMemoryOperOperSystemMemory {

	count1 := len(d)
	ret := make([]edpt.SystemMemoryOperOperSystemMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemoryOperOperSystemMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMemoryOperOperAflexMemory(d []interface{}) []edpt.SystemMemoryOperOperAflexMemory {

	count1 := len(d)
	ret := make([]edpt.SystemMemoryOperOperAflexMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemoryOperOperAflexMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMemoryOperOperSslMemory(d []interface{}) []edpt.SystemMemoryOperOperSslMemory {

	count1 := len(d)
	ret := make([]edpt.SystemMemoryOperOperSslMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemoryOperOperSslMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMemoryOperOperN2Memory(d []interface{}) []edpt.SystemMemoryOperOperN2Memory {

	count1 := len(d)
	ret := make([]edpt.SystemMemoryOperOperN2Memory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemoryOperOperN2Memory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemMemoryOperOperTcpMemory(d []interface{}) []edpt.SystemMemoryOperOperTcpMemory {

	count1 := len(d)
	ret := make([]edpt.SystemMemoryOperOperTcpMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemoryOperOperTcpMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemMemoryOper(d *schema.ResourceData) edpt.SystemMemoryOper {
	var ret edpt.SystemMemoryOper

	ret.Oper = getObjectSystemMemoryOperOper(d.Get("oper").([]interface{}))
	return ret
}
