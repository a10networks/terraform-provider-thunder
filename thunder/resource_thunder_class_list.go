package thunder

//Thunder resource ClassList

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceClassList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClassListCreate,
		UpdateContext: resourceClassListUpdate,
		ReadContext:   resourceClassListRead,
		DeleteContext: resourceClassListDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"file": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ipv4_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"glid_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lsn_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lsn_radius_profile": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gtp_rate_limit_policy_v4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"age": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ipv6_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"v6_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v6_glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_v6_glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v6_glid_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v6_lsn_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v6_lsn_radius_profile": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gtp_rate_limit_policy_v6": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"v6_age": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"dns": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_match_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dns_match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dns_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dns_glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_dns_glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dns_glid_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"str_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"str": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"str_lid_dummy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"str_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"str_glid_dummy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"str_glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_str_glid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"str_glid_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"value_str": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ac_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ac_match_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ac_key_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ac_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"gtp_rate_limit_policy_str": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating ClassList (Inside resourceClassListCreate) ")
		name1 := d.Get("name").(string)
		data := dataToClassList(d)
		logger.Println("[INFO] received formatted data from method data to ClassList --")
		d.SetId(name1)
		err := go_thunder.PostClassList(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceClassListRead(ctx, d, meta)

	}
	return diags
}

func resourceClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading ClassList (Inside resourceClassListRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetClassList(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying ClassList   (Inside resourceClassListUpdate) ")
		data := dataToClassList(d)
		logger.Println("[INFO] received formatted data from method data to ClassList ")
		err := go_thunder.PutClassList(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceClassListRead(ctx, d, meta)

	}
	return diags
}

func resourceClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceClassListDelete) " + name1)
		err := go_thunder.DeleteClassList(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToClassList(d *schema.ResourceData) go_thunder.ClassList {
	var vc go_thunder.ClassList
	var c go_thunder.ClassListInstance
	c.ClassListInstanceName = d.Get("name").(string)
	c.ClassListInstanceType = d.Get("type").(string)
	c.ClassListInstanceFile = d.Get("file").(int)

	ClassListInstanceIpv4ListCount := d.Get("ipv4_list.#").(int)
	c.ClassListInstanceIpv4ListIpv4Addr = make([]go_thunder.ClassListInstanceIpv4List, 0, ClassListInstanceIpv4ListCount)

	for i := 0; i < ClassListInstanceIpv4ListCount; i++ {
		var obj1 go_thunder.ClassListInstanceIpv4List
		prefix1 := fmt.Sprintf("ipv4_list.%d.", i)
		obj1.ClassListInstanceIpv4ListIpv4Addr = d.Get(prefix1 + "ipv4addr").(string)
		obj1.ClassListInstanceIpv4ListLid = d.Get(prefix1 + "lid").(int)
		obj1.ClassListInstanceIpv4ListGlid = d.Get(prefix1 + "glid").(int)
		obj1.ClassListInstanceIpv4ListSharedPartitionGlid = d.Get(prefix1 + "shared_partition_glid").(int)
		obj1.ClassListInstanceIpv4ListGlidShared = d.Get(prefix1 + "glid_shared").(int)
		obj1.ClassListInstanceIpv4ListLsnLid = d.Get(prefix1 + "lsn_lid").(int)
		obj1.ClassListInstanceIpv4ListLsnRadiusProfile = d.Get(prefix1 + "lsn_radius_profile").(int)
		obj1.ClassListInstanceIpv4ListGtpRateLimitPolicyV4 = d.Get(prefix1 + "gtp_rate_limit_policy_v4").(string)
		obj1.ClassListInstanceIpv4ListAge = d.Get(prefix1 + "age").(int)
		c.ClassListInstanceIpv4ListIpv4Addr = append(c.ClassListInstanceIpv4ListIpv4Addr, obj1)
	}

	ClassListInstanceIpv6ListCount := d.Get("ipv6_list.#").(int)
	c.ClassListInstanceIpv6ListIpv6Addr = make([]go_thunder.ClassListInstanceIpv6List, 0, ClassListInstanceIpv6ListCount)

	for i := 0; i < ClassListInstanceIpv6ListCount; i++ {
		var obj2 go_thunder.ClassListInstanceIpv6List
		prefix2 := fmt.Sprintf("ipv6_list.%d.", i)
		obj2.ClassListInstanceIpv6ListIpv6Addr = d.Get(prefix2 + "ipv6_addr").(string)
		obj2.ClassListInstanceIpv6ListV6Lid = d.Get(prefix2 + "v6_lid").(int)
		obj2.ClassListInstanceIpv6ListV6Glid = d.Get(prefix2 + "v6_glid").(int)
		obj2.ClassListInstanceIpv6ListSharedPartitionV6Glid = d.Get(prefix2 + "shared_partition_v6_glid").(int)
		obj2.ClassListInstanceIpv6ListV6GlidShared = d.Get(prefix2 + "v6_glid_shared").(int)
		obj2.ClassListInstanceIpv6ListV6LsnLid = d.Get(prefix2 + "v6_lsn_lid").(int)
		obj2.ClassListInstanceIpv6ListV6LsnRadiusProfile = d.Get(prefix2 + "v6_lsn_radius_profile").(int)
		obj2.ClassListInstanceIpv6ListGtpRateLimitPolicyV6 = d.Get(prefix2 + "gtp_rate_limit_policy_v6").(string)
		obj2.ClassListInstanceIpv6ListV6Age = d.Get(prefix2 + "v6_age").(int)
		c.ClassListInstanceIpv6ListIpv6Addr = append(c.ClassListInstanceIpv6ListIpv6Addr, obj2)
	}

	ClassListInstanceDNSCount := d.Get("dns.#").(int)
	c.ClassListInstanceDNSDNSMatchType = make([]go_thunder.ClassListInstanceDNS, 0, ClassListInstanceDNSCount)

	for i := 0; i < ClassListInstanceDNSCount; i++ {
		var obj3 go_thunder.ClassListInstanceDNS
		prefix3 := fmt.Sprintf("dns.%d.", i)
		obj3.ClassListInstanceDNSDNSMatchType = d.Get(prefix3 + "dns_match_type").(string)
		obj3.ClassListInstanceDNSDNSMatchString = d.Get(prefix3 + "dns_match_string").(string)
		obj3.ClassListInstanceDNSDNSLid = d.Get(prefix3 + "dns_lid").(int)
		obj3.ClassListInstanceDNSDNSGlid = d.Get(prefix3 + "dns_glid").(int)
		obj3.ClassListInstanceDNSSharedPartitionDNSGlid = d.Get(prefix3 + "shared_partition_dns_glid").(int)
		obj3.ClassListInstanceDNSDNSGlidShared = d.Get(prefix3 + "dns_glid_shared").(int)
		c.ClassListInstanceDNSDNSMatchType = append(c.ClassListInstanceDNSDNSMatchType, obj3)
	}

	ClassListInstanceStrListCount := d.Get("str_list.#").(int)
	c.ClassListInstanceStrListStr = make([]go_thunder.ClassListInstanceStrList, 0, ClassListInstanceStrListCount)

	for i := 0; i < ClassListInstanceStrListCount; i++ {
		var obj4 go_thunder.ClassListInstanceStrList
		prefix4 := fmt.Sprintf("str_list.%d.", i)
		obj4.ClassListInstanceStrListStr = d.Get(prefix4 + "str").(string)
		obj4.ClassListInstanceStrListStrLidDummy = d.Get(prefix4 + "str_lid_dummy").(int)
		obj4.ClassListInstanceStrListStrLid = d.Get(prefix4 + "str_lid").(int)
		obj4.ClassListInstanceStrListStrGlidDummy = d.Get(prefix4 + "str_glid_dummy").(int)
		obj4.ClassListInstanceStrListStrGlid = d.Get(prefix4 + "str_glid").(int)
		obj4.ClassListInstanceStrListSharedPartitionStrGlid = d.Get(prefix4 + "shared_partition_str_glid").(int)
		obj4.ClassListInstanceStrListStrGlidShared = d.Get(prefix4 + "str_glid_shared").(int)
		obj4.ClassListInstanceStrListValueStr = d.Get(prefix4 + "value_str").(string)
		c.ClassListInstanceStrListStr = append(c.ClassListInstanceStrListStr, obj4)
	}

	ClassListInstanceAcListCount := d.Get("ac_list.#").(int)
	c.ClassListInstanceAcListAcMatchType = make([]go_thunder.ClassListInstanceAcList, 0, ClassListInstanceAcListCount)

	for i := 0; i < ClassListInstanceAcListCount; i++ {
		var obj5 go_thunder.ClassListInstanceAcList
		prefix5 := fmt.Sprintf("ac_list.%d.", i)
		obj5.ClassListInstanceAcListAcMatchType = d.Get(prefix5 + "ac_match_type").(string)
		obj5.ClassListInstanceAcListAcKeyString = d.Get(prefix5 + "ac_key_string").(string)
		obj5.ClassListInstanceAcListAcValue = d.Get(prefix5 + "ac_value").(string)
		obj5.ClassListInstanceAcListGtpRateLimitPolicyStr = d.Get(prefix5 + "gtp_rate_limit_policy_str").(string)
		c.ClassListInstanceAcListAcMatchType = append(c.ClassListInstanceAcListAcMatchType, obj5)
	}

	c.ClassListInstanceUserTag = d.Get("user_tag").(string)

	vc.ClassListInstanceName = c
	return vc
}
