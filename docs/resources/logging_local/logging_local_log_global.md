---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_logging_local_log_global Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_logging_local_log_global:
  PLACEHOLDER
---

# thunder_logging_local_log_global (Resource)

`thunder_logging_local_log_global`: 

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_local_log_global" "thunder_logging_local_log_global" {
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'enqueue': Total local-log enqueue; 'enqueue-full': Total local-log queue full; 'enqueue-error': Total local-log enqueue error; 'dequeue': Total local-log dequeue; 'dequeue-error': Total local-log dequeue processing error; 'raw-log': Total local-log raw logs; 'raw-log-error': Total raw log logging error; 'log-summarized': Total raw log summarized; 'l1-log-summarized': Total layer 1 log summarized; 'l2-log-summarized': Total layer 2 log summarized; 'log-summarized-error': Total local-log summarization error; 'aam-db': Total local-log AAM raw database; 'ep-db': Total local-log EP raw database; 'fw-db': Total local-log Firewall raw database; 'aam-top-user-db': Total local-log AAM top user summary database; 'ep-top-user-db': Total local-log EP top user summary database; 'ep-top-src-db': Total local-log EP top client summary database; 'ep-top-dst-db': Total local-log EP top destination summary database; 'ep-top-domain-db': Total local-log EP top domain summary database; 'ep-top-web-category-db': Total local-log EP top web-category summary database; 'ep-top-host-db': Total local-log EP top host summary database; 'fw-top-app-db': Total local-log Firewall top application summary database; 'fw-top-src-db': Total local-log Firewall top source summary database; 'fw-top-app-src-db': Total local-log Firewall top application and source summary database; 'fw-top-category-db': Total local-log Firewall top category summary database; 'db-erro': Total local-log database create error; 'query': Total local-log axapi query; 'response': Total local-log axapi response; 'query-error': Total local-log axapi query error; 'fw-top-thr-db': Total local-log Firewall top threat summary database; 'fw-top-thr-src-db': Total local-log Firewall top threat and source summary database;

