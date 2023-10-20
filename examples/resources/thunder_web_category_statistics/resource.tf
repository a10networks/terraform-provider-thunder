provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_web_category_statistics" "resourceWebCategoryStatisticsTest" {
  sampling_enable {
    counters1 = "all"
  }
}