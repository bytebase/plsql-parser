ALTER TABLE table1 MODIFY PARTITION BY LIST(color)(
  partition RED VALUES ('RED'),
  partition BLUE VALUES ('BLUE')
);