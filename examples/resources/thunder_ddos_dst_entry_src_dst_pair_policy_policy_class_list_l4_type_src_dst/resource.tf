provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_l4_type_src_dst" "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_l4_type_src_dst" {
  dst_entry_name        = "test"
  src_based_policy_name = "test"
  class_list_name       = "test"
  deny                  = 1
  glid                  = "3"
  protocol              = "tcp"
  user_tag              = "testing"

}