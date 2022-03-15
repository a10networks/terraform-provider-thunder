---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_service_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_fw_service_group (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **health_check** (String)
- **id** (String) The ID of this resource.
- **member_list** (Block List) (see [below for nested schema](#nestedblock--member_list))
- **name** (String)
- **packet_capture_template** (String)
- **protocol** (String)
- **sampling_enable** (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- **traffic_replication_mirror_ip_repl** (Number)
- **user_tag** (String)
- **uuid** (String)

<a id="nestedblock--member_list"></a>
### Nested Schema for `member_list`

Optional:

- **name** (String)
- **packet_capture_template** (String)
- **port** (Number)
- **sampling_enable** (Block List) (see [below for nested schema](#nestedblock--member_list--sampling_enable))
- **user_tag** (String)
- **uuid** (String)

<a id="nestedblock--member_list--sampling_enable"></a>
### Nested Schema for `member_list.sampling_enable`

Optional:

- **counters1** (String)



<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- **counters1** (String)

