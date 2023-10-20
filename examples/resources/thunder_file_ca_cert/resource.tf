

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource "thunder_file_ca_cert" "test"{
    name = "CACERT_1"
    protocol = "scp"
    host = "10.64.3.186"
    path = "/var/www/html/server.pem"
    use_mgmt_port = 1
    password = "a10"
    username = "server1"
    certificate_type = "pem"
    overwrite = 1


}
