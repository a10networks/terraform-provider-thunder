---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_server_ldap Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_server_ldap: LDAP Authentication Server
  PLACEHOLDER
---

# thunder_aam_authentication_server_ldap (Resource)

`thunder_aam_authentication_server_ldap`: LDAP Authentication Server

__PLACEHOLDER__

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `instance_list` (Block List) (see [below for nested schema](#nestedblock--instance_list))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--instance_list"></a>
### Nested Schema for `instance_list`

Required:

- `name` (String) Specify LDAP authentication server name

Optional:

- `admin_dn` (String) The LDAP server's admin DN
- `admin_secret` (Number) Specify the LDAP server's admin secret password
- `auth_type` (String) 'ad': Active Directory. Default; 'open-ldap': OpenLDAP;
- `base` (String) Specify the LDAP server's search base
- `bind_with_dn` (Number) Enforce using DN for LDAP binding(All user input name will be used to create DN)
- `ca_cert` (String) Specify the LDAPS CA cert filename (Trusted LDAPS CA cert filename)
- `default_domain` (String) Specify default domain for LDAP
- `derive_bind_dn` (Block List, Max: 1) (see [below for nested schema](#nestedblock--instance_list--derive_bind_dn))
- `dn_attribute` (String) Specify Distinguished Name attribute, default is CN
- `health_check` (Number) Check server's health status
- `health_check_disable` (Number) Disable configured health check configuration
- `health_check_string` (String) Health monitor name
- `host` (Block List, Max: 1) (see [below for nested schema](#nestedblock--instance_list--host))
- `ldaps_conn_reuse_idle_timeout` (Number) Specify LDAPS connection reuse idle timeout value (in seconds) (Specify idle timeout value (in seconds), default is 0 (not reuse LDAPS connection))
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `port` (Number) Specify the LDAP server's authentication port, default is 389
- `port_hm` (String) Check port's health status
- `port_hm_disable` (Number) Disable configured port health check configuration
- `prompt_pw_change_before_exp` (Number) Prompt user to change password before expiration in N days. This option only takes effect when server type is AD (Prompt user to change password before expiration in N days, default is not to prompt the user)
- `protocol` (String) 'ldap': Use LDAP (default); 'ldaps': Use LDAP over SSL; 'starttls': Use LDAP StartTLS;
- `pwdmaxage` (Number) Specify the LDAP server's default password expiration time (in seconds) (The LDAP server's default password expiration time (in seconds), default is 0 (no expiration))
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--instance_list--sampling_enable))
- `secret_string` (String) secret password
- `timeout` (Number) Specify timout for LDAP, default is 10 seconds (The timeout, default is 10 seconds)
- `uuid` (String) uuid of the object

<a id="nestedblock--instance_list--derive_bind_dn"></a>
### Nested Schema for `instance_list.derive_bind_dn`

Optional:

- `username_attr` (String) Specify attribute name of username


<a id="nestedblock--instance_list--host"></a>
### Nested Schema for `instance_list.host`

Optional:

- `hostip` (String) Server's hostname(Length 1-31) or IP address
- `hostipv6` (String) Server's IPV6 address


<a id="nestedblock--instance_list--sampling_enable"></a>
### Nested Schema for `instance_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'admin-bind-success': Admin Bind Success; 'admin-bind-failure': Admin Bind Failure; 'bind-success': User Bind Success; 'bind-failure': User Bind Failure; 'search-success': Search Success; 'search-failure': Search Failure; 'authorize-success': Authorization Success; 'authorize-failure': Authorization Failure; 'timeout-error': Timeout; 'other-error': Other Error; 'request': Request; 'ssl-session-created': TLS/SSL Session Created; 'ssl-session-failure': TLS/SSL Session Failure; 'pw_expiry': Password expiry; 'pw_change_success': Password change success; 'pw_change_failure': Password change failure;



<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'admin-bind-success': Total Admin Bind Success; 'admin-bind-failure': Total Admin Bind Failure; 'bind-success': Total User Bind Success; 'bind-failure': Total User Bind Failure; 'search-success': Total Search Success; 'search-failure': Total Search Failure; 'authorize-success': Total Authorization Success; 'authorize-failure': Total Authorization Failure; 'timeout-error': Total Timeout; 'other-error': Total Other Error; 'request': Total Request; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response; 'job-start-error': Total Job Start Error; 'polling-control-error': Total Polling Control Error; 'ssl-session-created': TLS/SSL Session Created; 'ssl-session-failure': TLS/SSL Session Failure; 'ldaps-idle-conn-num': LDAPS Idle Connection Number; 'ldaps-inuse-conn-num': LDAPS In-use Connection Number; 'pw-expiry': Total Password expiry; 'pw-change-success': Total password change success; 'pw-change-failure': Total password change failure;

