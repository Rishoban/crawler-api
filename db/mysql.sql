-- MySQL dump 10.13  Distrib 8.0.41, for Win64 (x86_64)
--
-- Host: localhost    Database: crawler_db
-- ------------------------------------------------------
-- Server version	8.0.41

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `crawler_process`
--

DROP TABLE IF EXISTS `crawler_process`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `crawler_process` (
  `id` int NOT NULL AUTO_INCREMENT,
  `record_id` int NOT NULL,
  `component_name` varchar(45) NOT NULL,
  `object_info` json DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `last_updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` int NOT NULL,
  `last_updated_by` int NOT NULL,
  `object_status` varchar(45) NOT NULL,
  PRIMARY KEY (`id`,`record_id`,`component_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1757 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `crawler_process`
--

LOCK TABLES `crawler_process` WRITE;
/*!40000 ALTER TABLE `crawler_process` DISABLE KEYS */;
/*!40000 ALTER TABLE `crawler_process` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `crawler_url`
--

DROP TABLE IF EXISTS `crawler_url`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `crawler_url` (
  `id` int NOT NULL AUTO_INCREMENT,
  `object_info` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` int DEFAULT NULL,
  `last_updated_by` int DEFAULT NULL,
  `object_status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `crawler_url`
--

LOCK TABLES `crawler_url` WRITE;
/*!40000 ALTER TABLE `crawler_url` DISABLE KEYS */;
INSERT INTO `crawler_url` VALUES (1,NULL,'2025-07-12 08:54:57','2025-07-13 10:20:50',1,0,'Archived'),(2,NULL,'2025-07-13 10:24:55','2025-07-13 10:25:18',1,0,'Archived'),(3,'{\"url\": \"http://books.toscrape.com/\", \"title\": \"All products | Books to Scrape - Sandbox\", \"status\": \"Done\", \"checkbox\": false, \"headings\": {\"h1\": 1, \"h2\": 0, \"h3\": 20, \"h4\": 0, \"h5\": 0, \"h6\": 0}, \"broken_links\": [{\"url\": \"http://books.toscrape.com/\", \"status\": 200}], \"html_version\": \"<!DOCTYPE html>\", \"external_links\": 2, \"has_login_form\": false, \"internal_links\": 1, \"inaccessible_links\": 0, \"presenceOfLoginForm\": false}','2025-07-13 10:25:37','2025-07-13 10:25:53',1,1,'Active');
/*!40000 ALTER TABLE `crawler_url` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `object_info` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` int DEFAULT NULL,
  `last_updated_by` int DEFAULT NULL,
  `object_status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (7,'{\"email\": \"admin@skyell.com\", \"password\": \"$2a$10$YhlrVHQuQ1BKZ5pULv1bwOyLGs6k0fSQoudCQgOUvqj0A9bogexSq\", \"username\": \"admin@skyell.com\", \"firstName\": \"Administrator FN\"}','2025-07-10 12:15:47','2025-07-10 12:15:47',1,1,'Active');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-13 19:45:50
