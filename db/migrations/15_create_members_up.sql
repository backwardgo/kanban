CREATE TABLE members (
  id serial PRIMARY KEY,

  board_id int  NOT NULL references boards,
  user_id  int  NOT NULL references users,
  role     text NOT NULL,

  created_at timestamp without time zone NOT NULL DEFAULT now(),
  deleted_at timestamp without time zone     NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE TRIGGER stamp_updated_at BEFORE INSERT OR UPDATE
ON members FOR EACH ROW EXECUTE PROCEDURE updated_at_stamper();

CREATE UNIQUE INDEX index_members_on_board_id_and_user_id ON members (board_id, user_id);
