package vthunder

//vThunder resource http policy

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"util"
	"fmt"
)

func resourceSlbTemplateHttpPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceHttpPolicyCreate,
		Update: resourceHttpPolicyUpdate,
		Read:   resourceHttpPolicyRead,
		Delete: resourceHttpPolicyDelete,

		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"geo_location_match": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_template_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"geo_location_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"http_policy_match": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_under_host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"other_match_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"other_match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"match_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cookie_name": {
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
	}
}

func resourceHttpPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating http policy (Inside resourceHttpPolicyCreate    " + name)
		v := dataToHttpPolicy(name, d)
		d.SetId(name)
		go_vthunder.PostHttpPolicy(client.Token, v, client.Host)

		return resourceHttpPolicyRead(d, meta)
	}
	return nil
}

func resourceHttpPolicyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading http policy (Inside resourceHttpPolicyRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching htpp policy Read" + name)

		http_policy, err := go_vthunder.GetHttpPolicy(client.Token, name, client.Host)

		if http_policy == nil {
			logger.Println("[INFO] No http policy found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceHttpPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying http policy (Inside resourceHttpPolicyUpdate    " + name)
		v := dataToHttpPolicy(name, d)
		d.SetId(name)
		go_vthunder.PutHttpPolicy(client.Token, name, v, client.Host)

		return resourceHttpPolicyRead(d, meta)
	}
	return nil
}

func resourceHttpPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting http policy (Inside resourceHttpPolicyDelete) " + name)

		err := go_vthunder.DeleteHttpPolicy(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete http policy  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate http policy structure
func dataToHttpPolicy(name string, d *schema.ResourceData) go_vthunder.HttpPolicy {
	var s go_vthunder.HttpPolicy

	var sInstance go_vthunder.HttpPolicyInstance

	sInstance.CookieName = d.Get("cookie_name").(string)
	sInstance.Name = d.Get("name").(string)
	sInstance.UUID = d.Get("uuid").(string)
	sInstance.UserTag = d.Get("user_tag").(string)

	match_cnt := d.Get("http_policy_match.#").(int)
	sInstance.MatchString = make([]go_vthunder.HTTPPolicyMatch, 0, match_cnt)
	for i := 0; i < match_cnt; i++ {

		var match go_vthunder.HTTPPolicyMatch
		prefix := fmt.Sprintf("http_policy_match.%d", i)
		match.Type = d.Get(prefix + ".type").(string)
		match.MatchString = d.Get(prefix + ".match_string").(string)
		match.TemplateName = d.Get(prefix + ".template_name").(string)
		match.ServiceGroup = d.Get(prefix + ".service_group").(string)
		match.URLUnderHost = d.Get(prefix + ".url_under_host").(string)
		match.OtherMatchString = d.Get(prefix + ".other_match_string").(string)
		match.Template = d.Get(prefix + ".template").(string)
		match.OtherMatchType = d.Get(prefix + ".other_match_type").(string)
		match.MatchType = d.Get(prefix + ".match_type").(string)

		sInstance.MatchString = append(sInstance.MatchString, match)
	}

	geo_match_cnt := d.Get("geo_location_match.#").(int)
	sInstance.GeoLocation = make([]go_vthunder.GeoLocationMatch, 0, geo_match_cnt)
	for i := 0; i < geo_match_cnt; i++ {

		var match go_vthunder.GeoLocationMatch
		prefix := fmt.Sprintf("geo_location_match.%d", i)
		match.GeoLocation = d.Get(prefix + ".geo_location").(string)
		match.GeoLocationServiceGroup = d.Get(prefix + ".geo_location_service_group").(string)
		match.GeoLocationTemplate = d.Get(prefix + ".geo_location_template").(string)
		match.GeoLocationTemplateName = d.Get(prefix + ".geo_location_template_name").(string)

		sInstance.GeoLocation = append(sInstance.GeoLocation, match)
	}

	s.Name = sInstance

	return s
}
