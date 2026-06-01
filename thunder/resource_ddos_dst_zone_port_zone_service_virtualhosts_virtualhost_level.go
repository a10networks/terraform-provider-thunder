package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_virtualhosts_virtualhost_level`: Virtualhost Policy Level Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelDelete,

		Schema: map[string]*schema.Schema{
			"glid_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
			},
			"level_num": {
				Type: schema.TypeString, Required: true, Description: "'0': Default policy level;",
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
						"ssl_l4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
						},
						"tcp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
						},
					},
				},
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"vhost": {
				Type: schema.TypeString, Required: true, Description: "Vhost",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Tcp = in["tcp"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel {
	var ret edpt.DdosDstZonePortZoneServiceVirtualhostsVirtualhostLevel
	ret.Inst.GlidAction = d.Get("glid_action").(string)
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.SrcDefaultGlid = d.Get("src_default_glid").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZonePortZoneServiceVirtualhostsVirtualhostLevelZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Vhost = d.Get("vhost").(string)
	return ret
}
