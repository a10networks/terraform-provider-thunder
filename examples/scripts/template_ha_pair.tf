provider "thunder" {
  address = ""
  username = ""
  password = ""
}

resource "thunder_vrrp_vrid" "vrrp_vrid" {
	vrid_val = 0
	floating_ip {
		ip_address_cfg {
			ip_address = "10.0.2.4"
		}
		ip_address_cfg {
			ip_address = "10.0.3.5"
		}
	}
	blade_parameters {
		priority = 231
	}
}

resource "thunder_acos_cloud_integration_ecosystem_oracle" "oracle_resource" {
        service_label = "sg1"
        host_name = ""
        port = 443
        health_check_interval = "20"
        compartment_id = ""
        tenancy_id = ""
        user_id = ""
        fingerprint = ""
        private_key_path = ""
        action = "enable"
}

resource "thunder_vrrp_peer_group" "vrrp_peer_group" {
	peer {
		ip_peer_address_cfg {
			ip_peer_address = "10.0.2.2"
		}
	}
}

resource "thunder_vrrp_common" "vrrp_common" {
	device_id = 1
	set_id = 1
	action = "enable"
}

resource "thunder_virtual_server" "vs1" {
        name="vs2"
        ha_dynamic = 1
        ip_address="10.0.2.5"
        port_list {
                auto=1
                port_number=80
                protocol="tcp"
                service_group="${thunder_service_group.sg1.name}"
                ha_conn_mirror = 1
                on_syn = 1
        }
}

resource "thunder_service_group" "sg1" {
        name="sg1"
        protocol="TCP"
}
