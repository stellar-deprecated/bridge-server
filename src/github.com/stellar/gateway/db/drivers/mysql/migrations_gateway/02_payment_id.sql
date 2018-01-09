-- +migrate Up
ALTER TABLE `SentTransaction` ADD `payment_id` VARCHAR(255) NULL DEFAULT NULL AFTER `id`, ADD UNIQUE (`payment_id`) ;

-- +migrate Down
ALTER TABLE `SentTransaction` DROP `payment_id`;
