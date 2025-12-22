

CREATE TYPE conversation_status AS ENUM (
    'open',
    'pending',
    'resolved',
    'archived'
);

CREATE TYPE user_role AS ENUM (
    'admin',
    'staff',
    'superuser'
);

CREATE TABLE conversations (
    id integer NOT NULL,
    customer_id bigint,
    customer_full_name text,
    customer_email text,
    status text NOT NULL,
    resolved_at timestamp with time zone,
    archived_at timestamp with time zone,
    assigned_to bigint,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE SEQUENCE conversations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE conversations_id_seq OWNED BY conversations.id;

CREATE TABLE settings (
    key text NOT NULL,
    value jsonb NOT NULL
);

CREATE TABLE users (
    id bigint NOT NULL,
    email text NOT NULL,
    full_name text NOT NULL,
    password_hash text NOT NULL,
    password_salt text,
    is_active boolean DEFAULT true NOT NULL,
    is_password_expired boolean DEFAULT false NOT NULL,
    last_login_at timestamp with time zone,
    role text DEFAULT 'staff'::text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    password_changed_at timestamp without time zone,
    failed_login_attempts integer DEFAULT 0 NOT NULL
);

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE users_id_seq OWNED BY users.id;

ALTER TABLE ONLY conversations ALTER COLUMN id SET DEFAULT nextval('conversations_id_seq'::regclass);

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);

ALTER TABLE ONLY conversations
    ADD CONSTRAINT conversations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY conversations
    ADD CONSTRAINT conversations_assigned_to_fkey FOREIGN KEY (assigned_to) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE ONLY conversations
    ADD CONSTRAINT conversations_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE SET NULL;


-- Ensure db.version is set to latest migration
INSERT INTO settings (key, value)
VALUES ('db.version', to_jsonb(4::int))
ON CONFLICT (key)
DO UPDATE SET value = EXCLUDED.value;
	