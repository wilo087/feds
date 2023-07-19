DROP TABLE IF EXISTS `feeds`;

create table `feeds` (
  `id` int(11) unsigned not null auto_increment,
  `title` varchar(255) not null,
  `description` varchar(255) not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`)
);
