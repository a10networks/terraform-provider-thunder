package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAPartitionVridAllStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_partition_vrid_all_status_oper`: Operational Status for the object partition-vrid-all-status\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAPartitionVridAllStatusOperRead,

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
									"standby_device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"standby_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"standby_weight": {
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

func resourceVrrpAPartitionVridAllStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPartitionVridAllStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPartitionVridAllStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAPartitionVridAllStatusOperOper := setObjectVrrpAPartitionVridAllStatusOperOper(res)
		d.Set("oper", VrrpAPartitionVridAllStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAPartitionVridAllStatusOperOper(ret edpt.DataVrrpAPartitionVridAllStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all_partition_list": setSliceVrrpAPartitionVridAllStatusOperOperAllPartitionList(ret.DtVrrpAPartitionVridAllStatusOper.Oper.AllPartitionList),
		},
	}
}

func setSliceVrrpAPartitionVridAllStatusOperOperAllPartitionList(d []edpt.VrrpAPartitionVridAllStatusOperOperAllPartitionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["local_device_id"] = item.Local_device_id
		in["partition_name"] = item.PartitionName
		in["vrid"] = item.Vrid
		in["active_device_id"] = item.Active_device_id
		in["active_priority"] = item.Active_priority
		in["active_weight"] = item.Active_weight
		in["standby_device_id"] = item.Standby_device_id
		in["standby_priority"] = item.Standby_priority
		in["standby_weight"] = item.Standby_weight
		result = append(result, in)
	}
	return result
}

func getObjectVrrpAPartitionVridAllStatusOperOper(d []interface{}) edpt.VrrpAPartitionVridAllStatusOperOper {

	count1 := len(d)
	var ret edpt.VrrpAPartitionVridAllStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllPartitionList = getSliceVrrpAPartitionVridAllStatusOperOperAllPartitionList(in["all_partition_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAPartitionVridAllStatusOperOperAllPartitionList(d []interface{}) []edpt.VrrpAPartitionVridAllStatusOperOperAllPartitionList {

	count1 := len(d)
	ret := make([]edpt.VrrpAPartitionVridAllStatusOperOperAllPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAPartitionVridAllStatusOperOperAllPartitionList
		oi.Local_device_id = in["local_device_id"].(int)
		oi.PartitionName = in["partition_name"].(string)
		oi.Vrid = in["vrid"].(int)
		oi.Active_device_id = in["active_device_id"].(int)
		oi.Active_priority = in["active_priority"].(int)
		oi.Active_weight = in["active_weight"].(int)
		oi.Standby_device_id = in["standby_device_id"].(int)
		oi.Standby_priority = in["standby_priority"].(int)
		oi.Standby_weight = in["standby_weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAPartitionVridAllStatusOper(d *schema.ResourceData) edpt.VrrpAPartitionVridAllStatusOper {
	var ret edpt.VrrpAPartitionVridAllStatusOper

	ret.Oper = getObjectVrrpAPartitionVridAllStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
