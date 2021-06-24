package thunder

//Thunder resource WebCategoryCategoryList

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceWebCategoryCategoryList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebCategoryCategoryListCreate,
		Update: resourceWebCategoryCategoryListUpdate,
		Read:   resourceWebCategoryCategoryListRead,
		Delete: resourceWebCategoryCategoryListDelete,
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
	}
}

func resourceWebCategoryCategoryListCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating WebCategoryCategoryList (Inside resourceWebCategoryCategoryListCreate) ")
		name1 := d.Get("name").(string)
		data := dataToWebCategoryCategoryList(d)
		logger.Println("[INFO] received formatted data from method data to WebCategoryCategoryList --")
		d.SetId(name1)
		go_thunder.PostWebCategoryCategoryList(client.Token, data, client.Host)

		return resourceWebCategoryCategoryListRead(d, meta)

	}
	return nil
}

func resourceWebCategoryCategoryListRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading WebCategoryCategoryList (Inside resourceWebCategoryCategoryListRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetWebCategoryCategoryList(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceWebCategoryCategoryListUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying WebCategoryCategoryList   (Inside resourceWebCategoryCategoryListUpdate) ")
		data := dataToWebCategoryCategoryList(d)
		logger.Println("[INFO] received formatted data from method data to WebCategoryCategoryList ")
		go_thunder.PutWebCategoryCategoryList(client.Token, name1, data, client.Host)

		return resourceWebCategoryCategoryListRead(d, meta)

	}
	return nil
}

func resourceWebCategoryCategoryListDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceWebCategoryCategoryListDelete) " + name1)
		err := go_thunder.DeleteWebCategoryCategoryList(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToWebCategoryCategoryList(d *schema.ResourceData) go_thunder.WebCategoryCategoryList {
	var vc go_thunder.WebCategoryCategoryList
	var c go_thunder.WebCategoryCategoryListInstance
	c.Name = d.Get("name").(string)
	c.Uncategorized = d.Get("uncategorized").(int)
	c.RealEstate = d.Get("real_estate").(int)
	c.ComputerAndInternetSecurity = d.Get("computer_and_internet_security").(int)
	c.FinancialServices = d.Get("financial_services").(int)
	c.BusinessAndEconomy = d.Get("business_and_economy").(int)
	c.ComputerAndInternetInfo = d.Get("computer_and_internet_info").(int)
	c.Auctions = d.Get("auctions").(int)
	c.Shopping = d.Get("shopping").(int)
	c.CultAndOccult = d.Get("cult_and_occult").(int)
	c.Travel = d.Get("travel").(int)
	c.Drugs = d.Get("drugs").(int)
	c.AdultAndPornography = d.Get("adult_and_pornography").(int)
	c.HomeAndGarden = d.Get("home_and_garden").(int)
	c.Military = d.Get("military").(int)
	c.SocialNetwork = d.Get("social_network").(int)
	c.DeadSites = d.Get("dead_sites").(int)
	c.StockAdviceAndTools = d.Get("stock_advice_and_tools").(int)
	c.TrainingAndTools = d.Get("training_and_tools").(int)
	c.Dating = d.Get("dating").(int)
	c.SexEducation = d.Get("sex_education").(int)
	c.Religion = d.Get("religion").(int)
	c.EntertainmentAndArts = d.Get("entertainment_and_arts").(int)
	c.PersonalSitesAndBlogs = d.Get("personal_sites_and_blogs").(int)
	c.Legal = d.Get("legal").(int)
	c.LocalInformation = d.Get("local_information").(int)
	c.StreamingMedia = d.Get("streaming_media").(int)
	c.JobSearch = d.Get("job_search").(int)
	c.Gambling = d.Get("gambling").(int)
	c.Translation = d.Get("translation").(int)
	c.ReferenceAndResearch = d.Get("reference_and_research").(int)
	c.SharewareAndFreeware = d.Get("shareware_and_freeware").(int)
	c.PeerToPeer = d.Get("peer_to_peer").(int)
	c.Marijuana = d.Get("marijuana").(int)
	c.Hacking = d.Get("hacking").(int)
	c.Games = d.Get("games").(int)
	c.PhilosophyAndPolitics = d.Get("philosophy_and_politics").(int)
	c.Weapons = d.Get("weapons").(int)
	c.PayToSurf = d.Get("pay_to_surf").(int)
	c.HuntingAndFishing = d.Get("hunting_and_fishing").(int)
	c.Society = d.Get("society").(int)
	c.EducationalInstitutions = d.Get("educational_institutions").(int)
	c.OnlineGreetingCards = d.Get("online_greeting_cards").(int)
	c.Sports = d.Get("sports").(int)
	c.SwimsuitsAndIntimateApparel = d.Get("swimsuits_and_intimate_apparel").(int)
	c.Questionable = d.Get("questionable").(int)
	c.Kids = d.Get("kids").(int)
	c.HateAndRacism = d.Get("hate_and_racism").(int)
	c.PersonalStorage = d.Get("personal_storage").(int)
	c.Violence = d.Get("violence").(int)
	c.KeyloggersAndMonitoring = d.Get("keyloggers_and_monitoring").(int)
	c.SearchEngines = d.Get("search_engines").(int)
	c.InternetPortals = d.Get("internet_portals").(int)
	c.WebAdvertisements = d.Get("web_advertisements").(int)
	c.Cheating = d.Get("cheating").(int)
	c.Gross = d.Get("gross").(int)
	c.WebBasedEmail = d.Get("web_based_email").(int)
	c.MalwareSites = d.Get("malware_sites").(int)
	c.PhishingAndOtherFraud = d.Get("phishing_and_other_fraud").(int)
	c.ProxyAvoidAndAnonymizers = d.Get("proxy_avoid_and_anonymizers").(int)
	c.SpywareAndAdware = d.Get("spyware_and_adware").(int)
	c.Music = d.Get("music").(int)
	c.Government = d.Get("government").(int)
	c.Nudity = d.Get("nudity").(int)
	c.NewsAndMedia = d.Get("news_and_media").(int)
	c.Illegal = d.Get("illegal").(int)
	c.Cdns = d.Get("cdns").(int)
	c.InternetCommunications = d.Get("internet_communications").(int)
	c.BotNets = d.Get("bot_nets").(int)
	c.Abortion = d.Get("abortion").(int)
	c.HealthAndMedicine = d.Get("health_and_medicine").(int)
	c.ConfirmedSpamSources = d.Get("confirmed_spam_sources").(int)
	c.SpamUrls = d.Get("spam_urls").(int)
	c.UnconfirmedSpamSources = d.Get("unconfirmed_spam_sources").(int)
	c.OpenHTTPProxies = d.Get("open_http_proxies").(int)
	c.DynamicComment = d.Get("dynamic_comment").(int)
	c.ParkedDomains = d.Get("parked_domains").(int)
	c.AlcoholAndTobacco = d.Get("alcohol_and_tobacco").(int)
	c.PrivateIPAddresses = d.Get("private_ip_addresses").(int)
	c.ImageAndVideoSearch = d.Get("image_and_video_search").(int)
	c.FashionAndBeauty = d.Get("fashion_and_beauty").(int)
	c.RecreationAndHobbies = d.Get("recreation_and_hobbies").(int)
	c.MotorVehicles = d.Get("motor_vehicles").(int)
	c.WebHostingSites = d.Get("web_hosting_sites").(int)
	c.FoodAndDining = d.Get("food_and_dining").(int)
	c.UserTag = d.Get("user_tag").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.WebCategorySamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_thunder.WebCategorySamplingEnable
		prefix1 := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix1 + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Name = c
	return vc
}
