-- Schema-only definitions for sqlc.
--
-- Migrations may create some types inside DO blocks, which sqlc does not parse.
-- Keep this file minimal and only for objects referenced by queries.

CREATE TYPE status_responses AS ENUM ('pending', 'accepted', 'rejected');
