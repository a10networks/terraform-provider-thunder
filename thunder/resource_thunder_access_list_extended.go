package thunder

//Thunder resource AccessListExtended

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccessListExtended() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccessListExtendedCreate,
		UpdateContext: resourceAccessListExtendedUpdate,
		ReadContext:   resourceAccessListExtendedRead,
		DeleteContext: resourceAccessListExtendedDelete,
		Schema: map[string]*schema.Schema{
			"extd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extd_seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"extd_remark": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"extd_action": {
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
		},
	}
}

func resourceAccessListExtendedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating AccessListExtended (Inside resourceAccessListExtendedCreate) ")
		name1 := d.Get("extd").(int)
		data := dataToAccessListExtended(d)
		logger.Println("[INFO] received formatted data from method data to AccessListExtended --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostAccessListExtended(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceAccessListExtendedRead(ctx, d, meta)

	}
	return diags
}

func resourceAccessListExtendedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading AccessListExtended (Inside resourceAccessListExtendedRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetAccessListExtended(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceAccessListExtendedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying AccessListExtended   (Inside resourceAccessListExtendedUpdate) ")
		data := dataToAccessListExtended(d)
		logger.Println("[INFO] received formatted data from method data to AccessListExtended ")
		err := go_thunder.PutAccessListExtended(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceAccessListExtendedRead(ctx, d, meta)

	}
	return diags
}

func resourceAccessListExtendedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceAccessListExtendedDelete) " + name1)
		err := go_thunder.DeleteAccessListExtended(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToAccessListExtended(d *schema.ResourceData) go_thunder.AccessListExtended {
	var vc go_thunder.AccessListExtended
	var c go_thunder.AccessListExtendedInstance
	c.Extd = d.Get("extd").(int)

	RulesCount := d.Get("rules.#").(int)
	c.ExtdSeqNum = make([]go_thunder.AccessListRules, 0, RulesCount)

	for i := 0; i < RulesCount; i++ {
		var obj1 go_thunder.AccessListRules
		prefix1 := fmt.Sprintf("rules.%d.", i)
		obj1.ExtdSeqNum = d.Get(prefix1 + "extd_seq_num").(int)
		obj1.ExtdRemark = d.Get(prefix1 + "extd_remark").(string)
		obj1.ExtdAction = d.Get(prefix1 + "extd_action").(string)
		obj1.Icmp = d.Get(prefix1 + "icmp").(int)
		obj1.TCP = d.Get(prefix1 + "tcp").(int)
		obj1.UDP = d.Get(prefix1 + "udp").(int)
		obj1.IP = d.Get(prefix1 + "ip").(int)
		obj1.ServiceObjGroup = d.Get(prefix1 + "service_obj_group").(string)
		obj1.IcmpType = d.Get(prefix1 + "icmp_type").(int)
		obj1.AnyType = d.Get(prefix1 + "any_type").(int)
		obj1.SpecialType = d.Get(prefix1 + "special_type").(string)
		obj1.AnyCode = d.Get(prefix1 + "any_code").(int)
		obj1.IcmpCode = d.Get(prefix1 + "icmp_code").(int)
		obj1.SpecialCode = d.Get(prefix1 + "special_code").(string)
		obj1.SrcAny = d.Get(prefix1 + "src_any").(int)
		obj1.SrcHost = d.Get(prefix1 + "src_host").(string)
		obj1.SrcSubnet = d.Get(prefix1 + "src_subnet").(string)
		obj1.SrcMask = d.Get(prefix1 + "src_mask").(string)
		obj1.SrcObjectGroup = d.Get(prefix1 + "src_object_group").(string)
		obj1.SrcEq = d.Get(prefix1 + "src_eq").(int)
		obj1.SrcGt = d.Get(prefix1 + "src_gt").(int)
		obj1.SrcLt = d.Get(prefix1 + "src_lt").(int)
		obj1.SrcRange = d.Get(prefix1 + "src_range").(int)
		obj1.SrcPortEnd = d.Get(prefix1 + "src_port_end").(int)
		obj1.DstAny = d.Get(prefix1 + "dst_any").(int)
		obj1.DstHost = d.Get(prefix1 + "dst_host").(string)
		obj1.DstSubnet = d.Get(prefix1 + "dst_subnet").(string)
		obj1.DstMask = d.Get(prefix1 + "dst_mask").(string)
		obj1.DstObjectGroup = d.Get(prefix1 + "dst_object_group").(string)
		obj1.DstEq = d.Get(prefix1 + "dst_eq").(int)
		obj1.DstGt = d.Get(prefix1 + "dst_gt").(int)
		obj1.DstLt = d.Get(prefix1 + "dst_lt").(int)
		obj1.DstRange = d.Get(prefix1 + "dst_range").(int)
		obj1.DstPortEnd = d.Get(prefix1 + "dst_port_end").(int)
		obj1.Fragments = d.Get(prefix1 + "fragments").(int)
		obj1.Vlan = d.Get(prefix1 + "vlan").(int)
		obj1.Ethernet = d.Get(prefix1 + "ethernet").(int)
		obj1.Trunk = d.Get(prefix1 + "trunk").(int)
		obj1.Dscp = d.Get(prefix1 + "dscp").(int)
		obj1.Established = d.Get(prefix1 + "established").(int)
		obj1.AclLog = d.Get(prefix1 + "acl_log").(int)
		obj1.TransparentSessionOnly = d.Get(prefix1 + "transparent_session_only").(int)
		c.ExtdSeqNum = append(c.ExtdSeqNum, obj1)
	}

	vc.Extd = c
	return vc
}
