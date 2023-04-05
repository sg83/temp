CREATE TABLE IF NOT EXISTS articles (
  id serial not null unique,
  title varchar(500) not null,
  date VARCHAR not null,
  body text not null,
  tags TEXT[],
  primary key(id)
);

CREATE TABLE IF NOT EXISTS tags (
    tag VARCHAR PRIMARY KEY,
    count INTEGER NOT NULL,
    articles INTEGER[],
    related_tags TEXT[]
);

insert into articles(title, body, tags)
values
    ('latest science shows that potato chips are better for you than sugar',
     'some text, potentially containing simple markup about how potato chips are great ...', 
     '2023-04-05'
     {"health", "fitness", "science"}),
    ('Another Post', 
    'Yet another blog post about something exciting',
    '2023-04-05'
     {"health", "fitness", "science"});