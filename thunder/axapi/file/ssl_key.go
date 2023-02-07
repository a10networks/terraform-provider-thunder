package file

//pseudo endpoint for upload and delete SSL-key files

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	endpnt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type SslKey struct {
	Name        string
	Protocol    string
	Host        string
	Path        string
	Username    string
	Password    string
	UseMgmtPort int
	Overwrite   int
	Secured     int
}

func (p *SslKey) GetId() string {
	return p.Name
}

func (p *SslKey) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslKey::Post")

	realInst := endpnt.Import{}
	realInst.Inst.SslKey = p.Name
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	realInst.Inst.Secured = p.Secured
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *SslKey) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslKey::Get")
	err := oper.FileSslKeyGet(authToken, host, instId, logger)
	return err
}

func (p *SslKey) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslKey::Delet")
	obj := endpnt.PkiDelete{}
	obj.Inst.PrivateKey = instId
	return obj.Post(authToken, host, logger)
}
