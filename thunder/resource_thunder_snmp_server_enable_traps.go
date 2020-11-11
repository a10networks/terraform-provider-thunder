package thunder

//Thunder resource SnmpServerEnableTraps

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerEnableTraps() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsCreate,
		Update: resourceSnmpServerEnableTrapsUpdate,
		Read:   resourceSnmpServerEnableTrapsRead,
		Delete: resourceSnmpServerEnableTrapsDelete,
		Schema: map[string]*schema.Schema{
			"lldp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_change": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"resource_usage_warning": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ssl_cert_change": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ssl_cert_expire": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"system_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"connection_resource_event": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_port": {
							Type:        schema.TypeInt,
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
			"lsn": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
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
				},
			},
			"vrrp_a": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"standby": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"snmp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"linkup": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"linkdown": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"system": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
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
				},
			},
			"ssl": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_certificate_error": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"vcs": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state_change": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"routing": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bgp_established_notification": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"bgp_backward_trans_notification": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"isis": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isis_authentication_failure": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"isis_protocols_supported_mismatch": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_rejected_adjacency": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_max_area_addresses_mismatch": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_corrupted_lsp_detected": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_originating_lsp_buffer_size_mismatch": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_area_mismatch": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_lsp_too_large_to_propagate": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_own_lsp_purge": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_sequence_number_skip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_database_overload": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_attempt_to_exceed_max_sequence": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_id_len_mismatch": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_authentication_type_failure": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_version_skew": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_manual_address_drops": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"isis_adjacency_change": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf_lsdb_overflow": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ospf_nbr_state_change": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_if_state_change": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_virt_nbr_state_change": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_lsdb_approaching_overflow": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_if_auth_failure": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_virt_if_auth_failure": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_virt_if_config_error": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_virt_if_rx_bad_packet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_tx_retransmit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_virt_if_state_change": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_if_config_error": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_max_age_lsa": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_if_rx_bad_packet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_virt_if_tx_retransmit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ospf_originate_lsa": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"gslb": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"group": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"zone": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"site": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_ip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"slb": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_port_connratelimit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_selection_failure": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_group_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_conn_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_group_member_up": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_conn_resume": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_up": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_conn_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gateway_up": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_group_up": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"application_buffer_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_connratelimit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_connlimit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_group_member_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bw_rate_limit_exceed": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_disabled": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_up": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_port_connlimit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_port_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bw_rate_limit_resume": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gateway_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_up": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_port_up": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vip_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_conn_resume": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"network": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk_port_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
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

func resourceSnmpServerEnableTrapsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTraps (Inside resourceSnmpServerEnableTrapsCreate) ")

		data := dataToSnmpServerEnableTraps(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTraps --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTraps(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTraps (Inside resourceSnmpServerEnableTrapsRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTraps(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsRead(d, meta)
}

func resourceSnmpServerEnableTrapsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsRead(d, meta)
}
func dataToSnmpServerEnableTraps(d *schema.ResourceData) go_thunder.SnmpServerEnableTraps {
	var vc go_thunder.SnmpServerEnableTraps
	var c go_thunder.SnmpServerEnableTrapsInstance
	c.Lldp = d.Get("lldp").(int)
	c.All = d.Get("all").(int)

	var obj1 go_thunder.SnmpServerEnableSlbChange
	prefix1 := "slb_change.0."
	obj1.All = d.Get(prefix1 + "all").(int)
	obj1.ResourceUsageWarning = d.Get(prefix1 + "resource_usage_warning").(int)
	obj1.SslCertChange = d.Get(prefix1 + "ssl_cert_change").(int)
	obj1.SslCertExpire = d.Get(prefix1 + "ssl_cert_expire").(int)
	obj1.SystemThreshold = d.Get(prefix1 + "system_threshold").(int)
	obj1.Server = d.Get(prefix1 + "server").(int)
	obj1.Vip = d.Get(prefix1 + "vip").(int)
	obj1.ConnectionResourceEvent = d.Get(prefix1 + "connection_resource_event").(int)
	obj1.ServerPort = d.Get(prefix1 + "server_port").(int)
	obj1.VipPort = d.Get(prefix1 + "vip_port").(int)

	c.ConnectionResourceEvent = obj1

	var obj2 go_thunder.SnmpServerEnableLsn
	prefix2 := "lsn.0."
	obj2.All = d.Get(prefix2 + "all").(int)
	obj2.FixedNatPortMappingFileChange = d.Get(prefix2 + "fixed_nat_port_mapping_file_change").(int)
	obj2.PerIPPortUsageThreshold = d.Get(prefix2 + "per_ip_port_usage_threshold").(int)
	obj2.TotalPortUsageThreshold = d.Get(prefix2 + "total_port_usage_threshold").(int)
	obj2.MaxPortThreshold = d.Get(prefix2 + "max_port_threshold").(int)
	obj2.MaxIpportThreshold = d.Get(prefix2 + "max_ipport_threshold").(int)
	obj2.TrafficExceeded = d.Get(prefix2 + "traffic_exceeded").(int)

	c.FixedNatPortMappingFileChange = obj2

	var obj3 go_thunder.SnmpServerEnableVrrpA
	prefix3 := "vrrp_a.0."
	obj3.Active = d.Get(prefix3 + "active").(int)
	obj3.Standby = d.Get(prefix3 + "standby").(int)
	obj3.All = d.Get(prefix3 + "all").(int)

	c.Active = obj3

	var obj4 go_thunder.SnmpServerEnableSnmp
	prefix4 := "snmp.0."
	obj4.Linkup = d.Get(prefix4 + "linkup").(int)
	obj4.All = d.Get(prefix4 + "all").(int)
	obj4.Linkdown = d.Get(prefix4 + "linkdown").(int)

	c.Linkup = obj4

	var obj5 go_thunder.SnmpServerEnableSystem
	prefix5 := "system.0."
	obj5.All = d.Get(prefix5 + "all").(int)
	obj5.DataCPUHigh = d.Get(prefix5 + "data_cpu_high").(int)
	obj5.Power = d.Get(prefix5 + "power").(int)
	obj5.HighDiskUse = d.Get(prefix5 + "high_disk_use").(int)
	obj5.HighMemoryUse = d.Get(prefix5 + "high_memory_use").(int)
	obj5.ControlCPUHigh = d.Get(prefix5 + "control_cpu_high").(int)
	obj5.FileSysReadOnly = d.Get(prefix5 + "file_sys_read_only").(int)
	obj5.LowTemp = d.Get(prefix5 + "low_temp").(int)
	obj5.HighTemp = d.Get(prefix5 + "high_temp").(int)
	obj5.SecDisk = d.Get(prefix5 + "sec_disk").(int)
	obj5.LicenseManagement = d.Get(prefix5 + "license_management").(int)
	obj5.Start = d.Get(prefix5 + "start").(int)
	obj5.Fan = d.Get(prefix5 + "fan").(int)
	obj5.Shutdown = d.Get(prefix5 + "shutdown").(int)
	obj5.PriDisk = d.Get(prefix5 + "pri_disk").(int)
	obj5.SyslogSeverityOne = d.Get(prefix5 + "syslog_severity_one").(int)
	obj5.TacacsServerUpDown = d.Get(prefix5 + "tacacs_server_up_down").(int)
	obj5.SmpResourceEvent = d.Get(prefix5 + "smp_resource_event").(int)
	obj5.Restart = d.Get(prefix5 + "restart").(int)
	obj5.PacketDrop = d.Get(prefix5 + "packet_drop").(int)

	c.ControlCPUHigh = obj5

	var obj6 go_thunder.SnmpServerEnableSsl
	prefix6 := "ssl.0."
	obj6.ServerCertificateError = d.Get(prefix6 + "server_certificate_error").(int)

	c.ServerCertificateError = obj6

	var obj7 go_thunder.SnmpServerEnableVcs
	prefix7 := "vcs.0."
	obj7.StateChange = d.Get(prefix7 + "state_change").(int)

	c.StateChange = obj7

	var obj8 go_thunder.SnmpServerEnableRouting
	prefix8 := "routing.0."

	var obj8_1 go_thunder.SnmpServerEnableBgp
	prefix8_1 := prefix8 + "bgp.0."
	obj8_1.BgpEstablishedNotification = d.Get(prefix8_1 + "bgp_established_notification").(int)
	obj8_1.BgpBackwardTransNotification = d.Get(prefix8_1 + "bgp_backward_trans_notification").(int)

	obj8.BgpEstablishedNotification = obj8_1

	var obj8_2 go_thunder.SnmpServerEnableIsis
	prefix8_2 := prefix8 + "isis.0."
	obj8_2.IsisAuthenticationFailure = d.Get(prefix8_2 + "isis_authentication_failure").(int)
	obj8_2.IsisProtocolsSupportedMismatch = d.Get(prefix8_2 + "isis_protocols_supported_mismatch").(int)
	obj8_2.IsisRejectedAdjacency = d.Get(prefix8_2 + "isis_rejected_adjacency").(int)
	obj8_2.IsisMaxAreaAddressesMismatch = d.Get(prefix8_2 + "isis_max_area_addresses_mismatch").(int)
	obj8_2.IsisCorruptedLSPDetected = d.Get(prefix8_2 + "isis_corrupted_lsp_detected").(int)
	obj8_2.IsisOriginatingLSPBufferSizeMismatch = d.Get(prefix8_2 + "isis_originating_lsp_buffer_size_mismatch").(int)
	obj8_2.IsisAreaMismatch = d.Get(prefix8_2 + "isis_area_mismatch").(int)
	obj8_2.IsisLSPTooLargeToPropagate = d.Get(prefix8_2 + "isis_lsp_too_large_to_propagate").(int)
	obj8_2.IsisOwnLSPPurge = d.Get(prefix8_2 + "isis_own_lsp_purge").(int)
	obj8_2.IsisSequenceNumberSkip = d.Get(prefix8_2 + "isis_sequence_number_skip").(int)
	obj8_2.IsisDatabaseOverload = d.Get(prefix8_2 + "isis_database_overload").(int)
	obj8_2.IsisAttemptToExceedMaxSequence = d.Get(prefix8_2 + "isis_attempt_to_exceed_max_sequence").(int)
	obj8_2.IsisIDLenMismatch = d.Get(prefix8_2 + "isis_id_len_mismatch").(int)
	obj8_2.IsisAuthenticationTypeFailure = d.Get(prefix8_2 + "isis_authentication_type_failure").(int)
	obj8_2.IsisVersionSkew = d.Get(prefix8_2 + "isis_version_skew").(int)
	obj8_2.IsisManualAddressDrops = d.Get(prefix8_2 + "isis_manual_address_drops").(int)
	obj8_2.IsisAdjacencyChange = d.Get(prefix8_2 + "isis_adjacency_change").(int)

	obj8.IsisAuthenticationFailure = obj8_2

	var obj8_3 go_thunder.SnmpServerEnableOspf
	prefix8_3 := prefix8 + "ospf.0."
	obj8_3.OspfLsdbOverflow = d.Get(prefix8_3 + "ospf_lsdb_overflow").(int)
	obj8_3.OspfNbrStateChange = d.Get(prefix8_3 + "ospf_nbr_state_change").(int)
	obj8_3.OspfIfStateChange = d.Get(prefix8_3 + "ospf_if_state_change").(int)
	obj8_3.OspfVirtNbrStateChange = d.Get(prefix8_3 + "ospf_virt_nbr_state_change").(int)
	obj8_3.OspfLsdbApproachingOverflow = d.Get(prefix8_3 + "ospf_lsdb_approaching_overflow").(int)
	obj8_3.OspfIfAuthFailure = d.Get(prefix8_3 + "ospf_if_auth_failure").(int)
	obj8_3.OspfVirtIfAuthFailure = d.Get(prefix8_3 + "ospf_virt_if_auth_failure").(int)
	obj8_3.OspfVirtIfConfigError = d.Get(prefix8_3 + "ospf_virt_if_config_error").(int)
	obj8_3.OspfVirtIfRxBadPacket = d.Get(prefix8_3 + "ospf_virt_if_rx_bad_packet").(int)
	obj8_3.OspfTxRetransmit = d.Get(prefix8_3 + "ospf_tx_retransmit").(int)
	obj8_3.OspfVirtIfStateChange = d.Get(prefix8_3 + "ospf_virt_if_state_change").(int)
	obj8_3.OspfIfConfigError = d.Get(prefix8_3 + "ospf_if_config_error").(int)
	obj8_3.OspfMaxAgeLsa = d.Get(prefix8_3 + "ospf_max_age_lsa").(int)
	obj8_3.OspfIfRxBadPacket = d.Get(prefix8_3 + "ospf_if_rx_bad_packet").(int)
	obj8_3.OspfVirtIfTxRetransmit = d.Get(prefix8_3 + "ospf_virt_if_tx_retransmit").(int)
	obj8_3.OspfOriginateLsa = d.Get(prefix8_3 + "ospf_originate_lsa").(int)

	obj8.OspfLsdbOverflow = obj8_3

	c.Bgp = obj8

	var obj9 go_thunder.SnmpServerEnableGslb
	prefix9 := "gslb.0."
	obj9.All = d.Get(prefix9 + "all").(int)
	obj9.Group = d.Get(prefix9 + "group").(int)
	obj9.Zone = d.Get(prefix9 + "zone").(int)
	obj9.Site = d.Get(prefix9 + "site").(int)
	obj9.ServiceIP = d.Get(prefix9 + "service_ip").(int)

	c.Group = obj9

	var obj10 go_thunder.SnmpServerEnableSlb
	prefix10 := "slb.0."
	obj10.All = d.Get(prefix10 + "all").(int)
	obj10.ServerDown = d.Get(prefix10 + "server_down").(int)
	obj10.VipPortConnratelimit = d.Get(prefix10 + "vip_port_connratelimit").(int)
	obj10.ServerSelectionFailure = d.Get(prefix10 + "server_selection_failure").(int)
	obj10.ServiceGroupDown = d.Get(prefix10 + "service_group_down").(int)
	obj10.ServerConnLimit = d.Get(prefix10 + "server_conn_limit").(int)
	obj10.ServiceGroupMemberUp = d.Get(prefix10 + "service_group_member_up").(int)
	obj10.ServerConnResume = d.Get(prefix10 + "server_conn_resume").(int)
	obj10.ServiceUp = d.Get(prefix10 + "service_up").(int)
	obj10.ServiceConnLimit = d.Get(prefix10 + "service_conn_limit").(int)
	obj10.GatewayUp = d.Get(prefix10 + "gateway_up").(int)
	obj10.ServiceGroupUp = d.Get(prefix10 + "service_group_up").(int)
	obj10.ApplicationBufferLimit = d.Get(prefix10 + "application_buffer_limit").(int)
	obj10.VipConnratelimit = d.Get(prefix10 + "vip_connratelimit").(int)
	obj10.VipConnlimit = d.Get(prefix10 + "vip_connlimit").(int)
	obj10.ServiceGroupMemberDown = d.Get(prefix10 + "service_group_member_down").(int)
	obj10.ServiceDown = d.Get(prefix10 + "service_down").(int)
	obj10.BwRateLimitExceed = d.Get(prefix10 + "bw_rate_limit_exceed").(int)
	obj10.ServerDisabled = d.Get(prefix10 + "server_disabled").(int)
	obj10.ServerUp = d.Get(prefix10 + "server_up").(int)
	obj10.VipPortConnlimit = d.Get(prefix10 + "vip_port_connlimit").(int)
	obj10.VipPortDown = d.Get(prefix10 + "vip_port_down").(int)
	obj10.BwRateLimitResume = d.Get(prefix10 + "bw_rate_limit_resume").(int)
	obj10.GatewayDown = d.Get(prefix10 + "gateway_down").(int)
	obj10.VipUp = d.Get(prefix10 + "vip_up").(int)
	obj10.VipPortUp = d.Get(prefix10 + "vip_port_up").(int)
	obj10.VipDown = d.Get(prefix10 + "vip_down").(int)
	obj10.ServiceConnResume = d.Get(prefix10 + "service_conn_resume").(int)

	c.ApplicationBufferLimit = obj10

	var obj11 go_thunder.SnmpServerEnableNetwork
	prefix11 := "network.0."
	obj11.TrunkPortThreshold = d.Get(prefix11 + "trunk_port_threshold").(int)

	c.TrunkPortThreshold = obj11

	vc.Lldp = c
	return vc
}
