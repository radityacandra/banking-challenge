CREATE TABLE IF NOT EXISTS users (
  id VARCHAR PRIMARY KEY,
  name VARCHAR NOT NULL,
  phone_number VARCHAR NOT NULL,
  identity_number VARCHAR NOT NULL,
  created_at int8 NOT NULL,
  created_by VARCHAR NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE TABLE IF NOT EXISTS user_accounts (
  id VARCHAR PRIMARY KEY,
  user_id VARCHAR NOT NULL,
  account_number VARCHAR NOT NULL,
  total_balance int8 NOT NULL,
  created_at int8 NOT NULL,
  created_by VARCHAR NOT NULL,
  updated_at int8,
  updated_by VARCHAR
);

CREATE TABLE IF NOT EXISTS transactions (
  id VARCHAR PRIMARY KEY,
  user_account_id VARCHAR NOT NULL,
  transaction_type VARCHAR NOT NULL,
  amount int8 NOT NULL,
  created_at int8 NOT NULL,
  created_by VARCHAR NOT NULL,
  updated_at int8,
  updated_by VARCHAR
)