package thunder

//Thunder resource WebCategory

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategory() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebCategoryCreate,
		UpdateContext: resourceWebCategoryUpdate,
		ReadContext:   resourceWebCategoryRead,
		DeleteContext: resourceWebCategoryDelete,
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"database_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cloud_query_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cloud_query_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"db_update_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rtu_update_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rtu_update_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rtu_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_mgmt_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"remote_syslog_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"proxy_server": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"http_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"https_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auth_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"secret_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"intercepted_urls": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"bypassed_urls": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"url": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"license": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"category_list_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uncategorized": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"real_estate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"computer_and_internet_security": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"financial_services": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"business_and_economy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"computer_and_internet_info": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auctions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shopping": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cult_and_occult": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"travel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"drugs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"adult_and_pornography": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"home_and_garden": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"military": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"social_network": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dead_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"stock_advice_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"training_and_tools": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sex_education": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"religion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"entertainment_and_arts": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"personal_sites_and_blogs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"legal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"local_information": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"streaming_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"job_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gambling": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"translation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reference_and_research": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shareware_and_freeware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"peer_to_peer": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"marijuana": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hacking": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"games": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"philosophy_and_politics": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"weapons": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pay_to_surf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hunting_and_fishing": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"society": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"educational_institutions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"online_greeting_cards": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sports": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"swimsuits_and_intimate_apparel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"questionable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"kids": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hate_and_racism": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"personal_storage": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"violence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"keyloggers_and_monitoring": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"search_engines": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"internet_portals": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_advertisements": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cheating": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gross": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_based_email": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"malware_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"phishing_and_other_fraud": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"proxy_avoid_and_anonymizers": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"spyware_and_adware": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"music": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"government": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"nudity": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"news_and_media": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"illegal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cdns": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"internet_communications": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bot_nets": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"abortion": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"health_and_medicine": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"confirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"spam_urls": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"unconfirmed_spam_sources": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"open_http_proxies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dynamic_comment": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"parked_domains": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alcohol_and_tobacco": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"private_ip_addresses": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"image_and_video_search": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fashion_and_beauty": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"recreation_and_hobbies": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"motor_vehicles": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"web_hosting_sites": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"food_and_dining": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"statistics": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"reputation_scope_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"greater_than": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"greater_trustworthy": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"greater_low_risk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"greater_moderate_risk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"greater_suspicious": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"greater_malicious": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"greater_threshold": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"less_than": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"less_trustworthy": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"less_low_risk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"less_moderate_risk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"less_suspicious": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"less_malicious": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"less_threshold": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"web_reputation": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"intercepted_urls": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"bypassed_urls": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"url": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
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
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating WebCategory (Inside resourceWebCategoryCreate) ")

		data := dataToWebCategory(d)
		logger.Println("[INFO] received formatted data from method data to WebCategory --")
		d.SetId("1")
		err := go_thunder.PostWebCategory(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceWebCategoryRead(ctx, d, meta)

	}
	return diags
}

func resourceWebCategoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading WebCategory (Inside resourceWebCategoryRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetWebCategory(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceWebCategoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceWebCategoryRead(ctx, d, meta)
}

func resourceWebCategoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceWebCategoryRead(ctx, d, meta)
}
func dataToWebCategory(d *schema.ResourceData) go_thunder.WebCategorySelf {
	var vc go_thunder.WebCategorySelf
	var c go_thunder.WebCategorySelfInstance
	c.Server = d.Get("server").(string)
	c.DatabaseServer = d.Get("database_server").(string)
	c.Port = d.Get("port").(int)
	c.SslPort = d.Get("ssl_port").(int)
	c.ServerTimeout = d.Get("server_timeout").(int)
	c.CloudQueryDisable = d.Get("cloud_query_disable").(int)
	c.CloudQueryCacheSize = d.Get("cloud_query_cache_size").(int)
	c.DbUpdateTime = d.Get("db_update_time").(string)
	c.RtuUpdateDisable = d.Get("rtu_update_disable").(int)
	c.RtuUpdateInterval = d.Get("rtu_update_interval").(int)
	c.RtuCacheSize = d.Get("rtu_cache_size").(int)
	c.UseMgmtPort = d.Get("use_mgmt_port").(int)
	c.RemoteSyslogEnable = d.Get("remote_syslog_enable").(int)
	c.Enable = d.Get("enable").(int)

	var obj1 go_thunder.ProxyServerWebCategory
	prefix1 := "proxy_server.0."
	obj1.ProxyHost = d.Get(prefix1 + "proxy_host").(string)
	obj1.HTTPPort = d.Get(prefix1 + "http_port").(int)
	obj1.HTTPSPort = d.Get(prefix1 + "https_port").(int)
	obj1.AuthType = d.Get(prefix1 + "auth_type").(string)
	obj1.Domain = d.Get(prefix1 + "domain").(string)
	obj1.Username = d.Get(prefix1 + "username").(string)
	obj1.Password = d.Get(prefix1 + "password").(int)
	obj1.SecretString = d.Get(prefix1 + "secret_string").(string)

	c.ProxyHost = obj1

	// var obj2 go_thunder.InterceptedUrls
	// prefix := "intercepted_urls.0."

	// c. = obj2

	// var obj3 go_thunder.BypassedUrls
	// prefix := "bypassed_urls.0."

	// c. = obj3

	// var obj4 go_thunder.Url
	// prefix := "url.0."

	// c. = obj4

	// var obj5 go_thunder.License
	// prefix := "license.0."

	// c. = obj5

	CategoryListListCount := d.Get("category_list_list.#").(int)
	c.Name = make([]go_thunder.CategoryListList, 0, CategoryListListCount)

	for i := 0; i < CategoryListListCount; i++ {
		var obj6 go_thunder.CategoryListList
		prefix6 := fmt.Sprintf("category_list_list.%d.", i)
		obj6.Name = d.Get(prefix6 + "name").(string)
		obj6.Uncategorized = d.Get(prefix6 + "uncategorized").(int)
		obj6.RealEstate = d.Get(prefix6 + "real_estate").(int)
		obj6.ComputerAndInternetSecurity = d.Get(prefix6 + "computer_and_internet_security").(int)
		obj6.FinancialServices = d.Get(prefix6 + "financial_services").(int)
		obj6.BusinessAndEconomy = d.Get(prefix6 + "business_and_economy").(int)
		obj6.ComputerAndInternetInfo = d.Get(prefix6 + "computer_and_internet_info").(int)
		obj6.Auctions = d.Get(prefix6 + "auctions").(int)
		obj6.Shopping = d.Get(prefix6 + "shopping").(int)
		obj6.CultAndOccult = d.Get(prefix6 + "cult_and_occult").(int)
		obj6.Travel = d.Get(prefix6 + "travel").(int)
		obj6.Drugs = d.Get(prefix6 + "drugs").(int)
		obj6.AdultAndPornography = d.Get(prefix6 + "adult_and_pornography").(int)
		obj6.HomeAndGarden = d.Get(prefix6 + "home_and_garden").(int)
		obj6.Military = d.Get(prefix6 + "military").(int)
		obj6.SocialNetwork = d.Get(prefix6 + "social_network").(int)
		obj6.DeadSites = d.Get(prefix6 + "dead_sites").(int)
		obj6.StockAdviceAndTools = d.Get(prefix6 + "stock_advice_and_tools").(int)
		obj6.TrainingAndTools = d.Get(prefix6 + "training_and_tools").(int)
		obj6.Dating = d.Get(prefix6 + "dating").(int)
		obj6.SexEducation = d.Get(prefix6 + "sex_education").(int)
		obj6.Religion = d.Get(prefix6 + "religion").(int)
		obj6.EntertainmentAndArts = d.Get(prefix6 + "entertainment_and_arts").(int)
		obj6.PersonalSitesAndBlogs = d.Get(prefix6 + "personal_sites_and_blogs").(int)
		obj6.Legal = d.Get(prefix6 + "legal").(int)
		obj6.LocalInformation = d.Get(prefix6 + "local_information").(int)
		obj6.StreamingMedia = d.Get(prefix6 + "streaming_media").(int)
		obj6.JobSearch = d.Get(prefix6 + "job_search").(int)
		obj6.Gambling = d.Get(prefix6 + "gambling").(int)
		obj6.Translation = d.Get(prefix6 + "translation").(int)
		obj6.ReferenceAndResearch = d.Get(prefix6 + "reference_and_research").(int)
		obj6.SharewareAndFreeware = d.Get(prefix6 + "shareware_and_freeware").(int)
		obj6.PeerToPeer = d.Get(prefix6 + "peer_to_peer").(int)
		obj6.Marijuana = d.Get(prefix6 + "marijuana").(int)
		obj6.Hacking = d.Get(prefix6 + "hacking").(int)
		obj6.Games = d.Get(prefix6 + "games").(int)
		obj6.PhilosophyAndPolitics = d.Get(prefix6 + "philosophy_and_politics").(int)
		obj6.Weapons = d.Get(prefix6 + "weapons").(int)
		obj6.PayToSurf = d.Get(prefix6 + "pay_to_surf").(int)
		obj6.HuntingAndFishing = d.Get(prefix6 + "hunting_and_fishing").(int)
		obj6.Society = d.Get(prefix6 + "society").(int)
		obj6.EducationalInstitutions = d.Get(prefix6 + "educational_institutions").(int)
		obj6.OnlineGreetingCards = d.Get(prefix6 + "online_greeting_cards").(int)
		obj6.Sports = d.Get(prefix6 + "sports").(int)
		obj6.SwimsuitsAndIntimateApparel = d.Get(prefix6 + "swimsuits_and_intimate_apparel").(int)
		obj6.Questionable = d.Get(prefix6 + "questionable").(int)
		obj6.Kids = d.Get(prefix6 + "kids").(int)
		obj6.HateAndRacism = d.Get(prefix6 + "hate_and_racism").(int)
		obj6.PersonalStorage = d.Get(prefix6 + "personal_storage").(int)
		obj6.Violence = d.Get(prefix6 + "violence").(int)
		obj6.KeyloggersAndMonitoring = d.Get(prefix6 + "keyloggers_and_monitoring").(int)
		obj6.SearchEngines = d.Get(prefix6 + "search_engines").(int)
		obj6.InternetPortals = d.Get(prefix6 + "internet_portals").(int)
		obj6.WebAdvertisements = d.Get(prefix6 + "web_advertisements").(int)
		obj6.Cheating = d.Get(prefix6 + "cheating").(int)
		obj6.Gross = d.Get(prefix6 + "gross").(int)
		obj6.WebBasedEmail = d.Get(prefix6 + "web_based_email").(int)
		obj6.MalwareSites = d.Get(prefix6 + "malware_sites").(int)
		obj6.PhishingAndOtherFraud = d.Get(prefix6 + "phishing_and_other_fraud").(int)
		obj6.ProxyAvoidAndAnonymizers = d.Get(prefix6 + "proxy_avoid_and_anonymizers").(int)
		obj6.SpywareAndAdware = d.Get(prefix6 + "spyware_and_adware").(int)
		obj6.Music = d.Get(prefix6 + "music").(int)
		obj6.Government = d.Get(prefix6 + "government").(int)
		obj6.Nudity = d.Get(prefix6 + "nudity").(int)
		obj6.NewsAndMedia = d.Get(prefix6 + "news_and_media").(int)
		obj6.Illegal = d.Get(prefix6 + "illegal").(int)
		obj6.Cdns = d.Get(prefix6 + "cdns").(int)
		obj6.InternetCommunications = d.Get(prefix6 + "internet_communications").(int)
		obj6.BotNets = d.Get(prefix6 + "bot_nets").(int)
		obj6.Abortion = d.Get(prefix6 + "abortion").(int)
		obj6.HealthAndMedicine = d.Get(prefix6 + "health_and_medicine").(int)
		obj6.ConfirmedSpamSources = d.Get(prefix6 + "confirmed_spam_sources").(int)
		obj6.SpamUrls = d.Get(prefix6 + "spam_urls").(int)
		obj6.UnconfirmedSpamSources = d.Get(prefix6 + "unconfirmed_spam_sources").(int)
		obj6.OpenHTTPProxies = d.Get(prefix6 + "open_http_proxies").(int)
		obj6.DynamicComment = d.Get(prefix6 + "dynamic_comment").(int)
		obj6.ParkedDomains = d.Get(prefix6 + "parked_domains").(int)
		obj6.AlcoholAndTobacco = d.Get(prefix6 + "alcohol_and_tobacco").(int)
		obj6.PrivateIPAddresses = d.Get(prefix6 + "private_ip_addresses").(int)
		obj6.ImageAndVideoSearch = d.Get(prefix6 + "image_and_video_search").(int)
		obj6.FashionAndBeauty = d.Get(prefix6 + "fashion_and_beauty").(int)
		obj6.RecreationAndHobbies = d.Get(prefix6 + "recreation_and_hobbies").(int)
		obj6.MotorVehicles = d.Get(prefix6 + "motor_vehicles").(int)
		obj6.WebHostingSites = d.Get(prefix6 + "web_hosting_sites").(int)
		obj6.FoodAndDining = d.Get(prefix6 + "food_and_dining").(int)
		obj6.UserTag = d.Get(prefix6 + "user_tag").(string)

		SamplingEnableCount := d.Get(prefix6 + "sampling_enable.#").(int)
		obj6.Counters1 = make([]go_thunder.SamplingEnableWebCategory, 0, SamplingEnableCount)

		for i := 0; i < SamplingEnableCount; i++ {
			var obj6_1 go_thunder.SamplingEnableWebCategory
			prefix6_1 := fmt.Sprintf(prefix6+"sampling_enable.%d.", i)
			obj6_1.Counters1 = d.Get(prefix6_1 + "counters1").(string)
			obj6.Counters1 = append(obj6.Counters1, obj6_1)
		}

		c.Name = append(c.Name, obj6)
	}

	var obj7 go_thunder.Statistics
	prefix7 := "statistics.0."

	SamplingEnableCount := d.Get(prefix7 + "sampling_enable.#").(int)
	obj7.Counters1 = make([]go_thunder.SamplingEnableWebCategory, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj7_1 go_thunder.SamplingEnableWebCategory
		prefix7_1 := fmt.Sprintf(prefix7+"sampling_enable.%d.", i)
		obj7_1.Counters1 = d.Get(prefix7_1 + "counters1").(string)
		obj7.Counters1 = append(obj7.Counters1, obj7_1)
	}

	c.Counters1 = obj7

	ReputationScopeListCount := d.Get("reputation_scope_list.#").(int)
	c.GreaterTrustworthy = make([]go_thunder.ReputationScopeList, 0, ReputationScopeListCount)

	for i := 0; i < ReputationScopeListCount; i++ {
		var obj8 go_thunder.ReputationScopeList
		prefix8 := fmt.Sprintf("reputation_scope_list.%d.", i)
		obj8.Name = d.Get(prefix8 + "name").(string)

		var obj8_1 go_thunder.GreaterThan
		prefix8_1 := prefix8 + "greater_than.0."
		obj8_1.GreaterTrustworthy = d.Get(prefix8_1 + "greater_trustworthy").(int)
		obj8_1.GreaterLowRisk = d.Get(prefix8_1 + "greater_low_risk").(int)
		obj8_1.GreaterModerateRisk = d.Get(prefix8_1 + "greater_moderate_risk").(int)
		obj8_1.GreaterSuspicious = d.Get(prefix8_1 + "greater_suspicious").(int)
		obj8_1.GreaterMalicious = d.Get(prefix8_1 + "greater_malicious").(int)
		obj8_1.GreaterThreshold = d.Get(prefix8_1 + "greater_threshold").(int)

		obj8.GreaterTrustworthy = obj8_1

		var obj8_2 go_thunder.LessThan
		prefix8_2 := prefix8 + "less_than.0."
		obj8_2.LessTrustworthy = d.Get(prefix8_2 + "less_trustworthy").(int)
		obj8_2.LessLowRisk = d.Get(prefix8_2 + "less_low_risk").(int)
		obj8_2.LessModerateRisk = d.Get(prefix8_2 + "less_moderate_risk").(int)
		obj8_2.LessSuspicious = d.Get(prefix8_2 + "less_suspicious").(int)
		obj8_2.LessMalicious = d.Get(prefix8_2 + "less_malicious").(int)
		obj8_2.LessThreshold = d.Get(prefix8_2 + "less_threshold").(int)

		obj8.LessTrustworthy = obj8_2

		obj8.UserTag = d.Get(prefix8 + "user_tag").(string)

		SamplingEnableCount := d.Get(prefix8 + "sampling_enable.#").(int)
		obj8.Counters1 = make([]go_thunder.SamplingEnableWebCategory, 0, SamplingEnableCount)

		for i := 0; i < SamplingEnableCount; i++ {
			var obj8_3 go_thunder.SamplingEnableWebCategory
			prefix8_3 := fmt.Sprintf(prefix8+"sampling_enable.%d.", i)
			obj8_3.Counters1 = d.Get(prefix8_3 + "counters1").(string)
			obj8.Counters1 = append(obj8.Counters1, obj8_3)
		}

		c.GreaterTrustworthy = append(c.GreaterTrustworthy, obj8)
	}

	// var obj9 go_thunder.WebReputation
	// prefix9 := "web_reputation.0."

	// var obj9_1 go_thunder.InterceptedUrls
	// prefix9_1 :=prefix9 + "intercepted_urls.0."

	// obj9. = obj1

	// var obj2 go_thunder.BypassedUrls
	// prefix := "bypassed_urls.0."

	// obj9. = obj2

	// var obj3 go_thunder.Url
	// prefix := "url.0."

	// obj9. = obj3

	//c. = obj9

	vc.Server = c
	return vc
}
