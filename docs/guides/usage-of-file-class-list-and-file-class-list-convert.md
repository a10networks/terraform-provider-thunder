---
subcategory: ""
page_title: "Usage of file_class_list and file_class_list_convert - Thunder Provider"
description: |-
    Usage of file_class_list and file_class_list_convert.
---
 
# Usage of file_class_list and file_class_list_convert

In this provdier, we can use resources `thunder_file_class_list` and `thunder_file_class_list_convert` to do the operation of thunder CLI commands `import class-list` and `import class-list-convert`.
The imported file will be stored as a file-type class-list. You can check those files by

```
TH9049#sh class-list 
Name                                                             Type                     IP       Subnet   DNS      String   Location
AC_1                                                             ac                       0        0        0        2        file    
DNS_1                                                            dns                      0        0        2        0        file    
IPV4_1                                                           ipv4                     0        1        0        0        file    
IPV6_1                                                           ipv6                     0        1        0        0        file    
STR_1                                                            string                   0        0        0        2        file    
STR_CI_1                                                         string-case-insensitive  0        0        0        3        file    
Total: 6
```

# Example of Class-List Files

## AC Type

### A10 Format
For `thunder_file_class_list`
```
class-list AC_1 ac file

; AC (Total: 4)
contains AXYZ
ends-with B
equals CD
starts-with Efg
```

### Common Format
For `thunder_file_class_list_convert`
```
contains AXYZ
ends-with B
equals CD
starts-with Efg
```

## IPv4 Type

### A10 Format
For `thunder_file_class_list`
```
class-list "IPV4_1" ipv4 file
; IPv4

; IP subnet (Total: 1)
1.2.0.0/15
```

### Common Format
For `thunder_file_class_list_convert`
```
1.2.0.0/15
```

## IPv6 Type

### A10 Format
For `thunder_file_class_list`
```
class-list "IPV6_1" ipv6 file
; IPv6

; IP subnet (Total: 1)
a:b:c::/64
```

### Common Format
For `thunder_file_class_list_convert`
```
1.2.0.0/15
```

## String Type

### A10 Format
For `thunder_file_class_list`
```
class-list "STR_1" string file

; String (Total: 2)
str text-example
str pattern
```

### Common Format
For `thunder_file_class_list_convert`
```
str text-example
str pattern
```

## String-Case-Insensitive Type

### A10 Format
For `thunder_file_class_list`
```
class-list "STR_CI_1" string-case-insensitive file

; String (Total: 3)
str KKJ
str PATTERN
str TEXT-3
```

### Common Format
For `thunder_file_class_list_convert`
```
str KKJ
str PATTERN
str TEXT-3
```

## DNS Type

### A10 Format
For `thunder_file_class_list`
```
class-list "DNS_1" dns file

; DNS domain (Total: 3)
dns contains test-string
dns starts-with ddd
dns ends-with case-end
```

### Common Format
`thunder_file_class_list_convert` does not support DNS type.

```terraform
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
```
