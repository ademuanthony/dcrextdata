create table data (
id serial primary key not null
);

create table New_table (
        id serial primary key not null ,
        Exchangeid numeric,
		Globaltradeid numeric,
		Tradeid numeric,
		Timestamping Timestamp,
	    Quantity numeric,
		Price numeric,
		Total numeric,
		FillType VARCHAR (20),
		OrderType varchar (20)
);


create table historic_data (
id serial primary key not null ,
exchangeID numeric not null,
globaltradeid numeric not null,
tradeid numeric not null,
timestamp timestamp not null,
quantity  varchar (30) not null,
price numeric not null, 
total numeric not null,
fill_type varchar (20) not null,
order_type varchar (20) not null
);

create table chart_data (
    id serial primary key not null ,
    exchangeID integer ,
    date       timestamp,
    high    varchar (20),
    low     varchar (20),
    open1   varchar (20),
    close1   varchar (20),
    volume  varchar (20),
    quotevolume varchar (20),
    weightedaverage varchar (20)
);

create table POSData (
    id serial primary key not null ,
    POSid varchar (20),
    Apienabled varchar (10) ,
    APIVersionsSupported varchar (20) ,
    Network varchar (20),
    URL varchar (50) ,
    Launched numeric,
    LastUpdated numeric,
    Immature numeric,
    Live numeric,
    Voted numeric,
    Missed numeric,
    PoolFees numeric,
    ProportionLive numeric,
    ProportionMissed numeric,
    UserCount numeric,
    UserCountActive numeric

)

