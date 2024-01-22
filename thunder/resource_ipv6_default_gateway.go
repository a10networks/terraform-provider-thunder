package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6DefaultGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_default_gateway`: Default gateway address\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6DefaultGatewayCreate,
		UpdateContext: resourceIpv6DefaultGatewayUpdate,
		ReadContext:   resourceIpv6DefaultGatewayRead,
		DeleteContext: resourceIpv6DefaultGatewayDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_default_gateway": {
				Type: schema.TypeString, Optional: true, Description: "Default gateway address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6DefaultGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6DefaultGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6DefaultGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6DefaultGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6DefaultGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6DefaultGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6DefaultGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6DefaultGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6DefaultGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6DefaultGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6DefaultGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6DefaultGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6DefaultGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6DefaultGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6DefaultGateway(d *schema.ResourceData) edpt.Ipv6DefaultGateway {
	var ret edpt.Ipv6DefaultGateway
	ret.Inst.Ipv6DefaultGateway = d.Get("ipv6_default_gateway").(string)
	//omit uuid
	return ret
}
