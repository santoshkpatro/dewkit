INSERT INTO settings (key, value) VALUES
('db.version', to_jsonb(0::int)),
('app.baseUrl', to_jsonb('https://dewkit.app'::text)),
('app.orgName', to_jsonb('Dewkit'::text)),
('app.supportEmail', to_jsonb('support@dewkit.app'::text)),
('system.maintenance', to_jsonb(false))
ON CONFLICT (key) DO NOTHING;