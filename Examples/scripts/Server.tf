provider "vthunder" {
  address = ""
  username = ""
  password = ""
}

resource "vthunder_rib_route" "rib"{
ip_dest_addr=""
ip_mask=""
ip_nexthop_ipv4{
ip_next_hop=""
distance_nexthop_ip=
}
}

resource "vthunder_ethernet" "eth"{
ethernet_list{
ifnum=1
ip{
dhcp=1
address_list={
ipv4_address=""
ipv4_netmask=""
}
}
action="enable"
}
ethernet_list{
ifnum=2
ip{
dhcp=1
address_list={
ipv4_address=""
ipv4_netmask=""
}
}
action="enable"
}
}

resource "vthunder_server" "rs9" {
health_check_disable=1
name="rs9"
host=""
port_list {
health_check_disable=1
port_number=
protocol=""
} 
}

resource "vthunder_service_group" "sg9" {
name="sg9"
protocol=""
member_list {
name="${vthunder_server.rs9.name}"
port=
}
}

resource "vthunder_virtual_server" "vs9" {
name="vs9"
ip_address=""
port_list {
auto=1
port_number=
protocol=""
service_group="${vthunder_service_group.sg9.name}"
snat_on_vip=1
}
}
#TF_ACC=true go test -v vthunder/*