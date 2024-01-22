package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOffloadCpus() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_offload_cpus`: Enable Offload CPUs\n\n__PLACEHOLDER__",
		CreateContext: resourceOffloadCpusCreate,
		UpdateContext: resourceOffloadCpusUpdate,
		ReadContext:   resourceOffloadCpusRead,
		DeleteContext: resourceOffloadCpusDelete,

		Schema: map[string]*schema.Schema{
			"num_offload_cpus": {
				Type: schema.TypeInt, Optional: true, Description: "Set number of offload CPUs. Max limit is platform dependent.",
			},
		},
	}
}
func resourceOffloadCpusCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOffloadCpusCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOffloadCpus(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOffloadCpusRead(ctx, d, meta)
	}
	return diags
}

func resourceOffloadCpusUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOffloadCpusUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOffloadCpus(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOffloadCpusRead(ctx, d, meta)
	}
	return diags
}
func resourceOffloadCpusDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOffloadCpusDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOffloadCpus(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOffloadCpusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOffloadCpusRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOffloadCpus(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOffloadCpus(d *schema.ResourceData) edpt.OffloadCpus {
	var ret edpt.OffloadCpus
	ret.Inst.NumOffloadCpus = d.Get("num_offload_cpus").(int)
	return ret
}
