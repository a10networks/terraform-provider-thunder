

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource "thunder_file_ssl_key" "FileSslKey"{
    name = "FILESSLKEY"
    protocol = "scp"
    host = "10.64.3.186"
    path = "/var/www/html/server.key"
    password = "a10"
    username = "server1"
    overwrite = 1
    secured  = 0
    use_mgmt_port = 1   
}