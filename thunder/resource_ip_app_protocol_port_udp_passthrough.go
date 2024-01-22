package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortUdpPassthrough() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_udp_passthrough`: Controls UDP ports filtering on all interfaces, Default mode is enabled\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortUdpPassthroughCreate,
		UpdateContext: resourceIpAppProtocolPortUdpPassthroughUpdate,
		ReadContext:   resourceIpAppProtocolPortUdpPassthroughRead,
		DeleteContext: resourceIpAppProtocolPortUdpPassthroughDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable global port processing",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enables global port processing",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpAppProtocolPortUdpPassthroughCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPassthroughCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPassthrough(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPassthroughRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortUdpPassthroughUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPassthroughUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPassthrough(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPassthroughRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortUdpPassthroughDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPassthroughDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPassthrough(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortUdpPassthroughRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPassthroughRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPassthrough(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpAppProtocolPortUdpPassthrough(d *schema.ResourceData) edpt.IpAppProtocolPortUdpPassthrough {
	var ret edpt.IpAppProtocolPortUdpPassthrough
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
