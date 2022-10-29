package Functions

import "testing"

type testpairBinarizeMatrix struct {
    Matrix [][]float64
    Threshold float64
    Expected [][]int
}

/* TESTS */
func TestBinarizeMatrix(t *testing.T) {
    //Only include # that's above threshold.
    test1 := [][]float64{
        {0.0, 1.7, 2.3},
        {0.1, 0.0, 1.1},
        {0.2, 2.4, 0.0},
    }
    ans1 := [][]int{
        {0, 1, 1},
        {0, 0, 0},
        {0, 1, 0},
    }

    //Only keep the larger of between M[i][j] & M[j][i]
    test2 := [][]float64{
        {0.0, 4.2, 3.0, 5.5, 2.1},
        {5.3, 0.0, 0.5, 0.7, 5.5},
        {1.3, 2.2, 0.0, 3.2, 4.0},
        {1.1, 1.4, 2.5, 0.0, 3.3},
        {3.1, 2.4, 5.3, 4.3, 0.0},
    }
    ans2 := [][]int{
        {0, 0, 1, 1, 0},
        {1, 0, 0, 0, 1},
        {0, 0, 0, 1, 0},
        {0, 0, 0, 0, 0},
        {1, 0, 1, 1, 0},
    }

    //Only keep the larger of between M[i][j] & M[j][i]
    test3 := [][]float64{
        {0.0, 2.4, 0.3},
        {2.3, 0.0, 4.2},
        {0.5, 1.0, 0.0},
    }
    ans3 := [][]int{
        {0, 1, 0},
        {0, 0, 1},
        {1, 0, 0},
    }

    //Edge case: when m[i][j] == m[j][i], keep either one.
    test4 := [][]float64{
        {0.0, 5.0, 0.5},
        {5.0, 0.0, 2.4},
        {0.5, 2.4, 0.0},
    }
    ans4 := [][]int{
        {0, 0, 0},
        {1, 0, 0},
        {0, 1, 0},
    }
    ans4_alt := [][]int{
        {0, 1, 0},
        {0, 0, 1},
        {0, 0, 0},
    }

    tests := []testpairBinarizeMatrix{
        {test1, 1.5, ans1},     //testcase1
        {test2, 3.0, ans2},     //testcase2
        {test3, 0.5, ans3},     //testcase3
    }

    for i, test := range tests {
        result := BinarizeMatrix(test.Matrix, test.Threshold)
        if (!MtxEq(result, test.Expected)) {
            t.Error("BinarizeMatrix failed testcase#", i+1, "\n",
                    "Expecting:", test.Expected, "Actual outcome:", result)
        }
    }

    //edge case:
    result := BinarizeMatrix(test4, 0.6)
    if (!MtxEq(result, ans4) && !MtxEq(result, ans4_alt)) {
        t.Error("BinarizeMatrix failed testcase4", "\n",
                "Expecting:", ans4, "OR", ans4_alt, "Actual outcome", result)
    }
}
