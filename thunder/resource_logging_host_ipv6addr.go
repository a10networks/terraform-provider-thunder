package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingHostIpv6addr() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_host_ipv6addr`: ipv6 address of remote syslog server\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingHostIpv6addrCreate,
		UpdateContext: resourceLoggingHostIpv6addrUpdate,
		ReadContext:   resourceLoggingHostIpv6addrRead,
		DeleteContext: resourceLoggingHostIpv6addrDelete,

		Schema: map[string]*schema.Schema{
			"host_ipv6": {
				Type: schema.TypeString, Required: true, Description: "Set syslog host ipv6 address",
			},
			"over_tls": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable remote logging over TLS session",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 514, Description: "Set remote syslog port number",
			},
			"tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use TCP as transport protocol",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingHostIpv6addrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv6addrCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv6addr(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingHostIpv6addrRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingHostIpv6addrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv6addrUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv6addr(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingHostIpv6addrRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingHostIpv6addrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv6addrDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv6addr(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingHostIpv6addrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv6addrRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv6addr(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingHostIpv6addr(d *schema.ResourceData) edpt.LoggingHostIpv6addr {
	var ret edpt.LoggingHostIpv6addr
	ret.Inst.HostIpv6 = d.Get("host_ipv6").(string)
	ret.Inst.OverTls = d.Get("over_tls").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Tcp = d.Get("tcp").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
