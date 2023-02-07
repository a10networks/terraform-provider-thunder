package thunder

//Thunder resource SnmpServerEnableTrapsSystem

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSystem() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsSystemCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSystemUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSystemRead,
		DeleteContext: resourceSnmpServerEnableTrapsSystemDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"data_cpu_high": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"power": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"high_disk_use": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"high_memory_use": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"control_cpu_high": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"file_sys_read_only": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"low_temp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"high_temp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sec_disk": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"license_management": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fan": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shutdown": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pri_disk": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"syslog_severity_one": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tacacs_server_up_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"smp_resource_event": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"restart": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"packet_drop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSystem (Inside resourceSnmpServerEnableTrapsSystemCreate) ")

		data := dataToSnmpServerEnableTrapsSystem(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSystem --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsSystem(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsSystemRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSystem (Inside resourceSnmpServerEnableTrapsSystemRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSystem(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSnmpServerEnableTrapsSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSystemRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSystemRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsSystem(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsSystem {
	var vc go_thunder.SnmpServerEnableTrapsSystem
	var c go_thunder.SnmpServerEnableTrapsSystemInstance
	c.All = d.Get("all").(int)
	c.DataCPUHigh = d.Get("data_cpu_high").(int)
	c.Power = d.Get("power").(int)
	c.HighDiskUse = d.Get("high_disk_use").(int)
	c.HighMemoryUse = d.Get("high_memory_use").(int)
	c.ControlCPUHigh = d.Get("control_cpu_high").(int)
	c.FileSysReadOnly = d.Get("file_sys_read_only").(int)
	c.LowTemp = d.Get("low_temp").(int)
	c.HighTemp = d.Get("high_temp").(int)
	c.SecDisk = d.Get("sec_disk").(int)
	c.LicenseManagement = d.Get("license_management").(int)
	c.Start = d.Get("start").(int)
	c.Fan = d.Get("fan").(int)
	c.Shutdown = d.Get("shutdown").(int)
	c.PriDisk = d.Get("pri_disk").(int)
	c.SyslogSeverityOne = d.Get("syslog_severity_one").(int)
	c.TacacsServerUpDown = d.Get("tacacs_server_up_down").(int)
	c.SmpResourceEvent = d.Get("smp_resource_event").(int)
	c.Restart = d.Get("restart").(int)
	c.PacketDrop = d.Get("packet_drop").(int)

	vc.All = c
	return vc
}
