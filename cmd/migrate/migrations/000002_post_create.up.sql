CREATE TABLE if not exists posts (
	id BIGSERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	user_id BIGINT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now()
);
