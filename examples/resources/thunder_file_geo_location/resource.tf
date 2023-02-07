provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_file_geo_location" "test"{
    action = "import"
    file = "testing"
    file_handle = "/mnt/c/Users/smundhe/Workspace/thunder_terraform_provider/test_project/test2.csv"
}