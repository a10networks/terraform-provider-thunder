package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAStateCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_state_cpu_oper`: Operational Status for the object state-cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAStateCpuOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_sync_msg_per_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"min_sync_msg_per_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_query_msg_per_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"min_query_msg_per_packet": {
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

func resourceVrrpAStateCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAStateCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAStateCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAStateCpuOperOper := setObjectVrrpAStateCpuOperOper(res)
		d.Set("oper", VrrpAStateCpuOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAStateCpuOperOper(ret edpt.DataVrrpAStateCpuOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cpu_usage": setSliceVrrpAStateCpuOperOperCpuUsage(ret.DtVrrpAStateCpuOper.Oper.CpuUsage),
		},
	}
}

func setSliceVrrpAStateCpuOperOperCpuUsage(d []edpt.VrrpAStateCpuOperOperCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_id"] = item.CpuId
		in["max_sync_msg_per_packet"] = item.MaxSyncMsgPerPacket
		in["min_sync_msg_per_packet"] = item.MinSyncMsgPerPacket
		in["max_query_msg_per_packet"] = item.MaxQueryMsgPerPacket
		in["min_query_msg_per_packet"] = item.MinQueryMsgPerPacket
		result = append(result, in)
	}
	return result
}

func getObjectVrrpAStateCpuOperOper(d []interface{}) edpt.VrrpAStateCpuOperOper {

	count1 := len(d)
	var ret edpt.VrrpAStateCpuOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CpuUsage = getSliceVrrpAStateCpuOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAStateCpuOperOperCpuUsage(d []interface{}) []edpt.VrrpAStateCpuOperOperCpuUsage {

	count1 := len(d)
	ret := make([]edpt.VrrpAStateCpuOperOperCpuUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAStateCpuOperOperCpuUsage
		oi.CpuId = in["cpu_id"].(int)
		oi.MaxSyncMsgPerPacket = in["max_sync_msg_per_packet"].(int)
		oi.MinSyncMsgPerPacket = in["min_sync_msg_per_packet"].(int)
		oi.MaxQueryMsgPerPacket = in["max_query_msg_per_packet"].(int)
		oi.MinQueryMsgPerPacket = in["min_query_msg_per_packet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAStateCpuOper(d *schema.ResourceData) edpt.VrrpAStateCpuOper {
	var ret edpt.VrrpAStateCpuOper

	ret.Oper = getObjectVrrpAStateCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
