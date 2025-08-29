alter table users
add username varchar(100) not null;

alter table users 
add constraint unique unisque_username (username);