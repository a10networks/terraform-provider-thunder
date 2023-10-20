

provider  thunder  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource thunder_configure_sync "ConfigSync"{
    type =  "running"
    all_partitions =  0
    shared =  1
    partition_name =  "shared"
    address =  "10.64.3.183"
    auto_authentication =  1
    usr =  "admin"
    pwd =  "a10023"
}
