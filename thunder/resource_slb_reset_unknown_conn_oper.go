package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbResetUnknownConnOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_reset_unknown_conn_oper`: Operational Status for the object reset-unknown-conn\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbResetUnknownConnOperRead,

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

func resourceSlbResetUnknownConnOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbResetUnknownConnOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbResetUnknownConnOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbResetUnknownConnOperOper := setObjectSlbResetUnknownConnOperOper(res)
		d.Set("oper", SlbResetUnknownConnOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbResetUnknownConnOperOper(ret edpt.DataSlbResetUnknownConnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"unknown_conn_rate_limit":      ret.DtSlbResetUnknownConnOper.Oper.UnknownConnRateLimit,
			"unknown_conn_current_rate":    ret.DtSlbResetUnknownConnOper.Oper.UnknownConnCurrentRate,
			"unknown_conn_rate_limit_drop": ret.DtSlbResetUnknownConnOper.Oper.UnknownConnRateLimitDrop,
		},
	}
}

func getObjectSlbResetUnknownConnOperOper(d []interface{}) edpt.SlbResetUnknownConnOperOper {

	count1 := len(d)
	var ret edpt.SlbResetUnknownConnOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UnknownConnRateLimit = in["unknown_conn_rate_limit"].(int)
		ret.UnknownConnCurrentRate = in["unknown_conn_current_rate"].(int)
		ret.UnknownConnRateLimitDrop = in["unknown_conn_rate_limit_drop"].(int)
	}
	return ret
}

func dataToEndpointSlbResetUnknownConnOper(d *schema.ResourceData) edpt.SlbResetUnknownConnOper {
	var ret edpt.SlbResetUnknownConnOper

	ret.Oper = getObjectSlbResetUnknownConnOperOper(d.Get("oper").([]interface{}))
	return ret
}
