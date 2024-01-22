package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAllSessionsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_all_sessions_oper`: Operational Status for the object all-sessions\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAllSessionsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_pool_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnAllSessionsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAllSessionsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAllSessionsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAllSessionsOperOper := setObjectCgnv6LsnAllSessionsOperOper(res)
		d.Set("oper", Cgnv6LsnAllSessionsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAllSessionsOperOper(ret edpt.DataCgnv6LsnAllSessionsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status":        ret.DtCgnv6LsnAllSessionsOper.Oper.Status,
			"nat_pool_name": ret.DtCgnv6LsnAllSessionsOper.Oper.NatPoolName,
		},
	}
}

func getObjectCgnv6LsnAllSessionsOperOper(d []interface{}) edpt.Cgnv6LsnAllSessionsOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAllSessionsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(string)
		ret.NatPoolName = in["nat_pool_name"].(string)
	}
	return ret
}

func dataToEndpointCgnv6LsnAllSessionsOper(d *schema.ResourceData) edpt.Cgnv6LsnAllSessionsOper {
	var ret edpt.Cgnv6LsnAllSessionsOper

	ret.Oper = getObjectCgnv6LsnAllSessionsOperOper(d.Get("oper").([]interface{}))
	return ret
}
