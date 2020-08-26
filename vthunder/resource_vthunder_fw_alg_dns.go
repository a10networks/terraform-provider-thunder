package vthunder

//vThunder resource FwAlgDns

import (
    "github.com/go_vthunder/vthunder"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "util"
    "fmt"
    "strconv"
)

func resourceFwAlgDns() *schema.Resource {
    return &schema.Resource{
        Create: resourceFwAlgDnsCreate,
        Update: resourceFwAlgDnsUpdate,
        Read:   resourceFwAlgDnsRead,
        Delete: resourceFwAlgDnsDelete,
         Schema: map[string]*schema.Schema{ 
"default_port_disable":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
},
    } }
 
func resourceFwAlgDnsCreate(d *schema.ResourceData, meta interface{}) error {
     logger := util.GetLoggerInstance()
     client := meta.(vThunder)
      
    if client.Host != "" {
         logger.Println("[INFO] Creating FwAlgDns (Inside resourceFwAlgDnsCreate) ")
          
        data := dataToFwAlgDns(d)
         logger.Println("[INFO] received formatted data from method data to FwAlgDns --")
         d.SetId(strconv.Itoa('1'))
         go_vthunder.PostFwAlgDns(client.Token, data, client.Host)
 
        return resourceFwAlgDnsRead(d, meta)

    }
     return nil
 }
    
func resourceFwAlgDnsRead(d *schema.ResourceData, meta interface{}) error {
     logger := util.GetLoggerInstance()
     client := meta.(vThunder)
     logger.Println("[INFO] Reading FwAlgDns (Inside resourceFwAlgDnsRead)")
 
    if client.Host != "" {
         name := d.Id()
        logger.Println("[INFO] Fetching service Read" + name)
         data, err := go_vthunder.GetFwAlgDns(client.Token, client.Host)
         if data == nil {
             logger.Println("[INFO] No data found " + name)
             d.SetId("")
             return nil
         }
        return err
    }
     return nil
 }
    
func resourceFwAlgDnsUpdate(d *schema.ResourceData, meta interface{}) error {

    return resourceFwAlgDnsRead(d, meta)
}

func resourceFwAlgDnsDelete(d *schema.ResourceData, meta interface{}) error {

    return resourceFwAlgDnsRead(d, meta)
}
func dataToFwAlgDns(d *schema.ResourceData) go_vthunder.FwAlg {
    var vc go_vthunder.FwAlg
    var c go_vthunder.FwAlgDNSInstance
    c.DefaultPortDisable = d.Get("default_port_disable").(string)

    vc.DefaultPortDisable = c 
    return vc
}