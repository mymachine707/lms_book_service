CREATE TABLE IF NOT EXISTS "author" (
   "id" CHAR(36) PRIMARY KEY,
   "name"  varchar(255),
   "status" varchar(255) DEFAULT 'enabled',
   "created_at" TIMESTAMP DEFAULT now(),
   "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "category" (
   "id" CHAR(36) PRIMARY KEY,
   "title"  varchar(255),
   "status" varchar(255) DEFAULT 'enabled',
   "created_at" TIMESTAMP DEFAULT now(),
   "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "location" (
   "id" CHAR(36) PRIMARY KEY,
   "name"  varchar(255),
   "status" varchar(255) DEFAULT 'enabled',
   "created_at" TIMESTAMP DEFAULT now(),
   "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "book" (
   "id" CHAR(36) PRIMARY KEY,
   "name"  varchar(255),
   "author_id" CHAR(36),
   "category_id" CHAR(36),
   "location_id" CHAR(36),
   "ISBN" INT,
   "quantity" INT,
   "status" varchar(255) DEFAULT 'enabled',
   "created_at" TIMESTAMP DEFAULT now(),
   "updated_at" TIMESTAMP
);

ALTER TABLE "book" ADD CONSTRAINT fk_book_author FOREIGN KEY ("author_id") REFERENCES "author" ("id");
ALTER TABLE "book" ADD CONSTRAINT fk_book_category FOREIGN KEY ("category_id") REFERENCES "category" ("id");
ALTER TABLE "book" ADD CONSTRAINT fk_book_location FOREIGN KEY ("location_id") REFERENCES "location" ("id");

