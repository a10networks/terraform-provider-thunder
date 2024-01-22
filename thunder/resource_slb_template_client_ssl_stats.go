package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateClientSslStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_client_ssl_stats`: Statistics for the object client-ssl\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateClientSslStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Client SSL Template Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Type: schema.TypeInt, Optional: true, Description: "nudity join Entertainment and Arts",
						},
						"illegal_pornography": {
							Type: schema.TypeInt, Optional: true, Description: "illegal join Adult and Pornography",
						},
						"uncategorised": {
							Type: schema.TypeInt, Optional: true, Description: "uncategorised",
						},
						"other_category": {
							Type: schema.TypeInt, Optional: true, Description: "other category",
						},
						"trustworthy": {
							Type: schema.TypeInt, Optional: true, Description: "Trustworthy level(81-100)",
						},
						"low_risk": {
							Type: schema.TypeInt, Optional: true, Description: "Low-risk level(61-80)",
						},
						"moderate_risk": {
							Type: schema.TypeInt, Optional: true, Description: "Moderate-risk level(41-60)",
						},
						"suspicious": {
							Type: schema.TypeInt, Optional: true, Description: "Suspicious level(21-40)",
						},
						"malicious": {
							Type: schema.TypeInt, Optional: true, Description: "Malicious level(1-20)",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateClientSslStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSslStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateClientSslStatsStats := setObjectSlbTemplateClientSslStatsStats(res)
		d.Set("stats", SlbTemplateClientSslStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateClientSslStatsStats(ret edpt.DataSlbTemplateClientSslStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"real_estate":                    ret.DtSlbTemplateClientSslStats.Stats.RealEstate,
			"computer_and_internet_security": ret.DtSlbTemplateClientSslStats.Stats.ComputerAndInternetSecurity,
			"financial_services":             ret.DtSlbTemplateClientSslStats.Stats.FinancialServices,
			"business_and_economy":           ret.DtSlbTemplateClientSslStats.Stats.BusinessAndEconomy,
			"computer_and_internet_info":     ret.DtSlbTemplateClientSslStats.Stats.ComputerAndInternetInfo,
			"auctions":                       ret.DtSlbTemplateClientSslStats.Stats.Auctions,
			"shopping":                       ret.DtSlbTemplateClientSslStats.Stats.Shopping,
			"cult_and_occult":                ret.DtSlbTemplateClientSslStats.Stats.CultAndOccult,
			"travel":                         ret.DtSlbTemplateClientSslStats.Stats.Travel,
			"drugs":                          ret.DtSlbTemplateClientSslStats.Stats.Drugs,
			"adult_and_pornography":          ret.DtSlbTemplateClientSslStats.Stats.AdultAndPornography,
			"home_and_garden":                ret.DtSlbTemplateClientSslStats.Stats.HomeAndGarden,
			"military":                       ret.DtSlbTemplateClientSslStats.Stats.Military,
			"social_network":                 ret.DtSlbTemplateClientSslStats.Stats.SocialNetwork,
			"dead_sites":                     ret.DtSlbTemplateClientSslStats.Stats.DeadSites,
			"stock_advice_and_tools":         ret.DtSlbTemplateClientSslStats.Stats.StockAdviceAndTools,
			"training_and_tools":             ret.DtSlbTemplateClientSslStats.Stats.TrainingAndTools,
			"dating":                         ret.DtSlbTemplateClientSslStats.Stats.Dating,
			"sex_education":                  ret.DtSlbTemplateClientSslStats.Stats.SexEducation,
			"religion":                       ret.DtSlbTemplateClientSslStats.Stats.Religion,
			"entertainment_and_arts":         ret.DtSlbTemplateClientSslStats.Stats.EntertainmentAndArts,
			"personal_sites_and_blogs":       ret.DtSlbTemplateClientSslStats.Stats.PersonalSitesAndBlogs,
			"legal":                          ret.DtSlbTemplateClientSslStats.Stats.Legal,
			"local_information":              ret.DtSlbTemplateClientSslStats.Stats.LocalInformation,
			"streaming_media":                ret.DtSlbTemplateClientSslStats.Stats.StreamingMedia,
			"job_search":                     ret.DtSlbTemplateClientSslStats.Stats.JobSearch,
			"gambling":                       ret.DtSlbTemplateClientSslStats.Stats.Gambling,
			"translation":                    ret.DtSlbTemplateClientSslStats.Stats.Translation,
			"reference_and_research":         ret.DtSlbTemplateClientSslStats.Stats.ReferenceAndResearch,
			"shareware_and_freeware":         ret.DtSlbTemplateClientSslStats.Stats.SharewareAndFreeware,
			"peer_to_peer":                   ret.DtSlbTemplateClientSslStats.Stats.PeerToPeer,
			"marijuana":                      ret.DtSlbTemplateClientSslStats.Stats.Marijuana,
			"hacking":                        ret.DtSlbTemplateClientSslStats.Stats.Hacking,
			"games":                          ret.DtSlbTemplateClientSslStats.Stats.Games,
			"philosophy_and_politics":        ret.DtSlbTemplateClientSslStats.Stats.PhilosophyAndPolitics,
			"weapons":                        ret.DtSlbTemplateClientSslStats.Stats.Weapons,
			"pay_to_surf":                    ret.DtSlbTemplateClientSslStats.Stats.PayToSurf,
			"hunting_and_fishing":            ret.DtSlbTemplateClientSslStats.Stats.HuntingAndFishing,
			"society":                        ret.DtSlbTemplateClientSslStats.Stats.Society,
			"educational_institutions":       ret.DtSlbTemplateClientSslStats.Stats.EducationalInstitutions,
			"online_greeting_cards":          ret.DtSlbTemplateClientSslStats.Stats.OnlineGreetingCards,
			"sports":                         ret.DtSlbTemplateClientSslStats.Stats.Sports,
			"swimsuits_and_intimate_apparel": ret.DtSlbTemplateClientSslStats.Stats.SwimsuitsAndIntimateApparel,
			"questionable":                   ret.DtSlbTemplateClientSslStats.Stats.Questionable,
			"kids":                           ret.DtSlbTemplateClientSslStats.Stats.Kids,
			"hate_and_racism":                ret.DtSlbTemplateClientSslStats.Stats.HateAndRacism,
			"personal_storage":               ret.DtSlbTemplateClientSslStats.Stats.PersonalStorage,
			"violence":                       ret.DtSlbTemplateClientSslStats.Stats.Violence,
			"keyloggers_and_monitoring":      ret.DtSlbTemplateClientSslStats.Stats.KeyloggersAndMonitoring,
			"search_engines":                 ret.DtSlbTemplateClientSslStats.Stats.SearchEngines,
			"internet_portals":               ret.DtSlbTemplateClientSslStats.Stats.InternetPortals,
			"web_advertisements":             ret.DtSlbTemplateClientSslStats.Stats.WebAdvertisements,
			"cheating":                       ret.DtSlbTemplateClientSslStats.Stats.Cheating,
			"gross":                          ret.DtSlbTemplateClientSslStats.Stats.Gross,
			"web_based_email":                ret.DtSlbTemplateClientSslStats.Stats.WebBasedEmail,
			"malware_sites":                  ret.DtSlbTemplateClientSslStats.Stats.MalwareSites,
			"phishing_and_other_fraud":       ret.DtSlbTemplateClientSslStats.Stats.PhishingAndOtherFraud,
			"proxy_avoid_and_anonymizers":    ret.DtSlbTemplateClientSslStats.Stats.ProxyAvoidAndAnonymizers,
			"spyware_and_adware":             ret.DtSlbTemplateClientSslStats.Stats.SpywareAndAdware,
			"music":                          ret.DtSlbTemplateClientSslStats.Stats.Music,
			"government":                     ret.DtSlbTemplateClientSslStats.Stats.Government,
			"nudity":                         ret.DtSlbTemplateClientSslStats.Stats.Nudity,
			"news_and_media":                 ret.DtSlbTemplateClientSslStats.Stats.NewsAndMedia,
			"illegal":                        ret.DtSlbTemplateClientSslStats.Stats.Illegal,
			"cdns":                           ret.DtSlbTemplateClientSslStats.Stats.Cdns,
			"internet_communications":        ret.DtSlbTemplateClientSslStats.Stats.InternetCommunications,
			"bot_nets":                       ret.DtSlbTemplateClientSslStats.Stats.BotNets,
			"abortion":                       ret.DtSlbTemplateClientSslStats.Stats.Abortion,
			"health_and_medicine":            ret.DtSlbTemplateClientSslStats.Stats.HealthAndMedicine,
			"confirmed_spam_sources":         ret.DtSlbTemplateClientSslStats.Stats.ConfirmedSpamSources,
			"spam_urls":                      ret.DtSlbTemplateClientSslStats.Stats.SpamUrls,
			"unconfirmed_spam_sources":       ret.DtSlbTemplateClientSslStats.Stats.UnconfirmedSpamSources,
			"open_http_proxies":              ret.DtSlbTemplateClientSslStats.Stats.OpenHttpProxies,
			"dynamically_generated_content":  ret.DtSlbTemplateClientSslStats.Stats.DynamicallyGeneratedContent,
			"parked_domains":                 ret.DtSlbTemplateClientSslStats.Stats.ParkedDomains,
			"alcohol_and_tobacco":            ret.DtSlbTemplateClientSslStats.Stats.AlcoholAndTobacco,
			"private_ip_addresses":           ret.DtSlbTemplateClientSslStats.Stats.PrivateIpAddresses,
			"image_and_video_search":         ret.DtSlbTemplateClientSslStats.Stats.ImageAndVideoSearch,
			"fashion_and_beauty":             ret.DtSlbTemplateClientSslStats.Stats.FashionAndBeauty,
			"recreation_and_hobbies":         ret.DtSlbTemplateClientSslStats.Stats.RecreationAndHobbies,
			"motor_vehicles":                 ret.DtSlbTemplateClientSslStats.Stats.MotorVehicles,
			"web_hosting_sites":              ret.DtSlbTemplateClientSslStats.Stats.WebHostingSites,
			"food_and_dining":                ret.DtSlbTemplateClientSslStats.Stats.FoodAndDining,
			"nudity_artistic":                ret.DtSlbTemplateClientSslStats.Stats.NudityArtistic,
			"illegal_pornography":            ret.DtSlbTemplateClientSslStats.Stats.IllegalPornography,
			"uncategorised":                  ret.DtSlbTemplateClientSslStats.Stats.Uncategorised,
			"other_category":                 ret.DtSlbTemplateClientSslStats.Stats.OtherCategory,
			"trustworthy":                    ret.DtSlbTemplateClientSslStats.Stats.Trustworthy,
			"low_risk":                       ret.DtSlbTemplateClientSslStats.Stats.LowRisk,
			"moderate_risk":                  ret.DtSlbTemplateClientSslStats.Stats.ModerateRisk,
			"suspicious":                     ret.DtSlbTemplateClientSslStats.Stats.Suspicious,
			"malicious":                      ret.DtSlbTemplateClientSslStats.Stats.Malicious,
		},
	}
}

func getObjectSlbTemplateClientSslStatsStats(d []interface{}) edpt.SlbTemplateClientSslStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplateClientSslStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.Uncategorised = in["uncategorised"].(int)
		ret.OtherCategory = in["other_category"].(int)
		ret.Trustworthy = in["trustworthy"].(int)
		ret.LowRisk = in["low_risk"].(int)
		ret.ModerateRisk = in["moderate_risk"].(int)
		ret.Suspicious = in["suspicious"].(int)
		ret.Malicious = in["malicious"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateClientSslStats(d *schema.ResourceData) edpt.SlbTemplateClientSslStats {
	var ret edpt.SlbTemplateClientSslStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplateClientSslStatsStats(d.Get("stats").([]interface{}))
	return ret
}
