CREATE TABLE `status` (
  `id` bigint(24) NOT NULL AUTO_INCREMENT,
  `type` varchar(30) NOT NULL,
  `code` varchar(30) NOT NULL,
  `name` varchar(30) NOT NULL,
  `description` varchar(100) DEFAULT NULL,
  PRIMARY KEY  (`id`)
)