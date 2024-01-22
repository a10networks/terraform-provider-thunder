package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDynamicClassListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dynamic_class_list_oper`: Operational Status for the object dynamic-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDynamicClassListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"file_or_string": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4_total_single_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_total_subnet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_total_single_ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv6_total_subnet": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dns_total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"string_total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ac_total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"geo_location_total_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv4_lid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_glid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_lsn_lid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_lsn_radius_profile": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"ipv6_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6_lid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_glid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_lsn_lid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_lsn_radius_profile": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"dns_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_match_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dns_match_string": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dns_lid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dns_glid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dns_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"string_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"string_key": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"string_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"string_lid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"string_glid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"string_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"ac_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ac_match_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ac_match_string": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ac_match_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ac_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"geo_location_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_location": {
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

func resourceDdosDynamicClassListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDynamicClassListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDynamicClassListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDynamicClassListOperOper := setObjectDdosDynamicClassListOperOper(res)
		d.Set("oper", DdosDynamicClassListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDynamicClassListOperOper(ret edpt.DataDdosDynamicClassListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"type":                       ret.DtDdosDynamicClassListOper.Oper.Type,
			"file_or_string":             ret.DtDdosDynamicClassListOper.Oper.FileOrString,
			"user_tag":                   ret.DtDdosDynamicClassListOper.Oper.UserTag,
			"ipv4_total_single_ip":       ret.DtDdosDynamicClassListOper.Oper.Ipv4TotalSingleIp,
			"ipv4_total_subnet":          ret.DtDdosDynamicClassListOper.Oper.Ipv4TotalSubnet,
			"ipv6_total_single_ip":       ret.DtDdosDynamicClassListOper.Oper.Ipv6TotalSingleIp,
			"ipv6_total_subnet":          ret.DtDdosDynamicClassListOper.Oper.Ipv6TotalSubnet,
			"dns_total_entries":          ret.DtDdosDynamicClassListOper.Oper.DnsTotalEntries,
			"string_total_entries":       ret.DtDdosDynamicClassListOper.Oper.StringTotalEntries,
			"ac_total_entries":           ret.DtDdosDynamicClassListOper.Oper.AcTotalEntries,
			"geo_location_total_entries": ret.DtDdosDynamicClassListOper.Oper.GeoLocationTotalEntries,
			"ipv4_entries":               setSliceDdosDynamicClassListOperOperIpv4Entries(ret.DtDdosDynamicClassListOper.Oper.Ipv4Entries),
			"ipv6_entries":               setSliceDdosDynamicClassListOperOperIpv6Entries(ret.DtDdosDynamicClassListOper.Oper.Ipv6Entries),
			"dns_entries":                setSliceDdosDynamicClassListOperOperDnsEntries(ret.DtDdosDynamicClassListOper.Oper.DnsEntries),
			"string_entries":             setSliceDdosDynamicClassListOperOperStringEntries(ret.DtDdosDynamicClassListOper.Oper.StringEntries),
			"ac_entries":                 setSliceDdosDynamicClassListOperOperAcEntries(ret.DtDdosDynamicClassListOper.Oper.AcEntries),
			"geo_location_entries":       setSliceDdosDynamicClassListOperOperGeoLocationEntries(ret.DtDdosDynamicClassListOper.Oper.GeoLocationEntries),
		},
	}
}

func setSliceDdosDynamicClassListOperOperIpv4Entries(d []edpt.DdosDynamicClassListOperOperIpv4Entries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv4_lid"] = item.Ipv4Lid
		in["ipv4_glid"] = item.Ipv4Glid
		in["ipv4_lsn_lid"] = item.Ipv4LsnLid
		in["ipv4_lsn_radius_profile"] = item.Ipv4LsnRadiusProfile
		in["ipv4_hit_count"] = item.Ipv4HitCount
		in["ipv4_age"] = item.Ipv4Age
		result = append(result, in)
	}
	return result
}

func setSliceDdosDynamicClassListOperOperIpv6Entries(d []edpt.DdosDynamicClassListOperOperIpv6Entries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv6addr"] = item.Ipv6addr
		in["ipv6_lid"] = item.Ipv6Lid
		in["ipv6_glid"] = item.Ipv6Glid
		in["ipv6_lsn_lid"] = item.Ipv6LsnLid
		in["ipv6_lsn_radius_profile"] = item.Ipv6LsnRadiusProfile
		in["ipv6_hit_count"] = item.Ipv6HitCount
		in["ipv6_age"] = item.Ipv6Age
		result = append(result, in)
	}
	return result
}

func setSliceDdosDynamicClassListOperOperDnsEntries(d []edpt.DdosDynamicClassListOperOperDnsEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dns_match_type"] = item.DnsMatchType
		in["dns_match_string"] = item.DnsMatchString
		in["dns_lid"] = item.DnsLid
		in["dns_glid"] = item.DnsGlid
		in["dns_hit_count"] = item.DnsHitCount
		result = append(result, in)
	}
	return result
}

func setSliceDdosDynamicClassListOperOperStringEntries(d []edpt.DdosDynamicClassListOperOperStringEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["string_key"] = item.StringKey
		in["string_value"] = item.StringValue
		in["string_lid"] = item.StringLid
		in["string_glid"] = item.StringGlid
		in["string_hit_count"] = item.StringHitCount
		result = append(result, in)
	}
	return result
}

func setSliceDdosDynamicClassListOperOperAcEntries(d []edpt.DdosDynamicClassListOperOperAcEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ac_match_type"] = item.AcMatchType
		in["ac_match_string"] = item.AcMatchString
		in["ac_match_value"] = item.AcMatchValue
		in["ac_hit_count"] = item.AcHitCount
		result = append(result, in)
	}
	return result
}

func setSliceDdosDynamicClassListOperOperGeoLocationEntries(d []edpt.DdosDynamicClassListOperOperGeoLocationEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["geo_location"] = item.GeoLocation
		result = append(result, in)
	}
	return result
}

func getObjectDdosDynamicClassListOperOper(d []interface{}) edpt.DdosDynamicClassListOperOper {

	count1 := len(d)
	var ret edpt.DdosDynamicClassListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = in["type"].(string)
		ret.FileOrString = in["file_or_string"].(string)
		ret.UserTag = in["user_tag"].(string)
		ret.Ipv4TotalSingleIp = in["ipv4_total_single_ip"].(int)
		ret.Ipv4TotalSubnet = in["ipv4_total_subnet"].(int)
		ret.Ipv6TotalSingleIp = in["ipv6_total_single_ip"].(int)
		ret.Ipv6TotalSubnet = in["ipv6_total_subnet"].(int)
		ret.DnsTotalEntries = in["dns_total_entries"].(int)
		ret.StringTotalEntries = in["string_total_entries"].(int)
		ret.AcTotalEntries = in["ac_total_entries"].(int)
		ret.GeoLocationTotalEntries = in["geo_location_total_entries"].(int)
		ret.Ipv4Entries = getSliceDdosDynamicClassListOperOperIpv4Entries(in["ipv4_entries"].([]interface{}))
		ret.Ipv6Entries = getSliceDdosDynamicClassListOperOperIpv6Entries(in["ipv6_entries"].([]interface{}))
		ret.DnsEntries = getSliceDdosDynamicClassListOperOperDnsEntries(in["dns_entries"].([]interface{}))
		ret.StringEntries = getSliceDdosDynamicClassListOperOperStringEntries(in["string_entries"].([]interface{}))
		ret.AcEntries = getSliceDdosDynamicClassListOperOperAcEntries(in["ac_entries"].([]interface{}))
		ret.GeoLocationEntries = getSliceDdosDynamicClassListOperOperGeoLocationEntries(in["geo_location_entries"].([]interface{}))
	}
	return ret
}

func getSliceDdosDynamicClassListOperOperIpv4Entries(d []interface{}) []edpt.DdosDynamicClassListOperOperIpv4Entries {

	count1 := len(d)
	ret := make([]edpt.DdosDynamicClassListOperOperIpv4Entries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDynamicClassListOperOperIpv4Entries
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv4Lid = in["ipv4_lid"].(int)
		oi.Ipv4Glid = in["ipv4_glid"].(int)
		oi.Ipv4LsnLid = in["ipv4_lsn_lid"].(int)
		oi.Ipv4LsnRadiusProfile = in["ipv4_lsn_radius_profile"].(int)
		oi.Ipv4HitCount = in["ipv4_hit_count"].(int)
		oi.Ipv4Age = in["ipv4_age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDynamicClassListOperOperIpv6Entries(d []interface{}) []edpt.DdosDynamicClassListOperOperIpv6Entries {

	count1 := len(d)
	ret := make([]edpt.DdosDynamicClassListOperOperIpv6Entries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDynamicClassListOperOperIpv6Entries
		oi.Ipv6addr = in["ipv6addr"].(string)
		oi.Ipv6Lid = in["ipv6_lid"].(int)
		oi.Ipv6Glid = in["ipv6_glid"].(int)
		oi.Ipv6LsnLid = in["ipv6_lsn_lid"].(int)
		oi.Ipv6LsnRadiusProfile = in["ipv6_lsn_radius_profile"].(int)
		oi.Ipv6HitCount = in["ipv6_hit_count"].(int)
		oi.Ipv6Age = in["ipv6_age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDynamicClassListOperOperDnsEntries(d []interface{}) []edpt.DdosDynamicClassListOperOperDnsEntries {

	count1 := len(d)
	ret := make([]edpt.DdosDynamicClassListOperOperDnsEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDynamicClassListOperOperDnsEntries
		oi.DnsMatchType = in["dns_match_type"].(string)
		oi.DnsMatchString = in["dns_match_string"].(string)
		oi.DnsLid = in["dns_lid"].(int)
		oi.DnsGlid = in["dns_glid"].(int)
		oi.DnsHitCount = in["dns_hit_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDynamicClassListOperOperStringEntries(d []interface{}) []edpt.DdosDynamicClassListOperOperStringEntries {

	count1 := len(d)
	ret := make([]edpt.DdosDynamicClassListOperOperStringEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDynamicClassListOperOperStringEntries
		oi.StringKey = in["string_key"].(string)
		oi.StringValue = in["string_value"].(string)
		oi.StringLid = in["string_lid"].(int)
		oi.StringGlid = in["string_glid"].(int)
		oi.StringHitCount = in["string_hit_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDynamicClassListOperOperAcEntries(d []interface{}) []edpt.DdosDynamicClassListOperOperAcEntries {

	count1 := len(d)
	ret := make([]edpt.DdosDynamicClassListOperOperAcEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDynamicClassListOperOperAcEntries
		oi.AcMatchType = in["ac_match_type"].(string)
		oi.AcMatchString = in["ac_match_string"].(string)
		oi.AcMatchValue = in["ac_match_value"].(string)
		oi.AcHitCount = in["ac_hit_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDynamicClassListOperOperGeoLocationEntries(d []interface{}) []edpt.DdosDynamicClassListOperOperGeoLocationEntries {

	count1 := len(d)
	ret := make([]edpt.DdosDynamicClassListOperOperGeoLocationEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDynamicClassListOperOperGeoLocationEntries
		oi.GeoLocation = in["geo_location"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDynamicClassListOper(d *schema.ResourceData) edpt.DdosDynamicClassListOper {
	var ret edpt.DdosDynamicClassListOper

	ret.Oper = getObjectDdosDynamicClassListOperOper(d.Get("oper").([]interface{}))
	return ret
}
