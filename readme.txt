Copy the Assembly folder into your "go/src" directory.

Then, write each of the missing functions in the Functions folder. If you need a subroutine, please feel free to place it in the same file.

To test a function, you should "cd" into the Functions directory and run

go test -run ^TestFunctionName$

where TestFunctionName is based on the specific function you are testing.

For example, after writing each appropriate function, you should call each of the following tests.

go test -run ^TestGenerateRandomGenome$
go test -run ^TestKmerComposition$
go test -run ^TestShuffle$
go test -run ^TestScoreOverlapAlignment$
go test -run ^TestOverlapScoringMatrix$
go test -run ^TestBinarizeMatrix$
go test -run ^TestMakeOverlapNetwork$
go test -run ^TestAverageOutDegree$
go test -run ^TestSimulateReadsClean$
go test -run ^TestMinimizer$
go test -run ^TestMapToMinimizer$
go test -run ^TestGetExtendedNeighbors$
go test -run ^TestGetTrimmedNeighbors$
go test -run ^TestTrimNetwork$


To check that you have passed *all* tests, call "go test".

We will periodically navigate up ("cd ..") into the parent directory and run code calling our functions using main.go.