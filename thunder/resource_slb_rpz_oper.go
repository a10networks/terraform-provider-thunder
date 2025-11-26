package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRpzOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_rpz_oper`: Operational Status for the object rpz\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbRpzOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dns_template_bound": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"rpz_file_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rpz_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rpz_rule_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
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
					},
				},
			},
		},
	}
}

func resourceSlbRpzOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRpzOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRpzOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbRpzOperOper := setObjectSlbRpzOperOper(res)
		d.Set("oper", SlbRpzOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbRpzOperOper(ret edpt.DataSlbRpzOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter_entry":               ret.DtSlbRpzOper.Oper.Filter_entry,
			"file_list":                  setSliceSlbRpzOperOperFileList(ret.DtSlbRpzOper.Oper.FileList),
			"rpz_file_size_max":          ret.DtSlbRpzOper.Oper.RpzFileSizeMax,
			"rpz_count":                  ret.DtSlbRpzOper.Oper.RpzCount,
			"rpz_rule_count":             ret.DtSlbRpzOper.Oper.RpzRuleCount,
			"class_list":                 ret.DtSlbRpzOper.Oper.ClassList,
			"type":                       ret.DtSlbRpzOper.Oper.Type,
			"file_or_string":             ret.DtSlbRpzOper.Oper.FileOrString,
			"user_tag":                   ret.DtSlbRpzOper.Oper.UserTag,
			"ipv4_total_single_ip":       ret.DtSlbRpzOper.Oper.Ipv4TotalSingleIp,
			"ipv4_total_subnet":          ret.DtSlbRpzOper.Oper.Ipv4TotalSubnet,
			"ipv6_total_single_ip":       ret.DtSlbRpzOper.Oper.Ipv6TotalSingleIp,
			"ipv6_total_subnet":          ret.DtSlbRpzOper.Oper.Ipv6TotalSubnet,
			"dns_total_entries":          ret.DtSlbRpzOper.Oper.DnsTotalEntries,
			"string_total_entries":       ret.DtSlbRpzOper.Oper.StringTotalEntries,
			"ac_total_entries":           ret.DtSlbRpzOper.Oper.AcTotalEntries,
			"geo_location_total_entries": ret.DtSlbRpzOper.Oper.GeoLocationTotalEntries,
			"ipv4_entries":               setSliceSlbRpzOperOperIpv4Entries(ret.DtSlbRpzOper.Oper.Ipv4Entries),
			"ipv6_entries":               setSliceSlbRpzOperOperIpv6Entries(ret.DtSlbRpzOper.Oper.Ipv6Entries),
			"dns_entries":                setSliceSlbRpzOperOperDnsEntries(ret.DtSlbRpzOper.Oper.DnsEntries),
		},
	}
}

func setSliceSlbRpzOperOperFileList(d []edpt.SlbRpzOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["dns_template_bound"] = item.Dns_template_bound
		result = append(result, in)
	}
	return result
}

func setSliceSlbRpzOperOperIpv4Entries(d []edpt.SlbRpzOperOperIpv4Entries) []map[string]interface{} {
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

func setSliceSlbRpzOperOperIpv6Entries(d []edpt.SlbRpzOperOperIpv6Entries) []map[string]interface{} {
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

func setSliceSlbRpzOperOperDnsEntries(d []edpt.SlbRpzOperOperDnsEntries) []map[string]interface{} {
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

func getObjectSlbRpzOperOper(d []interface{}) edpt.SlbRpzOperOper {

	count1 := len(d)
	var ret edpt.SlbRpzOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Filter_entry = in["filter_entry"].(string)
		ret.FileList = getSliceSlbRpzOperOperFileList(in["file_list"].([]interface{}))
		ret.RpzFileSizeMax = in["rpz_file_size_max"].(int)
		ret.RpzCount = in["rpz_count"].(int)
		ret.RpzRuleCount = in["rpz_rule_count"].(int)
		ret.ClassList = in["class_list"].(string)
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
		ret.Ipv4Entries = getSliceSlbRpzOperOperIpv4Entries(in["ipv4_entries"].([]interface{}))
		ret.Ipv6Entries = getSliceSlbRpzOperOperIpv6Entries(in["ipv6_entries"].([]interface{}))
		ret.DnsEntries = getSliceSlbRpzOperOperDnsEntries(in["dns_entries"].([]interface{}))
	}
	return ret
}

func getSliceSlbRpzOperOperFileList(d []interface{}) []edpt.SlbRpzOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.SlbRpzOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRpzOperOperFileList
		oi.File = in["file"].(string)
		oi.Dns_template_bound = in["dns_template_bound"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbRpzOperOperIpv4Entries(d []interface{}) []edpt.SlbRpzOperOperIpv4Entries {

	count1 := len(d)
	ret := make([]edpt.SlbRpzOperOperIpv4Entries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRpzOperOperIpv4Entries
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

func getSliceSlbRpzOperOperIpv6Entries(d []interface{}) []edpt.SlbRpzOperOperIpv6Entries {

	count1 := len(d)
	ret := make([]edpt.SlbRpzOperOperIpv6Entries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRpzOperOperIpv6Entries
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

func getSliceSlbRpzOperOperDnsEntries(d []interface{}) []edpt.SlbRpzOperOperDnsEntries {

	count1 := len(d)
	ret := make([]edpt.SlbRpzOperOperDnsEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRpzOperOperDnsEntries
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

func dataToEndpointSlbRpzOper(d *schema.ResourceData) edpt.SlbRpzOper {
	var ret edpt.SlbRpzOper

	ret.Oper = getObjectSlbRpzOperOper(d.Get("oper").([]interface{}))
	return ret
}
