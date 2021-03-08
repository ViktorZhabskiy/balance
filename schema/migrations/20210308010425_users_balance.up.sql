CREATE TABLE "users_balance" (
    "id" serial NOT NULL,
    "user_id" integer NOT NULL,
	"balance" integer default 0,
	"currency" smallint NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT NOW(),
	"updated_at" timestamp,
	CONSTRAINT "users_balance_pk" PRIMARY KEY ("id"),
    CONSTRAINT "users_balance_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id")
)