package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslErrorOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_error_oper`: Operational Status for the object ssl-error\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslErrorOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
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
						"bad_protocol_version_number": {
							Type: schema.TypeInt, Optional: true, Description: "bad protocol version number",
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
						"wrong_number_of_key_bits": {
							Type: schema.TypeInt, Optional: true, Description: "wrong number of key bits",
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
						"wrong_version_number": {
							Type: schema.TypeInt, Optional: true, Description: "wrong version number",
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

func resourceSlbSslErrorOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslErrorOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslErrorOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslErrorOperOper := setObjectSlbSslErrorOperOper(res)
		d.Set("oper", SlbSslErrorOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslErrorOperOper(ret edpt.DataSlbSslErrorOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"app_data_in_handshake":                   ret.DtSlbSslErrorOper.Oper.AppDataInHandshake,
			"attempt_to_reuse_sess_in_diff_context":   ret.DtSlbSslErrorOper.Oper.AttemptToReuseSessInDiffContext,
			"bad_alert_record":                        ret.DtSlbSslErrorOper.Oper.BadAlertRecord,
			"bad_authentication_type":                 ret.DtSlbSslErrorOper.Oper.BadAuthenticationType,
			"bad_change_cipher_spec":                  ret.DtSlbSslErrorOper.Oper.BadChangeCipherSpec,
			"bad_checksum":                            ret.DtSlbSslErrorOper.Oper.BadChecksum,
			"bad_data_returned_by_callback":           ret.DtSlbSslErrorOper.Oper.BadDataReturnedByCallback,
			"bad_decompression":                       ret.DtSlbSslErrorOper.Oper.BadDecompression,
			"bad_dh_g_length":                         ret.DtSlbSslErrorOper.Oper.BadDhGLength,
			"bad_dh_pub_key_length":                   ret.DtSlbSslErrorOper.Oper.BadDhPubKeyLength,
			"bad_dh_p_length":                         ret.DtSlbSslErrorOper.Oper.BadDhPLength,
			"bad_digest_length":                       ret.DtSlbSslErrorOper.Oper.BadDigestLength,
			"bad_dsa_signature":                       ret.DtSlbSslErrorOper.Oper.BadDsaSignature,
			"bad_hello_request":                       ret.DtSlbSslErrorOper.Oper.BadHelloRequest,
			"bad_length":                              ret.DtSlbSslErrorOper.Oper.BadLength,
			"bad_mac_decode":                          ret.DtSlbSslErrorOper.Oper.BadMacDecode,
			"bad_message_type":                        ret.DtSlbSslErrorOper.Oper.BadMessageType,
			"bad_packet_length":                       ret.DtSlbSslErrorOper.Oper.BadPacketLength,
			"bad_protocol_version_number":             ret.DtSlbSslErrorOper.Oper.BadProtocolVersionNumber,
			"bad_response_argument":                   ret.DtSlbSslErrorOper.Oper.BadResponseArgument,
			"bad_rsa_decrypt":                         ret.DtSlbSslErrorOper.Oper.BadRsaDecrypt,
			"bad_rsa_encrypt":                         ret.DtSlbSslErrorOper.Oper.BadRsaEncrypt,
			"bad_rsa_e_length":                        ret.DtSlbSslErrorOper.Oper.BadRsaELength,
			"bad_rsa_modulus_length":                  ret.DtSlbSslErrorOper.Oper.BadRsaModulusLength,
			"bad_rsa_signature":                       ret.DtSlbSslErrorOper.Oper.BadRsaSignature,
			"bad_signature":                           ret.DtSlbSslErrorOper.Oper.BadSignature,
			"bad_ssl_filetype":                        ret.DtSlbSslErrorOper.Oper.BadSslFiletype,
			"bad_ssl_session_id_length":               ret.DtSlbSslErrorOper.Oper.BadSslSessionIdLength,
			"bad_state":                               ret.DtSlbSslErrorOper.Oper.BadState,
			"bad_write_retry":                         ret.DtSlbSslErrorOper.Oper.BadWriteRetry,
			"bio_not_set":                             ret.DtSlbSslErrorOper.Oper.BioNotSet,
			"block_cipher_pad_is_wrong":               ret.DtSlbSslErrorOper.Oper.BlockCipherPadIsWrong,
			"bn_lib":                                  ret.DtSlbSslErrorOper.Oper.BnLib,
			"ca_dn_length_mismatch":                   ret.DtSlbSslErrorOper.Oper.CaDnLengthMismatch,
			"ca_dn_too_long":                          ret.DtSlbSslErrorOper.Oper.CaDnTooLong,
			"ccs_received_early":                      ret.DtSlbSslErrorOper.Oper.CcsReceivedEarly,
			"certificate_verify_failed":               ret.DtSlbSslErrorOper.Oper.CertificateVerifyFailed,
			"cert_length_mismatch":                    ret.DtSlbSslErrorOper.Oper.CertLengthMismatch,
			"challenge_is_different":                  ret.DtSlbSslErrorOper.Oper.ChallengeIsDifferent,
			"cipher_code_wrong_length":                ret.DtSlbSslErrorOper.Oper.CipherCodeWrongLength,
			"cipher_or_hash_unavailable":              ret.DtSlbSslErrorOper.Oper.CipherOrHashUnavailable,
			"cipher_table_src_error":                  ret.DtSlbSslErrorOper.Oper.CipherTableSrcError,
			"compressed_length_too_long":              ret.DtSlbSslErrorOper.Oper.CompressedLengthTooLong,
			"compression_failure":                     ret.DtSlbSslErrorOper.Oper.CompressionFailure,
			"compression_library_error":               ret.DtSlbSslErrorOper.Oper.CompressionLibraryError,
			"connection_id_is_different":              ret.DtSlbSslErrorOper.Oper.ConnectionIdIsDifferent,
			"connection_type_not_set":                 ret.DtSlbSslErrorOper.Oper.ConnectionTypeNotSet,
			"data_between_ccs_and_finished":           ret.DtSlbSslErrorOper.Oper.DataBetweenCcsAndFinished,
			"data_length_too_long":                    ret.DtSlbSslErrorOper.Oper.DataLengthTooLong,
			"decryption_failed":                       ret.DtSlbSslErrorOper.Oper.DecryptionFailed,
			"decryption_failed_or_bad_record_mac":     ret.DtSlbSslErrorOper.Oper.DecryptionFailedOrBadRecordMac,
			"dh_public_value_length_is_wrong":         ret.DtSlbSslErrorOper.Oper.DhPublicValueLengthIsWrong,
			"digest_check_failed":                     ret.DtSlbSslErrorOper.Oper.DigestCheckFailed,
			"encrypted_length_too_long":               ret.DtSlbSslErrorOper.Oper.EncryptedLengthTooLong,
			"error_generating_tmp_rsa_key":            ret.DtSlbSslErrorOper.Oper.ErrorGeneratingTmpRsaKey,
			"error_in_received_cipher_list":           ret.DtSlbSslErrorOper.Oper.ErrorInReceivedCipherList,
			"excessive_message_size":                  ret.DtSlbSslErrorOper.Oper.ExcessiveMessageSize,
			"extra_data_in_message":                   ret.DtSlbSslErrorOper.Oper.ExtraDataInMessage,
			"got_a_fin_before_a_ccs":                  ret.DtSlbSslErrorOper.Oper.GotAFinBeforeACcs,
			"https_proxy_request":                     ret.DtSlbSslErrorOper.Oper.HttpsProxyRequest,
			"http_request":                            ret.DtSlbSslErrorOper.Oper.HttpRequest,
			"illegal_padding":                         ret.DtSlbSslErrorOper.Oper.IllegalPadding,
			"inappropriate_fallback":                  ret.DtSlbSslErrorOper.Oper.InappropriateFallback,
			"invalid_challenge_length":                ret.DtSlbSslErrorOper.Oper.InvalidChallengeLength,
			"invalid_command":                         ret.DtSlbSslErrorOper.Oper.InvalidCommand,
			"invalid_purpose":                         ret.DtSlbSslErrorOper.Oper.InvalidPurpose,
			"invalid_status_response":                 ret.DtSlbSslErrorOper.Oper.InvalidStatusResponse,
			"invalid_trust":                           ret.DtSlbSslErrorOper.Oper.InvalidTrust,
			"key_arg_too_long":                        ret.DtSlbSslErrorOper.Oper.KeyArgTooLong,
			"krb5":                                    ret.DtSlbSslErrorOper.Oper.Krb5,
			"krb5_client_cc_principal":                ret.DtSlbSslErrorOper.Oper.Krb5ClientCcPrincipal,
			"krb5_client_get_cred":                    ret.DtSlbSslErrorOper.Oper.Krb5ClientGetCred,
			"krb5_client_init":                        ret.DtSlbSslErrorOper.Oper.Krb5ClientInit,
			"krb5_client_mk_req":                      ret.DtSlbSslErrorOper.Oper.Krb5ClientMkReq,
			"krb5_server_bad_ticket":                  ret.DtSlbSslErrorOper.Oper.Krb5ServerBadTicket,
			"krb5_server_init":                        ret.DtSlbSslErrorOper.Oper.Krb5ServerInit,
			"krb5_server_rd_req":                      ret.DtSlbSslErrorOper.Oper.Krb5ServerRdReq,
			"krb5_server_tkt_expired":                 ret.DtSlbSslErrorOper.Oper.Krb5ServerTktExpired,
			"krb5_server_tkt_not_yet_valid":           ret.DtSlbSslErrorOper.Oper.Krb5ServerTktNotYetValid,
			"krb5_server_tkt_skew":                    ret.DtSlbSslErrorOper.Oper.Krb5ServerTktSkew,
			"length_mismatch":                         ret.DtSlbSslErrorOper.Oper.LengthMismatch,
			"length_too_short":                        ret.DtSlbSslErrorOper.Oper.LengthTooShort,
			"library_bug":                             ret.DtSlbSslErrorOper.Oper.LibraryBug,
			"library_has_no_ciphers":                  ret.DtSlbSslErrorOper.Oper.LibraryHasNoCiphers,
			"mast_key_too_long":                       ret.DtSlbSslErrorOper.Oper.MastKeyTooLong,
			"message_too_long":                        ret.DtSlbSslErrorOper.Oper.MessageTooLong,
			"missing_dh_dsa_cert":                     ret.DtSlbSslErrorOper.Oper.MissingDhDsaCert,
			"missing_dh_key":                          ret.DtSlbSslErrorOper.Oper.MissingDhKey,
			"missing_dh_rsa_cert":                     ret.DtSlbSslErrorOper.Oper.MissingDhRsaCert,
			"missing_dsa_signing_cert":                ret.DtSlbSslErrorOper.Oper.MissingDsaSigningCert,
			"missing_export_tmp_dh_key":               ret.DtSlbSslErrorOper.Oper.MissingExportTmpDhKey,
			"missing_export_tmp_rsa_key":              ret.DtSlbSslErrorOper.Oper.MissingExportTmpRsaKey,
			"missing_rsa_certificate":                 ret.DtSlbSslErrorOper.Oper.MissingRsaCertificate,
			"missing_rsa_encrypting_cert":             ret.DtSlbSslErrorOper.Oper.MissingRsaEncryptingCert,
			"missing_rsa_signing_cert":                ret.DtSlbSslErrorOper.Oper.MissingRsaSigningCert,
			"missing_tmp_dh_key":                      ret.DtSlbSslErrorOper.Oper.MissingTmpDhKey,
			"missing_tmp_rsa_key":                     ret.DtSlbSslErrorOper.Oper.MissingTmpRsaKey,
			"missing_tmp_rsa_pkey":                    ret.DtSlbSslErrorOper.Oper.MissingTmpRsaPkey,
			"missing_verify_message":                  ret.DtSlbSslErrorOper.Oper.MissingVerifyMessage,
			"non_sslv2_initial_packet":                ret.DtSlbSslErrorOper.Oper.NonSslv2InitialPacket,
			"no_certificates_returned":                ret.DtSlbSslErrorOper.Oper.NoCertificatesReturned,
			"no_certificate_assigned":                 ret.DtSlbSslErrorOper.Oper.NoCertificateAssigned,
			"no_certificate_returned":                 ret.DtSlbSslErrorOper.Oper.NoCertificateReturned,
			"no_certificate_set":                      ret.DtSlbSslErrorOper.Oper.NoCertificateSet,
			"no_certificate_specified":                ret.DtSlbSslErrorOper.Oper.NoCertificateSpecified,
			"no_ciphers_available":                    ret.DtSlbSslErrorOper.Oper.NoCiphersAvailable,
			"no_ciphers_passed":                       ret.DtSlbSslErrorOper.Oper.NoCiphersPassed,
			"no_ciphers_specified":                    ret.DtSlbSslErrorOper.Oper.NoCiphersSpecified,
			"no_cipher_list":                          ret.DtSlbSslErrorOper.Oper.NoCipherList,
			"no_cipher_match":                         ret.DtSlbSslErrorOper.Oper.NoCipherMatch,
			"no_client_cert_received":                 ret.DtSlbSslErrorOper.Oper.NoClientCertReceived,
			"no_compression_specified":                ret.DtSlbSslErrorOper.Oper.NoCompressionSpecified,
			"no_method_specified":                     ret.DtSlbSslErrorOper.Oper.NoMethodSpecified,
			"no_privatekey":                           ret.DtSlbSslErrorOper.Oper.NoPrivatekey,
			"no_private_key_assigned":                 ret.DtSlbSslErrorOper.Oper.NoPrivateKeyAssigned,
			"no_protocols_available":                  ret.DtSlbSslErrorOper.Oper.NoProtocolsAvailable,
			"no_publickey":                            ret.DtSlbSslErrorOper.Oper.NoPublickey,
			"no_shared_cipher":                        ret.DtSlbSslErrorOper.Oper.NoSharedCipher,
			"no_verify_callback":                      ret.DtSlbSslErrorOper.Oper.NoVerifyCallback,
			"null_ssl_ctx":                            ret.DtSlbSslErrorOper.Oper.NullSslCtx,
			"null_ssl_method_passed":                  ret.DtSlbSslErrorOper.Oper.NullSslMethodPassed,
			"old_session_cipher_not_returned":         ret.DtSlbSslErrorOper.Oper.OldSessionCipherNotReturned,
			"packet_length_too_long":                  ret.DtSlbSslErrorOper.Oper.PacketLengthTooLong,
			"path_too_long":                           ret.DtSlbSslErrorOper.Oper.PathTooLong,
			"peer_did_not_return_a_certificate":       ret.DtSlbSslErrorOper.Oper.PeerDidNotReturnACertificate,
			"peer_error":                              ret.DtSlbSslErrorOper.Oper.PeerError,
			"peer_error_certificate":                  ret.DtSlbSslErrorOper.Oper.PeerErrorCertificate,
			"peer_error_no_certificate":               ret.DtSlbSslErrorOper.Oper.PeerErrorNoCertificate,
			"peer_error_no_cipher":                    ret.DtSlbSslErrorOper.Oper.PeerErrorNoCipher,
			"peer_error_unsupported_certificate_type": ret.DtSlbSslErrorOper.Oper.PeerErrorUnsupportedCertificateType,
			"pre_mac_length_too_long":                 ret.DtSlbSslErrorOper.Oper.PreMacLengthTooLong,
			"problems_mapping_cipher_functions":       ret.DtSlbSslErrorOper.Oper.ProblemsMappingCipherFunctions,
			"protocol_is_shutdown":                    ret.DtSlbSslErrorOper.Oper.ProtocolIsShutdown,
			"public_key_encrypt_error":                ret.DtSlbSslErrorOper.Oper.PublicKeyEncryptError,
			"public_key_is_not_rsa":                   ret.DtSlbSslErrorOper.Oper.PublicKeyIsNotRsa,
			"public_key_not_rsa":                      ret.DtSlbSslErrorOper.Oper.PublicKeyNotRsa,
			"read_bio_not_set":                        ret.DtSlbSslErrorOper.Oper.ReadBioNotSet,
			"read_wrong_packet_type":                  ret.DtSlbSslErrorOper.Oper.ReadWrongPacketType,
			"record_length_mismatch":                  ret.DtSlbSslErrorOper.Oper.RecordLengthMismatch,
			"record_too_large":                        ret.DtSlbSslErrorOper.Oper.RecordTooLarge,
			"record_too_small":                        ret.DtSlbSslErrorOper.Oper.RecordTooSmall,
			"required_cipher_missing":                 ret.DtSlbSslErrorOper.Oper.RequiredCipherMissing,
			"reuse_cert_length_not_zero":              ret.DtSlbSslErrorOper.Oper.ReuseCertLengthNotZero,
			"reuse_cert_type_not_zero":                ret.DtSlbSslErrorOper.Oper.ReuseCertTypeNotZero,
			"reuse_cipher_list_not_zero":              ret.DtSlbSslErrorOper.Oper.ReuseCipherListNotZero,
			"scsv_received_when_renegotiating":        ret.DtSlbSslErrorOper.Oper.ScsvReceivedWhenRenegotiating,
			"session_id_context_uninitialized":        ret.DtSlbSslErrorOper.Oper.SessionIdContextUninitialized,
			"short_read":                              ret.DtSlbSslErrorOper.Oper.ShortRead,
			"signature_for_non_signing_certificate":   ret.DtSlbSslErrorOper.Oper.SignatureForNonSigningCertificate,
			"ssl23_doing_session_id_reuse":            ret.DtSlbSslErrorOper.Oper.Ssl23DoingSessionIdReuse,
			"ssl2_connection_id_too_long":             ret.DtSlbSslErrorOper.Oper.Ssl2ConnectionIdTooLong,
			"ssl3_session_id_too_long":                ret.DtSlbSslErrorOper.Oper.Ssl3SessionIdTooLong,
			"ssl3_session_id_too_short":               ret.DtSlbSslErrorOper.Oper.Ssl3SessionIdTooShort,
			"sslv3_alert_bad_certificate":             ret.DtSlbSslErrorOper.Oper.Sslv3AlertBadCertificate,
			"sslv3_alert_bad_record_mac":              ret.DtSlbSslErrorOper.Oper.Sslv3AlertBadRecordMac,
			"sslv3_alert_certificate_expired":         ret.DtSlbSslErrorOper.Oper.Sslv3AlertCertificateExpired,
			"sslv3_alert_certificate_revoked":         ret.DtSlbSslErrorOper.Oper.Sslv3AlertCertificateRevoked,
			"sslv3_alert_certificate_unknown":         ret.DtSlbSslErrorOper.Oper.Sslv3AlertCertificateUnknown,
			"sslv3_alert_decompression_failure":       ret.DtSlbSslErrorOper.Oper.Sslv3AlertDecompressionFailure,
			"sslv3_alert_handshake_failure":           ret.DtSlbSslErrorOper.Oper.Sslv3AlertHandshakeFailure,
			"sslv3_alert_illegal_parameter":           ret.DtSlbSslErrorOper.Oper.Sslv3AlertIllegalParameter,
			"sslv3_alert_no_certificate":              ret.DtSlbSslErrorOper.Oper.Sslv3AlertNoCertificate,
			"sslv3_alert_peer_error_cert":             ret.DtSlbSslErrorOper.Oper.Sslv3AlertPeerErrorCert,
			"sslv3_alert_peer_error_no_cert":          ret.DtSlbSslErrorOper.Oper.Sslv3AlertPeerErrorNoCert,
			"sslv3_alert_peer_error_no_cipher":        ret.DtSlbSslErrorOper.Oper.Sslv3AlertPeerErrorNoCipher,
			"sslv3_alert_peer_error_unsupp_cert_type": ret.DtSlbSslErrorOper.Oper.Sslv3AlertPeerErrorUnsuppCertType,
			"sslv3_alert_unexpected_msg":              ret.DtSlbSslErrorOper.Oper.Sslv3AlertUnexpectedMsg,
			"sslv3_alert_unknown_remote_err_type":     ret.DtSlbSslErrorOper.Oper.Sslv3AlertUnknownRemoteErrType,
			"sslv3_alert_unspported_cert":             ret.DtSlbSslErrorOper.Oper.Sslv3AlertUnspportedCert,
			"ssl_ctx_has_no_default_ssl_version":      ret.DtSlbSslErrorOper.Oper.SslCtxHasNoDefaultSslVersion,
			"ssl_handshake_failure":                   ret.DtSlbSslErrorOper.Oper.SslHandshakeFailure,
			"ssl_library_has_no_ciphers":              ret.DtSlbSslErrorOper.Oper.SslLibraryHasNoCiphers,
			"ssl_session_id_callback_failed":          ret.DtSlbSslErrorOper.Oper.SslSessionIdCallbackFailed,
			"ssl_session_id_conflict":                 ret.DtSlbSslErrorOper.Oper.SslSessionIdConflict,
			"ssl_session_id_context_too_long":         ret.DtSlbSslErrorOper.Oper.SslSessionIdContextTooLong,
			"ssl_session_id_has_bad_length":           ret.DtSlbSslErrorOper.Oper.SslSessionIdHasBadLength,
			"ssl_session_id_is_different":             ret.DtSlbSslErrorOper.Oper.SslSessionIdIsDifferent,
			"tlsv1_alert_access_denied":               ret.DtSlbSslErrorOper.Oper.Tlsv1AlertAccessDenied,
			"tlsv1_alert_decode_error":                ret.DtSlbSslErrorOper.Oper.Tlsv1AlertDecodeError,
			"tlsv1_alert_decryption_failed":           ret.DtSlbSslErrorOper.Oper.Tlsv1AlertDecryptionFailed,
			"tlsv1_alert_decrypt_error":               ret.DtSlbSslErrorOper.Oper.Tlsv1AlertDecryptError,
			"tlsv1_alert_export_restriction":          ret.DtSlbSslErrorOper.Oper.Tlsv1AlertExportRestriction,
			"tlsv1_alert_insufficient_security":       ret.DtSlbSslErrorOper.Oper.Tlsv1AlertInsufficientSecurity,
			"tlsv1_alert_internal_error":              ret.DtSlbSslErrorOper.Oper.Tlsv1AlertInternalError,
			"tlsv1_alert_no_renegotiation":            ret.DtSlbSslErrorOper.Oper.Tlsv1AlertNoRenegotiation,
			"tlsv1_alert_protocol_version":            ret.DtSlbSslErrorOper.Oper.Tlsv1AlertProtocolVersion,
			"tlsv1_alert_record_overflow":             ret.DtSlbSslErrorOper.Oper.Tlsv1AlertRecordOverflow,
			"tlsv1_alert_unknown_ca":                  ret.DtSlbSslErrorOper.Oper.Tlsv1AlertUnknownCa,
			"tlsv1_alert_user_cancelled":              ret.DtSlbSslErrorOper.Oper.Tlsv1AlertUserCancelled,
			"tls_client_cert_req_with_anon_cipher":    ret.DtSlbSslErrorOper.Oper.TlsClientCertReqWithAnonCipher,
			"tls_peer_did_not_respond_with_cert_list": ret.DtSlbSslErrorOper.Oper.TlsPeerDidNotRespondWithCertList,
			"tls_rsa_encrypted_value_length_is_wrong": ret.DtSlbSslErrorOper.Oper.TlsRsaEncryptedValueLengthIsWrong,
			"tried_to_use_unsupported_cipher":         ret.DtSlbSslErrorOper.Oper.TriedToUseUnsupportedCipher,
			"unable_to_decode_dh_certs":               ret.DtSlbSslErrorOper.Oper.UnableToDecodeDhCerts,
			"unable_to_extract_public_key":            ret.DtSlbSslErrorOper.Oper.UnableToExtractPublicKey,
			"unable_to_find_dh_parameters":            ret.DtSlbSslErrorOper.Oper.UnableToFindDhParameters,
			"unable_to_find_public_key_parameters":    ret.DtSlbSslErrorOper.Oper.UnableToFindPublicKeyParameters,
			"unable_to_find_ssl_method":               ret.DtSlbSslErrorOper.Oper.UnableToFindSslMethod,
			"unable_to_load_ssl2_md5_routines":        ret.DtSlbSslErrorOper.Oper.UnableToLoadSsl2Md5Routines,
			"unable_to_load_ssl3_md5_routines":        ret.DtSlbSslErrorOper.Oper.UnableToLoadSsl3Md5Routines,
			"unable_to_load_ssl3_sha1_routines":       ret.DtSlbSslErrorOper.Oper.UnableToLoadSsl3Sha1Routines,
			"unexpected_message":                      ret.DtSlbSslErrorOper.Oper.UnexpectedMessage,
			"unexpected_record":                       ret.DtSlbSslErrorOper.Oper.UnexpectedRecord,
			"uninitialized":                           ret.DtSlbSslErrorOper.Oper.Uninitialized,
			"unknown_alert_type":                      ret.DtSlbSslErrorOper.Oper.UnknownAlertType,
			"unknown_certificate_type":                ret.DtSlbSslErrorOper.Oper.UnknownCertificateType,
			"unknown_cipher_returned":                 ret.DtSlbSslErrorOper.Oper.UnknownCipherReturned,
			"unknown_cipher_type":                     ret.DtSlbSslErrorOper.Oper.UnknownCipherType,
			"unknown_key_exchange_type":               ret.DtSlbSslErrorOper.Oper.UnknownKeyExchangeType,
			"unknown_pkey_type":                       ret.DtSlbSslErrorOper.Oper.UnknownPkeyType,
			"unknown_protocol":                        ret.DtSlbSslErrorOper.Oper.UnknownProtocol,
			"unknown_remote_error_type":               ret.DtSlbSslErrorOper.Oper.UnknownRemoteErrorType,
			"unknown_ssl_version":                     ret.DtSlbSslErrorOper.Oper.UnknownSslVersion,
			"unknown_state":                           ret.DtSlbSslErrorOper.Oper.UnknownState,
			"unsupported_cipher":                      ret.DtSlbSslErrorOper.Oper.UnsupportedCipher,
			"unsupported_compression_algorithm":       ret.DtSlbSslErrorOper.Oper.UnsupportedCompressionAlgorithm,
			"unsupported_option":                      ret.DtSlbSslErrorOper.Oper.UnsupportedOption,
			"unsupported_protocol":                    ret.DtSlbSslErrorOper.Oper.UnsupportedProtocol,
			"unsupported_ssl_version":                 ret.DtSlbSslErrorOper.Oper.UnsupportedSslVersion,
			"unsupported_status_type":                 ret.DtSlbSslErrorOper.Oper.UnsupportedStatusType,
			"write_bio_not_set":                       ret.DtSlbSslErrorOper.Oper.WriteBioNotSet,
			"wrong_cipher_returned":                   ret.DtSlbSslErrorOper.Oper.WrongCipherReturned,
			"wrong_message_type":                      ret.DtSlbSslErrorOper.Oper.WrongMessageType,
			"wrong_number_of_key_bits":                ret.DtSlbSslErrorOper.Oper.WrongNumberOfKeyBits,
			"wrong_signature_length":                  ret.DtSlbSslErrorOper.Oper.WrongSignatureLength,
			"wrong_signature_size":                    ret.DtSlbSslErrorOper.Oper.WrongSignatureSize,
			"wrong_ssl_version":                       ret.DtSlbSslErrorOper.Oper.WrongSslVersion,
			"wrong_version_number":                    ret.DtSlbSslErrorOper.Oper.WrongVersionNumber,
			"x509_lib":                                ret.DtSlbSslErrorOper.Oper.X509Lib,
			"x509_verification_setup_problems":        ret.DtSlbSslErrorOper.Oper.X509VerificationSetupProblems,
			"clienthello_tlsext":                      ret.DtSlbSslErrorOper.Oper.ClienthelloTlsext,
			"parse_tlsext":                            ret.DtSlbSslErrorOper.Oper.ParseTlsext,
			"serverhello_tlsext":                      ret.DtSlbSslErrorOper.Oper.ServerhelloTlsext,
			"ssl3_ext_invalid_servername":             ret.DtSlbSslErrorOper.Oper.Ssl3ExtInvalidServername,
			"ssl3_ext_invalid_servername_type":        ret.DtSlbSslErrorOper.Oper.Ssl3ExtInvalidServernameType,
			"multiple_sgc_restarts":                   ret.DtSlbSslErrorOper.Oper.MultipleSgcRestarts,
			"tls_invalid_ecpointformat_list":          ret.DtSlbSslErrorOper.Oper.TlsInvalidEcpointformatList,
			"bad_ecc_cert":                            ret.DtSlbSslErrorOper.Oper.BadEccCert,
			"bad_ecdsa_sig":                           ret.DtSlbSslErrorOper.Oper.BadEcdsaSig,
			"bad_ecpoint":                             ret.DtSlbSslErrorOper.Oper.BadEcpoint,
			"cookie_mismatch":                         ret.DtSlbSslErrorOper.Oper.CookieMismatch,
			"unsupported_elliptic_curve":              ret.DtSlbSslErrorOper.Oper.UnsupportedEllipticCurve,
			"no_required_digest":                      ret.DtSlbSslErrorOper.Oper.NoRequiredDigest,
			"unsupported_digest_type":                 ret.DtSlbSslErrorOper.Oper.UnsupportedDigestType,
			"bad_handshake_length":                    ret.DtSlbSslErrorOper.Oper.BadHandshakeLength,
		},
	}
}

func getObjectSlbSslErrorOperOper(d []interface{}) edpt.SlbSslErrorOperOper {

	count1 := len(d)
	var ret edpt.SlbSslErrorOperOper
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
		ret.BadProtocolVersionNumber = in["bad_protocol_version_number"].(int)
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
		ret.WrongNumberOfKeyBits = in["wrong_number_of_key_bits"].(int)
		ret.WrongSignatureLength = in["wrong_signature_length"].(int)
		ret.WrongSignatureSize = in["wrong_signature_size"].(int)
		ret.WrongSslVersion = in["wrong_ssl_version"].(int)
		ret.WrongVersionNumber = in["wrong_version_number"].(int)
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

func dataToEndpointSlbSslErrorOper(d *schema.ResourceData) edpt.SlbSslErrorOper {
	var ret edpt.SlbSslErrorOper

	ret.Oper = getObjectSlbSslErrorOperOper(d.Get("oper").([]interface{}))
	return ret
}
