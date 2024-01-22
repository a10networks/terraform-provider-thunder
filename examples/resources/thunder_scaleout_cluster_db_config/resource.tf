provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_db_config" "thunder_scaleout_cluster_db_config" {

  broken_detect_timeout = 12000
  client_recv_timeout   = 13000
  clientport            = 21395
  elect_conn_timeout    = 1200
  initlimit             = 444
  loopback_intf_support = 1
  maxsessiontimeout     = 30000
  minsessiontimeout     = 100
  more_election_packet  = 1
  synclimit             = 182
  ticktime              = 3225
  cluster_id            = 2
}