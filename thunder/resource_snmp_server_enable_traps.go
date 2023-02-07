package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSnmpServerEnableTraps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps`: Enable SNMP traps\n\n",
		CreateContext: resourceSnmpServerEnableTrapsCreate,
		UpdateContext: resourceSnmpServerEnableTrapsUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRead,
		DeleteContext: resourceSnmpServerEnableTrapsDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all SNMP traps",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"gslb": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all GSLB traps",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"zone": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB zone related traps [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"site": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB site related traps [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"group": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB group related traps [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB service-ip related traps [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"lldp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp traps [shared partition only]",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"lsn": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all LSN group traps",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"total_port_usage_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when NAT total port usage reaches the threshold (default 655350000)",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"per_ip_port_usage_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when IP total port usage reaches the threshold (default 64512)",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"max_port_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 655350000, Description: "Maximum threshold",
							ValidateFunc: validation.IntBetween(10000, 655355000),
						},
						"max_ipport_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 64512, Description: "Maximum threshold",
							ValidateFunc: validation.IntBetween(10000, 64512),
						},
						"fixed_nat_port_mapping_file_change": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when fixed nat port mapping file change",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"traffic_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable LSN trap when NAT pool reaches the threshold",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable network trunk-port-threshold trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpEstablishedNotification traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"bgpbackwardtransnotification": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable bgpBackwardTransNotification traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"isis": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isisadjacencychange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAdjacencyChange traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisareamismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAreaMismatch traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisattempttoexceedmaxsequence": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAttemptToExceedMaxSequence traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisauthenticationfailure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAuthenticationFailure traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisauthenticationtypefailure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAuthenticationTypeFailure traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isiscorruptedlspdetected": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisCorruptedLSPDetected traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisdatabaseoverload": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisDatabaseOverload traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisidlenmismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisIDLenMismatch traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isislsptoolargetopropagate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisLSPTooLargeToPropagate traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isismanualaddressdrops": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisManualAddressDrops traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isismaxareaaddressesmismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisMaxAreaAddressesMismatch traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisoriginatinglspbuffersizemismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisOriginatingLSPBufferSizeMismatch traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisownlsppurge": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisOwnLSPPurge traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisprotocolssupportedmismatch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisProtocolsSupportedMismatch traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisrejectedadjacency": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisRejectedAdjacency traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isissequencenumberskip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisSequenceNumberSkip traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"isisversionskew": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisVersionSkew traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
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
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfAuthFailure traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfifconfigerror": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfConfigError traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfifrxbadpacket": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfRxBadPacket traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfifstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfStateChange traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospflsdbapproachingoverflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfLsdbApproachingOverflow traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospflsdboverflow": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfLsdbOverflow traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfmaxagelsa": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfMaxAgeLsa traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfnbrstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfNbrStateChange traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospforiginatelsa": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfOriginateLsa traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospftxretransmit": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfTxRetransmit traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfvirtifauthfailure": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfAuthFailure traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfvirtifconfigerror": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfConfigError traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfvirtifrxbadpacket": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfRxBadPacket traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfvirtifstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfStateChange traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfvirtiftxretransmit": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfTxRetransmit traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"ospfvirtnbrstatechange": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtNbrStateChange traps [shared partition only]",
										ValidateFunc: validation.IntBetween(0, 1),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"test_send_all_traps": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send all infra traps",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"cluster": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"single_node_mode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable single node status trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"election": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable election status trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"master_calling_re_election": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable re-election trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"node_status": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable active node status trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
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
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local device disabled trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"service_master": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable service-master trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"traffic_map_update": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable traffic map update trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
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
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Traffic-map distribution trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"vserver_traffic_map_update": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VServer Traffic-map trap [shared partition only]",
													ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"application_buffer_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable application buffer reach limit trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"gateway_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server gateway up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"gateway_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server gateway down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server_conn_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server connection limit trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server_conn_resume": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server connection resume trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server-down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server_disabled": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server-disabled trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server_selection_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server selection failure trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_conn_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service connection limit trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_conn_resume": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service connection resume trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_group_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_group_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_group_member_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-member-up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"service_group_member_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-member-down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_connlimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual server reach conn-limit trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_connratelimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual server reach conn-rate-limit trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual server down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_port_connlimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual port reach conn-limit trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_port_connratelimit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual port reach conn-rate-limit trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_port_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual port down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_port_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual port up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_up": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual server up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"bw_rate_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server/port bandwidth rate limit exceed trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"bw_rate_limit_resume": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server/port bandwidth rate limit resume trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"resource_usage_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable partition resource usage warning trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"connection_resource_event": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system connection resource event trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server create/delete trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server port create/delete trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"ssl_cert_change": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL certificate change trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"ssl_cert_expire": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL certificate expiring trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb vip create/delete trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vip_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb vip-port create/delete trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"system_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb system threshold trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"linkdown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SNMP link-down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"linkup": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SNMP link-up trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSL server certificate error trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all system group traps [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"control_cpu_high": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable control CPU usage high trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"data_cpu_high": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable data CPU usage high trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"fan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system fan trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"file_sys_read_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable file system read-only trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"high_disk_use": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system high disk usage trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"high_memory_use": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system high memory usage trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"high_temp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system high temperature trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"low_temp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system low temperature trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"license_management": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system license management traps [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"packet_drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system packet dropped trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"power": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system power supply trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"pri_disk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system primary hard disk trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"restart": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system restart trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"sec_disk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system secondary hard disk trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"shutdown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system shutdown trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"smp_resource_event": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system smp resource event trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"syslog_severity_one": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system syslog severity one messages trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"tacacs_server_up_down": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system TACACS monitor server up/down trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"start": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable system start trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VCS state change trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"active": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VRRP-A active trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"standby": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VRRP-A standby trap [shared partition only]",
							ValidateFunc: validation.IntBetween(0, 1),
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

func getObjectSnmpServerEnableTrapsGslb(d []interface{}) edpt.SnmpServerEnableTrapsGslb {
	var ret edpt.SnmpServerEnableTrapsGslb
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.All = in["all"].(int)
		ret.Zone = in["zone"].(int)
		ret.Site = in["site"].(int)
		ret.Group = in["group"].(int)
		ret.ServiceIp = in["service_ip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsLsn(d []interface{}) edpt.SnmpServerEnableTrapsLsn {
	var ret edpt.SnmpServerEnableTrapsLsn
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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

func getObjectSnmpServerEnableTrapsNetwork(d []interface{}) edpt.SnmpServerEnableTrapsNetwork {
	var ret edpt.SnmpServerEnableTrapsNetwork
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.TrunkPortThreshold = in["trunk_port_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRouting(d []interface{}) edpt.SnmpServerEnableTrapsRouting {
	var ret edpt.SnmpServerEnableTrapsRouting
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bgp = getObjectSnmpServerEnableTrapsRoutingBgp(in["bgp"].([]interface{}))
		ret.Isis = getObjectSnmpServerEnableTrapsRoutingIsis(in["isis"].([]interface{}))
		ret.Ospf = getObjectSnmpServerEnableTrapsRoutingOspf(in["ospf"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRoutingBgp(d []interface{}) edpt.SnmpServerEnableTrapsRoutingBgp {
	var ret edpt.SnmpServerEnableTrapsRoutingBgp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bgpestablishednotification = in["bgpestablishednotification"].(int)
		ret.Bgpbackwardtransnotification = in["bgpbackwardtransnotification"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRoutingIsis(d []interface{}) edpt.SnmpServerEnableTrapsRoutingIsis {
	var ret edpt.SnmpServerEnableTrapsRoutingIsis
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsRoutingOspf(d []interface{}) edpt.SnmpServerEnableTrapsRoutingOspf {
	var ret edpt.SnmpServerEnableTrapsRoutingOspf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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

func getObjectSnmpServerEnableTrapsScaleout(d []interface{}) edpt.SnmpServerEnableTrapsScaleout {
	var ret edpt.SnmpServerEnableTrapsScaleout
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Infrastructure = getObjectSnmpServerEnableTrapsScaleoutInfrastructure(in["infrastructure"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructure(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructure {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructure
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.All = in["all"].(int)
		ret.TestSendAllTraps = in["test_send_all_traps"].(int)
		//omit uuid
		ret.Cluster = getObjectSnmpServerEnableTrapsScaleoutInfrastructureCluster(in["cluster"].([]interface{}))
		ret.ServiceNode = getObjectSnmpServerEnableTrapsScaleoutInfrastructureServiceNode(in["service_node"].([]interface{}))
		ret.MasterNode = getObjectSnmpServerEnableTrapsScaleoutInfrastructureMasterNode(in["master_node"].([]interface{}))
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureCluster(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureCluster
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SingleNodeMode = in["single_node_mode"].(int)
		ret.Election = in["election"].(int)
		ret.MasterCallingReElection = in["master_calling_re_election"].(int)
		ret.NodeStatus = in["node_status"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureServiceNode(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureServiceNode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.LocalDeviceDisabled = in["local_device_disabled"].(int)
		ret.ServiceMaster = in["service_master"].(int)
		ret.TrafficMapUpdate = in["traffic_map_update"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsScaleoutInfrastructureMasterNode(d []interface{}) edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode {
	var ret edpt.SnmpServerEnableTrapsScaleoutInfrastructureMasterNode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.TrafficMapDistribution = in["traffic_map_distribution"].(int)
		ret.VserverTrafficMapUpdate = in["vserver_traffic_map_update"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSlb(d []interface{}) edpt.SnmpServerEnableTrapsSlb {
	var ret edpt.SnmpServerEnableTrapsSlb
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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

func getObjectSnmpServerEnableTrapsSlbChange(d []interface{}) edpt.SnmpServerEnableTrapsSlbChange {
	var ret edpt.SnmpServerEnableTrapsSlbChange
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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

func getObjectSnmpServerEnableTrapsSnmp(d []interface{}) edpt.SnmpServerEnableTrapsSnmp {
	var ret edpt.SnmpServerEnableTrapsSnmp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.All = in["all"].(int)
		ret.Linkdown = in["linkdown"].(int)
		ret.Linkup = in["linkup"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSsl(d []interface{}) edpt.SnmpServerEnableTrapsSsl {
	var ret edpt.SnmpServerEnableTrapsSsl
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.ServerCertificateError = in["server_certificate_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsSystem(d []interface{}) edpt.SnmpServerEnableTrapsSystem {
	var ret edpt.SnmpServerEnableTrapsSystem
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
	}
	return ret
}

func getObjectSnmpServerEnableTrapsVcs(d []interface{}) edpt.SnmpServerEnableTrapsVcs {
	var ret edpt.SnmpServerEnableTrapsVcs
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.StateChange = in["state_change"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSnmpServerEnableTrapsVrrpA(d []interface{}) edpt.SnmpServerEnableTrapsVrrpA {
	var ret edpt.SnmpServerEnableTrapsVrrpA
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
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
	ret.Inst.Gslb = getObjectSnmpServerEnableTrapsGslb(d.Get("gslb").([]interface{}))
	ret.Inst.Lldp = d.Get("lldp").(int)
	ret.Inst.Lsn = getObjectSnmpServerEnableTrapsLsn(d.Get("lsn").([]interface{}))
	ret.Inst.Network = getObjectSnmpServerEnableTrapsNetwork(d.Get("network").([]interface{}))
	ret.Inst.Routing = getObjectSnmpServerEnableTrapsRouting(d.Get("routing").([]interface{}))
	ret.Inst.Scaleout = getObjectSnmpServerEnableTrapsScaleout(d.Get("scaleout").([]interface{}))
	ret.Inst.Slb = getObjectSnmpServerEnableTrapsSlb(d.Get("slb").([]interface{}))
	ret.Inst.SlbChange = getObjectSnmpServerEnableTrapsSlbChange(d.Get("slb_change").([]interface{}))
	ret.Inst.Snmp = getObjectSnmpServerEnableTrapsSnmp(d.Get("snmp").([]interface{}))
	ret.Inst.Ssl = getObjectSnmpServerEnableTrapsSsl(d.Get("ssl").([]interface{}))
	ret.Inst.System = getObjectSnmpServerEnableTrapsSystem(d.Get("system").([]interface{}))
	//omit uuid
	ret.Inst.Vcs = getObjectSnmpServerEnableTrapsVcs(d.Get("vcs").([]interface{}))
	ret.Inst.VrrpA = getObjectSnmpServerEnableTrapsVrrpA(d.Get("vrrp_a").([]interface{}))
	return ret
}
