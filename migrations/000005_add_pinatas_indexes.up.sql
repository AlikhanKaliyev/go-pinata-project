CREATE INDEX IF NOT EXISTS pinatas_title_idx ON pinatas USING GIN (to_tsvector('simple', color));
CREATE INDEX IF NOT EXISTS pinatas_contents_idx ON pinatas USING GIN (contents);