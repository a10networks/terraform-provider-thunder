package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslScepCertStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_scep_cert_status_oper`: Operational Status for the object ssl-scep-cert-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslScepCertStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scep_certs": {
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

func resourceSlbSslScepCertStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslScepCertStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslScepCertStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslScepCertStatusOperOper := setObjectSlbSslScepCertStatusOperOper(res)
		d.Set("oper", SlbSslScepCertStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslScepCertStatusOperOper(ret edpt.DataSlbSslScepCertStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"scep_certs": setSliceSlbSslScepCertStatusOperOperScepCerts(ret.DtSlbSslScepCertStatusOper.Oper.ScepCerts),
		},
	}
}

func setSliceSlbSslScepCertStatusOperOperScepCerts(d []edpt.SlbSslScepCertStatusOperOperScepCerts) []map[string]interface{} {
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

func getObjectSlbSslScepCertStatusOperOper(d []interface{}) edpt.SlbSslScepCertStatusOperOper {

	count1 := len(d)
	var ret edpt.SlbSslScepCertStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ScepCerts = getSliceSlbSslScepCertStatusOperOperScepCerts(in["scep_certs"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslScepCertStatusOperOperScepCerts(d []interface{}) []edpt.SlbSslScepCertStatusOperOperScepCerts {

	count1 := len(d)
	ret := make([]edpt.SlbSslScepCertStatusOperOperScepCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslScepCertStatusOperOperScepCerts
		oi.Name = in["name"].(string)
		oi.Status = in["status"].(string)
		oi.Renew = in["renew"].(string)
		oi.Rotated = in["rotated"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslScepCertStatusOper(d *schema.ResourceData) edpt.SlbSslScepCertStatusOper {
	var ret edpt.SlbSslScepCertStatusOper

	ret.Oper = getObjectSlbSslScepCertStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
