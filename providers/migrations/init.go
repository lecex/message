package migrations

import (
	cpd "github.com/lecex/message/proto/config"
	tpd "github.com/lecex/message/proto/template"
	db "github.com/lecex/message/providers/database"
)

func init() {
	config()
	template()
	seeds()
}

// config 配置数据迁移
func config() {
	config := &cpd.Config{}
	if !db.DB.HasTable(&config) {
		db.DB.Exec(`
			CREATE TABLE configs (
			name varchar(32) NOT NULL COMMENT '配置名称',
			value json DEFAULT NULL COMMENT '配置内容',
			PRIMARY KEY (name)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// template 模板数据迁移
func template() {
	template := &tpd.Template{}
	if !db.DB.HasTable(&template) {
		db.DB.Exec(`
			CREATE TABLE templates (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,	
			event varchar(32) NOT NULL,
			name varchar(64) DEFAULT NULL,
			type varchar(64) DEFAULT NULL,
			template_code varchar(128) DEFAULT NULL,
			template_value varchar(128) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY event (event)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// seeds 填充文件
func seeds() {
	db.DB.Exec(`
		INSERT INTO templates ( event, name, type, template_code, template_value ) VALUES ('register_verify','用户注册验证码','sms','453946','')
	`)
	db.DB.Exec(`
		INSERT INTO configs ( name ) VALUES ('config')
	`)
}
