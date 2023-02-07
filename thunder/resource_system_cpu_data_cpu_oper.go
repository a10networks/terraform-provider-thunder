package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuDataCpuOper() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_cpu_data_cpu_oper`: Operational Status for the object data-cpu\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemCpuDataCpuOperCreate,
		UpdateContext: resourceSystemCpuDataCpuOperUpdate,
		ReadContext:   resourceSystemCpuDataCpuOperRead,
		DeleteContext: resourceSystemCpuDataCpuOperDelete,
		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number_of_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"number_of_data_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cpu_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"1_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"5_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"10_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"30_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"60_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcpu_str": {
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

func resourceSystemCpuDataCpuOperCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuDataCpuOperCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuDataCpuOper(d)
		d.SetId(obj.GetId())
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuDataCpuOperRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemCpuDataCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuDataCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuDataCpuOper(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemCpuDataCpuOperUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuDataCpuOperUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuDataCpuOper(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuDataCpuOperRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemCpuDataCpuOperDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuDataCpuOperDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuDataCpuOper(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemCpuDataCpuOperOper(d []interface{}) edpt.SystemCpuDataCpuOperOper {
	count := len(d)
	var ret edpt.SystemCpuDataCpuOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.NumberOfCpu = in["number_of_cpu"].(int)
		ret.NumberOfDataCpu = in["number_of_data_cpu"].(int)
		ret.CpuUsage = getSliceSystemCpuDataCpuOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuDataCpuOperOperCpuUsage(d []interface{}) []edpt.SystemCpuDataCpuOperOperCpuUsage {
	count := len(d)
	ret := make([]edpt.SystemCpuDataCpuOperOperCpuUsage, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuDataCpuOperOperCpuUsage
		oi.CpuId = in["cpu_id"].(int)
		oi.Sec1 = in["1_sec"].(int)
		oi.Sec5 = in["5_sec"].(int)
		oi.Sec10 = in["10_sec"].(int)
		oi.Sec30 = in["30_sec"].(int)
		oi.Sec60 = in["60_sec"].(int)
		oi.DcpuStr = in["dcpu_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuDataCpuOper(d *schema.ResourceData) edpt.SystemCpuDataCpuOper {
	var ret edpt.SystemCpuDataCpuOper
	ret.Inst.Oper = getObjectSystemCpuDataCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
