package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat46StatelessStaticDestMapping() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat46_stateless_static_dest_mapping`: Stateless NAT46 mapping (IPv4 <-> IPv6)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat46StatelessStaticDestMappingCreate,
		UpdateContext: resourceCgnv6Nat46StatelessStaticDestMappingUpdate,
		ReadContext:   resourceCgnv6Nat46StatelessStaticDestMappingRead,
		DeleteContext: resourceCgnv6Nat46StatelessStaticDestMappingDelete,

		Schema: map[string]*schema.Schema{
			"count1": {
				Type: schema.TypeInt, Optional: true, Description: "Set number of consecutive mappings (Number of mappings)",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share/Expose this mapping with other partitions",
			},
			"to_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send NATed traffic through shared partition",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_address": {
				Type: schema.TypeString, Required: true, Description: "IPv4 address",
			},
			"v6_address": {
				Type: schema.TypeString, Required: true, Description: "IPv6 address",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceCgnv6Nat46StatelessStaticDestMappingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessStaticDestMappingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessStaticDestMapping(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessStaticDestMappingRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat46StatelessStaticDestMappingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessStaticDestMappingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessStaticDestMapping(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessStaticDestMappingRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat46StatelessStaticDestMappingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessStaticDestMappingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessStaticDestMapping(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat46StatelessStaticDestMappingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessStaticDestMappingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessStaticDestMapping(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat46StatelessStaticDestMapping(d *schema.ResourceData) edpt.Cgnv6Nat46StatelessStaticDestMapping {
	var ret edpt.Cgnv6Nat46StatelessStaticDestMapping
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.ToShared = d.Get("to_shared").(int)
	//omit uuid
	ret.Inst.V4Address = d.Get("v4_address").(string)
	ret.Inst.V6Address = d.Get("v6_address").(string)
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
