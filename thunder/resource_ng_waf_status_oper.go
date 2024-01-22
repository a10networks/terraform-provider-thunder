package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNgWafStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ng_waf_status_oper`: Operational Status for the object status\n\n__PLACEHOLDER__",
		ReadContext: resourceNgWafStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ngwaf_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"partition_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"agent_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"access_key_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"secret_access_key": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cache_entries": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tracked_custom_signal": {
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

func resourceNgWafStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWafStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NgWafStatusOperOper := setObjectNgWafStatusOperOper(res)
		d.Set("oper", NgWafStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNgWafStatusOperOper(ret edpt.DataNgWafStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ngwaf_version":  ret.DtNgWafStatusOper.Oper.NgwafVersion,
			"partition_list": setSliceNgWafStatusOperOperPartitionList(ret.DtNgWafStatusOper.Oper.PartitionList),
		},
	}
}

func setSliceNgWafStatusOperOperPartitionList(d []edpt.NgWafStatusOperOperPartitionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["partition_name"] = item.PartitionName
		in["status"] = item.Status
		in["agent_name"] = item.AgentName
		in["access_key_id"] = item.AccessKeyId
		in["secret_access_key"] = item.SecretAccessKey
		in["cache_entries"] = item.CacheEntries
		in["tracked_custom_signal"] = item.TrackedCustomSignal
		result = append(result, in)
	}
	return result
}

func getObjectNgWafStatusOperOper(d []interface{}) edpt.NgWafStatusOperOper {

	count1 := len(d)
	var ret edpt.NgWafStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NgwafVersion = in["ngwaf_version"].(string)
		ret.PartitionList = getSliceNgWafStatusOperOperPartitionList(in["partition_list"].([]interface{}))
	}
	return ret
}

func getSliceNgWafStatusOperOperPartitionList(d []interface{}) []edpt.NgWafStatusOperOperPartitionList {

	count1 := len(d)
	ret := make([]edpt.NgWafStatusOperOperPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafStatusOperOperPartitionList
		oi.PartitionName = in["partition_name"].(string)
		oi.Status = in["status"].(string)
		oi.AgentName = in["agent_name"].(string)
		oi.AccessKeyId = in["access_key_id"].(string)
		oi.SecretAccessKey = in["secret_access_key"].(string)
		oi.CacheEntries = in["cache_entries"].(int)
		oi.TrackedCustomSignal = in["tracked_custom_signal"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNgWafStatusOper(d *schema.ResourceData) edpt.NgWafStatusOper {
	var ret edpt.NgWafStatusOper

	ret.Oper = getObjectNgWafStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
