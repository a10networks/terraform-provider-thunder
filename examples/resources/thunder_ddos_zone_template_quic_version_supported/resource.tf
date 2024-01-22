provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_quic_version_supported" "thunder_ddos_zone_template_quic_version_supported" {
  quic_tmpl_name = "testing"
  user_tag       = "113"
  version_action = "drop"
  version_end    = "3"
  version_start  = "1"

}