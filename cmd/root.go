package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

// rootcmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ggg",
	Short: "CLI tool for visualizing graph structure",
	Long: `hoge
This application is a tool to visualize graph structure.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(c *cobra.Command, args []string) {
		// フラグ名で値を取得する
		indexed, err := c.PersistentFlags().GetBool("indexed")
		if err != nil {
			fmt.Println(err)
			return
		}
		directed, err := c.PersistentFlags().GetBool("directed")
		if err != nil {
			fmt.Println(err)
			return
		}
		weighted, err := c.PersistentFlags().GetBool("weighted")
		if err != nil {
			fmt.Println(err)
			return
		}
		format, err := c.PersistentFlags().GetBool("format")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("indexed: ", indexed)
		fmt.Println("directed: ", directed)
		fmt.Println("weighted: ", weighted)
		fmt.Println("format: ", format)

		// validation をかけながら、入力を読む
		var url string
		if format {
			url, err = readGraph(indexed, directed, weighted)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// TODO 隣接行列に対応
			fmt.Println("隣接行列にはまだ対応していません")
		}

		if err := openbrowser(url); err != nil {
			log.Fatal(err)
		}
	},
}

func readGraph(indexed, directed, weighted bool) (string, error) {

	tf := map[bool]string{true: "true", false: "false"}

	hostUrl := "https://hello-world-494ec.firebaseapp.com/index.html"
	var queryUrl = bytes.NewBuffer(make([]byte, 0, 100))
	queryUrl.WriteString("indexed=" + tf[indexed] + "&directed=" + tf[directed] + "&weighted=" + tf[weighted])
	queryUrl.WriteString("&format=true&data=")

	var n, m int
	fmt.Scan(&n, &m)
	if n <= 0 {
		return "", errors.New("n must be positive integer")
	}

	if m < 0 {
		return "", errors.New("m must be non negative integer")
	}

	queryUrl.WriteString(strconv.Itoa(n) + "-" + strconv.Itoa(m))

	for i := 0; i < m; i++ {
		var a, b, c int

		if weighted {
			fmt.Scan(&a, &b, &c)
			queryUrl.WriteString("," + strconv.Itoa(a) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(c))
		} else {
			fmt.Scan(&a, &b)
			queryUrl.WriteString("," + strconv.Itoa(a) + "-" + strconv.Itoa(b))
		}

		if indexed {
			// a, b must be [1, n]
			for _, x := range []int{a, b} {
				if !(1 <= x && x <= n) {
					return "", errors.New(fmt.Sprintf("node must be in the range [%d %d]\n", 1, n))
				}
			}
		} else {
			// a, b must be [0, n - 1]
			for _, x := range []int{a, b} {
				if !(0 <= x && x <= n-1) {
					return "", errors.New(fmt.Sprintf("node must be in the range [%d %d]\n", 0, n-1))
				}
			}
		}
	}

	url := hostUrl + "?" + queryUrl.String()

	return url, nil
}

func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	rootCmd.PersistentFlags().BoolP("indexed", "i", true, "graph is 1-indexed")
	rootCmd.PersistentFlags().BoolP("directed", "d", false, "graph is directed")
	rootCmd.PersistentFlags().BoolP("weighted", "w", false, "graph is weighted")
	rootCmd.PersistentFlags().BoolP("format", "f", true, "graph format is normal")
}
