package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_ssl_error_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"app_data_in_handshake": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for app data in handshake",
			},
			"attempt_to_reuse_sess_in_diff_context": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for attempt to reuse sess in diff context",
			},
			"bad_alert_record": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad alert record",
			},
			"bad_authentication_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad authentication type",
			},
			"bad_change_cipher_spec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad change cipher spec",
			},
			"bad_checksum": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad checksum",
			},
			"bad_data_returned_by_callback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad data returned by callback",
			},
			"bad_decompression": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad decompression",
			},
			"bad_dh_g_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh g length",
			},
			"bad_dh_p_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh p length",
			},
			"bad_dh_pub_key_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh pub key length",
			},
			"bad_digest_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad digest length",
			},
			"bad_dsa_signature": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dsa signature",
			},
			"bad_ecc_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecc cert",
			},
			"bad_ecdsa_sig": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecdsa sig",
			},
			"bad_ecpoint": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecpoint",
			},
			"bad_handshake_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad handshake length",
			},
			"bad_hello_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad hello request",
			},
			"bad_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad length",
			},
			"bad_mac_decode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad mac decode",
			},
			"bad_message_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad message type",
			},
			"bad_packet_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad packet length",
			},
			"bad_protocol_version_counter": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad protocol version counter",
			},
			"bad_response_argument": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad response argument",
			},
			"bad_rsa_decrypt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa decrypt",
			},
			"bad_rsa_e_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa e length",
			},
			"bad_rsa_encrypt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa encrypt",
			},
			"bad_rsa_modulus_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa modulus length",
			},
			"bad_rsa_signature": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa signature",
			},
			"bad_signature": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad signature",
			},
			"bad_ssl_filetype": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ssl filetype",
			},
			"bad_ssl_session_id_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ssl session id length",
			},
			"bad_state": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad state",
			},
			"bad_write_retry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad write retry",
			},
			"bio_not_set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bio not set",
			},
			"block_cipher_pad_is_wrong": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for block cipher pad is wrong",
			},
			"bn_lib": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bn lib",
			},
			"ca_dn_length_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ca dn length mismatch",
			},
			"ca_dn_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ca dn too long",
			},
			"ccs_received_early": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ccs received early",
			},
			"cert_length_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cert length mismatch",
			},
			"certificate_verify_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for certificate verify failed",
			},
			"challenge_is_different": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for challenge is different",
			},
			"cipher_code_wrong_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher code wrong length",
			},
			"cipher_or_hash_unavailable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher or hash unavailable",
			},
			"cipher_table_src_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher table src error",
			},
			"clienthello_tlsext": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for clienthello tlsext",
			},
			"compressed_length_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compressed length too long",
			},
			"compression_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression failure",
			},
			"compression_library_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression library error",
			},
			"connection_id_is_different": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connection id is different",
			},
			"connection_type_not_set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connection type not set",
			},
			"cookie_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cookie mismatch",
			},
			"data_between_ccs_and_finished": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data between ccs and finished",
			},
			"data_length_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data length too long",
			},
			"decryption_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for decryption failed",
			},
			"decryption_failed_or_bad_record_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for decryption failed or bad record mac",
			},
			"dh_public_value_length_is_wrong": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dh public value length is wrong",
			},
			"digest_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for digest check failed",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"encrypted_length_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for encrypted length too long",
			},
			"error_generating_tmp_rsa_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error generating tmp rsa key",
			},
			"error_in_received_cipher_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error in received cipher list",
			},
			"excessive_message_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for excessive message size",
			},
			"extra_data_in_message": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for extra data in message",
			},
			"got_a_fin_before_a_ccs": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for got a fin before a ccs",
			},
			"http_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for http request",
			},
			"https_proxy_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for https proxy request",
			},
			"illegal_padding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for illegal padding",
			},
			"inappropriate_fallback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inappropriate fallback",
			},
			"invalid_challenge_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid challenge length",
			},
			"invalid_command": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid command",
			},
			"invalid_purpose": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid purpose",
			},
			"invalid_status_response": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid status response",
			},
			"invalid_trust": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid trust",
			},
			"key_arg_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for key arg too long",
			},
			"krb5": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5",
			},
			"krb5_client_cc_principal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client cc principal",
			},
			"krb5_client_get_cred": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client get cred",
			},
			"krb5_client_init": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client init",
			},
			"krb5_client_mk_req": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client mk_req",
			},
			"krb5_server_bad_ticket": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server bad ticket",
			},
			"krb5_server_init": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server init",
			},
			"krb5_server_rd_req": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server rd_req",
			},
			"krb5_server_tkt_expired": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt expired",
			},
			"krb5_server_tkt_not_yet_valid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt not yet valid",
			},
			"krb5_server_tkt_skew": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt skew",
			},
			"length_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for length mismatch",
			},
			"length_too_short": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for length too short",
			},
			"library_bug": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for library bug",
			},
			"library_has_no_ciphers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for library has no ciphers",
			},
			"mast_key_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for mast key too long",
			},
			"message_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for message too long",
			},
			"missing_dh_dsa_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh dsa cert",
			},
			"missing_dh_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh key",
			},
			"missing_dh_rsa_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh rsa cert",
			},
			"missing_dsa_signing_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dsa signing cert",
			},
			"missing_export_tmp_dh_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing export tmp dh key",
			},
			"missing_export_tmp_rsa_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing export tmp rsa key",
			},
			"missing_rsa_certificate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa certificate",
			},
			"missing_rsa_encrypting_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa encrypting cert",
			},
			"missing_rsa_signing_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa signing cert",
			},
			"missing_tmp_dh_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp dh key",
			},
			"missing_tmp_rsa_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp rsa key",
			},
			"missing_tmp_rsa_pkey": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp rsa pkey",
			},
			"missing_verify_message": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing verify message",
			},
			"multiple_sgc_restarts": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for multiple sgc restarts",
			},
			"no_certificate_assigned": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate assigned",
			},
			"no_certificate_returned": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate returned",
			},
			"no_certificate_set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate set",
			},
			"no_certificate_specified": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate specified",
			},
			"no_certificates_returned": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificates returned",
			},
			"no_cipher_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no cipher list",
			},
			"no_cipher_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no cipher match",
			},
			"no_ciphers_available": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers available",
			},
			"no_ciphers_passed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers passed",
			},
			"no_ciphers_specified": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers specified",
			},
			"no_client_cert_received": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no client cert received",
			},
			"no_compression_specified": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no compression specified",
			},
			"no_method_specified": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no method specified",
			},
			"no_private_key_assigned": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no private key assigned",
			},
			"no_privatekey": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no privatekey",
			},
			"no_protocols_available": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no protocols available",
			},
			"no_publickey": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no publickey",
			},
			"no_required_digest": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no required digest",
			},
			"no_shared_cipher": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no shared cipher",
			},
			"no_verify_callback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no verify callback",
			},
			"non_sslv2_initial_packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for non sslv2 initial packet",
			},
			"null_ssl_ctx": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for null ssl ctx",
			},
			"null_ssl_method_passed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for null ssl method passed",
			},
			"old_session_cipher_not_returned": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for old session cipher not returned",
			},
			"packet_length_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for packet length too long",
			},
			"parse_tlsext": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for parse tlsext",
			},
			"path_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for path too long",
			},
			"peer_did_not_return_a_certificate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer did not return a certificate",
			},
			"peer_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error",
			},
			"peer_error_certificate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error certificate",
			},
			"peer_error_no_certificate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error no certificate",
			},
			"peer_error_no_cipher": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error no cipher",
			},
			"peer_error_unsupported_certificate_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error unsupported certificate type",
			},
			"pre_mac_length_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for pre mac length too long",
			},
			"problems_mapping_cipher_functions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for problems mapping cipher functions",
			},
			"protocol_is_shutdown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for protocol is shutdown",
			},
			"public_key_encrypt_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key encrypt error",
			},
			"public_key_is_not_rsa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key is not rsa",
			},
			"public_key_not_rsa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key not rsa",
			},
			"read_bio_not_set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for read bio not set",
			},
			"read_wrong_packet_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for read wrong packet type",
			},
			"record_length_mismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record length mismatch",
			},
			"record_too_large": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record too large",
			},
			"record_too_small": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record too small",
			},
			"required_cipher_missing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for required cipher missing",
			},
			"reuse_cert_length_not_zero": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cert length not zero",
			},
			"reuse_cert_type_not_zero": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cert type not zero",
			},
			"reuse_cipher_list_not_zero": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cipher list not zero",
			},
			"scsv_received_when_renegotiating": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for scsv received when renegotiating",
			},
			"serverhello_tlsext": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for serverhello tlsext",
			},
			"session_id_context_uninitialized": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for session id context uninitialized",
			},
			"short_read": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for short read",
			},
			"signature_for_non_signing_certificate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for signature for non signing certificate",
			},
			"ssl_ctx_has_no_default_ssl_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl ctx has no default ssl version",
			},
			"ssl_handshake_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl handshake failure",
			},
			"ssl_library_has_no_ciphers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl library has no ciphers",
			},
			"ssl_session_id_callback_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id callback failed",
			},
			"ssl_session_id_conflict": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id conflict",
			},
			"ssl_session_id_context_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id context too long",
			},
			"ssl_session_id_has_bad_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id has bad length",
			},
			"ssl_session_id_is_different": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id is different",
			},
			"ssl2_connection_id_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl2 connection id too long",
			},
			"ssl23_doing_session_id_reuse": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl23 doing session id reuse",
			},
			"ssl3_ext_invalid_servername": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 ext invalid servername",
			},
			"ssl3_ext_invalid_servername_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 ext invalid servername type",
			},
			"ssl3_session_id_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 session id too long",
			},
			"ssl3_session_id_too_short": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 session id too short",
			},
			"sslv3_alert_bad_certificate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert bad certificate",
			},
			"sslv3_alert_bad_record_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert bad record mac",
			},
			"sslv3_alert_certificate_expired": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate expired",
			},
			"sslv3_alert_certificate_revoked": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate revoked",
			},
			"sslv3_alert_certificate_unknown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate unknown",
			},
			"sslv3_alert_decompression_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert decompression failure",
			},
			"sslv3_alert_handshake_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert handshake failure",
			},
			"sslv3_alert_illegal_parameter": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert illegal parameter",
			},
			"sslv3_alert_no_certificate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert no certificate",
			},
			"sslv3_alert_peer_error_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error cert",
			},
			"sslv3_alert_peer_error_no_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error no cert",
			},
			"sslv3_alert_peer_error_no_cipher": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error no cipher",
			},
			"sslv3_alert_peer_error_unsupp_cert_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error unsupp cert type",
			},
			"sslv3_alert_unexpected_msg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unexpected msg",
			},
			"sslv3_alert_unknown_remote_err_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unknown remote err type",
			},
			"sslv3_alert_unspported_cert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unspported cert",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
			},
			"tls_client_cert_req_with_anon_cipher": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls client cert req with anon cipher",
			},
			"tls_invalid_ecpointformat_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls invalid ecpointformat list",
			},
			"tls_peer_did_not_respond_with_cert_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls peer did not respond with cert list",
			},
			"tls_rsa_encrypted_value_length_is_wrong": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls rsa encrypted value length is wrong",
			},
			"tlsv1_alert_access_denied": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert access denied",
			},
			"tlsv1_alert_decode_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decode error",
			},
			"tlsv1_alert_decrypt_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decrypt error",
			},
			"tlsv1_alert_decryption_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decryption failed",
			},
			"tlsv1_alert_export_restriction": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert export restriction",
			},
			"tlsv1_alert_insufficient_security": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert insufficient security",
			},
			"tlsv1_alert_internal_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert internal error",
			},
			"tlsv1_alert_no_renegotiation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert no renegotiation",
			},
			"tlsv1_alert_protocol_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert protocol version",
			},
			"tlsv1_alert_record_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert record overflow",
			},
			"tlsv1_alert_unknown_ca": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert unknown ca",
			},
			"tlsv1_alert_user_cancelled": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert user cancelled",
			},
			"tried_to_use_unsupported_cipher": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tried to use unsupported cipher",
			},
			"unable_to_decode_dh_certs": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to decode dh certs",
			},
			"unable_to_extract_public_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to extract public key",
			},
			"unable_to_find_dh_parameters": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find dh parameters",
			},
			"unable_to_find_public_key_parameters": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find public key parameters",
			},
			"unable_to_find_ssl_method": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find ssl method",
			},
			"unable_to_load_ssl2_md5_routines": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl2 md5 routines",
			},
			"unable_to_load_ssl3_md5_routines": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl3 md5 routines",
			},
			"unable_to_load_ssl3_sha1_routines": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl3 sha1 routines",
			},
			"unexpected_message": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected message",
			},
			"unexpected_record": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected record",
			},
			"uninitialized": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for uninitialized",
			},
			"unknown_alert_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown alert type",
			},
			"unknown_certificate_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown certificate type",
			},
			"unknown_cipher_returned": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown cipher returned",
			},
			"unknown_cipher_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown cipher type",
			},
			"unknown_key_exchange_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown key exchange type",
			},
			"unknown_pkey_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown pkey type",
			},
			"unknown_protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown protocol",
			},
			"unknown_remote_error_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown remote error type",
			},
			"unknown_ssl_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown ssl version",
			},
			"unknown_state": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown state",
			},
			"unsupported_cipher": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported cipher",
			},
			"unsupported_compression_algorithm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported compression algorithm",
			},
			"unsupported_digest_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported digest type",
			},
			"unsupported_elliptic_curve": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported elliptic curve",
			},
			"unsupported_option": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported option",
			},
			"unsupported_protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported protocol",
			},
			"unsupported_ssl_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported ssl version",
			},
			"unsupported_status_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported status type",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"write_bio_not_set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for write bio not set",
			},
			"wrong_cipher_returned": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong cipher returned",
			},
			"wrong_counter_of_key_bits": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong counter of key bits",
			},
			"wrong_message_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong message type",
			},
			"wrong_signature_length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong signature length",
			},
			"wrong_signature_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong signature size",
			},
			"wrong_ssl_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong ssl version",
			},
			"wrong_version_counter": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong version counter",
			},
			"x509_lib": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for x509 lib",
			},
			"x509_verification_setup_problems": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for x509 verification setup problems",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate
	ret.Inst.AppDataInHandshake = d.Get("app_data_in_handshake").(int)
	ret.Inst.AttemptToReuseSessInDiffContext = d.Get("attempt_to_reuse_sess_in_diff_context").(int)
	ret.Inst.BadAlertRecord = d.Get("bad_alert_record").(int)
	ret.Inst.BadAuthenticationType = d.Get("bad_authentication_type").(int)
	ret.Inst.BadChangeCipherSpec = d.Get("bad_change_cipher_spec").(int)
	ret.Inst.BadChecksum = d.Get("bad_checksum").(int)
	ret.Inst.BadDataReturnedByCallback = d.Get("bad_data_returned_by_callback").(int)
	ret.Inst.BadDecompression = d.Get("bad_decompression").(int)
	ret.Inst.BadDhGLength = d.Get("bad_dh_g_length").(int)
	ret.Inst.BadDhPLength = d.Get("bad_dh_p_length").(int)
	ret.Inst.BadDhPubKeyLength = d.Get("bad_dh_pub_key_length").(int)
	ret.Inst.BadDigestLength = d.Get("bad_digest_length").(int)
	ret.Inst.BadDsaSignature = d.Get("bad_dsa_signature").(int)
	ret.Inst.BadEccCert = d.Get("bad_ecc_cert").(int)
	ret.Inst.BadEcdsaSig = d.Get("bad_ecdsa_sig").(int)
	ret.Inst.BadEcpoint = d.Get("bad_ecpoint").(int)
	ret.Inst.BadHandshakeLength = d.Get("bad_handshake_length").(int)
	ret.Inst.BadHelloRequest = d.Get("bad_hello_request").(int)
	ret.Inst.BadLength = d.Get("bad_length").(int)
	ret.Inst.BadMacDecode = d.Get("bad_mac_decode").(int)
	ret.Inst.BadMessageType = d.Get("bad_message_type").(int)
	ret.Inst.BadPacketLength = d.Get("bad_packet_length").(int)
	ret.Inst.BadProtocolVersionCounter = d.Get("bad_protocol_version_counter").(int)
	ret.Inst.BadResponseArgument = d.Get("bad_response_argument").(int)
	ret.Inst.BadRsaDecrypt = d.Get("bad_rsa_decrypt").(int)
	ret.Inst.BadRsaELength = d.Get("bad_rsa_e_length").(int)
	ret.Inst.BadRsaEncrypt = d.Get("bad_rsa_encrypt").(int)
	ret.Inst.BadRsaModulusLength = d.Get("bad_rsa_modulus_length").(int)
	ret.Inst.BadRsaSignature = d.Get("bad_rsa_signature").(int)
	ret.Inst.BadSignature = d.Get("bad_signature").(int)
	ret.Inst.BadSslFiletype = d.Get("bad_ssl_filetype").(int)
	ret.Inst.BadSslSessionIdLength = d.Get("bad_ssl_session_id_length").(int)
	ret.Inst.BadState = d.Get("bad_state").(int)
	ret.Inst.BadWriteRetry = d.Get("bad_write_retry").(int)
	ret.Inst.BioNotSet = d.Get("bio_not_set").(int)
	ret.Inst.BlockCipherPadIsWrong = d.Get("block_cipher_pad_is_wrong").(int)
	ret.Inst.BnLib = d.Get("bn_lib").(int)
	ret.Inst.CaDnLengthMismatch = d.Get("ca_dn_length_mismatch").(int)
	ret.Inst.CaDnTooLong = d.Get("ca_dn_too_long").(int)
	ret.Inst.CcsReceivedEarly = d.Get("ccs_received_early").(int)
	ret.Inst.CertLengthMismatch = d.Get("cert_length_mismatch").(int)
	ret.Inst.CertificateVerifyFailed = d.Get("certificate_verify_failed").(int)
	ret.Inst.ChallengeIsDifferent = d.Get("challenge_is_different").(int)
	ret.Inst.CipherCodeWrongLength = d.Get("cipher_code_wrong_length").(int)
	ret.Inst.CipherOrHashUnavailable = d.Get("cipher_or_hash_unavailable").(int)
	ret.Inst.CipherTableSrcError = d.Get("cipher_table_src_error").(int)
	ret.Inst.ClienthelloTlsext = d.Get("clienthello_tlsext").(int)
	ret.Inst.CompressedLengthTooLong = d.Get("compressed_length_too_long").(int)
	ret.Inst.CompressionFailure = d.Get("compression_failure").(int)
	ret.Inst.CompressionLibraryError = d.Get("compression_library_error").(int)
	ret.Inst.ConnectionIdIsDifferent = d.Get("connection_id_is_different").(int)
	ret.Inst.ConnectionTypeNotSet = d.Get("connection_type_not_set").(int)
	ret.Inst.CookieMismatch = d.Get("cookie_mismatch").(int)
	ret.Inst.DataBetweenCcsAndFinished = d.Get("data_between_ccs_and_finished").(int)
	ret.Inst.DataLengthTooLong = d.Get("data_length_too_long").(int)
	ret.Inst.DecryptionFailed = d.Get("decryption_failed").(int)
	ret.Inst.DecryptionFailedOrBadRecordMac = d.Get("decryption_failed_or_bad_record_mac").(int)
	ret.Inst.DhPublicValueLengthIsWrong = d.Get("dh_public_value_length_is_wrong").(int)
	ret.Inst.DigestCheckFailed = d.Get("digest_check_failed").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.EncryptedLengthTooLong = d.Get("encrypted_length_too_long").(int)
	ret.Inst.ErrorGeneratingTmpRsaKey = d.Get("error_generating_tmp_rsa_key").(int)
	ret.Inst.ErrorInReceivedCipherList = d.Get("error_in_received_cipher_list").(int)
	ret.Inst.ExcessiveMessageSize = d.Get("excessive_message_size").(int)
	ret.Inst.ExtraDataInMessage = d.Get("extra_data_in_message").(int)
	ret.Inst.GotAFinBeforeACcs = d.Get("got_a_fin_before_a_ccs").(int)
	ret.Inst.HttpRequest = d.Get("http_request").(int)
	ret.Inst.HttpsProxyRequest = d.Get("https_proxy_request").(int)
	ret.Inst.IllegalPadding = d.Get("illegal_padding").(int)
	ret.Inst.InappropriateFallback = d.Get("inappropriate_fallback").(int)
	ret.Inst.InvalidChallengeLength = d.Get("invalid_challenge_length").(int)
	ret.Inst.InvalidCommand = d.Get("invalid_command").(int)
	ret.Inst.InvalidPurpose = d.Get("invalid_purpose").(int)
	ret.Inst.InvalidStatusResponse = d.Get("invalid_status_response").(int)
	ret.Inst.InvalidTrust = d.Get("invalid_trust").(int)
	ret.Inst.KeyArgTooLong = d.Get("key_arg_too_long").(int)
	ret.Inst.Krb5 = d.Get("krb5").(int)
	ret.Inst.Krb5ClientCcPrincipal = d.Get("krb5_client_cc_principal").(int)
	ret.Inst.Krb5ClientGetCred = d.Get("krb5_client_get_cred").(int)
	ret.Inst.Krb5ClientInit = d.Get("krb5_client_init").(int)
	ret.Inst.Krb5ClientMkReq = d.Get("krb5_client_mk_req").(int)
	ret.Inst.Krb5ServerBadTicket = d.Get("krb5_server_bad_ticket").(int)
	ret.Inst.Krb5ServerInit = d.Get("krb5_server_init").(int)
	ret.Inst.Krb5ServerRdReq = d.Get("krb5_server_rd_req").(int)
	ret.Inst.Krb5ServerTktExpired = d.Get("krb5_server_tkt_expired").(int)
	ret.Inst.Krb5ServerTktNotYetValid = d.Get("krb5_server_tkt_not_yet_valid").(int)
	ret.Inst.Krb5ServerTktSkew = d.Get("krb5_server_tkt_skew").(int)
	ret.Inst.LengthMismatch = d.Get("length_mismatch").(int)
	ret.Inst.LengthTooShort = d.Get("length_too_short").(int)
	ret.Inst.LibraryBug = d.Get("library_bug").(int)
	ret.Inst.LibraryHasNoCiphers = d.Get("library_has_no_ciphers").(int)
	ret.Inst.MastKeyTooLong = d.Get("mast_key_too_long").(int)
	ret.Inst.MessageTooLong = d.Get("message_too_long").(int)
	ret.Inst.MissingDhDsaCert = d.Get("missing_dh_dsa_cert").(int)
	ret.Inst.MissingDhKey = d.Get("missing_dh_key").(int)
	ret.Inst.MissingDhRsaCert = d.Get("missing_dh_rsa_cert").(int)
	ret.Inst.MissingDsaSigningCert = d.Get("missing_dsa_signing_cert").(int)
	ret.Inst.MissingExportTmpDhKey = d.Get("missing_export_tmp_dh_key").(int)
	ret.Inst.MissingExportTmpRsaKey = d.Get("missing_export_tmp_rsa_key").(int)
	ret.Inst.MissingRsaCertificate = d.Get("missing_rsa_certificate").(int)
	ret.Inst.MissingRsaEncryptingCert = d.Get("missing_rsa_encrypting_cert").(int)
	ret.Inst.MissingRsaSigningCert = d.Get("missing_rsa_signing_cert").(int)
	ret.Inst.MissingTmpDhKey = d.Get("missing_tmp_dh_key").(int)
	ret.Inst.MissingTmpRsaKey = d.Get("missing_tmp_rsa_key").(int)
	ret.Inst.MissingTmpRsaPkey = d.Get("missing_tmp_rsa_pkey").(int)
	ret.Inst.MissingVerifyMessage = d.Get("missing_verify_message").(int)
	ret.Inst.MultipleSgcRestarts = d.Get("multiple_sgc_restarts").(int)
	ret.Inst.NoCertificateAssigned = d.Get("no_certificate_assigned").(int)
	ret.Inst.NoCertificateReturned = d.Get("no_certificate_returned").(int)
	ret.Inst.NoCertificateSet = d.Get("no_certificate_set").(int)
	ret.Inst.NoCertificateSpecified = d.Get("no_certificate_specified").(int)
	ret.Inst.NoCertificatesReturned = d.Get("no_certificates_returned").(int)
	ret.Inst.NoCipherList = d.Get("no_cipher_list").(int)
	ret.Inst.NoCipherMatch = d.Get("no_cipher_match").(int)
	ret.Inst.NoCiphersAvailable = d.Get("no_ciphers_available").(int)
	ret.Inst.NoCiphersPassed = d.Get("no_ciphers_passed").(int)
	ret.Inst.NoCiphersSpecified = d.Get("no_ciphers_specified").(int)
	ret.Inst.NoClientCertReceived = d.Get("no_client_cert_received").(int)
	ret.Inst.NoCompressionSpecified = d.Get("no_compression_specified").(int)
	ret.Inst.NoMethodSpecified = d.Get("no_method_specified").(int)
	ret.Inst.NoPrivateKeyAssigned = d.Get("no_private_key_assigned").(int)
	ret.Inst.NoPrivatekey = d.Get("no_privatekey").(int)
	ret.Inst.NoProtocolsAvailable = d.Get("no_protocols_available").(int)
	ret.Inst.NoPublickey = d.Get("no_publickey").(int)
	ret.Inst.NoRequiredDigest = d.Get("no_required_digest").(int)
	ret.Inst.NoSharedCipher = d.Get("no_shared_cipher").(int)
	ret.Inst.NoVerifyCallback = d.Get("no_verify_callback").(int)
	ret.Inst.NonSslv2InitialPacket = d.Get("non_sslv2_initial_packet").(int)
	ret.Inst.NullSslCtx = d.Get("null_ssl_ctx").(int)
	ret.Inst.NullSslMethodPassed = d.Get("null_ssl_method_passed").(int)
	ret.Inst.OldSessionCipherNotReturned = d.Get("old_session_cipher_not_returned").(int)
	ret.Inst.PacketLengthTooLong = d.Get("packet_length_too_long").(int)
	ret.Inst.ParseTlsext = d.Get("parse_tlsext").(int)
	ret.Inst.PathTooLong = d.Get("path_too_long").(int)
	ret.Inst.PeerDidNotReturnACertificate = d.Get("peer_did_not_return_a_certificate").(int)
	ret.Inst.PeerError = d.Get("peer_error").(int)
	ret.Inst.PeerErrorCertificate = d.Get("peer_error_certificate").(int)
	ret.Inst.PeerErrorNoCertificate = d.Get("peer_error_no_certificate").(int)
	ret.Inst.PeerErrorNoCipher = d.Get("peer_error_no_cipher").(int)
	ret.Inst.PeerErrorUnsupportedCertificateType = d.Get("peer_error_unsupported_certificate_type").(int)
	ret.Inst.PreMacLengthTooLong = d.Get("pre_mac_length_too_long").(int)
	ret.Inst.ProblemsMappingCipherFunctions = d.Get("problems_mapping_cipher_functions").(int)
	ret.Inst.ProtocolIsShutdown = d.Get("protocol_is_shutdown").(int)
	ret.Inst.PublicKeyEncryptError = d.Get("public_key_encrypt_error").(int)
	ret.Inst.PublicKeyIsNotRsa = d.Get("public_key_is_not_rsa").(int)
	ret.Inst.PublicKeyNotRsa = d.Get("public_key_not_rsa").(int)
	ret.Inst.ReadBioNotSet = d.Get("read_bio_not_set").(int)
	ret.Inst.ReadWrongPacketType = d.Get("read_wrong_packet_type").(int)
	ret.Inst.RecordLengthMismatch = d.Get("record_length_mismatch").(int)
	ret.Inst.RecordTooLarge = d.Get("record_too_large").(int)
	ret.Inst.RecordTooSmall = d.Get("record_too_small").(int)
	ret.Inst.RequiredCipherMissing = d.Get("required_cipher_missing").(int)
	ret.Inst.ReuseCertLengthNotZero = d.Get("reuse_cert_length_not_zero").(int)
	ret.Inst.ReuseCertTypeNotZero = d.Get("reuse_cert_type_not_zero").(int)
	ret.Inst.ReuseCipherListNotZero = d.Get("reuse_cipher_list_not_zero").(int)
	ret.Inst.ScsvReceivedWhenRenegotiating = d.Get("scsv_received_when_renegotiating").(int)
	ret.Inst.ServerhelloTlsext = d.Get("serverhello_tlsext").(int)
	ret.Inst.SessionIdContextUninitialized = d.Get("session_id_context_uninitialized").(int)
	ret.Inst.ShortRead = d.Get("short_read").(int)
	ret.Inst.SignatureForNonSigningCertificate = d.Get("signature_for_non_signing_certificate").(int)
	ret.Inst.SslCtxHasNoDefaultSslVersion = d.Get("ssl_ctx_has_no_default_ssl_version").(int)
	ret.Inst.SslHandshakeFailure = d.Get("ssl_handshake_failure").(int)
	ret.Inst.SslLibraryHasNoCiphers = d.Get("ssl_library_has_no_ciphers").(int)
	ret.Inst.SslSessionIdCallbackFailed = d.Get("ssl_session_id_callback_failed").(int)
	ret.Inst.SslSessionIdConflict = d.Get("ssl_session_id_conflict").(int)
	ret.Inst.SslSessionIdContextTooLong = d.Get("ssl_session_id_context_too_long").(int)
	ret.Inst.SslSessionIdHasBadLength = d.Get("ssl_session_id_has_bad_length").(int)
	ret.Inst.SslSessionIdIsDifferent = d.Get("ssl_session_id_is_different").(int)
	ret.Inst.Ssl2ConnectionIdTooLong = d.Get("ssl2_connection_id_too_long").(int)
	ret.Inst.Ssl23DoingSessionIdReuse = d.Get("ssl23_doing_session_id_reuse").(int)
	ret.Inst.Ssl3ExtInvalidServername = d.Get("ssl3_ext_invalid_servername").(int)
	ret.Inst.Ssl3ExtInvalidServernameType = d.Get("ssl3_ext_invalid_servername_type").(int)
	ret.Inst.Ssl3SessionIdTooLong = d.Get("ssl3_session_id_too_long").(int)
	ret.Inst.Ssl3SessionIdTooShort = d.Get("ssl3_session_id_too_short").(int)
	ret.Inst.Sslv3AlertBadCertificate = d.Get("sslv3_alert_bad_certificate").(int)
	ret.Inst.Sslv3AlertBadRecordMac = d.Get("sslv3_alert_bad_record_mac").(int)
	ret.Inst.Sslv3AlertCertificateExpired = d.Get("sslv3_alert_certificate_expired").(int)
	ret.Inst.Sslv3AlertCertificateRevoked = d.Get("sslv3_alert_certificate_revoked").(int)
	ret.Inst.Sslv3AlertCertificateUnknown = d.Get("sslv3_alert_certificate_unknown").(int)
	ret.Inst.Sslv3AlertDecompressionFailure = d.Get("sslv3_alert_decompression_failure").(int)
	ret.Inst.Sslv3AlertHandshakeFailure = d.Get("sslv3_alert_handshake_failure").(int)
	ret.Inst.Sslv3AlertIllegalParameter = d.Get("sslv3_alert_illegal_parameter").(int)
	ret.Inst.Sslv3AlertNoCertificate = d.Get("sslv3_alert_no_certificate").(int)
	ret.Inst.Sslv3AlertPeerErrorCert = d.Get("sslv3_alert_peer_error_cert").(int)
	ret.Inst.Sslv3AlertPeerErrorNoCert = d.Get("sslv3_alert_peer_error_no_cert").(int)
	ret.Inst.Sslv3AlertPeerErrorNoCipher = d.Get("sslv3_alert_peer_error_no_cipher").(int)
	ret.Inst.Sslv3AlertPeerErrorUnsuppCertType = d.Get("sslv3_alert_peer_error_unsupp_cert_type").(int)
	ret.Inst.Sslv3AlertUnexpectedMsg = d.Get("sslv3_alert_unexpected_msg").(int)
	ret.Inst.Sslv3AlertUnknownRemoteErrType = d.Get("sslv3_alert_unknown_remote_err_type").(int)
	ret.Inst.Sslv3AlertUnspportedCert = d.Get("sslv3_alert_unspported_cert").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	ret.Inst.TlsClientCertReqWithAnonCipher = d.Get("tls_client_cert_req_with_anon_cipher").(int)
	ret.Inst.TlsInvalidEcpointformatList = d.Get("tls_invalid_ecpointformat_list").(int)
	ret.Inst.TlsPeerDidNotRespondWithCertList = d.Get("tls_peer_did_not_respond_with_cert_list").(int)
	ret.Inst.TlsRsaEncryptedValueLengthIsWrong = d.Get("tls_rsa_encrypted_value_length_is_wrong").(int)
	ret.Inst.Tlsv1AlertAccessDenied = d.Get("tlsv1_alert_access_denied").(int)
	ret.Inst.Tlsv1AlertDecodeError = d.Get("tlsv1_alert_decode_error").(int)
	ret.Inst.Tlsv1AlertDecryptError = d.Get("tlsv1_alert_decrypt_error").(int)
	ret.Inst.Tlsv1AlertDecryptionFailed = d.Get("tlsv1_alert_decryption_failed").(int)
	ret.Inst.Tlsv1AlertExportRestriction = d.Get("tlsv1_alert_export_restriction").(int)
	ret.Inst.Tlsv1AlertInsufficientSecurity = d.Get("tlsv1_alert_insufficient_security").(int)
	ret.Inst.Tlsv1AlertInternalError = d.Get("tlsv1_alert_internal_error").(int)
	ret.Inst.Tlsv1AlertNoRenegotiation = d.Get("tlsv1_alert_no_renegotiation").(int)
	ret.Inst.Tlsv1AlertProtocolVersion = d.Get("tlsv1_alert_protocol_version").(int)
	ret.Inst.Tlsv1AlertRecordOverflow = d.Get("tlsv1_alert_record_overflow").(int)
	ret.Inst.Tlsv1AlertUnknownCa = d.Get("tlsv1_alert_unknown_ca").(int)
	ret.Inst.Tlsv1AlertUserCancelled = d.Get("tlsv1_alert_user_cancelled").(int)
	ret.Inst.TriedToUseUnsupportedCipher = d.Get("tried_to_use_unsupported_cipher").(int)
	ret.Inst.UnableToDecodeDhCerts = d.Get("unable_to_decode_dh_certs").(int)
	ret.Inst.UnableToExtractPublicKey = d.Get("unable_to_extract_public_key").(int)
	ret.Inst.UnableToFindDhParameters = d.Get("unable_to_find_dh_parameters").(int)
	ret.Inst.UnableToFindPublicKeyParameters = d.Get("unable_to_find_public_key_parameters").(int)
	ret.Inst.UnableToFindSslMethod = d.Get("unable_to_find_ssl_method").(int)
	ret.Inst.UnableToLoadSsl2Md5Routines = d.Get("unable_to_load_ssl2_md5_routines").(int)
	ret.Inst.UnableToLoadSsl3Md5Routines = d.Get("unable_to_load_ssl3_md5_routines").(int)
	ret.Inst.UnableToLoadSsl3Sha1Routines = d.Get("unable_to_load_ssl3_sha1_routines").(int)
	ret.Inst.UnexpectedMessage = d.Get("unexpected_message").(int)
	ret.Inst.UnexpectedRecord = d.Get("unexpected_record").(int)
	ret.Inst.Uninitialized = d.Get("uninitialized").(int)
	ret.Inst.UnknownAlertType = d.Get("unknown_alert_type").(int)
	ret.Inst.UnknownCertificateType = d.Get("unknown_certificate_type").(int)
	ret.Inst.UnknownCipherReturned = d.Get("unknown_cipher_returned").(int)
	ret.Inst.UnknownCipherType = d.Get("unknown_cipher_type").(int)
	ret.Inst.UnknownKeyExchangeType = d.Get("unknown_key_exchange_type").(int)
	ret.Inst.UnknownPkeyType = d.Get("unknown_pkey_type").(int)
	ret.Inst.UnknownProtocol = d.Get("unknown_protocol").(int)
	ret.Inst.UnknownRemoteErrorType = d.Get("unknown_remote_error_type").(int)
	ret.Inst.UnknownSslVersion = d.Get("unknown_ssl_version").(int)
	ret.Inst.UnknownState = d.Get("unknown_state").(int)
	ret.Inst.UnsupportedCipher = d.Get("unsupported_cipher").(int)
	ret.Inst.UnsupportedCompressionAlgorithm = d.Get("unsupported_compression_algorithm").(int)
	ret.Inst.UnsupportedDigestType = d.Get("unsupported_digest_type").(int)
	ret.Inst.UnsupportedEllipticCurve = d.Get("unsupported_elliptic_curve").(int)
	ret.Inst.UnsupportedOption = d.Get("unsupported_option").(int)
	ret.Inst.UnsupportedProtocol = d.Get("unsupported_protocol").(int)
	ret.Inst.UnsupportedSslVersion = d.Get("unsupported_ssl_version").(int)
	ret.Inst.UnsupportedStatusType = d.Get("unsupported_status_type").(int)
	//omit uuid
	ret.Inst.WriteBioNotSet = d.Get("write_bio_not_set").(int)
	ret.Inst.WrongCipherReturned = d.Get("wrong_cipher_returned").(int)
	ret.Inst.WrongCounterOfKeyBits = d.Get("wrong_counter_of_key_bits").(int)
	ret.Inst.WrongMessageType = d.Get("wrong_message_type").(int)
	ret.Inst.WrongSignatureLength = d.Get("wrong_signature_length").(int)
	ret.Inst.WrongSignatureSize = d.Get("wrong_signature_size").(int)
	ret.Inst.WrongSslVersion = d.Get("wrong_ssl_version").(int)
	ret.Inst.WrongVersionCounter = d.Get("wrong_version_counter").(int)
	ret.Inst.X509Lib = d.Get("x509_lib").(int)
	ret.Inst.X509VerificationSetupProblems = d.Get("x509_verification_setup_problems").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
