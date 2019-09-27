CREATE TABLE `btc` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `BlockChain_height` int(100) NOT NULL,
  `BlockChair_height` int(100) NOT NULL,
  `ViaBtc_height` int(100) NOT NULL,
  `BTCcom_height` int(100) NOT NULL,
  `Node_height` int(100) NOT NULL,
  `BlockChain_hash` varchar(100) NOT NULL,
  `BlockChair_hash` varchar(100) NOT NULL,
  `ViaBtc_hash` varchar(100) NOT NULL,
  `BTCcom_hash` varchar(100) NOT NULL,
  `Node_hash` varchar(100) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE `bch` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `BlockChair_height` int(100) NOT NULL,
  `Bitcoin_height` int(100) NOT NULL,
  `ViaBtc_height` int(100) NOT NULL,
  `BTCcom_height` int(100) NOT NULL,
  `Node_height` int(100) NOT NULL,
  `BlockChair_hash` varchar(100) NOT NULL,
  `Bitcoin_hash` varchar(100) NOT NULL,
  `ViaBtc_hash` varchar(100) NOT NULL,
  `BTCcom_hash` varchar(100) NOT NULL,
  `Node_hash` varchar(100) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE `ltc` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `BlockChair_height` int(100) NOT NULL,
  `ViaBtc_height` int(100) NOT NULL,
  `BlockCypher_height` int(100) NOT NULL,
  `BTCcom_height` int(100) NOT NULL,
  `Node_height` int(100) NOT NULL,
  `BlockChair_hash` varchar(100) NOT NULL,
  `ViaBtc_hash` varchar(100) NOT NULL,
  `BlockCypher_hash` varchar(100) NOT NULL,
  `BTCcom_hash` varchar(100) NOT NULL,
  `Node_hash` varchar(100) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE `eth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Etherscan_height` int(100) NOT NULL,
  `BlockChain_height` int(100) NOT NULL,
  `BlockChair_height` int(100) NOT NULL,
  `BTCcom_height` int(100) NOT NULL,
  `Node_height` int(100) NOT NULL,
  `Etherscan_hash` varchar(100) NOT NULL,
  `BlockChain_hash` varchar(100) NOT NULL,
  `BlockChair_hash` varchar(100) NOT NULL,
  `BTCcom_hash` varchar(100) NOT NULL,
  `Node_hash` varchar(100) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

