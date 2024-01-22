package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateClientSslOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_client_ssl_oper`: Operational Status for the object client-ssl\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateClientSslOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Client SSL Template Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_status_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cert_status_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cert_status_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cert_status_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cert_status_next_update": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cert_status_responder": {
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

func resourceSlbTemplateClientSslOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSslOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateClientSslOperOper := setObjectSlbTemplateClientSslOperOper(res)
		d.Set("oper", SlbTemplateClientSslOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateClientSslOperOper(ret edpt.DataSlbTemplateClientSslOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cert_status_list": setSliceSlbTemplateClientSslOperOperCertStatusList(ret.DtSlbTemplateClientSslOper.Oper.CertStatusList),
		},
	}
}

func setSliceSlbTemplateClientSslOperOperCertStatusList(d []edpt.SlbTemplateClientSslOperOperCertStatusList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cert_status_name"] = item.CertStatusName
		in["cert_status_status"] = item.CertStatusStatus
		in["cert_status_age"] = item.CertStatusAge
		in["cert_status_next_update"] = item.CertStatusNextUpdate
		in["cert_status_responder"] = item.CertStatusResponder
		result = append(result, in)
	}
	return result
}

func getObjectSlbTemplateClientSslOperOper(d []interface{}) edpt.SlbTemplateClientSslOperOper {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSslOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CertStatusList = getSliceSlbTemplateClientSslOperOperCertStatusList(in["cert_status_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateClientSslOperOperCertStatusList(d []interface{}) []edpt.SlbTemplateClientSslOperOperCertStatusList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateClientSslOperOperCertStatusList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateClientSslOperOperCertStatusList
		oi.CertStatusName = in["cert_status_name"].(string)
		oi.CertStatusStatus = in["cert_status_status"].(string)
		oi.CertStatusAge = in["cert_status_age"].(int)
		oi.CertStatusNextUpdate = in["cert_status_next_update"].(string)
		oi.CertStatusResponder = in["cert_status_responder"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateClientSslOper(d *schema.ResourceData) edpt.SlbTemplateClientSslOper {
	var ret edpt.SlbTemplateClientSslOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectSlbTemplateClientSslOperOper(d.Get("oper").([]interface{}))
	return ret
}
