package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiCertStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_pki_cert_stats_oper`: Operational Status for the object cert-stats\n\n__PLACEHOLDER__",
		ReadContext: resourcePkiCertStatsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"key_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ca_cert_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourcePkiCertStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCertStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCertStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		PkiCertStatsOperOper := setObjectPkiCertStatsOperOper(res)
		d.Set("oper", PkiCertStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPkiCertStatsOperOper(ret edpt.DataPkiCertStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cert_count":    ret.DtPkiCertStatsOper.Oper.CertCount,
			"key_count":     ret.DtPkiCertStatsOper.Oper.KeyCount,
			"ca_cert_count": ret.DtPkiCertStatsOper.Oper.CaCertCount,
			"partition":     ret.DtPkiCertStatsOper.Oper.Partition,
		},
	}
}

func getObjectPkiCertStatsOperOper(d []interface{}) edpt.PkiCertStatsOperOper {

	count1 := len(d)
	var ret edpt.PkiCertStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CertCount = in["cert_count"].(int)
		ret.KeyCount = in["key_count"].(int)
		ret.CaCertCount = in["ca_cert_count"].(int)
		ret.Partition = in["partition"].(string)
	}
	return ret
}

func dataToEndpointPkiCertStatsOper(d *schema.ResourceData) edpt.PkiCertStatsOper {
	var ret edpt.PkiCertStatsOper

	ret.Oper = getObjectPkiCertStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
