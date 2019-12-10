CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `auth_uid` varchar(255) NOT NULL,
  `auth_provider` varchar(20) NOT NULL,
  `icon_url` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp,
  `updated_at` datetime NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_users_on_uid_and_provider` (`auth_uid`,`auth_provider`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `calendars` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `year` smallint unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp,
  `updated_at` datetime NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_calendars_user_id` FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `entries` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `calendar_id` int unsigned NOT NULL,
  `day` tinyint unsigned NOT NULL,
  `comment` text NOT NULL,
  `url` text NOT NULL,
  `title` text NOT NULL,
  `image_url` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp,
  `updated_at` datetime NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_entries_on_calendar_id_and_day` (`calendar_id`,`day`) USING BTREE,
  CONSTRAINT `fk_entries_calendar_id` FOREIGN KEY (calendar_id) REFERENCES calendars (id) ON DELETE CASCADE,
  CONSTRAINT `fk_entries_user_id` FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
