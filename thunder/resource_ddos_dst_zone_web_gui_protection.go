package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneWebGuiProtection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_web_gui_protection`: Configure TPS Wizard Protection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneWebGuiProtectionCreate,
		UpdateContext: resourceDdosDstZoneWebGuiProtectionUpdate,
		ReadContext:   resourceDdosDstZoneWebGuiProtectionRead,
		DeleteContext: resourceDdosDstZoneWebGuiProtectionDelete,

		Schema: map[string]*schema.Schema{
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneWebGuiProtectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneWebGuiProtectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneWebGuiProtectionIpProto244(d []interface{}) edpt.DdosDstZoneWebGuiProtectionIpProto244 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtectionIpProto244
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProtoNameList = getSliceDdosDstZoneWebGuiProtectionIpProtoProtoNameList(in["proto_name_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionIpProtoProtoNameList(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionIpProtoProtoNameList
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneWebGuiProtectionPort245(d []interface{}) edpt.DdosDstZoneWebGuiProtectionPort245 {

	count1 := len(d)
	var ret edpt.DdosDstZoneWebGuiProtectionPort245
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneServiceList = getSliceDdosDstZoneWebGuiProtectionPortZoneServiceList(in["zone_service_list"].([]interface{}))
		ret.ZoneServiceOtherList = getSliceDdosDstZoneWebGuiProtectionPortZoneServiceOtherList(in["zone_service_other_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortZoneServiceList(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortZoneServiceList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortZoneServiceOtherList(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOtherList
		oi.PortOther = in["port_other"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Pbe = in["pbe"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneWebGuiProtectionPortRangeList(d []interface{}) []edpt.DdosDstZoneWebGuiProtectionPortRangeList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneWebGuiProtectionPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneWebGuiProtectionPortRangeList
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

func dataToEndpointDdosDstZoneWebGuiProtection(d *schema.ResourceData) edpt.DdosDstZoneWebGuiProtection {
	var ret edpt.DdosDstZoneWebGuiProtection
	ret.Inst.IpProto = getObjectDdosDstZoneWebGuiProtectionIpProto244(d.Get("ip_proto").([]interface{}))
	ret.Inst.Port = getObjectDdosDstZoneWebGuiProtectionPort245(d.Get("port").([]interface{}))
	ret.Inst.PortRangeList = getSliceDdosDstZoneWebGuiProtectionPortRangeList(d.Get("port_range_list").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
