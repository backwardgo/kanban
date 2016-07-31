CREATE TABLE users (
  id         serial PRIMARY KEY,
  first_name text   NOT NULL,
  last_name  text   NOT NULL,
  initials   text   NOT NULL,
  biography  text   NOT NULL,

  email           citext NOT NULL,
  password_digest bytea  NOT NULL,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  deleted_at timestamp without time zone     NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE TRIGGER stamp_updated_at BEFORE INSERT OR UPDATE
ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_stamper();

CREATE UNIQUE INDEX index_users_on_email ON users (lower(email));
