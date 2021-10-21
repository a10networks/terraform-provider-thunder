package thunder

//Thunder resource IpAccessList

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceIpAccessList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpAccessListCreate,
		UpdateContext: resourceIpAccessListUpdate,
		ReadContext:   resourceIpAccessListRead,
		DeleteContext: resourceIpAccessListDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"remark": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"icmp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"udp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_obj_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"icmp_type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"any_type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"special_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"any_code": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"icmp_code": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"special_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_any": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"src_host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_subnet": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_object_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_eq": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"src_gt": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"src_lt": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"src_range": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"src_port_end": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dst_any": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dst_host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_subnet": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_object_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_eq": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dst_gt": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dst_lt": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dst_range": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dst_port_end": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fragments": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ethernet": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"trunk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dscp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"established": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"transparent_session_only": {
							Type:        schema.TypeInt,
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

func resourceIpAccessListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating IpAccessList (Inside resourceIpAccessListCreate) ")
		name1 := d.Get("name").(string)
		data := dataToIpAccessList(d)
		logger.Println("[INFO] received formatted data from method data to IpAccessList --")
		d.SetId(name1)
		err := go_thunder.PostIpAccessList(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpAccessListRead(ctx, d, meta)

	}
	return diags
}

func resourceIpAccessListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpAccessList (Inside resourceIpAccessListRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetIpAccessList(client.Token, name1, client.Host)
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

func resourceIpAccessListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying IpAccessList   (Inside resourceIpAccessListUpdate) ")
		data := dataToIpAccessList(d)
		logger.Println("[INFO] received formatted data from method data to IpAccessList ")
		err := go_thunder.PutIpAccessList(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpAccessListRead(ctx, d, meta)

	}
	return diags
}

func resourceIpAccessListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceIpAccessListDelete) " + name1)
		err := go_thunder.DeleteIpAccessList(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToIpAccessList(d *schema.ResourceData) go_thunder.IpAccessList {
	var vc go_thunder.IpAccessList
	var c go_thunder.IPAccessListInstance
	c.IPAccessListInstanceName = d.Get("name").(string)

	IPAccessListInstanceRulesCount := d.Get("rules.#").(int)
	c.IPAccessListInstanceRulesSeqNum = make([]go_thunder.IPAccessListInstanceRules, 0, IPAccessListInstanceRulesCount)

	for i := 0; i < IPAccessListInstanceRulesCount; i++ {
		var obj1 go_thunder.IPAccessListInstanceRules
		prefix1 := fmt.Sprintf("rules.%d.", i)
		obj1.IPAccessListInstanceRulesSeqNum = d.Get(prefix1 + "seq_num").(int)
		obj1.IPAccessListInstanceRulesAction = d.Get(prefix1 + "action").(string)
		obj1.IPAccessListInstanceRulesRemark = d.Get(prefix1 + "remark").(string)
		obj1.IPAccessListInstanceRulesIcmp = d.Get(prefix1 + "icmp").(int)
		obj1.IPAccessListInstanceRulesTCP = d.Get(prefix1 + "tcp").(int)
		obj1.IPAccessListInstanceRulesUDP = d.Get(prefix1 + "udp").(int)
		obj1.IPAccessListInstanceRulesIP = d.Get(prefix1 + "ip").(int)
		obj1.IPAccessListInstanceRulesServiceObjGroup = d.Get(prefix1 + "service_obj_group").(string)
		obj1.IPAccessListInstanceRulesGeoLocation = d.Get(prefix1 + "geo_location").(string)
		obj1.IPAccessListInstanceRulesIcmpType = d.Get(prefix1 + "icmp_type").(int)
		obj1.IPAccessListInstanceRulesAnyType = d.Get(prefix1 + "any_type").(int)
		obj1.IPAccessListInstanceRulesSpecialType = d.Get(prefix1 + "special_type").(string)
		obj1.IPAccessListInstanceRulesAnyCode = d.Get(prefix1 + "any_code").(int)
		obj1.IPAccessListInstanceRulesIcmpCode = d.Get(prefix1 + "icmp_code").(int)
		obj1.IPAccessListInstanceRulesSpecialCode = d.Get(prefix1 + "special_code").(string)
		obj1.IPAccessListInstanceRulesSrcAny = d.Get(prefix1 + "src_any").(int)
		obj1.IPAccessListInstanceRulesSrcHost = d.Get(prefix1 + "src_host").(string)
		obj1.IPAccessListInstanceRulesSrcSubnet = d.Get(prefix1 + "src_subnet").(string)
		obj1.IPAccessListInstanceRulesSrcMask = d.Get(prefix1 + "src_mask").(string)
		obj1.IPAccessListInstanceRulesSrcObjectGroup = d.Get(prefix1 + "src_object_group").(string)
		obj1.IPAccessListInstanceRulesSrcEq = d.Get(prefix1 + "src_eq").(int)
		obj1.IPAccessListInstanceRulesSrcGt = d.Get(prefix1 + "src_gt").(int)
		obj1.IPAccessListInstanceRulesSrcLt = d.Get(prefix1 + "src_lt").(int)
		obj1.IPAccessListInstanceRulesSrcRange = d.Get(prefix1 + "src_range").(int)
		obj1.IPAccessListInstanceRulesSrcPortEnd = d.Get(prefix1 + "src_port_end").(int)
		obj1.IPAccessListInstanceRulesDstAny = d.Get(prefix1 + "dst_any").(int)
		obj1.IPAccessListInstanceRulesDstHost = d.Get(prefix1 + "dst_host").(string)
		obj1.IPAccessListInstanceRulesDstSubnet = d.Get(prefix1 + "dst_subnet").(string)
		obj1.IPAccessListInstanceRulesDstMask = d.Get(prefix1 + "dst_mask").(string)
		obj1.IPAccessListInstanceRulesDstObjectGroup = d.Get(prefix1 + "dst_object_group").(string)
		obj1.IPAccessListInstanceRulesDstEq = d.Get(prefix1 + "dst_eq").(int)
		obj1.IPAccessListInstanceRulesDstGt = d.Get(prefix1 + "dst_gt").(int)
		obj1.IPAccessListInstanceRulesDstLt = d.Get(prefix1 + "dst_lt").(int)
		obj1.IPAccessListInstanceRulesDstRange = d.Get(prefix1 + "dst_range").(int)
		obj1.IPAccessListInstanceRulesDstPortEnd = d.Get(prefix1 + "dst_port_end").(int)
		obj1.IPAccessListInstanceRulesFragments = d.Get(prefix1 + "fragments").(int)
		obj1.IPAccessListInstanceRulesVlan = d.Get(prefix1 + "vlan").(int)
		obj1.IPAccessListInstanceRulesEthernet = d.Get(prefix1 + "ethernet").(int)
		obj1.IPAccessListInstanceRulesTrunk = d.Get(prefix1 + "trunk").(int)
		obj1.IPAccessListInstanceRulesDscp = d.Get(prefix1 + "dscp").(int)
		obj1.IPAccessListInstanceRulesEstablished = d.Get(prefix1 + "established").(int)
		obj1.IPAccessListInstanceRulesAclLog = d.Get(prefix1 + "acl_log").(int)
		obj1.IPAccessListInstanceRulesTransparentSessionOnly = d.Get(prefix1 + "transparent_session_only").(int)
		c.IPAccessListInstanceRulesSeqNum = append(c.IPAccessListInstanceRulesSeqNum, obj1)
	}

	c.IPAccessListInstanceUserTag = d.Get("user_tag").(string)

	vc.IPAccessListInstanceName = c
	return vc
}
