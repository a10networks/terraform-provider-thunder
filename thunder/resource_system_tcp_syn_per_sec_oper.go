package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTcpSynPerSecOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_tcp_syn_per_sec_oper`: Operational Status for the object tcp-syn-per-sec\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTcpSynPerSecOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_syn_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemTcpSynPerSecOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpSynPerSecOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpSynPerSecOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTcpSynPerSecOperOper := setObjectSystemTcpSynPerSecOperOper(res)
		d.Set("oper", SystemTcpSynPerSecOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTcpSynPerSecOperOper(ret edpt.DataSystemTcpSynPerSecOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_syn_per_sec": ret.DtSystemTcpSynPerSecOper.Oper.TcpSynPerSec,
		},
	}
}

func getObjectSystemTcpSynPerSecOperOper(d []interface{}) edpt.SystemTcpSynPerSecOperOper {

	count1 := len(d)
	var ret edpt.SystemTcpSynPerSecOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpSynPerSec = in["tcp_syn_per_sec"].(int)
	}
	return ret
}

func dataToEndpointSystemTcpSynPerSecOper(d *schema.ResourceData) edpt.SystemTcpSynPerSecOper {
	var ret edpt.SystemTcpSynPerSecOper

	ret.Oper = getObjectSystemTcpSynPerSecOperOper(d.Get("oper").([]interface{}))
	return ret
}
