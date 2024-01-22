package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIoCpu() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_io_cpu`: IO CPU cores\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIoCpuCreate,
		UpdateContext: resourceSystemIoCpuUpdate,
		ReadContext:   resourceSystemIoCpuRead,
		DeleteContext: resourceSystemIoCpuDelete,

		Schema: map[string]*schema.Schema{
			"max_cores": {
				Type: schema.TypeInt, Optional: true, Description: "max number of IO cores (Specify number of cores)",
			},
		},
	}
}
func resourceSystemIoCpuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIoCpuCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIoCpu(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIoCpuRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIoCpuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIoCpuUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIoCpu(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIoCpuRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIoCpuDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIoCpuDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIoCpu(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIoCpuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIoCpuRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIoCpu(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIoCpu(d *schema.ResourceData) edpt.SystemIoCpu {
	var ret edpt.SystemIoCpu
	ret.Inst.MaxCores = d.Get("max_cores").(int)
	return ret
}
