package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMlbOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_mlb_oper`: Operational Status for the object mlb\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbMlbOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_msg_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_msg_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_conn_created": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_conn_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_conn_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_conn_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_conn_created": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_conn_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_conn_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_rerouted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mlb_dcmsg_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mlb_dcmsg_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mlb_dcmsg_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mlb_dcmsg_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mlb_dcmsg_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mlb_server_probe": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mlb_server_down": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbMlbOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMlbOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMlbOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbMlbOperOper := setObjectSlbMlbOperOper(res)
		d.Set("oper", SlbMlbOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbMlbOperOper(ret edpt.DataSlbMlbOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"l4_cpu_list": setSliceSlbMlbOperOperL4CpuList(ret.DtSlbMlbOper.Oper.L4CpuList),
			"cpu_count":   ret.DtSlbMlbOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbMlbOperOperL4CpuList(d []edpt.SlbMlbOperOperL4CpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["client_msg_sent"] = item.Client_msg_sent
		in["server_msg_received"] = item.Server_msg_received
		in["server_conn_created"] = item.Server_conn_created
		in["server_conn_rst"] = item.Server_conn_rst
		in["server_conn_failed"] = item.Server_conn_failed
		in["server_conn_closed"] = item.Server_conn_closed
		in["client_conn_created"] = item.Client_conn_created
		in["client_conn_closed"] = item.Client_conn_closed
		in["client_conn_not_found"] = item.Client_conn_not_found
		in["msg_dropped"] = item.Msg_dropped
		in["msg_rerouted"] = item.Msg_rerouted
		in["mlb_dcmsg_sent"] = item.Mlb_dcmsg_sent
		in["mlb_dcmsg_received"] = item.Mlb_dcmsg_received
		in["mlb_dcmsg_error"] = item.Mlb_dcmsg_error
		in["mlb_dcmsg_alloc"] = item.Mlb_dcmsg_alloc
		in["mlb_dcmsg_free"] = item.Mlb_dcmsg_free
		in["mlb_server_probe"] = item.Mlb_server_probe
		in["mlb_server_down"] = item.Mlb_server_down
		result = append(result, in)
	}
	return result
}

func getObjectSlbMlbOperOper(d []interface{}) edpt.SlbMlbOperOper {

	count1 := len(d)
	var ret edpt.SlbMlbOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4CpuList = getSliceSlbMlbOperOperL4CpuList(in["l4_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbMlbOperOperL4CpuList(d []interface{}) []edpt.SlbMlbOperOperL4CpuList {

	count1 := len(d)
	ret := make([]edpt.SlbMlbOperOperL4CpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbMlbOperOperL4CpuList
		oi.Client_msg_sent = in["client_msg_sent"].(int)
		oi.Server_msg_received = in["server_msg_received"].(int)
		oi.Server_conn_created = in["server_conn_created"].(int)
		oi.Server_conn_rst = in["server_conn_rst"].(int)
		oi.Server_conn_failed = in["server_conn_failed"].(int)
		oi.Server_conn_closed = in["server_conn_closed"].(int)
		oi.Client_conn_created = in["client_conn_created"].(int)
		oi.Client_conn_closed = in["client_conn_closed"].(int)
		oi.Client_conn_not_found = in["client_conn_not_found"].(int)
		oi.Msg_dropped = in["msg_dropped"].(int)
		oi.Msg_rerouted = in["msg_rerouted"].(int)
		oi.Mlb_dcmsg_sent = in["mlb_dcmsg_sent"].(int)
		oi.Mlb_dcmsg_received = in["mlb_dcmsg_received"].(int)
		oi.Mlb_dcmsg_error = in["mlb_dcmsg_error"].(int)
		oi.Mlb_dcmsg_alloc = in["mlb_dcmsg_alloc"].(int)
		oi.Mlb_dcmsg_free = in["mlb_dcmsg_free"].(int)
		oi.Mlb_server_probe = in["mlb_server_probe"].(int)
		oi.Mlb_server_down = in["mlb_server_down"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbMlbOper(d *schema.ResourceData) edpt.SlbMlbOper {
	var ret edpt.SlbMlbOper

	ret.Oper = getObjectSlbMlbOperOper(d.Get("oper").([]interface{}))
	return ret
}
