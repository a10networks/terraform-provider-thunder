package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameManualMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_name_manual_mode`: Enter manual-mode configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNameManualModeCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNameManualModeUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNameManualModeRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNameManualModeDelete,

		Schema: map[string]*schema.Schema{
			"config": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Manual-mode configuration;",
			},
			"glid_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
			},
			"src_default_glid": {
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
func resourceDdosDstZoneIpProtoProtoNameManualModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameManualModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameManualMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameManualModeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameManualModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameManualModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameManualMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNameManualModeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNameManualModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameManualModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameManualMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNameManualModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameManualModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameManualMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneIpProtoProtoNameManualModeZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameManualModeZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameManualModeZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpV4 = in["icmp_v4"].(string)
		ret.IcmpV6 = in["icmp_v6"].(string)
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNameManualMode(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameManualMode {
	var ret edpt.DdosDstZoneIpProtoProtoNameManualMode
	ret.Inst.Config = d.Get("config").(string)
	ret.Inst.GlidAction = d.Get("glid_action").(string)
	ret.Inst.SrcDefaultGlid = d.Get("src_default_glid").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNameManualModeZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
