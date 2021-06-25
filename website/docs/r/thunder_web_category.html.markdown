---
layout: "thunder"
page_title: "thunder: thunder_web_category"
sidebar_current: "docs-thunder-resource-web-category"
description: |-
    Provides details about thunder web category resource for A10
---

# thunder\_web\_category

`thunder_web_category` Provides details about thunder web category
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_web_category" "resourceWebCategoryTest" {
	server = "string"
database_server = "string"
port = 0
ssl_port = 0
server_timeout = 0
cloud_query_disable = 0
cloud_query_cache_size = 0
db_update_time = "string"
rtu_update_disable = 0
rtu_update_interval = 0
rtu_cache_size = 0
use_mgmt_port = 0
remote_syslog_enable = 0
enable = 0
uuid = "string"
proxy_server {  
 	proxy_host =  "string" 
	http_port =  0 
	https_port =  0 
	auth_type =  "string" 
	domain =  "string" 
	username =  "string" 
	password =  0 
	secret_string =  "string" 
	uuid =  "string" 
	}
intercepted_urls {  
 	uuid =  "string" 
	}
bypassed_urls {  
 	uuid =  "string" 
	}
url {  
 	uuid =  "string" 
	}
license {  
 	uuid =  "string" 
	}
category-list-list {   
	name =  "string" 
	uncategorized =  0 
	real_estate =  0 
	computer_and_internet_security =  0 
	financial_services =  0 
	business_and_economy =  0 
	computer_and_internet_info =  0 
	auctions =  0 
	shopping =  0 
	cult_and_occult =  0 
	travel =  0 
	drugs =  0 
	adult_and_pornography =  0 
	home_and_garden =  0 
	military =  0 
	social_network =  0 
	dead_sites =  0 
	stock_advice_and_tools =  0 
	training_and_tools =  0 
	dating =  0 
	sex_education =  0 
	religion =  0 
	entertainment_and_arts =  0 
	personal_sites_and_blogs =  0 
	legal =  0 
	local_information =  0 
	streaming_media =  0 
	job_search =  0 
	gambling =  0 
	translation =  0 
	reference_and_research =  0 
	shareware_and_freeware =  0 
	peer_to_peer =  0 
	marijuana =  0 
	hacking =  0 
	games =  0 
	philosophy_and_politics =  0 
	weapons =  0 
	pay_to_surf =  0 
	hunting_and_fishing =  0 
	society =  0 
	educational_institutions =  0 
	online_greeting_cards =  0 
	sports =  0 
	swimsuits_and_intimate_apparel =  0 
	questionable =  0 
	kids =  0 
	hate_and_racism =  0 
	personal_storage =  0 
	violence =  0 
	keyloggers_and_monitoring =  0 
	search_engines =  0 
	internet_portals =  0 
	web_advertisements =  0 
	cheating =  0 
	gross =  0 
	web_based_email =  0 
	malware_sites =  0 
	phishing_and_other_fraud =  0 
	proxy_avoid_and_anonymizers =  0 
	spyware_and_adware =  0 
	music =  0 
	government =  0 
	nudity =  0 
	news_and_media =  0 
	illegal =  0 
	cdns =  0 
	internet_communications =  0 
	bot_nets =  0 
	abortion =  0 
	health_and_medicine =  0 
	confirmed_spam_sources =  0 
	spam_urls =  0 
	unconfirmed_spam_sources =  0 
	open_http_proxies =  0 
	dynamic_comment =  0 
	parked_domains =  0 
	alcohol_and_tobacco =  0 
	private_ip_addresses =  0 
	image_and_video_search =  0 
	fashion_and_beauty =  0 
	recreation_and_hobbies =  0 
	motor_vehicles =  0 
	web_hosting_sites =  0 
	food_and_dining =  0 
	uuid =  "string" 
	user_tag =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
statistics {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
reputation-scope-list {   
	name =  "string" 
greater_than {  
 	greater_trustworthy =  0 
	greater_low_risk =  0 
	greater_moderate_risk =  0 
	greater_suspicious =  0 
	greater_malicious =  0 
	greater_threshold =  0 
	}
less_than {  
 	less_trustworthy =  0 
	less_low_risk =  0 
	less_moderate_risk =  0 
	less_suspicious =  0 
	less_malicious =  0 
	less_threshold =  0 
	}
	uuid =  "string" 
	user_tag =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
web_reputation {  
 	uuid =  "string" 
intercepted_urls {  
 	uuid =  "string" 
	}
bypassed_urls {  
 	uuid =  "string" 
	}
url {  
 	uuid =  "string" 
	}
	}
 
}

```

## Argument Reference

* `web-category` - Web-Category Commands
* `server` - BrightCloud Query Server
* `database-server` - BrightCloud Database Server
* `port` - BrightCloud Query Server Listening Port(default 80)
* `ssl-port` - BrightCloud Servers SSL Port(default 443)
* `server-timeout` - BrightCloud Servers Timeout in seconds (default: 15s)
* `cloud-query-disable` - Disables cloud queries for URL's not present in local database(default enable)
* `cloud-query-cache-size` - Maximum cache size for storing cloud query results
* `db-update-time` - Time of day to update database (default: 00:00)
* `rtu-update-disable` - Disables real time updates(default enable)
* `rtu-update-interval` - Interval to check for real time updates if enabled in mins(default 60)
* `rtu-cache-size` - Maximum cache size for storing RTU updates
* `use-mgmt-port` - Use management interface for all communication with BrightCloud
* `remote-syslog-enable` - Enable data plane logging to a remote syslog server
* `enable` - Enable BrightCloud SDK
* `uuid` - uuid of the object
* `proxy-host` - Proxy server hostname or IP address
* `http-port` - Proxy server HTTP port
* `https-port` - Proxy server HTTPS port(HTTP port will be used if not configured)
* `auth-type` - 'ntlm': NTLM authentication(default); 'basic': Basic authentication;
* `domain` - Realm for NTLM authentication
* `username` - Username for proxy authentication
* `password` - Password for proxy authentication
* `secret-string` - password value
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)
* `name` - Reputation Scope name
* `uncategorized` - Uncategorized URLs
* `real-estate` - Category Real Estate
* `computer-and-internet-security` - Category Computer and Internet Security
* `financial-services` - Category Financial Services
* `business-and-economy` - Category Business and Economy
* `computer-and-internet-info` - Category Computer and Internet Info
* `auctions` - Category Auctions
* `shopping` - Category Shopping
* `cult-and-occult` - Category Cult and Occult
* `travel` - Category Travel
* `drugs` - Category Abused Drugs
* `adult-and-pornography` - Category Adult and Pornography
* `home-and-garden` - Category Home and Garden
* `military` - Category Military
* `social-network` - Category Social Network
* `dead-sites` - Category Dead Sites (db Ops only)
* `stock-advice-and-tools` - Category Stock Advice and Tools
* `training-and-tools` - Category Training and Tools
* `dating` - Category Dating
* `sex-education` - Category Sex Education
* `religion` - Category Religion
* `entertainment-and-arts` - Category Entertainment and Arts
* `personal-sites-and-blogs` - Category Personal sites and Blogs
* `legal` - Category Legal
* `local-information` - Category Local Information
* `streaming-media` - Category Streaming Media
* `job-search` - Category Job Search
* `gambling` - Category Gambling
* `translation` - Category Translation
* `reference-and-research` - Category Reference and Research
* `shareware-and-freeware` - Category Shareware and Freeware
* `peer-to-peer` - Category Peer to Peer
* `marijuana` - Category Marijuana
* `hacking` - Category Hacking
* `games` - Category Games
* `philosophy-and-politics` - Category Philosophy and Political Advocacy
* `weapons` - Category Weapons
* `pay-to-surf` - Category Pay to Surf
* `hunting-and-fishing` - Category Hunting and Fishing
* `society` - Category Society
* `educational-institutions` - Category Educational Institutions
* `online-greeting-cards` - Category Online Greeting cards
* `sports` - Category Sports
* `swimsuits-and-intimate-apparel` - Category Swimsuits and Intimate Apparel
* `questionable` - Category Questionable
* `kids` - Category Kids
* `hate-and-racism` - Category Hate and Racism
* `personal-storage` - Category Personal Storage
* `violence` - Category Violence
* `keyloggers-and-monitoring` - Category Keyloggers and Monitoring
* `search-engines` - Category Search Engines
* `internet-portals` - Category Internet Portals
* `web-advertisements` - Category Web Advertisements
* `cheating` - Category Cheating
* `gross` - Category Gross
* `web-based-email` - Category Web based email
* `malware-sites` - Category Malware Sites
* `phishing-and-other-fraud` - Category Phishing and Other Frauds
* `proxy-avoid-and-anonymizers` - Category Proxy Avoid and Anonymizers
* `spyware-and-adware` - Category Spyware and Adware
* `music` - Category Music
* `government` - Category Government
* `nudity` - Category Nudity
* `news-and-media` - Category News and Media
* `illegal` - Category Illegal
* `cdns` - Category CDNs
* `internet-communications` - Category Internet Communications
* `bot-nets` - Category Bot Nets
* `abortion` - Category Abortion
* `health-and-medicine` - Category Health and Medicine
* `confirmed-spam-sources` - Category Confirmed SPAM Sources
* `spam-urls` - Category SPAM URLs
* `unconfirmed-spam-sources` - Category Unconfirmed SPAM Sources
* `open-http-proxies` - Category Open HTTP Proxies
* `dynamic-comment` - Category Dynamic Comment
* `parked-domains` - Category Parked Domains
* `alcohol-and-tobacco` - Category Alcohol and Tobacco
* `private-ip-addresses` - Category Private IP Addresses
* `image-and-video-search` - Category Image and Video Search
* `fashion-and-beauty` - Category Fashion and Beauty
* `recreation-and-hobbies` - Category Recreation and Hobbies
* `motor-vehicles` - Category Motor Vehicles
* `web-hosting-sites` - Category Web Hosting Sites
* `food-and-dining` - Category Food and Dining
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'trustworthy': Trustworthy level(81-100); 'low-risk': Low-risk level(61-80); 'moderate-risk': Moderate-risk level(41-60); 'suspicious': Suspicious level(21-40); 'malicious': Malicious level(1-20);
* `greater-trustworthy` - Reputation score is greater than or equal to 81
* `greater-low-risk` - Reputation score is greater than or equal to 61
* `greater-moderate-risk` - Reputation score is greater than or equal to 41
* `greater-suspicious` - Reputation score is greater than or equal to 21
* `greater-malicious` - Reputation score is greater than or equal to 1
* `greater-threshold` - Reputation score is greater than or equal to the customized score (1-100)
* `less-trustworthy` - Reputation score is less than or equal to 100
* `less-low-risk` - Reputation score is less than or equal to 80
* `less-moderate-risk` - Reputation score is less than or equal to 60
* `less-suspicious` - Reputation score is less than or equal to 40
* `less-malicious` - Reputation score is less than or equal to 20
* `less-threshold` - Reputation score is less than or equal to a customized value (1-100)

