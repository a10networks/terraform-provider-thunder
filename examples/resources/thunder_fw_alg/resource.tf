provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_alg" "thunder_fw_alg" {
  dns {
    default_port_disable = "default-port-disable"
  }
  esp {
    default_port_disable = "default-port-disable"
    sampling_enable {
      counters1 = "all"
    }
  }
  ftp {
    default_port_disable = "default-port-disable"
    sampling_enable {
      counters1 = "all"
    }
  }
  icmp {
    disable = "disable"
  }
  pptp {
    default_port_disable = "default-port-disable"
    sampling_enable {
      counters1 = "all"
    }
  }
  rtsp {
    default_port_disable = "default-port-disable"
    sampling_enable {
      counters1 = "all"
    }
  }
  sip {
    default_port_disable = "default-port-disable"
    sampling_enable {
      counters1 = "all"
    }
  }
  tftp {
    default_port_disable = "default-port-disable"
    sampling_enable {
      counters1 = "all"
    }
  }
}