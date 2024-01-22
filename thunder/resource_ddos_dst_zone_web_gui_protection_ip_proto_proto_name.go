package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneWebGuiProtectionIpProtoProtoName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_web_gui_protection_ip_proto_proto_name`: DDOS IP protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameCreate,
		UpdateContext: resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameUpdate,
		ReadContext:   resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameRead,
		DeleteContext: resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameDelete,

		Schema: map[string]*schema.Schema{
			"pbe": {
				Type: schema.TypeString, Optional: true, Description: "Peak Bandwidth Expected",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6;",
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
func resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionIpProtoProtoName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionIpProtoProtoName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionIpProtoProtoName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiProtectionIpProtoProtoNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiProtectionIpProtoProtoName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneWebGuiProtectionIpProtoProtoName(d *schema.ResourceData) edpt.DdosDstZoneWebGuiProtectionIpProtoProtoName {
	var ret edpt.DdosDstZoneWebGuiProtectionIpProtoProtoName
	ret.Inst.Pbe = d.Get("pbe").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
