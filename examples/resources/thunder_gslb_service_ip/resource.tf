provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_gslb_service_ip" "thunder_gslb_Service_ip_test" {
    node_name = "vs2"
    action = "disable"
    health_check_disable = 1
    health_check_protocol_disable = 1
    ip_address = "10.1.20.10"
    external_ip = "172.173.148.196"
    port_list {
        action1 = "disable"
        health_check_disable1 = 1
        health_check_protocol_disable1 = 1
        port_num = 80
        port_proto = "tcp"
        user_tag = "incedo"
    }
}
