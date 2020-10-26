package thunder

//Thunder resource TemplatePort

import (
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplatePort() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplatePortCreate,
		Update: resourceTemplatePortUpdate,
		Read:   resourceTemplatePortRead,
		Delete: resourceTemplatePortDelete,
		Schema: map[string]*schema.Schema{
			"source_nat": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"request_rate_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"till": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sub_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"down_timer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reassign": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_port_pool_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"inband_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"slow_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"initial_slow_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dampening_flaps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"times": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dscp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rate_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"every": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"add": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dynamic_member_priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_ssl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_pool": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"request_rate_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"request_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"flap_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"restore_svc_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dest_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"decrement": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"resel_on_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"down_grace_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"del_session_on_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplatePortCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplatePort (Inside resourceTemplatePortCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplatePort(d)
		logger.Println("[INFO] received formatted data from method data to TemplatePort --")
		d.SetId(name)
		go_thunder.PostTemplatePort(client.Token, data, client.Host)

		return resourceTemplatePortRead(d, meta)

	}
	return nil
}

func resourceTemplatePortRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading TemplatePort (Inside resourceTemplatePortRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplatePort(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplatePortUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplatePort   (Inside resourceTemplatePortUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplatePort(d)
		logger.Println("[INFO] received formatted data from method data to TemplatePort ")
		d.SetId(name)
		go_thunder.PutTemplatePort(client.Token, name, data, client.Host)

		return resourceTemplatePortRead(d, meta)

	}
	return nil
}

func resourceTemplatePortDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplatePortDelete) " + name)
		err := go_thunder.DeleteTemplatePort(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplatePort(d *schema.ResourceData) go_thunder.Port {
	var vc go_thunder.Port
	var c go_thunder.PortInstance
	c.Add = d.Get("add").(int)
	c.BwRateLimit = d.Get("bw_rate_limit").(int)
	c.BwRateLimitDuration = d.Get("bw_rate_limit_duration").(int)
	c.BwRateLimitNoLogging = d.Get("bw_rate_limit_no_logging").(int)
	c.BwRateLimitResume = d.Get("bw_rate_limit_resume").(int)
	c.ConnLimit = d.Get("conn_limit").(int)
	c.ConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	c.ConnRateLimit = d.Get("conn_rate_limit").(int)
	c.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.DampeningFlaps = d.Get("dampening_flaps").(int)
	c.Decrement = d.Get("decrement").(int)
	c.DelSessionOnServerDown = d.Get("del_session_on_server_down").(int)
	c.DestNat = d.Get("dest_nat").(int)
	c.DownGracePeriod = d.Get("down_grace_period").(int)
	c.DownTimer = d.Get("down_timer").(int)
	c.Dscp = d.Get("dscp").(int)
	c.DynamicMemberPriority = d.Get("dynamic_member_priority").(int)
	c.Every = d.Get("every").(int)
	c.ExtendedStats = d.Get("extended_stats").(int)
	c.FlapPeriod = d.Get("flap_period").(int)
	c.HealthCheck = d.Get("health_check").(string)
	c.HealthCheckDisable = d.Get("health_check_disable").(int)
	c.InbandHealthCheck = d.Get("inband_health_check").(int)
	c.InitialSlowStart = d.Get("initial_slow_start").(int)
	c.Name = d.Get("name").(string)
	c.NoSsl = d.Get("no_ssl").(int)
	c.RateInterval = d.Get("rate_interval").(string)
	c.Reassign = d.Get("reassign").(int)
	c.RequestRateInterval = d.Get("request_rate_interval").(string)
	c.RequestRateLimit = d.Get("request_rate_limit").(int)
	c.RequestRateNoLogging = d.Get("request_rate_no_logging").(int)
	c.ReselOnReset = d.Get("resel_on_reset").(int)
	c.Reset = d.Get("reset").(int)
	c.RestoreSvcTime = d.Get("restore_svc_time").(int)
	c.Resume = d.Get("resume").(int)
	c.Retry = d.Get("retry").(int)
	c.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	c.SlowStart = d.Get("slow_start").(int)
	c.SourceNat = d.Get("source_nat").(string)
	c.StatsDataAction = d.Get("stats_data_action").(string)
	c.SubGroup = d.Get("sub_group").(int)
	c.TemplatePortPoolShared = d.Get("template_port_pool_shared").(string)
	c.Till = d.Get("till").(int)
	c.Times = d.Get("times").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.Weight = d.Get("weight").(int)

	vc.UUID = c
	return vc
}
