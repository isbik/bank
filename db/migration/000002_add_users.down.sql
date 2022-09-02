ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IS EXISTS "owner_currency_key";
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IS EXISTS "accounts_owner  _fkey";
DROP TABLE IF EXISTS "users";