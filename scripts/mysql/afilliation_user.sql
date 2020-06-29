CREATE TABLE `afilliation_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user` varchar(70) DEFAULT NULL,
  `user_number` int(10) DEFAULT NULL,
  `afilliation_id` int(11) DEFAULT NULL,
  `aggregator` int(11) DEFAULT NULL,
  `terminal` int(2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8