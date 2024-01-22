package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneWebGuiProtectionPortZoneService() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_web_gui_protection_port_zone_service`: DDOS Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneWebGuiProtectionPortZoneServiceCreate,
		UpdateContext: resourceDdosDstZoneWebGuiProtectionPortZoneServiceUpdate,
		ReadContext:   resourceDdosDstZoneWebGuiProtectionPortZoneServiceRead,
		DeleteContext: resourceDdosDstZoneWebGuiProtectionPortZoneServiceDelete,

		Schema: map[string]*schema.Schema{
			"pbe": {
				Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port;",
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
func resourceDdosDstZoneWebGuiProtectionPortZoneServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneService(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionPortZoneServiceRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionPortZoneServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneService(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionPortZoneServiceRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneWebGuiProtectionPortZoneServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneService(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionPortZoneServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneService(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneWebGuiProtectionPortZoneService(d *schema.ResourceData) edpt.DdosDstZoneWebGuiProtectionPortZoneService {
	var ret edpt.DdosDstZoneWebGuiProtectionPortZoneService
	ret.Inst.Pbe = d.Get("pbe").(string)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
