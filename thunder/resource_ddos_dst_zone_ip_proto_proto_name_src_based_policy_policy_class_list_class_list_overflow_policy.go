package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_name_src_based_policy_policy_class_list_class_list_overflow_policy`: Configure dynamic entry count overflow policy for class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyDelete,

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
						"icmp_v4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
						},
						"icmp_v6": {
							Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
						},
						"ip_proto": {
							Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
						},
						"encap": {
							Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
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
		},
	}
}
func resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy {
	var ret edpt.DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicy
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.LogEnable = d.Get("log_enable").(int)
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListClassListOverflowPolicyZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	return ret
}
