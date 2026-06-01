package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneWebGui() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_web_gui`: Configure web settings for onbox GUI\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneWebGuiCreate,
		UpdateContext: resourceDdosDstZoneWebGuiUpdate,
		ReadContext:   resourceDdosDstZoneWebGuiRead,
		DeleteContext: resourceDdosDstZoneWebGuiDelete,

		Schema: map[string]*schema.Schema{
			"activated_after_learning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Activate it after learning",
			},
			"create_time": {
				Type: schema.TypeString, Optional: true, Description: "Configure create time",
			},
			"learning": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"duration": {
							Type: schema.TypeString, Optional: true, Default: "6hour", Description: "'1minute': 1 minute; '6hour': 6 hours; '12hour': 12 hours; '24hour': 24 hours; '7day': 7 days;",
						},
						"starting_time": {
							Type: schema.TypeString, Optional: true, Description: "Configure learning starting time",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"modify_time": {
				Type: schema.TypeString, Optional: true, Description: "Configure modify time",
			},
			"protection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_service_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port_num": {
													Type: schema.TypeInt, Required: true, Description: "Port Number",
												},
												"protocol": {
													Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port;",
												},
												"pbe": {
													Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"zone_service_other_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"port_other": {
													Type: schema.TypeString, Required: true, Description: "'other': other;",
												},
												"protocol": {
													Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
												},
												"pbe": {
													Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"ip_proto": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proto_name_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type: schema.TypeString, Required: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6;",
												},
												"pbe": {
													Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
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
						"port_range_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_range_start": {
										Type: schema.TypeInt, Required: true, Description: "Port-Range Start Port Number",
									},
									"port_range_end": {
										Type: schema.TypeInt, Required: true, Description: "Port-Range End Port Number",
									},
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port;",
									},
									"pbe": {
										Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
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
			"sensitivity": {
				Type: schema.TypeString, Optional: true, Default: "3", Description: "'5': Low; '3': Medium; '1.5': High;",
			},
			"status": {
				Type: schema.TypeString, Optional: true, Default: "newly", Description: "'newly': newly; 'learning': learning; 'learned': learned; 'activated': activated;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneWebGuiCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGui(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneWebGuiUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGui(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneWebGuiDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGui(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneWebGuiRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGui(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneWebGuiLearning273(d []interface{}) edpt.DdosDstZoneWebGuiLearning273 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiLearning273
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Duration = in["duration"].(string)
		ret.StartingTime = in["starting_time"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZoneWebGuiProtection274(d []interface{}) edpt.DdosDstZoneWebGuiProtection274 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtection274
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = getObjectDdosDstZoneWebGuiProtectionPort275(in["port"].([]interface{}))
		ret.IpProto = getObjectDdosDstZoneWebGuiProtectionIpProto278(in["ip_proto"].([]interface{}))
		ret.PortRangeList = getSliceDdosDstZoneWebGuiProtectionPortRangeList280(in["port_range_list"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneWebGuiProtectionPort275(d []interface{}) edpt.DdosDstZoneWebGuiProtectionPort275 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtectionPort275
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneServiceList = getSliceDdosDstZoneWebGuiProtectionPortZoneServiceList276(in["zone_service_list"].([]interface{}))
		ret.ZoneServiceOtherList = getSliceDdosDstZoneWebGuiProtectionPortZoneServiceOtherList277(in["zone_service_other_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortZoneServiceList276(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList276 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList276, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList276
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortZoneServiceOtherList277(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList277 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList277, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList277
		oi.PortOther = in["port_other"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneWebGuiProtectionIpProto278(d []interface{}) edpt.DdosDstZoneWebGuiProtectionIpProto278 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtectionIpProto278
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProtoNameList = getSliceDdosDstZoneWebGuiProtectionIpProtoProtoNameList279(in["proto_name_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionIpProtoProtoNameList279(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList279 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList279, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList279
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortRangeList280(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortRangeList280 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortRangeList280, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortRangeList280
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneWebGui(d *schema.ResourceData) edpt.DdosDstZoneWebGui {
	var ret edpt.DdosDstZoneWebGui
	ret.Inst.ActivatedAfterLearning = d.Get("activated_after_learning").(int)
	ret.Inst.CreateTime = d.Get("create_time").(string)
	ret.Inst.Learning = getObjectDdosDstZoneWebGuiLearning273(d.Get("learning").([]interface{}))
	ret.Inst.ModifyTime = d.Get("modify_time").(string)
	ret.Inst.Protection = getObjectDdosDstZoneWebGuiProtection274(d.Get("protection").([]interface{}))
	ret.Inst.Sensitivity = d.Get("sensitivity").(string)
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
