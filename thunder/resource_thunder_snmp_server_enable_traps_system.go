package thunder

//Thunder resource SnmpServerEnableTrapsSystem

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsSystemCreate,
		Update: resourceSnmpServerEnableTrapsSystemUpdate,
		Read:   resourceSnmpServerEnableTrapsSystemRead,
		Delete: resourceSnmpServerEnableTrapsSystemDelete,
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

func resourceSnmpServerEnableTrapsSystemCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSystem (Inside resourceSnmpServerEnableTrapsSystemCreate) ")

		data := dataToSnmpServerEnableTrapsSystem(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSystem --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsSystem(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsSystemRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsSystemRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSystem (Inside resourceSnmpServerEnableTrapsSystemRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSystem(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsSystemUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSystemRead(d, meta)
}

func resourceSnmpServerEnableTrapsSystemDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSystemRead(d, meta)
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
