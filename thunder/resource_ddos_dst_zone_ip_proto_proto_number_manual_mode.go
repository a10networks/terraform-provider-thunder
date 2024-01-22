package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberManualMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_number_manual_mode`: Enter manual-mode configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNumberManualModeCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNumberManualModeUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNumberManualModeRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNumberManualModeDelete,

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
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
		},
	}
}
func resourceDdosDstZoneIpProtoProtoNumberManualModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberManualModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberManualMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberManualModeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberManualModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberManualModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberManualMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberManualModeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNumberManualModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberManualModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberManualMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberManualModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberManualModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberManualMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneIpProtoProtoNumberManualModeZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberManualModeZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberManualModeZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberManualMode(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberManualMode {
	var ret edpt.DdosDstZoneIpProtoProtoNumberManualMode
	ret.Inst.Config = d.Get("config").(string)
	ret.Inst.GlidAction = d.Get("glid_action").(string)
	ret.Inst.SrcDefaultGlid = d.Get("src_default_glid").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberManualModeZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
