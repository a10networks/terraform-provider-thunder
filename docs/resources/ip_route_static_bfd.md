---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_route_static_bfd Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_ip_route_static_bfd (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **local_ip** (String) Local IP address
- **nexthop_ip** (String) Nexthop IP address

### Optional

- **action** (String) 'down': BFD down;  (BFD state)
- **id** (String) The ID of this resource.
- **template** (String) Configure tracking template (bind tracking template name)
- **threshold** (Number) action triggering threshold

### Read-Only

- **uuid** (String) uuid of the object

