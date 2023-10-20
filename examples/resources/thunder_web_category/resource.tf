provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_web_category" "resourceWebCategoryTest" {
  server                 = "test_server"
  database_server        = "test_db_server"
  port                   = 80
  ssl_port               = 443
  server_timeout         = 15
  cloud_query_disable    = 0
  cloud_query_cache_size = 500
  rtu_update_disable     = 0
  rtu_update_interval    = 60
  rtu_cache_size         = 500
  use_mgmt_port          = 0
  remote_syslog_enable   = 0
  enable                 = 0
  proxy_server {
    proxy_host = "192.168.50.10"
    http_port  = 80
    https_port = 443
    auth_type  = "basic"
    username   = "testing"
    password   = 0
  }
  category_list_list {
    name                           = "testingUser"
    user_tag                       = "abc"
    uncategorized                  = 1
    real_estate                    = 0
    computer_and_internet_security = 0
    financial_services             = 0
    business_and_economy           = 1
    computer_and_internet_info     = 0
    auctions                       = 1
    shopping                       = 1
    cult_and_occult                = 1
    travel                         = 1
    drugs                          = 0
    adult_and_pornography          = 1
    home_and_garden                = 0
    military                       = 1
    social_network                 = 0
    dead_sites                     = 0
    stock_advice_and_tools         = 0
    training_and_tools             = 0
    sex_education                  = 0
    dating                         = 0
    religion                       = 0
    entertainment_and_arts         = 0
    personal_sites_and_blogs       = 1
    legal                          = 0
    local_information              = 0
    streaming_media                = 0
    job_search                     = 0
    gambling                       = 0
    translation                    = 0
    reference_and_research         = 0
    shareware_and_freeware         = 1
    peer_to_peer                   = 0
    marijuana                      = 0
    hacking                        = 0
    games                          = 0
    philosophy_and_politics        = 1
    weapons                        = 0
    pay_to_surf                    = 0
    hunting_and_fishing            = 1
    society                        = 0
    educational_institutions       = 0
    online_greeting_cards          = 0
    sports                         = 0
    swimsuits_and_intimate_apparel = 1
    questionable                   = 0
    kids                           = 0
    hate_and_racism                = 0
    personal_storage               = 1
    violence                       = 0
    keyloggers_and_monitoring      = 0
    search_engines                 = 0
    internet_portals               = 0
    web_advertisements             = 1
    cheating                       = 1
    gross                          = 1
    web_based_email                = 1
    malware_sites                  = 1
    phishing_and_other_fraud       = 0
    proxy_avoid_and_anonymizers    = 0
    spyware_and_adware             = 0
    music                          = 1
    government                     = 1
    nudity                         = 1
    news_and_media                 = 1
    illegal                        = 1
    cdns                           = 1
    internet_communications        = 1
    bot_nets                       = 1
    abortion                       = 1
    health_and_medicine            = 1
    confirmed_spam_sources         = 1
    spam_urls                      = 1
    unconfirmed_spam_sources       = 1
    open_http_proxies              = 1
    dynamic_comment                = 1
    parked_domains                 = 1
    alcohol_and_tobacco            = 1
    private_ip_addresses           = 1
    image_and_video_search         = 1
    fashion_and_beauty             = 1
    recreation_and_hobbies         = 1
    motor_vehicles                 = 1
    web_hosting_sites              = 1
    food_and_dining                = 1
  }
  statistics {
    sampling_enable {
      counters1 = "all"
    }
  }
  reputation_scope_list {
    name = "testingUser"
    greater_than {
      greater_trustworthy   = 1
      greater_low_risk      = 0
      greater_moderate_risk = 0
      greater_suspicious    = 0
      greater_malicious     = 0
      greater_threshold     = 0
    }
    less_than {
      less_trustworthy   = 0
      less_low_risk      = 0
      less_moderate_risk = 0
      less_suspicious    = 0
      less_malicious     = 0
      less_threshold     = 0
    }
    user_tag = "abc"
  }
}
