

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource thunder_ethernet "test"{
  ethernet_list {
      ifnum = 1
      l3_vlan_fwd_disable = 1
      load_interval = 300
      mtu = 1500
      ip{
        dhcp = 1
        address_list{
            ipv4_address = "10.10.10.10"
            ipv4_netmask = "255.255.255.0"
          }
      }
  
      }
}
    