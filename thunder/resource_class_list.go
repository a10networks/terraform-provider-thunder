package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_class_list`: Configure classification list\n\n__PLACEHOLDER__",
		CreateContext: resourceClassListCreate,
		UpdateContext: resourceClassListUpdate,
		ReadContext:   resourceClassListRead,
		DeleteContext: resourceClassListDelete,
		Schema: map[string]*schema.Schema{
			"ac_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ac_match_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': String contains another string; 'ends-with': String ends with another string; 'equals': String equals another string; 'starts-with': String starts with another string;",
						},
						"ac_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Specify key string",
						},
						"ac_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify value string",
						},
						"gtp_rate_limit_policy_str": {
							Type: schema.TypeString, Optional: true, Description: "GTP Rate Limit Template Name",
						},
					},
				},
			},
			"dns": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_match_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Domain contains another string; 'ends-with': Domain ends with another string; 'starts-with': Domain starts-with another string;",
						},
						"dns_match_string": {
							Type: schema.TypeString, Optional: true, Description: "Domain name",
						},
						"dns_lid": {
							Type: schema.TypeInt, Optional: true, Description: "Use Limit ID defined in template (Specify LID index)",
						},
						"dns_glid": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID (Specify global LID index)",
						},
						"shared_partition_dns_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
						},
						"dns_glid_shared": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID",
						},
					},
				},
			},
			"file": {
				Type: schema.TypeInt, ForceNew: true, Optional: true, Default: 0, Description: "Create/Edit a class-list stored as a file",
			},
			"geo_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location": {
							Type: schema.TypeString, Optional: true, Description: "Specify geo-location",
						},
						"geo_location_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 geo-location",
						},
					},
				},
			},
			"ipv4_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address",
						},
						"lid": {
							Type: schema.TypeInt, Optional: true, Description: "Use Limit ID defined in template (Specify LID index)",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID (Specify global LID index)",
						},
						"shared_partition_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
						},
						"glid_shared": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID",
						},
						"lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
						},
						"lsn_radius_profile": {
							Type: schema.TypeInt, Optional: true, Description: "LSN RADIUS Profile Index",
						},
						"gtp_rate_limit_policy_v4": {
							Type: schema.TypeString, Optional: true, Description: "GTP Rate Limit Template Name",
						},
						"age": {
							Type: schema.TypeInt, Optional: true, Description: "Specify age in minutes",
						},
					},
				},
			},
			"ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 host or subnet",
						},
						"v6_lid": {
							Type: schema.TypeInt, Optional: true, Description: "Use Limit ID defined in template (Specify LID index)",
						},
						"v6_glid": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID (Specify global LID index)",
						},
						"shared_partition_v6_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
						},
						"v6_glid_shared": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID",
						},
						"v6_lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
						},
						"v6_lsn_radius_profile": {
							Type: schema.TypeInt, Optional: true, Description: "LSN RADIUS Profile Index",
						},
						"gtp_rate_limit_policy_v6": {
							Type: schema.TypeString, Optional: true, Description: "GTP Rate Limit Template Name",
						},
						"v6_age": {
							Type: schema.TypeInt, Optional: true, Description: "Specify age in minutes",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Specify name of the class list",
			},
			"str_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"str": {
							Type: schema.TypeString, Optional: true, Description: "Specify key string",
						},
						"str_lid_dummy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Limit ID defined in template",
						},
						"str_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LID index",
						},
						"str_glid_dummy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use global Limit ID",
						},
						"str_glid": {
							Type: schema.TypeString, Optional: true, Description: "Global LID index",
						},
						"shared_partition_str_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
						},
						"str_glid_shared": {
							Type: schema.TypeString, Optional: true, Description: "Use global Limit ID",
						},
						"value_str": {
							Type: schema.TypeString, Optional: true, Description: "Specify value string",
						},
					},
				},
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'ac': Make class-list type Aho-Corasick; 'dns': Make class-list type DNS; 'ipv4': Make class-list type IPv4; 'ipv6': Make class-list type IPv6; 'string': Make class-list type String; 'string-case-insensitive': Make class-list type String-case-insensitive. Case insensitive is applied to key string;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceClassListAcList(d []interface{}) []edpt.ClassListAcList {
	count := len(d)
	ret := make([]edpt.ClassListAcList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListAcList
		oi.AcMatchType = in["ac_match_type"].(string)
		oi.AcKeyString = in["ac_key_string"].(string)
		oi.AcValue = in["ac_value"].(string)
		oi.GtpRateLimitPolicyStr = in["gtp_rate_limit_policy_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListDns(d []interface{}) []edpt.ClassListDns {
	count := len(d)
	ret := make([]edpt.ClassListDns, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListDns
		oi.DnsMatchType = in["dns_match_type"].(string)
		oi.DnsMatchString = in["dns_match_string"].(string)
		oi.DnsLid = in["dns_lid"].(int)
		oi.DnsGlid = in["dns_glid"].(string)
		oi.SharedPartitionDnsGlid = in["shared_partition_dns_glid"].(int)
		oi.DnsGlidShared = in["dns_glid_shared"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListGeoList(d []interface{}) []edpt.ClassListGeoList {
	count := len(d)
	ret := make([]edpt.ClassListGeoList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListGeoList
		oi.GeoLocation = in["geo_location"].(string)
		oi.GeoLocationIpv6 = in["geo_location_ipv6"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListIpv4List(d []interface{}) []edpt.ClassListIpv4List {
	count := len(d)
	ret := make([]edpt.ClassListIpv4List, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListIpv4List
		oi.Ipv4addr = in["ipv4addr"].(string)
		oi.Lid = in["lid"].(int)
		oi.Glid = in["glid"].(string)
		oi.SharedPartitionGlid = in["shared_partition_glid"].(int)
		oi.GlidShared = in["glid_shared"].(string)
		oi.LsnLid = in["lsn_lid"].(int)
		oi.LsnRadiusProfile = in["lsn_radius_profile"].(int)
		oi.GtpRateLimitPolicyV4 = in["gtp_rate_limit_policy_v4"].(string)
		oi.Age = in["age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListIpv6List(d []interface{}) []edpt.ClassListIpv6List {
	count := len(d)
	ret := make([]edpt.ClassListIpv6List, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListIpv6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.V6Lid = in["v6_lid"].(int)
		oi.V6Glid = in["v6_glid"].(string)
		oi.SharedPartitionV6Glid = in["shared_partition_v6_glid"].(int)
		oi.V6GlidShared = in["v6_glid_shared"].(string)
		oi.V6LsnLid = in["v6_lsn_lid"].(int)
		oi.V6LsnRadiusProfile = in["v6_lsn_radius_profile"].(int)
		oi.GtpRateLimitPolicyV6 = in["gtp_rate_limit_policy_v6"].(string)
		oi.V6Age = in["v6_age"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceClassListStrList(d []interface{}) []edpt.ClassListStrList {
	count := len(d)
	ret := make([]edpt.ClassListStrList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ClassListStrList
		oi.Str = in["str"].(string)
		oi.StrLidDummy = in["str_lid_dummy"].(int)
		oi.StrLid = in["str_lid"].(int)
		oi.StrGlidDummy = in["str_glid_dummy"].(int)
		oi.StrGlid = in["str_glid"].(string)
		oi.SharedPartitionStrGlid = in["shared_partition_str_glid"].(int)
		oi.StrGlidShared = in["str_glid_shared"].(string)
		oi.ValueStr = in["value_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointClassList(d *schema.ResourceData) edpt.ClassList {
	var ret edpt.ClassList
	ret.Inst.AcList = getSliceClassListAcList(d.Get("ac_list").([]interface{}))
	ret.Inst.Dns = getSliceClassListDns(d.Get("dns").([]interface{}))
	ret.Inst.File = d.Get("file").(int)
	ret.Inst.GeoList = getSliceClassListGeoList(d.Get("geo_list").([]interface{}))
	ret.Inst.Ipv4List = getSliceClassListIpv4List(d.Get("ipv4_list").([]interface{}))
	ret.Inst.Ipv6List = getSliceClassListIpv6List(d.Get("ipv6_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.StrList = getSliceClassListStrList(d.Get("str_list").([]interface{}))
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
