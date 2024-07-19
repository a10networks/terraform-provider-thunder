---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_aam_authentication_portal_logon_fail Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_aam_authentication_portal_logon_fail: Logon fail page configuration
  PLACEHOLDER
---

# thunder_aam_authentication_portal_logon_fail (Resource)

`thunder_aam_authentication_portal_logon_fail`: Logon fail page configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_portal_logon_fail" "thunder_aam_authentication_portal_logon_fail" {

  name = "default-portal"
  background {
    bgfile  = "53"
    bgstyle = "tile"
  }
  fail_msg_cfg {
    fail_msg        = 1
    fail_text       = "36"
    fail_font       = 1
    fail_face       = "Arial"
    fail_size       = 3
    fail_color      = 1
    fail_color_name = "black"
  }
  title_cfg {
    title            = 1
    title_text       = "53"
    title_font       = 1
    title_face       = "Arial"
    title_size       = 5
    title_color      = 1
    title_color_name = "black"
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `background` (Block List, Max: 1) (see [below for nested schema](#nestedblock--background))
- `fail_msg_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--fail_msg_cfg))
- `title_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--title_cfg))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--background"></a>
### Nested Schema for `background`

Optional:

- `bgcolor_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `bgcolor_value` (String) Specify 6-digit HEX color value
- `bgfile` (String) Specify background image filename
- `bgstyle` (String) 'tile': Tile; 'stretch': Stretch; 'fit': Fit;


<a id="nestedblock--fail_msg_cfg"></a>
### Nested Schema for `fail_msg_cfg`

Optional:

- `fail_color` (Number) Specify font color (Default: black)
- `fail_color_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `fail_color_value` (String) Specify 6-digit HEX color value
- `fail_face` (String) 'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;
- `fail_font` (Number) Sepcify font (Default: Arial)
- `fail_font_custom` (String) Specify custom font
- `fail_msg` (Number) Configure logon failure message in default logon fail page
- `fail_size` (Number) Specify font size (Default: 3)
- `fail_text` (String) Specify logon failure message (Default: Login Failed!!)


<a id="nestedblock--title_cfg"></a>
### Nested Schema for `title_cfg`

Optional:

- `title` (Number) Configure title in default logon fail page
- `title_color` (Number) Specify font color (Default: black)
- `title_color_name` (String) 'aqua': aqua; 'black': black; 'blue': blue; 'fuchsia': fuchsia; 'gray': gray; 'green': green; 'lime': lime; 'maroon': maroon; 'navy': navy; 'olive': olive; 'orange': orange; 'purple': purple; 'red': red; 'silver': silver; 'teal': teal; 'white': white; 'yellow': yellow;
- `title_color_value` (String) Specify 6-digit HEX color value
- `title_face` (String) 'Arial': Arial; 'Courier_New': Courier New; 'Georgia': Georgia; 'Times_New_Roman': Times New Roman; 'Verdana': Verdana;
- `title_font` (Number) Sepcify font (Default: Arial)
- `title_font_custom` (String) Specify custom font
- `title_size` (Number) Specify font size (Default: 5)
- `title_text` (String) Specify title (Default: Try Too Many Times)

