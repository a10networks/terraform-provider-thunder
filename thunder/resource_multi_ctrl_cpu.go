package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMultiCtrlCpu() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_multi_ctrl_cpu`: Enable Multiple Control CPUs\n\n__PLACEHOLDER__",
		CreateContext: resourceMultiCtrlCpuCreate,
		UpdateContext: resourceMultiCtrlCpuUpdate,
		ReadContext:   resourceMultiCtrlCpuRead,
		DeleteContext: resourceMultiCtrlCpuDelete,

		Schema: map[string]*schema.Schema{
			"num_ctrl_cpus": {
				Type: schema.TypeInt, Optional: true, Description: "Set number of control CPUs. Default is 1, and max limit is platform dependent.",
			},
		},
	}
}
func resourceMultiCtrlCpuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMultiCtrlCpuCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMultiCtrlCpu(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMultiCtrlCpuRead(ctx, d, meta)
	}
	return diags
}

func resourceMultiCtrlCpuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMultiCtrlCpuUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMultiCtrlCpu(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMultiCtrlCpuRead(ctx, d, meta)
	}
	return diags
}
func resourceMultiCtrlCpuDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMultiCtrlCpuDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMultiCtrlCpu(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceMultiCtrlCpuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMultiCtrlCpuRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMultiCtrlCpu(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointMultiCtrlCpu(d *schema.ResourceData) edpt.MultiCtrlCpu {
	var ret edpt.MultiCtrlCpu
	ret.Inst.NumCtrlCpus = d.Get("num_ctrl_cpus").(int)
	return ret
}
