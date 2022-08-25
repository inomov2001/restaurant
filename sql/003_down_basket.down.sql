drop table basket;
create table basket(
                       chek_id UUID not null references chek(id),
                       product_id UUID not null references products(id)
);