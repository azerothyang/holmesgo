package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/jinzhu/gorm"
	"holmes/conf"
	"log"
)

var (
	db     *gorm.DB //master db
	readDB *gorm.DB //read db
)

type baseModel struct {
}

//init db
func InitDB() {
	var errM, errS error
	var dsnMaster, dsnSlave string
	var masterMaxIdle, masterMaxConn, slaveMaxIdle, slaveMaxConn int

	//mysql master
	dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local", //loc set the timezone
		conf.Conf.Database.Mysql.Master.User, conf.Conf.Database.Mysql.Master.Password, conf.Conf.Database.Mysql.Master.Host, conf.Conf.Database.Mysql.Master.Port, conf.Conf.Database.Mysql.Master.Database, conf.Conf.Database.Mysql.Master.Charset, conf.Conf.Database.Mysql.Master.Timeout)

	//mysql slave
	dsnSlave = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local",
		conf.Conf.Database.Mysql.Slave.User, conf.Conf.Database.Mysql.Slave.Password, conf.Conf.Database.Mysql.Slave.Host, conf.Conf.Database.Mysql.Slave.Port, conf.Conf.Database.Mysql.Slave.Database, conf.Conf.Database.Mysql.Slave.Charset, conf.Conf.Database.Mysql.Slave.Timeout)
	masterMaxIdle = conf.Conf.Database.Mysql.Master.MaxIdle
	masterMaxConn = conf.Conf.Database.Mysql.Master.MaxConn
	slaveMaxIdle = conf.Conf.Database.Mysql.Slave.MaxIdle
	slaveMaxConn = conf.Conf.Database.Mysql.Slave.MaxConn

	//gorm realizes mysql reconnect
	db, errM = gorm.Open("mysql", dsnMaster)
	readDB, errS = gorm.Open("mysql", dsnSlave)
	if errM != nil || errS != nil {
		log.Fatalln(errM, errS)
	}

	db.DB().SetMaxIdleConns(masterMaxIdle)
	db.DB().SetMaxOpenConns(masterMaxConn)
	//db.LogMode(true) enable log to debug, log can write into elasticsearch
	readDB.DB().SetMaxIdleConns(slaveMaxIdle)
	readDB.DB().SetMaxOpenConns(slaveMaxConn)
	//readDB.LogMode(true)
}
