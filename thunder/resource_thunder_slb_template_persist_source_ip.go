package thunder

//Thunder resource SlbTemplatePersistSourceIp

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSlbTemplatePersistSourceIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplatePersistSourceIpCreate,
		Update: resourceSlbTemplatePersistSourceIpUpdate,
		Read:   resourceSlbTemplatePersistSourceIpRead,
		Delete: resourceSlbTemplatePersistSourceIpDelete,
		Schema: map[string]*schema.Schema{
			"netmask6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"incl_dst_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"hash_persist": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"enforce_higher_priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dont_honor_conn_rules": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"primary_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"scan_all_members": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"netmask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"incl_sport": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"match_type": {
				Type:        schema.TypeInt,
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

func resourceSlbTemplatePersistSourceIpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplatePersistSourceIp (Inside resourceSlbTemplatePersistSourceIpCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplatePersistSourceIp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplatePersistSourceIp --")
		d.SetId(name1)
		go_thunder.PostSlbTemplatePersistSourceIp(client.Token, data, client.Host)

		return resourceSlbTemplatePersistSourceIpRead(d, meta)

	}
	return nil
}

func resourceSlbTemplatePersistSourceIpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplatePersistSourceIp (Inside resourceSlbTemplatePersistSourceIpRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read SlbTemplatePersistSourceIp ")
		data, err := go_thunder.GetSlbTemplatePersistSourceIp(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found SlbTemplatePersistSourceIp ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplatePersistSourceIpUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplatePersistSourceIp   (Inside resourceSlbTemplatePersistSourceIpUpdate) ")
		data := dataToSlbTemplatePersistSourceIp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplatePersistSourceIp ")
		go_thunder.PutSlbTemplatePersistSourceIp(client.Token, name1, data, client.Host)

		return resourceSlbTemplatePersistSourceIpRead(d, meta)

	}
	return nil
}

func resourceSlbTemplatePersistSourceIpDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplatePersistSourceIpDelete) ")
		err := go_thunder.DeleteSlbTemplatePersistSourceIp(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSlbTemplatePersistSourceIp(d *schema.ResourceData) go_thunder.SlbTemplatePersistSourceIp {
	var vc go_thunder.SlbTemplatePersistSourceIp
	var c go_thunder.SlbTemplatePersistSourceIPInstance
	c.Netmask6 = d.Get("netmask6").(int)
	c.InclDstIP = d.Get("incl_dst_ip").(int)
	c.HashPersist = d.Get("hash_persist").(int)
	c.Name = d.Get("name").(string)
	c.EnforceHigherPriority = d.Get("enforce_higher_priority").(int)
	c.DontHonorConnRules = d.Get("dont_honor_conn_rules").(int)
	c.PrimaryPort = d.Get("primary_port").(int)
	c.UserTag = d.Get("user_tag").(string)
	c.Server = d.Get("server").(int)
	c.ServiceGroup = d.Get("service_group").(int)
	c.Timeout = d.Get("timeout").(int)
	c.ScanAllMembers = d.Get("scan_all_members").(int)
	c.Netmask = d.Get("netmask").(string)
	c.InclSport = d.Get("incl_sport").(int)
	c.MatchType = d.Get("match_type").(int)

	vc.Netmask6 = c
	return vc
}
