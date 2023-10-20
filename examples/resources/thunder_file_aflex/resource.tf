

provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}


resource "thunder_file_aflex" "FileAflex" {
    host  = "10.64.3.186"
    name ="test2.txt"
    path = "/var/www/html/test2.txt"
    protocol = "scp"
    password =  "a10"
    use_mgmt_port = 1
    username = "server1"
    overwrite = 1
}