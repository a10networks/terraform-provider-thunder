---
layout: "thunder"
page_title: "thunder: thunder_web_category_category_list"
sidebar_current: "docs-thunder-resource-web-category-category-list"
description: |-
    Provides details about thunder web category category list resource for A10
---

# thunder\_web\_category\_category\_list

`thunder_web_category_category_list` Provides details about thunder web category category list
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_web_category_category_list" "resourceWebCategoryCategoryListTest" {
	name = "string"
uncategorized = 0
real_estate = 0
computer_and_internet_security = 0
financial_services = 0
business_and_economy = 0
computer_and_internet_info = 0
auctions = 0
shopping = 0
cult_and_occult = 0
travel = 0
drugs = 0
adult_and_pornography = 0
home_and_garden = 0
military = 0
social_network = 0
dead_sites = 0
stock_advice_and_tools = 0
training_and_tools = 0
dating = 0
sex_education = 0
religion = 0
entertainment_and_arts = 0
personal_sites_and_blogs = 0
legal = 0
local_information = 0
streaming_media = 0
job_search = 0
gambling = 0
translation = 0
reference_and_research = 0
shareware_and_freeware = 0
peer_to_peer = 0
marijuana = 0
hacking = 0
games = 0
philosophy_and_politics = 0
weapons = 0
pay_to_surf = 0
hunting_and_fishing = 0
society = 0
educational_institutions = 0
online_greeting_cards = 0
sports = 0
swimsuits_and_intimate_apparel = 0
questionable = 0
kids = 0
hate_and_racism = 0
personal_storage = 0
violence = 0
keyloggers_and_monitoring = 0
search_engines = 0
internet_portals = 0
web_advertisements = 0
cheating = 0
gross = 0
web_based_email = 0
malware_sites = 0
phishing_and_other_fraud = 0
proxy_avoid_and_anonymizers = 0
spyware_and_adware = 0
music = 0
government = 0
nudity = 0
news_and_media = 0
illegal = 0
cdns = 0
internet_communications = 0
bot_nets = 0
abortion = 0
health_and_medicine = 0
confirmed_spam_sources = 0
spam_urls = 0
unconfirmed_spam_sources = 0
open_http_proxies = 0
dynamic_comment = 0
parked_domains = 0
alcohol_and_tobacco = 0
private_ip_addresses = 0
image_and_video_search = 0
fashion_and_beauty = 0
recreation_and_hobbies = 0
motor_vehicles = 0
web_hosting_sites = 0
food_and_dining = 0
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
 
}

```

## Argument Reference

* `category-list` - List of web categories
* `name` - Web Category List name
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
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'uncategorized': uncategorized category; 'real-estate': real estate category; 'computer-and-internet-security': computer and internet security category; 'financial-services': financial services category; 'business-and-economy': business and economy category; 'computer-and-internet-info': computer and internet info category; 'auctions': auctions category; 'shopping': shopping category; 'cult-and-occult': cult and occult category; 'travel': travel category; 'drugs': drugs category; 'adult-and-pornography': adult and pornography category; 'home-and-garden': home and garden category; 'military': military category; 'social-network': social network category; 'dead-sites': dead sites category; 'stock-advice-and-tools': stock advice and tools category; 'training-and-tools': training and tools category; 'dating': dating category; 'sex-education': sex education category; 'religion': religion category; 'entertainment-and-arts': entertainment and arts category; 'personal-sites-and-blogs': personal sites and blogs category; 'legal': legal category; 'local-information': local information category; 'streaming-media': streaming media category; 'job-search': job search category; 'gambling': gambling category; 'translation': translation category; 'reference-and-research': reference and research category; 'shareware-and-freeware': shareware and freeware category; 'peer-to-peer': peer to peer category; 'marijuana': marijuana category; 'hacking': hacking category; 'games': games category; 'philosophy-and-politics': philosophy and politics category; 'weapons': weapons category; 'pay-to-surf': pay to surf category; 'hunting-and-fishing': hunting and fishing category; 'society': society category; 'educational-institutions': educational institutions category; 'online-greeting-cards': online greeting cards category; 'sports': sports category; 'swimsuits-and-intimate-apparel': swimsuits and intimate apparel category; 'questionable': questionable category; 'kids': kids category; 'hate-and-racism': hate and racism category; 'personal-storage': personal storage category; 'violence': violence category; 'keyloggers-and-monitoring': keyloggers and monitoring category; 'search-engines': search engines category; 'internet-portals': internet portals category; 'web-advertisements': web advertisements category; 'cheating': cheating category; 'gross': gross category; 'web-based-email': web based email category; 'malware-sites': malware sites category; 'phishing-and-other-fraud': phishing and other fraud category; 'proxy-avoid-and-anonymizers': proxy avoid and anonymizers category; 'spyware-and-adware': spyware and adware category; 'music': music category; 'government': government category; 'nudity': nudity category; 'news-and-media': news and media category; 'illegal': illegal category; 'CDNs': content delivery networks category; 'internet-communications': internet communications category; 'bot-nets': bot nets category; 'abortion': abortion category; 'health-and-medicine': health and medicine category; 'confirmed-SPAM-sources': confirmed SPAM sources category; 'SPAM-URLs': SPAM URLs category; 'unconfirmed-SPAM-sources': unconfirmed SPAM sources category; 'open-HTTP-proxies': open HTTP proxies category; 'dynamic-comment': dynamic comment category; 'parked-domains': parked domains category; 'alcohol-and-tobacco': alcohol and tobacco category; 'private-IP-addresses': private IP addresses category; 'image-and-video-search': image and video search category; 'fashion-and-beauty': fashion and beauty category; 'recreation-and-hobbies': recreation and hobbies category; 'motor-vehicles': motor vehicles category; 'web-hosting-sites': web hosting sites category; 'food-and-dining': food and dining category;

