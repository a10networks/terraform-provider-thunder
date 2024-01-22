provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_job_offload_stats" "thunder_system_job_offload_stats" {

}
output "get_system_job_offload_stats" {
  value = ["${data.thunder_system_job_offload_stats.thunder_system_job_offload_stats}"]
}