package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ScaleoutNatIpHashingScheme() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_scaleout_nat_ip_hashing_scheme`: Configure Scaleout NAT IP Hashing scheme\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6ScaleoutNatIpHashingSchemeCreate,
		UpdateContext: resourceCgnv6ScaleoutNatIpHashingSchemeUpdate,
		ReadContext:   resourceCgnv6ScaleoutNatIpHashingSchemeRead,
		DeleteContext: resourceCgnv6ScaleoutNatIpHashingSchemeDelete,

		Schema: map[string]*schema.Schema{
			"hashing_type": {
				Type: schema.TypeString, Optional: true, Default: "route-aggregation", Description: "'route-aggregation': Chunk contiguous NAT IPs for route aggregation(default); 'mod-user-groups': Hash NAT IPs by taking mod of user-groups;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6ScaleoutNatIpHashingSchemeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ScaleoutNatIpHashingSchemeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ScaleoutNatIpHashingScheme(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ScaleoutNatIpHashingSchemeRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6ScaleoutNatIpHashingSchemeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ScaleoutNatIpHashingSchemeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ScaleoutNatIpHashingScheme(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ScaleoutNatIpHashingSchemeRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6ScaleoutNatIpHashingSchemeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ScaleoutNatIpHashingSchemeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ScaleoutNatIpHashingScheme(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6ScaleoutNatIpHashingSchemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ScaleoutNatIpHashingSchemeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ScaleoutNatIpHashingScheme(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6ScaleoutNatIpHashingScheme(d *schema.ResourceData) edpt.Cgnv6ScaleoutNatIpHashingScheme {
	var ret edpt.Cgnv6ScaleoutNatIpHashingScheme
	ret.Inst.HashingType = d.Get("hashing_type").(string)
	//omit uuid
	return ret
}
