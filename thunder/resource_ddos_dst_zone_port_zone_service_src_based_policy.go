package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceSrcBasedPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_src_based_policy`: Configure src-based policy\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceSrcBasedPolicyRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyDelete,

		Schema: map[string]*schema.Schema{
			"policy_class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_name": {
							Type: schema.TypeString, Required: true, Description: "Class-list name",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
						},
						"log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
						},
						"log_periodic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
						},
						"max_dynamic_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic source zone service entry allowed for this class-list",
						},
						"zone_template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"quic": {
										Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
									},
									"dns": {
										Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
									},
									"http": {
										Type: schema.TypeString, Optional: true, Description: "DDOS http template",
									},
									"ssl_l4": {
										Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
									},
									"sip": {
										Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
									},
									"tcp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
									},
									"udp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
									},
									"encap": {
										Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
									},
									"logging": {
										Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;",
									},
								},
							},
						},
						"class_list_overflow_policy_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dummy_name": {
										Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy for class-list;",
									},
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
									},
									"log_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
									},
									"log_periodic": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
									},
									"zone_template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"quic": {
													Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
												},
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
												},
												"encap": {
													Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
												},
												"logging": {
													Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the policy",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceSrcBasedPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceSrcBasedPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceSrcBasedPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceSrcBasedPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListList(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicy(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceSrcBasedPolicy {
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicy
	ret.Inst.PolicyClassListList = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListList(d.Get("policy_class_list_list").([]interface{}))
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
