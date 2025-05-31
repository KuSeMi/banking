DROP DATABASE IF EXISTS banking;
CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
    `customer_id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `date_of_birth` date NOT NULL,
    `city` varchar(100) NOT NULL,
    `zipcode` varchar(10) NOT NULL,
    `status` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `customers` VALUES
(2000, 'Steve', '1978-12-15', 'Delhi', '110075', 1),
(2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550', 1),
(2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', 1),
(2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', 0),
(2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', 1),
(2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', 0);

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
    `account_id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL,
    `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account_type` varchar(10) NOT NULL,
    `amount` DECIMAL(10,2) NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    PRIMARY KEY (`account_id`),
    CONSTRAINT `account_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95476 DEFAULT CHARSET=latin1;

INSERT INTO `accounts` VALUES
(95470, 2000, '2020-08-22 10:27:22', 'Saving', 5000.75, 1),
(95471, 2001, '2020-06-15 10:27:22', 'Saving', 30200.50, 1),
(95472, 2002, '2020-08-09 10:27:22', 'Checking', 15000.25, 1),
(95473, 2000, '2020-06-03 10:27:22', 'Saving', 7800.00, 1),
(95474, 2004, '2020-02-27 10:27:22', 'Checking', 42000.30, 1),
(95475, 2005, '2020-03-20 10:27:22', 'Saving', 10000.00, 0);

DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
    `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL,
    `amount` DECIMAL(10,2) NOT NULL,
    `transaction_type` varchar(10) NOT NULL,
    `date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`transaction_id`),
    CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;