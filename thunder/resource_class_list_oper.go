package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceClassListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_class_list_oper`: Operational Status for the object class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceClassListOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the class list",
			},
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
									"ipv4_gtp_policy": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv4_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_rpz_type": {
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
									"ipv6_gtp_policy": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_rpz_type": {
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
									"dns_rpz_type": {
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
									"ac_gtp_policy": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceClassListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClassListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClassListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ClassListOperOper := setObjectClassListOperOper(res)
		d.Set("oper", ClassListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectClassListOperOper(ret edpt.DataClassListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"type":                       ret.DtClassListOper.Oper.Type,
			"file_or_string":             ret.DtClassListOper.Oper.FileOrString,
			"user_tag":                   ret.DtClassListOper.Oper.UserTag,
			"ipv4_total_single_ip":       ret.DtClassListOper.Oper.Ipv4TotalSingleIp,
			"ipv4_total_subnet":          ret.DtClassListOper.Oper.Ipv4TotalSubnet,
			"ipv6_total_single_ip":       ret.DtClassListOper.Oper.Ipv6TotalSingleIp,
			"ipv6_total_subnet":          ret.DtClassListOper.Oper.Ipv6TotalSubnet,
			"dns_total_entries":          ret.DtClassListOper.Oper.DnsTotalEntries,
			"string_total_entries":       ret.DtClassListOper.Oper.StringTotalEntries,
			"ac_total_entries":           ret.DtClassListOper.Oper.AcTotalEntries,
			"geo_location_total_entries": ret.DtClassListOper.Oper.GeoLocationTotalEntries,
			"ipv4_entries":               setSliceClassListOperOperIpv4Entries(ret.DtClassListOper.Oper.Ipv4Entries),
			"ipv6_entries":               setSliceClassListOperOperIpv6Entries(ret.DtClassListOper.Oper.Ipv6Entries),
			"dns_entries":                setSliceClassListOperOperDnsEntries(ret.DtClassListOper.Oper.DnsEntries),
			"string_entries":             setSliceClassListOperOperStringEntries(ret.DtClassListOper.Oper.StringEntries),
			"ac_entries":                 setSliceClassListOperOperAcEntries(ret.DtClassListOper.Oper.AcEntries),
			"geo_location_entries":       setSliceClassListOperOperGeoLocationEntries(ret.DtClassListOper.Oper.GeoLocationEntries),
		},
	}
}

func setSliceClassListOperOperIpv4Entries(d []edpt.ClassListOperOperIpv4Entries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv4_lid"] = item.Ipv4Lid
		in["ipv4_glid"] = item.Ipv4Glid
		in["ipv4_lsn_lid"] = item.Ipv4LsnLid
		in["ipv4_lsn_radius_profile"] = item.Ipv4LsnRadiusProfile
		in["ipv4_gtp_policy"] = item.Ipv4GtpPolicy
		in["ipv4_hit_count"] = item.Ipv4HitCount
		in["ipv4_age"] = item.Ipv4Age
		in["ipv4_rpz_type"] = item.Ipv4RpzType
		result = append(result, in)
	}
	return result
}

func setSliceClassListOperOperIpv6Entries(d []edpt.ClassListOperOperIpv6Entries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipv6addr"] = item.Ipv6addr
		in["ipv6_lid"] = item.Ipv6Lid
		in["ipv6_glid"] = item.Ipv6Glid
		in["ipv6_lsn_lid"] = item.Ipv6LsnLid
		in["ipv6_lsn_radius_profile"] = item.Ipv6LsnRadiusProfile
		in["ipv6_gtp_policy"] = item.Ipv6GtpPolicy
		in["ipv6_hit_count"] = item.Ipv6HitCount
		in["ipv6_age"] = item.Ipv6Age
		in["ipv6_rpz_type"] = item.Ipv6RpzType
		result = append(result, in)
	}
	return result
}

func setSliceClassListOperOperDnsEntries(d []edpt.ClassListOperOperDnsEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dns_match_type"] = item.DnsMatchType
		in["dns_match_string"] = item.DnsMatchString
		in["dns_lid"] = item.DnsLid
		in["dns_glid"] = item.DnsGlid
		in["dns_hit_count"] = item.DnsHitCount
		in["dns_rpz_type"] = item.DnsRpzType
		result = append(result, in)
	}
	return result
}

func setSliceClassListOperOperStringEntries(d []edpt.ClassListOperOperStringEntries) []map[string]interface{} {
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

func setSliceClassListOperOperAcEntries(d []edpt.ClassListOperOperAcEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ac_match_type"] = item.AcMatchType
		in["ac_match_string"] = item.AcMatchString
		in["ac_match_value"] = item.AcMatchValue
		in["ac_hit_count"] = item.AcHitCount
		in["ac_gtp_policy"] = item.AcGtpPolicy
		result = append(result, in)
	}
	return result
}

func setSliceClassListOperOperGeoLocationEntries(d []edpt.ClassListOperOperGeoLocationEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["geo_location"] = item.GeoLocation
		result = append(result, in)
	}
	return result
}

func getObjectClassListOperOper(d []interface{}) edpt.ClassListOperOper {

	count1 := len(d)
	var ret edpt.ClassListOperOper
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
		ret.Ipv4Entries = getSliceClassListOperOperIpv4Entries(in["ipv4_entries"].([]interface{}))
		ret.Ipv6Entries = getSliceClassListOperOperIpv6Entries(in["ipv6_entries"].([]interface{}))
		ret.DnsEntries = getSliceClassListOperOperDnsEntries(in["dns_entries"].([]interface{}))
		ret.StringEntries = getSliceClassListOperOperStringEntries(in["string_entries"].([]interface{}))
		ret.AcEntries = getSliceClassListOperOperAcEntries(in["ac_entries"].([]interface{}))
		ret.GeoLocationEntries = getSliceClassListOperOperGeoLocationEntries(in["geo_location_entries"].([]interface{}))
	}
	return ret
}

func getSliceClassListOperOperIpv4Entries(d []interface{}) []edpt.ClassListOperOperIpv4Entries {

	count1 := len(d)
	ret := make([]edpt.ClassListOperOperIpv4Entries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListOperOperIpv4Entries
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv4Lid = in["ipv4_lid"].(int)
		oi.Ipv4Glid = in["ipv4_glid"].(int)
		oi.Ipv4LsnLid = in["ipv4_lsn_lid"].(int)
		oi.Ipv4LsnRadiusProfile = in["ipv4_lsn_radius_profile"].(int)
		oi.Ipv4GtpPolicy = in["ipv4_gtp_policy"].(string)
		oi.Ipv4HitCount = in["ipv4_hit_count"].(int)
		oi.Ipv4Age = in["ipv4_age"].(int)
		oi.Ipv4RpzType = in["ipv4_rpz_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListOperOperIpv6Entries(d []interface{}) []edpt.ClassListOperOperIpv6Entries {

	count1 := len(d)
	ret := make([]edpt.ClassListOperOperIpv6Entries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListOperOperIpv6Entries
		oi.Ipv6addr = in["ipv6addr"].(string)
		oi.Ipv6Lid = in["ipv6_lid"].(int)
		oi.Ipv6Glid = in["ipv6_glid"].(int)
		oi.Ipv6LsnLid = in["ipv6_lsn_lid"].(int)
		oi.Ipv6LsnRadiusProfile = in["ipv6_lsn_radius_profile"].(int)
		oi.Ipv6GtpPolicy = in["ipv6_gtp_policy"].(string)
		oi.Ipv6HitCount = in["ipv6_hit_count"].(int)
		oi.Ipv6Age = in["ipv6_age"].(int)
		oi.Ipv6RpzType = in["ipv6_rpz_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListOperOperDnsEntries(d []interface{}) []edpt.ClassListOperOperDnsEntries {

	count1 := len(d)
	ret := make([]edpt.ClassListOperOperDnsEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListOperOperDnsEntries
		oi.DnsMatchType = in["dns_match_type"].(string)
		oi.DnsMatchString = in["dns_match_string"].(string)
		oi.DnsLid = in["dns_lid"].(int)
		oi.DnsGlid = in["dns_glid"].(int)
		oi.DnsHitCount = in["dns_hit_count"].(int)
		oi.DnsRpzType = in["dns_rpz_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListOperOperStringEntries(d []interface{}) []edpt.ClassListOperOperStringEntries {

	count1 := len(d)
	ret := make([]edpt.ClassListOperOperStringEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListOperOperStringEntries
		oi.StringKey = in["string_key"].(string)
		oi.StringValue = in["string_value"].(string)
		oi.StringLid = in["string_lid"].(int)
		oi.StringGlid = in["string_glid"].(int)
		oi.StringHitCount = in["string_hit_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListOperOperAcEntries(d []interface{}) []edpt.ClassListOperOperAcEntries {

	count1 := len(d)
	ret := make([]edpt.ClassListOperOperAcEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListOperOperAcEntries
		oi.AcMatchType = in["ac_match_type"].(string)
		oi.AcMatchString = in["ac_match_string"].(string)
		oi.AcMatchValue = in["ac_match_value"].(string)
		oi.AcHitCount = in["ac_hit_count"].(int)
		oi.AcGtpPolicy = in["ac_gtp_policy"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListOperOperGeoLocationEntries(d []interface{}) []edpt.ClassListOperOperGeoLocationEntries {

	count1 := len(d)
	ret := make([]edpt.ClassListOperOperGeoLocationEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListOperOperGeoLocationEntries
		oi.GeoLocation = in["geo_location"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointClassListOper(d *schema.ResourceData) edpt.ClassListOper {
	var ret edpt.ClassListOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectClassListOperOper(d.Get("oper").([]interface{}))
	return ret
}
