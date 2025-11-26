package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerTunnelStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_controller_tunnel_stats_oper`: Operational Status for the object tunnel-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceControllerTunnelStatsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"bytes_recieved": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"number_of_errors": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"uptime": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"error_message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceControllerTunnelStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerTunnelStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerTunnelStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ControllerTunnelStatsOperOper := setObjectControllerTunnelStatsOperOper(res)
		d.Set("oper", ControllerTunnelStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectControllerTunnelStatsOperOper(ret edpt.DataControllerTunnelStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status":           ret.DtControllerTunnelStatsOper.Oper.Status,
			"bytes_sent":       ret.DtControllerTunnelStatsOper.Oper.BytesSent,
			"bytes_recieved":   ret.DtControllerTunnelStatsOper.Oper.BytesRecieved,
			"number_of_errors": ret.DtControllerTunnelStatsOper.Oper.NumberOfErrors,
			"uptime":           ret.DtControllerTunnelStatsOper.Oper.Uptime,
			"error_message":    ret.DtControllerTunnelStatsOper.Oper.ErrorMessage,
		},
	}
}

func getObjectControllerTunnelStatsOperOper(d []interface{}) edpt.ControllerTunnelStatsOperOper {

	count1 := len(d)
	var ret edpt.ControllerTunnelStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(string)
		ret.BytesSent = in["bytes_sent"].(int)
		ret.BytesRecieved = in["bytes_recieved"].(int)
		ret.NumberOfErrors = in["number_of_errors"].(int)
		ret.Uptime = in["uptime"].(string)
		ret.ErrorMessage = in["error_message"].(string)
	}
	return ret
}

func dataToEndpointControllerTunnelStatsOper(d *schema.ResourceData) edpt.ControllerTunnelStatsOper {
	var ret edpt.ControllerTunnelStatsOper

	ret.Oper = getObjectControllerTunnelStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
