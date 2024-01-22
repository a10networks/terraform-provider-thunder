provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_detection_settings" "thunder_ddos_detection_settings" {
  ctrl_cpu_usage                   = 78
  de_escalation_quiet_time         = 22
  detection_window_size            = 2
  detector_mode                    = "standalone"
  export_interval                  = 20
  full_core_enable                 = 1
  histogram_de_escalate_percentage = 58
  histogram_escalate_percentage    = 39
  initial_learning_interval        = 22
  notification_debug_log           = "enable"
  pkt_sampling {
    override_rate = 24997
    assign_index  = 11
    assign_rate   = 41341989
  }
  standalone_settings {
    action = "disable"
    sflow {
      listening_port = 6343
    }
    netflow {
      listening_port          = 9996
      template_active_timeout = 30
    }
  }
  top_k_reset_interval = 35

}