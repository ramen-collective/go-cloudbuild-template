-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: 172.17.0.2    Database: template-api
-- ------------------------------------------------------
-- Server version	8.0.27

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (uuid, name, email, password, remember_token, created_at, updated_at) VALUES ('de2cae14-5c34-40a4-bd3c-37a8a8dbaebb', 'Jean', 'jean@gmail.com', '$2y$12$Nrxr9De4tfVZVx0uuayXsOoZO8XL7Mins4ZrsuZzW5OwuxUjCsYpG', '', NOW(), NOW()); /* password: test */
INSERT INTO `user` (uuid, name, email, password, remember_token, created_at, updated_at) VALUES ('21054bbd-4930-437b-8628-f2dba04e56bc', 'Marc', 'marc@gmail.com', '$2y$12$Vj9yrw6oCyR1DEqORfZy9emH3.3LeZfcfj6QEp9svWp0Qnl.coUG6', '', NOW(), NOW()); /* password: test */
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` (name, slug, created_at, updated_at) VALUES ('Default', 'default', NOW(), NOW());
INSERT INTO `role` (name, slug, created_at, updated_at) VALUES ('Editor', 'editor', NOW(), NOW());
INSERT INTO `role` (name, slug, created_at, updated_at) VALUES ('Administrator', 'admin', NOW(), NOW());
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `user_role` WRITE;
INSERT INTO `user_role` (user_uuid, role_id, created_at, updated_at) VALUES ('de2cae14-5c34-40a4-bd3c-37a8a8dbaebb', 2, NOW(), NOW());
/*!40000 ALTER TABLE `user_role` DISABLE KEYS */;
INSERT INTO `user_role` (user_uuid, role_id, created_at, updated_at) VALUES ('21054bbd-4930-437b-8628-f2dba04e56bc', 3, NOW(), NOW());
/*!40000 ALTER TABLE `user_role` ENABLE KEYS */;
UNLOCK TABLES;


LOCK TABLES `content` WRITE;
/*!40000 ALTER TABLE `content` DISABLE KEYS */;
INSERT INTO `content` (uuid, name, description, state, created_at, updated_at) VALUES ('b225a28d-30d3-4a2d-b12a-90356c9f8d8a', 'Content 1', 'This is the first content.', 'draft', NOW(), NOW());
INSERT INTO `content` (uuid, name, description, state, created_at, updated_at) VALUES ('1cec1abd-5162-4d73-87a1-896994cb687c', 'Content 2', 'Second content, is coming after the first.', 'draft', NOW(), NOW());
/*!40000 ALTER TABLE `content` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `collection` WRITE;
/*!40000 ALTER TABLE `collection` DISABLE KEYS */;
INSERT INTO `collection` (uuid, name, description, state, created_at, updated_at) VALUES ('92f09387-3662-4014-a7b9-e8b6834a6e9d', 'Collection 1', 'First collection to come.', 'draft', NOW(), NOW());
INSERT INTO `collection` (uuid, name, description, state, created_at, updated_at) VALUES ('4a16b748-7641-4819-84df-dfd4a7376e68', 'Collection 2', 'Here is the second collection.', 'draft', NOW(), NOW());
/*!40000 ALTER TABLE `collection` ENABLE KEYS */;
UNLOCK TABLES;

LOCK TABLES `collection_content` WRITE;
/*!40000 ALTER TABLE `collection_content` DISABLE KEYS */;
INSERT INTO `collection_content` (uuid, collection_uuid, content_uuid, position, created_at, updated_at) VALUES ('cf0ef9cf-df7b-4ae4-b77b-d36872143ea2', '92f09387-3662-4014-a7b9-e8b6834a6e9d', 'b225a28d-30d3-4a2d-b12a-90356c9f8d8a', 1, NOW(), NOW());
INSERT INTO `collection_content` (uuid, collection_uuid, content_uuid, position, created_at, updated_at) VALUES ('eafc60dd-cec1-4657-91db-9ff6d84cb0fd', '4a16b748-7641-4819-84df-dfd4a7376e68', '1cec1abd-5162-4d73-87a1-896994cb687c', 1, NOW(), NOW());
/*!40000 ALTER TABLE `collection_content` ENABLE KEYS */;

LOCK TABLES `tag` WRITE;
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
INSERT INTO `tag` (uuid, slug, name, created_at, updated_at) VALUES ('62681872-3eb7-437d-b306-669c9632fd75', 'first', 'First tag', NOW(), NOW());
INSERT INTO `tag` (uuid, slug, name, created_at, updated_at) VALUES ('8b5f9fe3-6e74-48b1-8c34-dba2ab11ff8c', 'second', 'Second tag', NOW(), NOW());
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;

LOCK TABLES `item_tag` WRITE;
/*!40000 ALTER TABLE `item_tag` DISABLE KEYS */;
INSERT INTO `item_tag` (uuid, item_uuid, item_type, tag_uuid, created_at, updated_at) VALUES ('98862272-62d7-430b-a8b5-e08fdf48c7fe', '62681872-3eb7-437d-b306-669c9632fd75', 'tag', '8b5f9fe3-6e74-48b1-8c34-dba2ab11ff8c', NOW(), NOW());
INSERT INTO `item_tag` (uuid, item_uuid, item_type, tag_uuid, created_at, updated_at) VALUES ('f28209e2-599d-4644-aba7-902374a753bd', 'eafc60dd-cec1-4657-91db-9ff6d84cb0fd', 'collection', '8b5f9fe3-6e74-48b1-8c34-dba2ab11ff8c', NOW(), NOW());
INSERT INTO `item_tag` (uuid, item_uuid, item_type, tag_uuid, created_at, updated_at) VALUES ('76e464ed-55c1-4f11-b8d3-5cf4be40cf3e', 'b225a28d-30d3-4a2d-b12a-90356c9f8d8a', 'content', '62681872-3eb7-437d-b306-669c9632fd75', NOW(), NOW());
/*!40000 ALTER TABLE `item_tag` ENABLE KEYS */;

UNLOCK TABLES;