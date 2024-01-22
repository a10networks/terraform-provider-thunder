package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDeleteCpuCore() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_delete_cpu_core`: delete CPU cores\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemDeleteCpuCoreCreate,
		UpdateContext: resourceSystemDeleteCpuCoreUpdate,
		ReadContext:   resourceSystemDeleteCpuCoreRead,
		DeleteContext: resourceSystemDeleteCpuCoreDelete,

		Schema: map[string]*schema.Schema{
			"core_index": {
				Type: schema.TypeInt, Optional: true, Description: "core index to be deleted (Specify core index)",
			},
		},
	}
}
func resourceSystemDeleteCpuCoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDeleteCpuCoreCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDeleteCpuCore(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDeleteCpuCoreRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemDeleteCpuCoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDeleteCpuCoreUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDeleteCpuCore(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDeleteCpuCoreRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemDeleteCpuCoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDeleteCpuCoreDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDeleteCpuCore(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemDeleteCpuCoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDeleteCpuCoreRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDeleteCpuCore(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemDeleteCpuCore(d *schema.ResourceData) edpt.SystemDeleteCpuCore {
	var ret edpt.SystemDeleteCpuCore
	ret.Inst.CoreIndex = d.Get("core_index").(int)
	return ret
}
