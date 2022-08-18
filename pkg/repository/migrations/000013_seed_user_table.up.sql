BEGIN;

INSERT INTO "user" (name, username, password, sex, email, role_id)
VALUES ('test', 'test', '$2a$14$JO.ESenz7vAX4W3BtlyxeOq87Q2VLXDrb.IF6OV7P7ALWVAoZ77MS', 'man', 'test@test.test', 1);

COMMIT;