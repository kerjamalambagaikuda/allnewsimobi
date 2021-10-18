CREATE TABLE `lookup` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type` varchar(30) NOT NULL,
  `code` varchar(30) NOT NULL,
  `name` varchar(40) NOT NULL,
  `priority` int(11) default NULL,
  `description` varchar(100) default NULL,
  `shortname` varchar(10) default NULL,
  `status` bigint(20) NOT NULL,
  `filter` varchar(30) default NULL,
  PRIMARY KEY  (`id`),
  UNIQUE KEY `Unique_type_code` (`type`,`code`),
  FOREIGN KEY FKLOOKUPSTATUS(`status`) REFERENCES status(`id`)
)