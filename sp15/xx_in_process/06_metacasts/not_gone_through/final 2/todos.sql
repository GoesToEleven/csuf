/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 90301
 Source Host           : localhost
 Source Database       : godos_development
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90301
 File Encoding         : utf-8

 Date: 01/25/2014 17:28:25 PM
*/

-- ----------------------------
--  Sequence structure for todos_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."todos_id_seq";
CREATE SEQUENCE "public"."todos_id_seq" INCREMENT 1 START 13 MAXVALUE 9223372036854775807 MINVALUE 1 CACHE 1;
ALTER TABLE "public"."todos_id_seq" OWNER TO "postgres";

-- ----------------------------
--  Table structure for todos
-- ----------------------------
DROP TABLE IF EXISTS "public"."todos";
CREATE TABLE "public"."todos" (
	"id" int4 NOT NULL DEFAULT nextval('todos_id_seq'::regclass),
	"subject" varchar(255) NOT NULL COLLATE "default",
	"description" text NOT NULL COLLATE "default",
	"completed" bool DEFAULT false,
	"created_at" timestamp(6) NOT NULL,
	"updated_at" timestamp(6) NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."todos" OWNER TO "postgres";

-- ----------------------------
--  Records of todos
-- ----------------------------
BEGIN;
INSERT INTO "public"."todos" VALUES ('1', 'Learn Go!', 'Go is really cool.', 'f', '2014-01-06 11:15:19', '2014-01-08 14:25:21');
INSERT INTO "public"."todos" VALUES ('2', 'Buy Milk', '', 'f', '2014-01-07 18:37:15', '2014-01-07 11:37:19');
INSERT INTO "public"."todos" VALUES ('3', 'Shovel Snow', 'Snow sucks, but someone needs to shovel it!', 't', '2014-01-05 10:46:58', '2014-01-08 10:38:00');
INSERT INTO "public"."todos" VALUES ('4', 'Watch Chuck', 'It''s a really fun show!', 'f', '2014-01-09 10:05:07.590939', '2014-01-09 10:05:07.590939');
COMMIT;


-- ----------------------------
--  Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."todos_id_seq" RESTART 14;
-- ----------------------------
--  Primary key structure for table todos
-- ----------------------------
ALTER TABLE "public"."todos" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

