provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_class_list" "resourceClassListTestipv4" {
  name = "Test1"
  type = "ipv4"
  file = 0
  ipv4_list {
    ipv4addr = "192.168.0.200/32"
    glid     = 20
    age      = 10
  }
}
resource "thunder_class_list" "resourceClassListTestipv6" {
  nameIpv6 = "Test2"
  typeIpv6 = "ipv6"
  fileIpv6 = 0
  ipv6_list {
    ipv6_addr = "4444::1/64"
    v6_glid   = "20"
  }
}
resource "thunder_class_list" "resourceClassListTestDNS" {
  namedns = "Test3"
  typedns = "dns"
  filedns = 0
  dns {
    dns_match_type   = "contains"
    dns_match_string = "DNS1"
    dns_lid          = 20
  }
  dns {
    dns_match_type_glid       = "contains"
    dns_match_string_glid     = "DNS2"
    dns_glid                  = "20"
    shared_partition_dns_glid = 0
  }

  dns {
    dns_match_type_ends_with   = "ends-with"
    dns_match_string_ends_with = "DNS3"
    dns_lid_ends_with          = 20
  }
}
resource "thunder_class_list" "resourceClassListTestString" {
  namestr = "Test4"
  typestr = "string"
  filestr = 0
  str_list {
    str1          = "String1"
    str_lid_dummy = 1
    str_lid       = 20
    value_str1    = "TestString"
  }
}