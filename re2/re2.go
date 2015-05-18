// re2.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// need more testing examples +++

package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type patternType struct {
	pattern string
	count   int64
}

var (
	collectionPatterns = []patternType{}
	g_verboseFlag      = false
)

func init() {
	// note that filenames are converted to lowercase before matching takes place
	// paterns should have longest match possiblity listed first below  (tar.z before .z)
	///////////////////////////////////////////////////////////////////////////
	// N E E E D E D

	// fmt.Printf("5\n")
	// file.log
	//	pat := ".*\\Q.\\Elog$" // just a placeholder
	//	collectionPatterns = append(collectionPatterns, patternType{ pat,0})

	/////////////////////////////////////////////////////////////////////////////
	// W O R K - I N - P R O G R E S S
	//  file.z -  recognized by dispatcher
	//	pat = ".*\\Q.\\E\\Qz\\E$" // Unix "pack" or compress
	//	collectionPatterns = append(collectionPatterns, patternType{ pat,0})

	//  file.tar.z - recognized by dispatcher
	pat := ".*\\Q.tar.z\\E$" // Unix "pack" or compress
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.jar - recognized by dispatcher
	pat = ".*\\Q.\\Ejar$" // java jar is somewhat similar to zip
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})
	////////////////////////////////////////////////////////////////////////////
	// D O N E

	// file.shar - needs more testing
	pat = ".*\\Q.\\Eshar$" // Shell Archive not compressed
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.shar.Z(z) - needs more testing
	pat = ".*\\Q.\\Eshar\\Q.\\Ez$" // Shell Archive not compressed
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.zip - testOK
	pat = ".*\\Q.zip\\E$" // PK-Zip usually compressed
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.tar - testOK
	pat = ".*\\Q.tar\\E$" // uncompressed tar
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.tgz - testOK
	pat = ".*\\Q.\\Etgz$" // tar compressed with gz
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.taz -
	pat = ".*\\Q.\\Etaz$" // tar compressed with gz
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.tar.gz - testOK
	pat = ".*\\Q.tar.gz\\E$" // tar compressed with gz
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.tbz - testOK
	pat = ".*\\Q.\\Etbz$" // tar compressed with bz2
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.tar.bz
	pat = ".*\\Q.tar.bz\\E$" // tar compressed with bz2
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.tar.bz2
	pat = ".*\\Q.tar.bz2\\E$" // tar compressed with bz2
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.tbz2
	pat = ".*\\Q.\\Etbz2$" // tar compressed with bz2
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})

	// file.bzip2
	pat = ".*\\Q.\\Ebzip2$" // tar compressed with bz2
	collectionPatterns = append(collectionPatterns, patternType{pat, 0})
}

func FileNameIsCollection(fname string) (bool, string) {
	//fmt.Printf("testing %s\n",fname)
	nameBytes := []byte(strings.ToLower(fname))
	maxNdx := len(collectionPatterns)
	for i := 0; i < maxNdx; i++ {
		pat := collectionPatterns[i]
		//		fun := dispatcher[i].Function
		isMatch, err := regexp.Match(pat.pattern, nameBytes)
		if err != nil {
			fmt.Printf("!Err-> ?pattern error in re2 %v\n", pat)
			log.Panicf(fmt.Sprintf("re2 pattern error %v", pat))
		}
		if isMatch {
			collectionPatterns[i].count++
			return true, pat.pattern
		} // only need the first match
	}
	// zipLog(fmt.Sprintf("!Err-> %s isn't a collection\n", fname))
	return false, "<no-match>"
}

func fatalErr(err error) {
	log.Panicf(fmt.Sprintf("%v", err))
}

func test2() {
	type testBlk struct {
		name  string
		iscol bool
	}

	var testBlocks = []testBlk{
		{"abc.jar", true},
		{"abc.jarHead", false},
		{"abc.JAR", true},
		{"abc.Jared", false},
		{"abc.JaR", true},
		{"abc.gz", false},
		{"abc.tgz", true},
		{"abc.gzip", false},
		{"abc.g.zip", true},
		{"abc.tar.gz", true},
		{"abc.shar", true},
		{"abc.Z", false},
		{"abc.z", false},
		{"abc.tbz", true},
		{"abc.tar.bz", true},
		{"abc.bzip2", true},
		{"abc.tbz2", true},
		{"abc.tar.bz2", true},
		{"abc.cpio", false},
		{"abc.cpio.gz", false},
		{"coverage-fail.tar", false},
	}
	fmt.Printf("is named file a collection? compare human & dispatcher\n")
	for _, blk := range testBlocks {
		name := blk.name
		iscol := blk.iscol
		iscollection, pat := FileNameIsCollection(name)
		if iscol != iscollection {
			fmt.Printf("%s %v  disagrees with %s <----- \n", name, iscol, pat)
		} else {
			fmt.Printf("%s %v  agrees with %s\n", name, iscol, pat)
		}
	}
}

func main() {
	//	test1()
	test2()
}
