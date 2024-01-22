package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsSystem() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_system`: Enable system group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsSystemCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSystemUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSystemRead,
		DeleteContext: resourceSnmpServerEnableTrapsSystemDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all system group traps",
			},
			"apps_global": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sessions_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable sessions threshold trap",
						},
						"cps_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CPS trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"control_cpu_high": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable control CPU usage high trap",
			},
			"data_cpu_high": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable data CPU usage high trap",
			},
			"fan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system fan trap",
			},
			"file_sys_read_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable file system read-only trap",
			},
			"high_disk_use": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system high disk usage trap",
			},
			"high_memory_use": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system high memory usage trap",
			},
			"high_temp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system high temperature trap",
			},
			"license_management": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system license management traps",
			},
			"low_temp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system low temperature trap",
			},
			"packet_drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system packet dropped trap",
			},
			"power": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system power supply trap",
			},
			"pri_disk": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system primary hard disk trap",
			},
			"restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system restart trap",
			},
			"sec_disk": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system secondary hard disk trap",
			},
			"shutdown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system shutdown trap",
			},
			"smp_resource_event": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system smp resource event trap",
			},
			"start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system start trap",
			},
			"syslog_severity_one": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system syslog severity one messages trap",
			},
			"tacacs_server_up_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system TACACS monitor server up/down trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystem(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSystemRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystem(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSystemRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystem(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSystemRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSystem(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSnmpServerEnableTrapsSystemAppsGlobal1482(d []interface{}) edpt.SnmpServerEnableTrapsSystemAppsGlobal1482 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsSystemAppsGlobal1482
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionsThreshold = in["sessions_threshold"].(int)
		ret.CpsThreshold = in["cps_threshold"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSnmpServerEnableTrapsSystem(d *schema.ResourceData) edpt.SnmpServerEnableTrapsSystem {
	var ret edpt.SnmpServerEnableTrapsSystem
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.AppsGlobal = getObjectSnmpServerEnableTrapsSystemAppsGlobal1482(d.Get("apps_global").([]interface{}))
	ret.Inst.ControlCpuHigh = d.Get("control_cpu_high").(int)
	ret.Inst.DataCpuHigh = d.Get("data_cpu_high").(int)
	ret.Inst.Fan = d.Get("fan").(int)
	ret.Inst.FileSysReadOnly = d.Get("file_sys_read_only").(int)
	ret.Inst.HighDiskUse = d.Get("high_disk_use").(int)
	ret.Inst.HighMemoryUse = d.Get("high_memory_use").(int)
	ret.Inst.HighTemp = d.Get("high_temp").(int)
	ret.Inst.LicenseManagement = d.Get("license_management").(int)
	ret.Inst.LowTemp = d.Get("low_temp").(int)
	ret.Inst.PacketDrop = d.Get("packet_drop").(int)
	ret.Inst.Power = d.Get("power").(int)
	ret.Inst.PriDisk = d.Get("pri_disk").(int)
	ret.Inst.Restart = d.Get("restart").(int)
	ret.Inst.SecDisk = d.Get("sec_disk").(int)
	ret.Inst.Shutdown = d.Get("shutdown").(int)
	ret.Inst.SmpResourceEvent = d.Get("smp_resource_event").(int)
	ret.Inst.Start = d.Get("start").(int)
	ret.Inst.SyslogSeverityOne = d.Get("syslog_severity_one").(int)
	ret.Inst.TacacsServerUpDown = d.Get("tacacs_server_up_down").(int)
	//omit uuid
	return ret
}
