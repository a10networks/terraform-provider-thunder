package thunder

//Thunder resource Server

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerCreate,
		UpdateContext: resourceServerUpdate,
		ReadContext:   resourceServerRead,
		DeleteContext: resourceServerDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_ipv6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fqdn_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"resolve_as": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"use_aam_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"trunk": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"external_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_server_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_server_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_link_cost": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"l2_health_check_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slow_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"spoofing_cache": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alternate_server": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alternate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
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
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"range": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_port": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_port_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_port_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_server_ssl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"no_ssl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_check": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_rport_health_check": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rport_health_check_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"health_check_follow_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"follow_port_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"health_check_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"support_http2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"conn_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"no_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"conn_resume": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"stats_data_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"extended_stats": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_port": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alternate": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"alternate_name": {
										Type:        schema.TypeString,
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
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
						"packet_capture_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating Server (Inside resourceServerCreate) ")
		name1 := d.Get("name").(string)
		data := dataToServer(d)
		logger.Println("[INFO] received formatted data from method data to Server --")
		d.SetId(name1)
		err := go_thunder.PostServer(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceServerRead(ctx, d, meta)

	}
	return diags
}

func resourceServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading Server (Inside resourceServerRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetServer(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying Server   (Inside resourceServerUpdate) ")
		data := dataToServer(d)
		logger.Println("[INFO] received formatted data from method data to Server ")
		err := go_thunder.PutServer(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceServerRead(ctx, d, meta)

	}
	return diags
}

func resourceServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceServerDelete) " + name1)
		err := go_thunder.DeleteServer(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToServer(d *schema.ResourceData) go_thunder.Server {
	var vc go_thunder.Server
	var c go_thunder.ServerInstance
	c.ServerInstanceName = d.Get("name").(string)
	c.ServerInstanceServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	c.ServerInstanceHost = d.Get("host").(string)
	c.ServerInstanceFqdnName = d.Get("fqdn_name").(string)
	c.ServerInstanceResolveAs = d.Get("resolve_as").(string)
	c.ServerInstanceUseAamServer = d.Get("use_aam_server").(int)
	c.ServerInstanceEthernet = d.Get("ethernet").(int)
	c.ServerInstanceTrunk = d.Get("trunk").(int)
	c.ServerInstanceAction = d.Get("action").(string)
	c.ServerInstanceExternalIP = d.Get("external_ip").(string)
	c.ServerInstanceIpv6 = d.Get("ipv6").(string)
	c.ServerInstanceTemplateServer = d.Get("template_server").(string)
	c.ServerInstanceSharedPartitionServerTemplate = d.Get("shared_partition_server_template").(int)
	c.ServerInstanceTemplateServerShared = d.Get("template_server_shared").(string)
	c.ServerInstanceTemplateLinkCost = d.Get("template_link_cost").(string)
	c.ServerInstanceHealthCheck = d.Get("health_check").(string)
	c.ServerInstanceL2HealthCheckPath = d.Get("l2_health_check_path").(string)
	c.ServerInstanceSharedPartitionHealthCheck = d.Get("shared_partition_health_check").(int)
	c.ServerInstanceHealthCheckShared = d.Get("health_check_shared").(string)
	c.ServerInstanceHealthCheckDisable = d.Get("health_check_disable").(int)
	c.ServerInstanceConnLimit = d.Get("conn_limit").(int)
	c.ServerInstanceNoLogging = d.Get("no_logging").(int)
	c.ServerInstanceConnResume = d.Get("conn_resume").(int)
	c.ServerInstanceWeight = d.Get("weight").(int)
	c.ServerInstanceSlowStart = d.Get("slow_start").(int)
	c.ServerInstanceSpoofingCache = d.Get("spoofing_cache").(int)
	c.ServerInstanceStatsDataAction = d.Get("stats_data_action").(string)
	c.ServerInstanceExtendedStats = d.Get("extended_stats").(int)

	ServerInstanceAlternateServerCount := d.Get("alternate_server.#").(int)
	c.ServerInstanceAlternateServerAlternate = make([]go_thunder.ServerInstanceAlternateServer, 0, ServerInstanceAlternateServerCount)

	for i := 0; i < ServerInstanceAlternateServerCount; i++ {
		var obj1 go_thunder.ServerInstanceAlternateServer
		prefix1 := fmt.Sprintf("alternate_server.%d.", i)
		obj1.ServerInstanceAlternateServerAlternate = d.Get(prefix1 + "alternate").(int)
		obj1.ServerInstanceAlternateServerAlternateName = d.Get(prefix1 + "alternate_name").(string)
		c.ServerInstanceAlternateServerAlternate = append(c.ServerInstanceAlternateServerAlternate, obj1)
	}

	c.ServerInstanceUserTag = d.Get("user_tag").(string)

	ServerInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.ServerInstanceSamplingEnableCounters1 = make([]go_thunder.ServerInstanceSamplingEnable, 0, ServerInstanceSamplingEnableCount)

	for i := 0; i < ServerInstanceSamplingEnableCount; i++ {
		var obj2 go_thunder.ServerInstanceSamplingEnable
		prefix2 := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.ServerInstanceSamplingEnableCounters1 = d.Get(prefix2 + "counters1").(string)
		c.ServerInstanceSamplingEnableCounters1 = append(c.ServerInstanceSamplingEnableCounters1, obj2)
	}

	ServerInstancePortListCount := d.Get("port_list.#").(int)
	c.ServerInstancePortListPortNumber = make([]go_thunder.ServerInstancePortList, 0, ServerInstancePortListCount)

	for i := 0; i < ServerInstancePortListCount; i++ {
		var obj3 go_thunder.ServerInstancePortList
		prefix3 := fmt.Sprintf("port_list.%d.", i)
		obj3.ServerInstancePortListPortNumber = d.Get(prefix3 + "port_number").(int)
		obj3.ServerInstancePortListProtocol = d.Get(prefix3 + "protocol").(string)
		obj3.ServerInstancePortListRange = d.Get(prefix3 + "range").(int)
		obj3.ServerInstancePortListTemplatePort = d.Get(prefix3 + "template_port").(string)
		obj3.ServerInstancePortListSharedPartitionPortTemplate = d.Get(prefix3 + "shared_partition_port_template").(int)
		obj3.ServerInstancePortListTemplatePortShared = d.Get(prefix3 + "template_port_shared").(string)
		obj3.ServerInstancePortListTemplateServerSsl = d.Get(prefix3 + "template_server_ssl").(string)
		obj3.ServerInstancePortListAction = d.Get(prefix3 + "action").(string)
		obj3.ServerInstancePortListNoSsl = d.Get(prefix3 + "no_ssl").(int)
		obj3.ServerInstancePortListHealthCheck = d.Get(prefix3 + "health_check").(string)
		obj3.ServerInstancePortListSharedRportHealthCheck = d.Get(prefix3 + "shared_rport_health_check").(int)
		obj3.ServerInstancePortListRportHealthCheckShared = d.Get(prefix3 + "rport_health_check_shared").(string)
		obj3.ServerInstancePortListHealthCheckFollowPort = d.Get(prefix3 + "health_check_follow_port").(int)
		obj3.ServerInstancePortListFollowPortProtocol = d.Get(prefix3 + "follow_port_protocol").(string)
		obj3.ServerInstancePortListHealthCheckDisable = d.Get(prefix3 + "health_check_disable").(int)
		obj3.ServerInstancePortListSupportHTTP2 = d.Get(prefix3 + "support_http2").(int)
		obj3.ServerInstancePortListWeight = d.Get(prefix3 + "weight").(int)
		obj3.ServerInstancePortListConnLimit = d.Get(prefix3 + "conn_limit").(int)
		obj3.ServerInstancePortListNoLogging = d.Get(prefix3 + "no_logging").(int)
		obj3.ServerInstancePortListConnResume = d.Get(prefix3 + "conn_resume").(int)
		obj3.ServerInstancePortListStatsDataAction = d.Get(prefix3 + "stats_data_action").(string)
		obj3.ServerInstancePortListExtendedStats = d.Get(prefix3 + "extended_stats").(int)

		ServerInstancePortListAlternatePortCount := d.Get(prefix3 + "alternate_port.#").(int)
		obj3.ServerInstancePortListAlternatePortAlternate = make([]go_thunder.ServerInstancePortListAlternatePort, 0, ServerInstancePortListAlternatePortCount)

		for i := 0; i < ServerInstancePortListAlternatePortCount; i++ {
			var obj3_1 go_thunder.ServerInstancePortListAlternatePort
			prefix3_1 := prefix3 + fmt.Sprintf("alternate_port.%d.", i)
			obj3_1.ServerInstancePortListAlternatePortAlternate = d.Get(prefix3_1 + "alternate").(int)
			obj3_1.ServerInstancePortListAlternatePortAlternateName = d.Get(prefix3_1 + "alternate_name").(string)
			obj3_1.ServerInstancePortListAlternatePortAlternateServerPort = d.Get(prefix3_1 + "alternate_server_port").(int)
			obj3.ServerInstancePortListAlternatePortAlternate = append(obj3.ServerInstancePortListAlternatePortAlternate, obj3_1)
		}

		var obj3_2 go_thunder.ServerInstancePortListAuthCfg
		prefix3_2 := prefix3 + "auth_cfg.0."
		obj3_2.ServerInstancePortListAuthCfgServicePrincipalName = d.Get(prefix3_2 + "service_principal_name").(string)

		obj3.ServerInstancePortListAuthCfgServicePrincipalName = obj3_2

		obj3.ServerInstancePortListUserTag = d.Get(prefix3 + "user_tag").(string)

		ServerInstancePortListSamplingEnableCount := d.Get(prefix3 + "sampling_enable.#").(int)
		obj3.ServerInstancePortListSamplingEnableCounters1 = make([]go_thunder.ServerInstancePortListSamplingEnable, 0, ServerInstancePortListSamplingEnableCount)

		for i := 0; i < ServerInstancePortListSamplingEnableCount; i++ {
			var obj3_3 go_thunder.ServerInstancePortListSamplingEnable
			prefix3_3 := prefix3 + fmt.Sprintf("sampling_enable.%d.", i)
			obj3_3.ServerInstancePortListSamplingEnableCounters1 = d.Get(prefix3_3 + "counters1").(string)
			obj3.ServerInstancePortListSamplingEnableCounters1 = append(obj3.ServerInstancePortListSamplingEnableCounters1, obj3_3)
		}

		obj3.ServerInstancePortListPacketCaptureTemplate = d.Get(prefix3 + "packet_capture_template").(string)
		c.ServerInstancePortListPortNumber = append(c.ServerInstancePortListPortNumber, obj3)
	}

	vc.ServerInstanceName = c
	return vc
}
