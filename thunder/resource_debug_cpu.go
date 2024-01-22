package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugCpu() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_cpu`: Debug CPU\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugCpuCreate,
		UpdateContext: resourceDebugCpuUpdate,
		ReadContext:   resourceDebugCpuRead,
		DeleteContext: resourceDebugCpuDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugCpuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCpuCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCpu(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCpuRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugCpuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCpuUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCpu(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugCpuRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugCpuDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCpuDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCpu(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugCpuRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugCpuRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugCpu(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugCpu(d *schema.ResourceData) edpt.DebugCpu {
	var ret edpt.DebugCpu
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
