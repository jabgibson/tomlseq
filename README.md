# tomlseq
a toml preprocessor for adding sequence to toml tables

## why?
While building out toml documents
I realized that if I wanted to id arrays of tables, without having
to manually enter an id for each table, I could instead write some
software to generate it. The TOML spec didn't require that arrays keep
their sequential order, and.. if I wanted to abstract the type
of each table to an interface, I wanted to preserve the order
of the table within its original form.


