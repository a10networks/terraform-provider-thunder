package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_udp`: System udp global config\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemUdpCreate,
		UpdateContext: resourceSystemUdpUpdate,
		ReadContext:   resourceSystemUdpRead,
		DeleteContext: resourceSystemUdpDelete,

		Schema: map[string]*schema.Schema{
			"skip_checksum_when_zero": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "When enabled, the system does not modify the L4 checksum if the original passthrough packet was received with UDP checksum valu",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemUdp(d *schema.ResourceData) edpt.SystemUdp {
	var ret edpt.SystemUdp
	ret.Inst.SkipChecksumWhenZero = d.Get("skip_checksum_when_zero").(int)
	//omit uuid
	return ret
}
