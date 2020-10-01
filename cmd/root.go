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

var rootCmd = &cobra.Command{
	Use:   "ggg",
	Short: "CLI tool for visualizing graph",
	Long: `
This application is a tool to visualize graph.
You can select

・1-indexed / 0-indexed
・directed / undirected
・weighted / unweighted

by optional flags.

`,
	Run: func(c *cobra.Command, args []string) {
		printLogo()
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
		/* matrix, err := c.PersistentFlags().GetBool("matrix")
		if err != nil {
			fmt.Println(err)
			return
		}
		*/

		fmt.Println("")
		fmt.Println("Option:")
		fmt.Println("  indexed: ", indexed)
		fmt.Println("  directed: ", directed)
		fmt.Println("  weighted: ", weighted)
		// fmt.Println("  matrix: ", matrix)
		printGraphImage(indexed, directed, weighted)
		printGraphFormat(indexed, directed, weighted, false)

		fmt.Println("")
		fmt.Println("please input your graph below...")

		// validation をかけながら、入力を読む
		var url string
		if false {
			// TODO 隣接行列に対応
			log.Fatal(errors.New("隣接行列にはまだ対応していません"))
		} else {
			url, err = readGraph(indexed, directed, weighted)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err := openbrowser(url); err != nil {
			log.Fatal(err)
		}

		fmt.Println("visualized correctly ✨✨")
	},
}

func printLogo() {
	fmt.Println("go GRAPH × GRAPH")
	fmt.Println("version 1.0.0")
}

func printGraphFormat(indexed, directed, weighted, matrix bool) {
	if matrix {
		return
	}

	n := 3
	m := 2

	u1 := 1
	v1 := 2
	w1 := 5

	u2 := 2
	v2 := 3
	w2 := 7

	if !indexed {
		u1--
		u2--
		v1--
		v2--
	}

	fmt.Println("")
	fmt.Println("Format:")
	fmt.Println("  ", n, m, "     (the number of nodes, the number of edges )")
	if weighted {
		fmt.Println("  ", u1, v1, w1, "   (edge informations.. )")
		fmt.Println("  ", u2, v2, w2)
	} else {
		fmt.Println("  ", u1, v1, "     (edge informations.. )")
		fmt.Println("  ", u2, v2)
	}
}

func printGraphImage(indexed, directed, weighted bool) {

	/*
	   ①  ----- ② ----- ③
	*/

	var nodeLeft, nodeMiddle, nodeRight string
	if indexed {
		nodeLeft = "①"
		nodeMiddle = "②"
		nodeRight = "③"
	} else {
		nodeLeft = "⓪"
		nodeMiddle = "①"
		nodeRight = "②"
	}

	var weightLeft, weightRight string
	if weighted {
		weightLeft = "5"
		weightRight = "7"
	} else {
		weightLeft = "-"
		weightRight = "-"
	}

	var arrow string
	if directed {
		arrow = ">"
	} else {
		arrow = "-"
	}

	fmt.Println("")
	fmt.Println("Image:")
	fmt.Println("  ---------------------")
	fmt.Println("  |                   |")
	fmt.Println("  |", nodeLeft, "--"+weightLeft+"-"+arrow, nodeMiddle, "--"+weightRight+"-"+arrow, nodeRight, "|")
	fmt.Println("  |                   |")
	fmt.Println("  ---------------------")
}

func readGraph(indexed, directed, weighted bool) (string, error) {

	tf := map[bool]string{true: "true", false: "false"}

	hostUrl := "https://hello-world-494ec.firebaseapp.com/index.html"
	var queryUrl = bytes.NewBuffer(make([]byte, 0, 100))
	queryUrl.WriteString("indexed=" + tf[indexed] + "&directed=" + tf[directed] + "&weighted=" + tf[weighted])
	queryUrl.WriteString("&format=true&data=")

	var n, m int
	fmt.Print(">>> ")
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("indexed", "i", true, "graph is 1-indexed（if you want to visualize 0-indexed graph, you should add `-i=false` option）")
	rootCmd.PersistentFlags().BoolP("directed", "d", false, "graph is directed")
	rootCmd.PersistentFlags().BoolP("weighted", "w", false, "graph is weighted")
	// rootCmd.PersistentFlags().BoolP("matrix", "m", false, "graph format is matrix")
}
