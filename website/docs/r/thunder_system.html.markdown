---
layout: "thunder"
page_title: "thunder: thunder_system"
sidebar_current: "docs-thunder-resource-system"
description: |-
    Provides details about thunder system resource for A10
---

# thunder\_system

`thunder_system` Provides details about thunder system
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_system" "resourceSystemTest" {
	anomaly_log = 0
attack_log = 0
ddos_attack = 0
ddos_log = 0
sockstress_disable = 0
promiscuous_mode = 0
glid = 0
module_ctrl_cpu = "string"
src_ip_hash_enable = 0
class_list_hitcount_enable = 0
geo_db_hitcount_enable = 0
domain_list_hitcount_enable = 0
dynamic_service_dns_socket_pool = 0
system_chassis_port_split_enable = 0
uuid = "string"
timeout_value {  
 	ftp =  0 
	scp =  0 
	sftp =  0 
	tftp =  0 
	http =  0 
	https =  0 
	uuid =  "string" 
	}
bandwidth {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
counter_lib_accounting {  
 	uuid =  "string" 
	}
control_cpu {  
 	uuid =  "string" 
	}
data_cpu {  
 	uuid =  "string" 
	}
mgmt_port {  
 	port_index =  0 
	mac_address =  "string" 
	pci_address =  "string" 
	}
shared_poll_mode {  
 	enable =  0 
	disable =  0 
	}
probe_network_devices {  
 	}
management_interface_mode {  
 	dedicated =  0 
	non_dedicated =  0 
	}
set_tcp_syn_per_sec {  
 	tcp_syn_value =  0 
	uuid =  "string" 
	}
add_port {  
 	port_index =  0 
	}
del_port {  
 	port_index =  0 
	}
modify_port {  
 	port_index =  0 
	port_number =  0 
	}
tls_1_3_mgmt {  
 	enable =  0 
	uuid =  "string" 
	}
multi_queue_support {  
 	enable =  0 
	}
add_cpu_core {  
 	core_index =  0 
	}
delete_cpu_core {  
 	core_index =  0 
	}
cpu_hyper_thread {  
 	enable =  0 
	disable =  0 
	}
io_cpu {  
 	max_cores =  0 
	}
link_monitor {  
 	enable =  0 
	disable =  0 
	}
lro {  
 	enable =  0 
	disable =  0 
	}
tso {  
 	enable =  0 
	disable =  0 
	}
port_list {  
 	uuid =  "string" 
	}
port_info {  
 	uuid =  "string" 
	}
inuse_port_list {  
 	uuid =  "string" 
	}
cpu_list {  
 	uuid =  "string" 
	}
cpu_map {  
 	uuid =  "string" 
	}
inuse_cpu_list {  
 	uuid =  "string" 
	}
set_rxtx_desc_size {  
 	port_index =  0 
	rxd_size =  0 
	txd_size =  0 
	}
set_rxtx_queue {  
 	port_index =  0 
	rxq_size =  0 
	txq_size =  0 
	}
template {  
 	template_policy =  "string" 
	uuid =  "string" 
	}
template_bind {  
 monitor-list {   
	id =  0 
clear-cfg {   
	sessions =  "string" 
	clear_all_sequence =  0 
	clear_sequence =  0 
	}
link-disable-cfg {   
	diseth =  0 
	dis_sequence =  0 
	}
link-enable-cfg {   
	enaeth =  0 
	ena_sequence =  0 
	}
	monitor_relation =  "string" 
link-up-cfg {   
	linkup_ethernet1 =  0 
	link_up_sequence1 =  0 
	linkup_ethernet2 =  0 
	link_up_sequence2 =  0 
	linkup_ethernet3 =  0 
	link_up_sequence3 =  0 
	}
link-down-cfg {   
	linkdown_ethernet1 =  0 
	link_down_sequence1 =  0 
	linkdown_ethernet2 =  0 
	link_down_sequence2 =  0 
	linkdown_ethernet3 =  0 
	link_down_sequence3 =  0 
	}
	uuid =  "string" 
	user_tag =  "string" 
	}
	}
mon_template {  
 monitor-list {   
	id =  0 
clear-cfg {   
	sessions =  "string" 
	clear_all_sequence =  0 
	clear_sequence =  0 
	}
link-disable-cfg {   
	diseth =  0 
	dis_sequence =  0 
	}
link-enable-cfg {   
	enaeth =  0 
	ena_sequence =  0 
	}
	monitor_relation =  "string" 
link-up-cfg {   
	linkup_ethernet1 =  0 
	link_up_sequence1 =  0 
	linkup_ethernet2 =  0 
	link_up_sequence2 =  0 
	linkup_ethernet3 =  0 
	link_up_sequence3 =  0 
	}
link-down-cfg {   
	linkdown_ethernet1 =  0 
	link_down_sequence1 =  0 
	linkdown_ethernet2 =  0 
	link_down_sequence2 =  0 
	linkdown_ethernet3 =  0 
	link_down_sequence3 =  0 
	}
	uuid =  "string" 
	user_tag =  "string" 
	}
link_block_as_down {  
 	enable =  0 
	uuid =  "string" 
	}
link_down_on_restart {  
 	enable =  0 
	uuid =  "string" 
	}
	}
memory {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
resource_usage {  
 	ssl_context_memory =  0 
	ssl_dma_memory =  0 
	nat_pool_addr_count =  0 
	l4_session_count =  0 
	auth_portal_html_file_size =  0 
	auth_portal_image_file_size =  0 
	max_aflex_file_size =  0 
	aflex_table_entry_count =  0 
	class_list_ipv6_addr_count =  0 
	class_list_ac_entry_count =  0 
	class_list_entry_count =  0 
	max_aflex_authz_collection_number =  0 
	radius_table_size =  0 
	authz_policy_number =  0 
	ipsec_sa_number =  0 
	ram_cache_memory_limit =  0 
	waf_template_count =  0 
	auth_session_count =  0 
	uuid =  "string" 
visibility {  
 	monitored_entity_count =  0 
	uuid =  "string" 
	}
	}
link_capability {  
 	enable =  0 
	uuid =  "string" 
	}
resource_accounting {  
 	uuid =  "string" 
template-list {   
	name =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
app_resources {  
 gslb_device_cfg {  
 	gslb_device_max =  0 
	gslb_device_min_guarantee =  0 
	}
gslb_geo_location_cfg {  
 	gslb_geo_location_max =  0 
	gslb_geo_location_min_guarantee =  0 
	}
gslb_ip_list_cfg {  
 	gslb_ip_list_max =  0 
	gslb_ip_list_min_guarantee =  0 
	}
gslb_policy_cfg {  
 	gslb_policy_max =  0 
	gslb_policy_min_guarantee =  0 
	}
gslb_service_cfg {  
 	gslb_service_max =  0 
	gslb_service_min_guarantee =  0 
	}
gslb_service_ip_cfg {  
 	gslb_service_ip_max =  0 
	gslb_service_ip_min_guarantee =  0 
	}
gslb_service_port_cfg {  
 	gslb_service_port_max =  0 
	gslb_service_port_min_guarantee =  0 
	}
gslb_site_cfg {  
 	gslb_site_max =  0 
	gslb_site_min_guarantee =  0 
	}
gslb_svc_group_cfg {  
 	gslb_svc_group_max =  0 
	gslb_svc_group_min_guarantee =  0 
	}
gslb_template_cfg {  
 	gslb_template_max =  0 
	gslb_template_min_guarantee =  0 
	}
gslb_zone_cfg {  
 	gslb_zone_max =  0 
	gslb_zone_min_guarantee =  0 
	}
health_monitor_cfg {  
 	health_monitor_max =  0 
	health_monitor_min_guarantee =  0 
	}
real_port_cfg {  
 	real_port_max =  0 
	real_port_min_guarantee =  0 
	}
real_server_cfg {  
 	real_server_max =  0 
	real_server_min_guarantee =  0 
	}
service_group_cfg {  
 	service_group_max =  0 
	service_group_min_guarantee =  0 
	}
virtual_server_cfg {  
 	virtual_server_max =  0 
	virtual_server_min_guarantee =  0 
	}
virtual_port_cfg {  
 	virtual_port_max =  0 
	virtual_port_min_guarantee =  0 
	}
cache_template_cfg {  
 	cache_template_max =  0 
	cache_template_min_guarantee =  0 
	}
client_ssl_template_cfg {  
 	client_ssl_template_max =  0 
	client_ssl_template_min_guarantee =  0 
	}
conn_reuse_template_cfg {  
 	conn_reuse_template_max =  0 
	conn_reuse_template_min_guarantee =  0 
	}
fast_tcp_template_cfg {  
 	fast_tcp_template_max =  0 
	fast_tcp_template_min_guarantee =  0 
	}
fast_udp_template_cfg {  
 	fast_udp_template_max =  0 
	fast_udp_template_min_guarantee =  0 
	}
fix_template_cfg {  
 	fix_template_max =  0 
	fix_template_min_guarantee =  0 
	}
http_template_cfg {  
 	http_template_max =  0 
	http_template_min_guarantee =  0 
	}
link_cost_template_cfg {  
 	link_cost_template_max =  0 
	link_cost_template_min_guarantee =  0 
	}
persist_cookie_template_cfg {  
 	persist_cookie_template_max =  0 
	persist_cookie_template_min_guarantee =  0 
	}
persist_srcip_template_cfg {  
 	persist_srcip_template_max =  0 
	persist_srcip_template_min_guarantee =  0 
	}
server_ssl_template_cfg {  
 	server_ssl_template_max =  0 
	server_ssl_template_min_guarantee =  0 
	}
proxy_template_cfg {  
 	proxy_template_max =  0 
	proxy_template_min_guarantee =  0 
	}
stream_template_cfg {  
 	stream_template_max =  0 
	stream_template_min_guarantee =  0 
	}
	threshold =  0 
	uuid =  "string" 
	}
network_resources {  
 static_ipv4_route_cfg {  
 	static_ipv4_route_max =  0 
	static_ipv4_route_min_guarantee =  0 
	}
static_ipv6_route_cfg {  
 	static_ipv6_route_max =  0 
	static_ipv6_route_min_guarantee =  0 
	}
ipv4_acl_line_cfg {  
 	ipv4_acl_line_max =  0 
	ipv4_acl_line_min_guarantee =  0 
	}
ipv6_acl_line_cfg {  
 	ipv6_acl_line_max =  0 
	ipv6_acl_line_min_guarantee =  0 
	}
static_arp_cfg {  
 	static_arp_max =  0 
	static_arp_min_guarantee =  0 
	}
static_neighbor_cfg {  
 	static_neighbor_max =  0 
	static_neighbor_min_guarantee =  0 
	}
static_mac_cfg {  
 	static_mac_max =  0 
	static_mac_min_guarantee =  0 
	}
object_group_cfg {  
 	object_group_max =  0 
	object_group_min_guarantee =  0 
	}
object_group_clause_cfg {  
 	object_group_clause_max =  0 
	object_group_clause_min_guarantee =  0 
	}
	threshold =  0 
	uuid =  "string" 
	}
system_resources {  
 bw_limit_cfg {  
 	bw_limit_max =  0 
	bw_limit_watermark_disable =  0 
	}
concurrent_session_limit_cfg {  
 	concurrent_session_limit_max =  0 
	}
l4_session_limit_cfg {  
 	l4_session_limit_max =  "string" 
	l4_session_limit_min_guarantee =  "string" 
	}
l4cps_limit_cfg {  
 	l4cps_limit_max =  0 
	}
l7cps_limit_cfg {  
 	l7cps_limit_max =  0 
	}
natcps_limit_cfg {  
 	natcps_limit_max =  0 
	}
fwcps_limit_cfg {  
 	fwcps_limit_max =  0 
	}
ssl_throughput_limit_cfg {  
 	ssl_throughput_limit_max =  0 
	ssl_throughput_limit_watermark_disable =  0 
	}
sslcps_limit_cfg {  
 	sslcps_limit_max =  0 
	}
	threshold =  0 
	uuid =  "string" 
	}
	}
	}
trunk {  
 load_balance {  
 	use_l3 =  0 
	use_l4 =  0 
	uuid =  "string" 
	}
	}
ports {  
 	link_detection_interval =  0 
	uuid =  "string" 
	}
table_integrity {  
 	table =  "string" 
	audit_action =  "string" 
	auto_sync_action =  "string" 
	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
ipsec {  
 	packet_round_robin =  0 
	crypto_core =  0 
	crypto_mem =  0 
	qat =  0 
	uuid =  "string" 
fpga_decrypt {  
 	action =  "string" 
	}
	}
spe_profile {  
 	action =  "string" 
	}
spe_status {  
 	uuid =  "string" 
	}
ssl_status {  
 	uuid =  "string" 
	}
deep_hrxq {  
 	enable =  0 
	}
hrxq_status {  
 	uuid =  "string" 
	}
cpu_load_sharing {  
 	disable =  0 
packets_per_second {  
 	min =  0 
	}
cpu_usage {  
 	low =  0 
	high =  0 
	}
	tcp =  0 
	udp =  0 
	uuid =  "string" 
	}
per_vlan_limit {  
 	bcast =  0 
	ipmcast =  0 
	mcast =  0 
	unknown_ucast =  0 
	uuid =  "string" 
	}
all_vlan_limit {  
 	bcast =  0 
	ipmcast =  0 
	mcast =  0 
	unknown_ucast =  0 
	uuid =  "string" 
	}
ve_mac_scheme {  
 	ve_mac_scheme_val =  "string" 
	uuid =  "string" 
	}
session_reclaim_limit {  
 	nscan_limit =  0 
	scan_freq =  0 
	uuid =  "string" 
	}
ssl_scv {  
 	enable =  0 
	uuid =  "string" 
	}
ssl_scv_verify_host {  
 	disable =  0 
	uuid =  "string" 
	}
ssl_scv_verify_crl_sign {  
 	enable =  0 
	uuid =  "string" 
	}
ssl_set_compatible_cipher {  
 	disable =  0 
	uuid =  "string" 
	}
hardware {  
 	uuid =  "string" 
	}
platformtype {  
 	uuid =  "string" 
	}
reboot {  
 	uuid =  "string" 
	}
shutdown {  
 	uuid =  "string" 
	}
environment {  
 	uuid =  "string" 
	}
hardware_forward {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
slb {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
	}
throughput {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
ipmi {  
 	reset =  0 
ip {  
 	ipv4_address =  "string" 
	ipv4_netmask =  "string" 
	default_gateway =  "string" 
	}
ipsrc {  
 	dhcp =  0 
	static =  0 
	}
user {  
 	add =  "string" 
	password =  "string" 
	administrator =  0 
	callback =  0 
	operator =  0 
	user =  0 
	disable =  "string" 
	privilege =  "string" 
	setname =  "string" 
	newname =  "string" 
	setpass =  "string" 
	newpass =  "string" 
	}
tool {  
 	cmd =  "string" 
	}
	}
queuing_buffer {  
 	enable =  0 
	uuid =  "string" 
	}
high_memory_l4_session {  
 	enable =  0 
	uuid =  "string" 
	}
trunk_hw_hash {  
 	mode =  0 
	uuid =  "string" 
	}
trunk_xaui_hw_hash {  
 	mode =  0 
	uuid =  "string" 
	}
upgrade_status {  
 	uuid =  "string" 
	}
guest_file {  
 	uuid =  "string" 
	}
cm_update_file_name_ref {  
 	source_name =  "string" 
	dest_name =  "string" 
	id =  0 
	}
core {  
 	uuid =  "string" 
	}
apps_global {  
 	log_session_on_established =  0 
	msl_time =  0 
	uuid =  "string" 
	}
shell_privileges {  
 	enable_shell_privileges =  0 
	uuid =  "string" 
	}
cosq_stats {  
 	uuid =  "string" 
	}
cosq_show {  
 	uuid =  "string" 
	}
shm_logging {  
 	enable =  0 
	uuid =  "string" 
	}
fw {  
 	application_mempool =  0 
	application_flow =  0 
	basic_dpi_enable =  0 
	uuid =  "string" 
	}
password_policy {  
 	complexity =  "string" 
	aging =  "string" 
	history =  "string" 
	min_pswd_len =  0 
	uuid =  "string" 
	}
radius {  
 server {  
 	listen_port =  0 
remote {  
 ip-list {   
	ip_list_name =  "string" 
	ip_list_secret =  0 
	ip_list_secret_string =  "string" 
	}
	}
	secret =  0 
	secret_string =  "string" 
	vrid =  0 
attribute {   
	attribute_value =  "string" 
	prefix_length =  "string" 
	prefix_vendor =  0 
	prefix_number =  0 
	name =  "string" 
	value =  "string" 
	custom_vendor =  0 
	custom_number =  0 
	vendor =  0 
	number =  0 
	}
	disable_reply =  0 
	accounting_start =  "string" 
	accounting_stop =  "string" 
	accounting_interim_update =  "string" 
	accounting_on =  "string" 
	attribute_name =  "string" 
	custom_attribute_name =  "string" 
	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
	}
geoloc-list-list {   
	name =  "string" 
	shared =  0 
include-geoloc-name-list {   
	include_geoloc_name_val =  "string" 
	}
exclude-geoloc-name-list {   
	exclude_geoloc_name_val =  "string" 
	}
	uuid =  "string" 
	user_tag =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
geoloc_name_helper {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
geolocation_file {  
 	uuid =  "string" 
error_info {  
 	uuid =  "string" 
	}
	}
geoloc {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
geo_location {  
 	geo_location_iana =  0 
	geo_location_geolite2_city =  0 
	geolite2_city_include_ipv6 =  0 
	geo_location_geolite2_country =  0 
	geolite2_country_include_ipv6 =  0 
geoloc-load-file-list {   
	geo_location_load_filename =  "string" 
	template_name =  "string" 
	}
	uuid =  "string" 
entry-list {   
	geo_locn_obj_name =  "string" 
geo-locn-multiple-addresses {   
	first_ip_address =  "string" 
	geol_ipv4_mask =  "string" 
	ip_addr2 =  "string" 
	first_ipv6_address =  "string" 
	geol_ipv6_mask =  0 
	ipv6_addr2 =  "string" 
	}
	uuid =  "string" 
	user_tag =  "string" 
	}
	}
tcp_syn_per_sec {  
 	uuid =  "string" 
	}
ip_threat_list {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
ipv4_source_list {  
 class-list-cfg {   
	class_list =  "string" 
	ip_threat_action_tmpl =  0 
	}
	uuid =  "string" 
	}
ipv4_dest_list {  
 class-list-cfg {   
	class_list =  "string" 
	ip_threat_action_tmpl =  0 
	}
	uuid =  "string" 
	}
ipv6_source_list {  
 class-list-cfg {   
	class_list =  "string" 
	ip_threat_action_tmpl =  0 
	}
	uuid =  "string" 
	}
ipv6_dest_list {  
 class-list-cfg {   
	class_list =  "string" 
	ip_threat_action_tmpl =  0 
	}
	uuid =  "string" 
	}
ipv4_internet_host_list {  
 class-list-cfg {   
	class_list =  "string" 
	ip_threat_action_tmpl =  0 
	}
	uuid =  "string" 
	}
ipv6_internet_host_list {  
 class-list-cfg {   
	class_list =  "string" 
	ip_threat_action_tmpl =  0 
	}
	uuid =  "string" 
	}
	}
fpga_drop {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
dpdk_stats {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
fpga_core_crc {  
 	monitor_disable =  0 
	reboot_enable =  0 
	uuid =  "string" 
	}
mfa_management {  
 	enable =  0 
	uuid =  "string" 
	}
mfa_validation_type {  
 	ca_cert =  "string" 
	uuid =  "string" 
	}
mfa_cert_store {  
 	cert_host =  "string" 
	protocol =  "string" 
	cert_store_path =  "string" 
	username =  "string" 
	passwd_string =  "string" 
	uuid =  "string" 
	}
mfa_auth {  
 	username =  "string" 
	second_factor =  "string" 
	}
q_in_q {  
 	inner_tpid =  "string" 
	outer_tpid =  "string" 
	enable_all_ports =  0 
	uuid =  "string" 
	}
port_count {  
 	port_count_kernel =  0 
	port_count_hm =  0 
	port_count_logging =  0 
	port_count_alg =  0 
	uuid =  "string" 
	}
health-check-list {   
	l2hm_hc_name =  "string" 
	method_l2bfd =  0 
	l2bfd_tx_interval =  0 
	l2bfd_rx_interval =  0 
	l2bfd_multiplier =  0 
	uuid =  "string" 
	user_tag =  "string" 
	}
path-list {   
	l2hm_path_name =  "string" 
	l2hm_vlan =  0 
	l2hm_setup_test_api =  0 
	ifpair_eth_start =  0 
	ifpair_eth_end =  0 
	ifpair_trunk_start =  0 
	ifpair_trunk_end =  0 
	l2hm_attach =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
	}
psu_info {  
 	uuid =  "string" 
	}
gui_image_list {  
 	uuid =  "string" 
	}
syslog_time_msec {  
 	enable_flag =  0 
	}
ipmi_service {  
 	disable =  0 
	uuid =  "string" 
	}
app_performance {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
ssl_req_q {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
tcp {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
rate_limit_reset_unknown_conn {  
 	pkt_rate_for_reset_unknown_conn =  0 
	log_for_reset_unknown_conn =  0 
	uuid =  "string" 
	}
	}
icmp {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
icmp6 {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
ip_stats {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
ip6_stats {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
domain_list_info {  
 	uuid =  "string" 
	}
ip_dns_cache {  
 	uuid =  "string" 
	}
bfd {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
icmp_rate {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
job_offload {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
dns {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
dns_cache {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
session {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
ndisc_ra {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
tcp_stats {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
telemetry_log {  
 top_k_source_list {  
 	uuid =  "string" 
	}
top_k_app_svc_list {  
 	uuid =  "string" 
	}
device_status {  
 	uuid =  "string" 
	}
environment {  
 	uuid =  "string" 
	}
partition_metrics {  
 	uuid =  "string" 
	}
	}
 
}

```

## Argument Reference

* `system` - Configure System Parameters
* `anomaly-log` - log system anomalies
* `attack-log` - log attack anomalies
* `ddos-attack` - System DDoS Attack
* `ddos-log` - log DDoS attack anomalies
* `sockstress-disable` - Disable sockstress protection
* `promiscuous-mode` - Run in promiscous mode settings
* `glid` - Apply limits to the whole system
* `module-ctrl-cpu` - 'high': high cpu usage; 'low': low cpu usage; 'medium': medium cpu usage;
* `src-ip-hash-enable` - Enable source ip hash
* `class-list-hitcount-enable` - Enable class list hit count
* `geo-db-hitcount-enable` - Enable Geolocation database hit count
* `domain-list-hitcount-enable` - Enable class list hit count
* `dynamic-service-dns-socket-pool` - Enable socket pool for dynamic-service DNS
* `system-chassis-port-split-enable` - Enable port split for the chassis
* `uuid` - uuid of the object
* `ftp` - set timeout to stop ftp transfer in seconds, 0 is no limit
* `scp` - set timeout to stop scp transfer in seconds, 0 is no limit
* `sftp` - set timeout to stop sftp transfer in seconds, 0 is no limit
* `tftp` - set timeout to stop tftp transfer in seconds, 0 is no limit
* `http` - set timeout to stop http transfer in seconds, 0 is no limit
* `https` - set timeout to stop https transfer in seconds, 0 is no limit
* `counters1` - 'all': all; 'connattempt': Connect initiated; 'connects': Connect established; 'drops': Connect dropped; 'conndrops': Embryonic connect dropped; 'closed': Connect closed; 'segstimed': Segs to get RTT; 'rttupdated': Update RTT; 'delack': Delayed acks sent; 'timeoutdrop': Conn dropped in rxmt timeout; 'rexmttimeo': Retransmit timeout; 'persisttimeo': Persist timeout; 'keeptimeo': Keepalive timeout; 'keepprobe': Keepalive probe sent; 'keepdrops': Connect dropped in keepalive; 'sndtotal': Total packet sent; 'sndpack': Data packet sent; 'sndbyte': Data bytes sent; 'sndrexmitpack': Data packet retransmit; 'sndrexmitbyte': Data byte retransmit; 'sndrexmitbad': Unnecessary packet retransmit; 'sndacks': Ack packet sent; 'sndprobe': Window probe sent; 'sndurg': URG packet sent; 'sndwinup': Window update packet sent; 'sndctrl': SYN|FIN|RST packet sent; 'sndrst': RST packet sent; 'sndfin': FIN packet sent; 'sndsyn': SYN packet sent; 'rcvtotal': Total packet received; 'rcvpack': Packet received; 'rcvbyte': Bytes received; 'rcvbadoff': Packet received with bad offset; 'rcvmemdrop': Packet dropped for lack of memory; 'rcvduppack': Duplicate packet received; 'rcvdupbyte': Duplicate bytes received; 'rcvpartduppack': Packet with some duplicate data; 'rcvpartdupbyte': Dup. bytes in part-dup. packets; 'rcvoopack': Out-of-order packet received; 'rcvoobyte': Out-of-order bytes received; 'rcvpackafterwin': Packets with data after window; 'rcvbyteafterwin': Bytes rcvd after window; 'rcvwinprobe': Rcvd window probe packet; 'rcvdupack': Rcvd duplicate acks; 'rcvacktoomuch': Rcvd acks for unsent data; 'rcvackpack': Rcvd ack packets; 'rcvackbyte': Bytes acked by rcvd acks; 'rcvwinupd': Rcvd window update packets; 'pawsdrop': Segments dropped due to PAWS; 'predack': Hdr predict for acks; 'preddat': Hdr predict for data pkts; 'persistdrop': Timeout in persist state; 'badrst': Ignored RST; 'finwait2_drops': Drop FIN_WAIT_2 connection after time limit; 'sack_recovery_episode': SACK recovery episodes; 'sack_rexmits': SACK rexmit segments; 'sack_rexmit_bytes': SACK rexmit bytes; 'sack_rcv_blocks': SACK received; 'sack_send_blocks': SACK sent; 'sndcack': Challenge ACK sent; 'cacklim': Challenge ACK limited; 'reassmemdrop': Packet dropped during reassembly; 'reasstimeout': Reassembly Time Out; 'cc_idle': Congestion control window set do to idle; 'cc_reduce': Congestion control window reduced by event; 'rcvdsack': Rcvd DSACK packets; 'a2brcvwnd': ATCP to BTCP receive window; 'a2bsackpresent': ATCP to BTCP SACK options present; 'a2bdupack': ATCP to BTCP Dup/OO ACK; 'a2brxdata': ATCP to BTCP Rxmitted data; 'a2btcpoptions': ATCP to BTCP unsupported TCP options; 'a2boodata': ATCP to BTCP oo data received; 'a2bpartialack': ATCP to BTCP partial ack received; 'a2bfsmtransition': ATCP to BTCP state machine transition; 'a2btransitionnum': ATCP to BTCP total transitions; 'b2atransitionnum': ATCP to BTCP total transitions; 'bad_iochan': IO Channel Modified; 'atcpforward': Adaptive TCP forward; 'atcpsent': Adaptive TCP sent; 'atcprexmitsadrop': Adaptive TCP transmit SA drops; 'atcpsendbackack': Adaptive TCP sendback ACK; 'atcprexmit': Adaptive TCP retransmits; 'atcpbuffallocfail': Adaptive TCP buffer allocation fails; 'a2bappbuffering': Transition to full stack on when application buffers too much data; 'atcpsendfail': Adaptive TCP sent fails; 'earlyrexmit': Early Retransmission sent; 'mburstlim': Maxburst limited tx; 'a2bsndwnd': ATCP to BTCP send window; 'proxyheaderv1': Proxy header v1; 'proxyheaderv2': Proxy header v2;
* `port-index` - port index to be configured (Specify port index)
* `mac-address` - mac-address to be configured as mgmt port
* `pci-address` - pci-address to be configured as mgmt port
* `enable` - Enable 2FA for management plane
* `disable` - Disable IPMI on platform
* `dedicated` - Set management interface in dedicated mode
* `non-dedicated` - Set management interface in non-dedicated mode
* `tcp-syn-value` - Configure Tcp SYN's per sec, default 70
* `port-number` - port number to be configured (Specify port number)
* `core-index` - core index to be deleted (Specify core index)
* `max-cores` - max number of IO cores (Specify number of cores)
* `rxd-size` - Set new rx-descriptor size
* `txd-size` - Set new tx-descriptor size
* `rxq-size` - Set number of new rx queues
* `txq-size` - Set number of new tx queues
* `template-policy` - Apply policy template to the whole system (Policy template name)
* `id` - Specify unique Partition id
* `sessions` - 'all': Clear all sessions; 'sequence': Sequence number;
* `clear-all-sequence` - Sequence number (Specify the port physical port number)
* `clear-sequence` - Specify the port physical port number
* `diseth` - Specify the physical port number (Ethernet interface number)
* `dis-sequence` - Sequence number (Specify the sequence number)
* `enaeth` - Specify the physical port number (Ethernet interface number)
* `ena-sequence` - Sequence number (Specify the sequence number)
* `monitor-relation` - 'monitor-and': Configures the monitors in current template to work with AND logic; 'monitor-or': Configures the monitors in current template to work with OR logic;
* `linkup-ethernet1` - Specify the port physical port number (Ethernet interface number)
* `link-up-sequence1` - Sequence number (Specify the sequence number)
* `linkup-ethernet2` - Specify the port physical port number (Ethernet interface number)
* `link-up-sequence2` - Sequence number (Specify the sequence number)
* `linkup-ethernet3` - Specify the port physical port number (Ethernet interface number)
* `link-up-sequence3` - Sequence number (Specify the sequece number)
* `linkdown-ethernet1` - Specify the port physical port number (Ethernet interface number)
* `link-down-sequence1` - Sequence number (Specify the sequence number)
* `linkdown-ethernet2` - Specify the port physical port number (Ethernet interface number)
* `link-down-sequence2` - Sequence number (Specify the seqeuence number)
* `linkdown-ethernet3` - Specify the port physical port number (Ethernet interface number)
* `link-down-sequence3` - Sequence number (Specify the sequence number)
* `user-tag` - Customized tag
* `ssl-context-memory` - Total SSL context memory needed in units of MB. Will be rounded to closest multiple of 2MB
* `ssl-dma-memory` - Total SSL DMA memory needed in units of MB. Will be rounded to closest multiple of 2MB
* `nat-pool-addr-count` - Total configurable NAT Pool addresses in the System
* `l4-session-count` - Total Sessions in the System
* `auth-portal-html-file-size` - Specify maximum html file size for each html page in auth portal (in KB)
* `auth-portal-image-file-size` - Specify maximum image file size for default portal (in KB)
* `max-aflex-file-size` - Set maximum aFleX file size (Maximum file size in KBytes, default is 32K)
* `aflex-table-entry-count` - Total aFleX table entry in the system (Total aFlex entry in the system)
* `class-list-ipv6-addr-count` - Total IPv6 addresses for class-list
* `class-list-ac-entry-count` - Total entries for AC class-list
* `class-list-entry-count` - Total entries for class-list
* `max-aflex-authz-collection-number` - Specify the maximum number of collections supported by aFleX authorization
* `radius-table-size` - Total configurable CGNV6 RADIUS Table entries
* `authz-policy-number` - Specify the maximum number of authorization policies
* `ipsec-sa-number` - Specify the maximum number of IPsec SA
* `ram-cache-memory-limit` - Specify the maximum memory used by ram cache
* `waf-template-count` - Total configurable WAF Templates in the System
* `auth-session-count` - Total auth sessions in the system
* `monitored-entity-count` - Total number of monitored entities for visibility
* `name` - Specify name of Geolocation list
* `gslb-device-max` - Enter the number of gslb-device allowed (gslb-device count (default is max-value))
* `gslb-device-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-geo-location-max` - Enter the number of gslb-geo-location allowed (gslb-geo-location count (default is max-value))
* `gslb-geo-location-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-ip-list-max` - Enter the number of gslb-ip-list allowed (gslb-ip-list count (default is max-value))
* `gslb-ip-list-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-policy-max` - Enter the number of gslb-policy allowed (gslb-policy count (default is max-value))
* `gslb-policy-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-service-max` - Enter the number of gslb-service allowed (gslb-service count (default is max-value)
* `gslb-service-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-service-ip-max` - Enter the number of gslb-service-ip allowed (gslb-service-ip count (default is max-value))
* `gslb-service-ip-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-service-port-max` - Enter the number of gslb-service-port allowed ( gslb-service-port count (default is max-value))
* `gslb-service-port-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-site-max` - Enter the number of gslb-site allowed (gslb-site count (default is max-value))
* `gslb-site-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-svc-group-max` - Enter the number of gslb-svc-group allowed (gslb-svc-group count (default is max-value))
* `gslb-svc-group-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-template-max` - Enter the number of gslb-template allowed (gslb-template count (default is max-value))
* `gslb-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `gslb-zone-max` - Enter the number of gslb-zone allowed (gslb-zone count (default is max-value))
* `gslb-zone-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `health-monitor-max` - Enter the number of health monitors allowed (health-monitor count (default is max-value))
* `health-monitor-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `real-port-max` - Enter the number of real-ports allowed (real-port count (default is max-value))
* `real-port-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `real-server-max` - Enter the number of real-servers allowed (real-server count (default is max-value))
* `real-server-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `service-group-max` - Enter the number of service groups allowed (service-group count (default is max-value))
* `service-group-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `virtual-server-max` - Enter the number of virtual-servers allowed (virtual-server count (default is max-value))
* `virtual-server-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `virtual-port-max` - Enter the number of virtual-port allowed (virtual-port count (default is max-value))
* `virtual-port-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `cache-template-max` - Enter the number of cache-template allowed (cache-template count (default is max-value))
* `cache-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `client-ssl-template-max` - Enter the number of client-ssl-template allowed (client-ssl-template count (default is max-value))
* `client-ssl-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `conn-reuse-template-max` - Enter the number of conn-reuse-template allowed (conn-reuse-template count (default is max-value))
* `conn-reuse-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `fast-tcp-template-max` - Enter the number of fast-tcp-template allowed (fast-tcp-template count (default is max-value))
* `fast-tcp-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `fast-udp-template-max` - Enter the number of fast-udp-template allowed (fast-udp-template count (default is max-value))
* `fast-udp-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `fix-template-max` - Enter the number of fix-template allowed (fix-template count (default is max-value))
* `fix-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `http-template-max` - Enter the number of http-template allowed (http-template count (default is max-value))
* `http-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `link-cost-template-max` - Enter the number of link-cost-template allowed (link-cost-template count (default is max-value))
* `link-cost-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `persist-cookie-template-max` - Enter the number of persist-cookie-template allowed (persist-cookie-template count (default is max-value))
* `persist-cookie-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `persist-srcip-template-max` - Enter the number of persist-srcip-template allowed (persist-source-ip-template count (default is max-value))
* `persist-srcip-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `server-ssl-template-max` - Enter the number of server-ssl-template allowed (server-ssl-template count (default is max-value))
* `server-ssl-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `proxy-template-max` - Enter the number of proxy-template allowed (server-ssl-template count (default is max-value))
* `proxy-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `stream-template-max` - Enter the number of stream-template allowed (server-ssl-template count (default is max-value))
* `stream-template-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `threshold` - Enter the threshold as a percentage (Threshold in percentage(default is 100%))
* `static-ipv4-route-max` - Enter the number of static ipv4 routes allowed (Static ipv4 routes (default is max-value))
* `static-ipv4-route-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `static-ipv6-route-max` - Enter the number of static ipv6 routes allowed (Static ipv6 routes (default is max-value))
* `static-ipv6-route-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `ipv4-acl-line-max` - Enter the number of ACL lines allowed (IPV4 ACL lines (default is max-value))
* `ipv4-acl-line-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `ipv6-acl-line-max` - Enter the number of ACL lines allowed (IPV6 ACL lines (default is max-value))
* `ipv6-acl-line-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `static-arp-max` - Enter the number of static arp entries allowed (Static arp (default is max-value))
* `static-arp-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `static-neighbor-max` - Enter the number of static neighbor entries allowed (Static neighbors (default is max-value))
* `static-neighbor-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `static-mac-max` - Enter the number of static MAC entries allowed (Static MACs (default is max-value))
* `static-mac-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `object-group-max` - Enter the number of object groups allowed (Object group (default is max-value))
* `object-group-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `object-group-clause-max` - Enter the number of object group clauses allowed (Object group clauses (default is max-value))
* `object-group-clause-min-guarantee` - Minimum guaranteed value ( Minimum guaranteed value)
* `bw-limit-max` - Enter the bandwidth limit in mbps (Bandwidth limit in Mbit/s (no limits applied by default))
* `bw-limit-watermark-disable` - Disable watermark (90% drop, keep existing sessions, drop  new sessions)
* `concurrent-session-limit-max` - Enter the Concurrent Session limit (cps) (Concurrent-Session cps limit (no limits applied by default))
* `l4-session-limit-max` - Enter the l4 session limit in % (0.01% to 99.99%) (Enter a number from 0.01 to 99.99 (up to 2 digits precision))
* `l4-session-limit-min-guarantee` - minimum guaranteed value in % (up to 2 digits precision) (Enter a number from 0 to 99.99)
* `l4cps-limit-max` - Enter the L4 cps limit (L4 cps limit (no limits applied by default))
* `l7cps-limit-max` - L7cps-limit (L7 cps limit (no limits applied by default))
* `natcps-limit-max` - Enter the Nat cps limit (NAT cps limit (no limits applied by default))
* `fwcps-limit-max` - Enter the Firewall cps limit (Firewall cps limit (no limits applied by default))
* `ssl-throughput-limit-max` - Enter the ssl throughput limit in mbps (SSL Througput limit in Mbit/s (no limits applied by default))
* `ssl-throughput-limit-watermark-disable` - Disable watermark (90% drop, keep existing sessions, drop  new sessions)
* `sslcps-limit-max` - Enter the SSL cps limit (SSL cps limit (no limits applied by default))
* `use-l3` - Layer-3 Header based load balancing
* `use-l4` - Layer-3/4 Header based load balancing
* `link-detection-interval` - Link detection interval in msecs
* `table` - 'all': All tables;
* `audit-action` - 'enable': Enable table integrity audit; 'disable': Disable table integrity audit;
* `auto-sync-action` - 'enable': Enable auto-sync; 'disable': Disable auto-sync;
* `packet-round-robin` - Enable packet round robin for IPsec packets
* `crypto-core` - Crypto cores assigned for IPsec processing
* `crypto-mem` - Crypto memory percentage assigned for IPsec processing (rounded to increments of 10)
* `qat` - HW assisted QAT SSL module
* `action` - 'ipv4-only': Enable IPv4 HW forward entries only; 'ipv6-only': Enable IPv6 HW forward entries only; 'ipv4-ipv6': Enable Both IPv4/IPv6 HW forward entries (shared);
* `min` - Minimum packets-per-second threshold (per CPU) before redistribution will take effect (Minimum packets-per-second threshold (per CPU) before redistribution will take effect (default: 100000))
* `low` - CPU usage threshold (percentage) that will restore the normal packet distribution (default: 60)
* `high` - CPU usage threshold (percentage) that will trigger the redistribution (default: 75)
* `tcp` - Disallow redistribution of new TCP sessions
* `udp` - Disallow redistribution of new UDP sessions
* `bcast` - broadcast packets (per second limit)
* `ipmcast` - IP multicast packets (per second limit)
* `mcast` - multicast packets (per second limit)
* `unknown-ucast` - unknown unicast packets (per second limit)
* `ve-mac-scheme-val` - 'hash-based': Hash-based using the VE number; 'round-robin': Round Robin scheme; 'system-mac': Use system MAC address;
* `nscan-limit` - smp session scan limit (number of smp sessions per scan)
* `scan-freq` - smp session scan frequency (scan per second)
* `reset` - Reset IPMI Controller
* `ipv4-address` - IP address
* `ipv4-netmask` - IP subnet mask
* `default-gateway` - Default gateway address
* `dhcp` - IP addr obtained by BMC running DHCP
* `static` - Manually configured static IP address
* `add` - Add a new IPMI user (IPMI User Name)
* `password` - Password
* `administrator` - Full control
* `callback` - Lowest privilege level
* `operator` - Most BMC commands are allowed
* `user` - Only 'benign' commands are allowed
* `privilege` - Change an existing IPMI user privilege (IPMI User Name)
* `setname` - Change User Name (Current IPMI User Name)
* `newname` - New IPMI User Name
* `setpass` - Change Password (IPMI User Name)
* `newpass` - New Password
* `cmd` - Command to execute in double quotes
* `mode` - Set HW hash mode, default is 6 (1:dst-mac 2:src-mac 3:src-dst-mac 4:src-ip 5:dst-ip 6:rtag6 7:rtag7)
* `source_name` - bind source name
* `dest_name` - bind dest name
* `log-session-on-established` - Send TCP session creation log on completion of 3-way handshake
* `msl-time` - Configure maximum session life, default is 2 seconds (1-40 seconds, default is 2 seconds)
* `enable-shell-privileges` - enable the shell privileges for a given customer
* `application-mempool` - Enable application memory pool
* `application-flow` - Number of flows
* `basic-dpi-enable` - Enable basic dpi
* `complexity` - 'Strict': Strict: Min length:8, Min Lower Case:2, Min Upper Case:2, Min Numbers:2, Min Special Character:1; 'Medium': Medium: Min length:6, Min Lower Case:2, Min Upper Case:2, Min Numbers:1, Min Special Character:1; 'Simple': Simple: Min length:4, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:0;
* `aging` - 'Strict': Strict: Max Age-60 Days; 'Medium': Medium: Max Age- 90 Days; 'Simple': Simple: Max Age-120 Days;
* `history` - 'Strict': Strict: Does not allow upto 5 old passwords; 'Medium': Medium: Does not allow upto 4 old passwords; 'Simple': Simple: Does not allow upto 3 old passwords;
* `min-pswd-len` - Configure custom password length
* `listen-port` - Configure the listen port of RADIUS server (default 1813) (Port number)
* `ip-list-name` - IP-list name
* `ip-list-secret` - Configure shared secret
* `ip-list-secret-string` - The RADIUS secret
* `ip-list-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)
* `secret` - Configure shared secret
* `secret-string` - The RADIUS secret
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)
* `vrid` - Join a VRRP-A failover group
* `attribute-value` - 'inside-ipv6-prefix': Framed IPv6 Prefix; 'inside-ip': Inside IP address; 'inside-ipv6': Inside IPv6 address; 'imei': International Mobile Equipment Identity (IMEI); 'imsi': International Mobile Subscriber Identity (IMSI); 'msisdn': Mobile Subscriber Integrated Services Digital Network-Number (MSISDN); 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3; 'custom4': Customized attribute 4; 'custom5': Customized attribute 5; 'custom6': Customized attribute 6;
* `prefix-length` - '32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;
* `prefix-vendor` - RADIUS vendor attribute information (RADIUS vendor ID)
* `prefix-number` - RADIUS attribute number
* `value` - 'hexadecimal': Type of attribute value is hexadecimal;
* `custom-vendor` - RADIUS vendor attribute information (RADIUS vendor ID)
* `custom-number` - RADIUS attribute number
* `vendor` - RADIUS vendor attribute information (RADIUS vendor ID)
* `number` - RADIUS attribute number
* `disable-reply` - Toggle option for RADIUS reply packet(Default: Accounting response will be sent)
* `accounting-start` - 'ignore': Ignore; 'append-entry': Append the AVPs to existing entry (default); 'replace-entry': Replace the AVPs of existing entry;
* `accounting-stop` - 'ignore': Ignore; 'delete-entry': Delete the entry (default); 'delete-entry-and-sessions': Delete the entry and data sessions associated(CGN only);
* `accounting-interim-update` - 'ignore': Ignore (default); 'append-entry': Append the AVPs to existing entry; 'replace-entry': Replace the AVPs of existing entry;
* `accounting-on` - 'ignore': Ignore (default); 'delete-entries-using-attribute': Delete entries matching attribute in RADIUS Table;
* `attribute-name` - 'msisdn': Clear using MSISDN; 'imei': Clear using IMEI; 'imsi': Clear using IMSI; 'custom1': Clear using CUSTOM1 attribute configured; 'custom2': Clear using CUSTOM2 attribute configured; 'custom3': Clear using CUSTOM3 attribute configured; 'custom4': Clear using CUSTOM4 attribute configured; 'custom5': Clear using CUSTOM5 attribute configured; 'custom6': Clear using CUSTOM6 attribute configured;
* `custom-attribute-name` - Clear using customized attribute
* `shared` - Enable sharing with other partitions
* `include-geoloc-name-val` - Geolocation name to add
* `exclude-geoloc-name-val` - Geolocation name to exclude
* `geo-location-iana` - Load built-in IANA Database
* `geo-location-geolite2-city` - Load built-in Maxmind GeoLite2-City database. Database available from http://www.maxmind.com
* `geolite2-city-include-ipv6` - Include IPv6 address
* `geo-location-geolite2-country` - Load built-in Maxmind GeoLite2-Country database. Database available from http://www.maxmind.com
* `geolite2-country-include-ipv6` - Include IPv6 address
* `geo-location-load-filename` - Specify file to be loaded
* `template-name` - CSV template to load this file
* `geo-locn-obj-name` - Specify geo-location name, section range is (1-15)
* `first-ip-address` - Specify IP information (Specify IP address)
* `geol-ipv4-mask` - Specify IPv4 mask
* `ip-addr2` - Specify IP address range
* `first-ipv6-address` - Specify IPv6 address
* `geol-ipv6-mask` - Specify IPv6 mask
* `ipv6-addr2` - Specify IPv6 address range
* `class-list` - Bind class-list (class-list name)
* `ip-threat-action-tmpl` - Bind ip-threat-action Template (ip-threat-action Template number)
* `monitor-disable` - Disable FPGA Core CRC error monitoring and act on it
* `reboot-enable` - Enable system reboot if system encounters FPGA Core CRC error
* `ca-cert` - Configure CA Certificate
* `cert-host` - Configure certificate store host
* `protocol` - 'tftp': Use tftp for connection; 'ftp': Use ftp for connection; 'scp': Use scp for connection; 'http': Use http for connection; 'https': Use https for connection; 'sftp': Use sftp for connection;
* `cert-store-path` - Configure certificate store path
* `username` - Username for MFA validation
* `passwd-string` - Certificate store host password
* `second-factor` - Input second factor paramter
* `inner-tpid` - TPID for inner VLAN (Inner TPID, 16 bit hex value, default is 8100)
* `outer-tpid` - TPID for outer VLAN (Outer TPID, 16 bit hex value, default is 8100)
* `enable-all-ports` - Enable 802.1QinQ on all physical ports
* `port-count-kernel` - Total Ports to be allocated for kernel.
* `port-count-hm` - Total Ports to be allocated for hm.
* `port-count-logging` - Total Ports to be allocated for logging.
* `port-count-alg` - Total Ports to be allocated for alg types.
* `l2hm-hc-name` - Monitor Name
* `method-l2bfd` - Method is l2bfd
* `l2bfd-tx-interval` - Transmit interval between BFD packets
* `l2bfd-rx-interval` - Minimum receive interval capability (Milliseconds (default: 800))
* `l2bfd-multiplier` - Multiplier value used to compute holddown (value used to multiply the interval (default: 4))
* `l2hm-path-name` - Monitor Name
* `l2hm-vlan` - VLAN id
* `l2hm-setup-test-api` - Test-API Interface (Ethernet Interface)
* `ifpair-eth-start` - Ethernet port (Interface number)
* `ifpair-eth-end` - Ethernet port
* `ifpair-trunk-start` - Trunk groups
* `ifpair-trunk-end` - Trunk Group
* `l2hm-attach` - Monitor Name
* `log-for-reset-unknown-conn` - Log when rate exceed

