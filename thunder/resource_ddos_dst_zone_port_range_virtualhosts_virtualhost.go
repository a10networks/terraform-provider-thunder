package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeVirtualhostsVirtualhost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_range_virtualhosts_virtualhost`: Configure mitigation for virtualhost\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortRangeVirtualhostsVirtualhostCreate,
		UpdateContext: resourceDdosDstZonePortRangeVirtualhostsVirtualhostUpdate,
		ReadContext:   resourceDdosDstZonePortRangeVirtualhostsVirtualhostRead,
		DeleteContext: resourceDdosDstZonePortRangeVirtualhostsVirtualhostDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"glid_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'ignore': Do nothing for glid exceed;",
						},
					},
				},
			},
			"level_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level_num": {
							Type: schema.TypeString, Required: true, Description: "'0': Default policy level;",
						},
						"src_default_glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"glid_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"servername": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': match servername extension when contains this string; 'ends-with': match servername extension when ends with this string; 'equals': match servername extension when equals this string; 'starts-with': match servername extension when starts with this string;",
						},
						"host_match_string": {
							Type: schema.TypeString, Optional: true, Description: "SNI String",
						},
					},
				},
			},
			"servername_list": {
				Type: schema.TypeString, Optional: true, Description: "Class List to match servername (Class List Name)",
			},
			"servername_match_any": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match when there is no SNI or other servernames are not matched",
			},
			"source_tracking": {
				Type: schema.TypeString, Optional: true, Default: "follow", Description: "'follow': enable creation of source entries when source-tracking-all is enabled (default); 'enable': enable creation of source entries on this virtualhost; 'disable': disable creation of source entries on this virtualhost;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vhost": {
				Type: schema.TypeString, Required: true, Description: "name for virtualhost",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
		},
	}
}
func resourceDdosDstZonePortRangeVirtualhostsVirtualhostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeVirtualhostsVirtualhostRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortRangeVirtualhostsVirtualhostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeVirtualhostsVirtualhostRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortRangeVirtualhostsVirtualhostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortRangeVirtualhostsVirtualhostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZonePortRangeVirtualhostsVirtualhostGlidCfg(d []interface{}) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostLevelList(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortRangeVirtualhostsVirtualhostLevelListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeVirtualhostsVirtualhostLevelListZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Tcp = in["tcp"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostServername(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostServername {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostServername, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostServername
		oi.MatchType = in["match_type"].(string)
		oi.HostMatchString = in["host_match_string"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhost(d *schema.ResourceData) edpt.DdosDstZonePortRangeVirtualhostsVirtualhost {
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhost
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.GlidCfg = getObjectDdosDstZonePortRangeVirtualhostsVirtualhostGlidCfg(d.Get("glid_cfg").([]interface{}))
	ret.Inst.LevelList = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostLevelList(d.Get("level_list").([]interface{}))
	ret.Inst.Servername = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostServername(d.Get("servername").([]interface{}))
	ret.Inst.ServernameList = d.Get("servername_list").(string)
	ret.Inst.ServernameMatchAny = d.Get("servername_match_any").(int)
	ret.Inst.SourceTracking = d.Get("source_tracking").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vhost = d.Get("vhost").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(string)
	return ret
}
