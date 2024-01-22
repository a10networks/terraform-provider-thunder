provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server" "thunder_aam_authentication_server" {

  ocsp {
    sampling_enable {
      counters1 = "all"
    }
    instance_list {
      name         = "test"
      url          = "http://192.168.0.1:80/"
      http_version = 1
      version_type = "1.1"
      sampling_enable {
        counters1 = "all"
      }
    }
  }
}