package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdDiskOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_disk_oper`: Operational Status for the object disk\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdDiskOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"end_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_disk": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"disk_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"disk_usage": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"alldynamic": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceRrdDiskOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdDiskOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdDiskOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdDiskOperOper := setObjectRrdDiskOperOper(res)
		d.Set("oper", RrdDiskOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdDiskOperOper(ret edpt.DataRrdDiskOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time": ret.DtRrdDiskOper.Oper.StartTime,
			"end_time":   ret.DtRrdDiskOper.Oper.EndTime,
			"total_disk": ret.DtRrdDiskOper.Oper.TotalDisk,
			"disk_usage": setSliceRrdDiskOperOperDiskUsage(ret.DtRrdDiskOper.Oper.DiskUsage),
			"alldynamic": ret.DtRrdDiskOper.Oper.Alldynamic,
		},
	}
}

func setSliceRrdDiskOperOperDiskUsage(d []edpt.RrdDiskOperOperDiskUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["disk_usage"] = item.DiskUsage
		result = append(result, in)
	}
	return result
}

func getObjectRrdDiskOperOper(d []interface{}) edpt.RrdDiskOperOper {

	count1 := len(d)
	var ret edpt.RrdDiskOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.TotalDisk = in["total_disk"].(string)
		ret.DiskUsage = getSliceRrdDiskOperOperDiskUsage(in["disk_usage"].([]interface{}))
		ret.Alldynamic = in["alldynamic"].(int)
	}
	return ret
}

func getSliceRrdDiskOperOperDiskUsage(d []interface{}) []edpt.RrdDiskOperOperDiskUsage {

	count1 := len(d)
	ret := make([]edpt.RrdDiskOperOperDiskUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdDiskOperOperDiskUsage
		oi.Time = in["time"].(int)
		oi.DiskUsage = in["disk_usage"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdDiskOper(d *schema.ResourceData) edpt.RrdDiskOper {
	var ret edpt.RrdDiskOper

	ret.Oper = getObjectRrdDiskOperOper(d.Get("oper").([]interface{}))
	return ret
}
