package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAsicDebugDump() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_asic_debug_dump`: Enable/Disable L2L3 ASIC traffic discard/drop events and Dump debug information\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemAsicDebugDumpCreate,
		UpdateContext: resourceSystemAsicDebugDumpUpdate,
		ReadContext:   resourceSystemAsicDebugDumpRead,
		DeleteContext: resourceSystemAsicDebugDumpDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable L2L3 ASIC traffic discard/drop events and Dump debug information",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemAsicDebugDumpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicDebugDumpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicDebugDump(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAsicDebugDumpRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemAsicDebugDumpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicDebugDumpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicDebugDump(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAsicDebugDumpRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemAsicDebugDumpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicDebugDumpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicDebugDump(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemAsicDebugDumpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicDebugDumpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicDebugDump(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemAsicDebugDump(d *schema.ResourceData) edpt.SystemAsicDebugDump {
	var ret edpt.SystemAsicDebugDump
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
