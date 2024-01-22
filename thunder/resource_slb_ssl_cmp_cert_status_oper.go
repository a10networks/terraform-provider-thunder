package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCmpCertStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_cmp_cert_status_oper`: Operational Status for the object ssl-cmp-cert-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCmpCertStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cmp_certs": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslCmpCertStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCmpCertStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCmpCertStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCmpCertStatusOperOper := setObjectSlbSslCmpCertStatusOperOper(res)
		d.Set("oper", SlbSslCmpCertStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCmpCertStatusOperOper(ret edpt.DataSlbSslCmpCertStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cmp_certs": setSliceSlbSslCmpCertStatusOperOperCmpCerts(ret.DtSlbSslCmpCertStatusOper.Oper.CmpCerts),
		},
	}
}

func setSliceSlbSslCmpCertStatusOperOperCmpCerts(d []edpt.SlbSslCmpCertStatusOperOperCmpCerts) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["status"] = item.Status
		in["renew"] = item.Renew
		in["rotated"] = item.Rotated
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslCmpCertStatusOperOper(d []interface{}) edpt.SlbSslCmpCertStatusOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCmpCertStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CmpCerts = getSliceSlbSslCmpCertStatusOperOperCmpCerts(in["cmp_certs"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslCmpCertStatusOperOperCmpCerts(d []interface{}) []edpt.SlbSslCmpCertStatusOperOperCmpCerts {

	count1 := len(d)
	ret := make([]edpt.SlbSslCmpCertStatusOperOperCmpCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslCmpCertStatusOperOperCmpCerts
		oi.Name = in["name"].(string)
		oi.Status = in["status"].(string)
		oi.Renew = in["renew"].(string)
		oi.Rotated = in["rotated"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslCmpCertStatusOper(d *schema.ResourceData) edpt.SlbSslCmpCertStatusOper {
	var ret edpt.SlbSslCmpCertStatusOper

	ret.Oper = getObjectSlbSslCmpCertStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
