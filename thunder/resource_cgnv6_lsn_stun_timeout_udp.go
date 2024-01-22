package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnStunTimeoutUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_stun_timeout_udp`: Set UDP STUN timeout\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnStunTimeoutUdpCreate,
		UpdateContext: resourceCgnv6LsnStunTimeoutUdpUpdate,
		ReadContext:   resourceCgnv6LsnStunTimeoutUdpRead,
		DeleteContext: resourceCgnv6LsnStunTimeoutUdpDelete,

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
func resourceCgnv6LsnStunTimeoutUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnStunTimeoutUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnStunTimeoutUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnStunTimeoutUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnStunTimeoutUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnStunTimeoutUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnStunTimeoutUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnStunTimeoutUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LsnStunTimeoutUdp(d *schema.ResourceData) edpt.Cgnv6LsnStunTimeoutUdp {
	var ret edpt.Cgnv6LsnStunTimeoutUdp
	ret.Inst.PortEnd = d.Get("port_end").(int)
	ret.Inst.PortStart = d.Get("port_start").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
