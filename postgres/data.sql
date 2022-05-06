insert into category(description)
values ('fruits');
insert into category(description)
values ('books');

insert into products(nome, price, quantity, amount, category)
values ('laranja', 1.80, 100, (1.80 *100), 1);
insert into products(nome, price, quantity, amount, category)
values ('banana', 1.20, 80, (1.20 *80), 1);

insert into products(nome, price, quantity, amount, category)
values ('Harry Potter', 20, 10, (20 *10), 2);
insert into products(nome, price, quantity, amount, category)
values ('Clean Code', 59.90, 10, (59.9 *10), 2);


select category.description, products.* from products 
         inner join category on category.id = products.category
         