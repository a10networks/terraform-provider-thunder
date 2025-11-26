package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6HealthCheckGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_health_check_gateway`: Configure LW-4over6 health-check gateway\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6HealthCheckGatewayCreate,
		UpdateContext: resourceCgnv6Lw4o6HealthCheckGatewayUpdate,
		ReadContext:   resourceCgnv6Lw4o6HealthCheckGatewayRead,
		DeleteContext: resourceCgnv6Lw4o6HealthCheckGatewayDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "LW-4over6 IPv4 Gateway",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "LW-4over6 IPv6 Gateway",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Lw4o6HealthCheckGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6HealthCheckGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6HealthCheckGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6HealthCheckGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6HealthCheckGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6HealthCheckGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6HealthCheckGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6HealthCheckGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6HealthCheckGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6HealthCheckGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6HealthCheckGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6HealthCheckGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6HealthCheckGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6HealthCheckGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Lw4o6HealthCheckGateway(d *schema.ResourceData) edpt.Cgnv6Lw4o6HealthCheckGateway {
	var ret edpt.Cgnv6Lw4o6HealthCheckGateway
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	//omit uuid
	return ret
}
