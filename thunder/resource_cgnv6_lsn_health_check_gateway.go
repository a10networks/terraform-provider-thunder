package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnHealthCheckGateway() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_health_check_gateway`: Configure LSN health-check gateway\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnHealthCheckGatewayCreate,
		UpdateContext: resourceCgnv6LsnHealthCheckGatewayUpdate,
		ReadContext:   resourceCgnv6LsnHealthCheckGatewayRead,
		DeleteContext: resourceCgnv6LsnHealthCheckGatewayDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "Specify IPv4 Gateway",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "Specify IPv6 Gateway",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnHealthCheckGatewayCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnHealthCheckGatewayCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnHealthCheckGateway(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnHealthCheckGatewayRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnHealthCheckGatewayUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnHealthCheckGatewayUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnHealthCheckGateway(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnHealthCheckGatewayRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnHealthCheckGatewayDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnHealthCheckGatewayDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnHealthCheckGateway(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnHealthCheckGatewayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnHealthCheckGatewayRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnHealthCheckGateway(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LsnHealthCheckGateway(d *schema.ResourceData) edpt.Cgnv6LsnHealthCheckGateway {
	var ret edpt.Cgnv6LsnHealthCheckGateway
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	//omit uuid
	return ret
}
