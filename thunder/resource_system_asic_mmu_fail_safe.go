package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAsicMmuFailSafe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_asic_mmu_fail_safe`: ASIC Fail Safe Global Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemAsicMmuFailSafeCreate,
		UpdateContext: resourceSystemAsicMmuFailSafeUpdate,
		ReadContext:   resourceSystemAsicMmuFailSafeRead,
		DeleteContext: resourceSystemAsicMmuFailSafeDelete,

		Schema: map[string]*schema.Schema{
			"inject_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inject MMU SER/Parity errors",
			},
			"monitor_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Fail-safe software error monitoring and act on it",
			},
			"monitor_interval": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "ASIC Fail-safe monitoring intervals in Seconds (Units of 1 Seconds (default 60))",
			},
			"reboot_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable system reboot if system encounters mmu error",
			},
			"recovery_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "ASIC Fail-safe recovery threshold in Errors (Units of 1 Errors (default 2))",
			},
			"test_pattern_type": {
				Type: schema.TypeString, Optional: true, Default: "lcb", Description: "'all-zeros': Inject all bits 0s in a byte; 'all-ones': Inject all bits 1s in a byte; 'lcb': Logical checker board; 'inverse-lcb': Inverse Logical checker board;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemAsicMmuFailSafeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicMmuFailSafeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicMmuFailSafe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAsicMmuFailSafeRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemAsicMmuFailSafeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicMmuFailSafeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicMmuFailSafe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAsicMmuFailSafeRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemAsicMmuFailSafeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicMmuFailSafeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicMmuFailSafe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemAsicMmuFailSafeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAsicMmuFailSafeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAsicMmuFailSafe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemAsicMmuFailSafe(d *schema.ResourceData) edpt.SystemAsicMmuFailSafe {
	var ret edpt.SystemAsicMmuFailSafe
	ret.Inst.InjectError = d.Get("inject_error").(int)
	ret.Inst.MonitorDisable = d.Get("monitor_disable").(int)
	ret.Inst.MonitorInterval = d.Get("monitor_interval").(int)
	ret.Inst.RebootDisable = d.Get("reboot_disable").(int)
	ret.Inst.RecoveryThreshold = d.Get("recovery_threshold").(int)
	ret.Inst.TestPatternType = d.Get("test_pattern_type").(string)
	//omit uuid
	return ret
}
