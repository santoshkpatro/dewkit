package migrations

import "github.com/jmoiron/sqlx"

var V8 = Migration{
	Version: 8,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			CREATE TYPE sender_type_enum AS ENUM ('customer', 'staff', 'system');
			CREATE TABLE messages (
				id TEXT PRIMARY KEY,

				conversation_id TEXT NOT NULL,
				sender_type sender_type_enum NOT NULL,
				sender_customer_id TEXT NULL,
  				sender_staff_id TEXT NULL,

				body TEXT NOT NULL,
				body_type VARCHAR(20) DEFAULT 'text',

				is_internal BOOLEAN DEFAULT FALSE,

				created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

				CONSTRAINT fk_messages_conversation
					FOREIGN KEY (conversation_id)
					REFERENCES conversations(id)
					ON DELETE CASCADE,

				CONSTRAINT fk_messages_customer
					FOREIGN KEY (sender_customer_id)
					REFERENCES users(id)
					ON DELETE SET NULL,

				CONSTRAINT fk_messages_staff
					FOREIGN KEY (sender_staff_id)
					REFERENCES users(id)
					ON DELETE SET NULL
			);
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			DROP TABLE messages;
			DROP TYPE sender_type_enum;
			`,
		)
		return err
	},
}
