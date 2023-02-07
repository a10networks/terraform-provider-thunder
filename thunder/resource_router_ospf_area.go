package thunder

import (
	"context"
	"fmt"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfArea() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ospf_area`: OSPF area parameters\n\n",
		CreateContext: resourceRouterOspfAreaCreate,
		UpdateContext: resourceRouterOspfAreaUpdate,
		ReadContext:   resourceRouterOspfAreaRead,
		DeleteContext: resourceRouterOspfAreaDelete,
		Schema: map[string]*schema.Schema{
			"process_id": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "OSPF process ID",
				ValidateFunc: validation.IntBetween(0, 65535),
			},
			"area_ipv4": {
				Type: schema.TypeString, Optional: true, ForceNew: true, Description: "OSPF area ID in IP address format",
				ValidateFunc: validation.IsIPv4Address,	ExactlyOneOf: []string{"area_ipv4", "area_num"},
			},
			"area_num": {
				Type: schema.TypeInt, Optional: true, ForceNew: true, Description: "OSPF area ID as a decimal value",
				ValidateFunc: validation.IntBetween(0, 4294967295), ExactlyOneOf: []string{"area_ipv4", "area_num"},
			},
			"auth_cfg": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"message_digest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use message-digest authentication",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			"default_cost": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)",
				ValidateFunc: validation.IntBetween(0, 16777215),
			},
			"filter_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter networks between OSPF areas",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Filter networks by access-list (Name of an access-list)",
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"acl_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Filter networks sent to this area; 'out': Filter networks sent from this area;",
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
						},
						"plist_name": {
							Type: schema.TypeString, Optional: true, Description: "Filter networks by prefix-list (Name of an IP prefix-list)",
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
						"plist_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Filter networks sent to this area; 'out': Filter networks sent from this area;",
							ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),
						},
					},
				},
			},
			"nssa_cfg": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nssa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a NSSA area",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"no_redistribution": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No redistribution into this NSSA area",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"no_summary": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not send summary LSA into NSSA",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"translator_role": {
							Type: schema.TypeString, Optional: true, Default: "candidate", Description: "'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;",
							ValidateFunc: validation.StringInSlice([]string{"always", "candidate", "never"}, false),
						},
						"default_information_originate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate Type 7 default into NSSA area",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "OSPF default metric (OSPF metric)",
							ValidateFunc: validation.IntBetween(0, 16777214),
						},
						"metric_type": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "OSPF metric type (OSPF metric type for default routes)",
							ValidateFunc: validation.IntBetween(1, 2),
						},
					},
				},
			},
			"range_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_range_prefix": {
							Type: schema.TypeString, Optional: true, Description: "Area range for IPv4 prefix",
							ValidateFunc: validation.IsCIDR,
						},
						"option": {
							Type: schema.TypeString, Optional: true, Default: "advertise", Description: "'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;",
							ValidateFunc: validation.StringInSlice([]string{"advertise", "not-advertise"}, false),
						},
					},
				},
			},
			"shortcut": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;",
				ValidateFunc: validation.StringInSlice([]string{"default", "disable", "enable"}, false),
			},
			"stub_cfg": {
				Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stub": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure OSPF area as stub",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"no_summary": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not inject inter-area routes into area",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_link_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_link_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "ID (IP addr) associated with virtual link neighbor",
							ValidateFunc: validation.IsIPv4Address,
						},
						"bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"hello_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Hello packet interval (Seconds)",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"dead_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Dead router detection time (Seconds)",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"retransmit_interval": {
							Type: schema.TypeInt, Optional: true, Description: "LSA retransmit interval (Seconds)",
							ValidateFunc: validation.IntBetween(1, 3600),
						},
						"transmit_delay": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "LSA transmission delay (Seconds)",
							ValidateFunc: validation.IntBetween(1, 3600),
						},
						"virtual_link_authentication": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"virtual_link_auth_type": {
							Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use null authentication;",
							ValidateFunc: validation.StringInSlice([]string{"message-digest", "null"}, false),
						},
						"authentication_key": {
							Type: schema.TypeString, Optional: true, Description: "Set authentication key (Authentication key (8 chars))",
							ValidateFunc: validation.StringLenBetween(1, 8),
						},
						"message_digest_key": {
							Type: schema.TypeInt, Optional: true, Description: "Set message digest key (Key ID)",
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"md5": {
							Type: schema.TypeString, Optional: true, Description: "Use MD5 algorithm (Authentication key (16 chars))",
							ValidateFunc: validation.StringLenBetween(1, 16),
						},
					},
				},
			},
		},
	}
}

func resourceRouterOspfAreaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfAreaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfArea(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfAreaRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterOspfAreaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfAreaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfArea(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterOspfAreaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfAreaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfArea(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfAreaRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterOspfAreaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfAreaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfArea(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterOspfArea(d *schema.ResourceData) edpt.RouterOspfArea {
	var ret edpt.RouterOspfArea
	count := 0
	ret.Inst.ProcessId = d.Get("process_id").(int)
	val, ok := d.GetOk("area_ipv4")
	if ok {
		ret.Inst.AreaIpv4 = val.(string)
		ret.Inst.UseAreaIpv4 = true
		ret.Inst.AreaNum = -1
	} else {
		ret.Inst.AreaNum = d.Get("area_num").(int)
		ret.Inst.UseAreaIpv4 = false
	}
	ret.Inst.AuthCfg.Authentication = d.Get("auth_cfg.0.authentication").(int)
	ret.Inst.AuthCfg.MessageDigest = d.Get("auth_cfg.0.message_digest").(int)
	ret.Inst.DefaultCost = d.Get("default_cost").(int)
	count = d.Get("filter_lists.#").(int)
	ret.Inst.FilterLists = make([]edpt.RouterOspfAreaFilterLists, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.RouterOspfAreaFilterLists
		prefix := fmt.Sprintf("filter_lists.%d.", i)
		obj.FilterList = d.Get(prefix + "filter_list").(int)
		obj.AclName = d.Get(prefix + "acl_name").(string)
		obj.AclDirection = d.Get(prefix + "acl_direction").(string)
		obj.PlistName = d.Get(prefix + "plist_name").(string)
		obj.PlistDirection = d.Get(prefix + "plist_direction").(string)
		ret.Inst.FilterLists = append(ret.Inst.FilterLists, obj)
	}
	ret.Inst.NssaCfg.Nssa = d.Get("nssa_cfg.0.nssa").(int)
	ret.Inst.NssaCfg.NoRedistribution = d.Get("nssa_cfg.0.no_redistribution").(int)
	ret.Inst.NssaCfg.NoSummary = d.Get("nssa_cfg.0.no_summary").(int)
	ret.Inst.NssaCfg.TranslatorRole = d.Get("nssa_cfg.0.translator_role").(string)
	ret.Inst.NssaCfg.DefaultInformationOriginate = d.Get("nssa_cfg.0.default_information_originate").(int)
	ret.Inst.NssaCfg.Metric = d.Get("nssa_cfg.0.metric").(int)
	ret.Inst.NssaCfg.MetricType = d.Get("nssa_cfg.0.metric_type").(int)
	count = d.Get("range_list.#").(int)
	ret.Inst.RangeList = make([]edpt.RouterOspfAreaRangeList, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.RouterOspfAreaRangeList
		prefix := fmt.Sprintf("range_list.%d.", i)
		obj.AreaRangePrefix = d.Get(prefix + "area_range_prefix").(string)
		obj.Option = d.Get(prefix + "option").(string)
		ret.Inst.RangeList = append(ret.Inst.RangeList, obj)
	}
	ret.Inst.Shortcut = d.Get("shortcut").(string)
	ret.Inst.StubCfg.Stub = d.Get("stub_cfg.0.stub").(int)
	ret.Inst.StubCfg.NoSummary = d.Get("stub_cfg.0.no_summary").(int)
	//omit uuid
	count = d.Get("virtual_link_list.#").(int)
	ret.Inst.VirtualLinkList = make([]edpt.RouterOspfAreaVirtualLinkList, 0, count)
	for i := 0; i < count; i++ {
		var obj edpt.RouterOspfAreaVirtualLinkList
		prefix := fmt.Sprintf("virtual_link_list.%d.", i)
		obj.VirtualLinkIpAddr = d.Get(prefix + "virtual_link_ip_addr").(string)
		obj.Bfd = d.Get(prefix + "bfd").(int)
		obj.HelloInterval = d.Get(prefix + "hello_interval").(int)
		obj.DeadInterval = d.Get(prefix + "dead_interval").(int)
		obj.RetransmitInterval = d.Get(prefix + "retransmit_interval").(int)
		obj.TransmitDelay = d.Get(prefix + "transmit_delay").(int)
		obj.VirtualLinkAuthentication = d.Get(prefix + "virtual_link_authentication").(int)
		obj.VirtualLinkAuthType = d.Get(prefix + "virtual_link_auth_type").(string)
		obj.AuthenticationKey = d.Get(prefix + "authentication_key").(string)
		obj.MessageDigestKey = d.Get(prefix + "message_digest_key").(int)
		obj.Md5 = d.Get(prefix + "md5").(string)
		ret.Inst.VirtualLinkList = append(ret.Inst.VirtualLinkList, obj)
	}
	return ret
}
