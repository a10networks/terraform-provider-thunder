package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_vrid_blade_parameters_tracking_options_gateway_ipv6_gateway`: IPv6 Gateway\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayCreate,
		UpdateContext: resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayUpdate,
		ReadContext:   resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayRead,
		DeleteContext: resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "IPV6 address",
			},
			"priority_cost": {
				Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid_val": {
				Type: schema.TypeString, Required: true, Description: "VridVal",
			},
		},
	}
}
func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway(d *schema.ResourceData) edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway {
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6Gateway
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.PriorityCost = d.Get("priority_cost").(int)
	//omit uuid
	ret.Inst.VridVal = d.Get("vrid_val").(string)
	return ret
}
