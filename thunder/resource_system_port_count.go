package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPortCount() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_port_count`: Port count related commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemPortCountCreate,
		UpdateContext: resourceSystemPortCountUpdate,
		ReadContext:   resourceSystemPortCountRead,
		DeleteContext: resourceSystemPortCountDelete,

		Schema: map[string]*schema.Schema{
			"port_count_alg": {
				Type: schema.TypeInt, Optional: true, Default: 6000, Description: "Total Ports to be allocated for alg types.",
			},
			"port_count_hm": {
				Type: schema.TypeInt, Optional: true, Default: 1024, Description: "Total Ports to be allocated for hm.",
			},
			"port_count_kernel": {
				Type: schema.TypeInt, Optional: true, Default: 18000, Description: "Total Ports to be allocated for kernel.",
			},
			"port_count_logging": {
				Type: schema.TypeInt, Optional: true, Default: 4096, Description: "Total Ports to be allocated for logging.",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemPortCountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortCountCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPortCount(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPortCountRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemPortCountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortCountUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPortCount(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPortCountRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemPortCountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortCountDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPortCount(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemPortCountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortCountRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPortCount(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemPortCount(d *schema.ResourceData) edpt.SystemPortCount {
	var ret edpt.SystemPortCount
	ret.Inst.PortCountAlg = d.Get("port_count_alg").(int)
	ret.Inst.PortCountHm = d.Get("port_count_hm").(int)
	ret.Inst.PortCountKernel = d.Get("port_count_kernel").(int)
	ret.Inst.PortCountLogging = d.Get("port_count_logging").(int)
	//omit uuid
	return ret
}
