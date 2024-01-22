package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_class_list_overflow_policy`: Configure dynamic entry count overflow policy for class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
			},
			"dummy_name": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy for class-list;",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"log_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable log periodic",
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
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "ClassListName",
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
func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate
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

func dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy {
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicy
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.LogEnable = d.Get("log_enable").(int)
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	return ret
}
