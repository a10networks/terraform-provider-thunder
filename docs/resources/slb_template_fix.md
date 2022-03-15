---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_fix Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_slb_template_fix (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **insert_client_ip** (Number)
- **logging** (String)
- **name** (String)
- **tag_switching** (Block List) (see [below for nested schema](#nestedblock--tag_switching))
- **user_tag** (String)
- **uuid** (String)

<a id="nestedblock--tag_switching"></a>
### Nested Schema for `tag_switching`

Optional:

- **equals** (String)
- **service_group** (String)
- **switching_type** (String)

