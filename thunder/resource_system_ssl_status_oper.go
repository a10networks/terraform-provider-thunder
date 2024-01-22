package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSslStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ssl_status_oper`: Operational Status for the object ssl-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemSslStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ssl_setup_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"num_chip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_aes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"crypto_support": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemSslStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemSslStatusOperOper := setObjectSystemSslStatusOperOper(res)
		d.Set("oper", SystemSslStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemSslStatusOperOper(ret edpt.DataSystemSslStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enable":           ret.DtSystemSslStatusOper.Oper.Enable,
			"ssl_setup_status": ret.DtSystemSslStatusOper.Oper.SslSetupStatus,
			"num_chip":         ret.DtSystemSslStatusOper.Oper.Num_chip,
			"num_aes":          ret.DtSystemSslStatusOper.Oper.Num_aes,
			"crypto_support":   ret.DtSystemSslStatusOper.Oper.CryptoSupport,
		},
	}
}

func getObjectSystemSslStatusOperOper(d []interface{}) edpt.SystemSslStatusOperOper {

	count1 := len(d)
	var ret edpt.SystemSslStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(string)
		ret.SslSetupStatus = in["ssl_setup_status"].(string)
		ret.Num_chip = in["num_chip"].(int)
		ret.Num_aes = in["num_aes"].(int)
		ret.CryptoSupport = in["crypto_support"].(string)
	}
	return ret
}

func dataToEndpointSystemSslStatusOper(d *schema.ResourceData) edpt.SystemSslStatusOper {
	var ret edpt.SystemSslStatusOper

	ret.Oper = getObjectSystemSslStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
