---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_radius_server"
sidebar_current: "docs-vthunder-resource-fw-radius-server"
description: |-
	Provides details about vthunder fw radius server resource for A10
---

# vthunder\_fw\_radius\_server

`vthunder_fw_radius_server` Provides details about vthunder fw radius server
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

// Put working JSON here
```

## Argument Reference

* `accounting_interim_update` - ‘ignore’: Ignore (default); ‘append-entry’: Append the AVPs to existing entry; ‘replace-entry’: Replace the AVPs of existing entry;
* `accounting_on` - ‘ignore’: Ignore (default); ‘delete-entries-using-attribute’: Delete entries matching attribute in RADIUS Table;
* `accounting_start` - ‘ignore’: Ignore; ‘append-entry’: Append the AVPs to existing entry (default); ‘replace-entry’: Replace the AVPs of existing entry;
* `accounting_stop` - ‘ignore’: Ignore; ‘delete-entry’: Delete the entry (default);
* `attribute_name` - ‘msisdn’: Clear using MSISDN; ‘imei’: Clear using IMEI; ‘imsi’: Clear using IMSI;
* `custom_attribute_name` - Clear using customized attribute
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)
* `listen_port` - Configure the listen port of RADIUS server (Port number)
* `secret` - Configure shared secret
* `secret_string` - The RADIUS secret
* `uuid` - uuid of the object
* `vrid` - Join a VRRP-A failover group
* `ip_list_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)
* `ip_list_name` - IP-list name
* `ip_list_secret` - Configure shared secret
* `ip_list_secret_string` - The RADIUS secret
* `counters1` - ‘all’: all; ‘msisdn-received’: MSISDN Received; ‘imei-received’: IMEI Received; ‘imsi-received’: IMSI Received; ‘custom-received’: Custom attribute Received; ‘radius-request-received’: RADIUS Request Received; ‘radius-request-dropped’: RADIUS Request Dropped (Malformed Packet); ‘request-bad-secret-dropped’: RADIUS Request Bad Secret Dropped; ‘request-no-key-vap-dropped’: RADIUS Request No Key Attribute Dropped; ‘request-malformed-dropped’: RADIUS Request Malformed Dropped; ‘request-ignored’: RADIUS Request Table Full Dropped; ‘radius-table-full’: RADIUS Request Dropped (Table Full); ‘secret-not-configured-dropped’: RADIUS Secret Not Configured Dropped; ‘ha-standby-dropped’: HA Standby Dropped; ‘ipv6-prefix-length-mismatch’: Framed IPV6 Prefix Length Mismatch; ‘invalid-key’: Radius Request has Invalid Key Field; ‘smp-mem-allocated’: RADIUS SMP Memory Allocated; ‘smp-mem-alloc-failed’: RADIUS SMP Memory Allocation Failed; ‘smp-mem-freed’: RADIUS SMP Memory Freed; ‘smp-created’: RADIUS SMP Created; ‘smp-in-rml’: RADIUS SMP in RML; ‘smp-deleted’: RADIUS SMP Deleted; ‘mem-allocated’: RADIUS Memory Allocated; ‘mem-alloc-failed’: RADIUS Memory Allocation Failed; ‘mem-freed’: RADIUS Memory Freed; ‘ha-sync-create-sent’: HA Record Sync Create Sent; ‘ha-sync-delete-sent’: HA Record Sync Delete Sent; ‘ha-sync-create-recv’: HA Record Sync Create Received; ‘ha-sync-delete-recv’: HA Record Sync Delete Received; ‘acct-on-filters-full’: RADIUS Acct On Request Ignored(Filters Full); ‘acct-on-dup-request’: Duplicate RADIUS Acct On Request; ‘ip-mismatch-delete’: Radius Entry IP Mismatch Delete; ‘ip-add-race-drop’: Radius Entry IP Add Race Drop; ‘ha-sync-no-key-vap-dropped’: HA Record Sync No key dropped; ‘inter-card-msg-fail-drop’: Inter-Card Message Fail Drop;
* `attribute_value` - ‘inside-ipv6-prefix’: Framed IPv6 Prefix; ‘inside-ip’: Inside IP address; ‘inside-ipv6’: Inside IPv6 address; ‘imei’: International Mobile Equipment Identity (IMEI); ‘imsi’: International Mobile Subscriber Identity (IMSI); ‘msisdn’: Mobile Subscriber Integrated Services Digital Network-Number (MSISDN); ‘custom1’: Customized attribute 1; ‘custom2’: Customized attribute 2; ‘custom3’: Customized attribute 3;
* `custom_number` - RADIUS attribute number
* `custom_vendor` - RADIUS vendor attribute information (RADIUS vendor ID)
* `name` - Customized attribute name
* `number` - RADIUS attribute number
* `prefix_length` - ‘32’: Prefix length 32; ‘48’: Prefix length 48; ‘64’: Prefix length 64; ‘80’: Prefix length 80; ‘96’: Prefix length 96; ‘112’: Prefix length 112;
* `prefix_number` - RADIUS attribute number
* `prefix_vendor` - RADIUS vendor attribute information (RADIUS vendor ID)
* `value` - ‘hexadecimal’: Type of attribute value is hexadecimal;
* `vendor` - RADIUS vendor attribute information (RADIUS vendor ID)

