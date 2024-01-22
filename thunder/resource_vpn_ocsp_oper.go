package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnOcspOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ocsp_oper`: Operational Status for the object ocsp\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnOcspOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ocsp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subject": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"issuer": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"validity": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"certificate_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total_ocsps": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVpnOcspOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnOcspOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnOcspOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnOcspOperOper := setObjectVpnOcspOperOper(res)
		d.Set("oper", VpnOcspOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnOcspOperOper(ret edpt.DataVpnOcspOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ocsp_list":   setSliceVpnOcspOperOperOcspList(ret.DtVpnOcspOper.Oper.OcspList),
			"total_ocsps": ret.DtVpnOcspOper.Oper.TotalOcsps,
		},
	}
}

func setSliceVpnOcspOperOperOcspList(d []edpt.VpnOcspOperOperOcspList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["subject"] = item.Subject
		in["issuer"] = item.Issuer
		in["validity"] = item.Validity
		in["certificate_status"] = item.CertificateStatus
		result = append(result, in)
	}
	return result
}

func getObjectVpnOcspOperOper(d []interface{}) edpt.VpnOcspOperOper {

	count1 := len(d)
	var ret edpt.VpnOcspOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OcspList = getSliceVpnOcspOperOperOcspList(in["ocsp_list"].([]interface{}))
		ret.TotalOcsps = in["total_ocsps"].(int)
	}
	return ret
}

func getSliceVpnOcspOperOperOcspList(d []interface{}) []edpt.VpnOcspOperOperOcspList {

	count1 := len(d)
	ret := make([]edpt.VpnOcspOperOperOcspList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOcspOperOperOcspList
		oi.Subject = in["subject"].(string)
		oi.Issuer = in["issuer"].(string)
		oi.Validity = in["validity"].(string)
		oi.CertificateStatus = in["certificate_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnOcspOper(d *schema.ResourceData) edpt.VpnOcspOper {
	var ret edpt.VpnOcspOper

	ret.Oper = getObjectVpnOcspOperOper(d.Get("oper").([]interface{}))
	return ret
}
