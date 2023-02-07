provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_health_monitor" "common" {
    method {
        radius {
            radius = 1
            radius_expect = 1
            radius_password_string = "radius"
            radius_port = 345
            radius_response_code = 13
            radius_secret = "secret"
            radius_username = "radius"
        }
    }
    ssl_version = 33
    name = "tf_test"
}
