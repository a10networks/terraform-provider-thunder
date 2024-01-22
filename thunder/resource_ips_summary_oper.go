package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpsSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ips_summary_oper`: Operational Status for the object summary\n\n__PLACEHOLDER__",
		ReadContext: resourceIpsSummaryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"a10_signature_source": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"a10_signature_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"custom_signature_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"protection": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceIpsSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpsSummaryOperOper := setObjectIpsSummaryOperOper(res)
		d.Set("oper", IpsSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpsSummaryOperOper(ret edpt.DataIpsSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"version":                ret.DtIpsSummaryOper.Oper.Version,
			"a10_signature_source":   ret.DtIpsSummaryOper.Oper.A10SignatureSource,
			"a10_signature_count":    ret.DtIpsSummaryOper.Oper.A10SignatureCount,
			"custom_signature_count": ret.DtIpsSummaryOper.Oper.CustomSignatureCount,
			"protection":             ret.DtIpsSummaryOper.Oper.Protection,
		},
	}
}

func getObjectIpsSummaryOperOper(d []interface{}) edpt.IpsSummaryOperOper {

	count1 := len(d)
	var ret edpt.IpsSummaryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Version = in["version"].(string)
		ret.A10SignatureSource = in["a10_signature_source"].(string)
		ret.A10SignatureCount = in["a10_signature_count"].(int)
		ret.CustomSignatureCount = in["custom_signature_count"].(int)
		ret.Protection = in["protection"].(string)
	}
	return ret
}

func dataToEndpointIpsSummaryOper(d *schema.ResourceData) edpt.IpsSummaryOper {
	var ret edpt.IpsSummaryOper

	ret.Oper = getObjectIpsSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
