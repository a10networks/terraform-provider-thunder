

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslErrorStats struct {
    
    Stats SlbSslErrorStatsStats `json:"stats"`

}
type DataSlbSslErrorStats struct {
    DtSlbSslErrorStats SlbSslErrorStats `json:"ssl-error"`
}


type SlbSslErrorStatsStats struct {
    AppDataInHandshake int `json:"app-data-in-handshake"`
    AttemptToReuseSessInDiffContext int `json:"attempt-to-reuse-sess-in-diff-context"`
    BadAlertRecord int `json:"bad-alert-record"`
    BadAuthenticationType int `json:"bad-authentication-type"`
    BadChangeCipherSpec int `json:"bad-change-cipher-spec"`
    BadChecksum int `json:"bad-checksum"`
    BadDataReturnedByCallback int `json:"bad-data-returned-by-callback"`
    BadDecompression int `json:"bad-decompression"`
    BadDhGLength int `json:"bad-dh-g-length"`
    BadDhPubKeyLength int `json:"bad-dh-pub-key-length"`
    BadDhPLength int `json:"bad-dh-p-length"`
    BadDigestLength int `json:"bad-digest-length"`
    BadDsaSignature int `json:"bad-dsa-signature"`
    BadHelloRequest int `json:"bad-hello-request"`
    BadLength int `json:"bad-length"`
    BadMacDecode int `json:"bad-mac-decode"`
    BadMessageType int `json:"bad-message-type"`
    BadPacketLength int `json:"bad-packet-length"`
    BadProtocolVersionCounter int `json:"bad-protocol-version-counter"`
    BadResponseArgument int `json:"bad-response-argument"`
    BadRsaDecrypt int `json:"bad-rsa-decrypt"`
    BadRsaEncrypt int `json:"bad-rsa-encrypt"`
    BadRsaELength int `json:"bad-rsa-e-length"`
    BadRsaModulusLength int `json:"bad-rsa-modulus-length"`
    BadRsaSignature int `json:"bad-rsa-signature"`
    BadSignature int `json:"bad-signature"`
    BadSslFiletype int `json:"bad-ssl-filetype"`
    BadSslSessionIdLength int `json:"bad-ssl-session-id-length"`
    BadState int `json:"bad-state"`
    BadWriteRetry int `json:"bad-write-retry"`
    BioNotSet int `json:"bio-not-set"`
    BlockCipherPadIsWrong int `json:"block-cipher-pad-is-wrong"`
    BnLib int `json:"bn-lib"`
    CaDnLengthMismatch int `json:"ca-dn-length-mismatch"`
    CaDnTooLong int `json:"ca-dn-too-long"`
    CcsReceivedEarly int `json:"ccs-received-early"`
    CertificateVerifyFailed int `json:"certificate-verify-failed"`
    CertLengthMismatch int `json:"cert-length-mismatch"`
    ChallengeIsDifferent int `json:"challenge-is-different"`
    CipherCodeWrongLength int `json:"cipher-code-wrong-length"`
    CipherOrHashUnavailable int `json:"cipher-or-hash-unavailable"`
    CipherTableSrcError int `json:"cipher-table-src-error"`
    CompressedLengthTooLong int `json:"compressed-length-too-long"`
    CompressionFailure int `json:"compression-failure"`
    CompressionLibraryError int `json:"compression-library-error"`
    ConnectionIdIsDifferent int `json:"connection-id-is-different"`
    ConnectionTypeNotSet int `json:"connection-type-not-set"`
    DataBetweenCcsAndFinished int `json:"data-between-ccs-and-finished"`
    DataLengthTooLong int `json:"data-length-too-long"`
    DecryptionFailed int `json:"decryption-failed"`
    DecryptionFailedOrBadRecordMac int `json:"decryption-failed-or-bad-record-mac"`
    DhPublicValueLengthIsWrong int `json:"dh-public-value-length-is-wrong"`
    DigestCheckFailed int `json:"digest-check-failed"`
    EncryptedLengthTooLong int `json:"encrypted-length-too-long"`
    ErrorGeneratingTmpRsaKey int `json:"error-generating-tmp-rsa-key"`
    ErrorInReceivedCipherList int `json:"error-in-received-cipher-list"`
    ExcessiveMessageSize int `json:"excessive-message-size"`
    ExtraDataInMessage int `json:"extra-data-in-message"`
    GotAFinBeforeACcs int `json:"got-a-fin-before-a-ccs"`
    HttpsProxyRequest int `json:"https-proxy-request"`
    HttpRequest int `json:"http-request"`
    IllegalPadding int `json:"illegal-padding"`
    InappropriateFallback int `json:"inappropriate-fallback"`
    InvalidChallengeLength int `json:"invalid-challenge-length"`
    InvalidCommand int `json:"invalid-command"`
    InvalidPurpose int `json:"invalid-purpose"`
    InvalidStatusResponse int `json:"invalid-status-response"`
    InvalidTrust int `json:"invalid-trust"`
    KeyArgTooLong int `json:"key-arg-too-long"`
    Krb5 int `json:"krb5"`
    Krb5ClientCcPrincipal int `json:"krb5-client-cc-principal"`
    Krb5ClientGetCred int `json:"krb5-client-get-cred"`
    Krb5ClientInit int `json:"krb5-client-init"`
    Krb5ClientMkReq int `json:"krb5-client-mk-req"`
    Krb5ServerBadTicket int `json:"krb5-server-bad-ticket"`
    Krb5ServerInit int `json:"krb5-server-init"`
    Krb5ServerRdReq int `json:"krb5-server-rd-req"`
    Krb5ServerTktExpired int `json:"krb5-server-tkt-expired"`
    Krb5ServerTktNotYetValid int `json:"krb5-server-tkt-not-yet-valid"`
    Krb5ServerTktSkew int `json:"krb5-server-tkt-skew"`
    LengthMismatch int `json:"length-mismatch"`
    LengthTooShort int `json:"length-too-short"`
    LibraryBug int `json:"library-bug"`
    LibraryHasNoCiphers int `json:"library-has-no-ciphers"`
    MastKeyTooLong int `json:"mast-key-too-long"`
    MessageTooLong int `json:"message-too-long"`
    MissingDhDsaCert int `json:"missing-dh-dsa-cert"`
    MissingDhKey int `json:"missing-dh-key"`
    MissingDhRsaCert int `json:"missing-dh-rsa-cert"`
    MissingDsaSigningCert int `json:"missing-dsa-signing-cert"`
    MissingExportTmpDhKey int `json:"missing-export-tmp-dh-key"`
    MissingExportTmpRsaKey int `json:"missing-export-tmp-rsa-key"`
    MissingRsaCertificate int `json:"missing-rsa-certificate"`
    MissingRsaEncryptingCert int `json:"missing-rsa-encrypting-cert"`
    MissingRsaSigningCert int `json:"missing-rsa-signing-cert"`
    MissingTmpDhKey int `json:"missing-tmp-dh-key"`
    MissingTmpRsaKey int `json:"missing-tmp-rsa-key"`
    MissingTmpRsaPkey int `json:"missing-tmp-rsa-pkey"`
    MissingVerifyMessage int `json:"missing-verify-message"`
    NonSslv2InitialPacket int `json:"non-sslv2-initial-packet"`
    NoCertificatesReturned int `json:"no-certificates-returned"`
    NoCertificateAssigned int `json:"no-certificate-assigned"`
    NoCertificateReturned int `json:"no-certificate-returned"`
    NoCertificateSet int `json:"no-certificate-set"`
    NoCertificateSpecified int `json:"no-certificate-specified"`
    NoCiphersAvailable int `json:"no-ciphers-available"`
    NoCiphersPassed int `json:"no-ciphers-passed"`
    NoCiphersSpecified int `json:"no-ciphers-specified"`
    NoCipherList int `json:"no-cipher-list"`
    NoCipherMatch int `json:"no-cipher-match"`
    NoClientCertReceived int `json:"no-client-cert-received"`
    NoCompressionSpecified int `json:"no-compression-specified"`
    NoMethodSpecified int `json:"no-method-specified"`
    NoPrivatekey int `json:"no-privatekey"`
    NoPrivateKeyAssigned int `json:"no-private-key-assigned"`
    NoProtocolsAvailable int `json:"no-protocols-available"`
    NoPublickey int `json:"no-publickey"`
    NoSharedCipher int `json:"no-shared-cipher"`
    NoVerifyCallback int `json:"no-verify-callback"`
    NullSslCtx int `json:"null-ssl-ctx"`
    NullSslMethodPassed int `json:"null-ssl-method-passed"`
    OldSessionCipherNotReturned int `json:"old-session-cipher-not-returned"`
    PacketLengthTooLong int `json:"packet-length-too-long"`
    PathTooLong int `json:"path-too-long"`
    PeerDidNotReturnACertificate int `json:"peer-did-not-return-a-certificate"`
    PeerError int `json:"peer-error"`
    PeerErrorCertificate int `json:"peer-error-certificate"`
    PeerErrorNoCertificate int `json:"peer-error-no-certificate"`
    PeerErrorNoCipher int `json:"peer-error-no-cipher"`
    PeerErrorUnsupportedCertificateType int `json:"peer-error-unsupported-certificate-type"`
    PreMacLengthTooLong int `json:"pre-mac-length-too-long"`
    ProblemsMappingCipherFunctions int `json:"problems-mapping-cipher-functions"`
    ProtocolIsShutdown int `json:"protocol-is-shutdown"`
    PublicKeyEncryptError int `json:"public-key-encrypt-error"`
    PublicKeyIsNotRsa int `json:"public-key-is-not-rsa"`
    PublicKeyNotRsa int `json:"public-key-not-rsa"`
    ReadBioNotSet int `json:"read-bio-not-set"`
    ReadWrongPacketType int `json:"read-wrong-packet-type"`
    RecordLengthMismatch int `json:"record-length-mismatch"`
    RecordTooLarge int `json:"record-too-large"`
    RecordTooSmall int `json:"record-too-small"`
    RequiredCipherMissing int `json:"required-cipher-missing"`
    ReuseCertLengthNotZero int `json:"reuse-cert-length-not-zero"`
    ReuseCertTypeNotZero int `json:"reuse-cert-type-not-zero"`
    ReuseCipherListNotZero int `json:"reuse-cipher-list-not-zero"`
    ScsvReceivedWhenRenegotiating int `json:"scsv-received-when-renegotiating"`
    SessionIdContextUninitialized int `json:"session-id-context-uninitialized"`
    ShortRead int `json:"short-read"`
    SignatureForNonSigningCertificate int `json:"signature-for-non-signing-certificate"`
    Ssl23DoingSessionIdReuse int `json:"ssl23-doing-session-id-reuse"`
    Ssl2ConnectionIdTooLong int `json:"ssl2-connection-id-too-long"`
    Ssl3SessionIdTooLong int `json:"ssl3-session-id-too-long"`
    Ssl3SessionIdTooShort int `json:"ssl3-session-id-too-short"`
    Sslv3AlertBadCertificate int `json:"sslv3-alert-bad-certificate"`
    Sslv3AlertBadRecordMac int `json:"sslv3-alert-bad-record-mac"`
    Sslv3AlertCertificateExpired int `json:"sslv3-alert-certificate-expired"`
    Sslv3AlertCertificateRevoked int `json:"sslv3-alert-certificate-revoked"`
    Sslv3AlertCertificateUnknown int `json:"sslv3-alert-certificate-unknown"`
    Sslv3AlertDecompressionFailure int `json:"sslv3-alert-decompression-failure"`
    Sslv3AlertHandshakeFailure int `json:"sslv3-alert-handshake-failure"`
    Sslv3AlertIllegalParameter int `json:"sslv3-alert-illegal-parameter"`
    Sslv3AlertNoCertificate int `json:"sslv3-alert-no-certificate"`
    Sslv3AlertPeerErrorCert int `json:"sslv3-alert-peer-error-cert"`
    Sslv3AlertPeerErrorNoCert int `json:"sslv3-alert-peer-error-no-cert"`
    Sslv3AlertPeerErrorNoCipher int `json:"sslv3-alert-peer-error-no-cipher"`
    Sslv3AlertPeerErrorUnsuppCertType int `json:"sslv3-alert-peer-error-unsupp-cert-type"`
    Sslv3AlertUnexpectedMsg int `json:"sslv3-alert-unexpected-msg"`
    Sslv3AlertUnknownRemoteErrType int `json:"sslv3-alert-unknown-remote-err-type"`
    Sslv3AlertUnspportedCert int `json:"sslv3-alert-unspported-cert"`
    SslCtxHasNoDefaultSslVersion int `json:"ssl-ctx-has-no-default-ssl-version"`
    SslHandshakeFailure int `json:"ssl-handshake-failure"`
    SslLibraryHasNoCiphers int `json:"ssl-library-has-no-ciphers"`
    SslSessionIdCallbackFailed int `json:"ssl-session-id-callback-failed"`
    SslSessionIdConflict int `json:"ssl-session-id-conflict"`
    SslSessionIdContextTooLong int `json:"ssl-session-id-context-too-long"`
    SslSessionIdHasBadLength int `json:"ssl-session-id-has-bad-length"`
    SslSessionIdIsDifferent int `json:"ssl-session-id-is-different"`
    Tlsv1AlertAccessDenied int `json:"tlsv1-alert-access-denied"`
    Tlsv1AlertDecodeError int `json:"tlsv1-alert-decode-error"`
    Tlsv1AlertDecryptionFailed int `json:"tlsv1-alert-decryption-failed"`
    Tlsv1AlertDecryptError int `json:"tlsv1-alert-decrypt-error"`
    Tlsv1AlertExportRestriction int `json:"tlsv1-alert-export-restriction"`
    Tlsv1AlertInsufficientSecurity int `json:"tlsv1-alert-insufficient-security"`
    Tlsv1AlertInternalError int `json:"tlsv1-alert-internal-error"`
    Tlsv1AlertNoRenegotiation int `json:"tlsv1-alert-no-renegotiation"`
    Tlsv1AlertProtocolVersion int `json:"tlsv1-alert-protocol-version"`
    Tlsv1AlertRecordOverflow int `json:"tlsv1-alert-record-overflow"`
    Tlsv1AlertUnknownCa int `json:"tlsv1-alert-unknown-ca"`
    Tlsv1AlertUserCancelled int `json:"tlsv1-alert-user-cancelled"`
    TlsClientCertReqWithAnonCipher int `json:"tls-client-cert-req-with-anon-cipher"`
    TlsPeerDidNotRespondWithCertList int `json:"tls-peer-did-not-respond-with-cert-list"`
    TlsRsaEncryptedValueLengthIsWrong int `json:"tls-rsa-encrypted-value-length-is-wrong"`
    TriedToUseUnsupportedCipher int `json:"tried-to-use-unsupported-cipher"`
    UnableToDecodeDhCerts int `json:"unable-to-decode-dh-certs"`
    UnableToExtractPublicKey int `json:"unable-to-extract-public-key"`
    UnableToFindDhParameters int `json:"unable-to-find-dh-parameters"`
    UnableToFindPublicKeyParameters int `json:"unable-to-find-public-key-parameters"`
    UnableToFindSslMethod int `json:"unable-to-find-ssl-method"`
    UnableToLoadSsl2Md5Routines int `json:"unable-to-load-ssl2-md5-routines"`
    UnableToLoadSsl3Md5Routines int `json:"unable-to-load-ssl3-md5-routines"`
    UnableToLoadSsl3Sha1Routines int `json:"unable-to-load-ssl3-sha1-routines"`
    UnexpectedMessage int `json:"unexpected-message"`
    UnexpectedRecord int `json:"unexpected-record"`
    Uninitialized int `json:"uninitialized"`
    UnknownAlertType int `json:"unknown-alert-type"`
    UnknownCertificateType int `json:"unknown-certificate-type"`
    UnknownCipherReturned int `json:"unknown-cipher-returned"`
    UnknownCipherType int `json:"unknown-cipher-type"`
    UnknownKeyExchangeType int `json:"unknown-key-exchange-type"`
    UnknownPkeyType int `json:"unknown-pkey-type"`
    UnknownProtocol int `json:"unknown-protocol"`
    UnknownRemoteErrorType int `json:"unknown-remote-error-type"`
    UnknownSslVersion int `json:"unknown-ssl-version"`
    UnknownState int `json:"unknown-state"`
    UnsupportedCipher int `json:"unsupported-cipher"`
    UnsupportedCompressionAlgorithm int `json:"unsupported-compression-algorithm"`
    UnsupportedOption int `json:"unsupported-option"`
    UnsupportedProtocol int `json:"unsupported-protocol"`
    UnsupportedSslVersion int `json:"unsupported-ssl-version"`
    UnsupportedStatusType int `json:"unsupported-status-type"`
    WriteBioNotSet int `json:"write-bio-not-set"`
    WrongCipherReturned int `json:"wrong-cipher-returned"`
    WrongMessageType int `json:"wrong-message-type"`
    WrongCounterOfKeyBits int `json:"wrong-counter-of-key-bits"`
    WrongSignatureLength int `json:"wrong-signature-length"`
    WrongSignatureSize int `json:"wrong-signature-size"`
    WrongSslVersion int `json:"wrong-ssl-version"`
    WrongVersionCounter int `json:"wrong-version-counter"`
    X509Lib int `json:"x509-lib"`
    X509VerificationSetupProblems int `json:"x509-verification-setup-problems"`
    ClienthelloTlsext int `json:"clienthello-tlsext"`
    ParseTlsext int `json:"parse-tlsext"`
    ServerhelloTlsext int `json:"serverhello-tlsext"`
    Ssl3ExtInvalidServername int `json:"ssl3-ext-invalid-servername"`
    Ssl3ExtInvalidServernameType int `json:"ssl3-ext-invalid-servername-type"`
    MultipleSgcRestarts int `json:"multiple-sgc-restarts"`
    TlsInvalidEcpointformatList int `json:"tls-invalid-ecpointformat-list"`
    BadEccCert int `json:"bad-ecc-cert"`
    BadEcdsaSig int `json:"bad-ecdsa-sig"`
    BadEcpoint int `json:"bad-ecpoint"`
    CookieMismatch int `json:"cookie-mismatch"`
    UnsupportedEllipticCurve int `json:"unsupported-elliptic-curve"`
    NoRequiredDigest int `json:"no-required-digest"`
    UnsupportedDigestType int `json:"unsupported-digest-type"`
    BadHandshakeLength int `json:"bad-handshake-length"`
}

func (p *SlbSslErrorStats) GetId() string{
    return "1"
}

func (p *SlbSslErrorStats) getPath() string{
    return "slb/ssl-error/stats"
}

func (p *SlbSslErrorStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslErrorStats,error) {
logger.Println("SlbSslErrorStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslErrorStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
