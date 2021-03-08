CREATE TABLE "currency" (
    "id" serial NOT NULL,
	"name" varchar(255) NOT NULL,
    UNIQUE ("name"),
	CONSTRAINT "currency_pk" PRIMARY KEY ("id")
);

INSERT INTO currency ("name") VALUES ('EUR');
INSERT INTO currency ("name") VALUES ('USD');