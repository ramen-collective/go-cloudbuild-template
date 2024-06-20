-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: 172.17.0.2    Database: template-api
-- ------------------------------------------------------
-- Server version	8.0.27

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (uuid, name, created_at, updated_at) VALUES ('de2cae14-5c34-40a4-bd3c-37a8a8dbaebb', 'Jean', NOW(), NOW());
INSERT INTO `user` (uuid, name, created_at, updated_at) VALUES ('21054bbd-4930-437b-8628-f2dba04e56bc', 'Marc', NOW(), NOW());
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;