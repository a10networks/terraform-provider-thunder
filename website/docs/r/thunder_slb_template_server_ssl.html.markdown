---
layout: "thunder"
page_title: "thunder: thunder_slb_template_server_ssl"
sidebar_current: "docs-thunder-resource-slb-template-server-ssl"
description: |-
    Provides details about thunder slb template server-ssl resource for A10
---

# thunder\_slb\_template\_server\_ssl

`thunder_slb_template_server_ssl` provides details about slb template server_ssl
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_server_ssl" "Slb_Template_Server_Ssl_Test" {

key = "string"
cipher_template = "string"
sslilogging = "string"
user_tag = "string"
passphrase = "string"
ocsp_stapling = 0
crl-certs {   
        crl =  "string" 
        }
uuid = "string"
certificate {  
        encrypted =  "Unknown Type: encrypted" 
        cert =  "string" 
        key =  "string" 
        passphrase =  "string" 
        uuid =  "string" 
        }
server_name = "string"
key_shared_str = "string"
template_cipher_shared = "string"
dgversion = 0
cert_shared_str = "string"
version = 0
ec-list {   
        ec =  "string" 
        }
encrypted = "Unknown Type: encrypted"
ssli_logging = 0
session_cache_size = 0
dh_type = "string"
use_client_sni = 0
renegotiation_disable = 0
forward_proxy_enable = 0
session_cache_timeout = 0
key_shared_encrypted = "Unknown Type: encrypted"
cipher-without-prio-list {   
        cipher_wo_prio =  "string" 
        }
ca-certs {   
        ca_cert =  "string" 
        ca_cert_partition_shared =  0 
        server_ocsp_sg =  "string" 
        server_ocsp_srvr =  "string" 
        }
name = "string"
shared_partition_cipher_template = 0
enable_tls_alert_logging = 0
session_ticket_enable = 0
alert_type = "string"
cert = "string"
handshake_logging_enable = 0
early_data = 0
server-certificate-error {   
        error_type =  "string" 
        }
close_notify = 0
key_shared_passphrase = "string"
 
}
```

## Argument Reference

* `alert_type` - ‘fatal’: Log fatal alerts;
* `cert` - Specify Client certificate (Certificate Name)
* `cert_shared_str` - Certificate Name
* `cipher_template` - Cipher Template (Cipher Config Name)
* `close_notify` - Send close notification when terminate connection
* `dgversion` - Lower TLS/SSL version can be downgraded
* `dh_type` - ‘1024’: 1024; ‘1024-dsa’: 1024-dsa; ‘2048’: 2048;
* `enable_tls_alert_logging` - Enable TLS alert logging
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `forward_proxy_enable` - Enable SSL forward proxy
* `handshake_logging_enable` - Enable SSL handshake logging
* `key` - Client private-key (Key Name)
* `key_shared_encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `key_shared_passphrase` - Password Phrase
* `key_shared_str` - Key Name
* `name` - Server SSL Template Name
* `ocsp_stapling` - Enable ocsp-stapling support
* `passphrase` - Password Phrase
* `renegotiation_disable` - Disable SSL renegotiation
* `session_cache_size` - Session Cache Size (Specify 0 to disable Session ID reuse.)
* `session_cache_timeout` - Session Cache Timeout (Timeout value, in seconds)
* `session_ticket_enable` - Enable server side session ticket support
* `shared_partition_cipher_template` - Reference a cipher template from shared partition
* `ssli_logging` - SSLi logging level, default is error logging only
* `sslilogging` - ‘disable’: Disable all logging; ‘all’: enable all logging(error, info);
* `template_cipher_shared` - Cipher Template Name
* `use_client_sni` - use client SNI
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `version` - TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)
* `crl` - Certificate Revocation Lists (Certificate Revocation Lists file name)
* `ec` - ‘secp256r1’: X9_62_prime256v1; ‘secp384r1’: secp384r1;
* `error_type` - ‘email’: Notify the error via email; ‘ignore’: Ignore the error, which mean the connection can continue; ‘logging’: Log the error; ‘trap’: Notify the error by SNMP trap;
* `ca_cert` - Specify CA certificate
* `ca_cert_partition_shared` - CA Certificate Partition Shared
* `server_ocsp_sg` - Specify service-group (Service group name)
* `server_ocsp_srvr` - Specify authentication server
* `cipher_wo_prio` - ‘SSL3_RSA_DES_192_CBC3_SHA’: SSL3_RSA_DES_192_CBC3_SHA; ‘SSL3_RSA_RC4_128_MD5’: SSL3_RSA_RC4_128_MD5; ‘SSL3_RSA_RC4_128_SHA’: SSL3_RSA_RC4_128_SHA; ‘TLS1_RSA_AES_128_SHA’: TLS1_RSA_AES_128_SHA; ‘TLS1_RSA_AES_256_SHA’: TLS1_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_SHA256’: TLS1_RSA_AES_128_SHA256; ‘TLS1_RSA_AES_256_SHA256’: TLS1_RSA_AES_256_SHA256; ‘TLS1_DHE_RSA_AES_128_GCM_SHA256’: TLS1_DHE_RSA_AES_128_GCM_SHA256; ‘TLS1_DHE_RSA_AES_128_SHA’: TLS1_DHE_RSA_AES_128_SHA; ‘TLS1_DHE_RSA_AES_128_SHA256’: TLS1_DHE_RSA_AES_128_SHA256; ‘TLS1_DHE_RSA_AES_256_GCM_SHA384’: TLS1_DHE_RSA_AES_256_GCM_SHA384; ‘TLS1_DHE_RSA_AES_256_SHA’: TLS1_DHE_RSA_AES_256_SHA; ‘TLS1_DHE_RSA_AES_256_SHA256’: TLS1_DHE_RSA_AES_256_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256’: TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_SHA’: TLS1_ECDHE_ECDSA_AES_128_SHA; ‘TLS1_ECDHE_ECDSA_AES_128_SHA256’: TLS1_ECDHE_ECDSA_AES_128_SHA256; ‘TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384’: TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_ECDSA_AES_256_SHA’: TLS1_ECDHE_ECDSA_AES_256_SHA; ‘TLS1_ECDHE_RSA_AES_128_GCM_SHA256’: TLS1_ECDHE_RSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_RSA_AES_128_SHA’: TLS1_ECDHE_RSA_AES_128_SHA; ‘TLS1_ECDHE_RSA_AES_128_SHA256’: TLS1_ECDHE_RSA_AES_128_SHA256; ‘TLS1_ECDHE_RSA_AES_256_GCM_SHA384’: TLS1_ECDHE_RSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_RSA_AES_256_SHA’: TLS1_ECDHE_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_GCM_SHA256’: TLS1_RSA_AES_128_GCM_SHA256; ‘TLS1_RSA_AES_256_GCM_SHA384’: TLS1_RSA_AES_256_GCM_SHA384;
