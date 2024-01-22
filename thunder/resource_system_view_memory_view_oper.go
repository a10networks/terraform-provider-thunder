package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewMemoryViewOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_memory_view_oper`: Operational Status for the object memory-view\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewMemoryViewOperRead,

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

func resourceSystemViewMemoryViewOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewMemoryViewOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewMemoryViewOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewMemoryViewOperOper := setObjectSystemViewMemoryViewOperOper(res)
		d.Set("oper", SystemViewMemoryViewOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewMemoryViewOperOper(ret edpt.DataSystemViewMemoryViewOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total":                ret.DtSystemViewMemoryViewOper.Oper.Total,
			"system_memory":        setSliceSystemViewMemoryViewOperOperSystemMemory(ret.DtSystemViewMemoryViewOper.Oper.SystemMemory),
			"system_memory_counts": ret.DtSystemViewMemoryViewOper.Oper.SystemMemoryCounts,
			"aflex_memory":         setSliceSystemViewMemoryViewOperOperAflexMemory(ret.DtSystemViewMemoryViewOper.Oper.AflexMemory),
			"aflex_memory_counts":  ret.DtSystemViewMemoryViewOper.Oper.AflexMemoryCounts,
			"ssl_memory":           setSliceSystemViewMemoryViewOperOperSslMemory(ret.DtSystemViewMemoryViewOper.Oper.SslMemory),
			"ssl_memory_counts":    ret.DtSystemViewMemoryViewOper.Oper.SslMemoryCounts,
			"n2_memory":            setSliceSystemViewMemoryViewOperOperN2Memory(ret.DtSystemViewMemoryViewOper.Oper.N2Memory),
			"n2_memory_counts":     ret.DtSystemViewMemoryViewOper.Oper.N2MemoryCounts,
			"tcp_memory":           setSliceSystemViewMemoryViewOperOperTcpMemory(ret.DtSystemViewMemoryViewOper.Oper.TcpMemory),
			"tcp_memory_counts":    ret.DtSystemViewMemoryViewOper.Oper.TcpMemoryCounts,
			"usage":                ret.DtSystemViewMemoryViewOper.Oper.Usage,
			"used":                 ret.DtSystemViewMemoryViewOper.Oper.Used,
			"free":                 ret.DtSystemViewMemoryViewOper.Oper.Free,
			"shared":               ret.DtSystemViewMemoryViewOper.Oper.Shared,
			"buffers":              ret.DtSystemViewMemoryViewOper.Oper.Buffers,
			"cached":               ret.DtSystemViewMemoryViewOper.Oper.Cached,
		},
	}
}

func setSliceSystemViewMemoryViewOperOperSystemMemory(d []edpt.SystemViewMemoryViewOperOperSystemMemory) []map[string]interface{} {
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

func setSliceSystemViewMemoryViewOperOperAflexMemory(d []edpt.SystemViewMemoryViewOperOperAflexMemory) []map[string]interface{} {
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

func setSliceSystemViewMemoryViewOperOperSslMemory(d []edpt.SystemViewMemoryViewOperOperSslMemory) []map[string]interface{} {
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

func setSliceSystemViewMemoryViewOperOperN2Memory(d []edpt.SystemViewMemoryViewOperOperN2Memory) []map[string]interface{} {
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

func setSliceSystemViewMemoryViewOperOperTcpMemory(d []edpt.SystemViewMemoryViewOperOperTcpMemory) []map[string]interface{} {
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

func getObjectSystemViewMemoryViewOperOper(d []interface{}) edpt.SystemViewMemoryViewOperOper {

	count1 := len(d)
	var ret edpt.SystemViewMemoryViewOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total = in["total"].(int)
		ret.SystemMemory = getSliceSystemViewMemoryViewOperOperSystemMemory(in["system_memory"].([]interface{}))
		ret.SystemMemoryCounts = in["system_memory_counts"].(int)
		ret.AflexMemory = getSliceSystemViewMemoryViewOperOperAflexMemory(in["aflex_memory"].([]interface{}))
		ret.AflexMemoryCounts = in["aflex_memory_counts"].(int)
		ret.SslMemory = getSliceSystemViewMemoryViewOperOperSslMemory(in["ssl_memory"].([]interface{}))
		ret.SslMemoryCounts = in["ssl_memory_counts"].(int)
		ret.N2Memory = getSliceSystemViewMemoryViewOperOperN2Memory(in["n2_memory"].([]interface{}))
		ret.N2MemoryCounts = in["n2_memory_counts"].(int)
		ret.TcpMemory = getSliceSystemViewMemoryViewOperOperTcpMemory(in["tcp_memory"].([]interface{}))
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

func getSliceSystemViewMemoryViewOperOperSystemMemory(d []interface{}) []edpt.SystemViewMemoryViewOperOperSystemMemory {

	count1 := len(d)
	ret := make([]edpt.SystemViewMemoryViewOperOperSystemMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewMemoryViewOperOperSystemMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemViewMemoryViewOperOperAflexMemory(d []interface{}) []edpt.SystemViewMemoryViewOperOperAflexMemory {

	count1 := len(d)
	ret := make([]edpt.SystemViewMemoryViewOperOperAflexMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewMemoryViewOperOperAflexMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemViewMemoryViewOperOperSslMemory(d []interface{}) []edpt.SystemViewMemoryViewOperOperSslMemory {

	count1 := len(d)
	ret := make([]edpt.SystemViewMemoryViewOperOperSslMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewMemoryViewOperOperSslMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemViewMemoryViewOperOperN2Memory(d []interface{}) []edpt.SystemViewMemoryViewOperOperN2Memory {

	count1 := len(d)
	ret := make([]edpt.SystemViewMemoryViewOperOperN2Memory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewMemoryViewOperOperN2Memory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemViewMemoryViewOperOperTcpMemory(d []interface{}) []edpt.SystemViewMemoryViewOperOperTcpMemory {

	count1 := len(d)
	ret := make([]edpt.SystemViewMemoryViewOperOperTcpMemory, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewMemoryViewOperOperTcpMemory
		oi.ObjectSize = in["object_size"].(int)
		oi.Allocated = in["allocated"].(int)
		oi.Max = in["max"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemViewMemoryViewOper(d *schema.ResourceData) edpt.SystemViewMemoryViewOper {
	var ret edpt.SystemViewMemoryViewOper

	ret.Oper = getObjectSystemViewMemoryViewOperOper(d.Get("oper").([]interface{}))
	return ret
}
