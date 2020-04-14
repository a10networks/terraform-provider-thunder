package vthunder

//vThunder resource SlbTemplateMonitor

import (
	"fmt"
	"log"
	"strconv"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateMonitorCreate,
		Update: resourceSlbTemplateMonitorUpdate,
		Read:   resourceSlbTemplateMonitorRead,
		Delete: resourceSlbTemplateMonitorDelete,
		Schema: map[string]*schema.Schema{
			"link_enable_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ena_sequence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"enaeth": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"link_disable_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dis_sequence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"diseth": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"clear_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sessions": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"clear_all_sequence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"clear_sequence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"link_up_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"linkup_ethernet2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"linkup_ethernet1": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"link_up_sequence2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"link_up_sequence3": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"linkup_ethernet3": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"link_up_sequence1": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"id2": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"link_down_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_down_sequence1": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"linkdown_ethernet2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"link_down_sequence2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"linkdown_ethernet3": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"link_down_sequence3": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"linkdown_ethernet1": {
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
			"monitor_relation": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateMonitor (Inside resourceSlbTemplateMonitorCreate) ")
		name := strconv.Itoa(d.Get("id2").(int))
		data := dataToSlbTemplateMonitor(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateMonitor --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateMonitor(client.Token, data, client.Host)

		return resourceSlbTemplateMonitorRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateMonitorRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateMonitor (Inside resourceSlbTemplateMonitorRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateMonitor(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateMonitor   (Inside resourceSlbTemplateMonitorUpdate) ")
		name := strconv.Itoa(d.Get("id2").(int))
		data := dataToSlbTemplateMonitor(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateMonitor ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateMonitor(client.Token, name, data, client.Host)

		return resourceSlbTemplateMonitorRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateMonitorDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateMonitor(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateMonitor(d *schema.ResourceData) go_vthunder.Monitor {
	var vc go_vthunder.Monitor
	var c go_vthunder.MonitorInstance

	ClearCfgCount := d.Get("clear_cfg.#").(int)
	c.ClearSequence = make([]go_vthunder.ClearCfg, 0, ClearCfgCount)

	for i := 0; i < ClearCfgCount; i++ {
		var obj1 go_vthunder.ClearCfg
		prefix := fmt.Sprintf("clear_cfg.%d.", i)
		obj1.ClearSequence = d.Get(prefix + "clear_sequence").(int)
		obj1.ClearAllSequence = d.Get(prefix + "clear_all_sequence").(int)
		obj1.Sessions = d.Get(prefix + "sessions").(string)
		c.ClearSequence = append(c.ClearSequence, obj1)
	}

	LinkEnableCfgCount := d.Get("link_enable_cfg.#").(int)
	c.EnaSequence = make([]go_vthunder.LinkEnableCfg, 0, LinkEnableCfgCount)

	for i := 0; i < LinkEnableCfgCount; i++ {
		var obj2 go_vthunder.LinkEnableCfg
		prefix := fmt.Sprintf("link_enable_cfg.%d.", i)
		obj2.EnaSequence = d.Get(prefix + "ena_sequence").(int)
		obj2.Enaeth = d.Get(prefix + "enaeth").(int)
		c.EnaSequence = append(c.EnaSequence, obj2)
	}

	LinkUpCfgCount := d.Get("link_up_cfg.#").(int)
	c.LinkupEthernet3 = make([]go_vthunder.LinkUpCfg, 0, LinkUpCfgCount)

	for i := 0; i < LinkUpCfgCount; i++ {
		var obj3 go_vthunder.LinkUpCfg
		prefix := fmt.Sprintf("link_up_cfg.%d.", i)
		obj3.LinkupEthernet3 = d.Get(prefix + "linkup_ethernet3").(int)
		obj3.LinkupEthernet2 = d.Get(prefix + "linkup_ethernet2").(int)
		obj3.LinkupEthernet1 = d.Get(prefix + "linkup_ethernet1").(int)
		obj3.LinkUpSequence1 = d.Get(prefix + "link_up_sequence1").(int)
		obj3.LinkUpSequence3 = d.Get(prefix + "link_up_sequence3").(int)
		obj3.LinkUpSequence2 = d.Get(prefix + "link_up_sequence2").(int)
		c.LinkupEthernet3 = append(c.LinkupEthernet3, obj3)
	}

	LinkDownCfgCount := d.Get("link_down_cfg.#").(int)
	c.LinkDownSequence1 = make([]go_vthunder.LinkDownCfg, 0, LinkDownCfgCount)

	for i := 0; i < LinkDownCfgCount; i++ {
		var obj4 go_vthunder.LinkDownCfg
		prefix := fmt.Sprintf("link_down_cfg.%d.", i)
		obj4.LinkDownSequence1 = d.Get(prefix + "link_down_sequence1").(int)
		obj4.LinkDownSequence2 = d.Get(prefix + "link_down_sequence2").(int)
		obj4.LinkDownSequence3 = d.Get(prefix + "link_down_sequence3").(int)
		obj4.LinkdownEthernet2 = d.Get(prefix + "linkdown_ethernet2").(int)
		obj4.LinkdownEthernet3 = d.Get(prefix + "linkdown_ethernet3").(int)
		obj4.LinkdownEthernet1 = d.Get(prefix + "linkdown_ethernet1").(int)
		c.LinkDownSequence1 = append(c.LinkDownSequence1, obj4)
	}

	c.UserTag = d.Get("user_tag").(string)

	LinkDisableCfgCount := d.Get("link_disable_cfg.#").(int)
	c.DisSequence = make([]go_vthunder.LinkDisableCfg, 0, LinkDisableCfgCount)

	for i := 0; i < LinkDisableCfgCount; i++ {
		var obj5 go_vthunder.LinkDisableCfg
		prefix := fmt.Sprintf("link_disable_cfg.%d.", i)
		obj5.DisSequence = d.Get(prefix + "dis_sequence").(int)
		obj5.Diseth = d.Get(prefix + "diseth").(int)
		c.DisSequence = append(c.DisSequence, obj5)
	}

	c.MonitorRelation = d.Get("monitor_relation").(string)
	c.ID = d.Get("id2").(int)

	vc.UUID = c
	return vc
}
