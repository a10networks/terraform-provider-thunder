package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingHostIpv4addr() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_host_ipv4addr`: host DNS name or ipv4 address of remote syslog server\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingHostIpv4addrCreate,
		UpdateContext: resourceLoggingHostIpv4addrUpdate,
		ReadContext:   resourceLoggingHostIpv4addrRead,
		DeleteContext: resourceLoggingHostIpv4addrDelete,

		Schema: map[string]*schema.Schema{
			"host_ipv4": {
				Type: schema.TypeString, Required: true, Description: "Set syslog host ip address",
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
func resourceLoggingHostIpv4addrCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv4addrCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv4addr(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingHostIpv4addrRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingHostIpv4addrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv4addrUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv4addr(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingHostIpv4addrRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingHostIpv4addrDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv4addrDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv4addr(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingHostIpv4addrRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingHostIpv4addrRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingHostIpv4addr(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingHostIpv4addr(d *schema.ResourceData) edpt.LoggingHostIpv4addr {
	var ret edpt.LoggingHostIpv4addr
	ret.Inst.HostIpv4 = d.Get("host_ipv4").(string)
	ret.Inst.OverTls = d.Get("over_tls").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.Tcp = d.Get("tcp").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
