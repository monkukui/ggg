package cmd

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/monkukui/ggg/lib/graph"
	"github.com/stretchr/testify/assert"
)

func TestReadGraph_invalidInput(t *testing.T) {

	tests := []struct {
		name     string
		data     string
		indexed  bool
		directed bool
		weighted bool
	}{
		{
			name:     "辺で指定された頂点が n を超えている（1-indexed）",
			data:     "3 1\n1 4\n",
			indexed:  true,
			directed: false,
			weighted: false,
		},
		{
			name:     "辺で指定された頂点が n を含む（0-indexed）",
			data:     "3 1\n1 3\n",
			indexed:  false,
			directed: false,
			weighted: false,
		},
		{
			name:     "辺で指定された頂点が 0 を含む（1-indexed）",
			data:     "3 1\n0 3\n",
			indexed:  true,
			directed: false,
			weighted: false,
		},
		{
			name:     "辺で指定された頂点が負の数を含む（0-indexed）",
			data:     "3 1\n-4 3\n",
			indexed:  false,
			directed: false,
			weighted: false,
		},
		{
			name:     "n が負の数（0-indexed）",
			data:     "-4 1\n-4 3\n",
			indexed:  false,
			directed: false,
			weighted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := bytes.NewBufferString(tt.data + "\n")
			output := new(bytes.Buffer)
			graph := &graph.Graph{
				Stdin:  input,
				Stdout: output,
				Stderr: output,
			}

			scanner := bufio.NewScanner(graph.Stdin)
			_, err := readGraph(tt.indexed, tt.directed, tt.weighted, scanner)
			assert.Error(t, err)
		})
	}
}

// indexed, directed, weighted の t/f の全パターンをテストする -> 8 通り
func TestReadGraph_validInput(t *testing.T) {

	hostURL := "https://hello-world-494ec.firebaseapp.com/index.html"

	tests := []struct {
		name          string
		data          string
		indexed       bool
		directed      bool
		weighted      bool
		expectedQuery string
	}{
		{
			name:          "0-indexed, undirected, unweighted で妥当な入力",
			data:          "3 2\n0 1\n1 2\n",
			indexed:       false,
			directed:      false,
			weighted:      false,
			expectedQuery: "indexed=false&directed=false&weighted=false&format=true&data=3-2,0-1,1-2",
		},
		{
			name:          "1-indexed, undirected, unweighted で妥当な入力",
			data:          "3 2\n1 2\n2 3\n",
			indexed:       true,
			directed:      false,
			weighted:      false,
			expectedQuery: "indexed=true&directed=false&weighted=false&format=true&data=3-2,1-2,2-3",
		},
		{
			name:          "0-indexed, directed, unweighted で妥当な入力",
			data:          "3 2\n0 1\n1 2\n",
			indexed:       false,
			directed:      true,
			weighted:      false,
			expectedQuery: "indexed=false&directed=true&weighted=false&format=true&data=3-2,0-1,1-2",
		},
		{
			name:          "1-indexed, directed, unweighted で妥当な入力",
			data:          "3 2\n1 2\n2 3\n",
			indexed:       true,
			directed:      true,
			weighted:      false,
			expectedQuery: "indexed=true&directed=true&weighted=false&format=true&data=3-2,1-2,2-3",
		},
		{
			name:          "0-indexed, undirected, weighted で妥当な入力",
			data:          "3 2\n0 1 5\n1 2 7\n",
			indexed:       false,
			directed:      false,
			weighted:      true,
			expectedQuery: "indexed=false&directed=false&weighted=true&format=true&data=3-2,0-1-5,1-2-7",
		},
		{
			name:          "1-indexed, undirected, weighted で妥当な入力",
			data:          "3 2\n1 2 5\n2 3 7\n",
			indexed:       true,
			directed:      false,
			weighted:      true,
			expectedQuery: "indexed=true&directed=false&weighted=true&format=true&data=3-2,1-2-5,2-3-7",
		},
		{
			name:          "0-indexed, directed, weighted で妥当な入力",
			data:          "3 2\n0 1 5\n1 2 7\n",
			indexed:       false,
			directed:      true,
			weighted:      true,
			expectedQuery: "indexed=false&directed=true&weighted=true&format=true&data=3-2,0-1-5,1-2-7",
		},
		{
			name:          "1-indexed, directed, weighted で妥当な入力",
			data:          "3 2\n1 2 5\n2 3 7\n",
			indexed:       true,
			directed:      true,
			weighted:      true,
			expectedQuery: "indexed=true&directed=true&weighted=true&format=true&data=3-2,1-2-5,2-3-7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := bytes.NewBufferString(tt.data + "\n")
			output := new(bytes.Buffer)
			graph := &graph.Graph{
				Stdin:  input,
				Stdout: output,
				Stderr: output,
			}

			scanner := bufio.NewScanner(graph.Stdin)
			url, err := readGraph(tt.indexed, tt.directed, tt.weighted, scanner)
			assert.NoError(t, err)

			assert.Equal(t, hostURL+"?"+tt.expectedQuery, url)
		})
	}
}
