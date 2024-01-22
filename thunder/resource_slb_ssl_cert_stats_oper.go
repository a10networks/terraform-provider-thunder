package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCertStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_cert_stats_oper`: Operational Status for the object ssl-cert-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCertStatsOperRead,

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

func resourceSlbSslCertStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCertStatsOperOper := setObjectSlbSslCertStatsOperOper(res)
		d.Set("oper", SlbSslCertStatsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCertStatsOperOper(ret edpt.DataSlbSslCertStatsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cert_count":    ret.DtSlbSslCertStatsOper.Oper.CertCount,
			"key_count":     ret.DtSlbSslCertStatsOper.Oper.KeyCount,
			"ca_cert_count": ret.DtSlbSslCertStatsOper.Oper.CaCertCount,
			"partition":     ret.DtSlbSslCertStatsOper.Oper.Partition,
		},
	}
}

func getObjectSlbSslCertStatsOperOper(d []interface{}) edpt.SlbSslCertStatsOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCertStatsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CertCount = in["cert_count"].(int)
		ret.KeyCount = in["key_count"].(int)
		ret.CaCertCount = in["ca_cert_count"].(int)
		ret.Partition = in["partition"].(string)
	}
	return ret
}

func dataToEndpointSlbSslCertStatsOper(d *schema.ResourceData) edpt.SlbSslCertStatsOper {
	var ret edpt.SlbSslCertStatsOper

	ret.Oper = getObjectSlbSslCertStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
