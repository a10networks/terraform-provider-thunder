package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnStunTimeoutTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_stun_timeout_tcp`: Set TCP STUN timeout\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnStunTimeoutTcpCreate,
		UpdateContext: resourceCgnv6LsnStunTimeoutTcpUpdate,
		ReadContext:   resourceCgnv6LsnStunTimeoutTcpRead,
		DeleteContext: resourceCgnv6LsnStunTimeoutTcpDelete,

		Schema: map[string]*schema.Schema{
			"port_end": {
				Type: schema.TypeInt, Required: true, Description: "Port Range (Port Range End)",
			},
			"port_start": {
				Type: schema.TypeInt, Required: true, Description: "Port Range (Port Range Start)",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Description: "STUN timeout in minutes (default: 2 minutes)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnStunTimeoutTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnStunTimeoutTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnStunTimeoutTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnStunTimeoutTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnStunTimeoutTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnStunTimeoutTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LsnStunTimeoutTcp(d *schema.ResourceData) edpt.Cgnv6LsnStunTimeoutTcp {
	var ret edpt.Cgnv6LsnStunTimeoutTcp
	ret.Inst.PortEnd = d.Get("port_end").(int)
	ret.Inst.PortStart = d.Get("port_start").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
