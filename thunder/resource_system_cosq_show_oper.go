package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCosqShowOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cosq_show_oper`: Operational Status for the object cosq-show\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCosqShowOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkt_percosq": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemCosqShowOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCosqShowOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCosqShowOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCosqShowOperOper := setObjectSystemCosqShowOperOper(res)
		d.Set("oper", SystemCosqShowOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCosqShowOperOper(ret edpt.DataSystemCosqShowOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pkt_percosq": ret.DtSystemCosqShowOper.Oper.PktPercosq,
		},
	}
}

func getObjectSystemCosqShowOperOper(d []interface{}) edpt.SystemCosqShowOperOper {

	count1 := len(d)
	var ret edpt.SystemCosqShowOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktPercosq = in["pkt_percosq"].(int)
	}
	return ret
}

func dataToEndpointSystemCosqShowOper(d *schema.ResourceData) edpt.SystemCosqShowOper {
	var ret edpt.SystemCosqShowOper

	ret.Oper = getObjectSystemCosqShowOperOper(d.Get("oper").([]interface{}))
	return ret
}
