package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbGamingProtocolCompliancePort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_gaming_protocol_compliance_port`: Gaming Protocol Compliance Port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbGamingProtocolCompliancePortCreate,
		UpdateContext: resourceSlbGamingProtocolCompliancePortUpdate,
		ReadContext:   resourceSlbGamingProtocolCompliancePortRead,
		DeleteContext: resourceSlbGamingProtocolCompliancePortDelete,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"range": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Port range",
			},
			"udp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "UDP",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbGamingProtocolCompliancePortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolCompliancePortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolCompliancePort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbGamingProtocolCompliancePortRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbGamingProtocolCompliancePortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolCompliancePortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolCompliancePort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbGamingProtocolCompliancePortRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbGamingProtocolCompliancePortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolCompliancePortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolCompliancePort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbGamingProtocolCompliancePortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbGamingProtocolCompliancePortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbGamingProtocolCompliancePort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbGamingProtocolCompliancePort(d *schema.ResourceData) edpt.SlbGamingProtocolCompliancePort {
	var ret edpt.SlbGamingProtocolCompliancePort
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Range = d.Get("range").(int)
	ret.Inst.Udp = d.Get("udp").(int)
	//omit uuid
	return ret
}
