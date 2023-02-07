provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_vrrp_common" "test" {
    set_id = 9
    device_id = 7
    hostid_append_to_vrid {  
        hostid_append_to_vrid_value = 6
    }
}
