package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_name_dynamic_entry_overflow_policy`: Configure overflow policy\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;",
			},
			"dummy_name": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Configure overflow policy;",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}
func resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy {
	var ret edpt.DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicy
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
