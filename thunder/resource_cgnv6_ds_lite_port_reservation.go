package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLitePortReservation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_port_reservation`: DS-Lite Static Port Reservation\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLitePortReservationCreate,
		UpdateContext: resourceCgnv6DsLitePortReservationUpdate,
		ReadContext:   resourceCgnv6DsLitePortReservationRead,
		DeleteContext: resourceCgnv6DsLitePortReservationDelete,

		Schema: map[string]*schema.Schema{
			"inside": {
				Type: schema.TypeString, Required: true, Description: "Inside User Address and Port Range (DS-Lite Inside User's Tunnel Source IPv6 Address)",
			},
			"inside_addr": {
				Type: schema.TypeString, Required: true, Description: "Inside User IP address",
			},
			"inside_end_port": {
				Type: schema.TypeInt, Required: true, Description: "Inside End Port",
			},
			"inside_start_port": {
				Type: schema.TypeInt, Required: true, Description: "Inside Start Port",
			},
			"nat": {
				Type: schema.TypeString, Required: true, Description: "NAT Port Range (NAT IP address)",
			},
			"nat_end_port": {
				Type: schema.TypeInt, Required: true, Description: "NAT End Port",
			},
			"nat_start_port": {
				Type: schema.TypeInt, Required: true, Description: "NAT Start Port",
			},
			"tunnel_dest_address": {
				Type: schema.TypeString, Required: true, Description: "DS-Lite Inside User's Tunnel Destination IPv6 Address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLitePortReservationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLitePortReservationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLitePortReservation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLitePortReservationRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLitePortReservationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLitePortReservationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLitePortReservation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLitePortReservationRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLitePortReservationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLitePortReservationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLitePortReservation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLitePortReservationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLitePortReservationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLitePortReservation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6DsLitePortReservation(d *schema.ResourceData) edpt.Cgnv6DsLitePortReservation {
	var ret edpt.Cgnv6DsLitePortReservation
	ret.Inst.Inside = d.Get("inside").(string)
	ret.Inst.InsideAddr = d.Get("inside_addr").(string)
	ret.Inst.InsideEndPort = d.Get("inside_end_port").(int)
	ret.Inst.InsideStartPort = d.Get("inside_start_port").(int)
	ret.Inst.Nat = d.Get("nat").(string)
	ret.Inst.NatEndPort = d.Get("nat_end_port").(int)
	ret.Inst.NatStartPort = d.Get("nat_start_port").(int)
	ret.Inst.TunnelDestAddress = d.Get("tunnel_dest_address").(string)
	//omit uuid
	return ret
}
