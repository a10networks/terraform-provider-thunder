package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiCertOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_pki_cert_oper`: Operational Status for the object cert\n\n__PLACEHOLDER__",
		ReadContext: resourcePkiCertOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sortby_name": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sortby_exp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"certs": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"serial": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"notbefore": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"notafter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"common_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"organization": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"subject": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"issuer": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"notafter_number": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"keysize": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"exact_match": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourcePkiCertOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCertOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCertOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		PkiCertOperOper := setObjectPkiCertOperOper(res)
		d.Set("oper", PkiCertOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPkiCertOperOper(ret edpt.DataPkiCertOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sortby_name": ret.DtPkiCertOper.Oper.SortbyName,
			"sortby_exp":  ret.DtPkiCertOper.Oper.SortbyExp,
			"certs":       setSlicePkiCertOperOperCerts(ret.DtPkiCertOper.Oper.Certs),
			"partition":   ret.DtPkiCertOper.Oper.Partition,
			"exact_match": ret.DtPkiCertOper.Oper.ExactMatch,
		},
	}
}

func setSlicePkiCertOperOperCerts(d []edpt.PkiCertOperOperCerts) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["type"] = item.Type
		in["serial"] = item.Serial
		in["notbefore"] = item.Notbefore
		in["notafter"] = item.Notafter
		in["common_name"] = item.CommonName
		in["organization"] = item.Organization
		in["subject"] = item.Subject
		in["issuer"] = item.Issuer
		in["notafter_number"] = item.NotafterNumber
		in["status"] = item.Status
		in["keysize"] = item.Keysize
		result = append(result, in)
	}
	return result
}

func getObjectPkiCertOperOper(d []interface{}) edpt.PkiCertOperOper {

	count1 := len(d)
	var ret edpt.PkiCertOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SortbyName = in["sortby_name"].(int)
		ret.SortbyExp = in["sortby_exp"].(int)
		ret.Certs = getSlicePkiCertOperOperCerts(in["certs"].([]interface{}))
		ret.Partition = in["partition"].(string)
		ret.ExactMatch = in["exact_match"].(int)
	}
	return ret
}

func getSlicePkiCertOperOperCerts(d []interface{}) []edpt.PkiCertOperOperCerts {

	count1 := len(d)
	ret := make([]edpt.PkiCertOperOperCerts, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PkiCertOperOperCerts
		oi.Name = in["name"].(string)
		oi.Type = in["type"].(string)
		oi.Serial = in["serial"].(string)
		oi.Notbefore = in["notbefore"].(string)
		oi.Notafter = in["notafter"].(string)
		oi.CommonName = in["common_name"].(string)
		oi.Organization = in["organization"].(string)
		oi.Subject = in["subject"].(string)
		oi.Issuer = in["issuer"].(string)
		oi.NotafterNumber = in["notafter_number"].(int)
		oi.Status = in["status"].(string)
		oi.Keysize = in["keysize"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointPkiCertOper(d *schema.ResourceData) edpt.PkiCertOper {
	var ret edpt.PkiCertOper

	ret.Oper = getObjectPkiCertOperOper(d.Get("oper").([]interface{}))
	return ret
}
