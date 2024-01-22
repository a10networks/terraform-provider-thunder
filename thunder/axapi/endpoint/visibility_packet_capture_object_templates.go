

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplates struct {
	Inst struct {

    AamAaaPolicyTmplList []VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList `json:"aam-aaa-policy-tmpl-list"`
    AamAuthCaptchaInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList `json:"aam-auth-captcha-inst-tmpl-list"`
    AamAuthLogonHttpInsTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList `json:"aam-auth-logon-http-ins-tmpl-list"`
    AamAuthRelayFormInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList `json:"aam-auth-relay-form-inst-tmpl-list"`
    AamAuthRelayHbaseInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList `json:"aam-auth-relay-hbase-inst-tmpl-list"`
    AamAuthRelayNtlmTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList `json:"aam-auth-relay-ntlm-tmpl-list"`
    AamAuthRelayWsFedTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList `json:"aam-auth-relay-ws-fed-tmpl-list"`
    AamAuthSamlIdProvTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList `json:"aam-auth-saml-id-prov-tmpl-list"`
    AamAuthSamlServiceProvTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList `json:"aam-auth-saml-service-prov-tmpl-list"`
    AamAuthServerLdapInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList `json:"aam-auth-server-ldap-inst-tmpl-list"`
    AamAuthServerOcspInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList `json:"aam-auth-server-ocsp-inst-tmpl-list"`
    AamAuthServerRadInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList `json:"aam-auth-server-rad-inst-tmpl-list"`
    AamAuthServerWinInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList `json:"aam-auth-server-win-inst-tmpl-list"`
    AamAuthServiceGroupMemTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList `json:"aam-auth-service-group-mem-tmpl-list"`
    AamAuthServiceGroupTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList `json:"aam-auth-service-group-tmpl-list"`
    AamJwtAuthorizationTmplList []VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList `json:"aam-jwt-authorization-tmpl-list"`
    Cgnv6Dns64VsPortTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList `json:"cgnv6-dns64-vs-port-tmpl-list"`
    Cgnv6EncapDomainTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList `json:"cgnv6-encap-domain-tmpl-list"`
    Cgnv6MapTransDomainTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList `json:"cgnv6-map-trans-domain-tmpl-list"`
    Cgnv6ServGroupTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList `json:"cgnv6-serv-group-tmpl-list"`
    Dns_vportTmplList []VisibilityPacketCaptureObjectTemplatesDns_vportTmplList `json:"dns_vport-tmpl-list"`
    FwServerPortTmplList []VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList `json:"fw-server-port-tmpl-list"`
    FwServiceGroupMemTmplList []VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList `json:"fw-service-group-mem-tmpl-list"`
    FwServiceGroupTmplList []VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList `json:"fw-service-group-tmpl-list"`
    ImapVportTmplList []VisibilityPacketCaptureObjectTemplatesImapVportTmplList `json:"imap-vport-tmpl-list"`
    InterfaceEthernetTmplList []VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList `json:"interface-ethernet-tmpl-list"`
    InterfaceTunnelTmplList []VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList `json:"interface-tunnel-tmpl-list"`
    NetflowMonitorTmplList []VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList `json:"netflow-monitor-tmpl-list"`
    Pop3VportTmplList []VisibilityPacketCaptureObjectTemplatesPop3VportTmplList `json:"pop3-vport-tmpl-list"`
    RuleSetTmplList []VisibilityPacketCaptureObjectTemplatesRuleSetTmplList `json:"rule-set-tmpl-list"`
    SlbPortTmplList []VisibilityPacketCaptureObjectTemplatesSlbPortTmplList `json:"slb-port-tmpl-list"`
    SlbTemplCacheTmplList []VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList `json:"slb-templ-cache-tmpl-list"`
    SlbVportTmplList []VisibilityPacketCaptureObjectTemplatesSlbVportTmplList `json:"slb-vport-tmpl-list"`
    SmtpVportTmplList []VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList `json:"smtp-vport-tmpl-list"`
    TemplGtpPlcyTmplList []VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList `json:"templ-gtp-plcy-tmpl-list"`
    Uuid string `json:"uuid"`

	} `json:"object-templates"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc struct {
    Error int `json:"error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Error int `json:"error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc struct {
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc struct {
    Spn_krb_faiure int `json:"spn_krb_faiure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Spn_krb_faiure int `json:"spn_krb_faiure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc struct {
    Invalid_srv_rsp int `json:"invalid_srv_rsp"`
    Post_fail int `json:"post_fail"`
    Invalid_cred int `json:"invalid_cred"`
    Bad_req int `json:"bad_req"`
    Not_fnd int `json:"not_fnd"`
    Error int `json:"error"`
    Other_error int `json:"other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Invalid_srv_rsp int `json:"invalid_srv_rsp"`
    Post_fail int `json:"post_fail"`
    Invalid_cred int `json:"invalid_cred"`
    Bad_req int `json:"bad_req"`
    Not_fnd int `json:"not_fnd"`
    Error int `json:"error"`
    Other_error int `json:"other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc struct {
    NoCreds int `json:"no-creds"`
    BadReq int `json:"bad-req"`
    Unauth int `json:"unauth"`
    Forbidden int `json:"forbidden"`
    NotFound int `json:"not-found"`
    ServerError int `json:"server-error"`
    Unavailable int `json:"unavailable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NoCreds int `json:"no-creds"`
    BadReq int `json:"bad-req"`
    Unauth int `json:"unauth"`
    Forbidden int `json:"forbidden"`
    NotFound int `json:"not-found"`
    ServerError int `json:"server-error"`
    Unavailable int `json:"unavailable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc struct {
    Failure int `json:"failure"`
    BufferAllocFail int `json:"buffer-alloc-fail"`
    EncodingFail int `json:"encoding-fail"`
    InsertHeaderFail int `json:"insert-header-fail"`
    ParseHeaderFail int `json:"parse-header-fail"`
    InternalError int `json:"internal-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure int `json:"failure"`
    BufferAllocFail int `json:"buffer-alloc-fail"`
    EncodingFail int `json:"encoding-fail"`
    InsertHeaderFail int `json:"insert-header-fail"`
    ParseHeaderFail int `json:"parse-header-fail"`
    InternalError int `json:"internal-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc struct {
    Failure int `json:"failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure int `json:"failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc struct {
    MdFail int `json:"md-fail"`
    AcsFail int `json:"acs-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MdFail int `json:"md-fail"`
    AcsFail int `json:"acs-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc struct {
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc struct {
    AdminBindFailure int `json:"admin-bind-failure"`
    BindFailure int `json:"bind-failure"`
    SearchFailure int `json:"search-failure"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    SslSessionFailure int `json:"ssl-session-failure"`
    Pw_change_failure int `json:"pw_change_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AdminBindFailure int `json:"admin-bind-failure"`
    BindFailure int `json:"bind-failure"`
    SearchFailure int `json:"search-failure"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    SslSessionFailure int `json:"ssl-session-failure"`
    Pw_change_failure int `json:"pw_change_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc struct {
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc struct {
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc struct {
    Krb_timeout_error int `json:"krb_timeout_error"`
    Krb_other_error int `json:"krb_other_error"`
    Krb_pw_expiry int `json:"krb_pw_expiry"`
    Krb_pw_change_failure int `json:"krb_pw_change_failure"`
    Ntlm_proto_nego_failure int `json:"ntlm_proto_nego_failure"`
    Ntlm_session_setup_failure int `json:"ntlm_session_setup_failure"`
    Ntlm_prepare_req_error int `json:"ntlm_prepare_req_error"`
    Ntlm_auth_failure int `json:"ntlm_auth_failure"`
    Ntlm_timeout_error int `json:"ntlm_timeout_error"`
    Ntlm_other_error int `json:"ntlm_other_error"`
    Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Krb_timeout_error int `json:"krb_timeout_error"`
    Krb_other_error int `json:"krb_other_error"`
    Krb_pw_expiry int `json:"krb_pw_expiry"`
    Krb_pw_change_failure int `json:"krb_pw_change_failure"`
    Ntlm_proto_nego_failure int `json:"ntlm_proto_nego_failure"`
    Ntlm_session_setup_failure int `json:"ntlm_session_setup_failure"`
    Ntlm_prepare_req_error int `json:"ntlm_prepare_req_error"`
    Ntlm_auth_failure int `json:"ntlm_auth_failure"`
    Ntlm_timeout_error int `json:"ntlm_timeout_error"`
    Ntlm_other_error int `json:"ntlm_other_error"`
    Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc struct {
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc struct {
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc struct {
    JwtAuthorizeFailure int `json:"jwt-authorize-failure"`
    JwtMissingToken int `json:"jwt-missing-token"`
    JwtMissingClaim int `json:"jwt-missing-claim"`
    JwtTokenExpired int `json:"jwt-token-expired"`
    JwtSignatureFailure int `json:"jwt-signature-failure"`
    JwtOtherError int `json:"jwt-other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    JwtAuthorizeFailure int `json:"jwt-authorize-failure"`
    JwtMissingToken int `json:"jwt-missing-token"`
    JwtMissingClaim int `json:"jwt-missing-claim"`
    JwtTokenExpired int `json:"jwt-token-expired"`
    JwtSignatureFailure int `json:"jwt-signature-failure"`
    JwtOtherError int `json:"jwt-other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc struct {
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc struct {
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc struct {
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc struct {
    Server_selection_fail_drop int `json:"server_selection_fail_drop"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_drop int `json:"server_selection_fail_drop"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc struct {
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Total_filter_drop int `json:"total_filter_drop"`
    Total_max_query_len_drop int `json:"total_max_query_len_drop"`
    Rcode_notimpl_receive int `json:"rcode_notimpl_receive"`
    Rcode_notimpl_response int `json:"rcode_notimpl_response"`
    Gslb_query_bad int `json:"gslb_query_bad"`
    Gslb_response_bad int `json:"gslb_response_bad"`
    Total_dns_filter_type_drop int `json:"total_dns_filter_type_drop"`
    Total_dns_filter_class_drop int `json:"total_dns_filter_class_drop"`
    Dns_filter_type_a_drop int `json:"dns_filter_type_a_drop"`
    Dns_filter_type_aaaa_drop int `json:"dns_filter_type_aaaa_drop"`
    Dns_filter_type_cname_drop int `json:"dns_filter_type_cname_drop"`
    Dns_filter_type_mx_drop int `json:"dns_filter_type_mx_drop"`
    Dns_filter_type_ns_drop int `json:"dns_filter_type_ns_drop"`
    Dns_filter_type_srv_drop int `json:"dns_filter_type_srv_drop"`
    Dns_filter_type_ptr_drop int `json:"dns_filter_type_ptr_drop"`
    Dns_filter_type_soa_drop int `json:"dns_filter_type_soa_drop"`
    Dns_filter_type_txt_drop int `json:"dns_filter_type_txt_drop"`
    Dns_filter_type_any_drop int `json:"dns_filter_type_any_drop"`
    Dns_filter_type_others_drop int `json:"dns_filter_type_others_drop"`
    Dns_filter_class_internet_drop int `json:"dns_filter_class_internet_drop"`
    Dns_filter_class_chaos_drop int `json:"dns_filter_class_chaos_drop"`
    Dns_filter_class_hesiod_drop int `json:"dns_filter_class_hesiod_drop"`
    Dns_filter_class_none_drop int `json:"dns_filter_class_none_drop"`
    Dns_filter_class_any_drop int `json:"dns_filter_class_any_drop"`
    Dns_filter_class_others_drop int `json:"dns_filter_class_others_drop"`
    Dns_rpz_action_drop int `json:"dns_rpz_action_drop"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Total_filter_drop int `json:"total_filter_drop"`
    Total_max_query_len_drop int `json:"total_max_query_len_drop"`
    Rcode_notimpl_receive int `json:"rcode_notimpl_receive"`
    Rcode_notimpl_response int `json:"rcode_notimpl_response"`
    Gslb_query_bad int `json:"gslb_query_bad"`
    Gslb_response_bad int `json:"gslb_response_bad"`
    Total_dns_filter_type_drop int `json:"total_dns_filter_type_drop"`
    Total_dns_filter_class_drop int `json:"total_dns_filter_class_drop"`
    Dns_filter_type_a_drop int `json:"dns_filter_type_a_drop"`
    Dns_filter_type_aaaa_drop int `json:"dns_filter_type_aaaa_drop"`
    Dns_filter_type_cname_drop int `json:"dns_filter_type_cname_drop"`
    Dns_filter_type_mx_drop int `json:"dns_filter_type_mx_drop"`
    Dns_filter_type_ns_drop int `json:"dns_filter_type_ns_drop"`
    Dns_filter_type_srv_drop int `json:"dns_filter_type_srv_drop"`
    Dns_filter_type_ptr_drop int `json:"dns_filter_type_ptr_drop"`
    Dns_filter_type_soa_drop int `json:"dns_filter_type_soa_drop"`
    Dns_filter_type_txt_drop int `json:"dns_filter_type_txt_drop"`
    Dns_filter_type_any_drop int `json:"dns_filter_type_any_drop"`
    Dns_filter_type_others_drop int `json:"dns_filter_type_others_drop"`
    Dns_filter_class_internet_drop int `json:"dns_filter_class_internet_drop"`
    Dns_filter_class_chaos_drop int `json:"dns_filter_class_chaos_drop"`
    Dns_filter_class_hesiod_drop int `json:"dns_filter_class_hesiod_drop"`
    Dns_filter_class_none_drop int `json:"dns_filter_class_none_drop"`
    Dns_filter_class_any_drop int `json:"dns_filter_class_any_drop"`
    Dns_filter_class_others_drop int `json:"dns_filter_class_others_drop"`
    Dns_rpz_action_drop int `json:"dns_rpz_action_drop"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc struct {
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc struct {
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc struct {
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesImapVportTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Auth_unsupported int `json:"auth_unsupported"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Auth_unsupported int `json:"auth_unsupported"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc struct {
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc struct {
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc struct {
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    SessionEventNat44RecordsSentFailur int `json:"session-event-nat44-records-sent-failur"`
    SessionEventNat64RecordsSentFailur int `json:"session-event-nat64-records-sent-failur"`
    SessionEventDsliteRecordsSentFailu int `json:"session-event-dslite-records-sent-failu"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    PortMappingDsliteRecordsSentFailur int `json:"port-mapping-dslite-records-sent-failur"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    PortBatchingNat44RecordsSentFailur int `json:"port-batching-nat44-records-sent-failur"`
    PortBatchingNat64RecordsSentFailur int `json:"port-batching-nat64-records-sent-failur"`
    PortBatchingDsliteRecordsSentFailu int `json:"port-batching-dslite-records-sent-failu"`
    PortBatchingV2Nat44RecordsSentFai int `json:"port-batching-v2-nat44-records-sent-fai"`
    PortBatchingV2Nat64RecordsSentFai int `json:"port-batching-v2-nat64-records-sent-fai"`
    PortBatchingV2DsliteRecordsSentFa int `json:"port-batching-v2-dslite-records-sent-fa"`
    CustomSessionEventNat44CreationRec int `json:"custom-session-event-nat44-creation-rec"`
    CustomSessionEventNat64CreationRec int `json:"custom-session-event-nat64-creation-rec"`
    CustomSessionEventDsliteCreationRe int `json:"custom-session-event-dslite-creation-re"`
    CustomSessionEventNat44DeletionRec int `json:"custom-session-event-nat44-deletion-rec"`
    CustomSessionEventNat64DeletionRec int `json:"custom-session-event-nat64-deletion-rec"`
    CustomSessionEventDsliteDeletionRe int `json:"custom-session-event-dslite-deletion-re"`
    CustomSessionEventFw4CreationRecor int `json:"custom-session-event-fw4-creation-recor"`
    CustomSessionEventFw6CreationRecor int `json:"custom-session-event-fw6-creation-recor"`
    CustomSessionEventFw4DeletionRecor int `json:"custom-session-event-fw4-deletion-recor"`
    CustomSessionEventFw6DeletionRecor int `json:"custom-session-event-fw6-deletion-recor"`
    CustomDenyResetEventFw4RecordsSen int `json:"custom-deny-reset-event-fw4-records-sen"`
    CustomDenyResetEventFw6RecordsSen int `json:"custom-deny-reset-event-fw6-records-sen"`
    CustomPortMappingNat44CreationReco int `json:"custom-port-mapping-nat44-creation-reco"`
    CustomPortMappingNat64CreationReco int `json:"custom-port-mapping-nat64-creation-reco"`
    CustomPortMappingDsliteCreationRec int `json:"custom-port-mapping-dslite-creation-rec"`
    CustomPortMappingNat44DeletionReco int `json:"custom-port-mapping-nat44-deletion-reco"`
    CustomPortMappingNat64DeletionReco int `json:"custom-port-mapping-nat64-deletion-reco"`
    CustomPortMappingDsliteDeletionRec int `json:"custom-port-mapping-dslite-deletion-rec"`
    CustomPortBatchingNat44CreationRec int `json:"custom-port-batching-nat44-creation-rec"`
    CustomPortBatchingNat64CreationRec int `json:"custom-port-batching-nat64-creation-rec"`
    CustomPortBatchingDsliteCreationRe int `json:"custom-port-batching-dslite-creation-re"`
    CustomPortBatchingNat44DeletionRec int `json:"custom-port-batching-nat44-deletion-rec"`
    CustomPortBatchingNat64DeletionRec int `json:"custom-port-batching-nat64-deletion-rec"`
    CustomPortBatchingDsliteDeletionRe int `json:"custom-port-batching-dslite-deletion-re"`
    CustomPortBatchingV2Nat44Creation int `json:"custom-port-batching-v2-nat44-creation-"`
    CustomPortBatchingV2Nat64Creation int `json:"custom-port-batching-v2-nat64-creation-"`
    CustomPortBatchingV2DsliteCreation int `json:"custom-port-batching-v2-dslite-creation"`
    CustomPortBatchingV2Nat44Deletion int `json:"custom-port-batching-v2-nat44-deletion-"`
    CustomPortBatchingV2Nat64Deletion int `json:"custom-port-batching-v2-nat64-deletion-"`
    CustomPortBatchingV2DsliteDeletion int `json:"custom-port-batching-v2-dslite-deletion"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent-"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent-"`
    CustomGtpDenyEventRecordsSentFail int `json:"custom-gtp-deny-event-records-sent-fail"`
    CustomGtpInfoEventRecordsSentFail int `json:"custom-gtp-info-event-records-sent-fail"`
    CustomFwIddosEntryCreatedRecordsS int `json:"custom-fw-iddos-entry-created-records-s"`
    CustomFwIddosEntryDeletedRecordsS int `json:"custom-fw-iddos-entry-deleted-records-s"`
    CustomFwSesnLimitExceededRecordsS int `json:"custom-fw-sesn-limit-exceeded-records-s"`
    CustomNatIddosL3EntryCreatedRecor int `json:"custom-nat-iddos-l3-entry-created-recor"`
    CustomNatIddosL3EntryDeletedRecor int `json:"custom-nat-iddos-l3-entry-deleted-recor"`
    CustomNatIddosL4EntryCreatedRecor int `json:"custom-nat-iddos-l4-entry-created-recor"`
    CustomNatIddosL4EntryDeletedRecor int `json:"custom-nat-iddos-l4-entry-deleted-recor"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    SessionEventNat44RecordsSentFailur int `json:"session-event-nat44-records-sent-failur"`
    SessionEventNat64RecordsSentFailur int `json:"session-event-nat64-records-sent-failur"`
    SessionEventDsliteRecordsSentFailu int `json:"session-event-dslite-records-sent-failu"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    PortMappingDsliteRecordsSentFailur int `json:"port-mapping-dslite-records-sent-failur"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    PortBatchingNat44RecordsSentFailur int `json:"port-batching-nat44-records-sent-failur"`
    PortBatchingNat64RecordsSentFailur int `json:"port-batching-nat64-records-sent-failur"`
    PortBatchingDsliteRecordsSentFailu int `json:"port-batching-dslite-records-sent-failu"`
    PortBatchingV2Nat44RecordsSentFai int `json:"port-batching-v2-nat44-records-sent-fai"`
    PortBatchingV2Nat64RecordsSentFai int `json:"port-batching-v2-nat64-records-sent-fai"`
    PortBatchingV2DsliteRecordsSentFa int `json:"port-batching-v2-dslite-records-sent-fa"`
    CustomSessionEventNat44CreationRec int `json:"custom-session-event-nat44-creation-rec"`
    CustomSessionEventNat64CreationRec int `json:"custom-session-event-nat64-creation-rec"`
    CustomSessionEventDsliteCreationRe int `json:"custom-session-event-dslite-creation-re"`
    CustomSessionEventNat44DeletionRec int `json:"custom-session-event-nat44-deletion-rec"`
    CustomSessionEventNat64DeletionRec int `json:"custom-session-event-nat64-deletion-rec"`
    CustomSessionEventDsliteDeletionRe int `json:"custom-session-event-dslite-deletion-re"`
    CustomSessionEventFw4CreationRecor int `json:"custom-session-event-fw4-creation-recor"`
    CustomSessionEventFw6CreationRecor int `json:"custom-session-event-fw6-creation-recor"`
    CustomSessionEventFw4DeletionRecor int `json:"custom-session-event-fw4-deletion-recor"`
    CustomSessionEventFw6DeletionRecor int `json:"custom-session-event-fw6-deletion-recor"`
    CustomDenyResetEventFw4RecordsSen int `json:"custom-deny-reset-event-fw4-records-sen"`
    CustomDenyResetEventFw6RecordsSen int `json:"custom-deny-reset-event-fw6-records-sen"`
    CustomPortMappingNat44CreationReco int `json:"custom-port-mapping-nat44-creation-reco"`
    CustomPortMappingNat64CreationReco int `json:"custom-port-mapping-nat64-creation-reco"`
    CustomPortMappingDsliteCreationRec int `json:"custom-port-mapping-dslite-creation-rec"`
    CustomPortMappingNat44DeletionReco int `json:"custom-port-mapping-nat44-deletion-reco"`
    CustomPortMappingNat64DeletionReco int `json:"custom-port-mapping-nat64-deletion-reco"`
    CustomPortMappingDsliteDeletionRec int `json:"custom-port-mapping-dslite-deletion-rec"`
    CustomPortBatchingNat44CreationRec int `json:"custom-port-batching-nat44-creation-rec"`
    CustomPortBatchingNat64CreationRec int `json:"custom-port-batching-nat64-creation-rec"`
    CustomPortBatchingDsliteCreationRe int `json:"custom-port-batching-dslite-creation-re"`
    CustomPortBatchingNat44DeletionRec int `json:"custom-port-batching-nat44-deletion-rec"`
    CustomPortBatchingNat64DeletionRec int `json:"custom-port-batching-nat64-deletion-rec"`
    CustomPortBatchingDsliteDeletionRe int `json:"custom-port-batching-dslite-deletion-re"`
    CustomPortBatchingV2Nat44Creation int `json:"custom-port-batching-v2-nat44-creation-"`
    CustomPortBatchingV2Nat64Creation int `json:"custom-port-batching-v2-nat64-creation-"`
    CustomPortBatchingV2DsliteCreation int `json:"custom-port-batching-v2-dslite-creation"`
    CustomPortBatchingV2Nat44Deletion int `json:"custom-port-batching-v2-nat44-deletion-"`
    CustomPortBatchingV2Nat64Deletion int `json:"custom-port-batching-v2-nat64-deletion-"`
    CustomPortBatchingV2DsliteDeletion int `json:"custom-port-batching-v2-dslite-deletion"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent-"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent-"`
    CustomGtpDenyEventRecordsSentFail int `json:"custom-gtp-deny-event-records-sent-fail"`
    CustomGtpInfoEventRecordsSentFail int `json:"custom-gtp-info-event-records-sent-fail"`
    CustomFwIddosEntryCreatedRecordsS int `json:"custom-fw-iddos-entry-created-records-s"`
    CustomFwIddosEntryDeletedRecordsS int `json:"custom-fw-iddos-entry-deleted-records-s"`
    CustomFwSesnLimitExceededRecordsS int `json:"custom-fw-sesn-limit-exceeded-records-s"`
    CustomNatIddosL3EntryCreatedRecor int `json:"custom-nat-iddos-l3-entry-created-recor"`
    CustomNatIddosL3EntryDeletedRecor int `json:"custom-nat-iddos-l3-entry-deleted-recor"`
    CustomNatIddosL4EntryCreatedRecor int `json:"custom-nat-iddos-l4-entry-created-recor"`
    CustomNatIddosL4EntryDeletedRecor int `json:"custom-nat-iddos-l4-entry-deleted-recor"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc struct {
    UnmatchedDrops int `json:"unmatched-drops"`
    Deny int `json:"deny"`
    Reset int `json:"reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    UnmatchedDrops int `json:"unmatched-drops"`
    Deny int `json:"deny"`
    Reset int `json:"reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc struct {
    Es_resp_300 int `json:"es_resp_300"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Resp3xx int `json:"resp-3xx"`
    Resp4xx int `json:"resp-4xx"`
    Resp5xx int `json:"resp-5xx"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_resp_300 int `json:"es_resp_300"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Resp3xx int `json:"resp-3xx"`
    Resp4xx int `json:"resp-4xx"`
    Resp5xx int `json:"resp-5xx"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc struct {
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc struct {
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Loc_deny int `json:"loc_deny"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Loc_deny int `json:"loc_deny"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc struct {
    No_proxy int `json:"no_proxy"`
    Parse_req_fail int `json:"parse_req_fail"`
    Server_select_fail int `json:"server_select_fail"`
    Forward_req_fail int `json:"forward_req_fail"`
    Forward_req_data_fail int `json:"forward_req_data_fail"`
    Snat_fail int `json:"snat_fail"`
    Send_client_service_not_ready int `json:"send_client_service_not_ready"`
    Recv_server_unknow_reply_code int `json:"recv_server_unknow_reply_code"`
    Read_request_line_fail int `json:"read_request_line_fail"`
    Get_all_headers_fail int `json:"get_all_headers_fail"`
    Too_many_headers int `json:"too_many_headers"`
    Line_too_long int `json:"line_too_long"`
    Line_extend_fail int `json:"line_extend_fail"`
    Line_table_extend_fail int `json:"line_table_extend_fail"`
    Parse_request_line_fail int `json:"parse_request_line_fail"`
    Insert_resonse_line_fail int `json:"insert_resonse_line_fail"`
    Remove_resonse_line_fail int `json:"remove_resonse_line_fail"`
    Parse_resonse_line_fail int `json:"parse_resonse_line_fail"`
    Server_starttls_fail int `json:"server_STARTTLS_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    No_proxy int `json:"no_proxy"`
    Parse_req_fail int `json:"parse_req_fail"`
    Server_select_fail int `json:"server_select_fail"`
    Forward_req_fail int `json:"forward_req_fail"`
    Forward_req_data_fail int `json:"forward_req_data_fail"`
    Snat_fail int `json:"snat_fail"`
    Send_client_service_not_ready int `json:"send_client_service_not_ready"`
    Recv_server_unknow_reply_code int `json:"recv_server_unknow_reply_code"`
    Read_request_line_fail int `json:"read_request_line_fail"`
    Get_all_headers_fail int `json:"get_all_headers_fail"`
    Too_many_headers int `json:"too_many_headers"`
    Line_too_long int `json:"line_too_long"`
    Line_extend_fail int `json:"line_extend_fail"`
    Line_table_extend_fail int `json:"line_table_extend_fail"`
    Parse_request_line_fail int `json:"parse_request_line_fail"`
    Insert_resonse_line_fail int `json:"insert_resonse_line_fail"`
    Remove_resonse_line_fail int `json:"remove_resonse_line_fail"`
    Parse_resonse_line_fail int `json:"parse_resonse_line_fail"`
    Server_starttls_fail int `json:"server_STARTTLS_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc struct {
    DropVldGtpIeRepeatCountExceed int `json:"drop-vld-gtp-ie-repeat-count-exceed"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
    DropVldMandatoryIeInGroupedIe int `json:"drop-vld-mandatory-ie-in-grouped-ie"`
    DropVldOutOfOrderIe int `json:"drop-vld-out-of-order-ie"`
    DropVldOutOfStateIe int `json:"drop-vld-out-of-state-ie"`
    DropVldReservedInformationElement int `json:"drop-vld-reserved-information-element"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    DropVldMessageLength int `json:"drop-vld-message-length"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldGtpUSpoofedSourceAddress int `json:"drop-vld-gtp-u-spoofed-source-address"`
    DropVldGtpBearerCountExceed int `json:"drop-vld-gtp-bearer-count-exceed"`
    DropVldGtpV2WrongLbiCreateBearer int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
    DropVldV0ReservedMessageDrop int `json:"drop-vld-v0-reserved-message-drop"`
    DropVldV1ReservedMessageDrop int `json:"drop-vld-v1-reserved-message-drop"`
    DropVldV2ReservedMessageDrop int `json:"drop-vld-v2-reserved-message-drop"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    DropVldGtpIeRepeatCountExceed int `json:"drop-vld-gtp-ie-repeat-count-exceed"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
    DropVldMandatoryIeInGroupedIe int `json:"drop-vld-mandatory-ie-in-grouped-ie"`
    DropVldOutOfOrderIe int `json:"drop-vld-out-of-order-ie"`
    DropVldOutOfStateIe int `json:"drop-vld-out-of-state-ie"`
    DropVldReservedInformationElement int `json:"drop-vld-reserved-information-element"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    DropVldMessageLength int `json:"drop-vld-message-length"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldGtpUSpoofedSourceAddress int `json:"drop-vld-gtp-u-spoofed-source-address"`
    DropVldGtpBearerCountExceed int `json:"drop-vld-gtp-bearer-count-exceed"`
    DropVldGtpV2WrongLbiCreateBearer int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
    DropVldV0ReservedMessageDrop int `json:"drop-vld-v0-reserved-message-drop"`
    DropVldV1ReservedMessageDrop int `json:"drop-vld-v1-reserved-message-drop"`
    DropVldV2ReservedMessageDrop int `json:"drop-vld-v2-reserved-message-drop"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureObjectTemplates) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplates) getPath() string{
    return "visibility/packet-capture/object-templates"
}

func (p *VisibilityPacketCaptureObjectTemplates) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplates::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplates) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplates::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *VisibilityPacketCaptureObjectTemplates) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplates::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplates) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplates::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
