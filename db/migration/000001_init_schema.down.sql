-- +migrate Down
-- SQL in section 'Up' is executed when this migration is applied

DROP TABLE IF EXISTS notices;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS noticeStatus;
DROP TABLE IF EXISTS requests;
DROP TABLE IF EXISTS requestStatus;
DROP TABLE IF EXISTS devices;
DROP TABLE IF EXISTS statusDevices;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS categories;