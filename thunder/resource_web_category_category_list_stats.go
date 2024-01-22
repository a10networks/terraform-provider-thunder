package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryCategoryListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_category_list_stats`: Statistics for the object category-list\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryCategoryListStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Web Category List name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uncategorized": {
							Type: schema.TypeInt, Optional: true, Description: "uncategorized category",
						},
						"real_estate": {
							Type: schema.TypeInt, Optional: true, Description: "real estate category",
						},
						"computer_and_internet_security": {
							Type: schema.TypeInt, Optional: true, Description: "computer and internet security category",
						},
						"financial_services": {
							Type: schema.TypeInt, Optional: true, Description: "financial services category",
						},
						"business_and_economy": {
							Type: schema.TypeInt, Optional: true, Description: "business and economy category",
						},
						"computer_and_internet_info": {
							Type: schema.TypeInt, Optional: true, Description: "computer and internet info category",
						},
						"auctions": {
							Type: schema.TypeInt, Optional: true, Description: "auctions category",
						},
						"shopping": {
							Type: schema.TypeInt, Optional: true, Description: "shopping category",
						},
						"cult_and_occult": {
							Type: schema.TypeInt, Optional: true, Description: "cult and occult category",
						},
						"travel": {
							Type: schema.TypeInt, Optional: true, Description: "travel category",
						},
						"drugs": {
							Type: schema.TypeInt, Optional: true, Description: "drugs category",
						},
						"adult_and_pornography": {
							Type: schema.TypeInt, Optional: true, Description: "adult and pornography category",
						},
						"home_and_garden": {
							Type: schema.TypeInt, Optional: true, Description: "home and garden category",
						},
						"military": {
							Type: schema.TypeInt, Optional: true, Description: "military category",
						},
						"social_network": {
							Type: schema.TypeInt, Optional: true, Description: "social network category",
						},
						"dead_sites": {
							Type: schema.TypeInt, Optional: true, Description: "dead sites category",
						},
						"stock_advice_and_tools": {
							Type: schema.TypeInt, Optional: true, Description: "stock advice and tools category",
						},
						"training_and_tools": {
							Type: schema.TypeInt, Optional: true, Description: "training and tools category",
						},
						"dating": {
							Type: schema.TypeInt, Optional: true, Description: "dating category",
						},
						"sex_education": {
							Type: schema.TypeInt, Optional: true, Description: "sex education category",
						},
						"religion": {
							Type: schema.TypeInt, Optional: true, Description: "religion category",
						},
						"entertainment_and_arts": {
							Type: schema.TypeInt, Optional: true, Description: "entertainment and arts category",
						},
						"personal_sites_and_blogs": {
							Type: schema.TypeInt, Optional: true, Description: "personal sites and blogs category",
						},
						"legal": {
							Type: schema.TypeInt, Optional: true, Description: "legal category",
						},
						"local_information": {
							Type: schema.TypeInt, Optional: true, Description: "local information category",
						},
						"streaming_media": {
							Type: schema.TypeInt, Optional: true, Description: "streaming media category",
						},
						"job_search": {
							Type: schema.TypeInt, Optional: true, Description: "job search category",
						},
						"gambling": {
							Type: schema.TypeInt, Optional: true, Description: "gambling category",
						},
						"translation": {
							Type: schema.TypeInt, Optional: true, Description: "translation category",
						},
						"reference_and_research": {
							Type: schema.TypeInt, Optional: true, Description: "reference and research category",
						},
						"shareware_and_freeware": {
							Type: schema.TypeInt, Optional: true, Description: "shareware and freeware category",
						},
						"peer_to_peer": {
							Type: schema.TypeInt, Optional: true, Description: "peer to peer category",
						},
						"marijuana": {
							Type: schema.TypeInt, Optional: true, Description: "marijuana category",
						},
						"hacking": {
							Type: schema.TypeInt, Optional: true, Description: "hacking category",
						},
						"games": {
							Type: schema.TypeInt, Optional: true, Description: "games category",
						},
						"philosophy_and_politics": {
							Type: schema.TypeInt, Optional: true, Description: "philosophy and politics category",
						},
						"weapons": {
							Type: schema.TypeInt, Optional: true, Description: "weapons category",
						},
						"pay_to_surf": {
							Type: schema.TypeInt, Optional: true, Description: "pay to surf category",
						},
						"hunting_and_fishing": {
							Type: schema.TypeInt, Optional: true, Description: "hunting and fishing category",
						},
						"society": {
							Type: schema.TypeInt, Optional: true, Description: "society category",
						},
						"educational_institutions": {
							Type: schema.TypeInt, Optional: true, Description: "educational institutions category",
						},
						"online_greeting_cards": {
							Type: schema.TypeInt, Optional: true, Description: "online greeting cards category",
						},
						"sports": {
							Type: schema.TypeInt, Optional: true, Description: "sports category",
						},
						"swimsuits_and_intimate_apparel": {
							Type: schema.TypeInt, Optional: true, Description: "swimsuits and intimate apparel category",
						},
						"questionable": {
							Type: schema.TypeInt, Optional: true, Description: "questionable category",
						},
						"kids": {
							Type: schema.TypeInt, Optional: true, Description: "kids category",
						},
						"hate_and_racism": {
							Type: schema.TypeInt, Optional: true, Description: "hate and racism category",
						},
						"personal_storage": {
							Type: schema.TypeInt, Optional: true, Description: "personal storage category",
						},
						"violence": {
							Type: schema.TypeInt, Optional: true, Description: "violence category",
						},
						"keyloggers_and_monitoring": {
							Type: schema.TypeInt, Optional: true, Description: "keyloggers and monitoring category",
						},
						"search_engines": {
							Type: schema.TypeInt, Optional: true, Description: "search engines category",
						},
						"internet_portals": {
							Type: schema.TypeInt, Optional: true, Description: "internet portals category",
						},
						"web_advertisements": {
							Type: schema.TypeInt, Optional: true, Description: "web advertisements category",
						},
						"cheating": {
							Type: schema.TypeInt, Optional: true, Description: "cheating category",
						},
						"gross": {
							Type: schema.TypeInt, Optional: true, Description: "gross category",
						},
						"web_based_email": {
							Type: schema.TypeInt, Optional: true, Description: "web based email category",
						},
						"malware_sites": {
							Type: schema.TypeInt, Optional: true, Description: "malware sites category",
						},
						"phishing_and_other_fraud": {
							Type: schema.TypeInt, Optional: true, Description: "phishing and other fraud category",
						},
						"proxy_avoid_and_anonymizers": {
							Type: schema.TypeInt, Optional: true, Description: "proxy avoid and anonymizers category",
						},
						"spyware_and_adware": {
							Type: schema.TypeInt, Optional: true, Description: "spyware and adware category",
						},
						"music": {
							Type: schema.TypeInt, Optional: true, Description: "music category",
						},
						"government": {
							Type: schema.TypeInt, Optional: true, Description: "government category",
						},
						"nudity": {
							Type: schema.TypeInt, Optional: true, Description: "nudity category",
						},
						"news_and_media": {
							Type: schema.TypeInt, Optional: true, Description: "news and media category",
						},
						"illegal": {
							Type: schema.TypeInt, Optional: true, Description: "illegal category",
						},
						"cdns": {
							Type: schema.TypeInt, Optional: true, Description: "content delivery networks category",
						},
						"internet_communications": {
							Type: schema.TypeInt, Optional: true, Description: "internet communications category",
						},
						"bot_nets": {
							Type: schema.TypeInt, Optional: true, Description: "bot nets category",
						},
						"abortion": {
							Type: schema.TypeInt, Optional: true, Description: "abortion category",
						},
						"health_and_medicine": {
							Type: schema.TypeInt, Optional: true, Description: "health and medicine category",
						},
						"confirmed_spam_sources": {
							Type: schema.TypeInt, Optional: true, Description: "confirmed SPAM sources category",
						},
						"spam_urls": {
							Type: schema.TypeInt, Optional: true, Description: "SPAM URLs category",
						},
						"unconfirmed_spam_sources": {
							Type: schema.TypeInt, Optional: true, Description: "unconfirmed SPAM sources category",
						},
						"open_http_proxies": {
							Type: schema.TypeInt, Optional: true, Description: "open HTTP proxies category",
						},
						"dynamically_generated_content": {
							Type: schema.TypeInt, Optional: true, Description: "dynamically generated content category",
						},
						"parked_domains": {
							Type: schema.TypeInt, Optional: true, Description: "parked domains category",
						},
						"alcohol_and_tobacco": {
							Type: schema.TypeInt, Optional: true, Description: "alcohol and tobacco category",
						},
						"private_ip_addresses": {
							Type: schema.TypeInt, Optional: true, Description: "private IP addresses category",
						},
						"image_and_video_search": {
							Type: schema.TypeInt, Optional: true, Description: "image and video search category",
						},
						"fashion_and_beauty": {
							Type: schema.TypeInt, Optional: true, Description: "fashion and beauty category",
						},
						"recreation_and_hobbies": {
							Type: schema.TypeInt, Optional: true, Description: "recreation and hobbies category",
						},
						"motor_vehicles": {
							Type: schema.TypeInt, Optional: true, Description: "motor vehicles category",
						},
						"web_hosting_sites": {
							Type: schema.TypeInt, Optional: true, Description: "web hosting sites category",
						},
						"food_and_dining": {
							Type: schema.TypeInt, Optional: true, Description: "food and dining category",
						},
						"nudity_artistic": {
							Type: schema.TypeInt, Optional: true, Description: "nudity join entertainment and arts",
						},
						"illegal_pornography": {
							Type: schema.TypeInt, Optional: true, Description: "illegal join adult and pornography",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryCategoryListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryCategoryListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryCategoryListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryCategoryListStatsStats := setObjectWebCategoryCategoryListStatsStats(res)
		d.Set("stats", WebCategoryCategoryListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryCategoryListStatsStats(ret edpt.DataWebCategoryCategoryListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"uncategorized":                  ret.DtWebCategoryCategoryListStats.Stats.Uncategorized,
			"real_estate":                    ret.DtWebCategoryCategoryListStats.Stats.RealEstate,
			"computer_and_internet_security": ret.DtWebCategoryCategoryListStats.Stats.ComputerAndInternetSecurity,
			"financial_services":             ret.DtWebCategoryCategoryListStats.Stats.FinancialServices,
			"business_and_economy":           ret.DtWebCategoryCategoryListStats.Stats.BusinessAndEconomy,
			"computer_and_internet_info":     ret.DtWebCategoryCategoryListStats.Stats.ComputerAndInternetInfo,
			"auctions":                       ret.DtWebCategoryCategoryListStats.Stats.Auctions,
			"shopping":                       ret.DtWebCategoryCategoryListStats.Stats.Shopping,
			"cult_and_occult":                ret.DtWebCategoryCategoryListStats.Stats.CultAndOccult,
			"travel":                         ret.DtWebCategoryCategoryListStats.Stats.Travel,
			"drugs":                          ret.DtWebCategoryCategoryListStats.Stats.Drugs,
			"adult_and_pornography":          ret.DtWebCategoryCategoryListStats.Stats.AdultAndPornography,
			"home_and_garden":                ret.DtWebCategoryCategoryListStats.Stats.HomeAndGarden,
			"military":                       ret.DtWebCategoryCategoryListStats.Stats.Military,
			"social_network":                 ret.DtWebCategoryCategoryListStats.Stats.SocialNetwork,
			"dead_sites":                     ret.DtWebCategoryCategoryListStats.Stats.DeadSites,
			"stock_advice_and_tools":         ret.DtWebCategoryCategoryListStats.Stats.StockAdviceAndTools,
			"training_and_tools":             ret.DtWebCategoryCategoryListStats.Stats.TrainingAndTools,
			"dating":                         ret.DtWebCategoryCategoryListStats.Stats.Dating,
			"sex_education":                  ret.DtWebCategoryCategoryListStats.Stats.SexEducation,
			"religion":                       ret.DtWebCategoryCategoryListStats.Stats.Religion,
			"entertainment_and_arts":         ret.DtWebCategoryCategoryListStats.Stats.EntertainmentAndArts,
			"personal_sites_and_blogs":       ret.DtWebCategoryCategoryListStats.Stats.PersonalSitesAndBlogs,
			"legal":                          ret.DtWebCategoryCategoryListStats.Stats.Legal,
			"local_information":              ret.DtWebCategoryCategoryListStats.Stats.LocalInformation,
			"streaming_media":                ret.DtWebCategoryCategoryListStats.Stats.StreamingMedia,
			"job_search":                     ret.DtWebCategoryCategoryListStats.Stats.JobSearch,
			"gambling":                       ret.DtWebCategoryCategoryListStats.Stats.Gambling,
			"translation":                    ret.DtWebCategoryCategoryListStats.Stats.Translation,
			"reference_and_research":         ret.DtWebCategoryCategoryListStats.Stats.ReferenceAndResearch,
			"shareware_and_freeware":         ret.DtWebCategoryCategoryListStats.Stats.SharewareAndFreeware,
			"peer_to_peer":                   ret.DtWebCategoryCategoryListStats.Stats.PeerToPeer,
			"marijuana":                      ret.DtWebCategoryCategoryListStats.Stats.Marijuana,
			"hacking":                        ret.DtWebCategoryCategoryListStats.Stats.Hacking,
			"games":                          ret.DtWebCategoryCategoryListStats.Stats.Games,
			"philosophy_and_politics":        ret.DtWebCategoryCategoryListStats.Stats.PhilosophyAndPolitics,
			"weapons":                        ret.DtWebCategoryCategoryListStats.Stats.Weapons,
			"pay_to_surf":                    ret.DtWebCategoryCategoryListStats.Stats.PayToSurf,
			"hunting_and_fishing":            ret.DtWebCategoryCategoryListStats.Stats.HuntingAndFishing,
			"society":                        ret.DtWebCategoryCategoryListStats.Stats.Society,
			"educational_institutions":       ret.DtWebCategoryCategoryListStats.Stats.EducationalInstitutions,
			"online_greeting_cards":          ret.DtWebCategoryCategoryListStats.Stats.OnlineGreetingCards,
			"sports":                         ret.DtWebCategoryCategoryListStats.Stats.Sports,
			"swimsuits_and_intimate_apparel": ret.DtWebCategoryCategoryListStats.Stats.SwimsuitsAndIntimateApparel,
			"questionable":                   ret.DtWebCategoryCategoryListStats.Stats.Questionable,
			"kids":                           ret.DtWebCategoryCategoryListStats.Stats.Kids,
			"hate_and_racism":                ret.DtWebCategoryCategoryListStats.Stats.HateAndRacism,
			"personal_storage":               ret.DtWebCategoryCategoryListStats.Stats.PersonalStorage,
			"violence":                       ret.DtWebCategoryCategoryListStats.Stats.Violence,
			"keyloggers_and_monitoring":      ret.DtWebCategoryCategoryListStats.Stats.KeyloggersAndMonitoring,
			"search_engines":                 ret.DtWebCategoryCategoryListStats.Stats.SearchEngines,
			"internet_portals":               ret.DtWebCategoryCategoryListStats.Stats.InternetPortals,
			"web_advertisements":             ret.DtWebCategoryCategoryListStats.Stats.WebAdvertisements,
			"cheating":                       ret.DtWebCategoryCategoryListStats.Stats.Cheating,
			"gross":                          ret.DtWebCategoryCategoryListStats.Stats.Gross,
			"web_based_email":                ret.DtWebCategoryCategoryListStats.Stats.WebBasedEmail,
			"malware_sites":                  ret.DtWebCategoryCategoryListStats.Stats.MalwareSites,
			"phishing_and_other_fraud":       ret.DtWebCategoryCategoryListStats.Stats.PhishingAndOtherFraud,
			"proxy_avoid_and_anonymizers":    ret.DtWebCategoryCategoryListStats.Stats.ProxyAvoidAndAnonymizers,
			"spyware_and_adware":             ret.DtWebCategoryCategoryListStats.Stats.SpywareAndAdware,
			"music":                          ret.DtWebCategoryCategoryListStats.Stats.Music,
			"government":                     ret.DtWebCategoryCategoryListStats.Stats.Government,
			"nudity":                         ret.DtWebCategoryCategoryListStats.Stats.Nudity,
			"news_and_media":                 ret.DtWebCategoryCategoryListStats.Stats.NewsAndMedia,
			"illegal":                        ret.DtWebCategoryCategoryListStats.Stats.Illegal,
			"cdns":                           ret.DtWebCategoryCategoryListStats.Stats.Cdns,
			"internet_communications":        ret.DtWebCategoryCategoryListStats.Stats.InternetCommunications,
			"bot_nets":                       ret.DtWebCategoryCategoryListStats.Stats.BotNets,
			"abortion":                       ret.DtWebCategoryCategoryListStats.Stats.Abortion,
			"health_and_medicine":            ret.DtWebCategoryCategoryListStats.Stats.HealthAndMedicine,
			"confirmed_spam_sources":         ret.DtWebCategoryCategoryListStats.Stats.ConfirmedSpamSources,
			"spam_urls":                      ret.DtWebCategoryCategoryListStats.Stats.SpamUrls,
			"unconfirmed_spam_sources":       ret.DtWebCategoryCategoryListStats.Stats.UnconfirmedSpamSources,
			"open_http_proxies":              ret.DtWebCategoryCategoryListStats.Stats.OpenHttpProxies,
			"dynamically_generated_content":  ret.DtWebCategoryCategoryListStats.Stats.DynamicallyGeneratedContent,
			"parked_domains":                 ret.DtWebCategoryCategoryListStats.Stats.ParkedDomains,
			"alcohol_and_tobacco":            ret.DtWebCategoryCategoryListStats.Stats.AlcoholAndTobacco,
			"private_ip_addresses":           ret.DtWebCategoryCategoryListStats.Stats.PrivateIpAddresses,
			"image_and_video_search":         ret.DtWebCategoryCategoryListStats.Stats.ImageAndVideoSearch,
			"fashion_and_beauty":             ret.DtWebCategoryCategoryListStats.Stats.FashionAndBeauty,
			"recreation_and_hobbies":         ret.DtWebCategoryCategoryListStats.Stats.RecreationAndHobbies,
			"motor_vehicles":                 ret.DtWebCategoryCategoryListStats.Stats.MotorVehicles,
			"web_hosting_sites":              ret.DtWebCategoryCategoryListStats.Stats.WebHostingSites,
			"food_and_dining":                ret.DtWebCategoryCategoryListStats.Stats.FoodAndDining,
			"nudity_artistic":                ret.DtWebCategoryCategoryListStats.Stats.NudityArtistic,
			"illegal_pornography":            ret.DtWebCategoryCategoryListStats.Stats.IllegalPornography,
		},
	}
}

func getObjectWebCategoryCategoryListStatsStats(d []interface{}) edpt.WebCategoryCategoryListStatsStats {

	count1 := len(d)
	var ret edpt.WebCategoryCategoryListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Uncategorized = in["uncategorized"].(int)
		ret.RealEstate = in["real_estate"].(int)
		ret.ComputerAndInternetSecurity = in["computer_and_internet_security"].(int)
		ret.FinancialServices = in["financial_services"].(int)
		ret.BusinessAndEconomy = in["business_and_economy"].(int)
		ret.ComputerAndInternetInfo = in["computer_and_internet_info"].(int)
		ret.Auctions = in["auctions"].(int)
		ret.Shopping = in["shopping"].(int)
		ret.CultAndOccult = in["cult_and_occult"].(int)
		ret.Travel = in["travel"].(int)
		ret.Drugs = in["drugs"].(int)
		ret.AdultAndPornography = in["adult_and_pornography"].(int)
		ret.HomeAndGarden = in["home_and_garden"].(int)
		ret.Military = in["military"].(int)
		ret.SocialNetwork = in["social_network"].(int)
		ret.DeadSites = in["dead_sites"].(int)
		ret.StockAdviceAndTools = in["stock_advice_and_tools"].(int)
		ret.TrainingAndTools = in["training_and_tools"].(int)
		ret.Dating = in["dating"].(int)
		ret.SexEducation = in["sex_education"].(int)
		ret.Religion = in["religion"].(int)
		ret.EntertainmentAndArts = in["entertainment_and_arts"].(int)
		ret.PersonalSitesAndBlogs = in["personal_sites_and_blogs"].(int)
		ret.Legal = in["legal"].(int)
		ret.LocalInformation = in["local_information"].(int)
		ret.StreamingMedia = in["streaming_media"].(int)
		ret.JobSearch = in["job_search"].(int)
		ret.Gambling = in["gambling"].(int)
		ret.Translation = in["translation"].(int)
		ret.ReferenceAndResearch = in["reference_and_research"].(int)
		ret.SharewareAndFreeware = in["shareware_and_freeware"].(int)
		ret.PeerToPeer = in["peer_to_peer"].(int)
		ret.Marijuana = in["marijuana"].(int)
		ret.Hacking = in["hacking"].(int)
		ret.Games = in["games"].(int)
		ret.PhilosophyAndPolitics = in["philosophy_and_politics"].(int)
		ret.Weapons = in["weapons"].(int)
		ret.PayToSurf = in["pay_to_surf"].(int)
		ret.HuntingAndFishing = in["hunting_and_fishing"].(int)
		ret.Society = in["society"].(int)
		ret.EducationalInstitutions = in["educational_institutions"].(int)
		ret.OnlineGreetingCards = in["online_greeting_cards"].(int)
		ret.Sports = in["sports"].(int)
		ret.SwimsuitsAndIntimateApparel = in["swimsuits_and_intimate_apparel"].(int)
		ret.Questionable = in["questionable"].(int)
		ret.Kids = in["kids"].(int)
		ret.HateAndRacism = in["hate_and_racism"].(int)
		ret.PersonalStorage = in["personal_storage"].(int)
		ret.Violence = in["violence"].(int)
		ret.KeyloggersAndMonitoring = in["keyloggers_and_monitoring"].(int)
		ret.SearchEngines = in["search_engines"].(int)
		ret.InternetPortals = in["internet_portals"].(int)
		ret.WebAdvertisements = in["web_advertisements"].(int)
		ret.Cheating = in["cheating"].(int)
		ret.Gross = in["gross"].(int)
		ret.WebBasedEmail = in["web_based_email"].(int)
		ret.MalwareSites = in["malware_sites"].(int)
		ret.PhishingAndOtherFraud = in["phishing_and_other_fraud"].(int)
		ret.ProxyAvoidAndAnonymizers = in["proxy_avoid_and_anonymizers"].(int)
		ret.SpywareAndAdware = in["spyware_and_adware"].(int)
		ret.Music = in["music"].(int)
		ret.Government = in["government"].(int)
		ret.Nudity = in["nudity"].(int)
		ret.NewsAndMedia = in["news_and_media"].(int)
		ret.Illegal = in["illegal"].(int)
		ret.Cdns = in["cdns"].(int)
		ret.InternetCommunications = in["internet_communications"].(int)
		ret.BotNets = in["bot_nets"].(int)
		ret.Abortion = in["abortion"].(int)
		ret.HealthAndMedicine = in["health_and_medicine"].(int)
		ret.ConfirmedSpamSources = in["confirmed_spam_sources"].(int)
		ret.SpamUrls = in["spam_urls"].(int)
		ret.UnconfirmedSpamSources = in["unconfirmed_spam_sources"].(int)
		ret.OpenHttpProxies = in["open_http_proxies"].(int)
		ret.DynamicallyGeneratedContent = in["dynamically_generated_content"].(int)
		ret.ParkedDomains = in["parked_domains"].(int)
		ret.AlcoholAndTobacco = in["alcohol_and_tobacco"].(int)
		ret.PrivateIpAddresses = in["private_ip_addresses"].(int)
		ret.ImageAndVideoSearch = in["image_and_video_search"].(int)
		ret.FashionAndBeauty = in["fashion_and_beauty"].(int)
		ret.RecreationAndHobbies = in["recreation_and_hobbies"].(int)
		ret.MotorVehicles = in["motor_vehicles"].(int)
		ret.WebHostingSites = in["web_hosting_sites"].(int)
		ret.FoodAndDining = in["food_and_dining"].(int)
		ret.NudityArtistic = in["nudity_artistic"].(int)
		ret.IllegalPornography = in["illegal_pornography"].(int)
	}
	return ret
}

func dataToEndpointWebCategoryCategoryListStats(d *schema.ResourceData) edpt.WebCategoryCategoryListStats {
	var ret edpt.WebCategoryCategoryListStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectWebCategoryCategoryListStatsStats(d.Get("stats").([]interface{}))
	return ret
}
