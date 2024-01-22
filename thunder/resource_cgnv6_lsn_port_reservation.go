package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPortReservation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_port_reservation`: Set Port Reservations\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnPortReservationCreate,
		UpdateContext: resourceCgnv6LsnPortReservationUpdate,
		ReadContext:   resourceCgnv6LsnPortReservationRead,
		DeleteContext: resourceCgnv6LsnPortReservationDelete,

		Schema: map[string]*schema.Schema{
			"inside": {
				Type: schema.TypeString, Required: true, Description: "Inside User Address and Port Range (Inside User IP address)",
			},
			"inside_port_end": {
				Type: schema.TypeInt, Required: true, Description: "Inside End Port",
			},
			"inside_port_start": {
				Type: schema.TypeInt, Required: true, Description: "Inside Start Port",
			},
			"nat": {
				Type: schema.TypeString, Required: true, Description: "NAT IP address",
			},
			"nat_port_end": {
				Type: schema.TypeInt, Required: true, Description: "NAT End Port",
			},
			"nat_port_start": {
				Type: schema.TypeInt, Required: true, Description: "NAT Start Port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnPortReservationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortReservationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortReservation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortReservationRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnPortReservationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortReservationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortReservation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortReservationRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnPortReservationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortReservationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortReservation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnPortReservationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortReservationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortReservation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LsnPortReservation(d *schema.ResourceData) edpt.Cgnv6LsnPortReservation {
	var ret edpt.Cgnv6LsnPortReservation
	ret.Inst.Inside = d.Get("inside").(string)
	ret.Inst.InsidePortEnd = d.Get("inside_port_end").(int)
	ret.Inst.InsidePortStart = d.Get("inside_port_start").(int)
	ret.Inst.Nat = d.Get("nat").(string)
	ret.Inst.NatPortEnd = d.Get("nat_port_end").(int)
	ret.Inst.NatPortStart = d.Get("nat_port_start").(int)
	//omit uuid
	return ret
}
