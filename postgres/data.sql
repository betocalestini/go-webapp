insert into category(description)
values ('Frutas');
insert into category(description)
values ('Livros');
insert into category(description)
values ('Carros');
insert into category(description)
values ('Informatica');

insert into products(name, price, quantity, amount, category)
values ('laranja', 1.80, 100, (1.80 *100), 1);
insert into products(name, price, quantity, amount, category)
values ('banana', 1.20, 80, (1.20 *80), 1);

insert into products(name, price, quantity, amount, category)
values ('Harry Potter', 20, 10, (20 *10), 2);
insert into products(name, price, quantity, amount, category)
values ('Clean Code', 59.90, 10, (59.9 *10), 2);

insert into products(name, price, quantity, amount, category)
values ('Ferrari', 3000000, 1, (3000000 *1), 3);
insert into products(name, price, quantity, amount, category)
values ('Jaguar', 800000, 10, (800000 *10), 3);

insert into products(name, price, quantity, amount, category)
values ('Notebook', 3000, 15, (3000 *1), 4);
insert into products(name, price, quantity, amount, category)
values ('Desktop', 5000, 10, (5000 *10), 4);


select category.description, products.* from products 
         inner join category on category.id = products.category
         