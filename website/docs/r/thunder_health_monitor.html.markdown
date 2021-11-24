---
layout: "thunder"
page_title: "thunder: thunder_health_monitor"
sidebar_current: "docs-thunder-resource-health-monitor"
description: |-
    Provides details about thunder health monitor resource for A10
---

# thunder\_health\_monitor

`thunder_health_monitor` Provides details about thunder health monitor
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_health_monitor" "resourceHealthMonitorTest" {
	name = "string"
dsr_l2_strict = 0
retry = 0
up_retry = 0
override_ipv4 = "string"
override_ipv6 = "string"
override_port = 0
passive = 0
status_code = "string"
passive_interval = 0
sample_threshold = 0
threshold = 0
strict_retry_on_server_err_resp = 0
disable_after_down = 0
interval = 0
timeout = 0
ssl_ciphers = "string"
ssl_ticket = 0
ssl_ticket_lifetime = 0
ssl_version = 0
ssl_dgversion = 0
uuid = "string"
user_tag = "string"
method {  
 icmp {  
 	icmp =  0 
	transparent =  0 
	ipv6 =  "string" 
	ip =  "string" 
	uuid =  "string" 
	}
tcp {  
 	method_tcp =  0 
	tcp_port =  0 
	port_halfopen =  0 
	port_send =  "string" 
port_resp {  
 	port_contains =  "string" 
	}
	maintenance =  0 
	maintenance_text =  "string" 
	uuid =  "string" 
	}
udp {  
 	udp =  0 
	udp_port =  0 
	force_up_with_single_healthcheck =  0 
	uuid =  "string" 
	}
http {  
 	http =  0 
	http_port =  0 
	http_expect =  0 
	http_response_code =  "string" 
	response_code_regex =  "string" 
	http_text =  "string" 
	text_regex =  "string" 
	http_host =  "string" 
	http_maintenance_code =  "string" 
	http_url =  0 
	url_type =  "string" 
	maintenance =  0 
	maintenance_text =  "string" 
	maintenance_text_regex =  "string" 
	url_path =  "string" 
	post_path =  "string" 
	post_type =  "string" 
	http_postdata =  "string" 
	http_postfile =  "string" 
	http_username =  "string" 
	http_password =  0 
	http_password_string =  "string" 
	http_kerberos_auth =  0 
	http_kerberos_realm =  "string" 
http_kerberos_kdc {  
 	http_kerberos_hostip =  "string" 
	http_kerberos_hostipv6 =  "string" 
	http_kerberos_port =  0 
	http_kerberos_portv6 =  0 
	}
	uuid =  "string" 
	}
ftp {  
 	ftp =  0 
	ftp_port =  0 
	ftp_username =  "string" 
	ftp_password =  0 
	ftp_password_string =  "string" 
	uuid =  "string" 
	}
snmp {  
 	snmp =  0 
	snmp_port =  0 
	community =  "string" 
oid {  
 	mib =  "string" 
	asn =  "string" 
	}
operation {  
 	oper_type =  "string" 
	}
	uuid =  "string" 
	}
smtp {  
 	smtp =  0 
	smtp_domain =  "string" 
	smtp_port =  0 
	smtp_starttls =  0 
	mail_from =  "string" 
	rcpt_to =  "string" 
	uuid =  "string" 
	}
dns {  
 	dns =  0 
	dns_ip_key =  0 
	dns_ipv4_addr =  "string" 
	dns_ipv6_addr =  "string" 
	dns_ipv4_port =  0 
dns_ipv4_expect {  
 	dns_ipv4_response =  "string" 
	dns_ipv4_fqdn =  "string" 
	}
	dns_ipv4_recurse =  "string" 
	dns_ipv4_tcp =  0 
	dns_ipv6_port =  0 
dns_ipv6_expect {  
 	dns_ipv6_response =  "string" 
	dns_ipv6_fqdn =  "string" 
	}
	dns_ipv6_recurse =  "string" 
	dns_ipv6_tcp =  0 
	dns_domain =  "string" 
	dns_domain_port =  0 
	dns_domain_type =  "string" 
dns_domain_expect {  
 	dns_domain_response =  "string" 
	dns_domain_fqdn =  "string" 
	dns_domain_ipv4 =  "string" 
	dns_domain_ipv6 =  "string" 
	}
	dns_domain_recurse =  "string" 
	dns_domain_tcp =  0 
	uuid =  "string" 
	}
pop3 {  
 	pop3 =  0 
	pop3_username =  "string" 
	pop3_password =  0 
	pop3_password_string =  "string" 
	pop3_port =  0 
	uuid =  "string" 
	}
imap {  
 	imap =  0 
	imap_port =  0 
	imap_username =  "string" 
	imap_password =  0 
	imap_password_string =  "string" 
	pwd_auth =  0 
	imap_plain =  0 
	imap_cram_md5 =  0 
	imap_login =  0 
	uuid =  "string" 
	}
sip {  
 	sip =  0 
	register =  0 
	sip_port =  0 
	expect_response_code =  "string" 
	sip_tcp =  0 
	uuid =  "string" 
	}
radius {  
 	radius =  0 
	radius_username =  "string" 
	radius_password_string =  "string" 
	radius_secret =  "string" 
	radius_port =  0 
	radius_expect =  0 
	radius_response_code =  "string" 
	uuid =  "string" 
	}
ldap {  
 	ldap =  0 
	ldap_port =  0 
	ldap_security =  "string" 
	ldap_binddn =  "string" 
	ldap_password =  0 
	ldap_password_string =  "string" 
	ldap_run_search =  0 
	base_dn =  "string" 
	ldap_query =  "string" 
	accept_res_ref =  0 
	accept_not_found =  0 
	uuid =  "string" 
	}
rtsp {  
 	rtsp =  0 
	rtspurl =  "string" 
	rtsp_port =  0 
	uuid =  "string" 
	}
database {  
 	database =  0 
	database_name =  "string" 
	db_name =  "string" 
	db_username =  "string" 
	db_password =  0 
	db_password_str =  "string" 
	db_send =  "string" 
	db_receive =  "string" 
	db_row =  0 
	db_column =  0 
	db_receive_integer =  0 
	db_row_integer =  0 
	db_column_integer =  0 
	uuid =  "string" 
	}
external {  
 	external =  0 
	ext_program =  "string" 
	shared_partition_program =  0 
	ext_program_shared =  "string" 
	ext_port =  0 
	ext_arguments =  "string" 
	ext_preference =  0 
	uuid =  "string" 
	}
ntp {  
 	ntp =  0 
	ntp_port =  0 
	uuid =  "string" 
	}
kerberos_kdc {  
 kerberos_cfg {  
 	kinit =  0 
	kinit_pricipal_name =  "string" 
	kinit_password =  "string" 
	kinit_kdc =  "string" 
	tcp_only =  0 
	kadmin =  0 
	kadmin_realm =  "string" 
	kadmin_pricipal_name =  "string" 
	kadmin_password =  "string" 
	kadmin_server =  "string" 
	kadmin_kdc =  "string" 
	kpasswd =  0 
	kpasswd_pricipal_name =  "string" 
	kpasswd_password =  "string" 
	kpasswd_server =  "string" 
	kpasswd_kdc =  "string" 
	}
	uuid =  "string" 
	}
https {  
 	https =  0 
	web_port =  0 
	disable_sslv2hello =  0 
	https_host =  "string" 
	sni =  0 
	https_expect =  0 
	https_response_code =  "string" 
	response_code_regex =  "string" 
	https_text =  "string" 
	text_regex =  "string" 
	https_url =  0 
	url_type =  "string" 
	url_path =  "string" 
	post_path =  "string" 
	post_type =  "string" 
	https_postdata =  "string" 
	https_postfile =  "string" 
	https_maintenance_code =  "string" 
	maintenance =  0 
	maintenance_text =  "string" 
	maintenance_text_regex =  "string" 
	https_username =  "string" 
	https_server_cert_name =  "string" 
	https_password =  0 
	https_password_string =  "string" 
	https_kerberos_auth =  0 
	https_kerberos_realm =  "string" 
https_kerberos_kdc {  
 	https_kerberos_hostip =  "string" 
	https_kerberos_hostipv6 =  "string" 
	https_kerberos_port =  0 
	https_kerberos_portv6 =  0 
	}
	cert_key_shared =  0 
	cert =  "string" 
	key =  "string" 
	key_pass_phrase =  0 
	key_phrase =  "string" 
	uuid =  "string" 
	}
tacplus {  
 	tacplus =  0 
	tacplus_username =  "string" 
	tacplus_password =  0 
	tacplus_password_string =  "string" 
	tacplus_secret =  0 
	tacplus_secret_string =  "string" 
	tacplus_port =  0 
	tacplus_type =  "string" 
	uuid =  "string" 
	}
compound {  
 	compound =  0 
	rpn_string =  "string" 
	uuid =  "string" 
	}
	}
 
}

```

## Argument Reference

* `monitor` - Define the Health Monitor object
* `name` - Monitor Name
* `dsr-l2-strict` - Enable strict L2dsr health-check
* `retry` - Specify the Healthcheck Retries (Retry Count (default 3))
* `up-retry` - Specify the Healthcheck Retries before declaring target up (Up-retry count (default 1))
* `override-ipv4` - Override implicitly inherited IPv4 address from target
* `override-ipv6` - Override implicitly inherited IPv6 address from target
* `override-port` - Override implicitly inherited port from target (Port number (1-65534))
* `passive` - Specify passive mode
* `status-code` - 'status-code-2xx': Enable passive mode with 2xx http status code; 'status-code-non-5xx': Enable passive mode with non-5xx http status code;
* `passive-interval` - Interval to do manual health checking while in passive mode (Specify value in seconds (Default is 10 s))
* `sample-threshold` - Number of samples in one epoch above which passive HC is enabled. If below or equal to the threshold, passive HC is disabled (Specify number of samples in one second (Default is 50). If the number of samples is 0, no action is taken)
* `threshold` - Threshold percentage above which passive mode is enabled (Specify percentage (Default is 75%))
* `strict-retry-on-server-err-resp` - Require strictly retry
* `disable-after-down` - Disable the target if health check failed
* `interval` - Specify the Healthcheck Interval (Interval Value, in seconds (default 5))
* `timeout` - Specify the Healthcheck Timeout (Timeout Value, in seconds(default 5), Timeout should be less than or equal to interval)
* `ssl-ciphers` - Specify OpenSSL Cipher Suite name(s) for Health check (OpenSSL Cipher Suite(s) (Eg: AES128-SHA256), if the cipher is invalid, would give information at HM down reason)
* `ssl-ticket` - Enable SSL-Ticket Session Resumption
* `ssl-ticket-lifetime` - SSL-Ticket lifetime (seconds)
* `ssl-version` - TLS/SSL version (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1, 33-TLSv1.2 and 34-TLSv1.3)
* `ssl-dgversion` - Lower TLS/SSL version can be downgraded
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `icmp` - ICMP type
* `transparent` - Apply transparent mode
* `ipv6` - Specify IPv6 address of destination behind monitored node
* `ip` - Specify IPv4 address of destination behind monitored node
* `method-tcp` - TCP type
* `tcp-port` - Specify TCP port (Specify port number)
* `port-halfopen` - Set TCP SYN check
* `port-send` - Send a string to server (Specify the string)
* `port-contains` - Mark server up if response string contains string (Specify the string)
* `maintenance` - Specify response text for maintenance
* `maintenance-text` - Specify text for maintenance
* `udp` - UDP type
* `udp-port` - Specify UDP port (Specify port number)
* `force-up-with-single-healthcheck` - Force Up with no response at the first time
* `http` - HTTP type
* `http-port` - Specify HTTP Port (Specify port number (default 80))
* `http-expect` - Specify what you expect from the response message
* `http-response-code` - Specify response code range (e.g. 200,400-430) (Format is xx,xx-xx (xx between [100, 899]))
* `response-code-regex` - Specify response code range with Regex (code with Regex, such as [2-5][0-9][0-9])
* `http-text` - Specify text expected
* `text-regex` - Specify text expected  with Regex
* `http-host` - Specify "Host:" header used in request (enclose IPv6 address in [])
* `http-maintenance-code` - Specify response code for maintenance (Format is xx,xx-xx (xx between [100, 899]))
* `http-url` - Specify URL string, default is GET /
* `url-type` - 'GET': HTTP GET method; 'POST': HTTP POST method; 'HEAD': HTTP HEAD method;
* `maintenance-text-regex` - Specify Regex text for maintenance
* `url-path` - Specify URL path, default is "/"
* `post-path` - Specify URL path, default is "/"
* `post-type` - 'postdata': Specify the HTTP post data; 'postfile': Specify the HTTP post data;
* `http-postdata` - Specify the HTTP post data (Input post data here)
* `http-postfile` - Specify the HTTP post data (Input post data file name here)
* `http-username` - Specify the username
* `http-password` - Specify the user password
* `http-password-string` - Specify password, '' means empty password
* `http-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `http-kerberos-auth` - Http Kerberos Auth
* `http-kerberos-realm` - Specify realm of Kerberos server
* `http-kerberos-hostip` - Kdc's hostname(length:1-31) or IP address
* `http-kerberos-hostipv6` - Server's IPV6 address
* `http-kerberos-port` - Specify the kdc port
* `http-kerberos-portv6` - Specify the kdc port
* `ftp` - FTP type
* `ftp-port` - Specify FTP port (Specify port number, default is 21)
* `ftp-username` - Specify the username
* `ftp-password` - Specify the user password
* `ftp-password-string` - Specify the user password, '' means empty password
* `ftp-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `snmp` - SNMP type
* `snmp-port` - Specify SNMP port, default is 161 (Port Number)
* `community` - Specify SNMP community, default is "public" (Community String)
* `mib` - 'sysDescr': The MIB-2 OID of system description, 1.1.0; 'sysUpTime': The MIB-2 OID of system up time, 1.3.0; 'sysName': The MIB-2 OID of system nume, 1.5.0;
* `asn` - Specify the format in ASN.1 style
* `oper-type` - 'getnext': Get-Next-Request command; 'get': Get-Request command;
* `smtp` - SMTP type
* `smtp-domain` - Specify domain name of 'helo' command
* `smtp-port` - Specify SMTP port, default is 25 (Port Number (default 25))
* `smtp-starttls` - Check the STARTTLS support at helo response
* `mail-from` - Specify SMTP Sender
* `rcpt-to` - Specify SMTP Receiver
* `dns` - DNS type
* `dns-ip-key` - Reverse DNS lookup (Specify IPv4 or IPv6 address)
* `dns-ipv4-addr` - Specify IPv4 address
* `dns-ipv6-addr` - Specify IPv6 address
* `dns-ipv4-port` - Specify DNS port, default is 53 (DNS Port(default 53))
* `dns-ipv4-response` - Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))
* `dns-ipv4-fqdn` - Specify fully qualified domain name expected in DNS response answer
* `dns-ipv4-recurse` - 'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;
* `dns-ipv4-tcp` - Configure DNS transport over TCP, default is UDP
* `dns-ipv6-port` - Specify DNS port, default is 53 (DNS Port(default 53))
* `dns-ipv6-response` - Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))
* `dns-ipv6-fqdn` - Specify fully qualified domain name expected in DNS response answer
* `dns-ipv6-recurse` - 'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;
* `dns-ipv6-tcp` - Configure DNS transport over TCP, default is UDP
* `dns-domain` - Specify fully qualified domain name of the host
* `dns-domain-port` - Specify DNS port, default is 53 (DNS Port(default 53))
* `dns-domain-type` - 'A': Used for storing Ipv4 address (default); 'CNAME': Canonical name for a DNS alias; 'SOA': Start of authority; 'PTR': Domain name pointer; 'MX': Mail exchanger; 'TXT': Text string; 'AAAA': Used for storing Ipv6 128-bits address;
* `dns-domain-response` - Specify response code range (e.g. 0,1-5) (Format is xx,xx-xx (xx between [0,15]))
* `dns-domain-fqdn` - Specify fully qualified domain name expected in DNS response answer
* `dns-domain-ipv4` - Specify expected resolved IPv4 address
* `dns-domain-ipv6` - Specify expected resolved IPv6 address
* `dns-domain-recurse` - 'enabled': Set the recursion bit; 'disabled': Clear the recursion bit;
* `dns-domain-tcp` - Configure DNS transport over TCP, default is UDP
* `pop3` - POP3 type
* `pop3-username` - Specify the username
* `pop3-password` - Specify the user password
* `pop3-password-string` - Specify the user password, '' means empty password
* `pop3-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `pop3-port` - Specify the POP3 port, default is 110 (Port Number (default 110))
* `imap` - IMAP type
* `imap-port` - Specify the IMAP port, default is 143 (Port Number (default 143))
* `imap-username` - Specify the username
* `imap-password` - Specify the user password
* `imap-password-string` - Configure password, '' means empty password
* `imap-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `pwd-auth` - Specify the Authentication method
* `imap-plain` - Plain text
* `imap-cram-md5` - Challenge-response authentication mechanism
* `imap-login` - Simple login
* `sip` - SIP type
* `register` - Send SIP REGISTER message, default is to send OPTION message
* `sip-port` - Specify the SIP port, default is 5060 (Port Number)
* `expect-response-code` - Specify accepted response codes (e.g. 200, 400-430, any) (Format is xxx,xxx-xxx,any (xxx between [100,899]))
* `sip-tcp` - Use TCP for transmission, default is UDP
* `radius` - RADIUS type
* `radius-username` - Specify the username
* `radius-password-string` - Configure password, '' means empty password
* `radius-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `radius-secret` - Configure shared secret of RADIUS server
* `radius-secret-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `radius-port` - Specify the RADIUS port, default is 1812 (Port number (default 1812))
* `radius-expect` - Specify what you expect from the response message
* `radius-response-code` - Specify response code range (e.g. 2,4-7) (Format is xx,xx-xx (xx between [1, 13]))
* `ldap` - LDAP type
* `ldap-port` - Specify the LDAP port (Speciry port number, default is 389, or 636 if LDAP over SSL)
* `ldap-security` - 'overssl': Set LDAP over SSL; 'StartTLS': LDAP switch to TLS;
* `ldap-binddn` - Specify the distinguished name for bindRequest (LDAP DN distinguished name)
* `ldap-password` - Specify the user password
* `ldap-password-string` - Configure password, '' means empty password
* `ldap-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `ldap-run-search` - Specify a query to be executed
* `base_dn` - Specify LDAP DN distinguished name
* `ldap-query` - LDAP query to be excuted
* `accept_res_ref` - Mark server up on receiving a search result reference response
* `accept_not_found` - Mark server up on receiving a not-found response
* `rtsp` - RTSP type
* `rtspurl` - Specify URL string (Specify the path on the server)
* `rtsp-port` - Specify RTSP port, default is 554 (Port Number (default 554))
* `database` - DATABASE type
* `database-name` - 'mssql': Specify MSSQL database; 'mysql': Specify MySQL database; 'oracle': Specify Oracle database; 'postgresql': Specify PostgreSQL database;
* `db-name` - Specify the database name
* `db-username` - Specify the username
* `db-password` - Specify the user password
* `db-password-str` - Configure password
* `db-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `db-send` - Specify the SQL query
* `db-receive` - Specify the response string
* `db-row` - Specify the row number for receiving
* `db-column` - Specify the column number for receiving
* `db-receive-integer` - Specify the response integer
* `db-row-integer` - Specify the row number for receiving
* `db-column-integer` - Specify the column number for receiving
* `external` - EXTERNAL type
* `ext-program` - Specify external application (Program name)
* `shared-partition-program` - external application from shared partition
* `ext-program-shared` - Specify external application (Program name)
* `ext-port` - Specify the server port (Port Number)
* `ext-arguments` - Specify external application's arguments (Application arguments)
* `ext-preference` - Get server's perference
* `ntp` - NTP type
* `ntp-port` - Specify the NTP port, default is 123 (Port Number (default 123))
* `kinit` - Kerberos KDC
* `kinit-pricipal-name` - Specify the principal name
* `kinit-password` - Password
* `kinit-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `kinit-kdc` - Specify the kdc server, host|ip [:port]
* `tcp-only` - Specify the kerberos tcp only
* `kadmin` - Kerberos admin
* `kadmin-realm` - Specify the realm
* `kadmin-pricipal-name` - Specify the principal name
* `kadmin-password` - Password
* `kadmin-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `kadmin-server` - Specify the admin server, host|ip [:port]
* `kadmin-kdc` - Specify the kdc server, host|ip [:port]
* `kpasswd` - Kerberos change passwd
* `kpasswd-pricipal-name` - Specify the principal name
* `kpasswd-password` - Password
* `kpasswd-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `kpasswd-server` - Specify the Kerberos password server, host|ip [:port]
* `kpasswd-kdc` - Specify the kdc server, host|ip [:port]
* `https` - HTTPS type
* `web-port` - Specify HTTPS port (Port Number (default 443))
* `disable-sslv2hello` - Disable SSLv2Hello for HTTPs
* `https-host` - Specify "Host:" header used in request (enclose IPv6 address in [])
* `sni` - Server Name Indication for HTTPs
* `https-expect` - Specify what you expect from the response message
* `https-response-code` - Specify response code range (e.g. 200,400-430) (Format is xx,xx-xx (xx between [100, 899])
* `https-text` - Specify text expected
* `https-url` - Specify URL string, default is GET /
* `https-postdata` - Specify the HTTP post data (Input post data here)
* `https-postfile` - Specify the HTTP post data (Input post data file name here)
* `https-maintenance-code` - Specify response code for maintenance (Format is xx,xx-xx (xx between [100, 899])
* `https-username` - Specify the username
* `https-server-cert-name` - Expect Server Cert commonName
* `https-password` - Specify the user password
* `https-password-string` - Configure password, '' means empty password
* `https-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `https-kerberos-auth` - Https Kerberos Auth
* `https-kerberos-realm` - Specify realm of Kerberos server
* `https-kerberos-hostip` - Kdc's hostname(length:1-31) or IP address
* `https-kerberos-hostipv6` - Server's IPV6 address
* `https-kerberos-port` - Specify the kdc port
* `https-kerberos-portv6` - Specify the kdc port
* `cert-key-shared` - Select shared partition
* `cert` - Specify client certificate (Certificate name)
* `key` - Specify client private key (Key name)
* `key-pass-phrase` - Client private key password phrase
* `key-phrase` - Password Phrase
* `https-key-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `tacplus` - TACACS+ type
* `tacplus-username` - Specify the username
* `tacplus-password` - Specify the user password
* `tacplus-password-string` - Configure password, '' means empty password
* `tacplus-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `tacplus-secret` - Specify the shared secret of TACACS+ server
* `tacplus-secret-string` - Shared Crypto Key
* `tacplus-port` - Specify the TACACS+ port, default 49 (Port number (default 49))
* `tacplus-type` - 'inbound-ascii-login': Specify Inbound ASCII Login type;
* `compound` - Compound type
* `rpn-string` - Enter Reverse Polish Notation, example: sub hm1 sub hm2 and

