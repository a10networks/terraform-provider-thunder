---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_password_policy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_password_policy: Configure Password Complexity, Passsword Aging, Password history under password policy
  PLACEHOLDER
---

# thunder_system_password_policy (Resource)

`thunder_system_password_policy`: Configure Password Complexity, Passsword Aging, Password history under password policy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_password_policy" "thunder_system_password_policy" {
  aging                        = "Strict"
  complexity                   = "Strict"
  forbid_consecutive_character = "0"
  history                      = "Strict"
  min_pswd_len                 = 8
  repeat_character_check       = "disable"
  username_check               = "disable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `aging` (String) 'Strict': Strict: Max Age-60 Days; 'Medium': Medium: Max Age- 90 Days; 'Simple': Simple: Max Age-120 Days;
- `complexity` (String) 'Strict': Strict: Min length:8, Min Lower Case:2, Min Upper Case:2, Min Numbers:2, Min Special Character:1, CHANGE Min 8 Characters; 'Medium': Medium: Min length:6, Min Lower Case:2, Min Upper Case:2, Min Numbers:1, Min Special Character:1, CHANGE Min 6 Characters; 'Default': Default: Min length:9, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:1, CHANGE Min 1 Characters; 'Simple': Simple: Min length:4, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:0, CHANGE Min 4 Characters;
- `forbid_consecutive_character` (String) '0': Will disable the check; '3': Three consecutive characters on keyboard will not be allowed.; '4': Four consecutive characters on keyboard will not be allowed.; '5': Five consecutive characters on keyboard will not be allowed.;
- `history` (String) 'Strict': Strict: Does not allow upto 5 old passwords; 'Medium': Medium: Does not allow upto 4 old passwords; 'Simple': Simple: Does not allow upto 3 old passwords;
- `min_pswd_len` (Number) Configure custom password length
- `repeat_character_check` (String) 'enable': Prohibition of consecutive repeated input of the same letter/number, case sensitive; 'disable': Will not check if the password contains repeat characters;
- `username_check` (String) 'enable': Prohibition to set password contains user account, case sensitive; 'disable': Will not check if the password contains user account;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

