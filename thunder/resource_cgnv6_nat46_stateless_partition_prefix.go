package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat46StatelessPartitionPrefix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat46_stateless_partition_prefix`: Configure NAT46 prefix for a L3V partition traffic\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat46StatelessPartitionPrefixCreate,
		UpdateContext: resourceCgnv6Nat46StatelessPartitionPrefixUpdate,
		ReadContext:   resourceCgnv6Nat46StatelessPartitionPrefixRead,
		DeleteContext: resourceCgnv6Nat46StatelessPartitionPrefixDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_prefix": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 prefix",
			},
			"partition": {
				Type: schema.TypeString, Required: true, Description: "Partition Name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceCgnv6Nat46StatelessPartitionPrefixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPartitionPrefixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPartitionPrefix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessPartitionPrefixRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat46StatelessPartitionPrefixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPartitionPrefixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPartitionPrefix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessPartitionPrefixRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat46StatelessPartitionPrefixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPartitionPrefixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPartitionPrefix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat46StatelessPartitionPrefixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPartitionPrefixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPartitionPrefix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat46StatelessPartitionPrefix(d *schema.ResourceData) edpt.Cgnv6Nat46StatelessPartitionPrefix {
	var ret edpt.Cgnv6Nat46StatelessPartitionPrefix
	ret.Inst.Ipv6Prefix = d.Get("ipv6_prefix").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
