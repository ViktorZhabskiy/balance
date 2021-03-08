CREATE TABLE "users" (
    "id" serial NOT NULL,
	"name" varchar(255) NOT NULL,
    UNIQUE ("name"),
	CONSTRAINT "users_pk" PRIMARY KEY ("id")
);

INSERT INTO users ("name") VALUES ('Viktor');
INSERT INTO users ("name") VALUES ('Alex');