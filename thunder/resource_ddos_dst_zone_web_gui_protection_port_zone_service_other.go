package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneWebGuiProtectionPortZoneServiceOther() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_web_gui_protection_port_zone_service_other`: DDOS Port & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherCreate,
		UpdateContext: resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherUpdate,
		ReadContext:   resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherRead,
		DeleteContext: resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherDelete,

		Schema: map[string]*schema.Schema{
			"pbe": {
				Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
			},
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "'other': other;",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
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
func resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneServiceOther(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneServiceOther(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneServiceOther(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortZoneServiceOtherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortZoneServiceOther(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneWebGuiProtectionPortZoneServiceOther(d *schema.ResourceData) edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOther {
	var ret edpt.DdosDstZoneWebGuiProtectionPortZoneServiceOther
	ret.Inst.Pbe = d.Get("pbe").(string)
	ret.Inst.PortOther = d.Get("port_other").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
