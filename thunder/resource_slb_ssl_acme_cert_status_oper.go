package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslAcmeCertStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_acme_cert_status_oper`: Operational Status for the object ssl-acme-cert-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslAcmeCertStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acme_certs": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"renew": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rotated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_enroll_renew_status": {
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

func resourceSlbSslAcmeCertStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslAcmeCertStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslAcmeCertStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslAcmeCertStatusOperOper := setObjectSlbSslAcmeCertStatusOperOper(res)
		d.Set("oper", SlbSslAcmeCertStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslAcmeCertStatusOperOper(ret edpt.DataSlbSslAcmeCertStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"acme_certs": setSliceSlbSslAcmeCertStatusOperOperAcmeCerts(ret.DtSlbSslAcmeCertStatusOper.Oper.AcmeCerts),
		},
	}
}

func setSliceSlbSslAcmeCertStatusOperOperAcmeCerts(d []edpt.SlbSslAcmeCertStatusOperOperAcmeCerts) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["status"] = item.Status
		in["renew"] = item.Renew
		in["rotated"] = item.Rotated
		in["last_enroll_renew_status"] = item.LastEnrollRenewStatus
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslAcmeCertStatusOperOper(d []interface{}) edpt.SlbSslAcmeCertStatusOperOper {

	count1 := len(d)
	var ret edpt.SlbSslAcmeCertStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AcmeCerts = getSliceSlbSslAcmeCertStatusOperOperAcmeCerts(in["acme_certs"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslAcmeCertStatusOperOperAcmeCerts(d []interface{}) []edpt.SlbSslAcmeCertStatusOperOperAcmeCerts {

	count1 := len(d)
	ret := make([]edpt.SlbSslAcmeCertStatusOperOperAcmeCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslAcmeCertStatusOperOperAcmeCerts
		oi.Name = in["name"].(string)
		oi.Status = in["status"].(string)
		oi.Renew = in["renew"].(string)
		oi.Rotated = in["rotated"].(int)
		oi.LastEnrollRenewStatus = in["last_enroll_renew_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslAcmeCertStatusOper(d *schema.ResourceData) edpt.SlbSslAcmeCertStatusOper {
	var ret edpt.SlbSslAcmeCertStatusOper

	ret.Oper = getObjectSlbSslAcmeCertStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
