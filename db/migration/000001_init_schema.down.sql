-- +migrate Down
-- SQL in section 'Up' is executed when this migration is applied

DROP TABLE facilityManage.categories;

DROP TABLE facilityManage.devices;

DROP TABLE facilityManage.gorp_migrations;

DROP TABLE facilityManage.noticeStatus;

DROP TABLE facilityManage.notices;

DROP TABLE facilityManage.requestStatus;

DROP TABLE facilityManage.requests;

DROP TABLE facilityManage.roles;

DROP TABLE facilityManage.schema_migrations;

DROP TABLE facilityManage.sessions;

DROP TABLE facilityManage.statusDevices;

DROP TABLE facilityManage.users;
