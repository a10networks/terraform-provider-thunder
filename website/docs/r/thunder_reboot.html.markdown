---
layout: "thunder"
page_title: "thunder: thunder_reboot"
sidebar_current: "docs-thunder-resource-reboot"
description: |-
	Provides details about thunder reboot resource for A10
---

# thunder\_reboot

`thunder_reboot` Provides details about thunder reboot
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_reboot" "Reboot_Test" {

all = 0
day_of_month = 0
reason_3 = "string"
reason_2 = "string"
in = "string"
month_2 = "string"
month = "string"
device = 0
reason = "string"
at = 0
time = "string"
cancel = 0
day_of_month_2 = 0
 
}
```

## Argument Reference

* `all` - Reboot all devices when VCS is enabled, or only this device itself if VCS is not enabled
* `at` - Reboot at a Specific time/date
* `cancel` - Cancel Pending Reboot
* `day_of_month` - Day of the Month
* `day_of_month_2` - Day of the Month
* `device` - Reboot a specific device when VCS is enabled (device id)
* `in` - Reboot after a time interval (Time in hours and minutes)
* `month` - ‘January’: Month of the year; ‘February’: Month of the year; ‘March’: Month of the year; ‘April’: Month of the year; ‘May’: Month of the year; ‘June’: Month of the year; ‘July’: Month of the year; ‘August’: Month of the year; ‘September’: Month of the year; ‘October’: Month of the year; ‘November’: Month of the year; ‘December’: Month of the year;
* `month_2` - ‘January’: Month of the year; ‘February’: Month of the year; ‘March’: Month of the year; ‘April’: Month of the year; ‘May’: Month of the year; ‘June’: Month of the year; ‘July’: Month of the r; ‘August’: Month of the year; ‘September’: Month of the year; ‘October’: Month of the year; ‘November’: Month of the year; ‘December’: Month of the year;
* `reason` - Reason for Reboot
* `reason_2` - Reason for Reboot
* `reason_3` - Reason for Reboot
* `time` - Time to Reboot (hh:mm)
