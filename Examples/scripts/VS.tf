provider "vthunder" {
  address = "18.229.28.140"
  username = "admin"
  password = "a10"
}

resource "vthunder_virtual_server" "server5" {
name="VS5"
ha_dynamic = 1
vrid = 2
ip_address="10.10.22.27"
port_list {
		port_number=6200
		protocol="tcp"
		service_group="sg1"
		
	}
port_list {
		port_number=6300
		protocol="tcp"
		service_group="sg1"
		
	}
}