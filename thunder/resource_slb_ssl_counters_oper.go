package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCountersOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_counters_oper`: Operational Status for the object ssl-counters\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCountersOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vserver": {
							Type: schema.TypeString, Optional: true, Description: "virtual server name",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Port",
						},
						"cumulative_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Cumulative SSL sessions",
						},
						"ssl3_rsa_des_192_cbc3_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "SSL3_RSA_DES_192_CBC3_SHA Cipher ID",
						},
						"ssl3_rsa_des_40_cbc_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "SSL3_RSA_DES_40_CBC_SHA Cipher ID",
						},
						"ssl3_rsa_des_64_cbc_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "SSL3_RSA_DES_64_CBC_SHA Cipher ID",
						},
						"ssl3_rsa_rc4_128_md5_id": {
							Type: schema.TypeString, Optional: true, Description: "SSL3_RSA_RC4_128_MD5 Cipher ID",
						},
						"ssl3_rsa_rc4_128_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "SSL3_RSA_RC4_128_SHA Cipher ID",
						},
						"ssl3_rsa_rc4_40_md5_id": {
							Type: schema.TypeString, Optional: true, Description: "SSL3_RSA_RC4_40_MD5 Cipher ID",
						},
						"tls1_dhe_rsa_aes_128_gcm_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_DHE_RSA_AES_128_GCM_SHA256 Cipher ID",
						},
						"tls1_dhe_rsa_aes_128_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_DHE_RSA_AES_128_SHA Cipher ID",
						},
						"tls1_dhe_rsa_aes_128_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_DHE_RSA_AES_128_SHA256 Cipher ID",
						},
						"tls1_dhe_rsa_aes_256_gcm_sha384_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_DHE_RSA_AES_256_GCM_SHA384 Cipher ID",
						},
						"tls1_dhe_rsa_aes_256_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_DHE_RSA_AES_256_SHA Cipher ID",
						},
						"tls1_dhe_rsa_aes_256_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_DHE_RSA_AES_256_SHA256 Cipher ID",
						},
						"tls1_ecdhe_ecdsa_aes_128_gcm_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256 Cipher ID",
						},
						"tls1_ecdhe_ecdsa_aes_128_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_SHA Cipher ID",
						},
						"tls1_ecdhe_ecdsa_aes_128_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_SHA256 Cipher ID",
						},
						"tls1_ecdhe_ecdsa_aes_256_sha384_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_SHA384 Cipher ID",
						},
						"tls1_ecdhe_ecdsa_aes_256_gcm_sha384_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384 Cipher ID",
						},
						"tls1_ecdhe_ecdsa_aes_256_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_SHA Cipher ID",
						},
						"tls1_ecdhe_rsa_aes_128_gcm_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_GCM_SHA256 Cipher ID",
						},
						"tls1_ecdhe_rsa_aes_128_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_SHA Cipher ID",
						},
						"tls1_ecdhe_rsa_aes_128_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_SHA256 Cipher ID",
						},
						"tls1_ecdhe_rsa_aes_256_sha384_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_SHA384 Cipher ID",
						},
						"tls1_ecdhe_rsa_aes_256_gcm_sha384_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_GCM_SHA384 Cipher ID",
						},
						"tls1_ecdhe_rsa_aes_256_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_SHA Cipher ID",
						},
						"tls1_rsa_aes_128_gcm_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_AES_128_GCM_SHA256 Cipher ID",
						},
						"tls1_rsa_aes_128_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_AES_128_SHA Cipher ID",
						},
						"tls1_rsa_aes_128_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_AES_128_SHA256 Cipher ID",
						},
						"tls1_rsa_aes_256_gcm_sha384_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_AES_256_GCM_SHA384 Cipher ID",
						},
						"tls1_rsa_aes_256_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_AES_256_SHA Cipher ID",
						},
						"tls1_rsa_aes_256_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_AES_256_SHA256 Cipher ID",
						},
						"tls1_rsa_export1024_rc4_56_md5_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_EXPORT1024_RC4_56_MD5 Cipher ID",
						},
						"tls1_rsa_export1024_rc4_56_sha_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_RSA_EXPORT1024_RC4_56_SHA Cipher ID",
						},
						"tls1_ecdhe_rsa_chacha20_poly1305_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256 Cipher ID",
						},
						"tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256 Cipher ID",
						},
						"tls1_dhe_rsa_chacha20_poly1305_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256 Cipher ID",
						},
						"tls13_aes_128_gcm_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS13_AES_128_GCM_SHA256 Cipher ID",
						},
						"tls13_aes_256_gcm_sha384_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS13_AES_256_GCM_SHA384 Cipher ID",
						},
						"tls13_chacha20_poly1305_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS13_CHACHA20_POLY1305_SHA256 Cipher ID",
						},
						"tls1_ecdhe_sm2_sms4_sm3_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_SM3 Cipher ID",
						},
						"tls1_ecdhe_sm2_sms4_gcm_sm3_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_GCM_SM3 Cipher ID",
						},
						"tls1_ecdhe_sm2_sms4_sha256_id": {
							Type: schema.TypeString, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_SHA256 Cipher ID",
						},
						"ssl3_rsa_des_192_cbc3_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_DES_192_CBC3_SHA Successes",
						},
						"ssl3_rsa_des_40_cbc_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_DES_40_CBC_SHA Successes",
						},
						"ssl3_rsa_des_64_cbc_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_DES_64_CBC_SHA Successes",
						},
						"ssl3_rsa_rc4_128_md5_successes": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_RC4_128_MD5 Successes",
						},
						"ssl3_rsa_rc4_128_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_RC4_128_SHA Successes",
						},
						"ssl3_rsa_rc4_40_md5_successes": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_RC4_40_MD5 Successes",
						},
						"tls1_dhe_rsa_aes_128_gcm_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_128_GCM_SHA256 Successes",
						},
						"tls1_dhe_rsa_aes_128_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_128_SHA Successes",
						},
						"tls1_dhe_rsa_aes_128_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_128_SHA256 Successes",
						},
						"tls1_dhe_rsa_aes_256_gcm_sha384_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_256_GCM_SHA384 Successes",
						},
						"tls1_dhe_rsa_aes_256_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_256_SHA Successes",
						},
						"tls1_dhe_rsa_aes_256_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_256_SHA256 Successes",
						},
						"tls1_ecdhe_ecdsa_aes_128_gcm_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256 Successes",
						},
						"tls1_ecdhe_ecdsa_aes_128_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_SHA Successes",
						},
						"tls1_ecdhe_ecdsa_aes_128_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_SHA256 Successes",
						},
						"tls1_ecdhe_ecdsa_aes_256_sha384_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_SHA384 Successes",
						},
						"tls1_ecdhe_ecdsa_aes_256_gcm_sha384_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384 Successes",
						},
						"tls1_ecdhe_ecdsa_aes_256_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_SHA Successes",
						},
						"tls1_ecdhe_rsa_aes_128_gcm_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_GCM_SHA256 Successes",
						},
						"tls1_ecdhe_rsa_aes_128_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_SHA Successes",
						},
						"tls1_ecdhe_rsa_aes_128_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_SHA256 Successes",
						},
						"tls1_ecdhe_rsa_aes_256_sha384_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_SHA384 Successes",
						},
						"tls1_ecdhe_rsa_aes_256_gcm_sha384_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_GCM_SHA384 Successes",
						},
						"tls1_ecdhe_rsa_aes_256_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_SHA Successes",
						},
						"tls1_rsa_aes_128_gcm_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_128_GCM_SHA256 Successes",
						},
						"tls1_rsa_aes_128_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_128_SHA Successes",
						},
						"tls1_rsa_aes_128_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_128_SHA256 Successes",
						},
						"tls1_rsa_aes_256_gcm_sha384_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_256_GCM_SHA384 Successes",
						},
						"tls1_rsa_aes_256_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_256_SHA Successes",
						},
						"tls1_rsa_aes_256_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_256_SHA256 Successes",
						},
						"tls1_rsa_export1024_rc4_56_md5_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_EXPORT1024_RC4_56_MD5 Successes",
						},
						"tls1_rsa_export1024_rc4_56_sha_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_EXPORT1024_RC4_56_SHA Successes",
						},
						"tls1_ecdhe_rsa_chacha20_poly1305_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256 Cipher successes",
						},
						"tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256 Cipher successes",
						},
						"tls1_dhe_rsa_chacha20_poly1305_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256 Cipher successes",
						},
						"tls13_aes_128_gcm_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13_AES_128_GCM_SHA256 cipher successes",
						},
						"tls13_aes_256_gcm_sha384_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13_AES_256_GCM_SHA384 cipher successes",
						},
						"tls13_chacha20_poly1305_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13_CHACHA20_POLY1305_SHA256 cipher successes",
						},
						"tls1_ecdhe_sm2_sms4_sm3_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_SM3 cipher successes",
						},
						"tls1_ecdhe_sm2_sms4_gcm_sm3_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_GCM_SM3 cipher successes",
						},
						"tls1_ecdhe_sm2_sms4_sha256_successes": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_SHA256 cipher successes",
						},
						"ssl3_rsa_des_192_cbc3_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_DES_192_CBC3_SHA Failures",
						},
						"ssl3_rsa_des_40_cbc_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_DES_40_CBC_SHA Failures",
						},
						"ssl3_rsa_des_64_cbc_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_DES_64_CBC_SHA Failures",
						},
						"ssl3_rsa_rc4_128_md5_failures": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_RC4_128_MD5 Failures",
						},
						"ssl3_rsa_rc4_128_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_RC4_128_SHA Failures",
						},
						"ssl3_rsa_rc4_40_md5_failures": {
							Type: schema.TypeInt, Optional: true, Description: "SSL3_RSA_RC4_40_MD5 Failures",
						},
						"tls1_dhe_rsa_aes_128_gcm_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_128_GCM_SHA256 Failures",
						},
						"tls1_dhe_rsa_aes_128_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_128_SHA Failures",
						},
						"tls1_dhe_rsa_aes_128_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_128_SHA256 Failures",
						},
						"tls1_dhe_rsa_aes_256_gcm_sha384_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_256_GCM_SHA384 Failures",
						},
						"tls1_dhe_rsa_aes_256_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_256_SHA Failures",
						},
						"tls1_dhe_rsa_aes_256_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_AES_256_SHA256 Failures",
						},
						"tls1_ecdhe_ecdsa_aes_128_gcm_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256 Failures",
						},
						"tls1_ecdhe_ecdsa_aes_128_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_SHA Failures",
						},
						"tls1_ecdhe_ecdsa_aes_128_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_128_SHA256 Failures",
						},
						"tls1_ecdhe_ecdsa_aes_256_sha384_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_SHA384 Failures",
						},
						"tls1_ecdhe_ecdsa_aes_256_gcm_sha384_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384 Failures",
						},
						"tls1_ecdhe_ecdsa_aes_256_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_AES_256_SHA Failures",
						},
						"tls1_ecdhe_rsa_aes_128_gcm_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_GCM_SHA256 Failures",
						},
						"tls1_ecdhe_rsa_aes_128_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_SHA Failures",
						},
						"tls1_ecdhe_rsa_aes_128_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_128_SHA256 Failures",
						},
						"tls1_ecdhe_rsa_aes_256_sha384_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_SHA384 Failures",
						},
						"tls1_ecdhe_rsa_aes_256_gcm_sha384_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_GCM_SHA384 Failures",
						},
						"tls1_ecdhe_rsa_aes_256_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_AES_256_SHA Failures",
						},
						"tls1_rsa_aes_128_gcm_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_128_GCM_SHA256 Failures",
						},
						"tls1_rsa_aes_128_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_128_SHA Failures",
						},
						"tls1_rsa_aes_128_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_128_SHA256 Failures",
						},
						"tls1_rsa_aes_256_gcm_sha384_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_256_GCM_SHA384 Failures",
						},
						"tls1_rsa_aes_256_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_256_SHA Failures",
						},
						"tls1_rsa_aes_256_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_AES_256_SHA256 Failures",
						},
						"tls1_rsa_export1024_rc4_56_md5_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_EXPORT1024_RC4_56_MD5 Failures",
						},
						"tls1_rsa_export1024_rc4_56_sha_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_RSA_EXPORT1024_RC4_56_SHA Failures",
						},
						"tls1_ecdhe_rsa_chacha20_poly1305_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256 Cipher failures",
						},
						"tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256 Cipher failures",
						},
						"tls1_dhe_rsa_chacha20_poly1305_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256 Cipher failures",
						},
						"tls13_aes_128_gcm_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13_AES_128_GCM_SHA256 cipher failures",
						},
						"tls13_aes_256_gcm_sha384_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13_AES_256_GCM_SHA384 cipher failures",
						},
						"tls13_chacha20_poly1305_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13_CHACHA20_POLY1305_SHA256 cipher failures",
						},
						"tls1_ecdhe_sm2_sms4_sm3_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_SM3 cipher failures",
						},
						"tls1_ecdhe_sm2_sms4_gcm_sm3_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_GCM_SM3 cipher failures",
						},
						"tls1_ecdhe_sm2_sms4_sha256_failures": {
							Type: schema.TypeInt, Optional: true, Description: "TLS1_ECDHE_SM2_WITH_SMS4_SHA256 cipher failures",
						},
						"kex_rsa_512_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful 512-bit RSA key exchanges",
						},
						"kex_rsa_1024_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful 1024-bit RSA key exchanges",
						},
						"kex_rsa_2048_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful 2048-bit RSA key exchanges",
						},
						"kex_rsa_4096_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful 4096-bit RSA key exchanges",
						},
						"kex_rsa_512_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed 512-bit RSA key exchanges",
						},
						"kex_rsa_1024_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed 1024-bit RSA key exchanges",
						},
						"kex_rsa_2048_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed 2048-bit RSA key exchanges",
						},
						"kex_rsa_4096_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed 4096-bit RSA key exchanges",
						},
						"kex_ecdhe_secp256r1_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful secp256r1 ECDHE key exchanges",
						},
						"kex_ecdhe_secp384r1_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful secp384r1 ECDHE key exchanges",
						},
						"kex_ecdhe_secp521r1_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful secp521r1 ECDHE key exchanges",
						},
						"kex_ecdhe_x25519_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful x25519 ECDHE key exchanges",
						},
						"kex_ecdhe_x448_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful x448 ECDHE key exchanges",
						},
						"kex_ecdhe_sm2_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful sm2p256v1 ECDHE key exchanges",
						},
						"kex_ecdhe_secp256r1_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed secp256r1 ECDHE key exchanges",
						},
						"kex_ecdhe_secp384r1_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed secp384r1 ECDHE key exchanges",
						},
						"kex_ecdhe_secp521r1_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed secp521r1 ECDHE key exchanges",
						},
						"kex_ecdhe_x25519_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed x25519 ECDHE key exchanges",
						},
						"kex_ecdhe_x448_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed x448 ECDHE key exchanges",
						},
						"kex_ecdhe_sm2_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed sm2p256v1 ECDHE key exchanges",
						},
						"kex_dhe_512_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful 512-bit DHE key exchanges",
						},
						"kex_dhe_1024_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful 1024-bit DHE key exchanges",
						},
						"kex_dhe_2048_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful 2048-bit DHE key exchanges",
						},
						"kex_dhe_512_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed 512-bit DHE key exchanges",
						},
						"kex_dhe_1024_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed 1024-bit DHE key exchanges",
						},
						"kex_dhe_2048_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed 2048-bit DHE key exchanges",
						},
						"ssl2_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful SSL2 connections",
						},
						"ssl3_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful SSL3 connections",
						},
						"tls10_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.0 connections",
						},
						"tls11_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.1 connections",
						},
						"tls12_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.2 connections",
						},
						"tls13_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.3 connections",
						},
						"ssl2_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SSL2 connections",
						},
						"ssl3_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SSL3 connections",
						},
						"tls10_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.0 connections",
						},
						"tls11_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.1 connections",
						},
						"tls12_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.2 connections",
						},
						"tls13_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.3 connections",
						},
						"sess_cache_new": {
							Type: schema.TypeInt, Optional: true, Description: "Session cache new entries",
						},
						"sess_cache_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Session cache hits",
						},
						"sess_cache_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Session cache misses",
						},
						"sess_cache_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Session cache timeouts",
						},
						"sess_cache_curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Session cache current connections",
						},
						"hs_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Total handshake failures",
						},
						"cert_vfy": {
							Type: schema.TypeInt, Optional: true, Description: "Sent certificate verify for authentication",
						},
						"hs_avg_time": {
							Type: schema.TypeInt, Optional: true, Description: "Average handshake time in milliseconds",
						},
						"sni_automap_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful SNI auto mappings",
						},
						"sni_automap_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SNI auto mappings",
						},
						"sni_automap_conn_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Conn closed before SNI auto mappings",
						},
						"sni_automap_max_active_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SNI auto map due to max active limit",
						},
						"sni_automap_missing_cert": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SNI auto map due to missing cert/key",
						},
						"sni_bypass_missing_cert": {
							Type: schema.TypeInt, Optional: true, Description: "SNI bypass event due to missing cert/key",
						},
						"sni_bypass_expired_cert": {
							Type: schema.TypeInt, Optional: true, Description: "SNI bypass event due to certificate expired",
						},
						"sni_bypass_explicit_list": {
							Type: schema.TypeInt, Optional: true, Description: "SNI bypass event due to matched explicit bypass list",
						},
						"renegotiation_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total renegotiations",
						},
						"renego_ssl2_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful SSL2 renegotiations",
						},
						"renego_ssl3_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful SSL3 renegotiations",
						},
						"renego_tls10_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.0 renegotiations",
						},
						"renego_tls11_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.1 renegotiations",
						},
						"renego_tls12_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.2 renegotiations",
						},
						"renego_tls13_successes": {
							Type: schema.TypeInt, Optional: true, Description: "Successful TLS1.3 renegotiations",
						},
						"renego_ssl2_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SSL2 renegotiations",
						},
						"renego_ssl3_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SSL3 renegotiations",
						},
						"renego_tls10_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.0 renegotiations",
						},
						"renego_tls11_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.1 renegotiations",
						},
						"renego_tls12_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.2 renegotiations",
						},
						"renego_tls13_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed TLS1.3 renegotiations",
						},
						"downgraded": {
							Type: schema.TypeInt, Optional: true, Description: "TLS version downgraded",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslCountersOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCountersOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCountersOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCountersOperOper := setObjectSlbSslCountersOperOper(res)
		d.Set("oper", SlbSslCountersOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCountersOperOper(ret edpt.DataSlbSslCountersOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vserver":                                             ret.DtSlbSslCountersOper.Oper.Vserver,
			"port":                                                ret.DtSlbSslCountersOper.Oper.Port,
			"cumulative_sessions":                                 ret.DtSlbSslCountersOper.Oper.CumulativeSessions,
			"ssl3_rsa_des_192_cbc3_sha_id":                        ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes192Cbc3ShaId,
			"ssl3_rsa_des_40_cbc_sha_id":                          ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes40CbcShaId,
			"ssl3_rsa_des_64_cbc_sha_id":                          ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes64CbcShaId,
			"ssl3_rsa_rc4_128_md5_id":                             ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc4128Md5Id,
			"ssl3_rsa_rc4_128_sha_id":                             ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc4128ShaId,
			"ssl3_rsa_rc4_40_md5_id":                              ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc440Md5Id,
			"tls1_dhe_rsa_aes_128_gcm_sha256_id":                  ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128GcmSha256Id,
			"tls1_dhe_rsa_aes_128_sha_id":                         ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128ShaId,
			"tls1_dhe_rsa_aes_128_sha256_id":                      ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128Sha256Id,
			"tls1_dhe_rsa_aes_256_gcm_sha384_id":                  ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256GcmSha384Id,
			"tls1_dhe_rsa_aes_256_sha_id":                         ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256ShaId,
			"tls1_dhe_rsa_aes_256_sha256_id":                      ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256Sha256Id,
			"tls1_ecdhe_ecdsa_aes_128_gcm_sha256_id":              ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128GcmSha256Id,
			"tls1_ecdhe_ecdsa_aes_128_sha_id":                     ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128ShaId,
			"tls1_ecdhe_ecdsa_aes_128_sha256_id":                  ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128Sha256Id,
			"tls1_ecdhe_ecdsa_aes_256_sha384_id":                  ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256Sha384Id,
			"tls1_ecdhe_ecdsa_aes_256_gcm_sha384_id":              ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256GcmSha384Id,
			"tls1_ecdhe_ecdsa_aes_256_sha_id":                     ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256ShaId,
			"tls1_ecdhe_rsa_aes_128_gcm_sha256_id":                ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128GcmSha256Id,
			"tls1_ecdhe_rsa_aes_128_sha_id":                       ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128ShaId,
			"tls1_ecdhe_rsa_aes_128_sha256_id":                    ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128Sha256Id,
			"tls1_ecdhe_rsa_aes_256_sha384_id":                    ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256Sha384Id,
			"tls1_ecdhe_rsa_aes_256_gcm_sha384_id":                ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256GcmSha384Id,
			"tls1_ecdhe_rsa_aes_256_sha_id":                       ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256ShaId,
			"tls1_rsa_aes_128_gcm_sha256_id":                      ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128GcmSha256Id,
			"tls1_rsa_aes_128_sha_id":                             ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128ShaId,
			"tls1_rsa_aes_128_sha256_id":                          ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128Sha256Id,
			"tls1_rsa_aes_256_gcm_sha384_id":                      ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256GcmSha384Id,
			"tls1_rsa_aes_256_sha_id":                             ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256ShaId,
			"tls1_rsa_aes_256_sha256_id":                          ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256Sha256Id,
			"tls1_rsa_export1024_rc4_56_md5_id":                   ret.DtSlbSslCountersOper.Oper.Tls1RsaExport1024Rc456Md5Id,
			"tls1_rsa_export1024_rc4_56_sha_id":                   ret.DtSlbSslCountersOper.Oper.Tls1RsaExport1024Rc456ShaId,
			"tls1_ecdhe_rsa_chacha20_poly1305_sha256_id":          ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaChacha20Poly1305Sha256Id,
			"tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_id":        ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaChacha20Poly1305Sha256Id,
			"tls1_dhe_rsa_chacha20_poly1305_sha256_id":            ret.DtSlbSslCountersOper.Oper.Tls1DheRsaChacha20Poly1305Sha256Id,
			"tls13_aes_128_gcm_sha256_id":                         ret.DtSlbSslCountersOper.Oper.Tls13Aes128GcmSha256Id,
			"tls13_aes_256_gcm_sha384_id":                         ret.DtSlbSslCountersOper.Oper.Tls13Aes256GcmSha384Id,
			"tls13_chacha20_poly1305_sha256_id":                   ret.DtSlbSslCountersOper.Oper.Tls13Chacha20Poly1305Sha256Id,
			"tls1_ecdhe_sm2_sms4_sm3_id":                          ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4Sm3Id,
			"tls1_ecdhe_sm2_sms4_gcm_sm3_id":                      ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4GcmSm3Id,
			"tls1_ecdhe_sm2_sms4_sha256_id":                       ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4Sha256Id,
			"ssl3_rsa_des_192_cbc3_sha_successes":                 ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes192Cbc3ShaSuccesses,
			"ssl3_rsa_des_40_cbc_sha_successes":                   ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes40CbcShaSuccesses,
			"ssl3_rsa_des_64_cbc_sha_successes":                   ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes64CbcShaSuccesses,
			"ssl3_rsa_rc4_128_md5_successes":                      ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc4128Md5Successes,
			"ssl3_rsa_rc4_128_sha_successes":                      ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc4128ShaSuccesses,
			"ssl3_rsa_rc4_40_md5_successes":                       ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc440Md5Successes,
			"tls1_dhe_rsa_aes_128_gcm_sha256_successes":           ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128GcmSha256Successes,
			"tls1_dhe_rsa_aes_128_sha_successes":                  ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128ShaSuccesses,
			"tls1_dhe_rsa_aes_128_sha256_successes":               ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128Sha256Successes,
			"tls1_dhe_rsa_aes_256_gcm_sha384_successes":           ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256GcmSha384Successes,
			"tls1_dhe_rsa_aes_256_sha_successes":                  ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256ShaSuccesses,
			"tls1_dhe_rsa_aes_256_sha256_successes":               ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256Sha256Successes,
			"tls1_ecdhe_ecdsa_aes_128_gcm_sha256_successes":       ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128GcmSha256Successes,
			"tls1_ecdhe_ecdsa_aes_128_sha_successes":              ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128ShaSuccesses,
			"tls1_ecdhe_ecdsa_aes_128_sha256_successes":           ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128Sha256Successes,
			"tls1_ecdhe_ecdsa_aes_256_sha384_successes":           ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256Sha384Successes,
			"tls1_ecdhe_ecdsa_aes_256_gcm_sha384_successes":       ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256GcmSha384Successes,
			"tls1_ecdhe_ecdsa_aes_256_sha_successes":              ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256ShaSuccesses,
			"tls1_ecdhe_rsa_aes_128_gcm_sha256_successes":         ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128GcmSha256Successes,
			"tls1_ecdhe_rsa_aes_128_sha_successes":                ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128ShaSuccesses,
			"tls1_ecdhe_rsa_aes_128_sha256_successes":             ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128Sha256Successes,
			"tls1_ecdhe_rsa_aes_256_sha384_successes":             ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256Sha384Successes,
			"tls1_ecdhe_rsa_aes_256_gcm_sha384_successes":         ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256GcmSha384Successes,
			"tls1_ecdhe_rsa_aes_256_sha_successes":                ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256ShaSuccesses,
			"tls1_rsa_aes_128_gcm_sha256_successes":               ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128GcmSha256Successes,
			"tls1_rsa_aes_128_sha_successes":                      ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128ShaSuccesses,
			"tls1_rsa_aes_128_sha256_successes":                   ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128Sha256Successes,
			"tls1_rsa_aes_256_gcm_sha384_successes":               ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256GcmSha384Successes,
			"tls1_rsa_aes_256_sha_successes":                      ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256ShaSuccesses,
			"tls1_rsa_aes_256_sha256_successes":                   ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256Sha256Successes,
			"tls1_rsa_export1024_rc4_56_md5_successes":            ret.DtSlbSslCountersOper.Oper.Tls1RsaExport1024Rc456Md5Successes,
			"tls1_rsa_export1024_rc4_56_sha_successes":            ret.DtSlbSslCountersOper.Oper.Tls1RsaExport1024Rc456ShaSuccesses,
			"tls1_ecdhe_rsa_chacha20_poly1305_sha256_successes":   ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaChacha20Poly1305Sha256Successes,
			"tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_successes": ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaChacha20Poly1305Sha256Successes,
			"tls1_dhe_rsa_chacha20_poly1305_sha256_successes":     ret.DtSlbSslCountersOper.Oper.Tls1DheRsaChacha20Poly1305Sha256Successes,
			"tls13_aes_128_gcm_sha256_successes":                  ret.DtSlbSslCountersOper.Oper.Tls13Aes128GcmSha256Successes,
			"tls13_aes_256_gcm_sha384_successes":                  ret.DtSlbSslCountersOper.Oper.Tls13Aes256GcmSha384Successes,
			"tls13_chacha20_poly1305_sha256_successes":            ret.DtSlbSslCountersOper.Oper.Tls13Chacha20Poly1305Sha256Successes,
			"tls1_ecdhe_sm2_sms4_sm3_successes":                   ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4Sm3Successes,
			"tls1_ecdhe_sm2_sms4_gcm_sm3_successes":               ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4GcmSm3Successes,
			"tls1_ecdhe_sm2_sms4_sha256_successes":                ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4Sha256Successes,
			"ssl3_rsa_des_192_cbc3_sha_failures":                  ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes192Cbc3ShaFailures,
			"ssl3_rsa_des_40_cbc_sha_failures":                    ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes40CbcShaFailures,
			"ssl3_rsa_des_64_cbc_sha_failures":                    ret.DtSlbSslCountersOper.Oper.Ssl3RsaDes64CbcShaFailures,
			"ssl3_rsa_rc4_128_md5_failures":                       ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc4128Md5Failures,
			"ssl3_rsa_rc4_128_sha_failures":                       ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc4128ShaFailures,
			"ssl3_rsa_rc4_40_md5_failures":                        ret.DtSlbSslCountersOper.Oper.Ssl3RsaRc440Md5Failures,
			"tls1_dhe_rsa_aes_128_gcm_sha256_failures":            ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128GcmSha256Failures,
			"tls1_dhe_rsa_aes_128_sha_failures":                   ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128ShaFailures,
			"tls1_dhe_rsa_aes_128_sha256_failures":                ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes128Sha256Failures,
			"tls1_dhe_rsa_aes_256_gcm_sha384_failures":            ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256GcmSha384Failures,
			"tls1_dhe_rsa_aes_256_sha_failures":                   ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256ShaFailures,
			"tls1_dhe_rsa_aes_256_sha256_failures":                ret.DtSlbSslCountersOper.Oper.Tls1DheRsaAes256Sha256Failures,
			"tls1_ecdhe_ecdsa_aes_128_gcm_sha256_failures":        ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128GcmSha256Failures,
			"tls1_ecdhe_ecdsa_aes_128_sha_failures":               ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128ShaFailures,
			"tls1_ecdhe_ecdsa_aes_128_sha256_failures":            ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes128Sha256Failures,
			"tls1_ecdhe_ecdsa_aes_256_sha384_failures":            ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256Sha384Failures,
			"tls1_ecdhe_ecdsa_aes_256_gcm_sha384_failures":        ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256GcmSha384Failures,
			"tls1_ecdhe_ecdsa_aes_256_sha_failures":               ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaAes256ShaFailures,
			"tls1_ecdhe_rsa_aes_128_gcm_sha256_failures":          ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128GcmSha256Failures,
			"tls1_ecdhe_rsa_aes_128_sha_failures":                 ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128ShaFailures,
			"tls1_ecdhe_rsa_aes_128_sha256_failures":              ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes128Sha256Failures,
			"tls1_ecdhe_rsa_aes_256_sha384_failures":              ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256Sha384Failures,
			"tls1_ecdhe_rsa_aes_256_gcm_sha384_failures":          ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256GcmSha384Failures,
			"tls1_ecdhe_rsa_aes_256_sha_failures":                 ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaAes256ShaFailures,
			"tls1_rsa_aes_128_gcm_sha256_failures":                ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128GcmSha256Failures,
			"tls1_rsa_aes_128_sha_failures":                       ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128ShaFailures,
			"tls1_rsa_aes_128_sha256_failures":                    ret.DtSlbSslCountersOper.Oper.Tls1RsaAes128Sha256Failures,
			"tls1_rsa_aes_256_gcm_sha384_failures":                ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256GcmSha384Failures,
			"tls1_rsa_aes_256_sha_failures":                       ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256ShaFailures,
			"tls1_rsa_aes_256_sha256_failures":                    ret.DtSlbSslCountersOper.Oper.Tls1RsaAes256Sha256Failures,
			"tls1_rsa_export1024_rc4_56_md5_failures":             ret.DtSlbSslCountersOper.Oper.Tls1RsaExport1024Rc456Md5Failures,
			"tls1_rsa_export1024_rc4_56_sha_failures":             ret.DtSlbSslCountersOper.Oper.Tls1RsaExport1024Rc456ShaFailures,
			"tls1_ecdhe_rsa_chacha20_poly1305_sha256_failures":    ret.DtSlbSslCountersOper.Oper.Tls1EcdheRsaChacha20Poly1305Sha256Failures,
			"tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_failures":  ret.DtSlbSslCountersOper.Oper.Tls1EcdheEcdsaChacha20Poly1305Sha256Failures,
			"tls1_dhe_rsa_chacha20_poly1305_sha256_failures":      ret.DtSlbSslCountersOper.Oper.Tls1DheRsaChacha20Poly1305Sha256Failures,
			"tls13_aes_128_gcm_sha256_failures":                   ret.DtSlbSslCountersOper.Oper.Tls13Aes128GcmSha256Failures,
			"tls13_aes_256_gcm_sha384_failures":                   ret.DtSlbSslCountersOper.Oper.Tls13Aes256GcmSha384Failures,
			"tls13_chacha20_poly1305_sha256_failures":             ret.DtSlbSslCountersOper.Oper.Tls13Chacha20Poly1305Sha256Failures,
			"tls1_ecdhe_sm2_sms4_sm3_failures":                    ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4Sm3Failures,
			"tls1_ecdhe_sm2_sms4_gcm_sm3_failures":                ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4GcmSm3Failures,
			"tls1_ecdhe_sm2_sms4_sha256_failures":                 ret.DtSlbSslCountersOper.Oper.Tls1EcdheSm2Sms4Sha256Failures,
			"kex_rsa_512_successes":                               ret.DtSlbSslCountersOper.Oper.KexRsa512Successes,
			"kex_rsa_1024_successes":                              ret.DtSlbSslCountersOper.Oper.KexRsa1024Successes,
			"kex_rsa_2048_successes":                              ret.DtSlbSslCountersOper.Oper.KexRsa2048Successes,
			"kex_rsa_4096_successes":                              ret.DtSlbSslCountersOper.Oper.KexRsa4096Successes,
			"kex_rsa_512_failures":                                ret.DtSlbSslCountersOper.Oper.KexRsa512Failures,
			"kex_rsa_1024_failures":                               ret.DtSlbSslCountersOper.Oper.KexRsa1024Failures,
			"kex_rsa_2048_failures":                               ret.DtSlbSslCountersOper.Oper.KexRsa2048Failures,
			"kex_rsa_4096_failures":                               ret.DtSlbSslCountersOper.Oper.KexRsa4096Failures,
			"kex_ecdhe_secp256r1_successes":                       ret.DtSlbSslCountersOper.Oper.KexEcdheSecp256r1Successes,
			"kex_ecdhe_secp384r1_successes":                       ret.DtSlbSslCountersOper.Oper.KexEcdheSecp384r1Successes,
			"kex_ecdhe_secp521r1_successes":                       ret.DtSlbSslCountersOper.Oper.KexEcdheSecp521r1Successes,
			"kex_ecdhe_x25519_successes":                          ret.DtSlbSslCountersOper.Oper.KexEcdheX25519Successes,
			"kex_ecdhe_x448_successes":                            ret.DtSlbSslCountersOper.Oper.KexEcdheX448Successes,
			"kex_ecdhe_sm2_successes":                             ret.DtSlbSslCountersOper.Oper.KexEcdheSm2Successes,
			"kex_ecdhe_secp256r1_failures":                        ret.DtSlbSslCountersOper.Oper.KexEcdheSecp256r1Failures,
			"kex_ecdhe_secp384r1_failures":                        ret.DtSlbSslCountersOper.Oper.KexEcdheSecp384r1Failures,
			"kex_ecdhe_secp521r1_failures":                        ret.DtSlbSslCountersOper.Oper.KexEcdheSecp521r1Failures,
			"kex_ecdhe_x25519_failures":                           ret.DtSlbSslCountersOper.Oper.KexEcdheX25519Failures,
			"kex_ecdhe_x448_failures":                             ret.DtSlbSslCountersOper.Oper.KexEcdheX448Failures,
			"kex_ecdhe_sm2_failures":                              ret.DtSlbSslCountersOper.Oper.KexEcdheSm2Failures,
			"kex_dhe_512_successes":                               ret.DtSlbSslCountersOper.Oper.KexDhe512Successes,
			"kex_dhe_1024_successes":                              ret.DtSlbSslCountersOper.Oper.KexDhe1024Successes,
			"kex_dhe_2048_successes":                              ret.DtSlbSslCountersOper.Oper.KexDhe2048Successes,
			"kex_dhe_512_failures":                                ret.DtSlbSslCountersOper.Oper.KexDhe512Failures,
			"kex_dhe_1024_failures":                               ret.DtSlbSslCountersOper.Oper.KexDhe1024Failures,
			"kex_dhe_2048_failures":                               ret.DtSlbSslCountersOper.Oper.KexDhe2048Failures,
			"ssl2_successes":                                      ret.DtSlbSslCountersOper.Oper.Ssl2Successes,
			"ssl3_successes":                                      ret.DtSlbSslCountersOper.Oper.Ssl3Successes,
			"tls10_successes":                                     ret.DtSlbSslCountersOper.Oper.Tls10Successes,
			"tls11_successes":                                     ret.DtSlbSslCountersOper.Oper.Tls11Successes,
			"tls12_successes":                                     ret.DtSlbSslCountersOper.Oper.Tls12Successes,
			"tls13_successes":                                     ret.DtSlbSslCountersOper.Oper.Tls13Successes,
			"ssl2_failures":                                       ret.DtSlbSslCountersOper.Oper.Ssl2Failures,
			"ssl3_failures":                                       ret.DtSlbSslCountersOper.Oper.Ssl3Failures,
			"tls10_failures":                                      ret.DtSlbSslCountersOper.Oper.Tls10Failures,
			"tls11_failures":                                      ret.DtSlbSslCountersOper.Oper.Tls11Failures,
			"tls12_failures":                                      ret.DtSlbSslCountersOper.Oper.Tls12Failures,
			"tls13_failures":                                      ret.DtSlbSslCountersOper.Oper.Tls13Failures,
			"sess_cache_new":                                      ret.DtSlbSslCountersOper.Oper.SessCacheNew,
			"sess_cache_hit":                                      ret.DtSlbSslCountersOper.Oper.SessCacheHit,
			"sess_cache_miss":                                     ret.DtSlbSslCountersOper.Oper.SessCacheMiss,
			"sess_cache_timeout":                                  ret.DtSlbSslCountersOper.Oper.SessCacheTimeout,
			"sess_cache_curr_conn":                                ret.DtSlbSslCountersOper.Oper.SessCacheCurrConn,
			"hs_failures":                                         ret.DtSlbSslCountersOper.Oper.HsFailures,
			"cert_vfy":                                            ret.DtSlbSslCountersOper.Oper.CertVfy,
			"hs_avg_time":                                         ret.DtSlbSslCountersOper.Oper.HsAvgTime,
			"sni_automap_successes":                               ret.DtSlbSslCountersOper.Oper.SniAutomapSuccesses,
			"sni_automap_failures":                                ret.DtSlbSslCountersOper.Oper.SniAutomapFailures,
			"sni_automap_conn_closed":                             ret.DtSlbSslCountersOper.Oper.SniAutomapConnClosed,
			"sni_automap_max_active_conn":                         ret.DtSlbSslCountersOper.Oper.SniAutomapMaxActiveConn,
			"sni_automap_missing_cert":                            ret.DtSlbSslCountersOper.Oper.SniAutomapMissingCert,
			"sni_bypass_missing_cert":                             ret.DtSlbSslCountersOper.Oper.SniBypassMissingCert,
			"sni_bypass_expired_cert":                             ret.DtSlbSslCountersOper.Oper.SniBypassExpiredCert,
			"sni_bypass_explicit_list":                            ret.DtSlbSslCountersOper.Oper.SniBypassExplicitList,
			"renegotiation_total":                                 ret.DtSlbSslCountersOper.Oper.RenegotiationTotal,
			"renego_ssl2_successes":                               ret.DtSlbSslCountersOper.Oper.RenegoSsl2Successes,
			"renego_ssl3_successes":                               ret.DtSlbSslCountersOper.Oper.RenegoSsl3Successes,
			"renego_tls10_successes":                              ret.DtSlbSslCountersOper.Oper.RenegoTls10Successes,
			"renego_tls11_successes":                              ret.DtSlbSslCountersOper.Oper.RenegoTls11Successes,
			"renego_tls12_successes":                              ret.DtSlbSslCountersOper.Oper.RenegoTls12Successes,
			"renego_tls13_successes":                              ret.DtSlbSslCountersOper.Oper.RenegoTls13Successes,
			"renego_ssl2_failures":                                ret.DtSlbSslCountersOper.Oper.RenegoSsl2Failures,
			"renego_ssl3_failures":                                ret.DtSlbSslCountersOper.Oper.RenegoSsl3Failures,
			"renego_tls10_failures":                               ret.DtSlbSslCountersOper.Oper.RenegoTls10Failures,
			"renego_tls11_failures":                               ret.DtSlbSslCountersOper.Oper.RenegoTls11Failures,
			"renego_tls12_failures":                               ret.DtSlbSslCountersOper.Oper.RenegoTls12Failures,
			"renego_tls13_failures":                               ret.DtSlbSslCountersOper.Oper.RenegoTls13Failures,
			"downgraded":                                          ret.DtSlbSslCountersOper.Oper.Downgraded,
		},
	}
}

func getObjectSlbSslCountersOperOper(d []interface{}) edpt.SlbSslCountersOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCountersOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Vserver = in["vserver"].(string)
		ret.Port = in["port"].(int)
		ret.CumulativeSessions = in["cumulative_sessions"].(int)
		ret.Ssl3RsaDes192Cbc3ShaId = in["ssl3_rsa_des_192_cbc3_sha_id"].(string)
		ret.Ssl3RsaDes40CbcShaId = in["ssl3_rsa_des_40_cbc_sha_id"].(string)
		ret.Ssl3RsaDes64CbcShaId = in["ssl3_rsa_des_64_cbc_sha_id"].(string)
		ret.Ssl3RsaRc4128Md5Id = in["ssl3_rsa_rc4_128_md5_id"].(string)
		ret.Ssl3RsaRc4128ShaId = in["ssl3_rsa_rc4_128_sha_id"].(string)
		ret.Ssl3RsaRc440Md5Id = in["ssl3_rsa_rc4_40_md5_id"].(string)
		ret.Tls1DheRsaAes128GcmSha256Id = in["tls1_dhe_rsa_aes_128_gcm_sha256_id"].(string)
		ret.Tls1DheRsaAes128ShaId = in["tls1_dhe_rsa_aes_128_sha_id"].(string)
		ret.Tls1DheRsaAes128Sha256Id = in["tls1_dhe_rsa_aes_128_sha256_id"].(string)
		ret.Tls1DheRsaAes256GcmSha384Id = in["tls1_dhe_rsa_aes_256_gcm_sha384_id"].(string)
		ret.Tls1DheRsaAes256ShaId = in["tls1_dhe_rsa_aes_256_sha_id"].(string)
		ret.Tls1DheRsaAes256Sha256Id = in["tls1_dhe_rsa_aes_256_sha256_id"].(string)
		ret.Tls1EcdheEcdsaAes128GcmSha256Id = in["tls1_ecdhe_ecdsa_aes_128_gcm_sha256_id"].(string)
		ret.Tls1EcdheEcdsaAes128ShaId = in["tls1_ecdhe_ecdsa_aes_128_sha_id"].(string)
		ret.Tls1EcdheEcdsaAes128Sha256Id = in["tls1_ecdhe_ecdsa_aes_128_sha256_id"].(string)
		ret.Tls1EcdheEcdsaAes256Sha384Id = in["tls1_ecdhe_ecdsa_aes_256_sha384_id"].(string)
		ret.Tls1EcdheEcdsaAes256GcmSha384Id = in["tls1_ecdhe_ecdsa_aes_256_gcm_sha384_id"].(string)
		ret.Tls1EcdheEcdsaAes256ShaId = in["tls1_ecdhe_ecdsa_aes_256_sha_id"].(string)
		ret.Tls1EcdheRsaAes128GcmSha256Id = in["tls1_ecdhe_rsa_aes_128_gcm_sha256_id"].(string)
		ret.Tls1EcdheRsaAes128ShaId = in["tls1_ecdhe_rsa_aes_128_sha_id"].(string)
		ret.Tls1EcdheRsaAes128Sha256Id = in["tls1_ecdhe_rsa_aes_128_sha256_id"].(string)
		ret.Tls1EcdheRsaAes256Sha384Id = in["tls1_ecdhe_rsa_aes_256_sha384_id"].(string)
		ret.Tls1EcdheRsaAes256GcmSha384Id = in["tls1_ecdhe_rsa_aes_256_gcm_sha384_id"].(string)
		ret.Tls1EcdheRsaAes256ShaId = in["tls1_ecdhe_rsa_aes_256_sha_id"].(string)
		ret.Tls1RsaAes128GcmSha256Id = in["tls1_rsa_aes_128_gcm_sha256_id"].(string)
		ret.Tls1RsaAes128ShaId = in["tls1_rsa_aes_128_sha_id"].(string)
		ret.Tls1RsaAes128Sha256Id = in["tls1_rsa_aes_128_sha256_id"].(string)
		ret.Tls1RsaAes256GcmSha384Id = in["tls1_rsa_aes_256_gcm_sha384_id"].(string)
		ret.Tls1RsaAes256ShaId = in["tls1_rsa_aes_256_sha_id"].(string)
		ret.Tls1RsaAes256Sha256Id = in["tls1_rsa_aes_256_sha256_id"].(string)
		ret.Tls1RsaExport1024Rc456Md5Id = in["tls1_rsa_export1024_rc4_56_md5_id"].(string)
		ret.Tls1RsaExport1024Rc456ShaId = in["tls1_rsa_export1024_rc4_56_sha_id"].(string)
		ret.Tls1EcdheRsaChacha20Poly1305Sha256Id = in["tls1_ecdhe_rsa_chacha20_poly1305_sha256_id"].(string)
		ret.Tls1EcdheEcdsaChacha20Poly1305Sha256Id = in["tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_id"].(string)
		ret.Tls1DheRsaChacha20Poly1305Sha256Id = in["tls1_dhe_rsa_chacha20_poly1305_sha256_id"].(string)
		ret.Tls13Aes128GcmSha256Id = in["tls13_aes_128_gcm_sha256_id"].(string)
		ret.Tls13Aes256GcmSha384Id = in["tls13_aes_256_gcm_sha384_id"].(string)
		ret.Tls13Chacha20Poly1305Sha256Id = in["tls13_chacha20_poly1305_sha256_id"].(string)
		ret.Tls1EcdheSm2Sms4Sm3Id = in["tls1_ecdhe_sm2_sms4_sm3_id"].(string)
		ret.Tls1EcdheSm2Sms4GcmSm3Id = in["tls1_ecdhe_sm2_sms4_gcm_sm3_id"].(string)
		ret.Tls1EcdheSm2Sms4Sha256Id = in["tls1_ecdhe_sm2_sms4_sha256_id"].(string)
		ret.Ssl3RsaDes192Cbc3ShaSuccesses = in["ssl3_rsa_des_192_cbc3_sha_successes"].(int)
		ret.Ssl3RsaDes40CbcShaSuccesses = in["ssl3_rsa_des_40_cbc_sha_successes"].(int)
		ret.Ssl3RsaDes64CbcShaSuccesses = in["ssl3_rsa_des_64_cbc_sha_successes"].(int)
		ret.Ssl3RsaRc4128Md5Successes = in["ssl3_rsa_rc4_128_md5_successes"].(int)
		ret.Ssl3RsaRc4128ShaSuccesses = in["ssl3_rsa_rc4_128_sha_successes"].(int)
		ret.Ssl3RsaRc440Md5Successes = in["ssl3_rsa_rc4_40_md5_successes"].(int)
		ret.Tls1DheRsaAes128GcmSha256Successes = in["tls1_dhe_rsa_aes_128_gcm_sha256_successes"].(int)
		ret.Tls1DheRsaAes128ShaSuccesses = in["tls1_dhe_rsa_aes_128_sha_successes"].(int)
		ret.Tls1DheRsaAes128Sha256Successes = in["tls1_dhe_rsa_aes_128_sha256_successes"].(int)
		ret.Tls1DheRsaAes256GcmSha384Successes = in["tls1_dhe_rsa_aes_256_gcm_sha384_successes"].(int)
		ret.Tls1DheRsaAes256ShaSuccesses = in["tls1_dhe_rsa_aes_256_sha_successes"].(int)
		ret.Tls1DheRsaAes256Sha256Successes = in["tls1_dhe_rsa_aes_256_sha256_successes"].(int)
		ret.Tls1EcdheEcdsaAes128GcmSha256Successes = in["tls1_ecdhe_ecdsa_aes_128_gcm_sha256_successes"].(int)
		ret.Tls1EcdheEcdsaAes128ShaSuccesses = in["tls1_ecdhe_ecdsa_aes_128_sha_successes"].(int)
		ret.Tls1EcdheEcdsaAes128Sha256Successes = in["tls1_ecdhe_ecdsa_aes_128_sha256_successes"].(int)
		ret.Tls1EcdheEcdsaAes256Sha384Successes = in["tls1_ecdhe_ecdsa_aes_256_sha384_successes"].(int)
		ret.Tls1EcdheEcdsaAes256GcmSha384Successes = in["tls1_ecdhe_ecdsa_aes_256_gcm_sha384_successes"].(int)
		ret.Tls1EcdheEcdsaAes256ShaSuccesses = in["tls1_ecdhe_ecdsa_aes_256_sha_successes"].(int)
		ret.Tls1EcdheRsaAes128GcmSha256Successes = in["tls1_ecdhe_rsa_aes_128_gcm_sha256_successes"].(int)
		ret.Tls1EcdheRsaAes128ShaSuccesses = in["tls1_ecdhe_rsa_aes_128_sha_successes"].(int)
		ret.Tls1EcdheRsaAes128Sha256Successes = in["tls1_ecdhe_rsa_aes_128_sha256_successes"].(int)
		ret.Tls1EcdheRsaAes256Sha384Successes = in["tls1_ecdhe_rsa_aes_256_sha384_successes"].(int)
		ret.Tls1EcdheRsaAes256GcmSha384Successes = in["tls1_ecdhe_rsa_aes_256_gcm_sha384_successes"].(int)
		ret.Tls1EcdheRsaAes256ShaSuccesses = in["tls1_ecdhe_rsa_aes_256_sha_successes"].(int)
		ret.Tls1RsaAes128GcmSha256Successes = in["tls1_rsa_aes_128_gcm_sha256_successes"].(int)
		ret.Tls1RsaAes128ShaSuccesses = in["tls1_rsa_aes_128_sha_successes"].(int)
		ret.Tls1RsaAes128Sha256Successes = in["tls1_rsa_aes_128_sha256_successes"].(int)
		ret.Tls1RsaAes256GcmSha384Successes = in["tls1_rsa_aes_256_gcm_sha384_successes"].(int)
		ret.Tls1RsaAes256ShaSuccesses = in["tls1_rsa_aes_256_sha_successes"].(int)
		ret.Tls1RsaAes256Sha256Successes = in["tls1_rsa_aes_256_sha256_successes"].(int)
		ret.Tls1RsaExport1024Rc456Md5Successes = in["tls1_rsa_export1024_rc4_56_md5_successes"].(int)
		ret.Tls1RsaExport1024Rc456ShaSuccesses = in["tls1_rsa_export1024_rc4_56_sha_successes"].(int)
		ret.Tls1EcdheRsaChacha20Poly1305Sha256Successes = in["tls1_ecdhe_rsa_chacha20_poly1305_sha256_successes"].(int)
		ret.Tls1EcdheEcdsaChacha20Poly1305Sha256Successes = in["tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_successes"].(int)
		ret.Tls1DheRsaChacha20Poly1305Sha256Successes = in["tls1_dhe_rsa_chacha20_poly1305_sha256_successes"].(int)
		ret.Tls13Aes128GcmSha256Successes = in["tls13_aes_128_gcm_sha256_successes"].(int)
		ret.Tls13Aes256GcmSha384Successes = in["tls13_aes_256_gcm_sha384_successes"].(int)
		ret.Tls13Chacha20Poly1305Sha256Successes = in["tls13_chacha20_poly1305_sha256_successes"].(int)
		ret.Tls1EcdheSm2Sms4Sm3Successes = in["tls1_ecdhe_sm2_sms4_sm3_successes"].(int)
		ret.Tls1EcdheSm2Sms4GcmSm3Successes = in["tls1_ecdhe_sm2_sms4_gcm_sm3_successes"].(int)
		ret.Tls1EcdheSm2Sms4Sha256Successes = in["tls1_ecdhe_sm2_sms4_sha256_successes"].(int)
		ret.Ssl3RsaDes192Cbc3ShaFailures = in["ssl3_rsa_des_192_cbc3_sha_failures"].(int)
		ret.Ssl3RsaDes40CbcShaFailures = in["ssl3_rsa_des_40_cbc_sha_failures"].(int)
		ret.Ssl3RsaDes64CbcShaFailures = in["ssl3_rsa_des_64_cbc_sha_failures"].(int)
		ret.Ssl3RsaRc4128Md5Failures = in["ssl3_rsa_rc4_128_md5_failures"].(int)
		ret.Ssl3RsaRc4128ShaFailures = in["ssl3_rsa_rc4_128_sha_failures"].(int)
		ret.Ssl3RsaRc440Md5Failures = in["ssl3_rsa_rc4_40_md5_failures"].(int)
		ret.Tls1DheRsaAes128GcmSha256Failures = in["tls1_dhe_rsa_aes_128_gcm_sha256_failures"].(int)
		ret.Tls1DheRsaAes128ShaFailures = in["tls1_dhe_rsa_aes_128_sha_failures"].(int)
		ret.Tls1DheRsaAes128Sha256Failures = in["tls1_dhe_rsa_aes_128_sha256_failures"].(int)
		ret.Tls1DheRsaAes256GcmSha384Failures = in["tls1_dhe_rsa_aes_256_gcm_sha384_failures"].(int)
		ret.Tls1DheRsaAes256ShaFailures = in["tls1_dhe_rsa_aes_256_sha_failures"].(int)
		ret.Tls1DheRsaAes256Sha256Failures = in["tls1_dhe_rsa_aes_256_sha256_failures"].(int)
		ret.Tls1EcdheEcdsaAes128GcmSha256Failures = in["tls1_ecdhe_ecdsa_aes_128_gcm_sha256_failures"].(int)
		ret.Tls1EcdheEcdsaAes128ShaFailures = in["tls1_ecdhe_ecdsa_aes_128_sha_failures"].(int)
		ret.Tls1EcdheEcdsaAes128Sha256Failures = in["tls1_ecdhe_ecdsa_aes_128_sha256_failures"].(int)
		ret.Tls1EcdheEcdsaAes256Sha384Failures = in["tls1_ecdhe_ecdsa_aes_256_sha384_failures"].(int)
		ret.Tls1EcdheEcdsaAes256GcmSha384Failures = in["tls1_ecdhe_ecdsa_aes_256_gcm_sha384_failures"].(int)
		ret.Tls1EcdheEcdsaAes256ShaFailures = in["tls1_ecdhe_ecdsa_aes_256_sha_failures"].(int)
		ret.Tls1EcdheRsaAes128GcmSha256Failures = in["tls1_ecdhe_rsa_aes_128_gcm_sha256_failures"].(int)
		ret.Tls1EcdheRsaAes128ShaFailures = in["tls1_ecdhe_rsa_aes_128_sha_failures"].(int)
		ret.Tls1EcdheRsaAes128Sha256Failures = in["tls1_ecdhe_rsa_aes_128_sha256_failures"].(int)
		ret.Tls1EcdheRsaAes256Sha384Failures = in["tls1_ecdhe_rsa_aes_256_sha384_failures"].(int)
		ret.Tls1EcdheRsaAes256GcmSha384Failures = in["tls1_ecdhe_rsa_aes_256_gcm_sha384_failures"].(int)
		ret.Tls1EcdheRsaAes256ShaFailures = in["tls1_ecdhe_rsa_aes_256_sha_failures"].(int)
		ret.Tls1RsaAes128GcmSha256Failures = in["tls1_rsa_aes_128_gcm_sha256_failures"].(int)
		ret.Tls1RsaAes128ShaFailures = in["tls1_rsa_aes_128_sha_failures"].(int)
		ret.Tls1RsaAes128Sha256Failures = in["tls1_rsa_aes_128_sha256_failures"].(int)
		ret.Tls1RsaAes256GcmSha384Failures = in["tls1_rsa_aes_256_gcm_sha384_failures"].(int)
		ret.Tls1RsaAes256ShaFailures = in["tls1_rsa_aes_256_sha_failures"].(int)
		ret.Tls1RsaAes256Sha256Failures = in["tls1_rsa_aes_256_sha256_failures"].(int)
		ret.Tls1RsaExport1024Rc456Md5Failures = in["tls1_rsa_export1024_rc4_56_md5_failures"].(int)
		ret.Tls1RsaExport1024Rc456ShaFailures = in["tls1_rsa_export1024_rc4_56_sha_failures"].(int)
		ret.Tls1EcdheRsaChacha20Poly1305Sha256Failures = in["tls1_ecdhe_rsa_chacha20_poly1305_sha256_failures"].(int)
		ret.Tls1EcdheEcdsaChacha20Poly1305Sha256Failures = in["tls1_ecdhe_ecdsa_chacha20_poly1305_sha256_failures"].(int)
		ret.Tls1DheRsaChacha20Poly1305Sha256Failures = in["tls1_dhe_rsa_chacha20_poly1305_sha256_failures"].(int)
		ret.Tls13Aes128GcmSha256Failures = in["tls13_aes_128_gcm_sha256_failures"].(int)
		ret.Tls13Aes256GcmSha384Failures = in["tls13_aes_256_gcm_sha384_failures"].(int)
		ret.Tls13Chacha20Poly1305Sha256Failures = in["tls13_chacha20_poly1305_sha256_failures"].(int)
		ret.Tls1EcdheSm2Sms4Sm3Failures = in["tls1_ecdhe_sm2_sms4_sm3_failures"].(int)
		ret.Tls1EcdheSm2Sms4GcmSm3Failures = in["tls1_ecdhe_sm2_sms4_gcm_sm3_failures"].(int)
		ret.Tls1EcdheSm2Sms4Sha256Failures = in["tls1_ecdhe_sm2_sms4_sha256_failures"].(int)
		ret.KexRsa512Successes = in["kex_rsa_512_successes"].(int)
		ret.KexRsa1024Successes = in["kex_rsa_1024_successes"].(int)
		ret.KexRsa2048Successes = in["kex_rsa_2048_successes"].(int)
		ret.KexRsa4096Successes = in["kex_rsa_4096_successes"].(int)
		ret.KexRsa512Failures = in["kex_rsa_512_failures"].(int)
		ret.KexRsa1024Failures = in["kex_rsa_1024_failures"].(int)
		ret.KexRsa2048Failures = in["kex_rsa_2048_failures"].(int)
		ret.KexRsa4096Failures = in["kex_rsa_4096_failures"].(int)
		ret.KexEcdheSecp256r1Successes = in["kex_ecdhe_secp256r1_successes"].(int)
		ret.KexEcdheSecp384r1Successes = in["kex_ecdhe_secp384r1_successes"].(int)
		ret.KexEcdheSecp521r1Successes = in["kex_ecdhe_secp521r1_successes"].(int)
		ret.KexEcdheX25519Successes = in["kex_ecdhe_x25519_successes"].(int)
		ret.KexEcdheX448Successes = in["kex_ecdhe_x448_successes"].(int)
		ret.KexEcdheSm2Successes = in["kex_ecdhe_sm2_successes"].(int)
		ret.KexEcdheSecp256r1Failures = in["kex_ecdhe_secp256r1_failures"].(int)
		ret.KexEcdheSecp384r1Failures = in["kex_ecdhe_secp384r1_failures"].(int)
		ret.KexEcdheSecp521r1Failures = in["kex_ecdhe_secp521r1_failures"].(int)
		ret.KexEcdheX25519Failures = in["kex_ecdhe_x25519_failures"].(int)
		ret.KexEcdheX448Failures = in["kex_ecdhe_x448_failures"].(int)
		ret.KexEcdheSm2Failures = in["kex_ecdhe_sm2_failures"].(int)
		ret.KexDhe512Successes = in["kex_dhe_512_successes"].(int)
		ret.KexDhe1024Successes = in["kex_dhe_1024_successes"].(int)
		ret.KexDhe2048Successes = in["kex_dhe_2048_successes"].(int)
		ret.KexDhe512Failures = in["kex_dhe_512_failures"].(int)
		ret.KexDhe1024Failures = in["kex_dhe_1024_failures"].(int)
		ret.KexDhe2048Failures = in["kex_dhe_2048_failures"].(int)
		ret.Ssl2Successes = in["ssl2_successes"].(int)
		ret.Ssl3Successes = in["ssl3_successes"].(int)
		ret.Tls10Successes = in["tls10_successes"].(int)
		ret.Tls11Successes = in["tls11_successes"].(int)
		ret.Tls12Successes = in["tls12_successes"].(int)
		ret.Tls13Successes = in["tls13_successes"].(int)
		ret.Ssl2Failures = in["ssl2_failures"].(int)
		ret.Ssl3Failures = in["ssl3_failures"].(int)
		ret.Tls10Failures = in["tls10_failures"].(int)
		ret.Tls11Failures = in["tls11_failures"].(int)
		ret.Tls12Failures = in["tls12_failures"].(int)
		ret.Tls13Failures = in["tls13_failures"].(int)
		ret.SessCacheNew = in["sess_cache_new"].(int)
		ret.SessCacheHit = in["sess_cache_hit"].(int)
		ret.SessCacheMiss = in["sess_cache_miss"].(int)
		ret.SessCacheTimeout = in["sess_cache_timeout"].(int)
		ret.SessCacheCurrConn = in["sess_cache_curr_conn"].(int)
		ret.HsFailures = in["hs_failures"].(int)
		ret.CertVfy = in["cert_vfy"].(int)
		ret.HsAvgTime = in["hs_avg_time"].(int)
		ret.SniAutomapSuccesses = in["sni_automap_successes"].(int)
		ret.SniAutomapFailures = in["sni_automap_failures"].(int)
		ret.SniAutomapConnClosed = in["sni_automap_conn_closed"].(int)
		ret.SniAutomapMaxActiveConn = in["sni_automap_max_active_conn"].(int)
		ret.SniAutomapMissingCert = in["sni_automap_missing_cert"].(int)
		ret.SniBypassMissingCert = in["sni_bypass_missing_cert"].(int)
		ret.SniBypassExpiredCert = in["sni_bypass_expired_cert"].(int)
		ret.SniBypassExplicitList = in["sni_bypass_explicit_list"].(int)
		ret.RenegotiationTotal = in["renegotiation_total"].(int)
		ret.RenegoSsl2Successes = in["renego_ssl2_successes"].(int)
		ret.RenegoSsl3Successes = in["renego_ssl3_successes"].(int)
		ret.RenegoTls10Successes = in["renego_tls10_successes"].(int)
		ret.RenegoTls11Successes = in["renego_tls11_successes"].(int)
		ret.RenegoTls12Successes = in["renego_tls12_successes"].(int)
		ret.RenegoTls13Successes = in["renego_tls13_successes"].(int)
		ret.RenegoSsl2Failures = in["renego_ssl2_failures"].(int)
		ret.RenegoSsl3Failures = in["renego_ssl3_failures"].(int)
		ret.RenegoTls10Failures = in["renego_tls10_failures"].(int)
		ret.RenegoTls11Failures = in["renego_tls11_failures"].(int)
		ret.RenegoTls12Failures = in["renego_tls12_failures"].(int)
		ret.RenegoTls13Failures = in["renego_tls13_failures"].(int)
		ret.Downgraded = in["downgraded"].(int)
	}
	return ret
}

func dataToEndpointSlbSslCountersOper(d *schema.ResourceData) edpt.SlbSslCountersOper {
	var ret edpt.SlbSslCountersOper

	ret.Oper = getObjectSlbSslCountersOperOper(d.Get("oper").([]interface{}))
	return ret
}
