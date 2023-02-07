package thunder

import (
	"context"
	"fmt"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_class_list` configures classification list\n\n",
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
							ValidateFunc: validation.StringInSlice([]string{"contains", "ends-with", "equals", "starts-with"}, false),
						},
						"ac_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Specify key string",
							ValidateFunc: validation.StringLenBetween(1, 255),
						},
						"ac_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify value string",
							ValidateFunc:  validation.StringLenBetween(1, 639),
						},
						"gtp_rate_limit_policy_str": {
							Type: schema.TypeString, Optional: true, Description: "GTP Rate Limit Template Name",
							ValidateFunc:  validation.StringLenBetween(1, 128),
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
							ValidateFunc: validation.StringInSlice([]string{"contains", "ends-with", "starts-with"}, false),
						},
						"dns_match_string": {
							Type: schema.TypeString, Optional: true, Description: "Domain name",
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"dns_lid": {
							Type: schema.TypeInt, Optional: true, Description: "Use Limit ID defined in template (Specify LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"dns_glid": {
							Type: schema.TypeInt, Optional: true, Description: "Use global Limit ID (Specify global LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"shared_partition_dns_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
							ValidateFunc:  validation.IntBetween(0, 1),
						},
						"dns_glid_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Use global Limit ID",
							ValidateFunc: validation.IntBetween(1, 1023),
						},
					},
				},
			},
			"file": {
				Type: schema.TypeInt, ForceNew: true, Optional: true, Default: 0, Description: "Create/Edit a class-list stored as a file",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"ipv4_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address",
							ValidateFunc: validation.IsCIDR,
						},
						"lid": {
							Type: schema.TypeInt, Optional: true, Description: "Use Limit ID defined in template (Specify LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"glid": {
							Type: schema.TypeInt, Optional: true, Description: "Use global Limit ID (Specify global LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"shared_partition_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
							ValidateFunc:  validation.IntBetween(0, 1),
						},
						"glid_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Use global Limit ID",
							ValidateFunc: validation.IntBetween(1, 1023),
						},
						"lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"lsn_radius_profile": {
							Type: schema.TypeInt, Optional: true, Description: "LSN RADIUS Profile Index",
							ValidateFunc:  validation.IntBetween(1, 16),
						},
						"gtp_rate_limit_policy_v4": {
							Type: schema.TypeString, Optional: true, Description: "GTP Rate Limit Template Name",
							ValidateFunc:  validation.StringLenBetween(1, 128),
						},
						"age": {
							Type: schema.TypeInt, Optional: true, Description: "Specify age in minutes",
							ValidateFunc: validation.IntBetween(1, 2000),
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
							ValidateFunc: axapi.IsIpv6AddressPlen,
						},
						"v6_lid": {
							Type: schema.TypeInt, Optional: true, Description: "Use Limit ID defined in template (Specify LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"v6_glid": {
							Type: schema.TypeInt, Optional: true, Description: "Use global Limit ID (Specify global LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"shared_partition_v6_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
							ValidateFunc:  validation.IntBetween(0, 1),
						},
						"v6_glid_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Use global Limit ID",
							ValidateFunc: validation.IntBetween(1, 1023),
						},
						"v6_lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN Limit ID (LID index)",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"v6_lsn_radius_profile": {
							Type: schema.TypeInt, Optional: true, Description: "LSN RADIUS Profile Index",
							ValidateFunc:  validation.IntBetween(1, 16),
						},
						"gtp_rate_limit_policy_v6": {
							Type: schema.TypeString, Optional: true, Description: "GTP Rate Limit Template Name",
							ValidateFunc:  validation.StringLenBetween(1, 128),
						},
						"v6_age": {
							Type: schema.TypeInt, Optional: true, Description: "Specify age in minutes",
							ValidateFunc: validation.IntBetween(1, 2000),
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Specify name of the class list",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"str_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"str": {
							Type: schema.TypeString, Optional: true, Description: "Specify key string",
							ValidateFunc: validation.StringLenBetween(1, 255),
						},
						"str_lid_dummy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Limit ID defined in template",
							ValidateFunc:  validation.IntBetween(0, 1),
						},
						"str_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LID index",
							ValidateFunc: validation.IntBetween(1, 1023),
						},
						"str_glid_dummy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use global Limit ID",
							ValidateFunc:  validation.IntBetween(0, 1),
						},
						"str_glid": {
							Type: schema.TypeInt, Optional: true, Description: "Global LID index",
							ValidateFunc:  validation.IntBetween(1, 1023),
						},
						"shared_partition_str_glid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a glid from shared partition",
							ValidateFunc:  validation.IntBetween(0, 1),
						},
						"str_glid_shared": {
							Type: schema.TypeInt, Optional: true, Description: "Use global Limit ID",
							ValidateFunc: validation.IntBetween(1, 1023),
						},
						"value_str": {
							Type: schema.TypeString, Optional: true, Description: "Specify value string",
							ValidateFunc: validation.StringLenBetween(1, 639),
						},
					},
				},
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'ac': Make class-list type Aho-Corasick; 'dns': Make class-list type DNS; 'ipv4': Make class-list type IPv4; 'ipv6': Make class-list type IPv6; 'string': Make class-list type String; 'string-case-insensitive': Make class-list type String-case-insensitive. Case insensitive is applied to key string;",
				ValidateFunc: validation.StringInSlice([]string{"ac", "dns", "ipv4", "ipv6", "string", "string-case-insensitive"}, false),
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
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
		obj := edpt.ClassList{}
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
		obj := edpt.ClassList{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointClassList(d *schema.ResourceData) edpt.ClassList {
	var ret edpt.ClassList
	count := 0
	count = d.Get("ac_list.#").(int)
	ret.Inst.AcList = make([]edpt.ClassListAcList, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.ClassListAcList
		prefix := fmt.Sprintf("ac_list.%d.", i)
		obj.AcMatchType = d.Get(prefix + "ac_match_type").(string)
		obj.AcKeyString = d.Get(prefix + "ac_key_string").(string)
		obj.AcValue = d.Get(prefix + "ac_value").(string)
		obj.GtpRateLimitPolicyStr = d.Get(prefix + "gtp_rate_limit_policy_str").(string)
		ret.Inst.AcList = append(ret.Inst.AcList, obj)
	}
	count = d.Get("dns.#").(int)
	ret.Inst.Dns = make([]edpt.ClassListDns, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.ClassListDns
		prefix := fmt.Sprintf("dns.%d.", i)
		obj.DnsMatchType = d.Get(prefix + "dns_match_type").(string)
		obj.DnsMatchString = d.Get(prefix + "dns_match_string").(string)
		obj.DnsLid = d.Get(prefix + "dns_lid").(int)
		obj.DnsGlid = d.Get(prefix + "dns_glid").(int)
		obj.SharedPartitionDnsGlid = d.Get(prefix + "shared_partition_dns_glid").(int)
		obj.DnsGlidShared = d.Get(prefix + "dns_glid_shared").(int)
		ret.Inst.Dns = append(ret.Inst.Dns, obj)
	}
	ret.Inst.File = d.Get("file").(int)
	count = d.Get("ipv4_list.#").(int)
	ret.Inst.Ipv4List = make([]edpt.ClassListIpv4List, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.ClassListIpv4List
		prefix := fmt.Sprintf("ipv4_list.%d.", i)
		obj.Ipv4addr = d.Get(prefix + "ipv4addr").(string)
		obj.Lid = d.Get(prefix + "lid").(int)
		obj.Glid = d.Get(prefix + "glid").(int)
		obj.SharedPartitionGlid = d.Get(prefix + "shared_partition_glid").(int)
		obj.GlidShared = d.Get(prefix + "glid_shared").(int)
		obj.LsnLid = d.Get(prefix + "lsn_lid").(int)
		obj.LsnRadiusProfile = d.Get(prefix + "lsn_radius_profile").(int)
		obj.GtpRateLimitPolicyV4 = d.Get(prefix + "gtp_rate_limit_policy_v4").(string)
		obj.Age = d.Get(prefix + "age").(int)
		ret.Inst.Ipv4List = append(ret.Inst.Ipv4List, obj)
	}
	count = d.Get("ipv6_list.#").(int)
	ret.Inst.Ipv6List = make([]edpt.ClassListIpv6List, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.ClassListIpv6List
		prefix := fmt.Sprintf("ipv6_list.%d.", i)
		obj.Ipv6Addr = d.Get(prefix + "ipv6_addr").(string)
		obj.V6Lid = d.Get(prefix + "v6_lid").(int)
		obj.V6Glid = d.Get(prefix + "v6_glid").(int)
		obj.SharedPartitionV6Glid = d.Get(prefix + "shared_partition_v6_glid").(int)
		obj.V6GlidShared = d.Get(prefix + "v6_glid_shared").(int)
		obj.V6LsnLid = d.Get(prefix + "v6_lsn_lid").(int)
		obj.V6LsnRadiusProfile = d.Get(prefix + "v6_lsn_radius_profile").(int)
		obj.GtpRateLimitPolicyV6 = d.Get(prefix + "gtp_rate_limit_policy_v6").(string)
		obj.V6Age = d.Get(prefix + "v6_age").(int)
		ret.Inst.Ipv6List = append(ret.Inst.Ipv6List, obj)
	}
	ret.Inst.Name = d.Get("name").(string)
	count = d.Get("str_list.#").(int)
	ret.Inst.StrList = make([]edpt.ClassListStrList, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.ClassListStrList
		prefix := fmt.Sprintf("str_list.%d.", i)
		obj.Str = d.Get(prefix + "str").(string)
		obj.StrLidDummy = d.Get(prefix + "str_lid_dummy").(int)
		obj.StrLid = d.Get(prefix + "str_lid").(int)
		obj.StrGlidDummy = d.Get(prefix + "str_glid_dummy").(int)
		obj.StrGlid = d.Get(prefix + "str_glid").(int)
		obj.SharedPartitionStrGlid = d.Get(prefix + "shared_partition_str_glid").(int)
		obj.StrGlidShared = d.Get(prefix + "str_glid_shared").(int)
		obj.ValueStr = d.Get(prefix + "value_str").(string)
		ret.Inst.StrList = append(ret.Inst.StrList, obj)
	}
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
