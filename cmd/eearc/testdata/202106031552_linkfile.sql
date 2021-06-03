CREATE TABLE file
(
    `id`   BINARY(16),
    `size` int(12),
    `name` varchar(255),
    PRIMARY KEY (`id`)
);

ALTER TABLE ticket
    ADD COLUMN (`file` BINARY(16)),
    ADD FOREIGN KEY (`file`) REFERENCES file (`id`);

