package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpThreatListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ip_threat_list_oper`: Operational Status for the object ip-threat-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIpThreatListOperRead,

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
									"match_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"in_spe": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
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

func resourceSystemIpThreatListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpThreatListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpThreatListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIpThreatListOperOper := setObjectSystemIpThreatListOperOper(res)
		d.Set("oper", SystemIpThreatListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemIpThreatListOperOper(ret edpt.DataSystemIpThreatListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entries_list": setSliceSystemIpThreatListOperOperEntriesList(ret.DtSystemIpThreatListOper.Oper.EntriesList),
			"v4_address":   ret.DtSystemIpThreatListOper.Oper.V4Address,
			"v4_netmask":   ret.DtSystemIpThreatListOper.Oper.V4Netmask,
			"v6_prefix":    ret.DtSystemIpThreatListOper.Oper.V6Prefix,
		},
	}
}

func setSliceSystemIpThreatListOperOperEntriesList(d []edpt.SystemIpThreatListOperOperEntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		in["match_type"] = item.MatchType
		in["in_spe"] = item.InSpe
		in["age"] = item.Age
		result = append(result, in)
	}
	return result
}

func getObjectSystemIpThreatListOperOper(d []interface{}) edpt.SystemIpThreatListOperOper {

	count1 := len(d)
	var ret edpt.SystemIpThreatListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntriesList = getSliceSystemIpThreatListOperOperEntriesList(in["entries_list"].([]interface{}))
		ret.V4Address = in["v4_address"].(string)
		ret.V4Netmask = in["v4_netmask"].(string)
		ret.V6Prefix = in["v6_prefix"].(string)
	}
	return ret
}

func getSliceSystemIpThreatListOperOperEntriesList(d []interface{}) []edpt.SystemIpThreatListOperOperEntriesList {

	count1 := len(d)
	ret := make([]edpt.SystemIpThreatListOperOperEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIpThreatListOperOperEntriesList
		oi.Ip = in["ip"].(string)
		oi.MatchType = in["match_type"].(string)
		oi.InSpe = in["in_spe"].(string)
		oi.Age = in["age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIpThreatListOper(d *schema.ResourceData) edpt.SystemIpThreatListOper {
	var ret edpt.SystemIpThreatListOper

	ret.Oper = getObjectSystemIpThreatListOperOper(d.Get("oper").([]interface{}))
	return ret
}
