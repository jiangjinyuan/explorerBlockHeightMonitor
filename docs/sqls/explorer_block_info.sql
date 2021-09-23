CREATE TABLE `explorer_block_info` (
  `coin` varchar(30) NOT NULL,
  `explorer_name` varchar(100) NOT NULL,
  `height` int(11) NOT NULL DEFAULT 0,
  `hash` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`coin`,`explorer_name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
