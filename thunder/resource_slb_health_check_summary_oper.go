package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthCheckSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_check_summary_oper`: Operational Status for the object health-check-summary\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthCheckSummaryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_up": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_down": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_port_up": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_port_down": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthCheckSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthCheckSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthCheckSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthCheckSummaryOperOper := setObjectSlbHealthCheckSummaryOperOper(res)
		d.Set("oper", SlbHealthCheckSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthCheckSummaryOperOper(ret edpt.DataSlbHealthCheckSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"server_up":        ret.DtSlbHealthCheckSummaryOper.Oper.ServerUp,
			"server_down":      ret.DtSlbHealthCheckSummaryOper.Oper.ServerDown,
			"server_port_up":   ret.DtSlbHealthCheckSummaryOper.Oper.ServerPortUp,
			"server_port_down": ret.DtSlbHealthCheckSummaryOper.Oper.ServerPortDown,
		},
	}
}

func getObjectSlbHealthCheckSummaryOperOper(d []interface{}) edpt.SlbHealthCheckSummaryOperOper {

	count1 := len(d)
	var ret edpt.SlbHealthCheckSummaryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerUp = in["server_up"].(int)
		ret.ServerDown = in["server_down"].(int)
		ret.ServerPortUp = in["server_port_up"].(int)
		ret.ServerPortDown = in["server_port_down"].(int)
	}
	return ret
}

func dataToEndpointSlbHealthCheckSummaryOper(d *schema.ResourceData) edpt.SlbHealthCheckSummaryOper {
	var ret edpt.SlbHealthCheckSummaryOper

	ret.Oper = getObjectSlbHealthCheckSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
