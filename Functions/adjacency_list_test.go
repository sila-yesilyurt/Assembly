package Functions

import (
    "testing"
)

//Test MakeOverlapNetwork
type testpairAdjacencyList struct {
    Reads []string
    Match float64
    Mis float64
    Gap float64
    Threshold float64
    Expected map[string][]string
}

func TestMakeOverlapNetwork(t *testing.T) {
    /* TESTS */
    var reads0 []string = []string{
        "ATAAC",
        "AACTT",
        "ACTAA",
    }
    var reads1 []string = []string{
        "ATTGA",
        "TTGAC",
        "TGACA",
        "GACAG",
        "ACAGT",
    }
    var reads2 []string = []string{
        "ATTGGA",
        "TGGAGG",
        "TGGGAA",
        "CCGGATT",
        "CGATCCGA",
        "ATTAAA",
        "GGGGTCAG",
        "GGGTTTCG",
    }
    var reads3 []string = []string {
        "ATTGAGGGATTA",
        "GAGGGCTCAACT",
        "GGCTCAAGGTCC",
        "AAGGTCCGAGGT",
    }
    var reads4 []string = []string{
        "ATTTCCTAAGAACT",
        "ATCCGATCCGAGG",
        "GATTTCGGGGAA",
        "CCTTTTTGGGA",
        "TTTTTCGATTTC",
        "GGGTCCGGACCTTA",
        "GGTCCAAATCCCGA",
        "GGTCCCAGGGGTC",
    }

    //Expected Answers
    ans0 := map[string][]string{
        "AACTT": {"ACTAA"},
        "ACTAA": {"ATAAC"},
        "ATAAC": {"AACTT"},
    }

    ans1 := map[string][]string{
        "ATTGA": {"TTGAC", "TGACA", "GACAG"},
        "GACAG": {"ACAGT"},
        "TGACA": {"GACAG", "ACAGT"},
        "TTGAC": {"TGACA", "GACAG", "ACAGT"},
    }

    ans2 := map[string][]string{
        "ATTGGA": {"TGGAGG", "TGGGAA", "ATTAAA"},
        "CCGGATT": {"ATTGGA", "ATTAAA", "GGGTTTCG"},
        "CGATCCGA": {"ATTGGA", "CCGGATT"},
        "GGGTTTCG": {"CGATCCGA"},
        "TGGAGG": {"TGGGAA", "GGGGTCAG", "GGGTTTCG"},
        "GGGGTCAG": {"GGGTTTCG"},
    }

    ans3 := map[string][]string{
        "AAGGTCCGAGGT": {"GAGGGCTCAACT"},
        "ATTGAGGGATTA": {"GAGGGCTCAACT", "GGCTCAAGGTCC"},
        "GAGGGCTCAACT": {"GGCTCAAGGTCC"},
        "GGCTCAAGGTCC": {"AAGGTCCGAGGT"},
    }

    ans4 := map[string][]string{
        "ATCCGATCCGAGG": {"GGTCCCAGGGGTC"},
        "CCTTTTTGGGA": {"TTTTTCGATTTC"},
        "GGTCCAAATCCCGA": {"ATCCGATCCGAGG"},
        "GGTCCCAGGGGTC": {"GGGTCCGGACCTTA"},
        "TTTTTCGATTTC": {"ATTTCCTAAGAACT", "GATTTCGGGGAA"},
    }

    tests := []testpairAdjacencyList{
        {reads0, 1, 1, 1, 2, ans0},     //testcase 0
        {reads1, 1, 1, 1, 2, ans1},     //testcase 1
        {reads2, 1, 1, 1, 2, ans2},     //testcase 2
        {reads3, 1, 1, 1, 2, ans3},     //testcase 3
        {reads4, 1, 1, 1, 5, ans4},     //testcase 4
    }

    for i, test := range tests {
        result := MakeOverlapNetwork(test.Reads, test.Match, test.Mis, test.Gap, test.Threshold)
        if (!GraphEq(result, test.Expected)) {
            t.Error("MakeOverlapNetwork() failed testcase#", i, "\n",
                    "Expecting:", test.Expected, "Actual outcome:", result)
        }
    }
}
