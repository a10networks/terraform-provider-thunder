provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_routing_ospf" "thunder_snmp_server_enable_traps_routing_ospf" {
  ospfifauthfailure           = 0
  ospfifconfigerror           = 0
  ospfifrxbadpacket           = 0
  ospfifstatechange           = 0
  ospflsdbapproachingoverflow = 0
  ospflsdboverflow            = 0
  ospfmaxagelsa               = 0
  ospfnbrstatechange          = 0
  ospforiginatelsa            = 0
  ospftxretransmit            = 0
  ospfvirtifauthfailure       = 0
  ospfvirtifconfigerror       = 0
  ospfvirtifrxbadpacket       = 0
  ospfvirtifstatechange       = 0
  ospfvirtiftxretransmit      = 0
  ospfvirtnbrstatechange      = 0
}