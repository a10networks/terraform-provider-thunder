package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeVirtualhosts() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_range_virtualhosts`: Configure virtualhost based mitigation for SSL services\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortRangeVirtualhostsCreate,
		UpdateContext: resourceDdosDstZonePortRangeVirtualhostsUpdate,
		ReadContext:   resourceDdosDstZonePortRangeVirtualhostsRead,
		DeleteContext: resourceDdosDstZonePortRangeVirtualhostsDelete,

		Schema: map[string]*schema.Schema{
			"source_tracking_all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enables creation of source entries for all virtualhosts",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vhosts_config": {
				Type: schema.TypeString, Required: true, Description: "'configuration': configure virtualhost based mitigation for ssl services;",
			},
			"virtualhost_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vhost": {
							Type: schema.TypeString, Required: true, Description: "name for virtualhost",
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
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}
func resourceDdosDstZonePortRangeVirtualhostsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhosts(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeVirtualhostsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortRangeVirtualhostsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhosts(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortRangeVirtualhostsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortRangeVirtualhostsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhosts(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortRangeVirtualhostsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhosts(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostList(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostList
		oi.Vhost = in["vhost"].(string)
		oi.Servername = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostListServername(in["servername"].([]interface{}))
		oi.ServernameList = in["servername_list"].(string)
		oi.ServernameMatchAny = in["servername_match_any"].(int)
		oi.SourceTracking = in["source_tracking"].(string)
		oi.GlidCfg = getObjectDdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg(in["glid_cfg"].([]interface{}))
		oi.Deny = in["deny"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.LevelList = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostListLevelList(in["level_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostListServername(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListServername {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListServername, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListServername
		oi.MatchType = in["match_type"].(string)
		oi.HostMatchString = in["host_match_string"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg(d []interface{}) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Glid = in["glid"].(string)
		ret.GlidAction = in["glid_action"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostListLevelList(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListLevelList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListLevelList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListLevelList
		oi.LevelNum = in["level_num"].(string)
		oi.SrcDefaultGlid = in["src_default_glid"].(string)
		oi.GlidAction = in["glid_action"].(string)
		oi.ZoneTemplate = getObjectDdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate(in["zone_template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate(d []interface{}) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Tcp = in["tcp"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeVirtualhosts(d *schema.ResourceData) edpt.DdosDstZonePortRangeVirtualhosts {
	var ret edpt.DdosDstZonePortRangeVirtualhosts
	ret.Inst.SourceTrackingAll = d.Get("source_tracking_all").(int)
	//omit uuid
	ret.Inst.VhostsConfig = d.Get("vhosts_config").(string)
	ret.Inst.VirtualhostList = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostList(d.Get("virtualhost_list").([]interface{}))
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(string)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
