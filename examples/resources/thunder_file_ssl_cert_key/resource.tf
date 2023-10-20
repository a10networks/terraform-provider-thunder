

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource "thunder_file_ssl_cert_key" "FileSslCertKey"{
    name = "CERT_KEY"
    protocol = "scp"
    host = "10.64.3.186"
    path = "/home/server1/test123.pem"
    password = "a10"
    username = "server1"
    overwrite = 1
    secured  = 0
    use_mgmt_port = 1   
}
