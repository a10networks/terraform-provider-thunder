---
layout: "thunder"
page_title: "thunder: thunder_slb_ssl_forward_proxy"
sidebar_current: "docs-thunder-resource-slb-ssl-forward-proxy"
description: |-
    Provides details about thunder SLB ssl forward proxy resource for A10
---

# thunder\_slb\_ssl\_forward\_proxy

`thunder_slb_ssl_forward_proxy` Provides details about thunder SLB ssl forward proxy
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_ssl_forward_proxy" "SSLForwardProxy" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'cert_create’: Certificates created; 'cert_expr’: Certificates expired; 'cert_hit’: Certificate cache hits; 'cert_miss’: Certificate cache miss; 'conn_bypass’: Connections bypassed; 'conn_inspect’: Connections inspected; 'bypass-failsafe-ssl-sessions’: Bypass Failsafe SSL sessions; 'bypass-sni-sessions’: Bypass SNI sessions; 'bypass-client-auth-sessions’: Bypass Client Auth sessions; 'failed-in-ssl-handshakes’: Failed in SSL handshakes; 'failed-in-crypto-operations’: Failed in crypto operations; 'failed-in-tcp’: Failed in TCP; 'failed-in-certificate-verification’: Failed in Certificate verification; 'failed-in-certificate-signing’: Failed in Certificate signing; 'invalid-ocsp-stapling-response’: Invalid OCSP Stapling Response; 'revoked-ocsp-response’: Revoked OCSP Response; 'unsupported-ssl-version’: Unsupported SSL version; 'certificates-in-cache’: Certificates in cache; 'connections-failed’: Connections failed; 'aflex-bypass’: Bypass triggered by aFleX; 'bypass-cert-subject-sessions’: Bypass Cert Subject sessions; 'bypass-cert-issuer-sessions’: Bypass Cert issuer sessions; 'bypass-cert-san-sessions’: Bypass Cert SAN sessions; 'bypass-no-sni-sessions’: Bypass NO SNI sessions; 'reset-no-sni-sessions’: Reset No SNI sessions; 'bypass-username-sessions’: Bypass Username sessions; 'bypass-ad-group-sessions’: Bypass AD-group sessions; 'cert_in_cache’: Certificates in cache; 'tot_conn_in_buff’: Total buffered async connections; 'curr_conn_in_buff’: Current buffered async connections;