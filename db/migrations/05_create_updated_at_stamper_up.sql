CREATE OR REPLACE FUNCTION updated_at_stamper()
RETURNS TRIGGER AS $$
BEGIN
	IF TG_OP = 'INSERT' THEN
		NEW.updated_at = NEW.created_at;
	ELSE
		NEW.updated_at = now();
	END IF;

	RETURN NEW;
END;
$$
language plpgsql;
