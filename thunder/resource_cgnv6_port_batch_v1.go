package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6PortBatchV1() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_port_batch_v1`: Configure Port-batch-v1 Behavior\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6PortBatchV1Create,
		UpdateContext: resourceCgnv6PortBatchV1Update,
		ReadContext:   resourceCgnv6PortBatchV1Read,
		DeleteContext: resourceCgnv6PortBatchV1Delete,

		Schema: map[string]*schema.Schema{
			"enable_port_batch_v1": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lsn port-batching v1 (default: disabled)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6PortBatchV1Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortBatchV1Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortBatchV1(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6PortBatchV1Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6PortBatchV1Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortBatchV1Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortBatchV1(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6PortBatchV1Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6PortBatchV1Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortBatchV1Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortBatchV1(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6PortBatchV1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6PortBatchV1Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6PortBatchV1(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6PortBatchV1(d *schema.ResourceData) edpt.Cgnv6PortBatchV1 {
	var ret edpt.Cgnv6PortBatchV1
	ret.Inst.EnablePortBatchV1 = d.Get("enable_port_batch_v1").(int)
	//omit uuid
	return ret
}
