package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfArea() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ospf_area`: OSPF area parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterOspfAreaCreate,
		UpdateContext: resourceRouterOspfAreaUpdate,
		ReadContext:   resourceRouterOspfAreaRead,
		DeleteContext: resourceRouterOspfAreaDelete,

		Schema: map[string]*schema.Schema{
			"area_ipv4": {
				Type: schema.TypeString, Required: true, Description: "OSPF area ID in IP address format",
			},
			"area_num": {
				Type: schema.TypeInt, Required: true, Description: "OSPF area ID as a decimal value",
			},
			"auth_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
						},
						"message_digest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use message-digest authentication",
						},
					},
				},
			},
			"default_cost": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)",
			},
			"filter_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter networks between OSPF areas",
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Filter networks by access-list (Name of an access-list)",
						},
						"acl_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Filter networks sent to this area; 'out': Filter networks sent from this area;",
						},
						"plist_name": {
							Type: schema.TypeString, Optional: true, Description: "Filter networks by prefix-list (Name of an IP prefix-list)",
						},
						"plist_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Filter networks sent to this area; 'out': Filter networks sent from this area;",
						},
					},
				},
			},
			"nssa_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nssa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a NSSA area",
						},
						"no_redistribution": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No redistribution into this NSSA area",
						},
						"no_summary": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not send summary LSA into NSSA",
						},
						"translator_role": {
							Type: schema.TypeString, Optional: true, Default: "candidate", Description: "'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;",
						},
						"default_information_originate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate Type 7 default into NSSA area",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "OSPF default metric (OSPF metric)",
						},
						"metric_type": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "OSPF metric type (OSPF metric type for default routes)",
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
						},
						"option": {
							Type: schema.TypeString, Optional: true, Default: "advertise", Description: "'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;",
						},
					},
				},
			},
			"shortcut": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;",
			},
			"stub_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stub": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure OSPF area as stub",
						},
						"no_summary": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not inject inter-area routes into area",
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
						},
						"bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
						},
						"hello_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Hello packet interval (Seconds)",
						},
						"dead_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Dead router detection time (Seconds)",
						},
						"retransmit_interval": {
							Type: schema.TypeInt, Optional: true, Description: "LSA retransmit interval (Seconds)",
						},
						"transmit_delay": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "LSA transmission delay (Seconds)",
						},
						"virtual_link_authentication": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
						},
						"virtual_link_auth_type": {
							Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use null authentication;",
						},
						"authentication_key": {
							Type: schema.TypeString, Optional: true, Description: "Set authentication key (Authentication key (8 chars))",
						},
						"message_digest_key": {
							Type: schema.TypeInt, Optional: true, Description: "Set message digest key (Key ID)",
						},
						"md5": {
							Type: schema.TypeString, Optional: true, Description: "Use MD5 algorithm (Authentication key (16 chars))",
						},
					},
				},
			},
			"process_id": {
				Type: schema.TypeString, Required: true, Description: "ProcessId",
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

func getObjectRouterOspfAreaAuthCfg(d []interface{}) edpt.RouterOspfAreaAuthCfg {

	count1 := len(d)
	var ret edpt.RouterOspfAreaAuthCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.MessageDigest = in["message_digest"].(int)
	}
	return ret
}

func getSliceRouterOspfAreaFilterLists(d []interface{}) []edpt.RouterOspfAreaFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterOspfAreaFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfAreaFilterLists
		oi.FilterList = in["filter_list"].(int)
		oi.AclName = in["acl_name"].(string)
		oi.AclDirection = in["acl_direction"].(string)
		oi.PlistName = in["plist_name"].(string)
		oi.PlistDirection = in["plist_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfAreaNssaCfg(d []interface{}) edpt.RouterOspfAreaNssaCfg {

	count1 := len(d)
	var ret edpt.RouterOspfAreaNssaCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nssa = in["nssa"].(int)
		ret.NoRedistribution = in["no_redistribution"].(int)
		ret.NoSummary = in["no_summary"].(int)
		ret.TranslatorRole = in["translator_role"].(string)
		ret.DefaultInformationOriginate = in["default_information_originate"].(int)
		ret.Metric = in["metric"].(int)
		ret.MetricType = in["metric_type"].(int)
	}
	return ret
}

func getSliceRouterOspfAreaRangeList(d []interface{}) []edpt.RouterOspfAreaRangeList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfAreaRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfAreaRangeList
		oi.AreaRangePrefix = in["area_range_prefix"].(string)
		oi.Option = in["option"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfAreaStubCfg(d []interface{}) edpt.RouterOspfAreaStubCfg {

	count1 := len(d)
	var ret edpt.RouterOspfAreaStubCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stub = in["stub"].(int)
		ret.NoSummary = in["no_summary"].(int)
	}
	return ret
}

func getSliceRouterOspfAreaVirtualLinkList(d []interface{}) []edpt.RouterOspfAreaVirtualLinkList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfAreaVirtualLinkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfAreaVirtualLinkList
		oi.VirtualLinkIpAddr = in["virtual_link_ip_addr"].(string)
		oi.Bfd = in["bfd"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.VirtualLinkAuthentication = in["virtual_link_authentication"].(int)
		oi.VirtualLinkAuthType = in["virtual_link_auth_type"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = in["md5"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterOspfArea(d *schema.ResourceData) edpt.RouterOspfArea {
	var ret edpt.RouterOspfArea
	ret.Inst.AreaIpv4 = d.Get("area_ipv4").(string)
	ret.Inst.AreaNum = d.Get("area_num").(int)
	ret.Inst.AuthCfg = getObjectRouterOspfAreaAuthCfg(d.Get("auth_cfg").([]interface{}))
	ret.Inst.DefaultCost = d.Get("default_cost").(int)
	ret.Inst.FilterLists = getSliceRouterOspfAreaFilterLists(d.Get("filter_lists").([]interface{}))
	ret.Inst.NssaCfg = getObjectRouterOspfAreaNssaCfg(d.Get("nssa_cfg").([]interface{}))
	ret.Inst.RangeList = getSliceRouterOspfAreaRangeList(d.Get("range_list").([]interface{}))
	ret.Inst.Shortcut = d.Get("shortcut").(string)
	ret.Inst.StubCfg = getObjectRouterOspfAreaStubCfg(d.Get("stub_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VirtualLinkList = getSliceRouterOspfAreaVirtualLinkList(d.Get("virtual_link_list").([]interface{}))
	ret.Inst.ProcessId = d.Get("process_id").(string)
	return ret
}
