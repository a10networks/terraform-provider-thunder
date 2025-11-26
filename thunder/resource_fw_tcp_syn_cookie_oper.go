package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpSynCookieOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_tcp_syn_cookie_oper`: Operational Status for the object syn-cookie\n\n__PLACEHOLDER__",
		ReadContext: resourceFwTcpSynCookieOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"syn_cookie_on": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwTcpSynCookieOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpSynCookieOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpSynCookieOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwTcpSynCookieOperOper := setObjectFwTcpSynCookieOperOper(res)
		d.Set("oper", FwTcpSynCookieOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwTcpSynCookieOperOper(ret edpt.DataFwTcpSynCookieOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"syn_cookie_on": ret.DtFwTcpSynCookieOper.Oper.Syn_cookie_on,
		},
	}
}

func getObjectFwTcpSynCookieOperOper(d []interface{}) edpt.FwTcpSynCookieOperOper {

	count1 := len(d)
	var ret edpt.FwTcpSynCookieOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn_cookie_on = in["syn_cookie_on"].(int)
	}
	return ret
}

func dataToEndpointFwTcpSynCookieOper(d *schema.ResourceData) edpt.FwTcpSynCookieOper {
	var ret edpt.FwTcpSynCookieOper

	ret.Oper = getObjectFwTcpSynCookieOperOper(d.Get("oper").([]interface{}))
	return ret
}
