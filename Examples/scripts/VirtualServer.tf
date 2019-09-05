provider "vthunder" {
  address = ""
  username = ""
  password = ""
}

resource "vthunder_virtual_server" "server5" {
name="VS5"
ha_dynamic = 1
vrid = 2
ip_address=""
port_list {
		port_number=
		protocol=""
		service_group="sg1"
		
	}
port_list {
		port_number=
		protocol=""
		service_group="sg1"
		
	}
}