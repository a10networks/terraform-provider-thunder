package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_number_dynamic_entry_overflow_policy`: Configure overflow policy\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyDelete,

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
						"ip_proto": {
							Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
						},
					},
				},
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
func resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy {
	var ret edpt.DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicy
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
