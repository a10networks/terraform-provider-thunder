package file

//pseudo endpoint for upload and delete aFlex files

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	endpnt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type Aflex struct {
	Name        string
	Protocol    string
	Host        string
	Path        string
	Username    string
	Password    string
	UseMgmtPort int
	Overwrite   int
}

func (p *Aflex) GetId() string {
	return p.Name
}

func (p *Aflex) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.Aflex::Post")

	realInst := endpnt.Import{}
	realInst.Inst.Aflex = p.Name
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *Aflex) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.Aflex::Get")
	err := oper.FileAflexGet(authToken, host, instId, logger)
	return err
}

func (p *Aflex) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.Aflex::Delet")
	obj := endpnt.FileAflex{}
	obj.Inst.Action = "delete"
	obj.Inst.File = instId
	return obj.Post(authToken, host, logger)
}
