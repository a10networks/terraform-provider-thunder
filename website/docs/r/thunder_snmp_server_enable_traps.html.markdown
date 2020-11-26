---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps"
description: |-
	Provides details about thunder snmp server enable traps resource for A10
---

# thunder\_snmp\_server\_enable\_traps

`thunder_snmp_server_enable_traps` Provides details about thunder snmp server enable traps
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_snmp_server_enable_traps" "Snmp_Server_Enable_Traps_Test" {
        lldp = 0
all = 0
slb_change {  
        all =  0 
        resource_usage_warning =  0 
        uuid =  "string" 
        ssl_cert_change =  0 
        ssl_cert_expire =  0 
        system_threshold =  0 
        server =  0 
        vip =  0 
        connection_resource_event =  0 
        server_port =  0 
        vip_port =  0 
        }
uuid = "string"
lsn {  
        all =  0 
        fixed_nat_port_mapping_file_change =  0 
        per_ip_port_usage_threshold =  0 
        uuid =  "string" 
        total_port_usage_threshold =  0 
        max_port_threshold =  0 
        max_ipport_threshold =  0 
        traffic_exceeded =  0 
        }
vrrp_a {  
        active =  0 
        standby =  0 
        all =  0 
        uuid =  "string" 
        }
snmp {  
        linkup =  0 
        all =  0 
        linkdown =  0 
        uuid =  "string" 
        }
system {  
        all =  0 
        data_cpu_high =  0 
        uuid =  "string" 
        power =  0 
        high_disk_use =  0 
        high_memory_use =  0 
        control_cpu_high =  0 
        file_sys_read_only =  0 
        low_temp =  0 
        high_temp =  0 
        sec_disk =  0 
        license_management =  0 
        start =  0 
        fan =  0 
        shutdown =  0 
        pri_disk =  0 
        syslog_severity_one =  0 
        tacacs_server_up_down =  0 
        smp_resource_event =  0 
        restart =  0 
        packet_drop =  0 
        }
ssl {  
        server_certificate_error =  0 
        uuid =  "string" 
        }
vcs {  
        state_change =  0 
        uuid =  "string" 
        }
routing {  
 bgp {  
        bgp_established_notification =  0 
        uuid =  "string" 
        bgp_backward_trans_notification =  0 
        }
isis {  
        isis_authentication_failure =  0 
        uuid =  "string" 
        isis_protocols_supported_mismatch =  0 
        isis_rejected_adjacency =  0 
        isis_max_area_addresses_mismatch =  0 
        isis_corrupted_lsp_detected =  0 
        isis_originating_lsp_buffer_size_mismatch =  0 
        isis_area_mismatch =  0 
        isis_lsp_too_large_to_propagate =  0 
        isis_own_lsp_purge =  0 
        isis_sequence_number_skip =  0 
        isis_database_overload =  0 
        isis_attempt_to_exceed_max_sequence =  0 
        isis_id_len_mismatch =  0 
        isis_authentication_type_failure =  0 
        isis_version_skew =  0 
        isis_manual_address_drops =  0 
        isis_adjacency_change =  0 
        }
ospf {  
        ospf_lsdb_overflow =  0 
        uuid =  "string" 
        ospf_nbr_state_change =  0 
        ospf_if_state_change =  0 
        ospf_virt_nbr_state_change =  0 
        ospf_lsdb_approaching_overflow =  0 
        ospf_if_auth_failure =  0 
        ospf_virt_if_auth_failure =  0 
        ospf_virt_if_config_error =  0 
        ospf_virt_if_rx_bad_packet =  0 
        ospf_tx_retransmit =  0 
        ospf_virt_if_state_change =  0 
        ospf_if_config_error =  0 
        ospf_max_age_lsa =  0 
        ospf_if_rx_bad_packet =  0 
        ospf_virt_if_tx_retransmit =  0 
        ospf_originate_lsa =  0 
        }
        }
gslb {  
        all =  0 
        group =  0 
        uuid =  "string" 
        zone =  0 
        site =  0 
        service_ip =  0 
        }
slb {  
        all =  0 
        server_down =  0 
        vip_port_conn_rate_limit =  0 
        server_selection_failure =  0 
        service_group_down =  0 
        server_conn_limit =  0 
        service_group_member_up =  0 
        uuid =  "string" 
        server_conn_resume =  0 
        service_up =  0 
        service_conn_limit =  0 
        gateway_up =  0 
        service_group_up =  0 
        application_buffer_limit =  0 
        vip_conn_rate_limit =  0 
        vip_conn_limit =  0 
        service_group_member_down =  0 
        service_down =  0 
        bw_rate_limit_exceed =  0 
        server_disabled =  0 
        server_up =  0 
        vip_port_conn_limit =  0 
        vip_port_down =  0 
        bw_rate_limit_resume =  0 
        gateway_down =  0 
        vip_up =  0 
        vip_port_up =  0 
        vip_down =  0 
        service_conn_resume =  0 
        }
network {  
        trunk_port_threshold =  0 
        uuid =  "string" 
        }
 
}

```

## Argument Reference

* `all` - Enable all SLB traps
* `lldp` - Enable lldp traps
* `uuid` - uuid of the object
* `connection_resource_event` - Enable system connection resource event trap
* `resource_usage_warning` - Enable partition resource usage warning trap
* `server` - Enable slb server create/delete trap
* `server_port` - Enable slb server port create/delete trap
* `ssl_cert_change` - Enable SSL certificate change trap
* `ssl_cert_expire` - Enable SSL certificate expiring trap
* `system_threshold` - Enable slb system threshold trap
* `vip` - Enable slb vip create/delete trap
* `vip_port` - Enable slb vip-port create/delete trap
* `fixed_nat_port_mapping_file_change` - Enable LSN trap when fixed nat port mapping file change
* `max_ipport_threshold` - Maximum threshold
* `max_port_threshold` - Maximum threshold
* `per_ip_port_usage_threshold` - Enable LSN trap when IP total port usage reaches the threshold (default 64512)
* `total_port_usage_threshold` - Enable LSN trap when NAT total port usage reaches the threshold (default 655350000)
* `traffic_exceeded` - Enable LSN trap when NAT pool reaches the threshold
* `active` - Enable VRRP-A active trap
* `standby` - Enable VRRP-A standby trap
* `linkdown` - Enable SNMP link-down trap
* `linkup` - Enable SNMP link-up trap
* `control_cpu_high` - Enable control CPU usage high trap
* `data_cpu_high` - Enable data CPU usage high trap
* `fan` - Enable system fan trap
* `file_sys_read_only` - Enable file system read-only trap
* `high_disk_use` - Enable system high disk usage trap
* `high_memory_use` - Enable system high memory usage trap
* `high_temp` - Enable system high temperature trap
* `license_management` - Enable system license management traps
* `low_temp` - Enable system low temperature trap
* `packet_drop` - Enable system packet dropped trap
* `power` - Enable system power supply trap
* `pri_disk` - Enable system primary hard disk trap
* `restart` - Enable system restart trap
* `sec_disk` - Enable system secondary hard disk trap
* `shutdown` - Enable system shutdown trap
* `smp_resource_event` - Enable system smp resource event trap
* `start` - Enable system start trap
* `tacacs_server_up_down` - Enable system TACACS monitor server up/down trap
* `server_certificate_error` - Enable SSL server certificate error trap
* `state_change` - Enable VCS state change trap
* `bgp_backward_trans_notification` - Enable bgpBackwardTransNotification traps
* `bgp_established_notification` - Enable bgp_established_notification traps
* `ospf_if_auth_failure` - Enable ospfIfAuthFailure traps
* `ospf_if_config_error` - Enable ospfIfConfigError traps
* `ospf_if_rx_bad_packet` - Enable ospfIfRxBadPacket traps
* `ospf_if_state_change` - Enable ospfIfStateChange traps
* `ospf_lsdb_approaching_overflow` - Enable ospfLsdbApproachingOverflow traps
* `ospf_lsdb_overflow` - Enable ospfLsdbOverflow traps
* `ospf_max_age_lsa` - Enable ospfMaxAgeLsa traps
* `ospf_nbr_state_change` - Enable ospfNbrStateChange traps
* `ospf_originate_lsa` - Enable ospfOriginateLsa traps
* `ospf_tx_retransmit` - Enable ospfTxRetransmit traps
* `ospf_virt_if_auth_failure` - Enable ospfVirtIfAuthFailure traps
* `ospf_virt_if_config_error` - Enable ospfVirtIfConfigError traps
* `ospf_virt_if_rx_bad_packet` - Enable ospfVirtIfRxBadPacket traps
* `ospf_virt_if_state_change` - Enable ospfVirtIfStateChange traps
* `ospf_virt_if_tx_retransmit` - Enable ospfVirtIfTxRetransmit traps
* `ospf_virt_nbr_state_change` - Enable ospfVirtNbrStateChange traps
* `isis_adjacency_change` - Enable isisAdjacencyChange traps
* `isis_area_mismatch` - Enable isisAreaMismatch traps
* `isis_attempt_to_exceed_max_sequence` - Enable isisAttemptToExceedMaxSequence traps
* `isis_authentication_failure` - Enable isisAuthenticationFailure traps
* `isis_authentication_type_failure` - Enable isisAuthenticationTypeFailure traps
* `isis_corrupted_lsp_detected` - Enable isisCorruptedLSPDetected traps
* `isis_database_overload` - Enable isisDatabaseOverload traps
* `isis_id_len_mismatch` - Enable isisIDLenMismatch traps
* `isis_lsp_too_large_to_propagate` - Enable isisLSPTooLargeToPropagate traps
* `isis_manual_address_drops` - Enable isisManualAddressDrops traps
* `isis_max_area_addresses_mismatch` - Enable isisMaxAreaAddressesMismatch traps
* `isis_originating_lsp_buffer_size_mismatch` - Enable isisOriginatingLSPBufferSizeMismatch traps
* `isis_own_lsp_purge` - Enable isisOwnLSPPurge traps
* `isis_protocols_supported_mismatch` - Enable isisProtocolsSupportedMismatch traps
* `isis_rejected_adjacency` - Enable isisRejectedAdjacency traps
* `isis_sequence_number_skip` - Enable isisSequenceNumberSkip traps
* `isis_version_skew` - Enable isisVersionSkew traps
* `group` - Enable GSLB group related traps
* `service_ip` - Enable GSLB service-ip related traps
* `site` - Enable GSLB site related traps
* `zone` - Enable GSLB zone related traps
* `application_buffer_limit` - Enable application buffer reach limit trap
* `bw_rate_limit_exceed` - Enable SLB server/port bandwidth rate limit exceed trap
* `bw_rate_limit_resume` - Enable SLB server/port bandwidth rate limit resume trap
* `gateway_down` - Enable SLB server gateway down trap
* `gateway_up` - Enable SLB server gateway up trap
* `server_conn_limit` - Enable SLB server connection limit trap
* `server_conn_resume` - Enable SLB server connection resume trap
* `server_disabled` - Enable SLB server-disabled trap
* `server_down` - Enable SLB server-down trap
* `server_selection_failure` - Enable SLB server selection failure trap
* `server_up` - Enable slb server up trap
* `service_conn_limit` - Enable SLB service connection limit trap
* `service_conn_resume` - Enable SLB service connection resume trap
* `service_down` - Enable SLB service-down trap
* `service_group_down` - Enable SLB service-group-down trap
* `service_group_member_down` - Enable SLB service-group-member-down trap
* `service_group_member_up` - Enable SLB service-group-member-up trap
* `service_group_up` - Enable SLB service-group-up trap
* `service_up` - Enable SLB service-up trap
* `vip_connlimit` - Enable the virtual server reach conn-limit trap
* `vip_connratelimit` - Enable the virtual server reach conn-rate-limit trap
* `vip_down` - Enable SLB virtual server down trap
* `vip_port_connlimit` - Enable the virtual port reach conn-limit trap
* `vip_port_connratelimit` - Enable the virtual port reach conn-rate-limit trap
* `vip_port_down` - Enable SLB virtual port down trap
* `vip_port_up` - Enable SLB virtual port up trap
* `vip_up` - Enable SLB virtual server up trap
* `trunk_port_threshold` - Enable network trunk-port-threshold trap

