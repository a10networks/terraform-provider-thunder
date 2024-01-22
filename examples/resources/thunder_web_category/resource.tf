provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_web_category" "thunder_web_category" {

  category_list_list {
    name                           = "test"
    uncategorized                  = 1
    real_estate                    = 1
    computer_and_internet_security = 1
    financial_services             = 1
    business_and_economy           = 1
    computer_and_internet_info     = 1
    auctions                       = 1
    shopping                       = 1
    cult_and_occult                = 1
    travel                         = 1
    drugs                          = 1
    adult_and_pornography          = 1
    home_and_garden                = 1
    military                       = 1
    social_network                 = 1
    dead_sites                     = 1
    stock_advice_and_tools         = 1
    training_and_tools             = 1
    dating                         = 1
    sex_education                  = 1
    religion                       = 1
    entertainment_and_arts         = 1
    personal_sites_and_blogs       = 1
    legal                          = 1
    local_information              = 1
    streaming_media                = 1
    job_search                     = 1
    gambling                       = 1
    translation                    = 1
    reference_and_research         = 1
    shareware_and_freeware         = 1
    peer_to_peer                   = 1
    marijuana                      = 1
    hacking                        = 1
    games                          = 1
    philosophy_and_politics        = 1
    weapons                        = 1
    pay_to_surf                    = 1
    hunting_and_fishing            = 1
    society                        = 1
    educational_institutions       = 1
    online_greeting_cards          = 1
    sports                         = 1
    swimsuits_and_intimate_apparel = 1
    questionable                   = 1
    kids                           = 1
    hate_and_racism                = 1
    personal_storage               = 1
    violence                       = 1
    keyloggers_and_monitoring      = 1
    search_engines                 = 1
    internet_portals               = 1
    web_advertisements             = 1
    cheating                       = 1
    gross                          = 1
    web_based_email                = 1
    malware_sites                  = 1
    phishing_and_other_fraud       = 1
    proxy_avoid_and_anonymizers    = 1
    spyware_and_adware             = 1
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
    spam_urls                      = 1
    dynamically_generated_content  = 1
    parked_domains                 = 1
    alcohol_and_tobacco            = 1
    image_and_video_search         = 1
    fashion_and_beauty             = 1
    recreation_and_hobbies         = 1
    motor_vehicles                 = 1
    web_hosting_sites              = 1
    nudity_artistic                = 1
    illegal_pornography            = 1
    user_tag                       = "97"
    sampling_enable {
      counters1 = "all"
    }
  }
  cloud_query_cache_size = 222
  cloud_query_disable    = 1
  database_server        = "133"
  online_check_disable   = 1
  port                   = 80
  proxy_server {
    proxy_host    = "16"
    http_port     = 37683
    https_port    = 46027
    auth_type     = "ntlm"
    username      = "99"
    password      = 1
    secret_string = "47"
  }
  remote_syslog_enable = 1
  reputation_scope_list {
    name = "test"
    greater_than {
      greater_trustworthy = 1
    }
    less_than {
      less_trustworthy = 1
    }
    user_tag = "9"
    sampling_enable {
      counters1 = "all"
    }
  }
  rtu_cache_size      = 222
  rtu_update_disable  = 1
  rtu_update_interval = 60
  server_timeout      = 15
  ssl_port            = 443
  statistics {
    sampling_enable {
      counters1 = "all"
    }
  }
}