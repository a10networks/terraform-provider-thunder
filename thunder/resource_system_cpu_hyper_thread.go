package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuHyperThread() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_cpu_hyper_thread`: Cpu Hyperthreading\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemCpuHyperThreadCreate,
		UpdateContext: resourceSystemCpuHyperThreadUpdate,
		ReadContext:   resourceSystemCpuHyperThreadRead,
		DeleteContext: resourceSystemCpuHyperThreadDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable CPU Hyperthreading",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPU Hyperthreading",
			},
		},
	}
}
func resourceSystemCpuHyperThreadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuHyperThreadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuHyperThread(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuHyperThreadRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemCpuHyperThreadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuHyperThreadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuHyperThread(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCpuHyperThreadRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemCpuHyperThreadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuHyperThreadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuHyperThread(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemCpuHyperThreadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuHyperThreadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuHyperThread(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemCpuHyperThread(d *schema.ResourceData) edpt.SystemCpuHyperThread {
	var ret edpt.SystemCpuHyperThread
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	return ret
}
