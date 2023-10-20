provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_radius_server" "test_thunder_fw_radius_server" {
  listen_port = 1211
  secret      = 1
  encrypted   = "CR7g9VpvefshUcltaTe06ONIm4Wr8dGKhaM7NZ2+rmpRCOQOmhzET7MC5CUQTNlP"
  attribute {
    attribute_value = "inside-ip"
    number          = 2
  }
  accounting_interim_update = "append-entry"
  accounting_on             = "delete-entries-using-attribute"
  attribute_name            = "imsi"
  sampling_enable {
    counters1 = "all"
  }
}