package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslErrorStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_error_stats`: Statistics for the object ssl-error\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslErrorStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_data_in_handshake": {
							Type: schema.TypeInt, Optional: true, Description: "app data in handshake",
						},
						"attempt_to_reuse_sess_in_diff_context": {
							Type: schema.TypeInt, Optional: true, Description: "attempt to reuse sess in diff context",
						},
						"bad_alert_record": {
							Type: schema.TypeInt, Optional: true, Description: "bad alert record",
						},
						"bad_authentication_type": {
							Type: schema.TypeInt, Optional: true, Description: "bad authentication type",
						},
						"bad_change_cipher_spec": {
							Type: schema.TypeInt, Optional: true, Description: "bad change cipher spec",
						},
						"bad_checksum": {
							Type: schema.TypeInt, Optional: true, Description: "bad checksum",
						},
						"bad_data_returned_by_callback": {
							Type: schema.TypeInt, Optional: true, Description: "bad data returned by callback",
						},
						"bad_decompression": {
							Type: schema.TypeInt, Optional: true, Description: "bad decompression",
						},
						"bad_dh_g_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad dh g length",
						},
						"bad_dh_pub_key_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad dh pub key length",
						},
						"bad_dh_p_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad dh p length",
						},
						"bad_digest_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad digest length",
						},
						"bad_dsa_signature": {
							Type: schema.TypeInt, Optional: true, Description: "bad dsa signature",
						},
						"bad_hello_request": {
							Type: schema.TypeInt, Optional: true, Description: "bad hello request",
						},
						"bad_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad length",
						},
						"bad_mac_decode": {
							Type: schema.TypeInt, Optional: true, Description: "bad mac decode",
						},
						"bad_message_type": {
							Type: schema.TypeInt, Optional: true, Description: "bad message type",
						},
						"bad_packet_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad packet length",
						},
						"bad_protocol_version_counter": {
							Type: schema.TypeInt, Optional: true, Description: "bad protocol version counter",
						},
						"bad_response_argument": {
							Type: schema.TypeInt, Optional: true, Description: "bad response argument",
						},
						"bad_rsa_decrypt": {
							Type: schema.TypeInt, Optional: true, Description: "bad rsa decrypt",
						},
						"bad_rsa_encrypt": {
							Type: schema.TypeInt, Optional: true, Description: "bad rsa encrypt",
						},
						"bad_rsa_e_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad rsa e length",
						},
						"bad_rsa_modulus_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad rsa modulus length",
						},
						"bad_rsa_signature": {
							Type: schema.TypeInt, Optional: true, Description: "bad rsa signature",
						},
						"bad_signature": {
							Type: schema.TypeInt, Optional: true, Description: "bad signature",
						},
						"bad_ssl_filetype": {
							Type: schema.TypeInt, Optional: true, Description: "bad ssl filetype",
						},
						"bad_ssl_session_id_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad ssl session id length",
						},
						"bad_state": {
							Type: schema.TypeInt, Optional: true, Description: "bad state",
						},
						"bad_write_retry": {
							Type: schema.TypeInt, Optional: true, Description: "bad write retry",
						},
						"bio_not_set": {
							Type: schema.TypeInt, Optional: true, Description: "bio not set",
						},
						"block_cipher_pad_is_wrong": {
							Type: schema.TypeInt, Optional: true, Description: "block cipher pad is wrong",
						},
						"bn_lib": {
							Type: schema.TypeInt, Optional: true, Description: "bn lib",
						},
						"ca_dn_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "ca dn length mismatch",
						},
						"ca_dn_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "ca dn too long",
						},
						"ccs_received_early": {
							Type: schema.TypeInt, Optional: true, Description: "ccs received early",
						},
						"certificate_verify_failed": {
							Type: schema.TypeInt, Optional: true, Description: "certificate verify failed",
						},
						"cert_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "cert length mismatch",
						},
						"challenge_is_different": {
							Type: schema.TypeInt, Optional: true, Description: "challenge is different",
						},
						"cipher_code_wrong_length": {
							Type: schema.TypeInt, Optional: true, Description: "cipher code wrong length",
						},
						"cipher_or_hash_unavailable": {
							Type: schema.TypeInt, Optional: true, Description: "cipher or hash unavailable",
						},
						"cipher_table_src_error": {
							Type: schema.TypeInt, Optional: true, Description: "cipher table src error",
						},
						"compressed_length_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "compressed length too long",
						},
						"compression_failure": {
							Type: schema.TypeInt, Optional: true, Description: "compression failure",
						},
						"compression_library_error": {
							Type: schema.TypeInt, Optional: true, Description: "compression library error",
						},
						"connection_id_is_different": {
							Type: schema.TypeInt, Optional: true, Description: "connection id is different",
						},
						"connection_type_not_set": {
							Type: schema.TypeInt, Optional: true, Description: "connection type not set",
						},
						"data_between_ccs_and_finished": {
							Type: schema.TypeInt, Optional: true, Description: "data between ccs and finished",
						},
						"data_length_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "data length too long",
						},
						"decryption_failed": {
							Type: schema.TypeInt, Optional: true, Description: "decryption failed",
						},
						"decryption_failed_or_bad_record_mac": {
							Type: schema.TypeInt, Optional: true, Description: "decryption failed or bad record mac",
						},
						"dh_public_value_length_is_wrong": {
							Type: schema.TypeInt, Optional: true, Description: "dh public value length is wrong",
						},
						"digest_check_failed": {
							Type: schema.TypeInt, Optional: true, Description: "digest check failed",
						},
						"encrypted_length_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "encrypted length too long",
						},
						"error_generating_tmp_rsa_key": {
							Type: schema.TypeInt, Optional: true, Description: "error generating tmp rsa key",
						},
						"error_in_received_cipher_list": {
							Type: schema.TypeInt, Optional: true, Description: "error in received cipher list",
						},
						"excessive_message_size": {
							Type: schema.TypeInt, Optional: true, Description: "excessive message size",
						},
						"extra_data_in_message": {
							Type: schema.TypeInt, Optional: true, Description: "extra data in message",
						},
						"got_a_fin_before_a_ccs": {
							Type: schema.TypeInt, Optional: true, Description: "got a fin before a ccs",
						},
						"https_proxy_request": {
							Type: schema.TypeInt, Optional: true, Description: "https proxy request",
						},
						"http_request": {
							Type: schema.TypeInt, Optional: true, Description: "http request",
						},
						"illegal_padding": {
							Type: schema.TypeInt, Optional: true, Description: "illegal padding",
						},
						"inappropriate_fallback": {
							Type: schema.TypeInt, Optional: true, Description: "inappropriate fallback",
						},
						"invalid_challenge_length": {
							Type: schema.TypeInt, Optional: true, Description: "invalid challenge length",
						},
						"invalid_command": {
							Type: schema.TypeInt, Optional: true, Description: "invalid command",
						},
						"invalid_purpose": {
							Type: schema.TypeInt, Optional: true, Description: "invalid purpose",
						},
						"invalid_status_response": {
							Type: schema.TypeInt, Optional: true, Description: "invalid status response",
						},
						"invalid_trust": {
							Type: schema.TypeInt, Optional: true, Description: "invalid trust",
						},
						"key_arg_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "key arg too long",
						},
						"krb5": {
							Type: schema.TypeInt, Optional: true, Description: "krb5",
						},
						"krb5_client_cc_principal": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 client cc principal",
						},
						"krb5_client_get_cred": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 client get cred",
						},
						"krb5_client_init": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 client init",
						},
						"krb5_client_mk_req": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 client mk_req",
						},
						"krb5_server_bad_ticket": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 server bad ticket",
						},
						"krb5_server_init": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 server init",
						},
						"krb5_server_rd_req": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 server rd_req",
						},
						"krb5_server_tkt_expired": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 server tkt expired",
						},
						"krb5_server_tkt_not_yet_valid": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 server tkt not yet valid",
						},
						"krb5_server_tkt_skew": {
							Type: schema.TypeInt, Optional: true, Description: "krb5 server tkt skew",
						},
						"length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "length mismatch",
						},
						"length_too_short": {
							Type: schema.TypeInt, Optional: true, Description: "length too short",
						},
						"library_bug": {
							Type: schema.TypeInt, Optional: true, Description: "library bug",
						},
						"library_has_no_ciphers": {
							Type: schema.TypeInt, Optional: true, Description: "library has no ciphers",
						},
						"mast_key_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "mast key too long",
						},
						"message_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "message too long",
						},
						"missing_dh_dsa_cert": {
							Type: schema.TypeInt, Optional: true, Description: "missing dh dsa cert",
						},
						"missing_dh_key": {
							Type: schema.TypeInt, Optional: true, Description: "missing dh key",
						},
						"missing_dh_rsa_cert": {
							Type: schema.TypeInt, Optional: true, Description: "missing dh rsa cert",
						},
						"missing_dsa_signing_cert": {
							Type: schema.TypeInt, Optional: true, Description: "missing dsa signing cert",
						},
						"missing_export_tmp_dh_key": {
							Type: schema.TypeInt, Optional: true, Description: "missing export tmp dh key",
						},
						"missing_export_tmp_rsa_key": {
							Type: schema.TypeInt, Optional: true, Description: "missing export tmp rsa key",
						},
						"missing_rsa_certificate": {
							Type: schema.TypeInt, Optional: true, Description: "missing rsa certificate",
						},
						"missing_rsa_encrypting_cert": {
							Type: schema.TypeInt, Optional: true, Description: "missing rsa encrypting cert",
						},
						"missing_rsa_signing_cert": {
							Type: schema.TypeInt, Optional: true, Description: "missing rsa signing cert",
						},
						"missing_tmp_dh_key": {
							Type: schema.TypeInt, Optional: true, Description: "missing tmp dh key",
						},
						"missing_tmp_rsa_key": {
							Type: schema.TypeInt, Optional: true, Description: "missing tmp rsa key",
						},
						"missing_tmp_rsa_pkey": {
							Type: schema.TypeInt, Optional: true, Description: "missing tmp rsa pkey",
						},
						"missing_verify_message": {
							Type: schema.TypeInt, Optional: true, Description: "missing verify message",
						},
						"non_sslv2_initial_packet": {
							Type: schema.TypeInt, Optional: true, Description: "non sslv2 initial packet",
						},
						"no_certificates_returned": {
							Type: schema.TypeInt, Optional: true, Description: "no certificates returned",
						},
						"no_certificate_assigned": {
							Type: schema.TypeInt, Optional: true, Description: "no certificate assigned",
						},
						"no_certificate_returned": {
							Type: schema.TypeInt, Optional: true, Description: "no certificate returned",
						},
						"no_certificate_set": {
							Type: schema.TypeInt, Optional: true, Description: "no certificate set",
						},
						"no_certificate_specified": {
							Type: schema.TypeInt, Optional: true, Description: "no certificate specified",
						},
						"no_ciphers_available": {
							Type: schema.TypeInt, Optional: true, Description: "no ciphers available",
						},
						"no_ciphers_passed": {
							Type: schema.TypeInt, Optional: true, Description: "no ciphers passed",
						},
						"no_ciphers_specified": {
							Type: schema.TypeInt, Optional: true, Description: "no ciphers specified",
						},
						"no_cipher_list": {
							Type: schema.TypeInt, Optional: true, Description: "no cipher list",
						},
						"no_cipher_match": {
							Type: schema.TypeInt, Optional: true, Description: "no cipher match",
						},
						"no_client_cert_received": {
							Type: schema.TypeInt, Optional: true, Description: "no client cert received",
						},
						"no_compression_specified": {
							Type: schema.TypeInt, Optional: true, Description: "no compression specified",
						},
						"no_method_specified": {
							Type: schema.TypeInt, Optional: true, Description: "no method specified",
						},
						"no_privatekey": {
							Type: schema.TypeInt, Optional: true, Description: "no privatekey",
						},
						"no_private_key_assigned": {
							Type: schema.TypeInt, Optional: true, Description: "no private key assigned",
						},
						"no_protocols_available": {
							Type: schema.TypeInt, Optional: true, Description: "no protocols available",
						},
						"no_publickey": {
							Type: schema.TypeInt, Optional: true, Description: "no publickey",
						},
						"no_shared_cipher": {
							Type: schema.TypeInt, Optional: true, Description: "no shared cipher",
						},
						"no_verify_callback": {
							Type: schema.TypeInt, Optional: true, Description: "no verify callback",
						},
						"null_ssl_ctx": {
							Type: schema.TypeInt, Optional: true, Description: "null ssl ctx",
						},
						"null_ssl_method_passed": {
							Type: schema.TypeInt, Optional: true, Description: "null ssl method passed",
						},
						"old_session_cipher_not_returned": {
							Type: schema.TypeInt, Optional: true, Description: "old session cipher not returned",
						},
						"packet_length_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "packet length too long",
						},
						"path_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "path too long",
						},
						"peer_did_not_return_a_certificate": {
							Type: schema.TypeInt, Optional: true, Description: "peer did not return a certificate",
						},
						"peer_error": {
							Type: schema.TypeInt, Optional: true, Description: "peer error",
						},
						"peer_error_certificate": {
							Type: schema.TypeInt, Optional: true, Description: "peer error certificate",
						},
						"peer_error_no_certificate": {
							Type: schema.TypeInt, Optional: true, Description: "peer error no certificate",
						},
						"peer_error_no_cipher": {
							Type: schema.TypeInt, Optional: true, Description: "peer error no cipher",
						},
						"peer_error_unsupported_certificate_type": {
							Type: schema.TypeInt, Optional: true, Description: "peer error unsupported certificate type",
						},
						"pre_mac_length_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "pre mac length too long",
						},
						"problems_mapping_cipher_functions": {
							Type: schema.TypeInt, Optional: true, Description: "problems mapping cipher functions",
						},
						"protocol_is_shutdown": {
							Type: schema.TypeInt, Optional: true, Description: "protocol is shutdown",
						},
						"public_key_encrypt_error": {
							Type: schema.TypeInt, Optional: true, Description: "public key encrypt error",
						},
						"public_key_is_not_rsa": {
							Type: schema.TypeInt, Optional: true, Description: "public key is not rsa",
						},
						"public_key_not_rsa": {
							Type: schema.TypeInt, Optional: true, Description: "public key not rsa",
						},
						"read_bio_not_set": {
							Type: schema.TypeInt, Optional: true, Description: "read bio not set",
						},
						"read_wrong_packet_type": {
							Type: schema.TypeInt, Optional: true, Description: "read wrong packet type",
						},
						"record_length_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "record length mismatch",
						},
						"record_too_large": {
							Type: schema.TypeInt, Optional: true, Description: "record too large",
						},
						"record_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "record too small",
						},
						"required_cipher_missing": {
							Type: schema.TypeInt, Optional: true, Description: "required cipher missing",
						},
						"reuse_cert_length_not_zero": {
							Type: schema.TypeInt, Optional: true, Description: "reuse cert length not zero",
						},
						"reuse_cert_type_not_zero": {
							Type: schema.TypeInt, Optional: true, Description: "reuse cert type not zero",
						},
						"reuse_cipher_list_not_zero": {
							Type: schema.TypeInt, Optional: true, Description: "reuse cipher list not zero",
						},
						"scsv_received_when_renegotiating": {
							Type: schema.TypeInt, Optional: true, Description: "scsv received when renegotiating",
						},
						"session_id_context_uninitialized": {
							Type: schema.TypeInt, Optional: true, Description: "session id context uninitialized",
						},
						"short_read": {
							Type: schema.TypeInt, Optional: true, Description: "short read",
						},
						"signature_for_non_signing_certificate": {
							Type: schema.TypeInt, Optional: true, Description: "signature for non signing certificate",
						},
						"ssl23_doing_session_id_reuse": {
							Type: schema.TypeInt, Optional: true, Description: "ssl23 doing session id reuse",
						},
						"ssl2_connection_id_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "ssl2 connection id too long",
						},
						"ssl3_session_id_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "ssl3 session id too long",
						},
						"ssl3_session_id_too_short": {
							Type: schema.TypeInt, Optional: true, Description: "ssl3 session id too short",
						},
						"sslv3_alert_bad_certificate": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert bad certificate",
						},
						"sslv3_alert_bad_record_mac": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert bad record mac",
						},
						"sslv3_alert_certificate_expired": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert certificate expired",
						},
						"sslv3_alert_certificate_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert certificate revoked",
						},
						"sslv3_alert_certificate_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert certificate unknown",
						},
						"sslv3_alert_decompression_failure": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert decompression failure",
						},
						"sslv3_alert_handshake_failure": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert handshake failure",
						},
						"sslv3_alert_illegal_parameter": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert illegal parameter",
						},
						"sslv3_alert_no_certificate": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert no certificate",
						},
						"sslv3_alert_peer_error_cert": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert peer error cert",
						},
						"sslv3_alert_peer_error_no_cert": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert peer error no cert",
						},
						"sslv3_alert_peer_error_no_cipher": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert peer error no cipher",
						},
						"sslv3_alert_peer_error_unsupp_cert_type": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert peer error unsupp cert type",
						},
						"sslv3_alert_unexpected_msg": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert unexpected msg",
						},
						"sslv3_alert_unknown_remote_err_type": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert unknown remote err type",
						},
						"sslv3_alert_unspported_cert": {
							Type: schema.TypeInt, Optional: true, Description: "sslv3 alert unspported cert",
						},
						"ssl_ctx_has_no_default_ssl_version": {
							Type: schema.TypeInt, Optional: true, Description: "ssl ctx has no default ssl version",
						},
						"ssl_handshake_failure": {
							Type: schema.TypeInt, Optional: true, Description: "ssl handshake failure",
						},
						"ssl_library_has_no_ciphers": {
							Type: schema.TypeInt, Optional: true, Description: "ssl library has no ciphers",
						},
						"ssl_session_id_callback_failed": {
							Type: schema.TypeInt, Optional: true, Description: "ssl session id callback failed",
						},
						"ssl_session_id_conflict": {
							Type: schema.TypeInt, Optional: true, Description: "ssl session id conflict",
						},
						"ssl_session_id_context_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "ssl session id context too long",
						},
						"ssl_session_id_has_bad_length": {
							Type: schema.TypeInt, Optional: true, Description: "ssl session id has bad length",
						},
						"ssl_session_id_is_different": {
							Type: schema.TypeInt, Optional: true, Description: "ssl session id is different",
						},
						"tlsv1_alert_access_denied": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert access denied",
						},
						"tlsv1_alert_decode_error": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert decode error",
						},
						"tlsv1_alert_decryption_failed": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert decryption failed",
						},
						"tlsv1_alert_decrypt_error": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert decrypt error",
						},
						"tlsv1_alert_export_restriction": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert export restriction",
						},
						"tlsv1_alert_insufficient_security": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert insufficient security",
						},
						"tlsv1_alert_internal_error": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert internal error",
						},
						"tlsv1_alert_no_renegotiation": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert no renegotiation",
						},
						"tlsv1_alert_protocol_version": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert protocol version",
						},
						"tlsv1_alert_record_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert record overflow",
						},
						"tlsv1_alert_unknown_ca": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert unknown ca",
						},
						"tlsv1_alert_user_cancelled": {
							Type: schema.TypeInt, Optional: true, Description: "tlsv1 alert user cancelled",
						},
						"tls_client_cert_req_with_anon_cipher": {
							Type: schema.TypeInt, Optional: true, Description: "tls client cert req with anon cipher",
						},
						"tls_peer_did_not_respond_with_cert_list": {
							Type: schema.TypeInt, Optional: true, Description: "tls peer did not respond with cert list",
						},
						"tls_rsa_encrypted_value_length_is_wrong": {
							Type: schema.TypeInt, Optional: true, Description: "tls rsa encrypted value length is wrong",
						},
						"tried_to_use_unsupported_cipher": {
							Type: schema.TypeInt, Optional: true, Description: "tried to use unsupported cipher",
						},
						"unable_to_decode_dh_certs": {
							Type: schema.TypeInt, Optional: true, Description: "unable to decode dh certs",
						},
						"unable_to_extract_public_key": {
							Type: schema.TypeInt, Optional: true, Description: "unable to extract public key",
						},
						"unable_to_find_dh_parameters": {
							Type: schema.TypeInt, Optional: true, Description: "unable to find dh parameters",
						},
						"unable_to_find_public_key_parameters": {
							Type: schema.TypeInt, Optional: true, Description: "unable to find public key parameters",
						},
						"unable_to_find_ssl_method": {
							Type: schema.TypeInt, Optional: true, Description: "unable to find ssl method",
						},
						"unable_to_load_ssl2_md5_routines": {
							Type: schema.TypeInt, Optional: true, Description: "unable to load ssl2 md5 routines",
						},
						"unable_to_load_ssl3_md5_routines": {
							Type: schema.TypeInt, Optional: true, Description: "unable to load ssl3 md5 routines",
						},
						"unable_to_load_ssl3_sha1_routines": {
							Type: schema.TypeInt, Optional: true, Description: "unable to load ssl3 sha1 routines",
						},
						"unexpected_message": {
							Type: schema.TypeInt, Optional: true, Description: "unexpected message",
						},
						"unexpected_record": {
							Type: schema.TypeInt, Optional: true, Description: "unexpected record",
						},
						"uninitialized": {
							Type: schema.TypeInt, Optional: true, Description: "uninitialized",
						},
						"unknown_alert_type": {
							Type: schema.TypeInt, Optional: true, Description: "unknown alert type",
						},
						"unknown_certificate_type": {
							Type: schema.TypeInt, Optional: true, Description: "unknown certificate type",
						},
						"unknown_cipher_returned": {
							Type: schema.TypeInt, Optional: true, Description: "unknown cipher returned",
						},
						"unknown_cipher_type": {
							Type: schema.TypeInt, Optional: true, Description: "unknown cipher type",
						},
						"unknown_key_exchange_type": {
							Type: schema.TypeInt, Optional: true, Description: "unknown key exchange type",
						},
						"unknown_pkey_type": {
							Type: schema.TypeInt, Optional: true, Description: "unknown pkey type",
						},
						"unknown_protocol": {
							Type: schema.TypeInt, Optional: true, Description: "unknown protocol",
						},
						"unknown_remote_error_type": {
							Type: schema.TypeInt, Optional: true, Description: "unknown remote error type",
						},
						"unknown_ssl_version": {
							Type: schema.TypeInt, Optional: true, Description: "unknown ssl version",
						},
						"unknown_state": {
							Type: schema.TypeInt, Optional: true, Description: "unknown state",
						},
						"unsupported_cipher": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported cipher",
						},
						"unsupported_compression_algorithm": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported compression algorithm",
						},
						"unsupported_option": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported option",
						},
						"unsupported_protocol": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported protocol",
						},
						"unsupported_ssl_version": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported ssl version",
						},
						"unsupported_status_type": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported status type",
						},
						"write_bio_not_set": {
							Type: schema.TypeInt, Optional: true, Description: "write bio not set",
						},
						"wrong_cipher_returned": {
							Type: schema.TypeInt, Optional: true, Description: "wrong cipher returned",
						},
						"wrong_message_type": {
							Type: schema.TypeInt, Optional: true, Description: "wrong message type",
						},
						"wrong_counter_of_key_bits": {
							Type: schema.TypeInt, Optional: true, Description: "wrong counter of key bits",
						},
						"wrong_signature_length": {
							Type: schema.TypeInt, Optional: true, Description: "wrong signature length",
						},
						"wrong_signature_size": {
							Type: schema.TypeInt, Optional: true, Description: "wrong signature size",
						},
						"wrong_ssl_version": {
							Type: schema.TypeInt, Optional: true, Description: "wrong ssl version",
						},
						"wrong_version_counter": {
							Type: schema.TypeInt, Optional: true, Description: "wrong version counter",
						},
						"x509_lib": {
							Type: schema.TypeInt, Optional: true, Description: "x509 lib",
						},
						"x509_verification_setup_problems": {
							Type: schema.TypeInt, Optional: true, Description: "x509 verification setup problems",
						},
						"clienthello_tlsext": {
							Type: schema.TypeInt, Optional: true, Description: "clienthello tlsext",
						},
						"parse_tlsext": {
							Type: schema.TypeInt, Optional: true, Description: "parse tlsext",
						},
						"serverhello_tlsext": {
							Type: schema.TypeInt, Optional: true, Description: "serverhello tlsext",
						},
						"ssl3_ext_invalid_servername": {
							Type: schema.TypeInt, Optional: true, Description: "ssl3 ext invalid servername",
						},
						"ssl3_ext_invalid_servername_type": {
							Type: schema.TypeInt, Optional: true, Description: "ssl3 ext invalid servername type",
						},
						"multiple_sgc_restarts": {
							Type: schema.TypeInt, Optional: true, Description: "multiple sgc restarts",
						},
						"tls_invalid_ecpointformat_list": {
							Type: schema.TypeInt, Optional: true, Description: "tls invalid ecpointformat list",
						},
						"bad_ecc_cert": {
							Type: schema.TypeInt, Optional: true, Description: "bad ecc cert",
						},
						"bad_ecdsa_sig": {
							Type: schema.TypeInt, Optional: true, Description: "bad ecdsa sig",
						},
						"bad_ecpoint": {
							Type: schema.TypeInt, Optional: true, Description: "bad ecpoint",
						},
						"cookie_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "cookie mismatch",
						},
						"unsupported_elliptic_curve": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported elliptic curve",
						},
						"no_required_digest": {
							Type: schema.TypeInt, Optional: true, Description: "no required digest",
						},
						"unsupported_digest_type": {
							Type: schema.TypeInt, Optional: true, Description: "unsupported digest type",
						},
						"bad_handshake_length": {
							Type: schema.TypeInt, Optional: true, Description: "bad handshake length",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslErrorStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslErrorStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslErrorStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslErrorStatsStats := setObjectSlbSslErrorStatsStats(res)
		d.Set("stats", SlbSslErrorStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslErrorStatsStats(ret edpt.DataSlbSslErrorStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"app_data_in_handshake":                   ret.DtSlbSslErrorStats.Stats.AppDataInHandshake,
			"attempt_to_reuse_sess_in_diff_context":   ret.DtSlbSslErrorStats.Stats.AttemptToReuseSessInDiffContext,
			"bad_alert_record":                        ret.DtSlbSslErrorStats.Stats.BadAlertRecord,
			"bad_authentication_type":                 ret.DtSlbSslErrorStats.Stats.BadAuthenticationType,
			"bad_change_cipher_spec":                  ret.DtSlbSslErrorStats.Stats.BadChangeCipherSpec,
			"bad_checksum":                            ret.DtSlbSslErrorStats.Stats.BadChecksum,
			"bad_data_returned_by_callback":           ret.DtSlbSslErrorStats.Stats.BadDataReturnedByCallback,
			"bad_decompression":                       ret.DtSlbSslErrorStats.Stats.BadDecompression,
			"bad_dh_g_length":                         ret.DtSlbSslErrorStats.Stats.BadDhGLength,
			"bad_dh_pub_key_length":                   ret.DtSlbSslErrorStats.Stats.BadDhPubKeyLength,
			"bad_dh_p_length":                         ret.DtSlbSslErrorStats.Stats.BadDhPLength,
			"bad_digest_length":                       ret.DtSlbSslErrorStats.Stats.BadDigestLength,
			"bad_dsa_signature":                       ret.DtSlbSslErrorStats.Stats.BadDsaSignature,
			"bad_hello_request":                       ret.DtSlbSslErrorStats.Stats.BadHelloRequest,
			"bad_length":                              ret.DtSlbSslErrorStats.Stats.BadLength,
			"bad_mac_decode":                          ret.DtSlbSslErrorStats.Stats.BadMacDecode,
			"bad_message_type":                        ret.DtSlbSslErrorStats.Stats.BadMessageType,
			"bad_packet_length":                       ret.DtSlbSslErrorStats.Stats.BadPacketLength,
			"bad_protocol_version_counter":            ret.DtSlbSslErrorStats.Stats.BadProtocolVersionCounter,
			"bad_response_argument":                   ret.DtSlbSslErrorStats.Stats.BadResponseArgument,
			"bad_rsa_decrypt":                         ret.DtSlbSslErrorStats.Stats.BadRsaDecrypt,
			"bad_rsa_encrypt":                         ret.DtSlbSslErrorStats.Stats.BadRsaEncrypt,
			"bad_rsa_e_length":                        ret.DtSlbSslErrorStats.Stats.BadRsaELength,
			"bad_rsa_modulus_length":                  ret.DtSlbSslErrorStats.Stats.BadRsaModulusLength,
			"bad_rsa_signature":                       ret.DtSlbSslErrorStats.Stats.BadRsaSignature,
			"bad_signature":                           ret.DtSlbSslErrorStats.Stats.BadSignature,
			"bad_ssl_filetype":                        ret.DtSlbSslErrorStats.Stats.BadSslFiletype,
			"bad_ssl_session_id_length":               ret.DtSlbSslErrorStats.Stats.BadSslSessionIdLength,
			"bad_state":                               ret.DtSlbSslErrorStats.Stats.BadState,
			"bad_write_retry":                         ret.DtSlbSslErrorStats.Stats.BadWriteRetry,
			"bio_not_set":                             ret.DtSlbSslErrorStats.Stats.BioNotSet,
			"block_cipher_pad_is_wrong":               ret.DtSlbSslErrorStats.Stats.BlockCipherPadIsWrong,
			"bn_lib":                                  ret.DtSlbSslErrorStats.Stats.BnLib,
			"ca_dn_length_mismatch":                   ret.DtSlbSslErrorStats.Stats.CaDnLengthMismatch,
			"ca_dn_too_long":                          ret.DtSlbSslErrorStats.Stats.CaDnTooLong,
			"ccs_received_early":                      ret.DtSlbSslErrorStats.Stats.CcsReceivedEarly,
			"certificate_verify_failed":               ret.DtSlbSslErrorStats.Stats.CertificateVerifyFailed,
			"cert_length_mismatch":                    ret.DtSlbSslErrorStats.Stats.CertLengthMismatch,
			"challenge_is_different":                  ret.DtSlbSslErrorStats.Stats.ChallengeIsDifferent,
			"cipher_code_wrong_length":                ret.DtSlbSslErrorStats.Stats.CipherCodeWrongLength,
			"cipher_or_hash_unavailable":              ret.DtSlbSslErrorStats.Stats.CipherOrHashUnavailable,
			"cipher_table_src_error":                  ret.DtSlbSslErrorStats.Stats.CipherTableSrcError,
			"compressed_length_too_long":              ret.DtSlbSslErrorStats.Stats.CompressedLengthTooLong,
			"compression_failure":                     ret.DtSlbSslErrorStats.Stats.CompressionFailure,
			"compression_library_error":               ret.DtSlbSslErrorStats.Stats.CompressionLibraryError,
			"connection_id_is_different":              ret.DtSlbSslErrorStats.Stats.ConnectionIdIsDifferent,
			"connection_type_not_set":                 ret.DtSlbSslErrorStats.Stats.ConnectionTypeNotSet,
			"data_between_ccs_and_finished":           ret.DtSlbSslErrorStats.Stats.DataBetweenCcsAndFinished,
			"data_length_too_long":                    ret.DtSlbSslErrorStats.Stats.DataLengthTooLong,
			"decryption_failed":                       ret.DtSlbSslErrorStats.Stats.DecryptionFailed,
			"decryption_failed_or_bad_record_mac":     ret.DtSlbSslErrorStats.Stats.DecryptionFailedOrBadRecordMac,
			"dh_public_value_length_is_wrong":         ret.DtSlbSslErrorStats.Stats.DhPublicValueLengthIsWrong,
			"digest_check_failed":                     ret.DtSlbSslErrorStats.Stats.DigestCheckFailed,
			"encrypted_length_too_long":               ret.DtSlbSslErrorStats.Stats.EncryptedLengthTooLong,
			"error_generating_tmp_rsa_key":            ret.DtSlbSslErrorStats.Stats.ErrorGeneratingTmpRsaKey,
			"error_in_received_cipher_list":           ret.DtSlbSslErrorStats.Stats.ErrorInReceivedCipherList,
			"excessive_message_size":                  ret.DtSlbSslErrorStats.Stats.ExcessiveMessageSize,
			"extra_data_in_message":                   ret.DtSlbSslErrorStats.Stats.ExtraDataInMessage,
			"got_a_fin_before_a_ccs":                  ret.DtSlbSslErrorStats.Stats.GotAFinBeforeACcs,
			"https_proxy_request":                     ret.DtSlbSslErrorStats.Stats.HttpsProxyRequest,
			"http_request":                            ret.DtSlbSslErrorStats.Stats.HttpRequest,
			"illegal_padding":                         ret.DtSlbSslErrorStats.Stats.IllegalPadding,
			"inappropriate_fallback":                  ret.DtSlbSslErrorStats.Stats.InappropriateFallback,
			"invalid_challenge_length":                ret.DtSlbSslErrorStats.Stats.InvalidChallengeLength,
			"invalid_command":                         ret.DtSlbSslErrorStats.Stats.InvalidCommand,
			"invalid_purpose":                         ret.DtSlbSslErrorStats.Stats.InvalidPurpose,
			"invalid_status_response":                 ret.DtSlbSslErrorStats.Stats.InvalidStatusResponse,
			"invalid_trust":                           ret.DtSlbSslErrorStats.Stats.InvalidTrust,
			"key_arg_too_long":                        ret.DtSlbSslErrorStats.Stats.KeyArgTooLong,
			"krb5":                                    ret.DtSlbSslErrorStats.Stats.Krb5,
			"krb5_client_cc_principal":                ret.DtSlbSslErrorStats.Stats.Krb5ClientCcPrincipal,
			"krb5_client_get_cred":                    ret.DtSlbSslErrorStats.Stats.Krb5ClientGetCred,
			"krb5_client_init":                        ret.DtSlbSslErrorStats.Stats.Krb5ClientInit,
			"krb5_client_mk_req":                      ret.DtSlbSslErrorStats.Stats.Krb5ClientMkReq,
			"krb5_server_bad_ticket":                  ret.DtSlbSslErrorStats.Stats.Krb5ServerBadTicket,
			"krb5_server_init":                        ret.DtSlbSslErrorStats.Stats.Krb5ServerInit,
			"krb5_server_rd_req":                      ret.DtSlbSslErrorStats.Stats.Krb5ServerRdReq,
			"krb5_server_tkt_expired":                 ret.DtSlbSslErrorStats.Stats.Krb5ServerTktExpired,
			"krb5_server_tkt_not_yet_valid":           ret.DtSlbSslErrorStats.Stats.Krb5ServerTktNotYetValid,
			"krb5_server_tkt_skew":                    ret.DtSlbSslErrorStats.Stats.Krb5ServerTktSkew,
			"length_mismatch":                         ret.DtSlbSslErrorStats.Stats.LengthMismatch,
			"length_too_short":                        ret.DtSlbSslErrorStats.Stats.LengthTooShort,
			"library_bug":                             ret.DtSlbSslErrorStats.Stats.LibraryBug,
			"library_has_no_ciphers":                  ret.DtSlbSslErrorStats.Stats.LibraryHasNoCiphers,
			"mast_key_too_long":                       ret.DtSlbSslErrorStats.Stats.MastKeyTooLong,
			"message_too_long":                        ret.DtSlbSslErrorStats.Stats.MessageTooLong,
			"missing_dh_dsa_cert":                     ret.DtSlbSslErrorStats.Stats.MissingDhDsaCert,
			"missing_dh_key":                          ret.DtSlbSslErrorStats.Stats.MissingDhKey,
			"missing_dh_rsa_cert":                     ret.DtSlbSslErrorStats.Stats.MissingDhRsaCert,
			"missing_dsa_signing_cert":                ret.DtSlbSslErrorStats.Stats.MissingDsaSigningCert,
			"missing_export_tmp_dh_key":               ret.DtSlbSslErrorStats.Stats.MissingExportTmpDhKey,
			"missing_export_tmp_rsa_key":              ret.DtSlbSslErrorStats.Stats.MissingExportTmpRsaKey,
			"missing_rsa_certificate":                 ret.DtSlbSslErrorStats.Stats.MissingRsaCertificate,
			"missing_rsa_encrypting_cert":             ret.DtSlbSslErrorStats.Stats.MissingRsaEncryptingCert,
			"missing_rsa_signing_cert":                ret.DtSlbSslErrorStats.Stats.MissingRsaSigningCert,
			"missing_tmp_dh_key":                      ret.DtSlbSslErrorStats.Stats.MissingTmpDhKey,
			"missing_tmp_rsa_key":                     ret.DtSlbSslErrorStats.Stats.MissingTmpRsaKey,
			"missing_tmp_rsa_pkey":                    ret.DtSlbSslErrorStats.Stats.MissingTmpRsaPkey,
			"missing_verify_message":                  ret.DtSlbSslErrorStats.Stats.MissingVerifyMessage,
			"non_sslv2_initial_packet":                ret.DtSlbSslErrorStats.Stats.NonSslv2InitialPacket,
			"no_certificates_returned":                ret.DtSlbSslErrorStats.Stats.NoCertificatesReturned,
			"no_certificate_assigned":                 ret.DtSlbSslErrorStats.Stats.NoCertificateAssigned,
			"no_certificate_returned":                 ret.DtSlbSslErrorStats.Stats.NoCertificateReturned,
			"no_certificate_set":                      ret.DtSlbSslErrorStats.Stats.NoCertificateSet,
			"no_certificate_specified":                ret.DtSlbSslErrorStats.Stats.NoCertificateSpecified,
			"no_ciphers_available":                    ret.DtSlbSslErrorStats.Stats.NoCiphersAvailable,
			"no_ciphers_passed":                       ret.DtSlbSslErrorStats.Stats.NoCiphersPassed,
			"no_ciphers_specified":                    ret.DtSlbSslErrorStats.Stats.NoCiphersSpecified,
			"no_cipher_list":                          ret.DtSlbSslErrorStats.Stats.NoCipherList,
			"no_cipher_match":                         ret.DtSlbSslErrorStats.Stats.NoCipherMatch,
			"no_client_cert_received":                 ret.DtSlbSslErrorStats.Stats.NoClientCertReceived,
			"no_compression_specified":                ret.DtSlbSslErrorStats.Stats.NoCompressionSpecified,
			"no_method_specified":                     ret.DtSlbSslErrorStats.Stats.NoMethodSpecified,
			"no_privatekey":                           ret.DtSlbSslErrorStats.Stats.NoPrivatekey,
			"no_private_key_assigned":                 ret.DtSlbSslErrorStats.Stats.NoPrivateKeyAssigned,
			"no_protocols_available":                  ret.DtSlbSslErrorStats.Stats.NoProtocolsAvailable,
			"no_publickey":                            ret.DtSlbSslErrorStats.Stats.NoPublickey,
			"no_shared_cipher":                        ret.DtSlbSslErrorStats.Stats.NoSharedCipher,
			"no_verify_callback":                      ret.DtSlbSslErrorStats.Stats.NoVerifyCallback,
			"null_ssl_ctx":                            ret.DtSlbSslErrorStats.Stats.NullSslCtx,
			"null_ssl_method_passed":                  ret.DtSlbSslErrorStats.Stats.NullSslMethodPassed,
			"old_session_cipher_not_returned":         ret.DtSlbSslErrorStats.Stats.OldSessionCipherNotReturned,
			"packet_length_too_long":                  ret.DtSlbSslErrorStats.Stats.PacketLengthTooLong,
			"path_too_long":                           ret.DtSlbSslErrorStats.Stats.PathTooLong,
			"peer_did_not_return_a_certificate":       ret.DtSlbSslErrorStats.Stats.PeerDidNotReturnACertificate,
			"peer_error":                              ret.DtSlbSslErrorStats.Stats.PeerError,
			"peer_error_certificate":                  ret.DtSlbSslErrorStats.Stats.PeerErrorCertificate,
			"peer_error_no_certificate":               ret.DtSlbSslErrorStats.Stats.PeerErrorNoCertificate,
			"peer_error_no_cipher":                    ret.DtSlbSslErrorStats.Stats.PeerErrorNoCipher,
			"peer_error_unsupported_certificate_type": ret.DtSlbSslErrorStats.Stats.PeerErrorUnsupportedCertificateType,
			"pre_mac_length_too_long":                 ret.DtSlbSslErrorStats.Stats.PreMacLengthTooLong,
			"problems_mapping_cipher_functions":       ret.DtSlbSslErrorStats.Stats.ProblemsMappingCipherFunctions,
			"protocol_is_shutdown":                    ret.DtSlbSslErrorStats.Stats.ProtocolIsShutdown,
			"public_key_encrypt_error":                ret.DtSlbSslErrorStats.Stats.PublicKeyEncryptError,
			"public_key_is_not_rsa":                   ret.DtSlbSslErrorStats.Stats.PublicKeyIsNotRsa,
			"public_key_not_rsa":                      ret.DtSlbSslErrorStats.Stats.PublicKeyNotRsa,
			"read_bio_not_set":                        ret.DtSlbSslErrorStats.Stats.ReadBioNotSet,
			"read_wrong_packet_type":                  ret.DtSlbSslErrorStats.Stats.ReadWrongPacketType,
			"record_length_mismatch":                  ret.DtSlbSslErrorStats.Stats.RecordLengthMismatch,
			"record_too_large":                        ret.DtSlbSslErrorStats.Stats.RecordTooLarge,
			"record_too_small":                        ret.DtSlbSslErrorStats.Stats.RecordTooSmall,
			"required_cipher_missing":                 ret.DtSlbSslErrorStats.Stats.RequiredCipherMissing,
			"reuse_cert_length_not_zero":              ret.DtSlbSslErrorStats.Stats.ReuseCertLengthNotZero,
			"reuse_cert_type_not_zero":                ret.DtSlbSslErrorStats.Stats.ReuseCertTypeNotZero,
			"reuse_cipher_list_not_zero":              ret.DtSlbSslErrorStats.Stats.ReuseCipherListNotZero,
			"scsv_received_when_renegotiating":        ret.DtSlbSslErrorStats.Stats.ScsvReceivedWhenRenegotiating,
			"session_id_context_uninitialized":        ret.DtSlbSslErrorStats.Stats.SessionIdContextUninitialized,
			"short_read":                              ret.DtSlbSslErrorStats.Stats.ShortRead,
			"signature_for_non_signing_certificate":   ret.DtSlbSslErrorStats.Stats.SignatureForNonSigningCertificate,
			"ssl23_doing_session_id_reuse":            ret.DtSlbSslErrorStats.Stats.Ssl23DoingSessionIdReuse,
			"ssl2_connection_id_too_long":             ret.DtSlbSslErrorStats.Stats.Ssl2ConnectionIdTooLong,
			"ssl3_session_id_too_long":                ret.DtSlbSslErrorStats.Stats.Ssl3SessionIdTooLong,
			"ssl3_session_id_too_short":               ret.DtSlbSslErrorStats.Stats.Ssl3SessionIdTooShort,
			"sslv3_alert_bad_certificate":             ret.DtSlbSslErrorStats.Stats.Sslv3AlertBadCertificate,
			"sslv3_alert_bad_record_mac":              ret.DtSlbSslErrorStats.Stats.Sslv3AlertBadRecordMac,
			"sslv3_alert_certificate_expired":         ret.DtSlbSslErrorStats.Stats.Sslv3AlertCertificateExpired,
			"sslv3_alert_certificate_revoked":         ret.DtSlbSslErrorStats.Stats.Sslv3AlertCertificateRevoked,
			"sslv3_alert_certificate_unknown":         ret.DtSlbSslErrorStats.Stats.Sslv3AlertCertificateUnknown,
			"sslv3_alert_decompression_failure":       ret.DtSlbSslErrorStats.Stats.Sslv3AlertDecompressionFailure,
			"sslv3_alert_handshake_failure":           ret.DtSlbSslErrorStats.Stats.Sslv3AlertHandshakeFailure,
			"sslv3_alert_illegal_parameter":           ret.DtSlbSslErrorStats.Stats.Sslv3AlertIllegalParameter,
			"sslv3_alert_no_certificate":              ret.DtSlbSslErrorStats.Stats.Sslv3AlertNoCertificate,
			"sslv3_alert_peer_error_cert":             ret.DtSlbSslErrorStats.Stats.Sslv3AlertPeerErrorCert,
			"sslv3_alert_peer_error_no_cert":          ret.DtSlbSslErrorStats.Stats.Sslv3AlertPeerErrorNoCert,
			"sslv3_alert_peer_error_no_cipher":        ret.DtSlbSslErrorStats.Stats.Sslv3AlertPeerErrorNoCipher,
			"sslv3_alert_peer_error_unsupp_cert_type": ret.DtSlbSslErrorStats.Stats.Sslv3AlertPeerErrorUnsuppCertType,
			"sslv3_alert_unexpected_msg":              ret.DtSlbSslErrorStats.Stats.Sslv3AlertUnexpectedMsg,
			"sslv3_alert_unknown_remote_err_type":     ret.DtSlbSslErrorStats.Stats.Sslv3AlertUnknownRemoteErrType,
			"sslv3_alert_unspported_cert":             ret.DtSlbSslErrorStats.Stats.Sslv3AlertUnspportedCert,
			"ssl_ctx_has_no_default_ssl_version":      ret.DtSlbSslErrorStats.Stats.SslCtxHasNoDefaultSslVersion,
			"ssl_handshake_failure":                   ret.DtSlbSslErrorStats.Stats.SslHandshakeFailure,
			"ssl_library_has_no_ciphers":              ret.DtSlbSslErrorStats.Stats.SslLibraryHasNoCiphers,
			"ssl_session_id_callback_failed":          ret.DtSlbSslErrorStats.Stats.SslSessionIdCallbackFailed,
			"ssl_session_id_conflict":                 ret.DtSlbSslErrorStats.Stats.SslSessionIdConflict,
			"ssl_session_id_context_too_long":         ret.DtSlbSslErrorStats.Stats.SslSessionIdContextTooLong,
			"ssl_session_id_has_bad_length":           ret.DtSlbSslErrorStats.Stats.SslSessionIdHasBadLength,
			"ssl_session_id_is_different":             ret.DtSlbSslErrorStats.Stats.SslSessionIdIsDifferent,
			"tlsv1_alert_access_denied":               ret.DtSlbSslErrorStats.Stats.Tlsv1AlertAccessDenied,
			"tlsv1_alert_decode_error":                ret.DtSlbSslErrorStats.Stats.Tlsv1AlertDecodeError,
			"tlsv1_alert_decryption_failed":           ret.DtSlbSslErrorStats.Stats.Tlsv1AlertDecryptionFailed,
			"tlsv1_alert_decrypt_error":               ret.DtSlbSslErrorStats.Stats.Tlsv1AlertDecryptError,
			"tlsv1_alert_export_restriction":          ret.DtSlbSslErrorStats.Stats.Tlsv1AlertExportRestriction,
			"tlsv1_alert_insufficient_security":       ret.DtSlbSslErrorStats.Stats.Tlsv1AlertInsufficientSecurity,
			"tlsv1_alert_internal_error":              ret.DtSlbSslErrorStats.Stats.Tlsv1AlertInternalError,
			"tlsv1_alert_no_renegotiation":            ret.DtSlbSslErrorStats.Stats.Tlsv1AlertNoRenegotiation,
			"tlsv1_alert_protocol_version":            ret.DtSlbSslErrorStats.Stats.Tlsv1AlertProtocolVersion,
			"tlsv1_alert_record_overflow":             ret.DtSlbSslErrorStats.Stats.Tlsv1AlertRecordOverflow,
			"tlsv1_alert_unknown_ca":                  ret.DtSlbSslErrorStats.Stats.Tlsv1AlertUnknownCa,
			"tlsv1_alert_user_cancelled":              ret.DtSlbSslErrorStats.Stats.Tlsv1AlertUserCancelled,
			"tls_client_cert_req_with_anon_cipher":    ret.DtSlbSslErrorStats.Stats.TlsClientCertReqWithAnonCipher,
			"tls_peer_did_not_respond_with_cert_list": ret.DtSlbSslErrorStats.Stats.TlsPeerDidNotRespondWithCertList,
			"tls_rsa_encrypted_value_length_is_wrong": ret.DtSlbSslErrorStats.Stats.TlsRsaEncryptedValueLengthIsWrong,
			"tried_to_use_unsupported_cipher":         ret.DtSlbSslErrorStats.Stats.TriedToUseUnsupportedCipher,
			"unable_to_decode_dh_certs":               ret.DtSlbSslErrorStats.Stats.UnableToDecodeDhCerts,
			"unable_to_extract_public_key":            ret.DtSlbSslErrorStats.Stats.UnableToExtractPublicKey,
			"unable_to_find_dh_parameters":            ret.DtSlbSslErrorStats.Stats.UnableToFindDhParameters,
			"unable_to_find_public_key_parameters":    ret.DtSlbSslErrorStats.Stats.UnableToFindPublicKeyParameters,
			"unable_to_find_ssl_method":               ret.DtSlbSslErrorStats.Stats.UnableToFindSslMethod,
			"unable_to_load_ssl2_md5_routines":        ret.DtSlbSslErrorStats.Stats.UnableToLoadSsl2Md5Routines,
			"unable_to_load_ssl3_md5_routines":        ret.DtSlbSslErrorStats.Stats.UnableToLoadSsl3Md5Routines,
			"unable_to_load_ssl3_sha1_routines":       ret.DtSlbSslErrorStats.Stats.UnableToLoadSsl3Sha1Routines,
			"unexpected_message":                      ret.DtSlbSslErrorStats.Stats.UnexpectedMessage,
			"unexpected_record":                       ret.DtSlbSslErrorStats.Stats.UnexpectedRecord,
			"uninitialized":                           ret.DtSlbSslErrorStats.Stats.Uninitialized,
			"unknown_alert_type":                      ret.DtSlbSslErrorStats.Stats.UnknownAlertType,
			"unknown_certificate_type":                ret.DtSlbSslErrorStats.Stats.UnknownCertificateType,
			"unknown_cipher_returned":                 ret.DtSlbSslErrorStats.Stats.UnknownCipherReturned,
			"unknown_cipher_type":                     ret.DtSlbSslErrorStats.Stats.UnknownCipherType,
			"unknown_key_exchange_type":               ret.DtSlbSslErrorStats.Stats.UnknownKeyExchangeType,
			"unknown_pkey_type":                       ret.DtSlbSslErrorStats.Stats.UnknownPkeyType,
			"unknown_protocol":                        ret.DtSlbSslErrorStats.Stats.UnknownProtocol,
			"unknown_remote_error_type":               ret.DtSlbSslErrorStats.Stats.UnknownRemoteErrorType,
			"unknown_ssl_version":                     ret.DtSlbSslErrorStats.Stats.UnknownSslVersion,
			"unknown_state":                           ret.DtSlbSslErrorStats.Stats.UnknownState,
			"unsupported_cipher":                      ret.DtSlbSslErrorStats.Stats.UnsupportedCipher,
			"unsupported_compression_algorithm":       ret.DtSlbSslErrorStats.Stats.UnsupportedCompressionAlgorithm,
			"unsupported_option":                      ret.DtSlbSslErrorStats.Stats.UnsupportedOption,
			"unsupported_protocol":                    ret.DtSlbSslErrorStats.Stats.UnsupportedProtocol,
			"unsupported_ssl_version":                 ret.DtSlbSslErrorStats.Stats.UnsupportedSslVersion,
			"unsupported_status_type":                 ret.DtSlbSslErrorStats.Stats.UnsupportedStatusType,
			"write_bio_not_set":                       ret.DtSlbSslErrorStats.Stats.WriteBioNotSet,
			"wrong_cipher_returned":                   ret.DtSlbSslErrorStats.Stats.WrongCipherReturned,
			"wrong_message_type":                      ret.DtSlbSslErrorStats.Stats.WrongMessageType,
			"wrong_counter_of_key_bits":               ret.DtSlbSslErrorStats.Stats.WrongCounterOfKeyBits,
			"wrong_signature_length":                  ret.DtSlbSslErrorStats.Stats.WrongSignatureLength,
			"wrong_signature_size":                    ret.DtSlbSslErrorStats.Stats.WrongSignatureSize,
			"wrong_ssl_version":                       ret.DtSlbSslErrorStats.Stats.WrongSslVersion,
			"wrong_version_counter":                   ret.DtSlbSslErrorStats.Stats.WrongVersionCounter,
			"x509_lib":                                ret.DtSlbSslErrorStats.Stats.X509Lib,
			"x509_verification_setup_problems":        ret.DtSlbSslErrorStats.Stats.X509VerificationSetupProblems,
			"clienthello_tlsext":                      ret.DtSlbSslErrorStats.Stats.ClienthelloTlsext,
			"parse_tlsext":                            ret.DtSlbSslErrorStats.Stats.ParseTlsext,
			"serverhello_tlsext":                      ret.DtSlbSslErrorStats.Stats.ServerhelloTlsext,
			"ssl3_ext_invalid_servername":             ret.DtSlbSslErrorStats.Stats.Ssl3ExtInvalidServername,
			"ssl3_ext_invalid_servername_type":        ret.DtSlbSslErrorStats.Stats.Ssl3ExtInvalidServernameType,
			"multiple_sgc_restarts":                   ret.DtSlbSslErrorStats.Stats.MultipleSgcRestarts,
			"tls_invalid_ecpointformat_list":          ret.DtSlbSslErrorStats.Stats.TlsInvalidEcpointformatList,
			"bad_ecc_cert":                            ret.DtSlbSslErrorStats.Stats.BadEccCert,
			"bad_ecdsa_sig":                           ret.DtSlbSslErrorStats.Stats.BadEcdsaSig,
			"bad_ecpoint":                             ret.DtSlbSslErrorStats.Stats.BadEcpoint,
			"cookie_mismatch":                         ret.DtSlbSslErrorStats.Stats.CookieMismatch,
			"unsupported_elliptic_curve":              ret.DtSlbSslErrorStats.Stats.UnsupportedEllipticCurve,
			"no_required_digest":                      ret.DtSlbSslErrorStats.Stats.NoRequiredDigest,
			"unsupported_digest_type":                 ret.DtSlbSslErrorStats.Stats.UnsupportedDigestType,
			"bad_handshake_length":                    ret.DtSlbSslErrorStats.Stats.BadHandshakeLength,
		},
	}
}

func getObjectSlbSslErrorStatsStats(d []interface{}) edpt.SlbSslErrorStatsStats {

	count1 := len(d)
	var ret edpt.SlbSslErrorStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AppDataInHandshake = in["app_data_in_handshake"].(int)
		ret.AttemptToReuseSessInDiffContext = in["attempt_to_reuse_sess_in_diff_context"].(int)
		ret.BadAlertRecord = in["bad_alert_record"].(int)
		ret.BadAuthenticationType = in["bad_authentication_type"].(int)
		ret.BadChangeCipherSpec = in["bad_change_cipher_spec"].(int)
		ret.BadChecksum = in["bad_checksum"].(int)
		ret.BadDataReturnedByCallback = in["bad_data_returned_by_callback"].(int)
		ret.BadDecompression = in["bad_decompression"].(int)
		ret.BadDhGLength = in["bad_dh_g_length"].(int)
		ret.BadDhPubKeyLength = in["bad_dh_pub_key_length"].(int)
		ret.BadDhPLength = in["bad_dh_p_length"].(int)
		ret.BadDigestLength = in["bad_digest_length"].(int)
		ret.BadDsaSignature = in["bad_dsa_signature"].(int)
		ret.BadHelloRequest = in["bad_hello_request"].(int)
		ret.BadLength = in["bad_length"].(int)
		ret.BadMacDecode = in["bad_mac_decode"].(int)
		ret.BadMessageType = in["bad_message_type"].(int)
		ret.BadPacketLength = in["bad_packet_length"].(int)
		ret.BadProtocolVersionCounter = in["bad_protocol_version_counter"].(int)
		ret.BadResponseArgument = in["bad_response_argument"].(int)
		ret.BadRsaDecrypt = in["bad_rsa_decrypt"].(int)
		ret.BadRsaEncrypt = in["bad_rsa_encrypt"].(int)
		ret.BadRsaELength = in["bad_rsa_e_length"].(int)
		ret.BadRsaModulusLength = in["bad_rsa_modulus_length"].(int)
		ret.BadRsaSignature = in["bad_rsa_signature"].(int)
		ret.BadSignature = in["bad_signature"].(int)
		ret.BadSslFiletype = in["bad_ssl_filetype"].(int)
		ret.BadSslSessionIdLength = in["bad_ssl_session_id_length"].(int)
		ret.BadState = in["bad_state"].(int)
		ret.BadWriteRetry = in["bad_write_retry"].(int)
		ret.BioNotSet = in["bio_not_set"].(int)
		ret.BlockCipherPadIsWrong = in["block_cipher_pad_is_wrong"].(int)
		ret.BnLib = in["bn_lib"].(int)
		ret.CaDnLengthMismatch = in["ca_dn_length_mismatch"].(int)
		ret.CaDnTooLong = in["ca_dn_too_long"].(int)
		ret.CcsReceivedEarly = in["ccs_received_early"].(int)
		ret.CertificateVerifyFailed = in["certificate_verify_failed"].(int)
		ret.CertLengthMismatch = in["cert_length_mismatch"].(int)
		ret.ChallengeIsDifferent = in["challenge_is_different"].(int)
		ret.CipherCodeWrongLength = in["cipher_code_wrong_length"].(int)
		ret.CipherOrHashUnavailable = in["cipher_or_hash_unavailable"].(int)
		ret.CipherTableSrcError = in["cipher_table_src_error"].(int)
		ret.CompressedLengthTooLong = in["compressed_length_too_long"].(int)
		ret.CompressionFailure = in["compression_failure"].(int)
		ret.CompressionLibraryError = in["compression_library_error"].(int)
		ret.ConnectionIdIsDifferent = in["connection_id_is_different"].(int)
		ret.ConnectionTypeNotSet = in["connection_type_not_set"].(int)
		ret.DataBetweenCcsAndFinished = in["data_between_ccs_and_finished"].(int)
		ret.DataLengthTooLong = in["data_length_too_long"].(int)
		ret.DecryptionFailed = in["decryption_failed"].(int)
		ret.DecryptionFailedOrBadRecordMac = in["decryption_failed_or_bad_record_mac"].(int)
		ret.DhPublicValueLengthIsWrong = in["dh_public_value_length_is_wrong"].(int)
		ret.DigestCheckFailed = in["digest_check_failed"].(int)
		ret.EncryptedLengthTooLong = in["encrypted_length_too_long"].(int)
		ret.ErrorGeneratingTmpRsaKey = in["error_generating_tmp_rsa_key"].(int)
		ret.ErrorInReceivedCipherList = in["error_in_received_cipher_list"].(int)
		ret.ExcessiveMessageSize = in["excessive_message_size"].(int)
		ret.ExtraDataInMessage = in["extra_data_in_message"].(int)
		ret.GotAFinBeforeACcs = in["got_a_fin_before_a_ccs"].(int)
		ret.HttpsProxyRequest = in["https_proxy_request"].(int)
		ret.HttpRequest = in["http_request"].(int)
		ret.IllegalPadding = in["illegal_padding"].(int)
		ret.InappropriateFallback = in["inappropriate_fallback"].(int)
		ret.InvalidChallengeLength = in["invalid_challenge_length"].(int)
		ret.InvalidCommand = in["invalid_command"].(int)
		ret.InvalidPurpose = in["invalid_purpose"].(int)
		ret.InvalidStatusResponse = in["invalid_status_response"].(int)
		ret.InvalidTrust = in["invalid_trust"].(int)
		ret.KeyArgTooLong = in["key_arg_too_long"].(int)
		ret.Krb5 = in["krb5"].(int)
		ret.Krb5ClientCcPrincipal = in["krb5_client_cc_principal"].(int)
		ret.Krb5ClientGetCred = in["krb5_client_get_cred"].(int)
		ret.Krb5ClientInit = in["krb5_client_init"].(int)
		ret.Krb5ClientMkReq = in["krb5_client_mk_req"].(int)
		ret.Krb5ServerBadTicket = in["krb5_server_bad_ticket"].(int)
		ret.Krb5ServerInit = in["krb5_server_init"].(int)
		ret.Krb5ServerRdReq = in["krb5_server_rd_req"].(int)
		ret.Krb5ServerTktExpired = in["krb5_server_tkt_expired"].(int)
		ret.Krb5ServerTktNotYetValid = in["krb5_server_tkt_not_yet_valid"].(int)
		ret.Krb5ServerTktSkew = in["krb5_server_tkt_skew"].(int)
		ret.LengthMismatch = in["length_mismatch"].(int)
		ret.LengthTooShort = in["length_too_short"].(int)
		ret.LibraryBug = in["library_bug"].(int)
		ret.LibraryHasNoCiphers = in["library_has_no_ciphers"].(int)
		ret.MastKeyTooLong = in["mast_key_too_long"].(int)
		ret.MessageTooLong = in["message_too_long"].(int)
		ret.MissingDhDsaCert = in["missing_dh_dsa_cert"].(int)
		ret.MissingDhKey = in["missing_dh_key"].(int)
		ret.MissingDhRsaCert = in["missing_dh_rsa_cert"].(int)
		ret.MissingDsaSigningCert = in["missing_dsa_signing_cert"].(int)
		ret.MissingExportTmpDhKey = in["missing_export_tmp_dh_key"].(int)
		ret.MissingExportTmpRsaKey = in["missing_export_tmp_rsa_key"].(int)
		ret.MissingRsaCertificate = in["missing_rsa_certificate"].(int)
		ret.MissingRsaEncryptingCert = in["missing_rsa_encrypting_cert"].(int)
		ret.MissingRsaSigningCert = in["missing_rsa_signing_cert"].(int)
		ret.MissingTmpDhKey = in["missing_tmp_dh_key"].(int)
		ret.MissingTmpRsaKey = in["missing_tmp_rsa_key"].(int)
		ret.MissingTmpRsaPkey = in["missing_tmp_rsa_pkey"].(int)
		ret.MissingVerifyMessage = in["missing_verify_message"].(int)
		ret.NonSslv2InitialPacket = in["non_sslv2_initial_packet"].(int)
		ret.NoCertificatesReturned = in["no_certificates_returned"].(int)
		ret.NoCertificateAssigned = in["no_certificate_assigned"].(int)
		ret.NoCertificateReturned = in["no_certificate_returned"].(int)
		ret.NoCertificateSet = in["no_certificate_set"].(int)
		ret.NoCertificateSpecified = in["no_certificate_specified"].(int)
		ret.NoCiphersAvailable = in["no_ciphers_available"].(int)
		ret.NoCiphersPassed = in["no_ciphers_passed"].(int)
		ret.NoCiphersSpecified = in["no_ciphers_specified"].(int)
		ret.NoCipherList = in["no_cipher_list"].(int)
		ret.NoCipherMatch = in["no_cipher_match"].(int)
		ret.NoClientCertReceived = in["no_client_cert_received"].(int)
		ret.NoCompressionSpecified = in["no_compression_specified"].(int)
		ret.NoMethodSpecified = in["no_method_specified"].(int)
		ret.NoPrivatekey = in["no_privatekey"].(int)
		ret.NoPrivateKeyAssigned = in["no_private_key_assigned"].(int)
		ret.NoProtocolsAvailable = in["no_protocols_available"].(int)
		ret.NoPublickey = in["no_publickey"].(int)
		ret.NoSharedCipher = in["no_shared_cipher"].(int)
		ret.NoVerifyCallback = in["no_verify_callback"].(int)
		ret.NullSslCtx = in["null_ssl_ctx"].(int)
		ret.NullSslMethodPassed = in["null_ssl_method_passed"].(int)
		ret.OldSessionCipherNotReturned = in["old_session_cipher_not_returned"].(int)
		ret.PacketLengthTooLong = in["packet_length_too_long"].(int)
		ret.PathTooLong = in["path_too_long"].(int)
		ret.PeerDidNotReturnACertificate = in["peer_did_not_return_a_certificate"].(int)
		ret.PeerError = in["peer_error"].(int)
		ret.PeerErrorCertificate = in["peer_error_certificate"].(int)
		ret.PeerErrorNoCertificate = in["peer_error_no_certificate"].(int)
		ret.PeerErrorNoCipher = in["peer_error_no_cipher"].(int)
		ret.PeerErrorUnsupportedCertificateType = in["peer_error_unsupported_certificate_type"].(int)
		ret.PreMacLengthTooLong = in["pre_mac_length_too_long"].(int)
		ret.ProblemsMappingCipherFunctions = in["problems_mapping_cipher_functions"].(int)
		ret.ProtocolIsShutdown = in["protocol_is_shutdown"].(int)
		ret.PublicKeyEncryptError = in["public_key_encrypt_error"].(int)
		ret.PublicKeyIsNotRsa = in["public_key_is_not_rsa"].(int)
		ret.PublicKeyNotRsa = in["public_key_not_rsa"].(int)
		ret.ReadBioNotSet = in["read_bio_not_set"].(int)
		ret.ReadWrongPacketType = in["read_wrong_packet_type"].(int)
		ret.RecordLengthMismatch = in["record_length_mismatch"].(int)
		ret.RecordTooLarge = in["record_too_large"].(int)
		ret.RecordTooSmall = in["record_too_small"].(int)
		ret.RequiredCipherMissing = in["required_cipher_missing"].(int)
		ret.ReuseCertLengthNotZero = in["reuse_cert_length_not_zero"].(int)
		ret.ReuseCertTypeNotZero = in["reuse_cert_type_not_zero"].(int)
		ret.ReuseCipherListNotZero = in["reuse_cipher_list_not_zero"].(int)
		ret.ScsvReceivedWhenRenegotiating = in["scsv_received_when_renegotiating"].(int)
		ret.SessionIdContextUninitialized = in["session_id_context_uninitialized"].(int)
		ret.ShortRead = in["short_read"].(int)
		ret.SignatureForNonSigningCertificate = in["signature_for_non_signing_certificate"].(int)
		ret.Ssl23DoingSessionIdReuse = in["ssl23_doing_session_id_reuse"].(int)
		ret.Ssl2ConnectionIdTooLong = in["ssl2_connection_id_too_long"].(int)
		ret.Ssl3SessionIdTooLong = in["ssl3_session_id_too_long"].(int)
		ret.Ssl3SessionIdTooShort = in["ssl3_session_id_too_short"].(int)
		ret.Sslv3AlertBadCertificate = in["sslv3_alert_bad_certificate"].(int)
		ret.Sslv3AlertBadRecordMac = in["sslv3_alert_bad_record_mac"].(int)
		ret.Sslv3AlertCertificateExpired = in["sslv3_alert_certificate_expired"].(int)
		ret.Sslv3AlertCertificateRevoked = in["sslv3_alert_certificate_revoked"].(int)
		ret.Sslv3AlertCertificateUnknown = in["sslv3_alert_certificate_unknown"].(int)
		ret.Sslv3AlertDecompressionFailure = in["sslv3_alert_decompression_failure"].(int)
		ret.Sslv3AlertHandshakeFailure = in["sslv3_alert_handshake_failure"].(int)
		ret.Sslv3AlertIllegalParameter = in["sslv3_alert_illegal_parameter"].(int)
		ret.Sslv3AlertNoCertificate = in["sslv3_alert_no_certificate"].(int)
		ret.Sslv3AlertPeerErrorCert = in["sslv3_alert_peer_error_cert"].(int)
		ret.Sslv3AlertPeerErrorNoCert = in["sslv3_alert_peer_error_no_cert"].(int)
		ret.Sslv3AlertPeerErrorNoCipher = in["sslv3_alert_peer_error_no_cipher"].(int)
		ret.Sslv3AlertPeerErrorUnsuppCertType = in["sslv3_alert_peer_error_unsupp_cert_type"].(int)
		ret.Sslv3AlertUnexpectedMsg = in["sslv3_alert_unexpected_msg"].(int)
		ret.Sslv3AlertUnknownRemoteErrType = in["sslv3_alert_unknown_remote_err_type"].(int)
		ret.Sslv3AlertUnspportedCert = in["sslv3_alert_unspported_cert"].(int)
		ret.SslCtxHasNoDefaultSslVersion = in["ssl_ctx_has_no_default_ssl_version"].(int)
		ret.SslHandshakeFailure = in["ssl_handshake_failure"].(int)
		ret.SslLibraryHasNoCiphers = in["ssl_library_has_no_ciphers"].(int)
		ret.SslSessionIdCallbackFailed = in["ssl_session_id_callback_failed"].(int)
		ret.SslSessionIdConflict = in["ssl_session_id_conflict"].(int)
		ret.SslSessionIdContextTooLong = in["ssl_session_id_context_too_long"].(int)
		ret.SslSessionIdHasBadLength = in["ssl_session_id_has_bad_length"].(int)
		ret.SslSessionIdIsDifferent = in["ssl_session_id_is_different"].(int)
		ret.Tlsv1AlertAccessDenied = in["tlsv1_alert_access_denied"].(int)
		ret.Tlsv1AlertDecodeError = in["tlsv1_alert_decode_error"].(int)
		ret.Tlsv1AlertDecryptionFailed = in["tlsv1_alert_decryption_failed"].(int)
		ret.Tlsv1AlertDecryptError = in["tlsv1_alert_decrypt_error"].(int)
		ret.Tlsv1AlertExportRestriction = in["tlsv1_alert_export_restriction"].(int)
		ret.Tlsv1AlertInsufficientSecurity = in["tlsv1_alert_insufficient_security"].(int)
		ret.Tlsv1AlertInternalError = in["tlsv1_alert_internal_error"].(int)
		ret.Tlsv1AlertNoRenegotiation = in["tlsv1_alert_no_renegotiation"].(int)
		ret.Tlsv1AlertProtocolVersion = in["tlsv1_alert_protocol_version"].(int)
		ret.Tlsv1AlertRecordOverflow = in["tlsv1_alert_record_overflow"].(int)
		ret.Tlsv1AlertUnknownCa = in["tlsv1_alert_unknown_ca"].(int)
		ret.Tlsv1AlertUserCancelled = in["tlsv1_alert_user_cancelled"].(int)
		ret.TlsClientCertReqWithAnonCipher = in["tls_client_cert_req_with_anon_cipher"].(int)
		ret.TlsPeerDidNotRespondWithCertList = in["tls_peer_did_not_respond_with_cert_list"].(int)
		ret.TlsRsaEncryptedValueLengthIsWrong = in["tls_rsa_encrypted_value_length_is_wrong"].(int)
		ret.TriedToUseUnsupportedCipher = in["tried_to_use_unsupported_cipher"].(int)
		ret.UnableToDecodeDhCerts = in["unable_to_decode_dh_certs"].(int)
		ret.UnableToExtractPublicKey = in["unable_to_extract_public_key"].(int)
		ret.UnableToFindDhParameters = in["unable_to_find_dh_parameters"].(int)
		ret.UnableToFindPublicKeyParameters = in["unable_to_find_public_key_parameters"].(int)
		ret.UnableToFindSslMethod = in["unable_to_find_ssl_method"].(int)
		ret.UnableToLoadSsl2Md5Routines = in["unable_to_load_ssl2_md5_routines"].(int)
		ret.UnableToLoadSsl3Md5Routines = in["unable_to_load_ssl3_md5_routines"].(int)
		ret.UnableToLoadSsl3Sha1Routines = in["unable_to_load_ssl3_sha1_routines"].(int)
		ret.UnexpectedMessage = in["unexpected_message"].(int)
		ret.UnexpectedRecord = in["unexpected_record"].(int)
		ret.Uninitialized = in["uninitialized"].(int)
		ret.UnknownAlertType = in["unknown_alert_type"].(int)
		ret.UnknownCertificateType = in["unknown_certificate_type"].(int)
		ret.UnknownCipherReturned = in["unknown_cipher_returned"].(int)
		ret.UnknownCipherType = in["unknown_cipher_type"].(int)
		ret.UnknownKeyExchangeType = in["unknown_key_exchange_type"].(int)
		ret.UnknownPkeyType = in["unknown_pkey_type"].(int)
		ret.UnknownProtocol = in["unknown_protocol"].(int)
		ret.UnknownRemoteErrorType = in["unknown_remote_error_type"].(int)
		ret.UnknownSslVersion = in["unknown_ssl_version"].(int)
		ret.UnknownState = in["unknown_state"].(int)
		ret.UnsupportedCipher = in["unsupported_cipher"].(int)
		ret.UnsupportedCompressionAlgorithm = in["unsupported_compression_algorithm"].(int)
		ret.UnsupportedOption = in["unsupported_option"].(int)
		ret.UnsupportedProtocol = in["unsupported_protocol"].(int)
		ret.UnsupportedSslVersion = in["unsupported_ssl_version"].(int)
		ret.UnsupportedStatusType = in["unsupported_status_type"].(int)
		ret.WriteBioNotSet = in["write_bio_not_set"].(int)
		ret.WrongCipherReturned = in["wrong_cipher_returned"].(int)
		ret.WrongMessageType = in["wrong_message_type"].(int)
		ret.WrongCounterOfKeyBits = in["wrong_counter_of_key_bits"].(int)
		ret.WrongSignatureLength = in["wrong_signature_length"].(int)
		ret.WrongSignatureSize = in["wrong_signature_size"].(int)
		ret.WrongSslVersion = in["wrong_ssl_version"].(int)
		ret.WrongVersionCounter = in["wrong_version_counter"].(int)
		ret.X509Lib = in["x509_lib"].(int)
		ret.X509VerificationSetupProblems = in["x509_verification_setup_problems"].(int)
		ret.ClienthelloTlsext = in["clienthello_tlsext"].(int)
		ret.ParseTlsext = in["parse_tlsext"].(int)
		ret.ServerhelloTlsext = in["serverhello_tlsext"].(int)
		ret.Ssl3ExtInvalidServername = in["ssl3_ext_invalid_servername"].(int)
		ret.Ssl3ExtInvalidServernameType = in["ssl3_ext_invalid_servername_type"].(int)
		ret.MultipleSgcRestarts = in["multiple_sgc_restarts"].(int)
		ret.TlsInvalidEcpointformatList = in["tls_invalid_ecpointformat_list"].(int)
		ret.BadEccCert = in["bad_ecc_cert"].(int)
		ret.BadEcdsaSig = in["bad_ecdsa_sig"].(int)
		ret.BadEcpoint = in["bad_ecpoint"].(int)
		ret.CookieMismatch = in["cookie_mismatch"].(int)
		ret.UnsupportedEllipticCurve = in["unsupported_elliptic_curve"].(int)
		ret.NoRequiredDigest = in["no_required_digest"].(int)
		ret.UnsupportedDigestType = in["unsupported_digest_type"].(int)
		ret.BadHandshakeLength = in["bad_handshake_length"].(int)
	}
	return ret
}

func dataToEndpointSlbSslErrorStats(d *schema.ResourceData) edpt.SlbSslErrorStats {
	var ret edpt.SlbSslErrorStats

	ret.Stats = getObjectSlbSslErrorStatsStats(d.Get("stats").([]interface{}))
	return ret
}
