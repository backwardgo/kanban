CREATE TABLE teams (
  id          serial PRIMARY KEY,
  name        text NOT NULL,
  slug        text NOT NULL,
  description text NOT NULL,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  deleted_at timestamp without time zone     NULL,
  updated_at timestamp without time zone NOT NULL,

  created_by  int  NOT NULL references users
);

CREATE TRIGGER stamp_updated_at BEFORE INSERT OR UPDATE
ON teams FOR EACH ROW EXECUTE PROCEDURE updated_at_stamper();

CREATE UNIQUE INDEX index_teams_on_slug ON teams (lower(slug));
