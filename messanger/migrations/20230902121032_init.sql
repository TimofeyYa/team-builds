-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password_hash" varchar NOT NULL,
    "update_at" timestamp,
    "created_at" timestamp DEFAULT (now())
);
CREATE UNIQUE INDEX users_item ON users(email);

CREATE TABLE "sessions" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" int NOT NULL,
    "refresh_token" varchar NOT NULL,
    "fingerprint" varchar NOT NULL,
    "expires_at" timestamp NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    FOREIGN KEY ( user_id ) references users(id)
);

CREATE TABLE "friends" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "friend_id" int UNIQUE NOT NULL,
  FOREIGN KEY ( user_id ) references users(id),
  FOREIGN KEY ( friend_id ) references users(id)
);

CREATE UNIQUE INDEX friend_item ON friends(user_id, friend_id);

CREATE TABLE "messages" (
  "id" BIGSERIAL PRIMARY KEY,
  "sendler_id" int NOT NULL,
  "recipient_id" int NOT NULL,
  "message" varchar NOT NULL,
  "is_read" boolean NOT NULL DEFAULT false,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  FOREIGN KEY ( sendler_id ) references users(id),
  FOREIGN KEY ( recipient_id ) references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX friend_item;
DROP TABLE friends;
DROP TABLE messages;
DROP TABLE sessions;
DROP TABLE users;
-- +goose StatementEnd
