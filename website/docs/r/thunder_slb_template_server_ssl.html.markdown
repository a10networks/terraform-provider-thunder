---
layout: "thunder"
page_title: "thunder: thunder_slb_template_server_ssl"
sidebar_current: "docs-thunder-resource-slb-template-server-ssl"
description: |-
    Provides details about thunder slb template server ssl resource for A10
---

# thunder\_slb\_template\_server\_ssl

`thunder_slb_template_server_ssl` Provides details about thunder slb template server ssl
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_server_ssl" "resourceSlbTemplateServerSslTest" {
	name = "string"
ca-certs {   
	ca_cert =  "string" 
	ca_cert_partition_shared =  0 
	server_ocsp_srvr =  "string" 
	server_ocsp_sg =  "string" 
	}
server_name = "string"
crl-certs {   
	crl =  "string" 
	crl_partition_shared =  0 
	}
cipher-without-prio-list {   
	cipher_wo_prio =  "string" 
	}
dh_type = "string"
ec-list {   
	ec =  "string" 
	}
enable_tls_alert_logging = 0
alert_type = "string"
handshake_logging_enable = 0
close_notify = 0
forward_proxy_enable = 0
session_ticket_enable = 0
version = 0
dgversion = 0
server-certificate-error {   
	error_type =  "string" 
	}
ssli_logging = 0
sslilogging = "string"
ocsp_stapling = 0
use_client_sni = 0
renegotiation_disable = 0
session_cache_size = 0
session_cache_timeout = 0
cipher_template = "string"
shared_partition_cipher_template = 0
template_cipher_shared = "string"
enable_ssli_ftp_alg = 0
early_data = 0
uuid = "string"
user_tag = "string"
certificate {  
 	cert =  "string" 
	key =  "string" 
	passphrase =  "string" 
	shared =  0 
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `server-ssl` - Server Side SSL Template
* `name` - Server SSL Template Name
* `ca-cert` - Specify CA certificate
* `ca-cert-partition-shared` - CA Certificate Partition Shared
* `server-ocsp-srvr` - Specify authentication server
* `server-ocsp-sg` - Specify service-group (Service group name)
* `server-name` - Specify Server Name
* `crl` - Certificate Revocation Lists (Certificate Revocation Lists file name)
* `crl-partition-shared` - Certificate Revocation Lists Partition Shared
* `cipher-wo-prio` - 'SSL3_RSA_DES_192_CBC3_SHA': TLS_RSA_WITH_3DES_EDE_CBC_SHA (0x000A); 'SSL3_RSA_RC4_128_MD5': TLS_RSA_WITH_RC4_128_MD5 (0x0004); 'SSL3_RSA_RC4_128_SHA': TLS_RSA_WITH_RC4_128_SHA (0x0005); 'TLS1_RSA_AES_128_SHA': TLS_RSA_WITH_AES_128_CBC_SHA (0x002F); 'TLS1_RSA_AES_256_SHA': TLS_RSA_WITH_AES_256_CBC_SHA (0x0035); 'TLS1_RSA_AES_128_SHA256': TLS_RSA_WITH_AES_128_CBC_SHA256 (0x003C); 'TLS1_RSA_AES_256_SHA256': TLS_RSA_WITH_AES_256_CBC_SHA256 (0x003D); 'TLS1_DHE_RSA_AES_128_GCM_SHA256': TLS_DHE_RSA_WITH_AES_128_GCM_SHA256 (0x009E); 'TLS1_DHE_RSA_AES_128_SHA': TLS_DHE_RSA_WITH_AES_128_CBC_SHA (0x0033); 'TLS1_DHE_RSA_AES_128_SHA256': TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 (0x0067); 'TLS1_DHE_RSA_AES_256_GCM_SHA384': TLS_DHE_RSA_WITH_AES_256_GCM_SHA384 (0x009F); 'TLS1_DHE_RSA_AES_256_SHA': TLS_DHE_RSA_WITH_AES_256_CBC_SHA (0x0039); 'TLS1_DHE_RSA_AES_256_SHA256': TLS_DHE_RSA_WITH_AES_256_CBC_SHA256 (0x006B); 'TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 (0xC02B); 'TLS1_ECDHE_ECDSA_AES_128_SHA': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA (0xC009); 'TLS1_ECDHE_ECDSA_AES_128_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 (0xC023); 'TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 (0xC02C); 'TLS1_ECDHE_ECDSA_AES_256_SHA': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA (0xC00A); 'TLS1_ECDHE_RSA_AES_128_GCM_SHA256': TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 (0xC02F); 'TLS1_ECDHE_RSA_AES_128_SHA': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA (0xC013); 'TLS1_ECDHE_RSA_AES_128_SHA256': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256 (0xC027); 'TLS1_ECDHE_RSA_AES_256_GCM_SHA384': TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 (0xC030); 'TLS1_ECDHE_RSA_AES_256_SHA': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA (0xC014); 'TLS1_RSA_AES_128_GCM_SHA256': TLS_RSA_WITH_AES_128_GCM_SHA256 (0x009C); 'TLS1_RSA_AES_256_GCM_SHA384': TLS_RSA_WITH_AES_256_GCM_SHA384 (0x009D); 'TLS1_ECDHE_RSA_AES_256_SHA384': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384 (0xC028); 'TLS1_ECDHE_ECDSA_AES_256_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384 (0xC024); 'TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA8); 'TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA9); 'TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256': TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCAA);
* `dh-type` - '1024': 1024; '1024-dsa': 1024-dsa; '2048': 2048;
* `ec` - 'secp256r1': X9_62_prime256v1; 'secp384r1': secp384r1;
* `enable-tls-alert-logging` - Enable TLS alert logging
* `alert-type` - 'fatal': Log fatal alerts;
* `handshake-logging-enable` - Enable SSL handshake logging
* `close-notify` - Send close notification when terminate connection
* `forward-proxy-enable` - Enable SSL forward proxy
* `session-ticket-enable` - Enable server side session ticket support
* `version` - TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)
* `dgversion` - Lower TLS/SSL version can be downgraded
* `error-type` - 'email': Notify the error via email; 'ignore': Ignore the error, which mean the connection can continue; 'logging': Log the error; 'trap': Notify the error by SNMP trap;
* `ssli-logging` - SSLi logging level, default is error logging only
* `sslilogging` - 'disable': Disable all logging; 'all': enable all logging(error, info);
* `ocsp-stapling` - Enable ocsp-stapling support
* `use-client-sni` - use client SNI
* `renegotiation-disable` - Disable SSL renegotiation
* `session-cache-size` - Session Cache Size (Maximum cache size. Default value 0 (Session ID reuse disabled))
* `session-cache-timeout` - Session Cache Timeout (Timeout value, in seconds. Default no timeout.)
* `cipher-template` - Cipher Template Name
* `shared-partition-cipher-template` - Reference a cipher template from shared partition
* `template-cipher-shared` - Cipher Template Name
* `enable-ssli-ftp-alg` - Enable SSLi FTP over TLS support at which port
* `early-data` - Enable TLS 1.3 early data (0-RTT)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `cert` - Certificate Name
* `key` - Client private-key (Key Name)
* `passphrase` - Password Phrase
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `shared` - Client Certificate and Key Partition Shared

