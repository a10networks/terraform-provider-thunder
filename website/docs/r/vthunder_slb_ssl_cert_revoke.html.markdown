---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_ssl_cert_revoke"
sidebar_current: "docs-vthunder-resource-slb-ssl-cert-revoke"
description: |-
    Provides details about vthunder SLB ssl cert revoke resource for A10
---

# vthunder\_slb\_ssl\_cert\_revoke

`vthunder_slb_ssl_cert_revoke` Provides details about vthunder SLB ssl cert revoke
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_ssl_cert_revoke" "SSLCertRevoke" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'ocsp_stapling_response_good’: OCSP stapling response good; 'ocsp_chain_status_good’: Certificate chain status good; 'ocsp_chain_status_revoked’: Certificate chain status revoked; 'ocsp_chain_status_unknown’: Certificate chain status unknown; 'ocsp_request’: OCSP requests; 'ocsp_response’: OCSP responses; 'ocsp_connection_error’: OCSP connection error; 'ocsp_uri_not_found’: OCSP URI not found; 'ocsp_uri_https’: Log OCSP URI https; 'ocsp_uri_unsupported’: OCSP URI unsupported; 'ocsp_response_status_good’: OCSP response status good; 'ocsp_response_status_revoked’: OCSP response status revoked; 'ocsp_response_status_unknown’: OCSP response status unknown; 'ocsp_cache_status_good’: OCSP cache status good; 'ocsp_cache_status_revoked’: OCSP cache status revoked; 'ocsp_cache_miss’: OCSP cache miss; 'ocsp_cache_expired’: OCSP cache expired; 'ocsp_other_error’: Log OCSP other errors; 'ocsp_response_no_nonce’: Log OCSP other errors; 'ocsp_response_nonce_error’: Log OCSP other errors; 'crl_request’: CRL requests; 'crl_response’: CRL responses; 'crl_connection_error’: CRL connection errors; 'crl_uri_not_found’: CRL URI not found; 'crl_uri_https’: CRL URI https; 'crl_uri_unsupported’: CRL URI unsupported; 'crl_response_status_good’: CRL response status good; 'crl_response_status_revoked’: CRL response status revoked; 'crl_response_status_unknown’: CRL response status unknown; 'crl_cache_status_good’: CRL cache status good; 'crl_cache_status_revoked’: CRL cache status revoked; 'crl_other_error’: CRL other errors;