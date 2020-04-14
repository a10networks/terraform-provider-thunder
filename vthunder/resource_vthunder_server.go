package vthunder

//vThunder resource Server

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Update: resourceServerUpdate,
		Read:   resourceServerRead,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
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
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": {
							Type:        schema.TypeString,
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
						"range": {
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
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_rport_health_check": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"conn_resume": {
							Type:        schema.TypeInt,
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
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"follow_port_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"health_check_follow_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"no_logging": {
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
									"alternate_server_port": {
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
						"port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						//						"auth_cfg": {
						//							Type:     schema.TypeList,
						//							Optional: true,
						//							MaxItems: 1,
						//							Elem: &schema.Resource{
						//								Schema: map[string]*schema.Schema{
						//									"service_principal_name": {
						//										Type:        schema.TypeString,
						//										Optional:    true,
						//										Description: "",
						//									},
						//								},
						//							},
						//						},
						"conn_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_port": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_server_ssl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"resolve_as": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_ipv6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"spoofing_cache": {
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
			"external_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
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
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"fqdn_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating Server (Inside resourceServerCreate    " + name)
		v := dataToServer(name, d)
		logger.Println("[INFO] received formatted data from method data to sg --" + v.Name.Name)
		d.SetId(name)
		go_vthunder.PostServer(client.Token, v, client.Host)

		return resourceServerRead(d, meta)
	}
	return nil
}

func resourceServerRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Server (Inside resourceServerRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching service Read" + name)

		server, err := go_vthunder.GetServer(client.Token, name, client.Host)

		if server == nil {
			logger.Println("[INFO] No server found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying server   (Inside resourceServerUpdate    " + name)
		v := dataToServer(name, d)
		logger.Println("[INFO] received formatted data from method data to sg --" + v.Name.Name)
		d.SetId(name)
		go_vthunder.PutServer(client.Token, name, v, client.Host)

		return resourceServerRead(d, meta)
	}
	return nil
}

func resourceServerDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting server (Inside resourceServerDelete) " + name)

		err := go_vthunder.DeleteServer(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete service group  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate Server structure
func dataToServer(name string, d *schema.ResourceData) go_vthunder.Server {
	//	logger := util.GetLoggerInstance()
	var s go_vthunder.Server

	var sInstance go_vthunder.ServerInstance

	sInstance.HealthCheckDisable = d.Get("health_check_disable").(int)
	sInstance.StatsDataAction = d.Get("stats_data_action").(string)
	sInstance.SlowStart = d.Get("slow_start").(int)
	sInstance.Weight = d.Get("weight").(int)
	sInstance.SpoofingCache = d.Get("spoofing_cache").(int)
	sInstance.ResolveAs = d.Get("resolve_as").(string)
	sInstance.ConnLimit = d.Get("conn_limit").(int)
	sInstance.UUID = d.Get("uuid").(string)
	sInstance.FqdnName = d.Get("fqdn_name").(string)
	sInstance.ExternalIP = d.Get("external_ip").(string)
	sInstance.HealthCheckShared = d.Get("health_check_shared").(string)
	sInstance.Ipv6 = d.Get("ipv6").(string)
	sInstance.TemplateServer = d.Get("template_server").(string)
	sInstance.ServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	sInstance.SharedPartitionHealthCheck = d.Get("shared_partition_health_check").(int)
	sInstance.Host = d.Get("host").(string)
	sInstance.ExtendedStats = d.Get("extended_stats").(int)
	sInstance.ConnResume = d.Get("conn_resume").(int)
	sInstance.Name = d.Get("name").(string)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.Action = d.Get("action").(string)
	sInstance.HealthCheck = d.Get("health_check").(string)
	sInstance.NoLogging = d.Get("no_logging").(int)

	portListCount := d.Get("port_list.#").(int)
	sInstance.PortNumber = make([]go_vthunder.PortLists, 0, portListCount)
	for i := 0; i < portListCount; i++ {
		var pl go_vthunder.PortLists
		prefix := fmt.Sprintf("port_list.%d", i)
		pl.HealthCheckDisable = d.Get(prefix + ".health_check_disable").(int)
		pl.Protocol = d.Get(prefix + ".protocol").(string)
		pl.Weight = d.Get(prefix + ".weight").(int)
		pl.SharedRportHealthCheck = d.Get(prefix + ".shared_rport_health_check").(int)
		pl.StatsDataAction = d.Get(prefix + ".stats_data_action").(string)
		pl.HealthCheckFollowPort = d.Get(prefix + ".health_check_follow_port").(int)
		pl.TemplatePort = d.Get(prefix + ".template_port").(string)
		pl.ConnLimit = d.Get(prefix + ".conn_limit").(int)
		pl.UUID = d.Get(prefix + ".uuid").(string)
		pl.SupportHTTP2 = d.Get(prefix + ".support_http2").(int)
		pl.NoSsl = d.Get(prefix + ".no_ssl").(int)
		pl.FollowPortProtocol = d.Get(prefix + ".follow_port_protocol").(string)
		pl.TemplateServerSsl = d.Get(prefix + ".template_server_ssl").(string)
		pl.PortNumber = d.Get(prefix + ".port_number").(int)
		pl.ExtendedStats = d.Get(prefix + ".extended_stats").(int)
		pl.RportHealthCheckShared = d.Get(prefix + ".rport_health_check_shared").(string)
		pl.ConnResume = d.Get(prefix + ".conn_resume").(int)
		pl.UserTag = d.Get(prefix + ".user_tag").(string)
		pl.Range = d.Get(prefix + ".range").(int)
		pl.Action = d.Get(prefix + ".action").(string)
		pl.HealthCheck = d.Get(prefix + ".health_check").(string)
		pl.NoLogging = d.Get(prefix + ".no_logging").(int)

		sampleCount := d.Get(prefix + ".sampling_enable.#").(int)
		pl.Counters1 = make([]go_vthunder.SamplingEnable, sampleCount, sampleCount)

		for x := 0; x < sampleCount; x++ {
			var s go_vthunder.SamplingEnable
			mapEntity(d.Get(fmt.Sprintf("%s.sampling_enable.%d", prefix, x)).(map[string]interface{}), &s)
			pl.Counters1[x] = s
		}

		altPortCt := d.Get(prefix + ".alternate_port.#").(int)
		pl.AlternateName = make([]go_vthunder.AlternatePort, altPortCt, altPortCt)

		for x := 0; x < altPortCt; x++ {
			var alt go_vthunder.AlternatePort
			mapEntity(d.Get(fmt.Sprintf("%s.alternate_port.%d", prefix, x)).(map[string]interface{}), &alt)
			pl.AlternateName[x] = alt
		}

		sInstance.PortNumber = append(sInstance.PortNumber, pl)
	}

	samplingCount := d.Get("sampling_enable.#").(int)
	sInstance.Counters1 = make([]go_vthunder.SamplingEnable, 0, samplingCount)
	for i := 0; i < samplingCount; i++ {
		var sm go_vthunder.SamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		sm.Counters1 = d.Get(prefix + ".counters1").(string)

		sInstance.Counters1 = append(sInstance.Counters1, sm)
	}

	altServerCt := d.Get("alternate_server.#").(int)
	sInstance.AlternateName = make([]go_vthunder.AlternateServer, 0, altServerCt)

	for i := 0; i < altServerCt; i++ {
		var altSer go_vthunder.AlternateServer
		prefix := fmt.Sprintf("alternate_server.%d", i)
		altSer.Alternate = d.Get(prefix + ".alternate").(int)
		altSer.AlternateName = d.Get(prefix + ".alternate_name").(string)

		sInstance.AlternateName = append(sInstance.AlternateName, altSer)
	}

	s.Name = sInstance

	return s
}
