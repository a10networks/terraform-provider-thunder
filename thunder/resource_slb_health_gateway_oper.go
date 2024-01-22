package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthGatewayOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_gateway_oper`: Operational Status for the object health-gateway\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthGatewayOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthGatewayOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthGatewayOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthGatewayOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthGatewayOperOper := setObjectSlbHealthGatewayOperOper(res)
		d.Set("oper", SlbHealthGatewayOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthGatewayOperOper(ret edpt.DataSlbHealthGatewayOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enabled":  ret.DtSlbHealthGatewayOper.Oper.Enabled,
			"interval": ret.DtSlbHealthGatewayOper.Oper.Interval,
			"timeout":  ret.DtSlbHealthGatewayOper.Oper.Timeout,
		},
	}
}

func getObjectSlbHealthGatewayOperOper(d []interface{}) edpt.SlbHealthGatewayOperOper {

	count1 := len(d)
	var ret edpt.SlbHealthGatewayOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enabled = in["enabled"].(int)
		ret.Interval = in["interval"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointSlbHealthGatewayOper(d *schema.ResourceData) edpt.SlbHealthGatewayOper {
	var ret edpt.SlbHealthGatewayOper

	ret.Oper = getObjectSlbHealthGatewayOperOper(d.Get("oper").([]interface{}))
	return ret
}
