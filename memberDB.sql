create table produk(
    id serial primary key NOT NULL,
    nama varchar(255),
    harga int
)
create table isiSaldo(
    id serial primary key NOT NULL,
    pembayaran varchar(255),
    nominal int,
    total int
)
create table register(
    id serial primary key NOT NULL,
    nama varchar(255),
    email varchar(255),
    password varchar(255),
    role varchar(10)
)
CREATE TABLE pemesanan (
	id serial primary key NOT NULL,
	idProduk int,
	jumlah int ,
    total int,
    constraint fk_PP foreign key (idProduk) references produk(id)
)
create table profil(
    id serial primary key NOT NULL,
    idRegister int,
    idSaldo int,
    constraint fk_PR foreign key (idRegister) references register(id),
    constraint fk_PS foreign key (idSaldo) references isiSaldo(id)
)
create table transaksi(
    id serial primary key NOT NULL,
    idUser int,
    idPesan int,
    tanggal timestamp,
    pembayaran varchar(255),
    status boolean,
    constraint fk_TP foreign key (idUser) references profil(id),
    constraint fk_TPe foreign key (idPesan) references pemesanan(id)
)
create table feedback(
    id serial primary key NOT NULL,
    idTransaksi int,
    idUser int,
    komen varchar(255),
    rating int,
    createAt timestamp, 
    updatedAt timestamp,
    constraint fk_PR foreign key (idTransaksi) references transaksi(id),
    constraint fk_PS foreign key (idUser) references profil(id)
)

ALTER TABLE isiSaldo
ADD tanggal timestamp;

