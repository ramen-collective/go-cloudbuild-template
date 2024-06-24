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
    `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
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
  `collection_uuid` varchar(36) NOT NULL,
  `content_uuid` varchar(36) NOT NULL,
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


--
-- Table structure for table `password_resets`
--

DROP TABLE IF EXISTS `password_resets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `password_resets` (
   `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
   `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
   `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `created_at` timestamp NULL DEFAULT NULL,
   KEY `password_resets_email_index` (`email`)
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `role`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `role` (
  `id` BIGINT UNSIGNED UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name`     varchar(64) NOT NULL,
  `slug` varchar(64) UNIQUE NOT NULL,
  `permissions` json NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_role`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_role` (
  `user_uuid` varchar(36) NOT NULL,
  `role_id` BIGINT UNSIGNED NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `uq_user_uuid_role_id_idx` (`user_uuid`, `role_id`),
  CONSTRAINT `fk_user_role_user_uuid` FOREIGN KEY (`user_uuid`) REFERENCES `user` (`uuid`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_role_role_id` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);
/*!40101 SET character_set_client = @saved_cs_client */;

CREATE TABLE `permission` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `resource` VARCHAR(255) NOT NULL,
    `create` BOOLEAN NOT NULL DEFAULT 0,
    `delete` BOOLEAN NOT NULL DEFAULT 0,
    `update` BOOLEAN NOT NULL DEFAULT 0,
    `view` BOOLEAN NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uq_name_idx` (`name`)
);


-- RBAC role - permission mapping table.
CREATE TABLE `role_permission` (
    `role_id` BIGINT UNSIGNED NOT NULL,
    `permission_id` BIGINT UNSIGNED NOT NULL,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY `uq_role_id_permission_id_idx` (`role_id`, `permission_id`),
    CONSTRAINT `fk_role_permission_role_id` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_role_permission_role_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-27 18:22:06