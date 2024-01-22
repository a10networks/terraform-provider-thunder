package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatIcmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_icmp`: Configure NAT ICMP settings\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatIcmpCreate,
		UpdateContext: resourceIpNatIcmpUpdate,
		ReadContext:   resourceIpNatIcmpRead,
		DeleteContext: resourceIpNatIcmpDelete,

		Schema: map[string]*schema.Schema{
			"always_source_nat_errors": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source NAT intermediate routers' IPs for ICMP errors (default: disabled)",
			},
			"respond_to_ping": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond to ICMP echo requests to NAT pool IPs (default: disabled)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpNatIcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatIcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatIcmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatIcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatIcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatIcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatIcmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatIcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatIcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatIcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatIcmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatIcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatIcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatIcmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatIcmp(d *schema.ResourceData) edpt.IpNatIcmp {
	var ret edpt.IpNatIcmp
	ret.Inst.AlwaysSourceNatErrors = d.Get("always_source_nat_errors").(int)
	ret.Inst.RespondToPing = d.Get("respond_to_ping").(int)
	//omit uuid
	return ret
}
