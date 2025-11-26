package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_range_virtualhosts_virtualhost_level`: Virtualhost Policy Level Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelCreate,
		UpdateContext: resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelUpdate,
		ReadContext:   resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelRead,
		DeleteContext: resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelDelete,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"vhost": {
				Type: schema.TypeString, Required: true, Description: "Vhost",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
		},
	}
}
func resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhostLevel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhostLevel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhostLevel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostLevelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhostLevel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZonePortRangeVirtualhostsVirtualhostLevelZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevelZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevelZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Tcp = in["tcp"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhostLevel(d *schema.ResourceData) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevel {
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevel
	ret.Inst.GlidAction = d.Get("glid_action").(string)
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.SrcDefaultGlid = d.Get("src_default_glid").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosDstZonePortRangeVirtualhostsVirtualhostLevelZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Vhost = d.Get("vhost").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(string)
	return ret
}
