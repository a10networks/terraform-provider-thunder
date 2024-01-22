package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_monitor`: Monitor debug output\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugMonitorCreate,
		UpdateContext: resourceDebugMonitorUpdate,
		ReadContext:   resourceDebugMonitorRead,
		DeleteContext: resourceDebugMonitorDelete,

		Schema: map[string]*schema.Schema{
			"all_slots": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Display debug output of both Master and Blade",
			},
			"cpuid": {
				Type: schema.TypeInt, Optional: true, Description: "CPU id to debug (0,1,...)",
			},
			"filename": {
				Type: schema.TypeString, Optional: true, Description: "Filename to save debug output",
			},
			"filesize": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "maximum file size to save debug messages (MB)",
			},
			"no_stop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Donot spawn another rimacli",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "timeout to stop debug monitor in minutes",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugMonitor(d *schema.ResourceData) edpt.DebugMonitor {
	var ret edpt.DebugMonitor
	ret.Inst.AllSlots = d.Get("all_slots").(int)
	ret.Inst.Cpuid = d.Get("cpuid").(int)
	ret.Inst.Filename = d.Get("filename").(string)
	ret.Inst.Filesize = d.Get("filesize").(int)
	ret.Inst.NoStop = d.Get("no_stop").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
