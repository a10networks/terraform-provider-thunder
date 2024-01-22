package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCertExpireCheckOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_cert_expire_check_oper`: Operational Status for the object ssl-cert-expire-check\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCertExpireCheckOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expire_check_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"email_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"email_address2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"before": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ssl_exception": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exception_cert": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslCertExpireCheckOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertExpireCheckOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertExpireCheckOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCertExpireCheckOperOper := setObjectSlbSslCertExpireCheckOperOper(res)
		d.Set("oper", SlbSslCertExpireCheckOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCertExpireCheckOperOper(ret edpt.DataSlbSslCertExpireCheckOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"expire_check_status": ret.DtSlbSslCertExpireCheckOper.Oper.ExpireCheckStatus,
			"email_address":       ret.DtSlbSslCertExpireCheckOper.Oper.EmailAddress,
			"email_address2":      ret.DtSlbSslCertExpireCheckOper.Oper.EmailAddress2,
			"before":              ret.DtSlbSslCertExpireCheckOper.Oper.Before,
			"interval":            ret.DtSlbSslCertExpireCheckOper.Oper.Interval,
			"ssl_exception":       setSliceSlbSslCertExpireCheckOperOperSslException(ret.DtSlbSslCertExpireCheckOper.Oper.SslException),
		},
	}
}

func setSliceSlbSslCertExpireCheckOperOperSslException(d []edpt.SlbSslCertExpireCheckOperOperSslException) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["exception_cert"] = item.ExceptionCert
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslCertExpireCheckOperOper(d []interface{}) edpt.SlbSslCertExpireCheckOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCertExpireCheckOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExpireCheckStatus = in["expire_check_status"].(string)
		ret.EmailAddress = in["email_address"].(string)
		ret.EmailAddress2 = in["email_address2"].(string)
		ret.Before = in["before"].(int)
		ret.Interval = in["interval"].(int)
		ret.SslException = getSliceSlbSslCertExpireCheckOperOperSslException(in["ssl_exception"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslCertExpireCheckOperOperSslException(d []interface{}) []edpt.SlbSslCertExpireCheckOperOperSslException {

	count1 := len(d)
	ret := make([]edpt.SlbSslCertExpireCheckOperOperSslException, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslCertExpireCheckOperOperSslException
		oi.ExceptionCert = in["exception_cert"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslCertExpireCheckOper(d *schema.ResourceData) edpt.SlbSslCertExpireCheckOper {
	var ret edpt.SlbSslCertExpireCheckOper

	ret.Oper = getObjectSlbSslCertExpireCheckOperOper(d.Get("oper").([]interface{}))
	return ret
}
