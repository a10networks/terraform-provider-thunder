provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_routing_isis" "thunder_snmp_server_enable_traps_routing_isis" {
  isisadjacencychange                  = 0
  isisareamismatch                     = 0
  isisattempttoexceedmaxsequence       = 0
  isisauthenticationfailure            = 0
  isisauthenticationtypefailure        = 0
  isiscorruptedlspdetected             = 0
  isisdatabaseoverload                 = 0
  isisidlenmismatch                    = 0
  isislsperrordetected                 = 0
  isislsptoolargetopropagate           = 0
  isismanualaddressdrops               = 0
  isismaxareaaddressesmismatch         = 0
  isisoriginatinglspbuffersizemismatch = 0
  isisownlsppurge                      = 0
  isisprotocolssupportedmismatch       = 0
  isisrejectedadjacency                = 0
  isissequencenumberskip               = 0
  isisversionskew                      = 0
}