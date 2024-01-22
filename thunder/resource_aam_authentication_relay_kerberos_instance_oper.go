package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayKerberosInstanceOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_kerberos_instance_oper`: Operational Status for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayKerberosInstanceOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Kerberos authentication relay name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ticket_cache": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"default_principal": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"item_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_principal": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"client_principal": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"start_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"end_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"renew_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"flags": {
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

func resourceAamAuthenticationRelayKerberosInstanceOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosInstanceOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberosInstanceOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayKerberosInstanceOperOper := setObjectAamAuthenticationRelayKerberosInstanceOperOper(res)
		d.Set("oper", AamAuthenticationRelayKerberosInstanceOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelayKerberosInstanceOperOper(ret edpt.DataAamAuthenticationRelayKerberosInstanceOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ticket_cache":      ret.DtAamAuthenticationRelayKerberosInstanceOper.Oper.TicketCache,
			"default_principal": ret.DtAamAuthenticationRelayKerberosInstanceOper.Oper.DefaultPrincipal,
			"item_list":         setSliceAamAuthenticationRelayKerberosInstanceOperOperItemList(ret.DtAamAuthenticationRelayKerberosInstanceOper.Oper.ItemList),
		},
	}
}

func setSliceAamAuthenticationRelayKerberosInstanceOperOperItemList(d []edpt.AamAuthenticationRelayKerberosInstanceOperOperItemList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["service_principal"] = item.ServicePrincipal
		in["client_principal"] = item.ClientPrincipal
		in["start_time"] = item.StartTime
		in["end_time"] = item.EndTime
		in["renew_time"] = item.RenewTime
		in["flags"] = item.Flags
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationRelayKerberosInstanceOperOper(d []interface{}) edpt.AamAuthenticationRelayKerberosInstanceOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayKerberosInstanceOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TicketCache = in["ticket_cache"].(string)
		ret.DefaultPrincipal = in["default_principal"].(string)
		ret.ItemList = getSliceAamAuthenticationRelayKerberosInstanceOperOperItemList(in["item_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationRelayKerberosInstanceOperOperItemList(d []interface{}) []edpt.AamAuthenticationRelayKerberosInstanceOperOperItemList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayKerberosInstanceOperOperItemList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayKerberosInstanceOperOperItemList
		oi.ServicePrincipal = in["service_principal"].(string)
		oi.ClientPrincipal = in["client_principal"].(string)
		oi.StartTime = in["start_time"].(string)
		oi.EndTime = in["end_time"].(string)
		oi.RenewTime = in["renew_time"].(string)
		oi.Flags = in["flags"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayKerberosInstanceOper(d *schema.ResourceData) edpt.AamAuthenticationRelayKerberosInstanceOper {
	var ret edpt.AamAuthenticationRelayKerberosInstanceOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectAamAuthenticationRelayKerberosInstanceOperOper(d.Get("oper").([]interface{}))
	return ret
}
