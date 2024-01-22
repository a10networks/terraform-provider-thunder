provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server_ldap" "thunder_aam_authentication_server_ldap" {
  instance_list {
    name = "test"
    host {
      hostip = "17"
    }
    base           = "75"
    port           = 389
    pwdmaxage      = 1
    admin_dn       = "37"
    admin_secret   = 1
    secret_string  = "87"
    timeout        = 10
    dn_attribute   = "cn"
    default_domain = "2"
    bind_with_dn   = 1
    derive_bind_dn {
      username_attr = "3"
    }
    protocol                      = "ldap"
    ca_cert                       = "56"
    ldaps_conn_reuse_idle_timeout = 1
    auth_type                     = "ad"
    prompt_pw_change_before_exp   = 970
    sampling_enable {
      counters1 = "all"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
}