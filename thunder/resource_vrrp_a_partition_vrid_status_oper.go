package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAPartitionVridStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_partition_vrid_status_oper`: Operational Status for the object partition-vrid-status\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAPartitionVridStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_partition_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"partition_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_weight": {
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

func resourceVrrpAPartitionVridStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPartitionVridStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPartitionVridStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAPartitionVridStatusOperOper := setObjectVrrpAPartitionVridStatusOperOper(res)
		d.Set("oper", VrrpAPartitionVridStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAPartitionVridStatusOperOper(ret edpt.DataVrrpAPartitionVridStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all_partition_list": setSliceVrrpAPartitionVridStatusOperOperAllPartitionList(ret.DtVrrpAPartitionVridStatusOper.Oper.AllPartitionList),
		},
	}
}

func setSliceVrrpAPartitionVridStatusOperOperAllPartitionList(d []edpt.VrrpAPartitionVridStatusOperOperAllPartitionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["local_device_id"] = item.Local_device_id
		in["partition_name"] = item.PartitionName
		in["vrid"] = item.Vrid
		in["active_device_id"] = item.Active_device_id
		in["active_priority"] = item.Active_priority
		in["active_weight"] = item.Active_weight
		result = append(result, in)
	}
	return result
}

func getObjectVrrpAPartitionVridStatusOperOper(d []interface{}) edpt.VrrpAPartitionVridStatusOperOper {

	count1 := len(d)
	var ret edpt.VrrpAPartitionVridStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllPartitionList = getSliceVrrpAPartitionVridStatusOperOperAllPartitionList(in["all_partition_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAPartitionVridStatusOperOperAllPartitionList(d []interface{}) []edpt.VrrpAPartitionVridStatusOperOperAllPartitionList {

	count1 := len(d)
	ret := make([]edpt.VrrpAPartitionVridStatusOperOperAllPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAPartitionVridStatusOperOperAllPartitionList
		oi.Local_device_id = in["local_device_id"].(int)
		oi.PartitionName = in["partition_name"].(string)
		oi.Vrid = in["vrid"].(int)
		oi.Active_device_id = in["active_device_id"].(int)
		oi.Active_priority = in["active_priority"].(int)
		oi.Active_weight = in["active_weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAPartitionVridStatusOper(d *schema.ResourceData) edpt.VrrpAPartitionVridStatusOper {
	var ret edpt.VrrpAPartitionVridStatusOper

	ret.Oper = getObjectVrrpAPartitionVridStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
