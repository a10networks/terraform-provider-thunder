package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartitionAvailableIdOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_partition_available_id_oper`: Operational Status for the object partition-available-id\n\n__PLACEHOLDER__",
		ReadContext: resourcePartitionAvailableIdOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"end": {
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

func resourcePartitionAvailableIdOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionAvailableIdOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionAvailableIdOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		PartitionAvailableIdOperOper := setObjectPartitionAvailableIdOperOper(res)
		d.Set("oper", PartitionAvailableIdOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPartitionAvailableIdOperOper(ret edpt.DataPartitionAvailableIdOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"range_list": setSlicePartitionAvailableIdOperOperRangeList(ret.DtPartitionAvailableIdOper.Oper.RangeList),
		},
	}
}

func setSlicePartitionAvailableIdOperOperRangeList(d []edpt.PartitionAvailableIdOperOperRangeList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["start"] = item.Start
		in["end"] = item.End
		result = append(result, in)
	}
	return result
}

func getObjectPartitionAvailableIdOperOper(d []interface{}) edpt.PartitionAvailableIdOperOper {

	count1 := len(d)
	var ret edpt.PartitionAvailableIdOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RangeList = getSlicePartitionAvailableIdOperOperRangeList(in["range_list"].([]interface{}))
	}
	return ret
}

func getSlicePartitionAvailableIdOperOperRangeList(d []interface{}) []edpt.PartitionAvailableIdOperOperRangeList {

	count1 := len(d)
	ret := make([]edpt.PartitionAvailableIdOperOperRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PartitionAvailableIdOperOperRangeList
		oi.Start = in["start"].(string)
		oi.End = in["end"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointPartitionAvailableIdOper(d *schema.ResourceData) edpt.PartitionAvailableIdOper {
	var ret edpt.PartitionAvailableIdOper

	ret.Oper = getObjectPartitionAvailableIdOperOper(d.Get("oper").([]interface{}))
	return ret
}
