package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnCrlOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_crl_oper`: Operational Status for the object crl\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnCrlOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crl_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subject": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"issuer": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"updates": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"serial": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"revoked": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"storage_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total_crls": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVpnCrlOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnCrlOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnCrlOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnCrlOperOper := setObjectVpnCrlOperOper(res)
		d.Set("oper", VpnCrlOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnCrlOperOper(ret edpt.DataVpnCrlOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"crl_list":   setSliceVpnCrlOperOperCrlList(ret.DtVpnCrlOper.Oper.CrlList),
			"total_crls": ret.DtVpnCrlOper.Oper.TotalCrls,
		},
	}
}

func setSliceVpnCrlOperOperCrlList(d []edpt.VpnCrlOperOperCrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["subject"] = item.Subject
		in["issuer"] = item.Issuer
		in["updates"] = item.Updates
		in["serial"] = item.Serial
		in["revoked"] = item.Revoked
		in["storage_type"] = item.StorageType
		result = append(result, in)
	}
	return result
}

func getObjectVpnCrlOperOper(d []interface{}) edpt.VpnCrlOperOper {

	count1 := len(d)
	var ret edpt.VpnCrlOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CrlList = getSliceVpnCrlOperOperCrlList(in["crl_list"].([]interface{}))
		ret.TotalCrls = in["total_crls"].(int)
	}
	return ret
}

func getSliceVpnCrlOperOperCrlList(d []interface{}) []edpt.VpnCrlOperOperCrlList {

	count1 := len(d)
	ret := make([]edpt.VpnCrlOperOperCrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnCrlOperOperCrlList
		oi.Subject = in["subject"].(string)
		oi.Issuer = in["issuer"].(string)
		oi.Updates = in["updates"].(string)
		oi.Serial = in["serial"].(string)
		oi.Revoked = in["revoked"].(string)
		oi.StorageType = in["storage_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnCrlOper(d *schema.ResourceData) edpt.VpnCrlOper {
	var ret edpt.VpnCrlOper

	ret.Oper = getObjectVpnCrlOperOper(d.Get("oper").([]interface{}))
	return ret
}
