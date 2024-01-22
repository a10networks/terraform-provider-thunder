package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneWebGuiProtectionPortRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_web_gui_protection_port_range`: DDOS Port-Range & Protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneWebGuiProtectionPortRangeCreate,
		UpdateContext: resourceDdosDstZoneWebGuiProtectionPortRangeUpdate,
		ReadContext:   resourceDdosDstZoneWebGuiProtectionPortRangeRead,
		DeleteContext: resourceDdosDstZoneWebGuiProtectionPortRangeDelete,

		Schema: map[string]*schema.Schema{
			"pbe": {
				Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
			},
			"port_range_end": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range End Port Number",
			},
			"port_range_start": {
				Type: schema.TypeInt, Required: true, Description: "Port-Range Start Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDdosDstZoneWebGuiProtectionPortRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionPortRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionPortRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionPortRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneWebGuiProtectionPortRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionPortRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionPortRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionPortRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneWebGuiProtectionPortRange(d *schema.ResourceData) edpt.DdosDstZoneWebGuiProtectionPortRange {
	var ret edpt.DdosDstZoneWebGuiProtectionPortRange
	ret.Inst.Pbe = d.Get("pbe").(string)
	ret.Inst.PortRangeEnd = d.Get("port_range_end").(int)
	ret.Inst.PortRangeStart = d.Get("port_range_start").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
