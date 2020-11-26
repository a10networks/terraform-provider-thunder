package thunder

//Thunder resource SnmpServerEnableTrapsLsn

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerEnableTrapsLsn() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsLsnCreate,
		Update: resourceSnmpServerEnableTrapsLsnUpdate,
		Read:   resourceSnmpServerEnableTrapsLsnRead,
		Delete: resourceSnmpServerEnableTrapsLsnDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fixed_nat_port_mapping_file_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"per_ip_port_usage_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"total_port_usage_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_port_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_ipport_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_exceeded": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsLsnCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsLsn (Inside resourceSnmpServerEnableTrapsLsnCreate) ")

		data := dataToSnmpServerEnableTrapsLsn(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsLsn --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsLsn(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsLsnRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsLsnRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsLsn (Inside resourceSnmpServerEnableTrapsLsnRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsLsn(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsLsnUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsLsnRead(d, meta)
}

func resourceSnmpServerEnableTrapsLsnDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsLsnRead(d, meta)
}
func dataToSnmpServerEnableTrapsLsn(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsLsn {
	var vc go_thunder.SnmpServerEnableTrapsLsn
	var c go_thunder.SnmpServerEnableTrapsLsnInstance
	c.All = d.Get("all").(int)
	c.FixedNatPortMappingFileChange = d.Get("fixed_nat_port_mapping_file_change").(int)
	c.PerIPPortUsageThreshold = d.Get("per_ip_port_usage_threshold").(int)
	c.TotalPortUsageThreshold = d.Get("total_port_usage_threshold").(int)
	c.MaxPortThreshold = d.Get("max_port_threshold").(int)
	c.MaxIpportThreshold = d.Get("max_ipport_threshold").(int)
	c.TrafficExceeded = d.Get("traffic_exceeded").(int)

	vc.All = c
	return vc
}
