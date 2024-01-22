package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortTcpPassthrough() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_tcp_passthrough`: Controls TCP ports filtering on all interfaces, Default mode is enabled\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortTcpPassthroughCreate,
		UpdateContext: resourceIpAppProtocolPortTcpPassthroughUpdate,
		ReadContext:   resourceIpAppProtocolPortTcpPassthroughRead,
		DeleteContext: resourceIpAppProtocolPortTcpPassthroughDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable passthrough mode",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enables passthrough mode",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpAppProtocolPortTcpPassthroughCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPassthroughCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPassthrough(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPassthroughRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortTcpPassthroughUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPassthroughUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPassthrough(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPassthroughRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortTcpPassthroughDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPassthroughDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPassthrough(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortTcpPassthroughRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPassthroughRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPassthrough(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpAppProtocolPortTcpPassthrough(d *schema.ResourceData) edpt.IpAppProtocolPortTcpPassthrough {
	var ret edpt.IpAppProtocolPortTcpPassthrough
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
