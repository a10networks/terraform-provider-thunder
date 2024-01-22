package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryCategoryList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_category_category_list`: List of web categories\n\n__PLACEHOLDER__",
		CreateContext: resourceWebCategoryCategoryListCreate,
		UpdateContext: resourceWebCategoryCategoryListUpdate,
		ReadContext:   resourceWebCategoryCategoryListRead,
		DeleteContext: resourceWebCategoryCategoryListDelete,

		Schema: map[string]*schema.Schema{
			"abortion": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abortion",
			},
			"adult_and_pornography": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Adult and Pornography",
			},
			"alcohol_and_tobacco": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Alcohol and Tobacco",
			},
			"auctions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Auctions",
			},
			"bot_nets": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Bot Nets",
			},
			"business_and_economy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Business and Economy",
			},
			"cdns": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category CDNs",
			},
			"cheating": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cheating",
			},
			"computer_and_internet_info": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Info",
			},
			"computer_and_internet_security": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Computer and Internet Security",
			},
			"cult_and_occult": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Cult and Occult",
			},
			"dating": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dating",
			},
			"dead_sites": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Dead Sites (db Ops only)",
			},
			"drugs": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Abused Drugs",
			},
			"dynamically_generated_content": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dynamically Generated Content",
			},
			"educational_institutions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Educational Institutions",
			},
			"entertainment_and_arts": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Entertainment and Arts",
			},
			"fashion_and_beauty": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Fashion and Beauty",
			},
			"financial_services": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Financial Services",
			},
			"gambling": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gambling",
			},
			"games": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Games",
			},
			"government": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Government",
			},
			"gross": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Gross",
			},
			"hacking": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hacking",
			},
			"hate_and_racism": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hate and Racism",
			},
			"health_and_medicine": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Health and Medicine",
			},
			"home_and_garden": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Home and Garden",
			},
			"hunting_and_fishing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Hunting and Fishing",
			},
			"illegal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal",
			},
			"illegal_pornography": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Illegal join Adult and Pornography",
			},
			"image_and_video_search": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Image and Video Search",
			},
			"internet_communications": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Communications",
			},
			"internet_portals": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Internet Portals",
			},
			"job_search": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Job Search",
			},
			"keyloggers_and_monitoring": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Keyloggers and Monitoring",
			},
			"kids": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Kids",
			},
			"legal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Legal",
			},
			"local_information": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Local Information",
			},
			"malware_sites": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Malware Sites",
			},
			"marijuana": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Marijuana",
			},
			"military": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Military",
			},
			"motor_vehicles": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Motor Vehicles",
			},
			"music": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Music",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Web Category List name",
			},
			"news_and_media": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category News and Media",
			},
			"nudity": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Nudity",
			},
			"nudity_artistic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Nudity join Entertainment and Arts",
			},
			"online_greeting_cards": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Online Greeting cards",
			},
			"parked_domains": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Parked Domains",
			},
			"pay_to_surf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Pay to Surf",
			},
			"peer_to_peer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Peer to Peer",
			},
			"personal_sites_and_blogs": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal sites and Blogs",
			},
			"personal_storage": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Personal Storage",
			},
			"philosophy_and_politics": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Philosophy and Political Advocacy",
			},
			"phishing_and_other_fraud": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Phishing and Other Frauds",
			},
			"proxy_avoid_and_anonymizers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Proxy Avoid and Anonymizers",
			},
			"questionable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Questionable",
			},
			"real_estate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Real Estate",
			},
			"recreation_and_hobbies": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Recreation and Hobbies",
			},
			"reference_and_research": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Reference and Research",
			},
			"religion": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Religion",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'uncategorized': uncategorized category; 'real-estate': real estate category; 'computer-and-internet-security': computer and internet security category; 'financial-services': financial services category; 'business-and-economy': business and economy category; 'computer-and-internet-info': computer and internet info category; 'auctions': auctions category; 'shopping': shopping category; 'cult-and-occult': cult and occult category; 'travel': travel category; 'drugs': drugs category; 'adult-and-pornography': adult and pornography category; 'home-and-garden': home and garden category; 'military': military category; 'social-network': social network category; 'dead-sites': dead sites category; 'stock-advice-and-tools': stock advice and tools category; 'training-and-tools': training and tools category; 'dating': dating category; 'sex-education': sex education category; 'religion': religion category; 'entertainment-and-arts': entertainment and arts category; 'personal-sites-and-blogs': personal sites and blogs category; 'legal': legal category; 'local-information': local information category; 'streaming-media': streaming media category; 'job-search': job search category; 'gambling': gambling category; 'translation': translation category; 'reference-and-research': reference and research category; 'shareware-and-freeware': shareware and freeware category; 'peer-to-peer': peer to peer category; 'marijuana': marijuana category; 'hacking': hacking category; 'games': games category; 'philosophy-and-politics': philosophy and politics category; 'weapons': weapons category; 'pay-to-surf': pay to surf category; 'hunting-and-fishing': hunting and fishing category; 'society': society category; 'educational-institutions': educational institutions category; 'online-greeting-cards': online greeting cards category; 'sports': sports category; 'swimsuits-and-intimate-apparel': swimsuits and intimate apparel category; 'questionable': questionable category; 'kids': kids category; 'hate-and-racism': hate and racism category; 'personal-storage': personal storage category; 'violence': violence category; 'keyloggers-and-monitoring': keyloggers and monitoring category; 'search-engines': search engines category; 'internet-portals': internet portals category; 'web-advertisements': web advertisements category; 'cheating': cheating category; 'gross': gross category; 'web-based-email': web based email category; 'malware-sites': malware sites category; 'phishing-and-other-fraud': phishing and other fraud category; 'proxy-avoid-and-anonymizers': proxy avoid and anonymizers category; 'spyware-and-adware': spyware and adware category; 'music': music category; 'government': government category; 'nudity': nudity category; 'news-and-media': news and media category; 'illegal': illegal category; 'CDNs': content delivery networks category; 'internet-communications': internet communications category; 'bot-nets': bot nets category; 'abortion': abortion category; 'health-and-medicine': health and medicine category; 'confirmed-SPAM-sources': confirmed SPAM sources category; 'SPAM-URLs': SPAM URLs category; 'unconfirmed-SPAM-sources': unconfirmed SPAM sources category; 'open-HTTP-proxies': open HTTP proxies category; 'dynamically-generated-content': dynamically generated content category; 'parked-domains': parked domains category; 'alcohol-and-tobacco': alcohol and tobacco category; 'private-IP-addresses': private IP addresses category; 'image-and-video-search': image and video search category; 'fashion-and-beauty': fashion and beauty category; 'recreation-and-hobbies': recreation and hobbies category; 'motor-vehicles': motor vehicles category; 'web-hosting-sites': web hosting sites category; 'food-and-dining': food and dining category; 'nudity-artistic': nudity join entertainment and arts; 'illegal-pornography': illegal join adult and pornography;",
						},
					},
				},
			},
			"search_engines": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Search Engines",
			},
			"sex_education": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sex Education",
			},
			"shareware_and_freeware": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shareware and Freeware",
			},
			"shopping": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Shopping",
			},
			"social_network": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Social Network",
			},
			"society": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Society",
			},
			"spam_urls": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category SPAM URLs",
			},
			"sports": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Sports",
			},
			"spyware_and_adware": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Spyware and Adware",
			},
			"stock_advice_and_tools": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Stock Advice and Tools",
			},
			"streaming_media": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Streaming Media",
			},
			"swimsuits_and_intimate_apparel": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Swimsuits and Intimate Apparel",
			},
			"training_and_tools": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Training and Tools",
			},
			"translation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Translation",
			},
			"travel": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Travel",
			},
			"uncategorized": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Uncategorized URLs",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"violence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Violence",
			},
			"weapons": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Weapons",
			},
			"web_advertisements": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Advertisements",
			},
			"web_based_email": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web based email",
			},
			"web_hosting_sites": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Category Web Hosting Sites",
			},
		},
	}
}
func resourceWebCategoryCategoryListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryCategoryListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryCategoryList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryCategoryListRead(ctx, d, meta)
	}
	return diags
}

func resourceWebCategoryCategoryListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryCategoryListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryCategoryList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebCategoryCategoryListRead(ctx, d, meta)
	}
	return diags
}
func resourceWebCategoryCategoryListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryCategoryListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryCategoryList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebCategoryCategoryListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryCategoryListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryCategoryList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceWebCategoryCategoryListSamplingEnable(d []interface{}) []edpt.WebCategoryCategoryListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.WebCategoryCategoryListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryCategoryListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryCategoryList(d *schema.ResourceData) edpt.WebCategoryCategoryList {
	var ret edpt.WebCategoryCategoryList
	ret.Inst.Abortion = d.Get("abortion").(int)
	ret.Inst.AdultAndPornography = d.Get("adult_and_pornography").(int)
	ret.Inst.AlcoholAndTobacco = d.Get("alcohol_and_tobacco").(int)
	ret.Inst.Auctions = d.Get("auctions").(int)
	ret.Inst.BotNets = d.Get("bot_nets").(int)
	ret.Inst.BusinessAndEconomy = d.Get("business_and_economy").(int)
	ret.Inst.Cdns = d.Get("cdns").(int)
	ret.Inst.Cheating = d.Get("cheating").(int)
	ret.Inst.ComputerAndInternetInfo = d.Get("computer_and_internet_info").(int)
	ret.Inst.ComputerAndInternetSecurity = d.Get("computer_and_internet_security").(int)
	ret.Inst.CultAndOccult = d.Get("cult_and_occult").(int)
	ret.Inst.Dating = d.Get("dating").(int)
	ret.Inst.DeadSites = d.Get("dead_sites").(int)
	ret.Inst.Drugs = d.Get("drugs").(int)
	ret.Inst.DynamicallyGeneratedContent = d.Get("dynamically_generated_content").(int)
	ret.Inst.EducationalInstitutions = d.Get("educational_institutions").(int)
	ret.Inst.EntertainmentAndArts = d.Get("entertainment_and_arts").(int)
	ret.Inst.FashionAndBeauty = d.Get("fashion_and_beauty").(int)
	ret.Inst.FinancialServices = d.Get("financial_services").(int)
	ret.Inst.Gambling = d.Get("gambling").(int)
	ret.Inst.Games = d.Get("games").(int)
	ret.Inst.Government = d.Get("government").(int)
	ret.Inst.Gross = d.Get("gross").(int)
	ret.Inst.Hacking = d.Get("hacking").(int)
	ret.Inst.HateAndRacism = d.Get("hate_and_racism").(int)
	ret.Inst.HealthAndMedicine = d.Get("health_and_medicine").(int)
	ret.Inst.HomeAndGarden = d.Get("home_and_garden").(int)
	ret.Inst.HuntingAndFishing = d.Get("hunting_and_fishing").(int)
	ret.Inst.Illegal = d.Get("illegal").(int)
	ret.Inst.IllegalPornography = d.Get("illegal_pornography").(int)
	ret.Inst.ImageAndVideoSearch = d.Get("image_and_video_search").(int)
	ret.Inst.InternetCommunications = d.Get("internet_communications").(int)
	ret.Inst.InternetPortals = d.Get("internet_portals").(int)
	ret.Inst.JobSearch = d.Get("job_search").(int)
	ret.Inst.KeyloggersAndMonitoring = d.Get("keyloggers_and_monitoring").(int)
	ret.Inst.Kids = d.Get("kids").(int)
	ret.Inst.Legal = d.Get("legal").(int)
	ret.Inst.LocalInformation = d.Get("local_information").(int)
	ret.Inst.MalwareSites = d.Get("malware_sites").(int)
	ret.Inst.Marijuana = d.Get("marijuana").(int)
	ret.Inst.Military = d.Get("military").(int)
	ret.Inst.MotorVehicles = d.Get("motor_vehicles").(int)
	ret.Inst.Music = d.Get("music").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NewsAndMedia = d.Get("news_and_media").(int)
	ret.Inst.Nudity = d.Get("nudity").(int)
	ret.Inst.NudityArtistic = d.Get("nudity_artistic").(int)
	ret.Inst.OnlineGreetingCards = d.Get("online_greeting_cards").(int)
	ret.Inst.ParkedDomains = d.Get("parked_domains").(int)
	ret.Inst.PayToSurf = d.Get("pay_to_surf").(int)
	ret.Inst.PeerToPeer = d.Get("peer_to_peer").(int)
	ret.Inst.PersonalSitesAndBlogs = d.Get("personal_sites_and_blogs").(int)
	ret.Inst.PersonalStorage = d.Get("personal_storage").(int)
	ret.Inst.PhilosophyAndPolitics = d.Get("philosophy_and_politics").(int)
	ret.Inst.PhishingAndOtherFraud = d.Get("phishing_and_other_fraud").(int)
	ret.Inst.ProxyAvoidAndAnonymizers = d.Get("proxy_avoid_and_anonymizers").(int)
	ret.Inst.Questionable = d.Get("questionable").(int)
	ret.Inst.RealEstate = d.Get("real_estate").(int)
	ret.Inst.RecreationAndHobbies = d.Get("recreation_and_hobbies").(int)
	ret.Inst.ReferenceAndResearch = d.Get("reference_and_research").(int)
	ret.Inst.Religion = d.Get("religion").(int)
	ret.Inst.SamplingEnable = getSliceWebCategoryCategoryListSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SearchEngines = d.Get("search_engines").(int)
	ret.Inst.SexEducation = d.Get("sex_education").(int)
	ret.Inst.SharewareAndFreeware = d.Get("shareware_and_freeware").(int)
	ret.Inst.Shopping = d.Get("shopping").(int)
	ret.Inst.SocialNetwork = d.Get("social_network").(int)
	ret.Inst.Society = d.Get("society").(int)
	ret.Inst.SpamUrls = d.Get("spam_urls").(int)
	ret.Inst.Sports = d.Get("sports").(int)
	ret.Inst.SpywareAndAdware = d.Get("spyware_and_adware").(int)
	ret.Inst.StockAdviceAndTools = d.Get("stock_advice_and_tools").(int)
	ret.Inst.StreamingMedia = d.Get("streaming_media").(int)
	ret.Inst.SwimsuitsAndIntimateApparel = d.Get("swimsuits_and_intimate_apparel").(int)
	ret.Inst.TrainingAndTools = d.Get("training_and_tools").(int)
	ret.Inst.Translation = d.Get("translation").(int)
	ret.Inst.Travel = d.Get("travel").(int)
	ret.Inst.Uncategorized = d.Get("uncategorized").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Violence = d.Get("violence").(int)
	ret.Inst.Weapons = d.Get("weapons").(int)
	ret.Inst.WebAdvertisements = d.Get("web_advertisements").(int)
	ret.Inst.WebBasedEmail = d.Get("web_based_email").(int)
	ret.Inst.WebHostingSites = d.Get("web_hosting_sites").(int)
	return ret
}
