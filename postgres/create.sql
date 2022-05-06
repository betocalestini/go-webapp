create table if not exists category(
  id serial primary key,
  description varchar(100) not null
);
create table if not exists products(
  id bigserial primary key,
  nome varchar(255) not null,
  price decimal not null,
  quantity integer default 0,
  amount decimal default 0.0,
  category bigint not null,
  constraint products_category_fk foreign key(category) references category(id)
);