---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_saml_service_provider_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_saml_service_provider_stats: Statistics for the object service-provider
  PLACEHOLDER
---

# thunder_aam_authentication_saml_service_provider_stats (Data Source)

`thunder_aam_authentication_saml_service_provider_stats`: Statistics for the object service-provider

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_aam_authentication_saml_service_provider_stats" "thunder_aam_authentication_saml_service_provider_stats" {

  name = "test"
}
output "get_aam_authentication_saml_service_provider_stats" {
  value = ["${data.thunder_aam_authentication_saml_service_provider_stats.thunder_aam_authentication_saml_service_provider_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify SAML authentication service provider name

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `acs_authz_fail` (Number) SAML Single-Sign-On Authorization Fail
- `acs_error` (Number) SAML Single-Sign-On Error
- `acs_req` (Number) SAML Single-Sign-On Request
- `acs_success` (Number) SAML Single-Sign-On Success
- `glo_slo_success` (Number) Total Global Logout Success
- `loc_slo_success` (Number) Total Local Logout Success
- `login_auth_req` (Number) Login Authentication Request
- `login_auth_resp` (Number) Login Authentication Response
- `other_error` (Number) Other Error
- `par_slo_success` (Number) Total Partial Logout Success
- `slo_error` (Number) Single Logout Error
- `slo_req` (Number) Single Logout Request
- `slo_success` (Number) Single Logout Success
- `sp_metadata_export_req` (Number) Metadata Export Request
- `sp_metadata_export_success` (Number) Metadata Export Success
- `sp_slo_req` (Number) SP-initiated Single Logout Request

