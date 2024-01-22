provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_zone_template_http_malformed_http" "thunder_ddos_zone_template_http_malformed_http" {
  http_tmpl_name                       = "testing"
  malformed_http                       = "check"
  malformed_http_action                = "drop"
  malformed_http_bad_chunk_mon_enabled = 0
  malformed_http_max_content_length    = 4294967295
  malformed_http_max_header_name_size  = 64
  malformed_http_max_line_size         = 32512
  malformed_http_max_num_headers       = 90
  malformed_http_max_req_line_size     = 32512

}