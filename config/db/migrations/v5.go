package migrations

import "github.com/jmoiron/sqlx"

var V5 = Migration{
	Version: 5,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(`
			CREATE TYPE project_member_role AS ENUM ('admin', 'collaborator');

			CREATE TABLE projects (
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL,
				description TEXT,
				code TEXT NOT NULL UNIQUE,
				created_by_id TEXT NOT NULL,
				created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

				CONSTRAINT fk_projects_created_by
					FOREIGN KEY (created_by_id)
					REFERENCES users(id)
					ON DELETE CASCADE
			);

			CREATE TABLE project_members (
				id TEXT PRIMARY KEY,
				project_id TEXT NOT NULL,
				user_id TEXT NOT NULL,
				role project_member_role NOT NULL,
				created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

				CONSTRAINT fk_project_members_project
					FOREIGN KEY (project_id)
					REFERENCES projects(id)
					ON DELETE CASCADE,

				CONSTRAINT fk_project_members_user
					FOREIGN KEY (user_id)
					REFERENCES users(id)
					ON DELETE CASCADE,

				CONSTRAINT project_members_unique
					UNIQUE (project_id, user_id)
			);

			ALTER TABLE conversations
			ADD COLUMN project_id TEXT;

			ALTER TABLE conversations
			ADD CONSTRAINT fk_conversations_project
				FOREIGN KEY (project_id)
				REFERENCES projects(id)
				ON DELETE SET NULL;
	`)
		return err
	},

	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE conversations
			DROP CONSTRAINT IF EXISTS fk_conversations_project;

			ALTER TABLE conversations
			DROP COLUMN IF EXISTS project_id;

			DROP TABLE IF EXISTS project_members;
			DROP TABLE IF EXISTS projects;

			DROP TYPE IF EXISTS project_member_role;
			`,
		)
		return err
	},
}
