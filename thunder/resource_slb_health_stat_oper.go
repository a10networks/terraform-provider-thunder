package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthStatOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_stat_oper`: Operational Status for the object health-stat\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthStatOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"health_monitor": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"up_cause": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"down_cause": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"down_state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reason": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retries": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"up_retries": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"partition_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ssl_version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ssl_cipher": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ssl_ticket": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"num_pins": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_pins_stat_up": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_pins_stat_down": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_pins_stat_unkn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_pins_stat_else": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_ssl_tickets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_stat": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"method": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"clear_ssl_ticket": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"monitor": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthStatOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthStatOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthStatOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthStatOperOper := setObjectSlbHealthStatOperOper(res)
		d.Set("oper", SlbHealthStatOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthStatOperOper(ret edpt.DataSlbHealthStatOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"health_check_list":  setSliceSlbHealthStatOperOperHealthCheckList(ret.DtSlbHealthStatOper.Oper.HealthCheckList),
			"num_pins":           ret.DtSlbHealthStatOper.Oper.Num_pins,
			"num_pins_stat_up":   ret.DtSlbHealthStatOper.Oper.Num_pins_stat_up,
			"num_pins_stat_down": ret.DtSlbHealthStatOper.Oper.Num_pins_stat_down,
			"num_pins_stat_unkn": ret.DtSlbHealthStatOper.Oper.Num_pins_stat_unkn,
			"num_pins_stat_else": ret.DtSlbHealthStatOper.Oper.Num_pins_stat_else,
			"num_ssl_tickets":    ret.DtSlbHealthStatOper.Oper.Num_ssl_tickets,
			"total_stat":         ret.DtSlbHealthStatOper.Oper.Total_stat,
			"method":             ret.DtSlbHealthStatOper.Oper.Method,
			"clear_ssl_ticket":   ret.DtSlbHealthStatOper.Oper.Clear_ssl_ticket,
			"monitor":            ret.DtSlbHealthStatOper.Oper.Monitor,
		},
	}
}

func setSliceSlbHealthStatOperOperHealthCheckList(d []edpt.SlbHealthStatOperOperHealthCheckList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_address"] = item.IpAddress
		in["port"] = item.Port
		in["health_monitor"] = item.HealthMonitor
		in["status"] = item.Status
		in["up_cause"] = item.UpCause
		in["down_cause"] = item.DownCause
		in["down_state"] = item.DownState
		in["reason"] = item.Reason
		in["total_retry"] = item.TotalRetry
		in["retries"] = item.Retries
		in["up_retries"] = item.UpRetries
		in["partition_id"] = item.PartitionId
		in["server"] = item.Server
		in["ssl_version"] = item.SslVersion
		in["ssl_cipher"] = item.SslCipher
		in["ssl_ticket"] = item.SslTicket
		result = append(result, in)
	}
	return result
}

func getObjectSlbHealthStatOperOper(d []interface{}) edpt.SlbHealthStatOperOper {

	count1 := len(d)
	var ret edpt.SlbHealthStatOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HealthCheckList = getSliceSlbHealthStatOperOperHealthCheckList(in["health_check_list"].([]interface{}))
		ret.Num_pins = in["num_pins"].(int)
		ret.Num_pins_stat_up = in["num_pins_stat_up"].(int)
		ret.Num_pins_stat_down = in["num_pins_stat_down"].(int)
		ret.Num_pins_stat_unkn = in["num_pins_stat_unkn"].(int)
		ret.Num_pins_stat_else = in["num_pins_stat_else"].(int)
		ret.Num_ssl_tickets = in["num_ssl_tickets"].(int)
		ret.Total_stat = in["total_stat"].(int)
		ret.Method = in["method"].(string)
		ret.Clear_ssl_ticket = in["clear_ssl_ticket"].(int)
		ret.Monitor = in["monitor"].(string)
	}
	return ret
}

func getSliceSlbHealthStatOperOperHealthCheckList(d []interface{}) []edpt.SlbHealthStatOperOperHealthCheckList {

	count1 := len(d)
	ret := make([]edpt.SlbHealthStatOperOperHealthCheckList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHealthStatOperOperHealthCheckList
		oi.IpAddress = in["ip_address"].(string)
		oi.Port = in["port"].(string)
		oi.HealthMonitor = in["health_monitor"].(string)
		oi.Status = in["status"].(string)
		oi.UpCause = in["up_cause"].(int)
		oi.DownCause = in["down_cause"].(int)
		oi.DownState = in["down_state"].(int)
		oi.Reason = in["reason"].(string)
		oi.TotalRetry = in["total_retry"].(int)
		oi.Retries = in["retries"].(int)
		oi.UpRetries = in["up_retries"].(int)
		oi.PartitionId = in["partition_id"].(int)
		oi.Server = in["server"].(string)
		oi.SslVersion = in["ssl_version"].(string)
		oi.SslCipher = in["ssl_cipher"].(string)
		oi.SslTicket = in["ssl_ticket"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHealthStatOper(d *schema.ResourceData) edpt.SlbHealthStatOper {
	var ret edpt.SlbHealthStatOper

	ret.Oper = getObjectSlbHealthStatOperOper(d.Get("oper").([]interface{}))
	return ret
}
