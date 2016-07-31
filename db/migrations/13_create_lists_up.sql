CREATE TABLE lists (
  id      serial PRIMARY KEY,
  team_id int  NOT NULL references teams,

  board_id int  NOT NULL references boards,
  title    text NOT NULL,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  deleted_at timestamp without time zone     NULL,
  updated_at timestamp without time zone NOT NULL,

  created_by  int  NOT NULL references users
);

CREATE TRIGGER stamp_updated_at BEFORE INSERT OR UPDATE
ON lists FOR EACH ROW EXECUTE PROCEDURE updated_at_stamper();
