CREATE TABLE "users_balance" (
    "id" serial NOT NULL,
    "user_id" integer NOT NULL,
	"balance" integer default 0,
	"currency_id" smallint NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT NOW(),
	"updated_at" timestamp,
    UNIQUE ("user_id", "currency_id"),
	CONSTRAINT "users_balance_pk" PRIMARY KEY ("id"),
    CONSTRAINT "users_balance_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id"),
    CONSTRAINT "users_balance_fk1" FOREIGN KEY ("currency_id") REFERENCES "currency"("id")
);

INSERT INTO users_balance (user_id, balance, currency_id) VALUES (1, 0, 1);
INSERT INTO users_balance (user_id, balance, currency_id) VALUES (1, 0, 2);

INSERT INTO users_balance (user_id, balance, currency_id) VALUES (2, 0, 1);
INSERT INTO users_balance (user_id, balance, currency_id) VALUES (2, 0, 2);