package parser_test

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	plsqlparser "github.com/bytebase/plsql-parser"
	"github.com/stretchr/testify/require"
)

type CustomErrorListener struct {
	errors int
}

func NewCustomErrorListener() *CustomErrorListener {
	return new(CustomErrorListener)
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors += 1
	antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (l *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func TestPLSQLParser(t *testing.T) {
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	examplesSQLScript, err := os.ReadDir("examples-sql-script")
	require.NoError(t, err)

	files := append(examples, examplesSQLScript...)

	for i, file := range files {
		if file.Name() != "1.sql" {
			// continue
			break
		}
		if strings.HasSuffix(file.Name(), ".sql.tree") {
			continue
		}
		var filePath string
		if i < len(examples) {
			filePath = path.Join("examples", file.Name())
		} else {
			filePath = path.Join("examples-sql-script", file.Name())
		}
		t.Run(filePath, func(t *testing.T) {
			// t.Parallel()
			// input, err := antlr.NewFileStream(filePath)
			// require.NoError(t, err)

			// open filePath and read each line.
			f, err := os.Open(filePath)
			require.NoError(t, err)
			defer f.Close()

			scanner := bufio.NewScanner(f)
			total := 0

			// Read each line of the file
			for scanner.Scan() {
				// Print the line
				text := scanner.Text()
				// lexer := plsqlparser.NewPlSqlLexer(input)
				lexer := plsqlparser.NewPlSqlLexer(antlr.NewInputStream(text))

				stream := antlr.NewCommonTokenStream(lexer, 0)
				p := plsqlparser.NewPlSqlParser(stream)
				p.SetVersion12(true)

				lexerErrors := &CustomErrorListener{}
				lexer.RemoveErrorListeners()
				lexer.AddErrorListener(lexerErrors)

				parserErrors := &CustomErrorListener{}
				p.RemoveErrorListeners()
				p.AddErrorListener(parserErrors)

				// var trees []plsqlparser.ISingle_sqlContext
				p.BuildParseTrees = true
				// number := 0
				// last := 0
				// fmt.Println("Start")
				_ = p.Sql_script()
				// fmt.Println(tree.GetStop().GetTokenIndex())
				// for {
				// 	tree := p.Single_sql()
				// 	trees = append(trees, tree)
				// 	number++
				// 	if last == tree.GetStop().GetTokenIndex() {
				// 		break
				// 	}
				// 	last = tree.GetStop().GetTokenIndex()
				// 	fmt.Println(last, " total: ", number)
				// 	if tree.GetStop().GetTokenType() == antlr.TokenEOF {
				// 		break
				// 	}
				// }
				// fmt.Println(trees[len(trees)-1].GetStop().GetTokenIndex())
				require.Equal(t, 0, lexerErrors.errors)
				require.Equal(t, 0, parserErrors.errors)
				total++
			}

			fmt.Println(total)
		})
	}
}
