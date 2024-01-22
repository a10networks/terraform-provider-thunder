package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTcpRateLimitResetUnknownConnOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_tcp_rate_limit_reset_unknown_conn_oper`: Operational Status for the object rate-limit-reset-unknown-conn\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTcpRateLimitResetUnknownConnOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"unknown_conn_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"unknown_conn_current_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"unknown_conn_rate_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTcpRateLimitResetUnknownConnOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpRateLimitResetUnknownConnOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpRateLimitResetUnknownConnOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTcpRateLimitResetUnknownConnOperOper := setObjectSystemTcpRateLimitResetUnknownConnOperOper(res)
		d.Set("oper", SystemTcpRateLimitResetUnknownConnOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTcpRateLimitResetUnknownConnOperOper(ret edpt.DataSystemTcpRateLimitResetUnknownConnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"unknown_conn_rate_limit":      ret.DtSystemTcpRateLimitResetUnknownConnOper.Oper.UnknownConnRateLimit,
			"unknown_conn_current_rate":    ret.DtSystemTcpRateLimitResetUnknownConnOper.Oper.UnknownConnCurrentRate,
			"unknown_conn_rate_limit_drop": ret.DtSystemTcpRateLimitResetUnknownConnOper.Oper.UnknownConnRateLimitDrop,
		},
	}
}

func getObjectSystemTcpRateLimitResetUnknownConnOperOper(d []interface{}) edpt.SystemTcpRateLimitResetUnknownConnOperOper {

	count1 := len(d)
	var ret edpt.SystemTcpRateLimitResetUnknownConnOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UnknownConnRateLimit = in["unknown_conn_rate_limit"].(int)
		ret.UnknownConnCurrentRate = in["unknown_conn_current_rate"].(int)
		ret.UnknownConnRateLimitDrop = in["unknown_conn_rate_limit_drop"].(int)
	}
	return ret
}

func dataToEndpointSystemTcpRateLimitResetUnknownConnOper(d *schema.ResourceData) edpt.SystemTcpRateLimitResetUnknownConnOper {
	var ret edpt.SystemTcpRateLimitResetUnknownConnOper

	ret.Oper = getObjectSystemTcpRateLimitResetUnknownConnOperOper(d.Get("oper").([]interface{}))
	return ret
}
