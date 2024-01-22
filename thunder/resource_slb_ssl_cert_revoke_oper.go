package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCertRevokeOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_cert_revoke_oper`: Operational Status for the object ssl-cert-revoke\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCertRevokeOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vserver": {
							Type: schema.TypeString, Optional: true, Description: "DELETE method filter: Virtual Server Name",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "DELETE method filter: Virtual Port",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslCertRevokeOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertRevokeOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertRevokeOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCertRevokeOperOper := setObjectSlbSslCertRevokeOperOper(res)
		d.Set("oper", SlbSslCertRevokeOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCertRevokeOperOper(ret edpt.DataSlbSslCertRevokeOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vserver": ret.DtSlbSslCertRevokeOper.Oper.Vserver,
			"port":    ret.DtSlbSslCertRevokeOper.Oper.Port,
		},
	}
}

func getObjectSlbSslCertRevokeOperOper(d []interface{}) edpt.SlbSslCertRevokeOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCertRevokeOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vserver = in["vserver"].(string)
		ret.Port = in["port"].(int)
	}
	return ret
}

func dataToEndpointSlbSslCertRevokeOper(d *schema.ResourceData) edpt.SlbSslCertRevokeOper {
	var ret edpt.SlbSslCertRevokeOper

	ret.Oper = getObjectSlbSslCertRevokeOperOper(d.Get("oper").([]interface{}))
	return ret
}
