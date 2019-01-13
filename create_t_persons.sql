CREATE TABLE `t_persons` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL COMMENT '名前',
  `adress` varchar(45) DEFAULT NULL COMMENT '住所',
  `str2` text COMMENT 'str2',
  `age` int(11) NOT NULL DEFAULT '0',
  `int1` int(11) NOT NULL DEFAULT '0',
  `int2` int(11) NOT NULL DEFAULT '0',
  `int3` int(11) NOT NULL DEFAULT '0',
  `int4` int(11) NOT NULL DEFAULT '0',
  `update_time` datetime DEFAULT NULL,
  `del` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

