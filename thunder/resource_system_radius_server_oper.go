package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemRadiusServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_radius_server_oper`: Operational Status for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemRadiusServerOperRead,

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

func resourceSystemRadiusServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemRadiusServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemRadiusServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemRadiusServerOperOper := setObjectSystemRadiusServerOperOper(res)
		d.Set("oper", SystemRadiusServerOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemRadiusServerOperOper(ret edpt.DataSystemRadiusServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"radius_table_entries_list": setSliceSystemRadiusServerOperOperRadiusTableEntriesList(ret.DtSystemRadiusServerOper.Oper.RadiusTableEntriesList),
			"total_entries":             ret.DtSystemRadiusServerOper.Oper.TotalEntries,
			"custom_attr_name":          ret.DtSystemRadiusServerOper.Oper.CustomAttrName,
			"custom_attr_value":         ret.DtSystemRadiusServerOper.Oper.CustomAttrValue,
			"starts_with":               ret.DtSystemRadiusServerOper.Oper.StartsWith,
			"case_insensitive":          ret.DtSystemRadiusServerOper.Oper.CaseInsensitive,
		},
	}
}

func setSliceSystemRadiusServerOperOperRadiusTableEntriesList(d []edpt.SystemRadiusServerOperOperRadiusTableEntriesList) []map[string]interface{} {
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

func getObjectSystemRadiusServerOperOper(d []interface{}) edpt.SystemRadiusServerOperOper {

	count1 := len(d)
	var ret edpt.SystemRadiusServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusTableEntriesList = getSliceSystemRadiusServerOperOperRadiusTableEntriesList(in["radius_table_entries_list"].([]interface{}))
		ret.TotalEntries = in["total_entries"].(int)
		ret.CustomAttrName = in["custom_attr_name"].(string)
		ret.CustomAttrValue = in["custom_attr_value"].(string)
		ret.StartsWith = in["starts_with"].(int)
		ret.CaseInsensitive = in["case_insensitive"].(int)
	}
	return ret
}

func getSliceSystemRadiusServerOperOperRadiusTableEntriesList(d []interface{}) []edpt.SystemRadiusServerOperOperRadiusTableEntriesList {

	count1 := len(d)
	ret := make([]edpt.SystemRadiusServerOperOperRadiusTableEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemRadiusServerOperOperRadiusTableEntriesList
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

func dataToEndpointSystemRadiusServerOper(d *schema.ResourceData) edpt.SystemRadiusServerOper {
	var ret edpt.SystemRadiusServerOper

	ret.Oper = getObjectSystemRadiusServerOperOper(d.Get("oper").([]interface{}))
	return ret
}
