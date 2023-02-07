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

resource "thunder_file_class_list" "IPV4" {
    name = "IPV4_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/IPV4_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_IPV4" {
    provider = thunder.L3V_A
    name = "IPV4_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/IPV4_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "IPV6" {
    name = "IPV6_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/IPV6_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_IPV6" {
    provider = thunder.L3V_A
    name = "IPV6_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/IPV6_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "DNS" {
    name = "DNS_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/DNS_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_DNS" {
    provider = thunder.L3V_A
    name = "DNS_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/DNS_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "AC" {
    name = "AC_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/AC_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_AC" {
    provider = thunder.L3V_A
    name = "AC_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/AC_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "STR" {
    name = "STR_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/STR_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_STR" {
    provider = thunder.L3V_A
    name = "STR_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/STR_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "STR_CI" {
    name = "STR_CI_1"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/STR_CI_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list" "L3V_STR_CI" {
    provider = thunder.L3V_A
    name = "STR_CI_1"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/STR_CI_1.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "IPV4" {
    name = "IPV4_2"
    class_list_type = "ipv4"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/IPV4_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_IPV4" {
    provider = thunder.L3V_A
    name = "IPV4_2"
    class_list_type = "ipv4"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/IPV4_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "IPV6" {
    name = "IPV6_2"
    class_list_type = "ipv6"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/IPV6_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_IPV6" {
    provider = thunder.L3V_A
    name = "IPV6_2"
    class_list_type = "ipv6"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/IPV6_1X.txt"
    use_mgmt_port = 1
}

# thunder_file_class_list_convert does not support dns type

resource "thunder_file_class_list_convert" "AC" {
    name = "AC_2"
    class_list_type = "ac"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/AC_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_AC" {
    provider = thunder.L3V_A
    name = "AC_2"
    class_list_type = "ac"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/AC_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "STR" {
    name = "STR_2"
    class_list_type = "string"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/STR_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_STR" {
    provider = thunder.L3V_A
    name = "STR_2"
    class_list_type = "string"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/STR_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "STR_CI" {
    name = "STR_CI_2"
    class_list_type = "string-case-insensitive"
    protocol = "http"
    host = "192.168.92.200"
    path = "/class-list-files/STR_CI_1X.txt"
    use_mgmt_port = 1
}

resource "thunder_file_class_list_convert" "L3V_STR_CI" {
    provider = thunder.L3V_A
    name = "STR_CI_2"
    class_list_type = "string-case-insensitive"
    protocol = "https"
    host = "192.168.92.200"
    path = "/class-list-files/STR_CI_1X.txt"
    use_mgmt_port = 1
}
