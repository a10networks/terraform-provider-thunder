package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTokenAuthenticationSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_token_authentication_summary_oper`: Operational Status for the object summary\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosTokenAuthenticationSummaryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"player_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosTokenAuthenticationSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosTokenAuthenticationSummaryOperOper := setObjectDdosTokenAuthenticationSummaryOperOper(res)
		d.Set("oper", DdosTokenAuthenticationSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosTokenAuthenticationSummaryOperOper(ret edpt.DataDdosTokenAuthenticationSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"player_mode": ret.DtDdosTokenAuthenticationSummaryOper.Oper.PlayerMode,
		},
	}
}

func getObjectDdosTokenAuthenticationSummaryOperOper(d []interface{}) edpt.DdosTokenAuthenticationSummaryOperOper {

	count1 := len(d)
	var ret edpt.DdosTokenAuthenticationSummaryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PlayerMode = in["player_mode"].(string)
	}
	return ret
}

func dataToEndpointDdosTokenAuthenticationSummaryOper(d *schema.ResourceData) edpt.DdosTokenAuthenticationSummaryOper {
	var ret edpt.DdosTokenAuthenticationSummaryOper

	ret.Oper = getObjectDdosTokenAuthenticationSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
