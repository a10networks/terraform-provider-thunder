package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRadiusServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_radius_server_oper`: Operational Status for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRadiusServerOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"radius_table_entries_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"inside_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"prefix_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msisdn": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"imei": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"imsi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"custom1_attr_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"custom2_attr_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"custom3_attr_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"custom4_attr_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"custom5_attr_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"custom6_attr_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_obsolete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"custom_attr_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"custom_attr_value": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"starts_with": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"case_insensitive": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnRadiusServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRadiusServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRadiusServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRadiusServerOperOper := setObjectCgnv6LsnRadiusServerOperOper(res)
		d.Set("oper", Cgnv6LsnRadiusServerOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRadiusServerOperOper(ret edpt.DataCgnv6LsnRadiusServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"radius_table_entries_list": setSliceCgnv6LsnRadiusServerOperOperRadiusTableEntriesList(ret.DtCgnv6LsnRadiusServerOper.Oper.RadiusTableEntriesList),
			"total_entries":             ret.DtCgnv6LsnRadiusServerOper.Oper.TotalEntries,
			"custom_attr_name":          ret.DtCgnv6LsnRadiusServerOper.Oper.CustomAttrName,
			"custom_attr_value":         ret.DtCgnv6LsnRadiusServerOper.Oper.CustomAttrValue,
			"starts_with":               ret.DtCgnv6LsnRadiusServerOper.Oper.StartsWith,
			"case_insensitive":          ret.DtCgnv6LsnRadiusServerOper.Oper.CaseInsensitive,
		},
	}
}

func setSliceCgnv6LsnRadiusServerOperOperRadiusTableEntriesList(d []edpt.Cgnv6LsnRadiusServerOperOperRadiusTableEntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["inside_ip"] = item.InsideIp
		in["inside_ipv6"] = item.InsideIpv6
		in["prefix_len"] = item.PrefixLen
		in["msisdn"] = item.Msisdn
		in["imei"] = item.Imei
		in["imsi"] = item.Imsi
		in["custom1_attr_value"] = item.Custom1AttrValue
		in["custom2_attr_value"] = item.Custom2AttrValue
		in["custom3_attr_value"] = item.Custom3AttrValue
		in["custom4_attr_value"] = item.Custom4AttrValue
		in["custom5_attr_value"] = item.Custom5AttrValue
		in["custom6_attr_value"] = item.Custom6AttrValue
		in["is_obsolete"] = item.IsObsolete
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6LsnRadiusServerOperOper(d []interface{}) edpt.Cgnv6LsnRadiusServerOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRadiusServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusTableEntriesList = getSliceCgnv6LsnRadiusServerOperOperRadiusTableEntriesList(in["radius_table_entries_list"].([]interface{}))
		ret.TotalEntries = in["total_entries"].(int)
		ret.CustomAttrName = in["custom_attr_name"].(string)
		ret.CustomAttrValue = in["custom_attr_value"].(string)
		ret.StartsWith = in["starts_with"].(int)
		ret.CaseInsensitive = in["case_insensitive"].(int)
	}
	return ret
}

func getSliceCgnv6LsnRadiusServerOperOperRadiusTableEntriesList(d []interface{}) []edpt.Cgnv6LsnRadiusServerOperOperRadiusTableEntriesList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRadiusServerOperOperRadiusTableEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRadiusServerOperOperRadiusTableEntriesList
		oi.InsideIp = in["inside_ip"].(string)
		oi.InsideIpv6 = in["inside_ipv6"].(string)
		oi.PrefixLen = in["prefix_len"].(int)
		oi.Msisdn = in["msisdn"].(string)
		oi.Imei = in["imei"].(string)
		oi.Imsi = in["imsi"].(string)
		oi.Custom1AttrValue = in["custom1_attr_value"].(string)
		oi.Custom2AttrValue = in["custom2_attr_value"].(string)
		oi.Custom3AttrValue = in["custom3_attr_value"].(string)
		oi.Custom4AttrValue = in["custom4_attr_value"].(string)
		oi.Custom5AttrValue = in["custom5_attr_value"].(string)
		oi.Custom6AttrValue = in["custom6_attr_value"].(string)
		oi.IsObsolete = in["is_obsolete"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnRadiusServerOper(d *schema.ResourceData) edpt.Cgnv6LsnRadiusServerOper {
	var ret edpt.Cgnv6LsnRadiusServerOper

	ret.Oper = getObjectCgnv6LsnRadiusServerOperOper(d.Get("oper").([]interface{}))
	return ret
}
