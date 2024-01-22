provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_offload_cpus" "thunder_offload_cpus" {

  num_offload_cpus = 2

}