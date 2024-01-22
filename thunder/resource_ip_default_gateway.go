package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpDefaultGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_default_gateway`: Transparent mode gateway\n\n__PLACEHOLDER__",
		CreateContext: resourceIpDefaultGatewayCreate,
		UpdateContext: resourceIpDefaultGatewayUpdate,
		ReadContext:   resourceIpDefaultGatewayRead,
		DeleteContext: resourceIpDefaultGatewayDelete,

		Schema: map[string]*schema.Schema{
			"gateway_ip": {
				Type: schema.TypeString, Optional: true, Description: "Default gateway address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpDefaultGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDefaultGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDefaultGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDefaultGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceIpDefaultGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDefaultGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDefaultGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpDefaultGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceIpDefaultGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDefaultGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDefaultGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpDefaultGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpDefaultGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpDefaultGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpDefaultGateway(d *schema.ResourceData) edpt.IpDefaultGateway {
	var ret edpt.IpDefaultGateway
	ret.Inst.GatewayIp = d.Get("gateway_ip").(string)
	//omit uuid
	return ret
}
