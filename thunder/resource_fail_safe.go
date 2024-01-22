package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFailSafe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fail_safe`: Fail Safe Global Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFailSafeCreate,
		UpdateContext: resourceFailSafeUpdate,
		ReadContext:   resourceFailSafeRead,
		DeleteContext: resourceFailSafeDelete,

		Schema: map[string]*schema.Schema{
			"config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dataplane_recovery_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "dataplane hung detection timeout before ACOS is restarted (in seconds)",
			},
			"disable_failsafe": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Default: "all", Description: "'all': Disable All; 'io-buffer': Disable I/O Buffer; 'session-memory': Disable Session Memory; 'system-memory': Disable System Memory;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"fpga_buff_recovery_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "FPGA buffers recovery threshold (Units of 256 buffers (default 2))",
			},
			"fpga_monitor_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "FPGA monitor feature enable",
			},
			"fpga_monitor_forced_reboot": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "FPGA monitor forced reboot in error condition",
			},
			"fpga_monitor_interval": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "FPGA monitor packet interval (seconds) (Numbers of seconds between sending packets (default 1))",
			},
			"fpga_monitor_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "FPGA monitor packet missed for error condition (Numbers of missed monitor packets before setting error condition (default 30))",
			},
			"hw_error_monitor": {
				Type: schema.TypeString, Optional: true, Default: "hw-error-monitor-enable", Description: "'hw-error-monitor-disable': Disable fail-safe hardware error monitor; 'hw-error-monitor-enable': Enable fail-safe hardware error monitor;",
			},
			"hw_error_recovery_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Hardware error recovery timeout (minutes) (waiting time of recovery from hardware errors (default 0))",
			},
			"hw_ssl_timeout_monitor": {
				Type: schema.TypeString, Optional: true, Default: "hw-ssl-timeout-monitor-enable", Description: "'hw-ssl-timeout-monitor-disable': Disable fail-safe hardware SSL timeout monitor; 'hw-ssl-timeout-monitor-enable': Enable fail-safe hardware SSL timeout monitor;",
			},
			"kill": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Stop the traffic and log the event",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log the event",
			},
			"session_mem_recovery_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Session memory recovery threshold (percentage) (Percentage of available session memory (default 30%))",
			},
			"sw_error_monitor_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable fail-safe software error monitor",
			},
			"sw_error_recovery_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Software error recovery timeout (minutes) (waiting time of recovery from software errors (default 3))",
			},
			"total_memory_size_check": {
				Type: schema.TypeInt, Optional: true, Description: "Check total memory size of current system (Size of memory (GB))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFailSafeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFailSafeRead(ctx, d, meta)
	}
	return diags
}

func resourceFailSafeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFailSafeRead(ctx, d, meta)
	}
	return diags
}
func resourceFailSafeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFailSafeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFailSafeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFailSafe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFailSafeConfig349(d []interface{}) edpt.FailSafeConfig349 {

	var ret edpt.FailSafeConfig349
	return ret
}

func getObjectFailSafeDisableFailsafe350(d []interface{}) edpt.FailSafeDisableFailsafe350 {

	count1 := len(d)
	var ret edpt.FailSafeDisableFailsafe350
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointFailSafe(d *schema.ResourceData) edpt.FailSafe {
	var ret edpt.FailSafe
	ret.Inst.Config = getObjectFailSafeConfig349(d.Get("config").([]interface{}))
	ret.Inst.DataplaneRecoveryTimeout = d.Get("dataplane_recovery_timeout").(int)
	ret.Inst.DisableFailsafe = getObjectFailSafeDisableFailsafe350(d.Get("disable_failsafe").([]interface{}))
	ret.Inst.FpgaBuffRecoveryThreshold = d.Get("fpga_buff_recovery_threshold").(int)
	ret.Inst.FpgaMonitorEnable = d.Get("fpga_monitor_enable").(int)
	ret.Inst.FpgaMonitorForcedReboot = d.Get("fpga_monitor_forced_reboot").(int)
	ret.Inst.FpgaMonitorInterval = d.Get("fpga_monitor_interval").(int)
	ret.Inst.FpgaMonitorThreshold = d.Get("fpga_monitor_threshold").(int)
	ret.Inst.HwErrorMonitor = d.Get("hw_error_monitor").(string)
	ret.Inst.HwErrorRecoveryTimeout = d.Get("hw_error_recovery_timeout").(int)
	ret.Inst.HwSslTimeoutMonitor = d.Get("hw_ssl_timeout_monitor").(string)
	ret.Inst.Kill = d.Get("kill").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.SessionMemRecoveryThreshold = d.Get("session_mem_recovery_threshold").(int)
	ret.Inst.SwErrorMonitorEnable = d.Get("sw_error_monitor_enable").(int)
	ret.Inst.SwErrorRecoveryTimeout = d.Get("sw_error_recovery_timeout").(int)
	ret.Inst.TotalMemorySizeCheck = d.Get("total_memory_size_check").(int)
	//omit uuid
	return ret
}
