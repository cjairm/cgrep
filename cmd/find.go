package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	ArgsString  string
	RelativeDir string
)

var findCmd = &cobra.Command{
	Use:   "find [string-to-find]",
	Short: "find file",
	Long:  "find file",
	Args:  argsRun,
	Run:   run,
}

func argsRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a string you are looking for")
	}
	ArgsString = strings.ToLower(strings.Join(args, ""))
	return nil
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	RelativeDir = wd
	visitAllFilesInDir(wd, func(pathToFile string) {
		err = readFile(pathToFile)
		if err != nil {
			log.Fatal(err)
			return
		}
	})
}

func visitAllFilesInDir(wd string, funcPathToFile func(pathToFile string)) {
	err := filepath.Walk(wd, func(path string, file fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		// fmt.Printf("File %q has type %q\n", path, ext)

		if !file.IsDir() && ext != "" && ext != ".zip" {
			funcPathToFile(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(wd string) error {
	file, err := os.Open(wd)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	conincidencesFound := []string{}
	lineNum := 1
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		foundLine := lookUpWord(line)
		if foundLine != "" {
			lineHighLighted := highlightLine(ArgsString, foundLine)
			numLinehighlighted := highlightNumLine(lineNum)
			conincidencesFound = append(
				conincidencesFound,
				fmt.Sprintf("%s: %s", numLinehighlighted, lineHighLighted),
			)
		}
		lineNum++
	}
	if len(conincidencesFound) > 0 {
		pathToShow := strings.Replace(wd, RelativeDir+"/", "", 1)
		fmt.Printf("\033[31m" + pathToShow + "\033[0m\n")
		for _, lineFound := range conincidencesFound {
			fmt.Printf(lineFound)
		}
		fmt.Printf("\n")
	}
	return nil
}

func lookUpWord(line string) string {
	lineLooking := strings.ToLower(line)
	if strings.Contains(lineLooking, ArgsString) {
		return line
	}
	return ""
}

func highlightLine(word, line string) string {
	startTag := "\033[35m"
	endTag := "\033[0m"

	highlightedStr := strings.Replace(line, word, startTag+word+endTag, -1)

	return highlightedStr
}

func highlightNumLine(lineNum int) string {
	startTag := "\033[32m"
	endTag := "\033[0m"

	return startTag + fmt.Sprintf("%d", lineNum) + endTag
}
