package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodExternal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_external`: EXTERNAL type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodExternalCreate,
		UpdateContext: resourceHealthMonitorMethodExternalUpdate,
		ReadContext:   resourceHealthMonitorMethodExternalRead,
		DeleteContext: resourceHealthMonitorMethodExternalDelete,

		Schema: map[string]*schema.Schema{
			"ext_arguments": {
				Type: schema.TypeString, Optional: true, Description: "Specify external application's arguments (Application arguments)",
			},
			"ext_port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the server port (Port Number)",
			},
			"ext_preference": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Get server's perference",
			},
			"ext_program": {
				Type: schema.TypeString, Optional: true, Description: "Specify external application (Program name)",
			},
			"ext_program_shared": {
				Type: schema.TypeString, Optional: true, Description: "Specify external application (Program name)",
			},
			"external": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "EXTERNAL type",
			},
			"shared_partition_program": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "external application from shared partition",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodExternalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodExternalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodExternal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodExternalRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodExternalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodExternalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodExternal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodExternalRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodExternalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodExternalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodExternal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodExternalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodExternalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodExternal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodExternal(d *schema.ResourceData) edpt.HealthMonitorMethodExternal {
	var ret edpt.HealthMonitorMethodExternal
	ret.Inst.ExtArguments = d.Get("ext_arguments").(string)
	ret.Inst.ExtPort = d.Get("ext_port").(int)
	ret.Inst.ExtPreference = d.Get("ext_preference").(int)
	ret.Inst.ExtProgram = d.Get("ext_program").(string)
	ret.Inst.ExtProgramShared = d.Get("ext_program_shared").(string)
	ret.Inst.External = d.Get("external").(int)
	ret.Inst.SharedPartitionProgram = d.Get("shared_partition_program").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
