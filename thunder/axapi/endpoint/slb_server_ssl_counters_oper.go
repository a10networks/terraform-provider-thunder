

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbServerSslCountersOper struct {
    
    Oper SlbServerSslCountersOperOper `json:"oper"`

}
type DataSlbServerSslCountersOper struct {
    DtSlbServerSslCountersOper SlbServerSslCountersOper `json:"server-ssl-counters"`
}


type SlbServerSslCountersOperOper struct {
    Vserver string `json:"vserver"`
    Port int `json:"port"`
    CumulativeSessions int `json:"cumulative-sessions"`
    Ssl3RsaDes192Cbc3ShaId string `json:"ssl3-rsa-des-192-cbc3-sha-id"`
    Ssl3RsaDes40CbcShaId string `json:"ssl3-rsa-des-40-cbc-sha-id"`
    Ssl3RsaDes64CbcShaId string `json:"ssl3-rsa-des-64-cbc-sha-id"`
    Ssl3RsaRc4128Md5Id string `json:"ssl3-rsa-rc4-128-md5-id"`
    Ssl3RsaRc4128ShaId string `json:"ssl3-rsa-rc4-128-sha-id"`
    Ssl3RsaRc440Md5Id string `json:"ssl3-rsa-rc4-40-md5-id"`
    Tls1DheRsaAes128GcmSha256Id string `json:"tls1-dhe-rsa-aes-128-gcm-sha256-id"`
    Tls1DheRsaAes128ShaId string `json:"tls1-dhe-rsa-aes-128-sha-id"`
    Tls1DheRsaAes128Sha256Id string `json:"tls1-dhe-rsa-aes-128-sha256-id"`
    Tls1DheRsaAes256GcmSha384Id string `json:"tls1-dhe-rsa-aes-256-gcm-sha384-id"`
    Tls1DheRsaAes256ShaId string `json:"tls1-dhe-rsa-aes-256-sha-id"`
    Tls1DheRsaAes256Sha256Id string `json:"tls1-dhe-rsa-aes-256-sha256-id"`
    Tls1EcdheEcdsaAes128GcmSha256Id string `json:"tls1-ecdhe-ecdsa-aes-128-gcm-sha256-id"`
    Tls1EcdheEcdsaAes128ShaId string `json:"tls1-ecdhe-ecdsa-aes-128-sha-id"`
    Tls1EcdheEcdsaAes128Sha256Id string `json:"tls1-ecdhe-ecdsa-aes-128-sha256-id"`
    Tls1EcdheEcdsaAes256Sha384Id string `json:"tls1-ecdhe-ecdsa-aes-256-sha384-id"`
    Tls1EcdheEcdsaAes256GcmSha384Id string `json:"tls1-ecdhe-ecdsa-aes-256-gcm-sha384-id"`
    Tls1EcdheEcdsaAes256ShaId string `json:"tls1-ecdhe-ecdsa-aes-256-sha-id"`
    Tls1EcdheRsaAes128GcmSha256Id string `json:"tls1-ecdhe-rsa-aes-128-gcm-sha256-id"`
    Tls1EcdheRsaAes128ShaId string `json:"tls1-ecdhe-rsa-aes-128-sha-id"`
    Tls1EcdheRsaAes128Sha256Id string `json:"tls1-ecdhe-rsa-aes-128-sha256-id"`
    Tls1EcdheRsaAes256Sha384Id string `json:"tls1-ecdhe-rsa-aes-256-sha384-id"`
    Tls1EcdheRsaAes256GcmSha384Id string `json:"tls1-ecdhe-rsa-aes-256-gcm-sha384-id"`
    Tls1EcdheRsaAes256ShaId string `json:"tls1-ecdhe-rsa-aes-256-sha-id"`
    Tls1RsaAes128GcmSha256Id string `json:"tls1-rsa-aes-128-gcm-sha256-id"`
    Tls1RsaAes128ShaId string `json:"tls1-rsa-aes-128-sha-id"`
    Tls1RsaAes128Sha256Id string `json:"tls1-rsa-aes-128-sha256-id"`
    Tls1RsaAes256GcmSha384Id string `json:"tls1-rsa-aes-256-gcm-sha384-id"`
    Tls1RsaAes256ShaId string `json:"tls1-rsa-aes-256-sha-id"`
    Tls1RsaAes256Sha256Id string `json:"tls1-rsa-aes-256-sha256-id"`
    Tls1RsaExport1024Rc456Md5Id string `json:"tls1-rsa-export1024-rc4-56-md5-id"`
    Tls1RsaExport1024Rc456ShaId string `json:"tls1-rsa-export1024-rc4-56-sha-id"`
    Tls1EcdheRsaChacha20Poly1305Sha256Id string `json:"tls1-ecdhe-rsa-chacha20-poly1305-sha256-id"`
    Tls1EcdheEcdsaChacha20Poly1305Sha256Id string `json:"tls1-ecdhe-ecdsa-chacha20-poly1305-sha256-id"`
    Tls1DheRsaChacha20Poly1305Sha256Id string `json:"tls1-dhe-rsa-chacha20-poly1305-sha256-id"`
    Tls13Aes128GcmSha256Id string `json:"tls13-aes-128-gcm-sha256-id"`
    Tls13Aes256GcmSha384Id string `json:"tls13-aes-256-gcm-sha384-id"`
    Tls13Chacha20Poly1305Sha256Id string `json:"tls13-chacha20-poly1305-sha256-id"`
    Tls1EcdheSm2Sms4Sm3Id string `json:"tls1-ecdhe-sm2-sms4-sm3-id"`
    Tls1EcdheSm2Sms4GcmSm3Id string `json:"tls1-ecdhe-sm2-sms4-gcm-sm3-id"`
    Tls1EcdheSm2Sms4Sha256Id string `json:"tls1-ecdhe-sm2-sms4-sha256-id"`
    Ssl3RsaDes192Cbc3ShaSuccesses int `json:"ssl3-rsa-des-192-cbc3-sha-successes"`
    Ssl3RsaDes40CbcShaSuccesses int `json:"ssl3-rsa-des-40-cbc-sha-successes"`
    Ssl3RsaDes64CbcShaSuccesses int `json:"ssl3-rsa-des-64-cbc-sha-successes"`
    Ssl3RsaRc4128Md5Successes int `json:"ssl3-rsa-rc4-128-md5-successes"`
    Ssl3RsaRc4128ShaSuccesses int `json:"ssl3-rsa-rc4-128-sha-successes"`
    Ssl3RsaRc440Md5Successes int `json:"ssl3-rsa-rc4-40-md5-successes"`
    Tls1DheRsaAes128GcmSha256Successes int `json:"tls1-dhe-rsa-aes-128-gcm-sha256-successes"`
    Tls1DheRsaAes128ShaSuccesses int `json:"tls1-dhe-rsa-aes-128-sha-successes"`
    Tls1DheRsaAes128Sha256Successes int `json:"tls1-dhe-rsa-aes-128-sha256-successes"`
    Tls1DheRsaAes256GcmSha384Successes int `json:"tls1-dhe-rsa-aes-256-gcm-sha384-successes"`
    Tls1DheRsaAes256ShaSuccesses int `json:"tls1-dhe-rsa-aes-256-sha-successes"`
    Tls1DheRsaAes256Sha256Successes int `json:"tls1-dhe-rsa-aes-256-sha256-successes"`
    Tls1EcdheEcdsaAes128GcmSha256Successes int `json:"tls1-ecdhe-ecdsa-aes-128-gcm-sha256-successes"`
    Tls1EcdheEcdsaAes128ShaSuccesses int `json:"tls1-ecdhe-ecdsa-aes-128-sha-successes"`
    Tls1EcdheEcdsaAes128Sha256Successes int `json:"tls1-ecdhe-ecdsa-aes-128-sha256-successes"`
    Tls1EcdheEcdsaAes256Sha384Successes int `json:"tls1-ecdhe-ecdsa-aes-256-sha384-successes"`
    Tls1EcdheEcdsaAes256GcmSha384Successes int `json:"tls1-ecdhe-ecdsa-aes-256-gcm-sha384-successes"`
    Tls1EcdheEcdsaAes256ShaSuccesses int `json:"tls1-ecdhe-ecdsa-aes-256-sha-successes"`
    Tls1EcdheRsaAes128GcmSha256Successes int `json:"tls1-ecdhe-rsa-aes-128-gcm-sha256-successes"`
    Tls1EcdheRsaAes128ShaSuccesses int `json:"tls1-ecdhe-rsa-aes-128-sha-successes"`
    Tls1EcdheRsaAes128Sha256Successes int `json:"tls1-ecdhe-rsa-aes-128-sha256-successes"`
    Tls1EcdheRsaAes256Sha384Successes int `json:"tls1-ecdhe-rsa-aes-256-sha384-successes"`
    Tls1EcdheRsaAes256GcmSha384Successes int `json:"tls1-ecdhe-rsa-aes-256-gcm-sha384-successes"`
    Tls1EcdheRsaAes256ShaSuccesses int `json:"tls1-ecdhe-rsa-aes-256-sha-successes"`
    Tls1RsaAes128GcmSha256Successes int `json:"tls1-rsa-aes-128-gcm-sha256-successes"`
    Tls1RsaAes128ShaSuccesses int `json:"tls1-rsa-aes-128-sha-successes"`
    Tls1RsaAes128Sha256Successes int `json:"tls1-rsa-aes-128-sha256-successes"`
    Tls1RsaAes256GcmSha384Successes int `json:"tls1-rsa-aes-256-gcm-sha384-successes"`
    Tls1RsaAes256ShaSuccesses int `json:"tls1-rsa-aes-256-sha-successes"`
    Tls1RsaAes256Sha256Successes int `json:"tls1-rsa-aes-256-sha256-successes"`
    Tls1RsaExport1024Rc456Md5Successes int `json:"tls1-rsa-export1024-rc4-56-md5-successes"`
    Tls1RsaExport1024Rc456ShaSuccesses int `json:"tls1-rsa-export1024-rc4-56-sha-successes"`
    Tls1EcdheRsaChacha20Poly1305Sha256Successes int `json:"tls1-ecdhe-rsa-chacha20-poly1305-sha256-successes"`
    Tls1EcdheEcdsaChacha20Poly1305Sha256Successes int `json:"tls1-ecdhe-ecdsa-chacha20-poly1305-sha256-successes"`
    Tls1DheRsaChacha20Poly1305Sha256Successes int `json:"tls1-dhe-rsa-chacha20-poly1305-sha256-successes"`
    Tls13Aes128GcmSha256Successes int `json:"tls13-aes-128-gcm-sha256-successes"`
    Tls13Aes256GcmSha384Successes int `json:"tls13-aes-256-gcm-sha384-successes"`
    Tls13Chacha20Poly1305Sha256Successes int `json:"tls13-chacha20-poly1305-sha256-successes"`
    Tls1EcdheSm2Sms4Sm3Successes int `json:"tls1-ecdhe-sm2-sms4-sm3-successes"`
    Tls1EcdheSm2Sms4GcmSm3Successes int `json:"tls1-ecdhe-sm2-sms4-gcm-sm3-successes"`
    Tls1EcdheSm2Sms4Sha256Successes int `json:"tls1-ecdhe-sm2-sms4-sha256-successes"`
    Ssl3RsaDes192Cbc3ShaFailures int `json:"ssl3-rsa-des-192-cbc3-sha-failures"`
    Ssl3RsaDes40CbcShaFailures int `json:"ssl3-rsa-des-40-cbc-sha-failures"`
    Ssl3RsaDes64CbcShaFailures int `json:"ssl3-rsa-des-64-cbc-sha-failures"`
    Ssl3RsaRc4128Md5Failures int `json:"ssl3-rsa-rc4-128-md5-failures"`
    Ssl3RsaRc4128ShaFailures int `json:"ssl3-rsa-rc4-128-sha-failures"`
    Ssl3RsaRc440Md5Failures int `json:"ssl3-rsa-rc4-40-md5-failures"`
    Tls1DheRsaAes128GcmSha256Failures int `json:"tls1-dhe-rsa-aes-128-gcm-sha256-failures"`
    Tls1DheRsaAes128ShaFailures int `json:"tls1-dhe-rsa-aes-128-sha-failures"`
    Tls1DheRsaAes128Sha256Failures int `json:"tls1-dhe-rsa-aes-128-sha256-failures"`
    Tls1DheRsaAes256GcmSha384Failures int `json:"tls1-dhe-rsa-aes-256-gcm-sha384-failures"`
    Tls1DheRsaAes256ShaFailures int `json:"tls1-dhe-rsa-aes-256-sha-failures"`
    Tls1DheRsaAes256Sha256Failures int `json:"tls1-dhe-rsa-aes-256-sha256-failures"`
    Tls1EcdheEcdsaAes128GcmSha256Failures int `json:"tls1-ecdhe-ecdsa-aes-128-gcm-sha256-failures"`
    Tls1EcdheEcdsaAes128ShaFailures int `json:"tls1-ecdhe-ecdsa-aes-128-sha-failures"`
    Tls1EcdheEcdsaAes128Sha256Failures int `json:"tls1-ecdhe-ecdsa-aes-128-sha256-failures"`
    Tls1EcdheEcdsaAes256Sha384Failures int `json:"tls1-ecdhe-ecdsa-aes-256-sha384-failures"`
    Tls1EcdheEcdsaAes256GcmSha384Failures int `json:"tls1-ecdhe-ecdsa-aes-256-gcm-sha384-failures"`
    Tls1EcdheEcdsaAes256ShaFailures int `json:"tls1-ecdhe-ecdsa-aes-256-sha-failures"`
    Tls1EcdheRsaAes128GcmSha256Failures int `json:"tls1-ecdhe-rsa-aes-128-gcm-sha256-failures"`
    Tls1EcdheRsaAes128ShaFailures int `json:"tls1-ecdhe-rsa-aes-128-sha-failures"`
    Tls1EcdheRsaAes128Sha256Failures int `json:"tls1-ecdhe-rsa-aes-128-sha256-failures"`
    Tls1EcdheRsaAes256Sha384Failures int `json:"tls1-ecdhe-rsa-aes-256-sha384-failures"`
    Tls1EcdheRsaAes256GcmSha384Failures int `json:"tls1-ecdhe-rsa-aes-256-gcm-sha384-failures"`
    Tls1EcdheRsaAes256ShaFailures int `json:"tls1-ecdhe-rsa-aes-256-sha-failures"`
    Tls1RsaAes128GcmSha256Failures int `json:"tls1-rsa-aes-128-gcm-sha256-failures"`
    Tls1RsaAes128ShaFailures int `json:"tls1-rsa-aes-128-sha-failures"`
    Tls1RsaAes128Sha256Failures int `json:"tls1-rsa-aes-128-sha256-failures"`
    Tls1RsaAes256GcmSha384Failures int `json:"tls1-rsa-aes-256-gcm-sha384-failures"`
    Tls1RsaAes256ShaFailures int `json:"tls1-rsa-aes-256-sha-failures"`
    Tls1RsaAes256Sha256Failures int `json:"tls1-rsa-aes-256-sha256-failures"`
    Tls1RsaExport1024Rc456Md5Failures int `json:"tls1-rsa-export1024-rc4-56-md5-failures"`
    Tls1RsaExport1024Rc456ShaFailures int `json:"tls1-rsa-export1024-rc4-56-sha-failures"`
    Tls1EcdheRsaChacha20Poly1305Sha256Failures int `json:"tls1-ecdhe-rsa-chacha20-poly1305-sha256-failures"`
    Tls1EcdheEcdsaChacha20Poly1305Sha256Failures int `json:"tls1-ecdhe-ecdsa-chacha20-poly1305-sha256-failures"`
    Tls1DheRsaChacha20Poly1305Sha256Failures int `json:"tls1-dhe-rsa-chacha20-poly1305-sha256-failures"`
    Tls13Aes128GcmSha256Failures int `json:"tls13-aes-128-gcm-sha256-failures"`
    Tls13Aes256GcmSha384Failures int `json:"tls13-aes-256-gcm-sha384-failures"`
    Tls13Chacha20Poly1305Sha256Failures int `json:"tls13-chacha20-poly1305-sha256-failures"`
    Tls1EcdheSm2Sms4Sm3Failures int `json:"tls1-ecdhe-sm2-sms4-sm3-failures"`
    Tls1EcdheSm2Sms4GcmSm3Failures int `json:"tls1-ecdhe-sm2-sms4-gcm-sm3-failures"`
    Tls1EcdheSm2Sms4Sha256Failures int `json:"tls1-ecdhe-sm2-sms4-sha256-failures"`
    KexRsa512Successes int `json:"kex-rsa-512-successes"`
    KexRsa1024Successes int `json:"kex-rsa-1024-successes"`
    KexRsa2048Successes int `json:"kex-rsa-2048-successes"`
    KexRsa4096Successes int `json:"kex-rsa-4096-successes"`
    KexRsa512Failures int `json:"kex-rsa-512-failures"`
    KexRsa1024Failures int `json:"kex-rsa-1024-failures"`
    KexRsa2048Failures int `json:"kex-rsa-2048-failures"`
    KexRsa4096Failures int `json:"kex-rsa-4096-failures"`
    KexEcdheSecp256r1Successes int `json:"kex-ecdhe-secp256r1-successes"`
    KexEcdheSecp384r1Successes int `json:"kex-ecdhe-secp384r1-successes"`
    KexEcdheSecp521r1Successes int `json:"kex-ecdhe-secp521r1-successes"`
    KexEcdheX25519Successes int `json:"kex-ecdhe-x25519-successes"`
    KexEcdheX448Successes int `json:"kex-ecdhe-x448-successes"`
    KexEcdheSm2Successes int `json:"kex-ecdhe-sm2-successes"`
    KexEcdheSecp256r1Failures int `json:"kex-ecdhe-secp256r1-failures"`
    KexEcdheSecp384r1Failures int `json:"kex-ecdhe-secp384r1-failures"`
    KexEcdheSecp521r1Failures int `json:"kex-ecdhe-secp521r1-failures"`
    KexEcdheX25519Failures int `json:"kex-ecdhe-x25519-failures"`
    KexEcdheX448Failures int `json:"kex-ecdhe-x448-failures"`
    KexEcdheSm2Failures int `json:"kex-ecdhe-sm2-failures"`
    KexDhe512Successes int `json:"kex-dhe-512-successes"`
    KexDhe1024Successes int `json:"kex-dhe-1024-successes"`
    KexDhe2048Successes int `json:"kex-dhe-2048-successes"`
    KexDhe512Failures int `json:"kex-dhe-512-failures"`
    KexDhe1024Failures int `json:"kex-dhe-1024-failures"`
    KexDhe2048Failures int `json:"kex-dhe-2048-failures"`
    Ssl2Successes int `json:"ssl2-successes"`
    Ssl3Successes int `json:"ssl3-successes"`
    Tls10Successes int `json:"tls10-successes"`
    Tls11Successes int `json:"tls11-successes"`
    Tls12Successes int `json:"tls12-successes"`
    Tls13Successes int `json:"tls13-successes"`
    Ssl2Failures int `json:"ssl2-failures"`
    Ssl3Failures int `json:"ssl3-failures"`
    Tls10Failures int `json:"tls10-failures"`
    Tls11Failures int `json:"tls11-failures"`
    Tls12Failures int `json:"tls12-failures"`
    Tls13Failures int `json:"tls13-failures"`
    SessCacheNew int `json:"sess-cache-new"`
    SessCacheHit int `json:"sess-cache-hit"`
    SessCacheMiss int `json:"sess-cache-miss"`
    SessCacheTimeout int `json:"sess-cache-timeout"`
    SessCacheCurrConn int `json:"sess-cache-curr-conn"`
    HsFailures int `json:"hs-failures"`
    CertVfy int `json:"cert-vfy"`
    HsAvgTime int `json:"hs-avg-time"`
    RenegotiationTotal int `json:"renegotiation-total"`
    RenegoSsl2Successes int `json:"renego-ssl2-successes"`
    RenegoSsl3Successes int `json:"renego-ssl3-successes"`
    RenegoTls10Successes int `json:"renego-tls10-successes"`
    RenegoTls11Successes int `json:"renego-tls11-successes"`
    RenegoTls12Successes int `json:"renego-tls12-successes"`
    RenegoTls13Successes int `json:"renego-tls13-successes"`
    RenegoSsl2Failures int `json:"renego-ssl2-failures"`
    RenegoSsl3Failures int `json:"renego-ssl3-failures"`
    RenegoTls10Failures int `json:"renego-tls10-failures"`
    RenegoTls11Failures int `json:"renego-tls11-failures"`
    RenegoTls12Failures int `json:"renego-tls12-failures"`
    RenegoTls13Failures int `json:"renego-tls13-failures"`
    Downgraded int `json:"downgraded"`
}

func (p *SlbServerSslCountersOper) GetId() string{
    return "1"
}

func (p *SlbServerSslCountersOper) getPath() string{
    return "slb/server-ssl-counters/oper"
}

func (p *SlbServerSslCountersOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbServerSslCountersOper,error) {
logger.Println("SlbServerSslCountersOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbServerSslCountersOper
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
