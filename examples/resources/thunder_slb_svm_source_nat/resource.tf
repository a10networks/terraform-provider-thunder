provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_svm_source_nat" "test_thunder_slb_svm_source_nat" {
  pool = "pool1"
}