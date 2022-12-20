

BEGIN;

INSERT INTO "author" ("id", "name") VALUES ('b9a39905-d3d4-4025-8905-21aa625a1aef', 'Leo Tolstoy') ON CONFLICT DO NOTHING;
INSERT INTO "author" ("id", "name") VALUES ('b9a39905-d3d4-4025-8905-22aa625a1aef', 'Gustav Flaubert') ON CONFLICT DO NOTHING;
INSERT INTO "author" ("id", "name") VALUES ('b9a39905-d3d4-4025-8905-23aa625a1aef', 'Vladimir Nabokov') ON CONFLICT DO NOTHING;
INSERT INTO "author" ("id", "name") VALUES ('b9a39905-d3d4-4025-8905-24aa625a1aef', 'Mark Twain') ON CONFLICT DO NOTHING;
INSERT INTO "author" ("id", "name") VALUES ('b9a39905-d3d4-4025-8905-25aa625a1aef', 'William Shakespeare') ON CONFLICT DO NOTHING;
INSERT INTO "author" ("id", "name") VALUES ('b9a39905-d3d4-4025-8905-26aa625a1aef', 'F.Scott Fizgerald') ON CONFLICT DO NOTHING;
INSERT INTO "author" ("id", "name") VALUES ('b9a39905-d3d4-4025-8905-27aa625a1aef', 'Marcel Proust') ON CONFLICT DO NOTHING;

INSERT INTO "category" ("id", "title") VALUES ('4b19a81c-f4b5-482b-bcc4-7d517a5cfd50', 'Realist novel') ON CONFLICT DO NOTHING;
INSERT INTO "category" ("id", "title") VALUES ('4b19a81c-f4b5-482b-bcc4-7d527a5cfd50', 'Roman epopeya') ON CONFLICT DO NOTHING;
INSERT INTO "category" ("id", "title") VALUES ('4b19a81c-f4b5-482b-bcc4-7d537a5cfd50', 'Novel') ON CONFLICT DO NOTHING;
INSERT INTO "category" ("id", "title") VALUES ('4b19a81c-f4b5-482b-bcc4-7d547a5cfd50', 'Picaresque novel') ON CONFLICT DO NOTHING;
INSERT INTO "category" ("id", "title") VALUES ('4b19a81c-f4b5-482b-bcc4-7d557a5cfd50', 'Shakespearean tragedy') ON CONFLICT DO NOTHING;
INSERT INTO "category" ("id", "title") VALUES ('4b19a81c-f4b5-482b-bcc4-7d567a5cfd50', 'Tragedy') ON CONFLICT DO NOTHING;
INSERT INTO "category" ("id", "title") VALUES ('4b19a81c-f4b5-482b-bcc4-7d577a5cfd50', 'Modernist') ON CONFLICT DO NOTHING;

INSERT INTO "location" ("id", "name") VALUES ('1319540f-fe91-4f8b-8bbe-93cfae5ed381', 'A1') ON CONFLICT DO NOTHING;
INSERT INTO "location" ("id", "name") VALUES ('1319540f-fe92-4f8b-8bbe-93cfae5ed381', 'B1') ON CONFLICT DO NOTHING;
INSERT INTO "location" ("id", "name") VALUES ('1319540f-fe93-4f8b-8bbe-93cfae5ed381', 'C1') ON CONFLICT DO NOTHING;
INSERT INTO "location" ("id", "name") VALUES ('1319540f-fe94-4f8b-8bbe-93cfae5ed381', 'D1') ON CONFLICT DO NOTHING;
INSERT INTO "location" ("id", "name") VALUES ('1319540f-fe95-4f8b-8bbe-93cfae5ed381', 'E1') ON CONFLICT DO NOTHING;
INSERT INTO "location" ("id", "name") VALUES ('1319540f-fe96-4f8b-8bbe-93cfae5ed381', 'F1') ON CONFLICT DO NOTHING;
INSERT INTO "location" ("id", "name") VALUES ('1319540f-fe97-4f8b-8bbe-93cfae5ed381', 'G1') ON CONFLICT DO NOTHING;
COMMIT;

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc1', 'Anna Karenina', 'b9a39905-d3d4-4025-8905-21aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d517a5cfd50', '1319540f-fe91-4f8b-8bbe-93cfae5ed381',123,10);

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc2', 'Madame Bovary', 'b9a39905-d3d4-4025-8905-22aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d517a5cfd50', '1319540f-fe91-4f8b-8bbe-93cfae5ed381',123,10);

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc3', 'War and Peace', 'b9a39905-d3d4-4025-8905-21aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d527a5cfd50', '1319540f-fe92-4f8b-8bbe-93cfae5ed381',123,10);

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc4', 'Lolita', 'b9a39905-d3d4-4025-8905-23aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d537a5cfd50', '1319540f-fe93-4f8b-8bbe-93cfae5ed381',123,10);

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc5', 'The Adventures of Huckleberry Finn', 'b9a39905-d3d4-4025-8905-24aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d547a5cfd50', '1319540f-fe94-4f8b-8bbe-93cfae5ed381',123,10);

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc6', 'Hamlet', 'b9a39905-d3d4-4025-8905-25aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d557a5cfd50', '1319540f-fe95-4f8b-8bbe-93cfae5ed381',123,10);

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc7', 'The Great Gatsby', 'b9a39905-d3d4-4025-8905-26aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d567a5cfd50', '1319540f-fe96-4f8b-8bbe-93cfae5ed381',123,10);

INSERT INTO "book" ("id", "name", "author_id", "category_id", "location_id", "ISBN", "quantity") VALUES 
(
	'433d6ba4-3d1d-4256-80d7-830bd6b30cc8', 'In Search of Lost Time', 'b9a39905-d3d4-4025-8905-27aa625a1aef', '4b19a81c-f4b5-482b-bcc4-7d577a5cfd50', '1319540f-fe97-4f8b-8bbe-93cfae5ed381',123,10);


-- Anna Karenina by Leo Tolstoy. // Realist novel
-- Madame Bovary by Gustav Flaubert. // Realist novel
-- War and Peace by Leo Tolstoy. // roman epopeya
-- Lolita by Vladimir Nabokov. // novel
-- The Adventures of Huckleberry Finn by Mark Twain. //  Picaresque novel
-- Hamlet by William Shakespeare. // Shakespearean tragedy
-- The Great Gatsby by F. Scott Fizgerald. //  Tragedy
-- In Search of Lost Time by Marcel Proust // Modernist