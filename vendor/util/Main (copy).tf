provider "vthunder" {
  address = "18.229.28.140"
  username = "admin"
  password = "a10"
}

resource "vthunder_virtual_server" "server4" {
name="VS4"
ip_address="10.10.10.10"
enable_disable_action="enable"
redistribution_flagged=1
arp_disable=1
stats_data_action="stats-data-enable"
extended_stats=1
disable_vip_adv=1
uuid="ID4"
port_list {
		name="portlist4"
		protocol="tcp"
		stats_data_action="stats-data-enable"
		persist_type="src-dst-ip-swap-persist"
		def_selection_if_pref_failed="def-selection-if-pref-failed"
		action="enable"
		alt_protocol2="tcp"
		alt_protocol1="http"
		precedence=1
		optimization_level="0"
		acl_name_list = {
			acl_name="aclname4"
		}
	}
}