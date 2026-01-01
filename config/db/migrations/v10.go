package migrations

import "github.com/jmoiron/sqlx"

var V10 = Migration{
	Version: 10,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			CREATE EXTENSION IF NOT EXISTS pgcrypto;

			CREATE OR REPLACE FUNCTION NewID(prefix text)
			RETURNS text
			LANGUAGE plpgsql
			AS $$
			DECLARE
			    crockford text := '0123456789ABCDEFGHJKMNPQRSTVWXYZ';

			    -- Timestamp
			    time_ms   bigint;
			    time_bits bit(48);

			    -- Randomness
			    rand_bytes bytea;
			    rand_bits_text text := '';
			    rand_bits bit(80);

			    -- ULID
			    ulid_bits bit(128);
			    result_ulid text := '';

			    i int;
			    idx int;
			BEGIN
			    -- 48-bit timestamp (milliseconds since Unix epoch)
			    time_ms := FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000);
			    time_bits := time_ms::bit(48);

			    -- 80 bits of randomness
			    rand_bytes := gen_random_bytes(10);

			    -- Convert random bytes → 80-bit string
			    FOR i IN 0..9 LOOP
			        rand_bits_text :=
			            rand_bits_text ||
			            (get_byte(rand_bytes, i)::bit(8))::text;
			    END LOOP;

			    rand_bits := rand_bits_text::bit(80);

			    -- Combine timestamp + randomness
			    ulid_bits := time_bits || rand_bits;

			    -- Crockford Base32 encoding (26 chars × 5 bits)
			    FOR i IN 0..25 LOOP
			        idx := substring(ulid_bits FROM i * 5 + 1 FOR 5)::int;
			        result_ulid := result_ulid || substr(crockford, idx + 1, 1);
			    END LOOP;

			    -- Lowercase to match Go / Python ULID output
			    RETURN prefix || '_' || lower(result_ulid);
			END;
			$$;
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			DROP FUNCTION IF EXISTS NewID(text);
			`,
		)
		return err
	},
}
