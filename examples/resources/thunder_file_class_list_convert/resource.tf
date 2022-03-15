provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

provider "thunder" {
    alias = "L3V_A"
    address  = var.dut9049
    username = var.username
    password = var.password
    partition = var.l3v_1
}

resource "thunder_file_class_list_convert" "IPV4_1" {
    name = "IPV4_1"
    class_list_type = "ipv4"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/IPV4_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_IPV4_1" {
    provider = thunder.L3V_A
    name = "IPV4_1"
    class_list_type = "ipv4"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/IPV4_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "IPV6_1" {
    name = "IPV6_1"
    class_list_type = "ipv6"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/IPV6_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_IPV6_1" {
    provider = thunder.L3V_A
    name = "IPV6_1"
    class_list_type = "ipv6"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/IPV6_1X.txt"
    use_mgmt_port = 1
}

# thunder_file_class_list_convert does not support dns type

resource "thunder_file_class_list_convert" "AC_1" {
    name = "AC_1"
    class_list_type = "ac"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/AC_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_AC_1" {
    provider = thunder.L3V_A
    name = "AC_1"
    class_list_type = "ac"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/AC_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "STR_1" {
    name = "STR_1"
    class_list_type = "string"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/STR_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_STR_1" {
    provider = thunder.L3V_A
    name = "STR_1"
    class_list_type = "string"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/STR_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "STR_CI_1" {
    name = "STR_CI_1"
    class_list_type = "string-case-insensitive"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/STR_CI_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_STR_CI_1" {
    provider = thunder.L3V_A
    name = "STR_CI_1"
    class_list_type = "string-case-insensitive"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/STR_CI_1X.txt"
    use_mgmt_port = 1
}
