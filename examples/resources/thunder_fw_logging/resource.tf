provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_logging" "test_thunder_fw_logging" {
  name = "test_logging_template"
  sampling_enable {
    counters1 = "all"
  }
  gtp {
    sampling_enable {
      counters1 = "all"
    }
  }
}