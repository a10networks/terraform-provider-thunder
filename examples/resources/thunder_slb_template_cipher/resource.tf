provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_slb_template_cipher" "item1" {
    name = "ITEM_1"
    cipher_cfg {
        cipher_suite = "SSL3_RSA_DES_192_CBC3_SHA"
        priority = 15
    }
    cipher_cfg {
        cipher_suite = "TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256"
    }
    cipher13_cfg {
        cipher13_suite = "TLS_CHACHA20_POLY1305_SHA256"
        priority = 90
    }
}

