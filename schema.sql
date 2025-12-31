

CREATE TYPE conversation_status AS ENUM (
    'open',
    'pending',
    'resolved',
    'archived'
);

CREATE TYPE project_member_role AS ENUM (
    'admin',
    'collaborator'
);

CREATE TYPE sender_type_enum AS ENUM (
    'customer',
    'staff',
    'system'
);

CREATE TYPE user_role AS ENUM (
    'admin',
    'staff',
    'superuser'
);

CREATE TABLE conversations (
    id text NOT NULL,
    customer_id text,
    customer_full_name text,
    customer_email text,
    status text NOT NULL,
    resolved_at timestamp with time zone,
    archived_at timestamp with time zone,
    assigned_to text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    project_id text
);

CREATE TABLE messages (
    id text NOT NULL,
    conversation_id text NOT NULL,
    sender_type sender_type_enum NOT NULL,
    sender_customer_id text,
    sender_staff_id text,
    body text NOT NULL,
    body_type character varying(20) DEFAULT 'text'::character varying,
    is_internal boolean DEFAULT false,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE project_members (
    id text NOT NULL,
    project_id text NOT NULL,
    user_id text NOT NULL,
    role project_member_role NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE projects (
    id text NOT NULL,
    name text NOT NULL,
    description text,
    code text NOT NULL,
    created_by_id text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE settings (
    key text NOT NULL,
    value jsonb NOT NULL
);

CREATE TABLE users (
    id text NOT NULL,
    email text NOT NULL,
    full_name text NOT NULL,
    password_hash text,
    password_salt text,
    is_active boolean DEFAULT true NOT NULL,
    is_password_expired boolean DEFAULT false NOT NULL,
    last_login_at timestamp with time zone,
    role text DEFAULT 'staff'::text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    password_changed_at timestamp without time zone,
    failed_login_attempts integer DEFAULT 0 NOT NULL,
    is_customer boolean DEFAULT false NOT NULL,
    customer_identifier text
);

ALTER TABLE ONLY conversations
    ADD CONSTRAINT conversations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY project_members
    ADD CONSTRAINT project_members_pkey PRIMARY KEY (id);

ALTER TABLE ONLY project_members
    ADD CONSTRAINT project_members_unique UNIQUE (project_id, user_id);

ALTER TABLE ONLY projects
    ADD CONSTRAINT projects_code_key UNIQUE (code);

ALTER TABLE ONLY projects
    ADD CONSTRAINT projects_name_unique UNIQUE (name);

ALTER TABLE ONLY projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);

ALTER TABLE ONLY settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (key);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_customer_identifier_unique UNIQUE (customer_identifier);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE INDEX idx_users_is_customer ON users USING btree (is_customer);

ALTER TABLE ONLY conversations
    ADD CONSTRAINT conversations_assigned_to_fkey FOREIGN KEY (assigned_to) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE ONLY conversations
    ADD CONSTRAINT conversations_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE ONLY conversations
    ADD CONSTRAINT fk_conversations_project FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE SET NULL;

ALTER TABLE ONLY messages
    ADD CONSTRAINT fk_messages_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE;

ALTER TABLE ONLY messages
    ADD CONSTRAINT fk_messages_customer FOREIGN KEY (sender_customer_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE ONLY messages
    ADD CONSTRAINT fk_messages_staff FOREIGN KEY (sender_staff_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE ONLY project_members
    ADD CONSTRAINT fk_project_members_project FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE;

ALTER TABLE ONLY project_members
    ADD CONSTRAINT fk_project_members_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE ONLY projects
    ADD CONSTRAINT fk_projects_created_by FOREIGN KEY (created_by_id) REFERENCES users(id) ON DELETE CASCADE;



-- Seed data
INSERT INTO settings (key, value) VALUES
('db.version', to_jsonb(0::int)),
('app.baseUrl', to_jsonb('https://dewkit.app'::text)),
('app.orgName', to_jsonb('Dewkit'::text)),
('app.supportEmail', to_jsonb('support@dewkit.app'::text)),
('system.maintenance', to_jsonb(false))
ON CONFLICT (key) DO NOTHING;
-- Ensure db.version is set to latest migration
UPDATE settings
SET value = to_jsonb(8::int)
WHERE key = 'db.version';
