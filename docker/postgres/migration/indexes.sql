\connect bank

CREATE INDEX IF NOT EXISTS idx_accounts_user_id ON accounts (user_id);
CREATE INDEX IF NOT EXISTS idx_accounts_currency_id ON accounts (currency_id);
CREATE INDEX IF NOT EXISTS idx_cards_user_id ON cards (user_id);
CREATE INDEX IF NOT EXISTS idx_cards_account_id ON cards (account_id);
CREATE INDEX IF NOT EXISTS idx_credits_user_id ON credits (user_id);
CREATE INDEX IF NOT EXISTS idx_credits_currency_id ON credits (currency_id);
CREATE INDEX IF NOT EXISTS idx_deposits_user_id ON deposits (user_id);
CREATE INDEX IF NOT EXISTS idx_deposits_currency_id ON deposits (currency_id);
CREATE INDEX IF NOT EXISTS idx_currencies_symbol ON currencies (symbol);
CREATE INDEX IF NOT EXISTS idx_currencies_minor_units ON currencies (minor_units);

\connect bank_test

CREATE INDEX IF NOT EXISTS idx_accounts_user_id ON accounts (user_id);
CREATE INDEX IF NOT EXISTS idx_accounts_currency_id ON accounts (currency_id);
CREATE INDEX IF NOT EXISTS idx_cards_user_id ON cards (user_id);
CREATE INDEX IF NOT EXISTS idx_cards_account_id ON cards (account_id);
CREATE INDEX IF NOT EXISTS idx_credits_user_id ON credits (user_id);
CREATE INDEX IF NOT EXISTS idx_credits_currency_id ON credits (currency_id);
CREATE INDEX IF NOT EXISTS idx_deposits_user_id ON deposits (user_id);
CREATE INDEX IF NOT EXISTS idx_deposits_currency_id ON deposits (currency_id);
CREATE INDEX IF NOT EXISTS idx_currencies_symbol ON currencies (symbol);
CREATE INDEX IF NOT EXISTS idx_currencies_minor_units ON currencies (minor_units);
