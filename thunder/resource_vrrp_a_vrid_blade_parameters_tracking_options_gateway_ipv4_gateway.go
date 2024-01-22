package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_vrid_blade_parameters_tracking_options_gateway_ipv4_gateway`: IPv4 Gateway\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayCreate,
		UpdateContext: resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayUpdate,
		ReadContext:   resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayRead,
		DeleteContext: resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type: schema.TypeString, Required: true, Description: "IP Address",
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
func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway(d *schema.ResourceData) edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway {
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4Gateway
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.PriorityCost = d.Get("priority_cost").(int)
	//omit uuid
	ret.Inst.VridVal = d.Get("vrid_val").(string)
	return ret
}
