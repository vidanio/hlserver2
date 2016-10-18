/*
Navicat SQLite Data Transfer

Source Server         : Prueba
Source Server Version : 30808
Source Host           : :0

Target Server Type    : SQLite
Target Server Version : 30808
File Encoding         : 65001

Date: 2016-10-03 16:21:15
*/

PRAGMA foreign_keys = OFF;

-- ----------------------------
-- Table structure for resumen
-- ----------------------------
DROP TABLE IF EXISTS "main"."resumen";
CREATE TABLE "resumen" (
"id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
"username"  TEXT(255),
"streamname"  TEXT(255),
"os"  TEXT(7),
"isocode"  TEXT(5),
"time"  INTEGER,
"kilobytes"  INTEGER,
"count"  INTEGER,
"hour"  INTEGER,
"minutes"  INTEGER,
"date"  TEXT(255)
);

-- ----------------------------
-- Records of resumen
-- ----------------------------

-- ----------------------------
-- Table structure for sqlite_sequence
-- ----------------------------
DROP TABLE IF EXISTS "main"."sqlite_sequence";
CREATE TABLE sqlite_sequence(name,seq);

-- ----------------------------
-- Records of sqlite_sequence
-- ----------------------------
INSERT INTO "main"."sqlite_sequence" VALUES ('resumen', 1);
