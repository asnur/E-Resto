/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE `tbl_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `icon` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

CREATE TABLE `tbl_meja` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `no_meja` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

CREATE TABLE `tbl_menu` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `price` bigint(20) DEFAULT NULL,
  `id_category` bigint(20) DEFAULT NULL,
  `foto` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_tbl_menu_category` (`id_category`),
  CONSTRAINT `fk_tbl_menu_category` FOREIGN KEY (`id_category`) REFERENCES `tbl_category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

CREATE TABLE `tbl_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `no_meja` bigint(20) DEFAULT NULL,
  `total_price` bigint(20) DEFAULT NULL,
  `id_payement` bigint(20) DEFAULT NULL,
  `status` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_tbl_order_table` (`no_meja`),
  KEY `fk_tbl_order_payement` (`id_payement`),
  CONSTRAINT `fk_tbl_order_payement` FOREIGN KEY (`id_payement`) REFERENCES `tbl_payement` (`id`),
  CONSTRAINT `fk_tbl_order_table` FOREIGN KEY (`no_meja`) REFERENCES `tbl_meja` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `tbl_order_detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_order` bigint(20) DEFAULT NULL,
  `id_menu` bigint(20) DEFAULT NULL,
  `quantity` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_tbl_order_detail_menu` (`id_menu`),
  KEY `fk_tbl_order_order_details` (`id_order`),
  CONSTRAINT `fk_tbl_order_detail_menu` FOREIGN KEY (`id_menu`) REFERENCES `tbl_menu` (`id`),
  CONSTRAINT `fk_tbl_order_order_details` FOREIGN KEY (`id_order`) REFERENCES `tbl_order` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `tbl_payement` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

INSERT INTO `tbl_category` (`id`, `name`, `icon`) VALUES
(1, 'Makanan', 'bibimbap.png');
INSERT INTO `tbl_category` (`id`, `name`, `icon`) VALUES
(2, 'Minuman', 'drink.png');


INSERT INTO `tbl_meja` (`id`, `no_meja`) VALUES
(1, 1);
INSERT INTO `tbl_meja` (`id`, `no_meja`) VALUES
(2, 2);
INSERT INTO `tbl_meja` (`id`, `no_meja`) VALUES
(3, 3);
INSERT INTO `tbl_meja` (`id`, `no_meja`) VALUES
(4, 4),
(5, 5),
(6, 6);

INSERT INTO `tbl_menu` (`id`, `name`, `price`, `id_category`, `foto`) VALUES
(1, 'Es Teh', 2000, 2, 'es_teh.jpg');
INSERT INTO `tbl_menu` (`id`, `name`, `price`, `id_category`, `foto`) VALUES
(2, 'Mie', 10000, 1, 'mie.jpg');
INSERT INTO `tbl_menu` (`id`, `name`, `price`, `id_category`, `foto`) VALUES
(3, 'Bakso', 15000, 1, 'bakso.jpg');
INSERT INTO `tbl_menu` (`id`, `name`, `price`, `id_category`, `foto`) VALUES
(4, 'Es Coklat', 3000, 2, 'es_coklat.jpeg'),
(5, 'Kopi', 2000, 2, 'kopi.jpg'),
(6, 'Rujak Kangkung', 20000, 1, 'rujak_kangkung.jpg');





INSERT INTO `tbl_payement` (`id`, `name`) VALUES
(1, 'Gopay');
INSERT INTO `tbl_payement` (`id`, `name`) VALUES
(2, 'BCA');
INSERT INTO `tbl_payement` (`id`, `name`) VALUES
(3, 'BNI');


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;