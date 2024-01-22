package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetMgmtSnmpEngineidOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_net_mgmt_snmp_engineID_oper`: Operational Status for the object engineID\n\n__PLACEHOLDER__",
		ReadContext: resourceNetMgmtSnmpEngineidOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engineid": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceNetMgmtSnmpEngineidOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetMgmtSnmpEngineidOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetMgmtSnmpEngineidOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NetMgmtSnmpEngineidOperOper := setObjectNetMgmtSnmpEngineidOperOper(res)
		d.Set("oper", NetMgmtSnmpEngineidOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNetMgmtSnmpEngineidOperOper(ret edpt.DataNetMgmtSnmpEngineidOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"engineid": ret.DtNetMgmtSnmpEngineidOper.Oper.Engineid,
		},
	}
}

func getObjectNetMgmtSnmpEngineidOperOper(d []interface{}) edpt.NetMgmtSnmpEngineidOperOper {

	count1 := len(d)
	var ret edpt.NetMgmtSnmpEngineidOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Engineid = in["engineid"].(string)
	}
	return ret
}

func dataToEndpointNetMgmtSnmpEngineidOper(d *schema.ResourceData) edpt.NetMgmtSnmpEngineidOper {
	var ret edpt.NetMgmtSnmpEngineidOper

	ret.Oper = getObjectNetMgmtSnmpEngineidOperOper(d.Get("oper").([]interface{}))
	return ret
}
