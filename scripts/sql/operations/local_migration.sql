-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: 0.0.0.0.0    Database: api-template
-- ------------------------------------------------------
-- Server version 8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `language`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `language` (
    `id` BIGINT UNSIGNED UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` varchar(64) NOT NULL,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;

CREATE TABLE `user` (
    `uuid` varchar(36) UNIQUE PRIMARY KEY NOT NULL,
    `name` varchar(64) NOT NULL,
    `email` varchar(64) UNIQUE NOT NULL,
    `password` varchar(64) NOT NULL,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `content`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;

CREATE TABLE `content` (
    `uuid` varchar(36) UNIQUE PRIMARY KEY NOT NULL,
    `name` varchar(255) NOT NULL,
    `description` TEXT,
    `state` enum('draft', 'validated', 'published') NOT NULL DEFAULT 'draft',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `collection`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `collection` (
  `uuid` varchar(36) UNIQUE PRIMARY KEY NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `state` enum('draft', 'validated', 'published') NOT NULL DEFAULT 'draft',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `collection_content`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `collection_content` (
  `uuid` varchar(36) UNIQUE PRIMARY KEY NOT NULL,
  `collection_uuid` varchar(64) NOT NULL,
  `content_uuid` varchar(64) NOT NULL,
  `position` int NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `uq_collection_uuid_content_uuid_idx` (`collection_uuid`,`content_uuid`),
  KEY `fk_collection_content_content_uuid_idx` (`content_uuid`),
  KEY `collection_uuid_position_idx` (`collection_uuid`,`position`),
  CONSTRAINT `fk_collection_content_collection_uuid` FOREIGN KEY (`collection_uuid`) REFERENCES `collection` (`uuid`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_collection_content_content_uuid` FOREIGN KEY (`content_uuid`) REFERENCES `content` (`uuid`) ON DELETE RESTRICT ON UPDATE RESTRICT
);
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `tag`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;

CREATE TABLE `tag` (
    `uuid` varchar(36) UNIQUE PRIMARY KEY NOT NULL,
    `slug` varchar(64) NOT NULL,
    `name` varchar(255) NOT NULL,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `item_tag`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `item_tag` (
  `uuid` varchar(36) UNIQUE PRIMARY KEY NOT NULL,
  `item_uuid` varchar(36) NOT NULL,
  `item_type` enum('tag','content','collection') NOT NULL,
  `tag_uuid` varchar(36) NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `uq_item_uuid_item_type_tag_uuid_idx` (`item_uuid`,`item_type`,`tag_uuid`),
  KEY `fk_item_tag_tag_iuud_idx` (`tag_uuid`),
  CONSTRAINT `fk_item_tag_tag_uuid` FOREIGN KEY (`tag_uuid`) REFERENCES `tag` (`uuid`)
);
/*!40101 SET character_set_client = @saved_cs_client */;


/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-27 18:22:06