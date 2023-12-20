
CREATE TABLE IF NOT EXISTS "branch" (
  "id" UUID PRIMARY KEY NOT NULL,
  "name" VARCHAR(50) NOT NULL,
  "address" VARCHAR(50) NOT NULL,
  "phone_number" VARCHAR(50) NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "client" (
  "id" UUID PRIMARY KEY NOT NULL,
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "phone_number" VARCHAR(50) NOT NULL,
  "birthday" TIMESTAMP NOT NULL,
  "is_active" BOOLEAN DEFAULT FALSE NOT NULL,
  "gender" VARCHAR(20) NOT NULL,
  "branch_id" UUID REFERENCES "branch"("id") NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "product"(
  "id" UUID PRIMARY KEY NOT NULL,
  "name" VARCHAR(50) NOT NULL,
  "selling_price" DECIMAL NOT NULL,
  "branch_id" UUID REFERENCES "branch"("id") NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "coming"(
  "id" UUID PRIMARY KEY NOT NULL,
  "increment_id" VARCHAR(50) NOT NULL,
  "dated" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "branch_id" UUID REFERENCES "branch"("id") NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "picking_sheet"(
  "id" UUID PRIMARY KEY NOT NULL,
  "increment_id" VARCHAR(15) NOT NULL,
  "product_id" UUID REFERENCES "product"("id") NOT NULL,
  "coming_id" UUID REFERENCES "coming"("id") NOT NULL,
  "price" DECIMAL NOT NULL,
  "quantity" INTEGER NOT NULL,
  "total" DECIMAL NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "remainder"(
  "id" UUID PRIMARY KEY NOT NULL,
  "name" VARCHAR(50) UNIQUE NOT NULL,
  "quantity" INTEGER NOT NULL,
  "arrival_price" DECIMAL NOT NULL,
  "selling_price" DECIMAL NOT NULL,
  "branch_id" UUID REFERENCES "branch"("id") NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);



CREATE TABLE IF NOT EXISTS "sale"(
  "id" UUID PRIMARY KEY NOT NULL,
  "increment_id" VARCHAR(15) NOT NULL,
  "client_id" UUID REFERENCES "client"("id") NOT NULL,
  "branch_id" UUID REFERENCES "branch"("id") NOT NULL,
  "total" DECIMAL NOT NULL,
  "debt" DECIMAL NOT NULL,
  "paid"  DECIMAL NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "sale_product"(
  "id" UUID PRIMARY KEY NOT NULL,
  "increment_id" VARCHAR(15) NOT NULL,
  "product_id" UUID REFERENCES "product"("id") NOT NULL,
  "sale_id" UUID REFERENCES "sale"("id") NOT NULL,
  "price" DECIMAL NOT NULL,
  "quantity" INTEGER NOT NULL,
  "total" DECIMAL NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);

SELECT * FROM sale_product AS sp 

JOIN sale AS s ON sp.sale_id = s.id
