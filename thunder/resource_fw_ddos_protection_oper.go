package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwDdosProtectionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_ddos_protection_oper`: Operational Status for the object ddos-protection\n\n__PLACEHOLDER__",
		ReadContext: resourceFwDdosProtectionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entries_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"prefix": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rule_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"expiration": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hints": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hash": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"v4_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"v4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"v6_prefix": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwDdosProtectionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwDdosProtectionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwDdosProtectionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwDdosProtectionOperOper := setObjectFwDdosProtectionOperOper(res)
		d.Set("oper", FwDdosProtectionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwDdosProtectionOperOper(ret edpt.DataFwDdosProtectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entries_list": setSliceFwDdosProtectionOperOperEntriesList(ret.DtFwDdosProtectionOper.Oper.EntriesList),
			"details":      ret.DtFwDdosProtectionOper.Oper.Details,
			"v4_address":   ret.DtFwDdosProtectionOper.Oper.V4Address,
			"v4_netmask":   ret.DtFwDdosProtectionOper.Oper.V4Netmask,
			"v6_prefix":    ret.DtFwDdosProtectionOper.Oper.V6Prefix,
		},
	}
}

func setSliceFwDdosProtectionOperOperEntriesList(d []edpt.FwDdosProtectionOperOperEntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		in["prefix"] = item.Prefix
		in["rule_name"] = item.RuleName
		in["pps"] = item.Pps
		in["expiration"] = item.Expiration
		in["hints"] = item.Hints
		in["hash"] = item.Hash
		in["lid"] = item.Lid
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func getObjectFwDdosProtectionOperOper(d []interface{}) edpt.FwDdosProtectionOperOper {

	count1 := len(d)
	var ret edpt.FwDdosProtectionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntriesList = getSliceFwDdosProtectionOperOperEntriesList(in["entries_list"].([]interface{}))
		ret.Details = in["details"].(int)
		ret.V4Address = in["v4_address"].(string)
		ret.V4Netmask = in["v4_netmask"].(string)
		ret.V6Prefix = in["v6_prefix"].(string)
	}
	return ret
}

func getSliceFwDdosProtectionOperOperEntriesList(d []interface{}) []edpt.FwDdosProtectionOperOperEntriesList {

	count1 := len(d)
	ret := make([]edpt.FwDdosProtectionOperOperEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwDdosProtectionOperOperEntriesList
		oi.Ip = in["ip"].(string)
		oi.Prefix = in["prefix"].(int)
		oi.RuleName = in["rule_name"].(string)
		oi.Pps = in["pps"].(int)
		oi.Expiration = in["expiration"].(int)
		oi.Hints = in["hints"].(string)
		oi.Hash = in["hash"].(int)
		oi.Lid = in["lid"].(int)
		oi.Rate = in["rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwDdosProtectionOper(d *schema.ResourceData) edpt.FwDdosProtectionOper {
	var ret edpt.FwDdosProtectionOper

	ret.Oper = getObjectFwDdosProtectionOperOper(d.Get("oper").([]interface{}))
	return ret
}
