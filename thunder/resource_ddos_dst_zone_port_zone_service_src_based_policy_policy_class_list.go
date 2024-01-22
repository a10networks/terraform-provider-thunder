package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list`: Configure class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
			},
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
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
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"glid_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Glid = in["glid"].(string)
		oi.Action = in["action"].(string)
		oi.LogEnable = in["log_enable"].(int)
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.ZoneTemplate = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyListZoneTemplate
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

func getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListSamplingEnable(d []interface{}) []edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListZoneTemplate
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

func dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList {
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassList
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.ClassListOverflowPolicyList = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyList(d.Get("class_list_overflow_policy_list").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.GlidAction = d.Get("glid_action").(string)
	ret.Inst.LogEnable = d.Get("log_enable").(int)
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.MaxDynamicEntryCount = d.Get("max_dynamic_entry_count").(int)
	ret.Inst.SamplingEnable = getSliceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	return ret
}
