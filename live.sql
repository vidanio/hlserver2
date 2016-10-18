/*
Navicat SQLite Data Transfer

Source Server         : nueva prueba
Source Server Version : 30808
Source Host           : :0

Target Server Type    : SQLite
Target Server Version : 30808
File Encoding         : 65001

Date: 2016-10-18 19:03:06
*/

PRAGMA foreign_keys = OFF;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS "main"."admin";
CREATE TABLE "admin" (
"id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
"username"  TEXT(255),
"password"  TEXT(255),
"type"  INTEGER,
"status"  INTEGER
);

-- ----------------------------
-- Table structure for encoders
-- ----------------------------
DROP TABLE IF EXISTS "main"."encoders";
CREATE TABLE "encoders" (
"id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
"username"  TEXT(255),
"streamname"  TEXT(255),
"time"  INTEGER,
"bitrate"  INTEGER,
"ip"  TEXT(255),
"info"  TEXT(255),
"isocode"  TEXT(4),
"country"  TEXT(255),
"region"  TEXT(255),
"city"  TEXT(255),
"timezone"  TEXT(255),
"lat"  TEXT(255),
"long"  TEXT(255),
"timestamp"  INTEGER
);

-- ----------------------------
-- Table structure for players
-- ----------------------------
DROP TABLE IF EXISTS "main"."players";
CREATE TABLE "players" (
"id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
"username"  TEXT(255),
"streamname"  TEXT(255),
"os"  TEXT(7),
"ipproxy"  TEXT(255),
"ipclient"  TEXT(255),
"isocode"  TEXT(4),
"country"  TEXT(255),
"region"  TEXT(255),
"city"  TEXT(255),
"timezone"  TEXT(255),
"lat"  TEXT(255),
"long"  TEXT(255),
"timestamp"  INTEGER,
"time"  INTEGER,
"kilobytes"  INTEGER,
"total_time"  INTEGER
);

-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "main"."sqlite_sequence";
CREATE TABLE sqlite_sequence(name,seq);
