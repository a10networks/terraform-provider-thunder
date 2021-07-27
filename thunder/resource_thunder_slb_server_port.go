package thunder

//Thunder resource SlbServerPort

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerPort() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbServerPortCreate,
		UpdateContext: resourceSlbServerPortUpdate,
		ReadContext:   resourceSlbServerPortRead,
		DeleteContext: resourceSlbServerPortDelete,
		Schema: map[string]*schema.Schema{
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_rport_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_follow_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_port_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"support_http2": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"no_ssl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"follow_port_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_server_ssl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"alternate_port": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alternate_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"alternate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_server_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"port_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rport_health_check_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"range": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_principal_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_port_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbServerPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbServerPort (Inside resourceSlbServerPortCreate) ")
		name2 := d.Get("port-number").(int)

		name1 := d.Get("name").(string)
		name3 := d.Get("protocol").(string)
		data := dataToSlbServerPort(d)
		logger.Println("[INFO] received formatted data from method data to SlbServerPort --")
		d.SetId(name1 + "," + strconv.Itoa(name2) + "," + name3)
		err := go_thunder.PostSlbServerPort(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbServerPortRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbServerPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbServerPort (Inside resourceSlbServerPortRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]

		logger.Println("[INFO] Fetching service Read server/port")
		data, err := go_thunder.GetSlbServerPort(client.Token, name1, name2, name3, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found slb/server/port")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbServerPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]

		logger.Println("[INFO] Modifying SlbServerPort   (Inside resourceSlbServerPortUpdate) ")
		data := dataToSlbServerPort(d)
		logger.Println("[INFO] received formatted data from method data to SlbServerPort ")
		err := go_thunder.PutSlbServerPort(client.Token, name1, name2, name3, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbServerPortRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbServerPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]

		logger.Println("[INFO] Deleting instance (Inside resourceSlbServerPortDelete) ")
		err := go_thunder.DeleteSlbServerPort(client.Token, name1, name2, name3, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToSlbServerPort(d *schema.ResourceData) go_thunder.SlbServerPort {
	var vc go_thunder.SlbServerPort
	var c go_thunder.SlbServerPortInstance
	c.HealthCheckDisable = d.Get("health_check_disable").(int)
	c.Protocol = d.Get("protocol").(string)
	c.Weight = d.Get("weight").(int)
	c.SharedRportHealthCheck = d.Get("shared_rport_health_check").(int)
	c.StatsDataAction = d.Get("stats_data_action").(string)
	c.HealthCheckFollowPort = d.Get("health_check_follow_port").(int)
	c.TemplatePort = d.Get("template_port").(string)
	c.ConnLimit = d.Get("conn_limit").(int)
	c.TemplatePortShared = d.Get("template_port_shared").(string)
	c.SupportHTTP2 = d.Get("support_http2").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SlbServerSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.SlbServerSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.NoSsl = d.Get("no_ssl").(int)
	c.FollowPortProtocol = d.Get("follow_port_protocol").(string)
	c.TemplateServerSsl = d.Get("template_server_ssl").(string)

	AlternatePortCount := d.Get("alternate_port.#").(int)
	c.AlternateName = make([]go_thunder.SlbServerAlternatePort, 0, AlternatePortCount)

	for i := 0; i < AlternatePortCount; i++ {
		var obj2 go_thunder.SlbServerAlternatePort
		prefix := fmt.Sprintf("alternate_port.%d.", i)
		obj2.AlternateName = d.Get(prefix + "alternate_name").(string)
		obj2.Alternate = d.Get(prefix + "alternate").(int)
		obj2.AlternateServerPort = d.Get(prefix + "alternate_server_port").(int)
		c.AlternateName = append(c.AlternateName, obj2)
	}

	c.PortNumber = d.Get("port_number").(int)
	c.ExtendedStats = d.Get("extended_stats").(int)
	c.RportHealthCheckShared = d.Get("rport_health_check_shared").(string)
	c.ConnResume = d.Get("conn_resume").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.Range = d.Get("range").(int)

	var obj3 go_thunder.SlbServerAuthCfg
	prefix := "auth_cfg.0."
	obj3.ServicePrincipalName = d.Get(prefix + "service_principal_name").(string)

	c.ServicePrincipalName = obj3

	c.Action = d.Get("action").(string)
	c.NoLogging = d.Get("no_logging").(int)
	c.HealthCheck = d.Get("health_check").(string)
	c.SharedPartitionPortTemplate = d.Get("shared_partition_port_template").(int)

	vc.HealthCheckDisable = c
	return vc
}
