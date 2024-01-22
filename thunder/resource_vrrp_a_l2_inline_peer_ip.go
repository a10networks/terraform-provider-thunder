package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAL2InlinePeerIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_l2_inline_peer_ip`: Peer IP for Layer 2 inline mode\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAL2InlinePeerIpCreate,
		UpdateContext: resourceVrrpAL2InlinePeerIpUpdate,
		ReadContext:   resourceVrrpAL2InlinePeerIpRead,
		DeleteContext: resourceVrrpAL2InlinePeerIpDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type: schema.TypeString, Required: true, Description: "IP address (IPv4 address)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVrrpAL2InlinePeerIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL2InlinePeerIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL2InlinePeerIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAL2InlinePeerIpRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAL2InlinePeerIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL2InlinePeerIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL2InlinePeerIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAL2InlinePeerIpRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAL2InlinePeerIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL2InlinePeerIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL2InlinePeerIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAL2InlinePeerIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAL2InlinePeerIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAL2InlinePeerIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAL2InlinePeerIp(d *schema.ResourceData) edpt.VrrpAL2InlinePeerIp {
	var ret edpt.VrrpAL2InlinePeerIp
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	//omit uuid
	return ret
}
