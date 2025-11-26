package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHrxqStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_hrxq_status_oper`: Operational Status for the object hrxq-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemHrxqStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deep_hrxq_enable": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hrxq_num_chunks": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemHrxqStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHrxqStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHrxqStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemHrxqStatusOperOper := setObjectSystemHrxqStatusOperOper(res)
		d.Set("oper", SystemHrxqStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemHrxqStatusOperOper(ret edpt.DataSystemHrxqStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"deep_hrxq_enable": ret.DtSystemHrxqStatusOper.Oper.DeepHrxqEnable,
			"hrxq_num_chunks":  ret.DtSystemHrxqStatusOper.Oper.Hrxq_num_chunks,
		},
	}
}

func getObjectSystemHrxqStatusOperOper(d []interface{}) edpt.SystemHrxqStatusOperOper {

	count1 := len(d)
	var ret edpt.SystemHrxqStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DeepHrxqEnable = in["deep_hrxq_enable"].(string)
		ret.Hrxq_num_chunks = in["hrxq_num_chunks"].(int)
	}
	return ret
}

func dataToEndpointSystemHrxqStatusOper(d *schema.ResourceData) edpt.SystemHrxqStatusOper {
	var ret edpt.SystemHrxqStatusOper

	ret.Oper = getObjectSystemHrxqStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
