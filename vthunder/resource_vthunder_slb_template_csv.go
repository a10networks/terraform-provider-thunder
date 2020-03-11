package vthunder

//vThunder resource SlbTemplateCSV

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateCSV() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateCSVCreate,
		Update: resourceSlbTemplateCSVUpdate,
		Read:   resourceSlbTemplateCSVRead,
		Delete: resourceSlbTemplateCSVDelete,
		Schema: map[string]*schema.Schema{
			"delim_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"multiple_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"csv_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"csv_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"delim_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateCSVCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateCSV (Inside resourceSlbTemplateCSVCreate) ")
		name := d.Get("csv_name").(string)
		data := dataToSlbTemplateCSV(d)
		logger.Println("[INFO] received V from method data to SlbTemplateCSV --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateCSV(client.Token, data, client.Host)

		return resourceSlbTemplateCSVRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateCSVRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateCSV (Inside resourceSlbTemplateCSVRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateCSV(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateCSVUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateCSV   (Inside resourceSlbTemplateCSVUpdate) ")
		name := d.Get("csv_name").(string)
		data := dataToSlbTemplateCSV(d)
		logger.Println("[INFO] received V from method data to SlbTemplateCSV ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateCSV(client.Token, name, data, client.Host)

		return resourceSlbTemplateCSVRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateCSVDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateCSVDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateCSV(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateCSV(d *schema.ResourceData) go_vthunder.CSV {
	var vc go_vthunder.CSV
	var c go_vthunder.CSVInstance

	c.CsvName = d.Get("csv_name").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Ipv6Enable = d.Get("ipv6_enable").(int)
	c.DelimNum = d.Get("delim_num").(int)
	c.DelimChar = d.Get("delim_char").(string)

	MultipleFieldsCount := d.Get("multiple_fields.#").(int)
	c.Field = make([]go_vthunder.MultipleFields, 0, MultipleFieldsCount)

	for i := 0; i < MultipleFieldsCount; i++ {
		var obj1 go_vthunder.MultipleFields
		prefix := fmt.Sprintf("multiple_fields.%d.", i)
		obj1.Field = d.Get(prefix + "field").(int)
		obj1.CsvType = d.Get(prefix + "csv_type").(string)
		c.Field = append(c.Field, obj1)
	}

	vc.UUID = c
	return vc
}
