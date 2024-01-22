package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_number_src_based_policy`: Configure src-based policy\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyDelete,

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
									"logging": {
										Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
									},
									"ip_proto": {
										Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
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
												"ip_proto": {
													Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
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
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
		},
	}
}
func resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberSrcBasedPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberSrcBasedPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberSrcBasedPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberSrcBasedPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListListClassListOverflowPolicyListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberSrcBasedPolicy(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicy {
	var ret edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicy
	ret.Inst.PolicyClassListList = getSliceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListList(d.Get("policy_class_list_list").([]interface{}))
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
