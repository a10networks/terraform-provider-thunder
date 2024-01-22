provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_encap" "thunder_ddos_zone_template_encap" {
  encap_tmpl_name    = "test"
  preserve_source_ip = 1
  tunnel_encap {
    ip_cfg {
      ip_encap = 1
      always {
        ipv4_addr = "10.10.10.10"
      }
    }
  }
  user_tag = "7"

}