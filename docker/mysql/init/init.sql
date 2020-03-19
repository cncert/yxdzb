CREATE DATABASE
IF NOT EXISTS yxdzb DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
CREATE USER 'yxdzb'@'%' IDENTIFIED BY '12345678';
GRANT All privileges ON yxdzb.* TO 'yxdzb'@'%';
