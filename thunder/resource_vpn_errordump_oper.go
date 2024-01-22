package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnErrordumpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_errordump_oper`: Operational Status for the object errordump\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnErrordumpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipsec_error_dump_path": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVpnErrordumpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnErrordumpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnErrordumpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnErrordumpOperOper := setObjectVpnErrordumpOperOper(res)
		d.Set("oper", VpnErrordumpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnErrordumpOperOper(ret edpt.DataVpnErrordumpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipsec_error_dump_path": ret.DtVpnErrordumpOper.Oper.IpsecErrorDumpPath,
		},
	}
}

func getObjectVpnErrordumpOperOper(d []interface{}) edpt.VpnErrordumpOperOper {

	count1 := len(d)
	var ret edpt.VpnErrordumpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpsecErrorDumpPath = in["ipsec_error_dump_path"].(string)
	}
	return ret
}

func dataToEndpointVpnErrordumpOper(d *schema.ResourceData) edpt.VpnErrordumpOper {
	var ret edpt.VpnErrordumpOper

	ret.Oper = getObjectVpnErrordumpOperOper(d.Get("oper").([]interface{}))
	return ret
}
