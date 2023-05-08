The plsql-parser is a parser for Oracle 11g/12c PL/SQL. It is based on the [ANTLR4](https://github.com/antlr/antlr4) and use the grammar from [antlr4-grammars-plsql](https://github.com/antlr/grammars-v4/tree/master/sql/plsql).

## Build

Before build, you need to install the ANTLR4.

requirements:
- https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- https://github.com/antlr/antlr4/blob/master/doc/go-target.md

```bash
./build.sh
```

## Update grammar

### Manually change the grammar file in this project

1. Update the `PlSqlLexer.g4` and `PlSqlParser.g4` in root directory.
2. run `./build.sh` to generate the parser code.

### From antlr4-grammars-plsql

1. Clone the `PlSqlLexer.g4` and `PlSqlParser.g4` grammar files from https://github.com/antlr/grammars-v4/tree/master/sql/plsql into `parser` folder.
2. Clone the https://github.com/antlr/grammars-v4/blob/master/sql/plsql/Go/transformGrammar.py into root directory.
3. Run the `transformGrammar.py` to generate the `PlSqlLexer.g4` and `PlSqlParser.g4` in `parser` directory.
4. Copy the `PlSqlLexer.g4` and `PlSqlParser.g4` to root directory.
5. run `./build.sh` to generate the parser code.

## Test the parser

Run `TestPLSQLParser` in `parser_test.go` to test the parser.

```bash
go test -v
```

## References

- Oracle SQL Language Reference: https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Introduction-to-Oracle-SQL.html
- ANTLR4 Getting Started https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- ANTLR4 Go Garget https://github.com/antlr/antlr4/blob/master/doc/go-target.md