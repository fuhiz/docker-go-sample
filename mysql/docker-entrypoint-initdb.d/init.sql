CREATE DATABASE IF NOT EXISTS `go_sample` COLLATE 'utf8mb4_general_ci' ;
CREATE DATABASE IF NOT EXISTS `go_sample_test` COLLATE 'utf8mb4_general_ci' ;

GRANT ALL ON `go_sample`.* TO 'localuser'@'%' ;
GRANT ALL ON `go_sample_test`.* TO 'localuser'@'%' ;

FLUSH PRIVILEGES ;
