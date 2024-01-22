package axapi

import (
	"log"
	"os"
)

type ThunderLog struct {
	filename string
	*log.Logger
}

func CreateLogger(partition_name string) *ThunderLog {
	fname := "Thunder_" + partition_name + ".txt"
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	return &ThunderLog{
		filename: fname,
		Logger:   log.New(file, "["+partition_name+"] ", log.Lshortfile),
	}
}
