package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTimezoneOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_timezone_oper`: Operational Status for the object timezone\n\n__PLACEHOLDER__",
		ReadContext: resourceTimezoneOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"std_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"deny_dst": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"location": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceTimezoneOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTimezoneOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTimezoneOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TimezoneOperOper := setObjectTimezoneOperOper(res)
		d.Set("oper", TimezoneOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectTimezoneOperOper(ret edpt.DataTimezoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"std_name": ret.DtTimezoneOper.Oper.StdName,
			"dst_name": ret.DtTimezoneOper.Oper.DstName,
			"deny_dst": ret.DtTimezoneOper.Oper.DenyDst,
			"location": ret.DtTimezoneOper.Oper.Location,
		},
	}
}

func getObjectTimezoneOperOper(d []interface{}) edpt.TimezoneOperOper {

	count1 := len(d)
	var ret edpt.TimezoneOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StdName = in["std_name"].(string)
		ret.DstName = in["dst_name"].(string)
		ret.DenyDst = in["deny_dst"].(string)
		ret.Location = in["location"].(string)
	}
	return ret
}

func dataToEndpointTimezoneOper(d *schema.ResourceData) edpt.TimezoneOper {
	var ret edpt.TimezoneOper

	ret.Oper = getObjectTimezoneOperOper(d.Get("oper").([]interface{}))
	return ret
}
