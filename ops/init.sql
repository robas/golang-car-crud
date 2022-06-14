create database if not exists car_project;
use car_project;
create table if not exists car(id varchar(50), brand varchar(255), model varchar(255), doorquantity integer, created_at datetime, updated_at datetime, PRIMARY KEY(`id`) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;)
