provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_encap" "thunder_ddos_template_encap" {
  encap_tmpl_name    = "39"
  preserve_source_ip = 1
  tunnel_encap {
    gre_cfg {
      gre_encap = 1
      gre_always {
        gre_ipv4 = "10.10.10.10"
        key_ipv4 = "1"
      }
    }
  }
  user_tag = "65"
}