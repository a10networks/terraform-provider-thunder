package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategory() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_category`: Web-Category Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceWebCategoryCreate,
		UpdateContext: resourceWebCategoryUpdate,
		ReadContext:   resourceWebCategoryRead,
		DeleteContext: resourceWebCategoryDelete,

		Schema: map[string]*schema.Schema{
			"bypassed_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"category_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Web Category List name",
						},
						"uncategorized": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Uncategorized URLs",
						},
						"real_estate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Real Estate",
						},
						"computer_and_internet_security": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Security",
						},
						"financial_services": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Financial Services",
						},
						"business_and_economy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Business and Economy",
						},
						"computer_and_internet_info": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Info",
						},
						"auctions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Auctions",
						},
						"shopping": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shopping",
						},
						"cult_and_occult": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cult and Occult",
						},
						"travel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Travel",
						},
						"drugs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abused Drugs",
						},
						"adult_and_pornography": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Adult and Pornography",
						},
						"home_and_garden": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Home and Garden",
						},
						"military": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Military",
						},
						"social_network": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Social Network",
						},
						"dead_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dead Sites (db Ops only)",
						},
						"stock_advice_and_tools": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Stock Advice and Tools",
						},
						"training_and_tools": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Training and Tools",
						},
						"dating": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dating",
						},
						"sex_education": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sex Education",
						},
						"religion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Religion",
						},
						"entertainment_and_arts": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Entertainment and Arts",
						},
						"personal_sites_and_blogs": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal sites and Blogs",
						},
						"legal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Legal",
						},
						"local_information": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Local Information",
						},
						"streaming_media": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Streaming Media",
						},
						"job_search": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Job Search",
						},
						"gambling": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gambling",
						},
						"translation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Translation",
						},
						"reference_and_research": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Reference and Research",
						},
						"shareware_and_freeware": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shareware and Freeware",
						},
						"peer_to_peer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Peer to Peer",
						},
						"marijuana": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Marijuana",
						},
						"hacking": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hacking",
						},
						"games": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Games",
						},
						"philosophy_and_politics": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Philosophy and Political Advocacy",
						},
						"weapons": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Weapons",
						},
						"pay_to_surf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Pay to Surf",
						},
						"hunting_and_fishing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hunting and Fishing",
						},
						"society": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Society",
						},
						"educational_institutions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Educational Institutions",
						},
						"online_greeting_cards": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Online Greeting cards",
						},
						"sports": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sports",
						},
						"swimsuits_and_intimate_apparel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Swimsuits and Intimate Apparel",
						},
						"questionable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Questionable",
						},
						"kids": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Kids",
						},
						"hate_and_racism": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hate and Racism",
						},
						"personal_storage": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal Storage",
						},
						"violence": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Violence",
						},
						"keyloggers_and_monitoring": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Keyloggers and Monitoring",
						},
						"search_engines": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Search Engines",
						},
						"internet_portals": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Portals",
						},
						"web_advertisements": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Advertisements",
						},
						"cheating": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cheating",
						},
						"gross": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gross",
						},
						"web_based_email": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web based email",
						},
						"malware_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Malware Sites",
						},
						"phishing_and_other_fraud": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Phishing and Other Frauds",
						},
						"proxy_avoid_and_anonymizers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Proxy Avoid and Anonymizers",
						},
						"spyware_and_adware": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Spyware and Adware",
						},
						"music": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Music",
						},
						"government": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Government",
						},
						"nudity": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Nudity",
						},
						"news_and_media": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category News and Media",
						},
						"illegal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal",
						},
						"cdns": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category CDNs",
						},
						"internet_communications": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Communications",
						},
						"bot_nets": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Bot Nets",
						},
						"abortion": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abortion",
						},
						"health_and_medicine": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Health and Medicine",
						},
						"spam_urls": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category SPAM URLs",
						},
						"dynamically_generated_content": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dynamically Generated Content",
						},
						"parked_domains": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Parked Domains",
						},
						"alcohol_and_tobacco": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Alcohol and Tobacco",
						},
						"image_and_video_search": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Image and Video Search",
						},
						"fashion_and_beauty": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Fashion and Beauty",
						},
						"recreation_and_hobbies": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Recreation and Hobbies",
						},
						"motor_vehicles": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Motor Vehicles",
						},
						"web_hosting_sites": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Hosting Sites",
						},
						"self_harm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Self Harm",
						},
						"dns_over_https": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category DNS over HTTPs",
						},
						"low_thc_cannabis_products": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Low-THC Cannabis Products",
						},
						"generative_ai": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Generative AI",
						},
						"nudity_artistic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Artistic Nudity",
						},
						"illegal_pornography": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal Pornography eg. Child Sexual Abuse",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'uncategorized': uncategorized category; 'real-estate': real estate category; 'computer-and-internet-security': computer and internet security category; 'financial-services': financial services category; 'business-and-economy': business and economy category; 'computer-and-internet-info': computer and internet info category; 'auctions': auctions category; 'shopping': shopping category; 'cult-and-occult': cult and occult category; 'travel': travel category; 'drugs': drugs category; 'adult-and-pornography': adult and pornography category; 'home-and-garden': home and garden category; 'military': military category; 'social-network': social network category; 'dead-sites': dead sites category; 'stock-advice-and-tools': stock advice and tools category; 'training-and-tools': training and tools category; 'dating': dating category; 'sex-education': sex education category; 'religion': religion category; 'entertainment-and-arts': entertainment and arts category; 'personal-sites-and-blogs': personal sites and blogs category; 'legal': legal category; 'local-information': local information category; 'streaming-media': streaming media category; 'job-search': job search category; 'gambling': gambling category; 'translation': translation category; 'reference-and-research': reference and research category; 'shareware-and-freeware': shareware and freeware category; 'peer-to-peer': peer to peer category; 'marijuana': marijuana category; 'hacking': hacking category; 'games': games category; 'philosophy-and-politics': philosophy and politics category; 'weapons': weapons category; 'pay-to-surf': pay to surf category; 'hunting-and-fishing': hunting and fishing category; 'society': society category; 'educational-institutions': educational institutions category; 'online-greeting-cards': online greeting cards category; 'sports': sports category; 'swimsuits-and-intimate-apparel': swimsuits and intimate apparel category; 'questionable': questionable category; 'kids': kids category; 'hate-and-racism': hate and racism category; 'personal-storage': personal storage category; 'violence': violence category; 'keyloggers-and-monitoring': keyloggers and monitoring category; 'search-engines': search engines category; 'internet-portals': internet portals category; 'web-advertisements': web advertisements category; 'cheating': cheating category; 'gross': gross category; 'web-based-email': web based email category; 'malware-sites': malware sites category; 'phishing-and-other-fraud': phishing and other fraud category; 'proxy-avoid-and-anonymizers': proxy avoid and anonymizers category; 'spyware-and-adware': spyware and adware category; 'music': music category; 'government': government category; 'nudity': nudity category; 'news-and-media': news and media category; 'illegal': illegal category; 'CDNs': content delivery networks category; 'internet-communications': internet communications category; 'bot-nets': bot nets category; 'abortion': abortion category; 'health-and-medicine': health and medicine category; 'confirmed-SPAM-sources': confirmed SPAM sources category; 'SPAM-URLs': SPAM URLs category; 'unconfirmed-SPAM-sources': unconfirmed SPAM sources category; 'open-HTTP-proxies': open HTTP proxies category; 'dynamically-generated-content': dynamically generated content category; 'parked-domains': parked domains category; 'alcohol-and-tobacco': alcohol and tobacco category; 'private-IP-addresses': private IP addresses category; 'image-and-video-search': image and video search category; 'fashion-and-beauty': fashion and beauty category; 'recreation-and-hobbies': recreation and hobbies category; 'motor-vehicles': motor vehicles category; 'web-hosting-sites': web hosting sites category; 'food-and-dining': food and dining category; 'dummy-item': dummy item category; 'self-harm': self harm category; 'dns-over-https': dns over https category; 'low-thc-cannabis-products': low-thc cannabis products category; 'generative-ai': generative ai; 'nudity-artistic': artistic nudity; 'illegal-pornography': illegal pornography eg. child sexual abuse;",
									},
								},
							},
						},
					},
				},
			},
			"cloud_query_cache_size": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum cache size for storing cloud query results, default: 1",
			},
			"cloud_query_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disables cloud queries for URL's not present in local database(default enable)",
			},
			"database_server": {
				Type: schema.TypeString, Optional: true, Description: "BrightCloud Database Server",
			},
			"db_update_time": {
				Type: schema.TypeString, Optional: true, Description: "Time of day to update database (default: 00:00)",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BrightCloud SDK",
			},
			"intercepted_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"license": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"online_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disables online queries for license. By default it is enabled.",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "BrightCloud Query Server Listening Port(default 80)",
			},
			"proxy_server": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_host": {
							Type: schema.TypeString, Optional: true, Description: "Proxy server hostname or IP address",
						},
						"http_port": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy server HTTP port",
						},
						"https_port": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy server HTTPS port(HTTP port will be used if not configured)",
						},
						"auth_type": {
							Type: schema.TypeString, Optional: true, Default: "ntlm", Description: "'ntlm': NTLM authentication(default); 'basic': Basic authentication;",
						},
						"domain": {
							Type: schema.TypeString, Optional: true, Description: "Realm for NTLM authentication",
						},
						"username": {
							Type: schema.TypeString, Optional: true, Description: "Username for proxy authentication",
						},
						"password": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Password for proxy authentication",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "password value",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"remote_syslog_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable data plane logging to a remote syslog server",
			},
			"reputation_scope_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Reputation Scope name",
						},
						"greater_than": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"greater_trustworthy": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 81",
									},
									"greater_low_risk": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 61",
									},
									"greater_moderate_risk": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 41",
									},
									"greater_suspicious": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 21",
									},
									"greater_malicious": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is greater than or equal to 1",
									},
									"greater_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "Reputation score is greater than or equal to the customized score (1-100)",
									},
								},
							},
						},
						"less_than": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"less_trustworthy": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 100",
									},
									"less_low_risk": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 80",
									},
									"less_moderate_risk": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 60",
									},
									"less_suspicious": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 40",
									},
									"less_malicious": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reputation score is less than or equal to 20",
									},
									"less_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "Reputation score is less than or equal to a customized value (1-100)",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'trustworthy': Trustworthy level(81-100); 'low-risk': Low-risk level(61-80); 'moderate-risk': Moderate-risk level(41-60); 'suspicious': Suspicious level(21-40); 'malicious': Malicious level(1-20);",
									},
								},
							},
						},
					},
				},
			},
			"rtu_cache_size": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum cache size for storing RTU updates, default: 1",
			},
			"rtu_update_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disables real time updates(default enable)",
			},
			"rtu_update_interval": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Interval to check for real time updates if enabled in mins(default 60)",
			},
			"server": {
				Type: schema.TypeString, Optional: true, Description: "BrightCloud Query Server",
			},
			"server_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "BrightCloud Servers Timeout in seconds (default: 15s)",
			},
			"ssl_port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "BrightCloud Servers SSL Port(default 443)",
			},
			"statistics": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'db-lookup': db-lookup; 'cloud-cache-lookup': cloud-cache-lookup; 'cloud-lookup': cloud-lookup; 'rtu-lookup': rtu-lookup; 'lookup-latency': lookup-latency; 'db-mem': db-mem; 'rtu-cache-mem': rtu-cache-mem; 'lookup-cache-mem': lookup-cache-mem;",
									},
								},
							},
						},
					},
				},
			},
			"url": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management interface for all communication with BrightCloud",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"web_reputation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"intercepted_urls": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"bypassed_urls": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"url": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func resourceWebCategoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategory(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryRead(ctx, d, meta)
	}
	return diags
}

func resourceWebCategoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategory(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryRead(ctx, d, meta)
	}
	return diags
}
func resourceWebCategoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategory(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebCategoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategory(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectWebCategoryBypassedUrls3760(d []interface{}) edpt.WebCategoryBypassedUrls3760 {

	var ret edpt.WebCategoryBypassedUrls3760
	return ret
}

func getSliceWebCategoryCategoryListList(d []interface{}) []edpt.WebCategoryCategoryListList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryCategoryListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryCategoryListList
		oi.Name = in["name"].(string)
		oi.Uncategorized = in["uncategorized"].(int)
		oi.RealEstate = in["real_estate"].(int)
		oi.ComputerAndInternetSecurity = in["computer_and_internet_security"].(int)
		oi.FinancialServices = in["financial_services"].(int)
		oi.BusinessAndEconomy = in["business_and_economy"].(int)
		oi.ComputerAndInternetInfo = in["computer_and_internet_info"].(int)
		oi.Auctions = in["auctions"].(int)
		oi.Shopping = in["shopping"].(int)
		oi.CultAndOccult = in["cult_and_occult"].(int)
		oi.Travel = in["travel"].(int)
		oi.Drugs = in["drugs"].(int)
		oi.AdultAndPornography = in["adult_and_pornography"].(int)
		oi.HomeAndGarden = in["home_and_garden"].(int)
		oi.Military = in["military"].(int)
		oi.SocialNetwork = in["social_network"].(int)
		oi.DeadSites = in["dead_sites"].(int)
		oi.StockAdviceAndTools = in["stock_advice_and_tools"].(int)
		oi.TrainingAndTools = in["training_and_tools"].(int)
		oi.Dating = in["dating"].(int)
		oi.SexEducation = in["sex_education"].(int)
		oi.Religion = in["religion"].(int)
		oi.EntertainmentAndArts = in["entertainment_and_arts"].(int)
		oi.PersonalSitesAndBlogs = in["personal_sites_and_blogs"].(int)
		oi.Legal = in["legal"].(int)
		oi.LocalInformation = in["local_information"].(int)
		oi.StreamingMedia = in["streaming_media"].(int)
		oi.JobSearch = in["job_search"].(int)
		oi.Gambling = in["gambling"].(int)
		oi.Translation = in["translation"].(int)
		oi.ReferenceAndResearch = in["reference_and_research"].(int)
		oi.SharewareAndFreeware = in["shareware_and_freeware"].(int)
		oi.PeerToPeer = in["peer_to_peer"].(int)
		oi.Marijuana = in["marijuana"].(int)
		oi.Hacking = in["hacking"].(int)
		oi.Games = in["games"].(int)
		oi.PhilosophyAndPolitics = in["philosophy_and_politics"].(int)
		oi.Weapons = in["weapons"].(int)
		oi.PayToSurf = in["pay_to_surf"].(int)
		oi.HuntingAndFishing = in["hunting_and_fishing"].(int)
		oi.Society = in["society"].(int)
		oi.EducationalInstitutions = in["educational_institutions"].(int)
		oi.OnlineGreetingCards = in["online_greeting_cards"].(int)
		oi.Sports = in["sports"].(int)
		oi.SwimsuitsAndIntimateApparel = in["swimsuits_and_intimate_apparel"].(int)
		oi.Questionable = in["questionable"].(int)
		oi.Kids = in["kids"].(int)
		oi.HateAndRacism = in["hate_and_racism"].(int)
		oi.PersonalStorage = in["personal_storage"].(int)
		oi.Violence = in["violence"].(int)
		oi.KeyloggersAndMonitoring = in["keyloggers_and_monitoring"].(int)
		oi.SearchEngines = in["search_engines"].(int)
		oi.InternetPortals = in["internet_portals"].(int)
		oi.WebAdvertisements = in["web_advertisements"].(int)
		oi.Cheating = in["cheating"].(int)
		oi.Gross = in["gross"].(int)
		oi.WebBasedEmail = in["web_based_email"].(int)
		oi.MalwareSites = in["malware_sites"].(int)
		oi.PhishingAndOtherFraud = in["phishing_and_other_fraud"].(int)
		oi.ProxyAvoidAndAnonymizers = in["proxy_avoid_and_anonymizers"].(int)
		oi.SpywareAndAdware = in["spyware_and_adware"].(int)
		oi.Music = in["music"].(int)
		oi.Government = in["government"].(int)
		oi.Nudity = in["nudity"].(int)
		oi.NewsAndMedia = in["news_and_media"].(int)
		oi.Illegal = in["illegal"].(int)
		oi.Cdns = in["cdns"].(int)
		oi.InternetCommunications = in["internet_communications"].(int)
		oi.BotNets = in["bot_nets"].(int)
		oi.Abortion = in["abortion"].(int)
		oi.HealthAndMedicine = in["health_and_medicine"].(int)
		oi.SpamUrls = in["spam_urls"].(int)
		oi.DynamicallyGeneratedContent = in["dynamically_generated_content"].(int)
		oi.ParkedDomains = in["parked_domains"].(int)
		oi.AlcoholAndTobacco = in["alcohol_and_tobacco"].(int)
		oi.ImageAndVideoSearch = in["image_and_video_search"].(int)
		oi.FashionAndBeauty = in["fashion_and_beauty"].(int)
		oi.RecreationAndHobbies = in["recreation_and_hobbies"].(int)
		oi.MotorVehicles = in["motor_vehicles"].(int)
		oi.WebHostingSites = in["web_hosting_sites"].(int)
		oi.SelfHarm = in["self_harm"].(int)
		oi.DnsOverHttps = in["dns_over_https"].(int)
		oi.LowThcCannabisProducts = in["low_thc_cannabis_products"].(int)
		oi.GenerativeAi = in["generative_ai"].(int)
		oi.NudityArtistic = in["nudity_artistic"].(int)
		oi.IllegalPornography = in["illegal_pornography"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceWebCategoryCategoryListListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceWebCategoryCategoryListListSamplingEnable(d []interface{}) []edpt.WebCategoryCategoryListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.WebCategoryCategoryListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryCategoryListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryInterceptedUrls3761(d []interface{}) edpt.WebCategoryInterceptedUrls3761 {

	var ret edpt.WebCategoryInterceptedUrls3761
	return ret
}

func getObjectWebCategoryLicense3762(d []interface{}) edpt.WebCategoryLicense3762 {

	var ret edpt.WebCategoryLicense3762
	return ret
}

func getObjectWebCategoryProxyServer3763(d []interface{}) edpt.WebCategoryProxyServer3763 {

	count1 := len(d)
	var ret edpt.WebCategoryProxyServer3763
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyHost = in["proxy_host"].(string)
		ret.HttpPort = in["http_port"].(int)
		ret.HttpsPort = in["https_port"].(int)
		ret.AuthType = in["auth_type"].(string)
		ret.Domain = in["domain"].(string)
		ret.Username = in["username"].(string)
		ret.Password = in["password"].(int)
		ret.SecretString = in["secret_string"].(string)
		//omit encrypted
		//omit uuid
	}
	return ret
}

func getSliceWebCategoryReputationScopeList(d []interface{}) []edpt.WebCategoryReputationScopeList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryReputationScopeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryReputationScopeList
		oi.Name = in["name"].(string)
		oi.GreaterThan = getObjectWebCategoryReputationScopeListGreaterThan(in["greater_than"].([]interface{}))
		oi.LessThan = getObjectWebCategoryReputationScopeListLessThan(in["less_than"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceWebCategoryReputationScopeListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryReputationScopeListGreaterThan(d []interface{}) edpt.WebCategoryReputationScopeListGreaterThan {

	count1 := len(d)
	var ret edpt.WebCategoryReputationScopeListGreaterThan
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreaterTrustworthy = in["greater_trustworthy"].(int)
		ret.GreaterLowRisk = in["greater_low_risk"].(int)
		ret.GreaterModerateRisk = in["greater_moderate_risk"].(int)
		ret.GreaterSuspicious = in["greater_suspicious"].(int)
		ret.GreaterMalicious = in["greater_malicious"].(int)
		ret.GreaterThreshold = in["greater_threshold"].(int)
	}
	return ret
}

func getObjectWebCategoryReputationScopeListLessThan(d []interface{}) edpt.WebCategoryReputationScopeListLessThan {

	count1 := len(d)
	var ret edpt.WebCategoryReputationScopeListLessThan
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LessTrustworthy = in["less_trustworthy"].(int)
		ret.LessLowRisk = in["less_low_risk"].(int)
		ret.LessModerateRisk = in["less_moderate_risk"].(int)
		ret.LessSuspicious = in["less_suspicious"].(int)
		ret.LessMalicious = in["less_malicious"].(int)
		ret.LessThreshold = in["less_threshold"].(int)
	}
	return ret
}

func getSliceWebCategoryReputationScopeListSamplingEnable(d []interface{}) []edpt.WebCategoryReputationScopeListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.WebCategoryReputationScopeListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryReputationScopeListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryStatistics3764(d []interface{}) edpt.WebCategoryStatistics3764 {

	count1 := len(d)
	var ret edpt.WebCategoryStatistics3764
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceWebCategoryStatisticsSamplingEnable3765(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceWebCategoryStatisticsSamplingEnable3765(d []interface{}) []edpt.WebCategoryStatisticsSamplingEnable3765 {

	count1 := len(d)
	ret := make([]edpt.WebCategoryStatisticsSamplingEnable3765, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryStatisticsSamplingEnable3765
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryUrl3766(d []interface{}) edpt.WebCategoryUrl3766 {

	var ret edpt.WebCategoryUrl3766
	return ret
}

func getObjectWebCategoryWebReputation3767(d []interface{}) edpt.WebCategoryWebReputation3767 {

	count1 := len(d)
	var ret edpt.WebCategoryWebReputation3767
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.InterceptedUrls = getObjectWebCategoryWebReputationInterceptedUrls3768(in["intercepted_urls"].([]interface{}))
		ret.BypassedUrls = getObjectWebCategoryWebReputationBypassedUrls3769(in["bypassed_urls"].([]interface{}))
		ret.Url = getObjectWebCategoryWebReputationUrl3770(in["url"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryWebReputationInterceptedUrls3768(d []interface{}) edpt.WebCategoryWebReputationInterceptedUrls3768 {

	var ret edpt.WebCategoryWebReputationInterceptedUrls3768
	return ret
}

func getObjectWebCategoryWebReputationBypassedUrls3769(d []interface{}) edpt.WebCategoryWebReputationBypassedUrls3769 {

	var ret edpt.WebCategoryWebReputationBypassedUrls3769
	return ret
}

func getObjectWebCategoryWebReputationUrl3770(d []interface{}) edpt.WebCategoryWebReputationUrl3770 {

	var ret edpt.WebCategoryWebReputationUrl3770
	return ret
}

func dataToEndpointWebCategory(d *schema.ResourceData) edpt.WebCategory {
	var ret edpt.WebCategory
	ret.Inst.BypassedUrls = getObjectWebCategoryBypassedUrls3760(d.Get("bypassed_urls").([]interface{}))
	ret.Inst.CategoryListList = getSliceWebCategoryCategoryListList(d.Get("category_list_list").([]interface{}))
	ret.Inst.CloudQueryCacheSize = d.Get("cloud_query_cache_size").(int)
	ret.Inst.CloudQueryDisable = d.Get("cloud_query_disable").(int)
	ret.Inst.DatabaseServer = d.Get("database_server").(string)
	ret.Inst.DbUpdateTime = d.Get("db_update_time").(string)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.InterceptedUrls = getObjectWebCategoryInterceptedUrls3761(d.Get("intercepted_urls").([]interface{}))
	ret.Inst.License = getObjectWebCategoryLicense3762(d.Get("license").([]interface{}))
	ret.Inst.OnlineCheckDisable = d.Get("online_check_disable").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.ProxyServer = getObjectWebCategoryProxyServer3763(d.Get("proxy_server").([]interface{}))
	ret.Inst.RemoteSyslogEnable = d.Get("remote_syslog_enable").(int)
	ret.Inst.ReputationScopeList = getSliceWebCategoryReputationScopeList(d.Get("reputation_scope_list").([]interface{}))
	ret.Inst.RtuCacheSize = d.Get("rtu_cache_size").(int)
	ret.Inst.RtuUpdateDisable = d.Get("rtu_update_disable").(int)
	ret.Inst.RtuUpdateInterval = d.Get("rtu_update_interval").(int)
	ret.Inst.Server = d.Get("server").(string)
	ret.Inst.ServerTimeout = d.Get("server_timeout").(int)
	ret.Inst.SslPort = d.Get("ssl_port").(int)
	ret.Inst.Statistics = getObjectWebCategoryStatistics3764(d.Get("statistics").([]interface{}))
	ret.Inst.Url = getObjectWebCategoryUrl3766(d.Get("url").([]interface{}))
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	ret.Inst.WebReputation = getObjectWebCategoryWebReputation3767(d.Get("web_reputation").([]interface{}))
	return ret
}
