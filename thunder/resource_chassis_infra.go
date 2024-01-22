package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChassisInfra() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_chassis_infra`: Chassis infrastructure commands\n\n__PLACEHOLDER__",
		CreateContext: resourceChassisInfraCreate,
		UpdateContext: resourceChassisInfraUpdate,
		ReadContext:   resourceChassisInfraRead,
		DeleteContext: resourceChassisInfraDelete,

		Schema: map[string]*schema.Schema{
			"debug_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable chassis infrastruture debugging",
			},
			"debug_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable chassis infrastruture debugging",
			},
			"debug_status": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show chassis infrastruture debugging status",
			},
			"detailed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Give Chassis filesystem info( USED BY TAC ONLY )",
			},
			"sys_sync": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Synchronize the Master and Blade filesystems (For A10 TAC Use Only)",
			},
			"system_sync_verify": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Validate chassis filesytem synchronization status (For A10 TAC Use Only)",
			},
		},
	}
}
func resourceChassisInfraCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfraCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfra(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceChassisInfraRead(ctx, d, meta)
	}
	return diags
}

func resourceChassisInfraUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfraUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfra(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceChassisInfraRead(ctx, d, meta)
	}
	return diags
}
func resourceChassisInfraDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfraDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfra(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceChassisInfraRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfraRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfra(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointChassisInfra(d *schema.ResourceData) edpt.ChassisInfra {
	var ret edpt.ChassisInfra
	ret.Inst.DebugDisable = d.Get("debug_disable").(int)
	ret.Inst.DebugEnable = d.Get("debug_enable").(int)
	ret.Inst.DebugStatus = d.Get("debug_status").(int)
	ret.Inst.Detailed = d.Get("detailed").(int)
	ret.Inst.SysSync = d.Get("sys_sync").(int)
	ret.Inst.SystemSyncVerify = d.Get("system_sync_verify").(int)
	return ret
}
