package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationPasswordRetryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_password_retry_oper`: Operational Status for the object password-retry\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationPasswordRetryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"logon": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pw_failure": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"locked_out": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"logon_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationPasswordRetryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationPasswordRetryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationPasswordRetryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationPasswordRetryOperOper := setObjectAamAuthenticationPasswordRetryOperOper(res)
		d.Set("oper", AamAuthenticationPasswordRetryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationPasswordRetryOperOper(ret edpt.DataAamAuthenticationPasswordRetryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list": setSliceAamAuthenticationPasswordRetryOperOperEntryList(ret.DtAamAuthenticationPasswordRetryOper.Oper.EntryList),
			"logon_name": ret.DtAamAuthenticationPasswordRetryOper.Oper.LogonName,
			"name":       ret.DtAamAuthenticationPasswordRetryOper.Oper.Name,
		},
	}
}

func setSliceAamAuthenticationPasswordRetryOperOperEntryList(d []edpt.AamAuthenticationPasswordRetryOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["account"] = item.Account
		in["logon"] = item.Logon
		in["pw_failure"] = item.PwFailure
		in["ttl"] = item.Ttl
		in["locked_out"] = item.LockedOut
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationPasswordRetryOperOper(d []interface{}) edpt.AamAuthenticationPasswordRetryOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationPasswordRetryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceAamAuthenticationPasswordRetryOperOperEntryList(in["entry_list"].([]interface{}))
		ret.LogonName = in["logon_name"].(string)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getSliceAamAuthenticationPasswordRetryOperOperEntryList(d []interface{}) []edpt.AamAuthenticationPasswordRetryOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationPasswordRetryOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationPasswordRetryOperOperEntryList
		oi.Account = in["account"].(string)
		oi.Logon = in["logon"].(string)
		oi.PwFailure = in["pw_failure"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.LockedOut = in["locked_out"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationPasswordRetryOper(d *schema.ResourceData) edpt.AamAuthenticationPasswordRetryOper {
	var ret edpt.AamAuthenticationPasswordRetryOper

	ret.Oper = getObjectAamAuthenticationPasswordRetryOperOper(d.Get("oper").([]interface{}))
	return ret
}
