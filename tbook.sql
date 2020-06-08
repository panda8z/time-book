

DROP TABLE if exist  `tb_note`;

CREATE TABLE `tb_note` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(100) DEFAULT '' COMMENT '类型',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by`  int(10) DEFAULT 0 COMMENT '创建人id',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Note实体';

DROP TABLE if exist  `tb_auth`;

CREATE TABLE `tb_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='认证实体';

INSERT INTO `tbook`.`tb_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');