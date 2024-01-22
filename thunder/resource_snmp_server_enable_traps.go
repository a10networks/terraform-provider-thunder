package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTraps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps`: Enable SNMP traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsCreate,
		UpdateContext: resourceSnmpServerEnableTrapsUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRead,
		DeleteContext: resourceSnmpServerEnableTrapsDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all SNMP traps",
			},
			"gslb": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all GSLB traps",
						},
						"zone": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB zone related traps",
						},
						"site": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB site related traps",
						},
						"group": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB group related traps",
						},
						"service_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB service-ip related traps",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"lldp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp traps",
			},
			"lsn": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all LSN group traps",
						},
						"total_port_usage_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when NAT total port usage reaches the threshold (default 655350000)",
						},
						"per_ip_port_usage_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when IP total port usage reaches the threshold (default 64512)",
						},
						"max_port_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 655350000, Description: "Maximum threshold",
						},
						"max_ipport_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 64512, Description: "Maximum threshold",
						},
						"fixed_nat_port_mapping_file_change": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when fixed nat port mapping file change",
						},
						"traffic_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when NAT pool reaches the threshold",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"network": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk_port_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable network trunk-port-threshold trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"routing": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bgpestablishednotification": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpEstablishedNotification traps",
									},
									"bgpbackwardtransnotification": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpBackwardTransNotification traps",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"ax": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bgpestablishednotification": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpEstablishedNotification traps",
												},
												"bgpbackwardtransnotification": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpBackwardTransNotification traps",
												},
												"bgpprefixthresholdexceedednotification": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpPrefixThresholdExceededNotification traps",
												},
												"bgpprefixthresholdclearnotification": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpPrefixThresholdClearNotification traps",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"isis": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isisadjacencychange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAdjacencyChange traps",
									},
									"isisareamismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAreaMismatch traps",
									},
									"isisattempttoexceedmaxsequence": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAttemptToExceedMaxSequence traps",
									},
									"isisauthenticationfailure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAuthenticationFailure traps",
									},
									"isisauthenticationtypefailure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAuthenticationTypeFailure traps",
									},
									"isiscorruptedlspdetected": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisCorruptedLSPDetected traps",
									},
									"isisdatabaseoverload": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisDatabaseOverload traps",
									},
									"isisidlenmismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisIDLenMismatch traps",
									},
									"isislsptoolargetopropagate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisLSPTooLargeToPropagate traps",
									},
									"isismanualaddressdrops": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisManualAddressDrops traps",
									},
									"isismaxareaaddressesmismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisMaxAreaAddressesMismatch traps",
									},
									"isisoriginatinglspbuffersizemismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisOriginatingLSPBufferSizeMismatch traps",
									},
									"isisownlsppurge": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisOwnLSPPurge traps",
									},
									"isisprotocolssupportedmismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisProtocolsSupportedMismatch traps",
									},
									"isisrejectedadjacency": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisRejectedAdjacency traps",
									},
									"isissequencenumberskip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisSequenceNumberSkip traps",
									},
									"isisversionskew": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisVersionSkew traps",
									},
									"isislsperrordetected": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisLSPErrorDetected traps",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ospf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospfifauthfailure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfAuthFailure traps",
									},
									"ospfifconfigerror": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfConfigError traps",
									},
									"ospfifrxbadpacket": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfRxBadPacket traps",
									},
									"ospfifstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfStateChange traps",
									},
									"ospflsdbapproachingoverflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfLsdbApproachingOverflow traps",
									},
									"ospflsdboverflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfLsdbOverflow traps",
									},
									"ospfmaxagelsa": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfMaxAgeLsa traps",
									},
									"ospfnbrstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfNbrStateChange traps",
									},
									"ospforiginatelsa": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfOriginateLsa traps",
									},
									"ospftxretransmit": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfTxRetransmit traps",
									},
									"ospfvirtifauthfailure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfAuthFailure traps",
									},
									"ospfvirtifconfigerror": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfConfigError traps",
									},
									"ospfvirtifrxbadpacket": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfRxBadPacket traps",
									},
									"ospfvirtifstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfStateChange traps",
									},
									"ospfvirtiftxretransmit": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfTxRetransmit traps",
									},
									"ospfvirtnbrstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtNbrStateChange traps",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"scaleout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"infrastructure": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"all": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all infra traps",
									},
									"test_send_all_traps": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send all infra traps",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"cluster": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"election": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable election status trap",
												},
												"master_calling_re_election": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable re-election trap",
												},
												"node_status": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable active node status trap",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"service_node": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"local_device_disabled": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local device disabled trap",
												},
												"service_master": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable service-master trap",
												},
												"traffic_map_update": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable traffic map update trap",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"master_node": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"traffic_map_distribution": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Traffic-map distribution trap",
												},
												"vserver_traffic_map_update": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VServer Traffic-map trap",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"slb": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all SLB traps",
						},
						"application_buffer_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable application buffer reach limit trap",
						},
						"gateway_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server gateway up trap",
						},
						"gateway_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server gateway down trap",
						},
						"server_conn_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server connection limit trap",
						},
						"server_conn_resume": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server connection resume trap",
						},
						"server_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server up trap",
						},
						"server_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server-down trap",
						},
						"server_disabled": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server-disabled trap",
						},
						"server_selection_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server selection failure trap",
						},
						"service_conn_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service connection limit trap",
						},
						"service_conn_resume": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service connection resume trap",
						},
						"service_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-down trap",
						},
						"service_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-up trap",
						},
						"service_group_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-up trap",
						},
						"service_group_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-down trap",
						},
						"service_group_member_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-member-up trap",
						},
						"service_group_member_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-member-down trap",
						},
						"vip_connlimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual server reach conn-limit trap",
						},
						"vip_connratelimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual server reach conn-rate-limit trap",
						},
						"vip_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual server down trap",
						},
						"vip_port_connlimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual port reach conn-limit trap",
						},
						"vip_port_connratelimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual port reach conn-rate-limit trap",
						},
						"vip_port_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual port down trap",
						},
						"vip_port_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual port up trap",
						},
						"vip_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual server up trap",
						},
						"bw_rate_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server/port bandwidth rate limit exceed trap",
						},
						"bw_rate_limit_resume": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server/port bandwidth rate limit resume trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"slb_change": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all system group traps",
						},
						"resource_usage_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable partition resource usage warning trap",
						},
						"connection_resource_event": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system connection resource event trap",
						},
						"server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server create/delete trap",
						},
						"server_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server port create/delete trap",
						},
						"ssl_cert_change": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL certificate change trap",
						},
						"ssl_cert_expire": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL certificate expiring trap",
						},
						"vip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb vip create/delete trap",
						},
						"vip_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb vip-port create/delete trap",
						},
						"system_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb system threshold trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"snmp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all SNMP group traps",
						},
						"linkdown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SNMP link-down trap",
						},
						"linkup": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SNMP link-up trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ssl": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_certificate_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL server certificate error trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"system": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all system group traps",
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
						"low_temp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system low temperature trap",
						},
						"license_management": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system license management traps",
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
						"syslog_severity_one": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system syslog severity one messages trap",
						},
						"tacacs_server_up_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system TACACS monitor server up/down trap",
						},
						"start": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system start trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vcs": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state_change": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VCS state change trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"vrrp_a": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all VRRP-A group traps",
						},
						"active": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VRRP-A active trap",
						},
						"standby": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VRRP-A standby trap",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceSnmpServerEnableTrapsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTraps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTraps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTraps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTraps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSnmpServerEnableTrapsGslb1483(d []interface{}) edpt.SnmpServerEnableTrapsGslb1483 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsGslb1483
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.Zone = in["zone"].(int)
		ret.Site = in["site"].(int)
		ret.Group = in["group"].(int)
		ret.ServiceIp = in["service_ip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsLsn1484(d []interface{}) edpt.SnmpServerEnableTrapsLsn1484 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsLsn1484
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.TotalPortUsageThreshold = in["total_port_usage_threshold"].(int)
		ret.PerIpPortUsageThreshold = in["per_ip_port_usage_threshold"].(int)
		ret.MaxPortThreshold = in["max_port_threshold"].(int)
		ret.MaxIpportThreshold = in["max_ipport_threshold"].(int)
		ret.FixedNatPortMappingFileChange = in["fixed_nat_port_mapping_file_change"].(int)
		ret.TrafficExceeded = in["traffic_exceeded"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsNetwork1485(d []interface{}) edpt.SnmpServerEnableTrapsNetwork1485 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsNetwork1485
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TrunkPortThreshold = in["trunk_port_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRouting1486(d []interface{}) edpt.SnmpServerEnableTrapsRouting1486 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsRouting1486
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgp = getObjectSnmpServerEnableTrapsRoutingBgp1487(in["bgp"].([]interface{}))
		ret.Isis = getObjectSnmpServerEnableTrapsRoutingIsis1489(in["isis"].([]interface{}))
		ret.Ospf = getObjectSnmpServerEnableTrapsRoutingOspf1490(in["ospf"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRoutingBgp1487(d []interface{}) edpt.SnmpServerEnableTrapsRoutingBgp1487 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsRoutingBgp1487
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgpestablishednotification = in["bgpestablishednotification"].(int)
		ret.Bgpbackwardtransnotification = in["bgpbackwardtransnotification"].(int)
		//omit uuid
		ret.Ax = getObjectSnmpServerEnableTrapsRoutingBgpAx1488(in["ax"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRoutingBgpAx1488(d []interface{}) edpt.SnmpServerEnableTrapsRoutingBgpAx1488 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsRoutingBgpAx1488
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bgpestablishednotification = in["bgpestablishednotification"].(int)
		ret.Bgpbackwardtransnotification = in["bgpbackwardtransnotification"].(int)
		ret.Bgpprefixthresholdexceedednotification = in["bgpprefixthresholdexceedednotification"].(int)
		ret.Bgpprefixthresholdclearnotification = in["bgpprefixthresholdclearnotification"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRoutingIsis1489(d []interface{}) edpt.SnmpServerEnableTrapsRoutingIsis1489 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsRoutingIsis1489
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isisadjacencychange = in["isisadjacencychange"].(int)
		ret.Isisareamismatch = in["isisareamismatch"].(int)
		ret.Isisattempttoexceedmaxsequence = in["isisattempttoexceedmaxsequence"].(int)
		ret.Isisauthenticationfailure = in["isisauthenticationfailure"].(int)
		ret.Isisauthenticationtypefailure = in["isisauthenticationtypefailure"].(int)
		ret.Isiscorruptedlspdetected = in["isiscorruptedlspdetected"].(int)
		ret.Isisdatabaseoverload = in["isisdatabaseoverload"].(int)
		ret.Isisidlenmismatch = in["isisidlenmismatch"].(int)
		ret.Isislsptoolargetopropagate = in["isislsptoolargetopropagate"].(int)
		ret.Isismanualaddressdrops = in["isismanualaddressdrops"].(int)
		ret.Isismaxareaaddressesmismatch = in["isismaxareaaddressesmismatch"].(int)
		ret.Isisoriginatinglspbuffersizemismatch = in["isisoriginatinglspbuffersizemismatch"].(int)
		ret.Isisownlsppurge = in["isisownlsppurge"].(int)
		ret.Isisprotocolssupportedmismatch = in["isisprotocolssupportedmismatch"].(int)
		ret.Isisrejectedadjacency = in["isisrejectedadjacency"].(int)
		ret.Isissequencenumberskip = in["isissequencenumberskip"].(int)
		ret.Isisversionskew = in["isisversionskew"].(int)
		ret.Isislsperrordetected = in["isislsperrordetected"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRoutingOspf1490(d []interface{}) edpt.SnmpServerEnableTrapsRoutingOspf1490 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsRoutingOspf1490
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospfifauthfailure = in["ospfifauthfailure"].(int)
		ret.Ospfifconfigerror = in["ospfifconfigerror"].(int)
		ret.Ospfifrxbadpacket = in["ospfifrxbadpacket"].(int)
		ret.Ospfifstatechange = in["ospfifstatechange"].(int)
		ret.Ospflsdbapproachingoverflow = in["ospflsdbapproachingoverflow"].(int)
		ret.Ospflsdboverflow = in["ospflsdboverflow"].(int)
		ret.Ospfmaxagelsa = in["ospfmaxagelsa"].(int)
		ret.Ospfnbrstatechange = in["ospfnbrstatechange"].(int)
		ret.Ospforiginatelsa = in["ospforiginatelsa"].(int)
		ret.Ospftxretransmit = in["ospftxretransmit"].(int)
		ret.Ospfvirtifauthfailure = in["ospfvirtifauthfailure"].(int)
		ret.Ospfvirtifconfigerror = in["ospfvirtifconfigerror"].(int)
		ret.Ospfvirtifrxbadpacket = in["ospfvirtifrxbadpacket"].(int)
		ret.Ospfvirtifstatechange = in["ospfvirtifstatechange"].(int)
		ret.Ospfvirtiftxretransmit = in["ospfvirtiftxretransmit"].(int)
		ret.Ospfvirtnbrstatechange = in["ospfvirtnbrstatechange"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleout1491(d []interface{}) edpt.SnmpServerEnableTrapsScaleout1491 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleout1491
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Infrastructure = getObjectSnmpServerEnableTrapsScaleoutInfrastructure1492(in["infrastructure"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructure1492(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructure1492 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructure1492
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.TestSendAllTraps = in["test_send_all_traps"].(int)
		//omit uuid
		ret.Cluster = getObjectSnmpServerEnableTrapsScaleoutInfrastructureCluster1493(in["cluster"].([]interface{}))
		ret.ServiceNode = getObjectSnmpServerEnableTrapsScaleoutInfrastructureServiceNode1494(in["service_node"].([]interface{}))
		ret.MasterNode = getObjectSnmpServerEnableTrapsScaleoutInfrastructureMasterNode1495(in["master_node"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureCluster1493(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster1493 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster1493
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Election = in["election"].(int)
		ret.MasterCallingReElection = in["master_calling_re_election"].(int)
		ret.NodeStatus = in["node_status"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureServiceNode1494(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1494 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1494
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalDeviceDisabled = in["local_device_disabled"].(int)
		ret.ServiceMaster = in["service_master"].(int)
		ret.TrafficMapUpdate = in["traffic_map_update"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureMasterNode1495(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1495 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1495
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TrafficMapDistribution = in["traffic_map_distribution"].(int)
		ret.VserverTrafficMapUpdate = in["vserver_traffic_map_update"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSlb1496(d []interface{}) edpt.SnmpServerEnableTrapsSlb1496 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsSlb1496
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.ApplicationBufferLimit = in["application_buffer_limit"].(int)
		ret.GatewayUp = in["gateway_up"].(int)
		ret.GatewayDown = in["gateway_down"].(int)
		ret.ServerConnLimit = in["server_conn_limit"].(int)
		ret.ServerConnResume = in["server_conn_resume"].(int)
		ret.ServerUp = in["server_up"].(int)
		ret.ServerDown = in["server_down"].(int)
		ret.ServerDisabled = in["server_disabled"].(int)
		ret.ServerSelectionFailure = in["server_selection_failure"].(int)
		ret.ServiceConnLimit = in["service_conn_limit"].(int)
		ret.ServiceConnResume = in["service_conn_resume"].(int)
		ret.ServiceDown = in["service_down"].(int)
		ret.ServiceUp = in["service_up"].(int)
		ret.ServiceGroupUp = in["service_group_up"].(int)
		ret.ServiceGroupDown = in["service_group_down"].(int)
		ret.ServiceGroupMemberUp = in["service_group_member_up"].(int)
		ret.ServiceGroupMemberDown = in["service_group_member_down"].(int)
		ret.VipConnlimit = in["vip_connlimit"].(int)
		ret.VipConnratelimit = in["vip_connratelimit"].(int)
		ret.VipDown = in["vip_down"].(int)
		ret.VipPortConnlimit = in["vip_port_connlimit"].(int)
		ret.VipPortConnratelimit = in["vip_port_connratelimit"].(int)
		ret.VipPortDown = in["vip_port_down"].(int)
		ret.VipPortUp = in["vip_port_up"].(int)
		ret.VipUp = in["vip_up"].(int)
		ret.BwRateLimitExceed = in["bw_rate_limit_exceed"].(int)
		ret.BwRateLimitResume = in["bw_rate_limit_resume"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSlbChange1497(d []interface{}) edpt.SnmpServerEnableTrapsSlbChange1497 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsSlbChange1497
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.ResourceUsageWarning = in["resource_usage_warning"].(int)
		ret.ConnectionResourceEvent = in["connection_resource_event"].(int)
		ret.Server = in["server"].(int)
		ret.ServerPort = in["server_port"].(int)
		ret.SslCertChange = in["ssl_cert_change"].(int)
		ret.SslCertExpire = in["ssl_cert_expire"].(int)
		ret.Vip = in["vip"].(int)
		ret.VipPort = in["vip_port"].(int)
		ret.SystemThreshold = in["system_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSnmp1498(d []interface{}) edpt.SnmpServerEnableTrapsSnmp1498 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsSnmp1498
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.Linkdown = in["linkdown"].(int)
		ret.Linkup = in["linkup"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSsl1499(d []interface{}) edpt.SnmpServerEnableTrapsSsl1499 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsSsl1499
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServerCertificateError = in["server_certificate_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSystem1500(d []interface{}) edpt.SnmpServerEnableTrapsSystem1500 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsSystem1500
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.ControlCpuHigh = in["control_cpu_high"].(int)
		ret.DataCpuHigh = in["data_cpu_high"].(int)
		ret.Fan = in["fan"].(int)
		ret.FileSysReadOnly = in["file_sys_read_only"].(int)
		ret.HighDiskUse = in["high_disk_use"].(int)
		ret.HighMemoryUse = in["high_memory_use"].(int)
		ret.HighTemp = in["high_temp"].(int)
		ret.LowTemp = in["low_temp"].(int)
		ret.LicenseManagement = in["license_management"].(int)
		ret.PacketDrop = in["packet_drop"].(int)
		ret.Power = in["power"].(int)
		ret.PriDisk = in["pri_disk"].(int)
		ret.Restart = in["restart"].(int)
		ret.SecDisk = in["sec_disk"].(int)
		ret.Shutdown = in["shutdown"].(int)
		ret.SmpResourceEvent = in["smp_resource_event"].(int)
		ret.SyslogSeverityOne = in["syslog_severity_one"].(int)
		ret.TacacsServerUpDown = in["tacacs_server_up_down"].(int)
		ret.Start = in["start"].(int)
		//omit uuid
		ret.AppsGlobal = getObjectSnmpServerEnableTrapsSystemAppsGlobal1501(in["apps_global"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSystemAppsGlobal1501(d []interface{}) edpt.SnmpServerEnableTrapsSystemAppsGlobal1501 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsSystemAppsGlobal1501
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionsThreshold = in["sessions_threshold"].(int)
		ret.CpsThreshold = in["cps_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsVcs1502(d []interface{}) edpt.SnmpServerEnableTrapsVcs1502 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsVcs1502
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StateChange = in["state_change"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsVrrpA1503(d []interface{}) edpt.SnmpServerEnableTrapsVrrpA1503 {

	count1 := len(d)
	var ret edpt.SnmpServerEnableTrapsVrrpA1503
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.All = in["all"].(int)
		ret.Active = in["active"].(int)
		ret.Standby = in["standby"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSnmpServerEnableTraps(d *schema.ResourceData) edpt.SnmpServerEnableTraps {
	var ret edpt.SnmpServerEnableTraps
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Gslb = getObjectSnmpServerEnableTrapsGslb1483(d.Get("gslb").([]interface{}))
	ret.Inst.Lldp = d.Get("lldp").(int)
	ret.Inst.Lsn = getObjectSnmpServerEnableTrapsLsn1484(d.Get("lsn").([]interface{}))
	ret.Inst.Network = getObjectSnmpServerEnableTrapsNetwork1485(d.Get("network").([]interface{}))
	ret.Inst.Routing = getObjectSnmpServerEnableTrapsRouting1486(d.Get("routing").([]interface{}))
	ret.Inst.Scaleout = getObjectSnmpServerEnableTrapsScaleout1491(d.Get("scaleout").([]interface{}))
	ret.Inst.Slb = getObjectSnmpServerEnableTrapsSlb1496(d.Get("slb").([]interface{}))
	ret.Inst.SlbChange = getObjectSnmpServerEnableTrapsSlbChange1497(d.Get("slb_change").([]interface{}))
	ret.Inst.Snmp = getObjectSnmpServerEnableTrapsSnmp1498(d.Get("snmp").([]interface{}))
	ret.Inst.Ssl = getObjectSnmpServerEnableTrapsSsl1499(d.Get("ssl").([]interface{}))
	ret.Inst.System = getObjectSnmpServerEnableTrapsSystem1500(d.Get("system").([]interface{}))
	//omit uuid
	ret.Inst.Vcs = getObjectSnmpServerEnableTrapsVcs1502(d.Get("vcs").([]interface{}))
	ret.Inst.VrrpA = getObjectSnmpServerEnableTrapsVrrpA1503(d.Get("vrrp_a").([]interface{}))
	return ret
}
