CREATE TABLE "users_balance_transactions" (
    "id" serial NOT NULL,
    "user_id" integer NOT NULL,
	"balance_id" integer NOT NULL,
	"balance_before" integer NOT NULL,
	"balance_after" integer NOT NULL,
	"time_placed" timestamp NOT NULL,
	"transaction_type" smallint NOT NULL,
	"created_at" timestamp NOT NULL DEFAULT NOW(),
	CONSTRAINT "users_balance_transactions_pk" PRIMARY KEY ("id"),
    CONSTRAINT "users_balance_transactions_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id"),
    CONSTRAINT "users_balance_transactions_fk1" FOREIGN KEY ("balance_id") REFERENCES "users_balance"("id")
)