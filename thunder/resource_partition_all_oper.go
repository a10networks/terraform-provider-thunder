package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartitionAllOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_partition_all_oper`: Operational Status for the object partition-all\n\n__PLACEHOLDER__",
		ReadContext: resourcePartitionAllOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"partition_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"parent_l3v": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"admin_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"active_partition_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"manageable": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourcePartitionAllOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionAllOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionAllOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		PartitionAllOperOper := setObjectPartitionAllOperOper(res)
		d.Set("oper", PartitionAllOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPartitionAllOperOper(ret edpt.DataPartitionAllOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"partition_list":         setSlicePartitionAllOperOperPartitionList(ret.DtPartitionAllOper.Oper.PartitionList),
			"active_partition_count": ret.DtPartitionAllOper.Oper.ActivePartitionCount,
			"manageable":             ret.DtPartitionAllOper.Oper.Manageable,
		},
	}
}

func setSlicePartitionAllOperOperPartitionList(d []edpt.PartitionAllOperOperPartitionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["partition_name"] = item.PartitionName
		in["partition_id"] = item.PartitionId
		in["partition_type"] = item.PartitionType
		in["parent_l3v"] = item.ParentL3v
		in["app_type"] = item.AppType
		in["admin_count"] = item.AdminCount
		in["status"] = item.Status
		result = append(result, in)
	}
	return result
}

func getObjectPartitionAllOperOper(d []interface{}) edpt.PartitionAllOperOper {

	count1 := len(d)
	var ret edpt.PartitionAllOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PartitionList = getSlicePartitionAllOperOperPartitionList(in["partition_list"].([]interface{}))
		ret.ActivePartitionCount = in["active_partition_count"].(int)
		ret.Manageable = in["manageable"].(int)
	}
	return ret
}

func getSlicePartitionAllOperOperPartitionList(d []interface{}) []edpt.PartitionAllOperOperPartitionList {

	count1 := len(d)
	ret := make([]edpt.PartitionAllOperOperPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PartitionAllOperOperPartitionList
		oi.PartitionName = in["partition_name"].(string)
		oi.PartitionId = in["partition_id"].(int)
		oi.PartitionType = in["partition_type"].(string)
		oi.ParentL3v = in["parent_l3v"].(string)
		oi.AppType = in["app_type"].(string)
		oi.AdminCount = in["admin_count"].(int)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointPartitionAllOper(d *schema.ResourceData) edpt.PartitionAllOper {
	var ret edpt.PartitionAllOper

	ret.Oper = getObjectPartitionAllOperOper(d.Get("oper").([]interface{}))
	return ret
}
