package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecIpsecGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vpn_ipsec_ipsec_gateway`: Configure IKE gateway to use for IPsec\n\n__PLACEHOLDER__",
		CreateContext: resourceVpnIpsecIpsecGatewayCreate,
		UpdateContext: resourceVpnIpsecIpsecGatewayUpdate,
		ReadContext:   resourceVpnIpsecIpsecGatewayRead,
		DeleteContext: resourceVpnIpsecIpsecGatewayDelete,

		Schema: map[string]*schema.Schema{
			"ike_gateway": {
				Type: schema.TypeString, Optional: true, Description: "Gateway to use for IPsec SA",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVpnIpsecIpsecGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecIpsecGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecIpsecGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecIpsecGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceVpnIpsecIpsecGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecIpsecGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecIpsecGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecIpsecGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceVpnIpsecIpsecGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecIpsecGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecIpsecGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVpnIpsecIpsecGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecIpsecGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecIpsecGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVpnIpsecIpsecGateway(d *schema.ResourceData) edpt.VpnIpsecIpsecGateway {
	var ret edpt.VpnIpsecIpsecGateway
	ret.Inst.IkeGateway = d.Get("ike_gateway").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
