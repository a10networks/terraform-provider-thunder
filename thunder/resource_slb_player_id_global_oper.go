package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPlayerIdGlobalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_player_id_global_oper`: Operational Status for the object player-id-global\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPlayerIdGlobalOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"time_to_active": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"table_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPlayerIdGlobalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdGlobalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdGlobalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPlayerIdGlobalOperOper := setObjectSlbPlayerIdGlobalOperOper(res)
		d.Set("oper", SlbPlayerIdGlobalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPlayerIdGlobalOperOper(ret edpt.DataSlbPlayerIdGlobalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":          ret.DtSlbPlayerIdGlobalOper.Oper.State,
			"time_to_active": ret.DtSlbPlayerIdGlobalOper.Oper.Time_to_active,
			"table_count":    ret.DtSlbPlayerIdGlobalOper.Oper.Table_count,
		},
	}
}

func getObjectSlbPlayerIdGlobalOperOper(d []interface{}) edpt.SlbPlayerIdGlobalOperOper {

	count1 := len(d)
	var ret edpt.SlbPlayerIdGlobalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Time_to_active = in["time_to_active"].(int)
		ret.Table_count = in["table_count"].(int)
	}
	return ret
}

func dataToEndpointSlbPlayerIdGlobalOper(d *schema.ResourceData) edpt.SlbPlayerIdGlobalOper {
	var ret edpt.SlbPlayerIdGlobalOper

	ret.Oper = getObjectSlbPlayerIdGlobalOperOper(d.Get("oper").([]interface{}))
	return ret
}
