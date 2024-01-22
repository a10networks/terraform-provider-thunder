package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCrlOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_crl_oper`: Operational Status for the object ssl-crl\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCrlOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vserver": {
							Type: schema.TypeString, Optional: true, Description: "virtual server name",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Port",
						},
						"crl_info": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"issuer": {
										Type: schema.TypeString, Optional: true, Description: "Issuer",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "Status",
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

func resourceSlbSslCrlOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCrlOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCrlOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCrlOperOper := setObjectSlbSslCrlOperOper(res)
		d.Set("oper", SlbSslCrlOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCrlOperOper(ret edpt.DataSlbSslCrlOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vserver":  ret.DtSlbSslCrlOper.Oper.Vserver,
			"port":     ret.DtSlbSslCrlOper.Oper.Port,
			"crl_info": setSliceSlbSslCrlOperOperCrlInfo(ret.DtSlbSslCrlOper.Oper.CrlInfo),
		},
	}
}

func setSliceSlbSslCrlOperOperCrlInfo(d []edpt.SlbSslCrlOperOperCrlInfo) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["issuer"] = item.Issuer
		in["status"] = item.Status
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslCrlOperOper(d []interface{}) edpt.SlbSslCrlOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCrlOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vserver = in["vserver"].(string)
		ret.Port = in["port"].(int)
		ret.CrlInfo = getSliceSlbSslCrlOperOperCrlInfo(in["crl_info"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslCrlOperOperCrlInfo(d []interface{}) []edpt.SlbSslCrlOperOperCrlInfo {

	count1 := len(d)
	ret := make([]edpt.SlbSslCrlOperOperCrlInfo, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslCrlOperOperCrlInfo
		oi.Issuer = in["issuer"].(string)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslCrlOper(d *schema.ResourceData) edpt.SlbSslCrlOper {
	var ret edpt.SlbSslCrlOper

	ret.Oper = getObjectSlbSslCrlOperOper(d.Get("oper").([]interface{}))
	return ret
}
